# runcmd
just golang binary that runs commands from url or local file and logs output


# usage

```
Usage of ./runcmd:
  -file string
        Text file
  -url string
        URL
```

# example

```
ubuntu:~/runcmd$ cat cmd.txt 
ls
date

ubuntu:~/runcmd$ ./main -file cmd.txt 
2022-04-08 18:45:18.349384602 -0700 PDT m=+0.000543554 , ls
2022-04-08 18:45:18.351011802 -0700 PDT m=+0.002170771 , cmd.txt
main
main.go
README.md

2022-04-08 18:45:18.351107496 -0700 PDT m=+0.002266450 , date
2022-04-08 18:45:18.352342846 -0700 PDT m=+0.003501829 , Fri 08 Apr 2022 06:45:18 PM PDT
```
