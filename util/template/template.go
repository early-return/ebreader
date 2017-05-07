package template

import (
	"ebreader/config"
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"os"
)

type (
	//NavMap 解析后的图书目录表
	NavMap struct {
		Title  string `xml:"docTitle>text"`
		Author string `xml:"docAuthor>text"`
		Navs   []nav  `xml:"navMap>navPoint"`
	}

	nav struct {
		Title   string `xml:"navLabel>text"`
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

	// 如果原目录中已存在index.html，则重命名为index.bak.html
	indexPath := config.Path + "/index.html"
	os.Rename(indexPath, config.Path+"/index.bak.html")

	err = parseTemplate(indexPath)
	if err != nil {
		return err
	}

	return nil
}

//渲染模板
func parseTemplate(file string) error {
	t := template.New("template")
	t = t.Funcs(template.FuncMap{"getFirstSrc": getFirstSrc})
	t, err := t.Parse(htmlTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, navMap)
	if err != nil {
		return err
	}
	return nil
}

func getFirstSrc(navs []nav) string {
	if navs[0].Src.URL == "index.html" {
		return "index.bak.html"
	}
	return navs[0].Src.URL
}

//parseToc 解析图书的目录
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
	return nil
}
