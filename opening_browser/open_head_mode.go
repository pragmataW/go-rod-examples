package openingbrowser

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func OpenHeadModBrowser() {
	//Create a launcher url with launcher package
	l := launcher.New().Headless(false).MustLaunch()

	//create browser with launcher url
	browser := rod.New().ControlURL(l).MustConnect()

	//get page
	page := browser.MustPage("https://www.google.com")

	//wait'til page load
	time.Sleep(time.Second * 2)

	//print page url
	fmt.Println(page.MustInfo().URL)
}
