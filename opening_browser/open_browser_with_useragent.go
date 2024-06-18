package openingbrowser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func OpenBrowserWithUserAgent() {
	//create urlLauncher with head mode
	urlLauncher := launcher.New().
		Headless(false).
		MustLaunch()

	//get browser instance
	browser := rod.New().ControlURL(urlLauncher).MustConnect()

	//override page userAgent
	page := browser.MustPage("https://www.google.com").MustSetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	})

	//wait page until all animations loaded
	page.MustWaitStable()

	//click element, take input and press enter
	page.MustElement("#APjFqb").MustInput("what is my user agent").MustType(input.Enter)

	page.MustWaitStable()
}
