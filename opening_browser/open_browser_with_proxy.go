package openingbrowser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func OpenBrowserWithProxy() {
	launcherUrl := launcher.New().Proxy("38.154.227.167:5868").Headless(false).MustLaunch()

	browser := rod.New().ControlURL(launcherUrl).MustConnect()

	page := browser.MustPage("https://www.google.com")

	page.MustWaitStable()

	page.MustElement("#APjFqb").MustInput("what is my ip address").MustType(input.Enter)

	page.MustWaitStable()
}
