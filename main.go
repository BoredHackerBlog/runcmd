package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main(){
	var textfile string
	var url string
	var logfile string

	flag.StringVar(&textfile, "file", "", "Text file")
	flag.StringVar(&url, "url", "", "URL")
	flag.StringVar(&logfile, "log", "", "Output log file - required")

	flag.Parse()

	if logfile == "" {
		fmt.Println("output log file name is required")
		os.Exit(1)
	}

	var commands string

	if textfile != "" {
		content, err := ioutil.ReadFile(textfile)
		if err != nil {
			fmt.Println(err)
			return
		}
		commands = string(content)
	} else if url != "" {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		commands = string(content)
	} else {
		fmt.Println("file or url are required")
		return
	}

	f, err := os.Create(logfile) //https://golangbot.com/write-files/
    if err != nil {
        fmt.Println(err)
        return
    }

	commands = strings.Replace(commands, "\r\n", "\n", -1)
	commands_split := strings.Split(commands, "\n")
	
	var clean_commands []string

	for _, value := range commands_split {
		if value != ""{
			clean_commands = append(clean_commands, value)
		}
	}

	for _, command := range clean_commands {
		args := strings.Fields(command)
		cmd := exec.Command(args[0], args[1:]...)
		
		command = command + "\n"
		_, err = f.WriteString(command)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}

		_, err = f.WriteString(string(output))
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}

    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

}
