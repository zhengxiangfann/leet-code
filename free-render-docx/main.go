package main

import (
	docx "github.com/lukasjarosch/go-docx"
)

func main() {
	replaceMap := docx.PlaceholderMap{
		"Fund_cnname": "测试测试测试",
	}

	// read and parse the template docx
	doc, err := docx.Open("glv_template1.docx")
	if err != nil {
		panic(err)
	}

	// replace the keys with values from replaceMap
	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		panic(err)
	}

	// write out a new file
	err = doc.WriteToFile("replaced.docx")
	if err != nil {
		panic(err)
	}
}
