package inputandclick

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func ScrollPage() {
	launcherUrl := launcher.New().Headless(false).MustLaunch()
	browser := rod.New().ControlURL(launcherUrl).MustConnect()
	page := browser.MustPage("https://www.google.com")

	//APjFqb is google input's css selector, write hello world and send enter key
	page.MustElement("#APjFqb").MustInput("Hello World!").MustType(input.Enter)

	//scroll 2000 pixels for 10 times
	for i := 0; i < 300; i++ {
		err := page.Mouse.Scroll(0, 2000, 10)
		if err != nil {
			log.Fatal(err)
		}
	}

}
