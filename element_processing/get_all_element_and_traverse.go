package elementprocessing

import (
	"fmt"
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func GetAllElementsAndTraverse() {
	launcherUrl := launcher.New().Headless(false).MustLaunch()
	browser := rod.New().ControlURL(launcherUrl).MustConnect()
	page := browser.MustPage("https://www.google.com")

	//APjFqb is google input's css selector, write hello world and send enter key
	page.MustElement("#APjFqb").MustInput("Hello World!").MustType(input.Enter)

	//scroll 2000 pixels for 10 times
	for i := 0; i < 300; i++ {
		// first argument for scroll x axis, second for y axis, last argument for the number of steps
		err := page.Mouse.Scroll(0, 2000, 100)
		if err != nil {
			log.Fatal(err)
		}
	}

	//get all a tags with MustElements not MustElement
	links := page.MustElements("a")

	//traverse with range
	for _, link := range links {
		//get href attribute
		href, err := link.Attribute("href")
		if err != nil {
			log.Fatal(err)
		}

		//if href attribute != nil => print it
		if href != nil {
			fmt.Println(*href)
		}
	}
}
