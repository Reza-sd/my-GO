#ls -l | sed -n '1~2p'
ls -l | awk NR%2==0