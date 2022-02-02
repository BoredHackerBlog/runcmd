# runcmd
just golang binary that runs commands from url or local file and logs output


# usage

```
Usage of ./runcmd:
  -file string
        Text file
  -log string
        Output log file - required
  -url string
        URL
```

# example

```
~/runcmd$ cat commands.txt 
ls
date

~/runcmd$ ./runcmd -log output.txt -file commands.txt

~/runcmd$ cat output.txt
ls
commands.txt
output.txt
runcmd
date
Tue 01 Feb 2022 10:25:50 PM EST
```
