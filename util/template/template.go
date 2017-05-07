package template

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	//NavMap 解析后的图书目录表
	NavMap struct {
		Title string `xml:"docTitle>text"`
		Navs  []nav  `xml:"navMap>navPoint"`
	}

	nav struct {
		Text    string `xml:"navLabel>text"`
		Src     src    `xml:"content"`
		SubNavs []nav  `xml:"navPoint"`
	}

	src struct {
		URL string `xml:"src,attr"`
	}
)

var (
	path   string
	navMap NavMap
)

//Build 构造页面框架
func Build(p string) error {
	path = p

	err := parseToc()
	if err != nil {
		return err
	}

	return nil
}

//ParseToc 解析图书的目录
func parseToc() error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, &navMap)
	if err != nil {
		return err
	}
	fmt.Println("Title: " + navMap.Title)
	subNav := navMap.Navs[0]
	for _, nav := range subNav.SubNavs {
		fmt.Printf("%s: %s\n", nav.Text, nav.Src.URL)
		for _, nav := range nav.SubNavs {
			fmt.Printf("\t%s: %s\n", nav.Text, nav.Src.URL)
		}
	}
	return nil
}
