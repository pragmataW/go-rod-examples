package elementprocessing

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
)

func GetImgFromElement() {
	launcherUrl := launcher.New().Headless(false).MustLaunch()
	page := rod.New().ControlURL(launcherUrl).MustConnect().MustPage("https://www.wikipedia.org/")

	page.MustElement("#searchInput").MustInput("earth")
	page.MustElement("#search-form > fieldset > button").MustClick()

	el := page.MustElement("#mw-content-text > div.mw-content-ltr.mw-parser-output > table.infobox > tbody > tr:nth-child(1) > td > span > a > img")

	_ = utils.OutputFile("b.png", el.MustResource())
}
