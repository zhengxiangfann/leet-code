package main

import (
	"fmt"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"log"
	"strings"
)

const apiKEY = `347a2c74b262eb4b5ee8e3b50a28283d03defd356bd1745116f74c9ea3cef9f5`

func init() {
	// The customer name needs to match the entry that is embedded in the signed key.
	//customerName := `test-95271`

	// Good to load the license key in `init`. Needs to be done prior to using the library, otherwise operations
	// will result in an error.

	err := license.SetMeteredKey(apiKEY)

	if err != nil {
		panic(err)
	}
}

func main() {
	fileSrc := "glv_template1.docx"
	fileDest := "output.docx"
	newContext := map[string]string{
		"Fund_cnname":  "Value1",
		"letters_date": "2023-06-09",
	}
	// 打开模板文件
	doc, err := document.Open(fileSrc)
	if err != nil {
		log.Fatalf("Failed to open template file: %s", err)
	}
	// 替换占位符
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			for key, value := range newContext {
				// 使用占位符（例如{{variable1}}）替换文本
				text := r.Text()
				fmt.Println(text)
				if strings.Contains(text, "{{"+key+"}}") {
					fmt.Printf("find :", text)
					newText := strings.ReplaceAll(text, "{{"+key+"}}", value)
					r.ClearContent()
					r.AddText(newText)
				}
			}
		}
	}
	for _, v := range doc.Headers() {
		for _, h := range v.Paragraphs() {
			for _, r := range h.Runs() {
				for key, value := range newContext {
					// 使用占位符（例如{{variable1}}）替换文本
					text := r.Text()
					fmt.Println(text)
					if strings.Contains(text, "{{"+key+"}}") {
						fmt.Printf("find :", text)
						newText := strings.ReplaceAll(text, "{{"+key+"}}", value)
						r.ClearContent()
						r.AddText(newText)
					}
				}
			}
		}
	}
	//// 保存渲染后的文档到目标文件
	//err = os.Remove(fileDest)
	//if err != nil && !os.IsNotExist(err) {
	//	log.Fatalf("Failed to remove existing destination file: %s", err)
	//}
	err = doc.SaveToFile(fileDest)
	if err != nil {
		log.Fatalf("Failed to save the rendered document: %s", err)
	}
	fmt.Println("Rendered document saved successfully!")

	aDoc, err := document.Open(fileDest)
	if err != nil {
		log.Fatalf("无法打开 a.docx 文件：%s", err)
	}

	bDoc, err := document.Open("test-gf-2023.docx")
	if err != nil {
		log.Fatalf("无法打开 b.docx 文件：%s", err)
	}

	for _, p := range bDoc.Paragraphs() {
		aDoc.AddParagraph().AddEndnote(p.Style())

	}

	if err = aDoc.Validate(); err != nil {
		log.Fatalf("合并文档后的验证失败：%s", err)
	}

	if err := aDoc.SaveToFile("merged.docx"); err != nil {
		log.Fatalf("保存合并后的文档失败：%s", err)
	}

	fmt.Println("文档合并成功！")

}
