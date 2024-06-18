package inputandclick

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func KeyCombinations() {
	launcherUrl := launcher.New().Headless(false).MustLaunch()
	browser := rod.New().ControlURL(launcherUrl).MustConnect()
	page := browser.MustPage("https://pragmanotes.co.uk/")

	page.KeyActions().Press(input.ControlRight).Type('A').MustDo()
	time.Sleep(time.Second * 10)
}
