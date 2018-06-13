package handler

import (
	"net/http"
	"github.com/mrtomyum/nopadol/incentive"
	"html/template"
)

// New creates new domain1 handler
func New(ic incentive.Service) http.Handler {
	c := ctrl{}
	c.ic = ic

	c.templates = make(map[string]*template.Template)
	c.templates["index"] = template.Must(template.ParseFiles("template/incentive/index.tmpl"))

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(c.Index))

	return mux
}

type ctrl struct {
	templates map[string]*template.Template
	ic         incentive.Service
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