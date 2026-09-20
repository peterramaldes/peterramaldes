package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func buildFromGlob(file, glob string, data map[string]interface{}) error {
	out, err := os.Create(file)
	defer out.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseGlob(glob)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}

func buildFromFile(file, tmpl string, data map[string]interface{}) error {
	out, err := os.Create(file)
	defer out.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}

func main() {
	data := map[string]interface{}{}
	buf, err := os.ReadFile("data.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return
	}
	entries, err := os.ReadDir("tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		name := entry.Name()
		var err error
		if entry.IsDir() {
			err = buildFromGlob(name, filepath.Join("tmpl", name, "*.html"), data)
		} else {
			err = buildFromFile(name, filepath.Join("tmpl", name), data)
		}
		if err != nil {
			fmt.Println(name, err)
		}
	}
}
