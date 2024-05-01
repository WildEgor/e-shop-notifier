package template

import (
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"path"
)

const templatePath = "internal/templates/"

type TemplateCache struct {
	templates map[string]*template.Template
}

func (t *TemplateCache) Init() {

	t.templates = make(map[string]*template.Template)

	tp := path.Join(templatePath)

	files, err := os.ReadDir(tp)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	for _, file := range files {
		tml, err := template.ParseFiles(path.Join(templatePath, file.Name()))
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		t.templates[file.Name()] = tml
	}
}

func (t *TemplateCache) GetByName(name string) *template.Template {
	fn := fmt.Sprintf("%s.html", name)

	return t.templates[fn]
}
