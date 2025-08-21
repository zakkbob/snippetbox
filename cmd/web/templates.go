package main

import "github.com/zakkbob/snippetbox/internal/models"

// A holding structure for data we want to pass to templates, since only one piece of dynamic data can be passed
type templateData struct {
	Snippet models.Snippet
}
