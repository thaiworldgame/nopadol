package handler

import (
	"net/http"
	"github.com/mrtomyum/nopadol/sale"
	"html/template"
)

// New creates new domain1 handler
func New(s sale.Service) http.Handler {
	c := ctrl{}
	c.s = s

	c.templates = make(map[string]*template.Template)
	c.templates["index"] = template.Must(template.ParseFiles("template/sale/index.tmpl"))

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(c.Index))

	return mux
}

type ctrl struct {
	templates map[string]*template.Template
	s         sale.Service
}

func (c *ctrl) render(w http.ResponseWriter, name string, data interface{}) {
	tmpl := c.templates[name]
	if tmpl == nil {
		// this can panic, since it should never happened in production
		panic("template not found")
	}

	w.Header().Set("Content-Type", "text/html; chatset=utf-8")
	tmpl.Execute(w, data)
}

func (c *ctrl) Index(w http.ResponseWriter, r *http.Request) {
	c.render(w, "index", nil)
}