package cmd

import (
	"fmt"
	"os"
	"text/template"
)

func writeProjectTemplate(path, tmplContent, name string) {
	tmpl := template.Must(template.New("").Parse(tmplContent))
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("❌ Failed to create file:", path)
		return
	}
	defer f.Close()
	tmpl.Execute(f, struct{ ProjectName string }{ProjectName: name})
}

func writeModuleTemplate(path, tmplContent string, data any) {
	funcMap := template.FuncMap{
		"title": func(s string) string {
			return string(upperFirst(s))
		},
	}
	tmpl := template.Must(template.New("").Funcs(funcMap).Parse(tmplContent))
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("❌ Failed to create file:", path)
		return
	}
	defer f.Close()
	tmpl.Execute(f, data)
}

func writeFile(path, content string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("❌ Failed to cteate file:", path)
		return
	}
	defer f.Close()
	f.WriteString(content)
}
