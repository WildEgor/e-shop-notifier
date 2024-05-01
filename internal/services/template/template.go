package template

import (
	"bytes"
	"errors"
)

var (
	ErrTemplateNotFound = errors.New("cannot find template")
	ErrParseTemplate    = errors.New("cannot parse template")
)

type TemplateService struct {
	cache *TemplateCache
}

func NewTemplateService() *TemplateService {

	cache := &TemplateCache{}
	cache.Init()

	return &TemplateService{
		cache,
	}
}

func (t *TemplateService) Build(name string, data interface{}) (string, error) {

	tmpl := t.cache.GetByName(name)
	if tmpl == nil {
		return "", ErrTemplateNotFound
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return "", ErrParseTemplate
	}

	return buf.String(), nil
}
