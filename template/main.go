package main

import (
	_ "Go/template/routers"
	"github.com/astaxie/beego"
	"path/filepath"
	"github.com/Joker/jade"
	"html/template"
	"Go/template/utils"
	"fmt"
	"github.com/yosssi/ace"
	"strings"
)

func main() {
	addJadeTemplate()
	addAceTemplate()
	beego.Run()
}
func addAceTemplate()  {
	beego.AddTemplateEngine("ace", func(root, path string, funcs template.FuncMap) (*template.Template, error) {
		aceOptions := &ace.Options{DynamicReload: true, FuncMap: funcs}
		aceBasePath := filepath.Join(root, "base")
		aceInnerPath := filepath.Join(root, strings.TrimSuffix(path, ".ace"))

		tpl, err := ace.Load(aceBasePath, aceInnerPath, aceOptions)

		if err != nil {
			return nil, fmt.Errorf("error loading ace template: %v", err)
		}

		return tpl, nil
	})
	
}

func addJadeTemplate() {
	beego.AddTemplateEngine("jade", func(root, path string, funcs template.FuncMap) (*template.Template, error) {
		jadePath := filepath.Join(root, path)
		content, err := utils.ReadFile(jadePath)
		fmt.Println(content)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		tpl, err := jade.Parse("name_of_tpl", content)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		fmt.Println("html:\n%s",tpl)
		tmp := template.New("Person template")
		tmp, err = tmp.Parse(tpl)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		fmt.Println(tmp)
		return tmp, err

	})
}


