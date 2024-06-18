package inputandclick

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Input(){
	launcherUrl := launcher.New().Headless(false).MustLaunch()
	browser := rod.New().ControlURL(launcherUrl).MustConnect()

	page := browser.MustPage("https://pragmanotes.co.uk/login.html")
	
	page.MustWaitStable()

	//MustElement gets element with css selector and return element, mustInput can fill element with argument
	page.MustElement("#mail").MustInput("example@gmail.com")
	page.MustElement("#password").MustInput("123123123")

	time.Sleep(time.Second * 10)

}