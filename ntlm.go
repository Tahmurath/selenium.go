package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main2() {

	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("service: %v\n", service)
	//defer service.Stop()
	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		//"--headless", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")

	if err != nil {
		log.Fatal("Error:", err)
	}

	// visit the target page
	//driver.get(“http://username:password@url”);
	err = driver.Get("http://migraven%5Cadministrator%3AMIG%40rasht.2024@win-mm8glctgeq2.migraven.ir/")
	if err != nil {
		log.Fatal("Error:", err)
	}
}
