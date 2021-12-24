package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error reading $HOME: %s", err)
		os.Exit(1)
	}

	// grab command line args
	setConfigPath := flag.String("set-path", "", "Path for where timelog file is stored")
	flag.Parse()

	// open cofig file
	f, err := os.OpenFile(homeDir+"/.logtime", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening ~/.logtime file: %s", err)
		os.Exit(1)
	}
	defer f.Close()

	timeLogPath, err := ioutil.ReadFile(homeDir + "/.logtime")
	if err != nil {
		fmt.Printf("Error reading ~/.logtime file: %s", err)
		os.Exit(1)
	}
	// write default path if not set
	if len(timeLogPath) == 0 {
		timeLogPath = []byte(homeDir + "/timelog.txt")
		_, err = f.WriteString(homeDir + "/timelog.txt")
		if err != nil {
			fmt.Printf("Error writing to config file: %s", err)
		}
	}
	// check if we are setting timelog file locationgo env
	var toLog []string
	if *setConfigPath != "" {
		toLog = os.Args[2:]
		// write new path to ~/.logtime
		_, err = f.WriteString(*setConfigPath)
		if err != nil {
			fmt.Printf("Error writing to ~/.logtime file: %s", err)
			os.Exit(1)
		}
		timeLogPath = []byte(*setConfigPath + "/timelog.txt")
	} else {
		toLog = os.Args[1:]
	}

	// write to timelog
	timeLogFile, err := os.OpenFile(string(timeLogPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening timelog file: %s", err)
		os.Exit(1)
	}
	defer timeLogFile.Close()
	logString := ""
	for _, word := range toLog {
		logString = logString + word + " "
	}
	timeLogFile.WriteString(time.Now().Format(time.ANSIC) + ": " + logString + "\n")
	os.Exit(0)
}
