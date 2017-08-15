package main

import (
	"cowper/bibleweb"
	"fmt"
)

func main() {
	fmt.Println("HOHO")

	bible, err := bibleweb.NewAPI()
	if err != nil {
		println(err)
	}
	version := bibleweb.Version("eng-ESV")
	psalmChapterNum := 1

	psalm, err := bible.GetPsalm(version, psalmChapterNum)
	if err != nil {
		println(err)
	}
	text, err := psalm.GetNakedText()
	if err != nil {
		println(err)
	}
	println("foo")
	println(text)
}
