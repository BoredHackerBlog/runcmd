package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"os/exec"
	//"time"
)

func main(){
	var textfile string
	var url string

	flag.StringVar(&textfile, "file", "", "Text file")
	flag.StringVar(&url, "url", "", "URL")

	flag.Parse()

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
		
		//fmt.Println(time.Now(), ",", command)

		//output, err := cmd.CombinedOutput()
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println(time.Now(), ",", string(output))
		fmt.Println(string(output))
	}

}
