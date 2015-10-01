package main

import(
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "strings"
    "math/rand"
    "time"
    "image"
    "log"
    "image/jpeg"
    "image/png"
)

var target string

func getImages(dir string) ([]string){
    var images []string
    files, _ := ioutil.ReadDir(dir)
    for _, f := range files {
            if strings.HasSuffix(f.Name(), ".png")  || strings.HasSuffix(f.Name(), ".jpg") {
                images = append(images, dir + f.Name())
            }
    }
    return images
}

func ImageRead(ImageFile string) (image image.Image) {
     file, err := os.Open(ImageFile)
     if err != nil {
      log.Fatal(err)
     }
     img, err := jpeg.Decode(file)
     if err != nil {
      log.Fatal(err)
     }
     file.Close()

     return img
}

func Formatpng(img image.Image) {
 out, err := os.Create(target)
     if err != nil {
      fmt.Println(err)
      os.Exit(1)
     }
     err = png.Encode(out, img)
     if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
}

func CopyPng(dst, src string) error {
    in, err := os.Open(src)
    if err != nil { return err }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil { return err }
    defer out.Close()
    _, err = io.Copy(out, in)
    cerr := out.Close()
    if err != nil { return err }
    return cerr
}

func ProcessImage(img string){
    if strings.HasSuffix(img, ".png") {
            CopyPng(target, img)
    }else{
        Formatpng(ImageRead(img))
    }
}

func main() {

    helpMsg := "Use lock-gallery -d [image directory] -t [target directory]"

    directory := os.Args[1:]

    if os.Args[3] == "-t" {
        target = os.Args[4]
       if directory[0] == "-d" {
                images := getImages(directory[1])
                numImages := len(images)
                s1 := rand.NewSource(time.Now().UnixNano())
                r1 := rand.New(s1)
                pos := r1.Intn(numImages - 1)

                ProcessImage(images[pos])

        }else{
            fmt.Println(helpMsg)
        }
    }else{
             fmt.Println(helpMsg)
    }

}
