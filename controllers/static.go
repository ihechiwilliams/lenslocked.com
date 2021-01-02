package controllers

import "github.com/charleyvibez/lenslocked.com/views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView(
			"bootstrap", "views/static/contact.gohtml"),
		Faq: views.NewView(
			"bootstrap", "views/static/faq.gohtml"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}
