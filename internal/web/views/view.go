package views

import (
	"fmt"
	"net/http"
	"text/template"
)

type View struct {
	mainTmpl *template.Template
}

func NewView() (*View, error) {
	view := &View{
		mainTmpl: template.New("mainTmpl"),
	}

	_, err := view.mainTmpl.ParseFiles("web/bootstrap.html", "web/navbar.html")
	if err != nil {
		return nil, fmt.Errorf("View.NewView: %w", err)
	}

	return view, nil
}

// init static files
func (v *View) InitRoutes(mux *http.ServeMux) {
	fsStatic := http.FileServer(http.Dir("web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fsStatic))
}

func (v *View) NewTemplate(file string) (*template.Template, error) {
	newTmpl, err := v.mainTmpl.Clone()
	if err != nil {
		return nil, fmt.Errorf("View.NewTemplate: %w", err)
	}

	fullFile := fmt.Sprintf("web/%s.html", file)
	_, err = newTmpl.ParseFiles(fullFile)
	if err != nil {
		return nil, fmt.Errorf("View.NewTemplate: %w", err)
	}

	return newTmpl, nil
}

func (v *View) ErrorTemplate() (*template.Template, error) {
	newTmpl, err := v.mainTmpl.Clone()
	if err != nil {
		return nil, fmt.Errorf("View.ErrorTemplate: %w", err)
	}

	_, err = newTmpl.ParseFiles("web/error.html")
	if err != nil {
		return nil, fmt.Errorf("View.ErrorTemplate: %w", err)
	}

	return newTmpl, nil
}
