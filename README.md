# Lock Gallery
Little utility to randomize lock screen for i3

I like to have different lock screens, so I wrote a little util that I run everytime I lock my screen. I use i3, an to lock the screen I do:

    # extract of my ~/.i3/config
    bindsym $mod+l exec /home/ivan/dotfiles/bin/lockme.sh
    exec xautolock -time 5 -locker /home/ivan/dotfiles/bin/lockme.sh &
  
You can find [lockme.sh](https://github.com/ipedrazas/dotfiles/blob/master/bin/lockme.sh) in my dotfiles.

As you can see in the `lockme.sh` file, to modify your lock screen wallpaper you just have to call this tool and pass 2 parameters: your target file, and your directory with images

    ./lock-gallery -d ~/backgrounds/ ~/.i3/screen_locked.png
