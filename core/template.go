package core

import "html/template"

const (
	// AuthFormTemplate name
	AuthFormTemplate string = "authForm.html"
)

// Use GetTemplates to get an instance
var twitchTemplates *template.Template

// GetTemplates return the reference to template.Template containing the twitch template
func GetTemplates() *template.Template {
	if twitchTemplates == nil {
		twitchTemplates = template.Must(template.ParseGlob("./core/template/*.html"))
	}

	return twitchTemplates
}

// GetTemplate will fetch the requested template with the context
func GetTemplate(templateName string) *template.Template {
	template := GetTemplates().Lookup(templateName)
	return template
}
