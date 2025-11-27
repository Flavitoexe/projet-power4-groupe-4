package game

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var listTemplates *template.Template

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var buffer bytes.Buffer

	errRender := listTemplates.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message%s", http.StatusInternalServerError,
			url.QueryEscape("Une erreur est survenue lors du cgargement de la page")), http.StatusSeeOther)
		return
	}
	buffer.WriteTo(w)
}
