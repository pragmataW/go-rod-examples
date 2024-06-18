package elementprocessing

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func GetTextFromElement(){
	launcherUrl := launcher.New().Headless(false).MustLaunch()
	browser := rod.New().ControlURL(launcherUrl).MustConnect()
	element := browser.MustPage("https://pragmanotes.co.uk/").MustElement("div.feature:nth-child(2)")
	fmt.Println(element.MustText())
}