package elementprocessing

import (
	"fmt"

	"github.com/go-rod/rod"
)

func MostUsedElementSelectors() {
	page := rod.New().MustConnect().MustPage("https://pragmanotes.co.uk/")

	//x path
	element1 := page.MustElementX("/html/body/main/section[2]/div[1]")

	//css selector
	element2 := page.MustElement("div.feature:nth-child(2)")

	//with text first argument for tag, second argument for text
	element3 := page.MustElementR("a", "Kayıt Ol")

	//And theirs s posix for get all elements. like page.MustElements(a) -> gets all a tags in a list

	fmt.Println(element1.MustText())
	fmt.Println(element2.MustText())
	fmt.Println(element3.MustText())
}
