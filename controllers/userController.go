package controllers

import (
	"github.com/Muart-C/webcbo/config"
	"github.com/Muart-C/webcbo/repository"
	"github.com/Muart-C/webcbo/views"
	"html/template"
	"net/http"
)

//DASHBOARD
//Dashboard view definition
func Dashboard(w http.ResponseWriter, r *http.Request) error{
	v := views.New(r)
	v.Render(w,"admin/dashboard/index")
	return nil
}

//ListProjects view definition
func ListUsers(w http.ResponseWriter, r *http.Request) error{

	v := views.New(r)
	pageNum := config.GetPageNum(r)

	p, er := views.NewPaginator(pageNum, config.PerPage, r.URL)
	if er != nil {
		 RespondWithError(w,http.StatusNotFound,"error happened")
	}


	users,numPages,err := repository.FetchUsers(p.Start, p.Limit)
	if err != nil {
		RespondWithError(w,http.StatusNotFound,"error happened while retrieving users")
	}

	pagination := p.Render(numPages)
	if err != nil {
		RespondWithError(w,http.StatusNotFound,"error happened while during pagination")
	}
	v.Vars["Users"] = users
	v.Vars["Pagination"] = template.HTML(pagination)
	v.Render(w, "admin/users/index")

	return nil
}

