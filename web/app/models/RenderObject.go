package models

import "html/template"

type RenderObject struct {
	Title string
	Name  template.HTML
	Lang  string
}
