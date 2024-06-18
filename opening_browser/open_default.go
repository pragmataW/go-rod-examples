package openingbrowser

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

//Opening Default Browser
func OpenDefaultBrowser() {
	//create new browser
	browser := rod.New().MustConnect()

	//get page from browser
	page := browser.MustPage("https://www.google.com")
	
	//wait until load page
	time.Sleep(time.Second * 2)

	//get page URL
	fmt.Println(page.MustInfo().URL)
}
