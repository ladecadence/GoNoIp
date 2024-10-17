package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ladecadence/GoNoIp/pkg/config"
	"github.com/ladecadence/GoNoIp/pkg/update"
)

func main() {
	// check args
	configFile := flag.String("c", "", "Configuration file name")
	testMode := flag.Bool("t", false, "Test mode")

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println(os.Args[0] + " -c <config file> [-t]")
		fmt.Println("  -t : test mode (don't send updates)")
	}
	flag.Parse()
	if *configFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// check config file permissions
	fileInfo, err := os.Stat(*configFile)
	if err != nil {
		fmt.Printf("Error accessing config file: %v", err)
		os.Exit(1)
	}
	if fileInfo.Mode()&(1<<2) != 0 {
		fmt.Printf("Config file is world readable, change permissions.\n")
		os.Exit(1)
	}

	config, err := config.GetConfig(*configFile)
	if err != nil {
		fmt.Printf("Config file error: %s\n", err.Error())
		os.Exit(1)
	}

	if *testMode {
		fmt.Println("test file ok")
	} else {
		for _, host := range config.Hosts {
			ok := update.Update(host)
			fmt.Printf("Host: %s : %s\n", host.Hostname, ok)
		}
	}

}
