package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"main.go/connection"
)

type Project struct {
	ID int
	Title string
	StartDate time.Time
	EndDate time.Time
	Duration string
	Desc string
	Javascript bool
	React bool
	PHP bool
	Java bool
	Image string
	FormatDateStart string
	FormatDateEnd string
}

func main() {
	connection.DatabaseConnect()
	e := echo.New()

	e.Static("/public", "public")

	// Routing method GET
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/projects", projects) // projects
	e.GET("/project/:id", projectDetail) // detail project
	e.GET("/formAddProjects", formAddProjects) // tambah project
	e.GET("/testimonial", testimonial)
	e.GET("/editProjects/:id", editProjects)
	
	// Routing Method POST
	e.POST("/addProjects", addProjects)
	e.POST("/editProjects/:id", editProjectsForm)
	e.POST("/deleteProject/:id", deleteProject)

	e.Logger.Fatal(e.Start(":5000"))

}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// Contact
func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// Detail Project
func projects(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, start_date, end_date, duration, description, javascript, react, php, java FROM tb_projects")
	
	var result []Project
	for data.Next() {
		var each = Project{}

		
		err := data.Scan(&each.ID, &each.Title, &each.StartDate, &each.EndDate, &each.Duration, &each.Desc, &each.Javascript, &each.React, &each.PHP, &each.Java)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		
		each.FormatDateStart = each.StartDate.Format("2 January 2006")
		each.FormatDateEnd = each.EndDate.Format("2 January 2006")

		result = append(result, each)
	}

	projects := map[string]interface{} {
		"Projects": result,
	}

	var tmpl, err = template.ParseFiles("views/projects.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	

	return tmpl.Execute(c.Response(), projects)
}

// Projects
func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	
	var projectDetail = Project{}
	
	 err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, duration, description, javascript, react, php, java, image FROM tb_projects WHERE id=$1", id).Scan(&projectDetail.ID, &projectDetail.Title, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Duration, &projectDetail.Desc, &projectDetail.Javascript, &projectDetail.React, &projectDetail.PHP, &projectDetail.Java, &projectDetail.Image)

	 projectDetail.FormatDateStart = projectDetail.StartDate.Format("2 January 2006")
	projectDetail.FormatDateEnd = projectDetail.EndDate.Format("2 January 2006")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{} {
		"Projects" : projectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/project.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

// nambah project
func formAddProjects(c echo.Context) error {

	var tmpl, err = template.ParseFiles("views/add-projects.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return tmpl.Execute(c.Response(), nil)
}

func addProjects(c echo.Context) error {
	title := c.FormValue("title")
	desc := c.FormValue("desc")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := calcDuration(startDate, endDate)
	// Author := c.FormValue("author")
	image := c.FormValue("image")

	var javascript bool
	if c.FormValue("javascript") == "yes" {
		javascript = true
	}
	var react bool
	if c.FormValue("react") == "yes" {
		react = true
	}
	var php bool
	if c.FormValue("php") == "yes" {
		php = true
	}
	var java bool
	if c.FormValue("java") == "yes" {
		java = true
	}

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (title, start_date, end_date, duration, description, javascript, react, php, java, image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", title, startDate,  endDate, duration , desc, javascript, react, php, java, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.Redirect(http.StatusMovedPermanently, "/projects")
}

// Edit Project
func editProjects(c echo.Context) error {

	id,_ := strconv.Atoi(c.Param("id"))


	var projectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, duration, description, javascript, react, php, java, image FROM tb_projects WHERE id=$1", id).Scan(&projectDetail.ID, &projectDetail.Title, &projectDetail.StartDate, &projectDetail.EndDate, &projectDetail.Duration, &projectDetail.Desc,&projectDetail.Javascript, &projectDetail.React, &projectDetail.PHP, &projectDetail.Java, &projectDetail.Image)

	data := map[string]interface{}{
		"Projects": projectDetail,
		"StartDate": projectDetail.StartDate.Format("2006-01-02"),
		"EndDate":   projectDetail.EndDate.Format("2006-01-02"),
	}

	
	var tmpl, errTemplate = template.ParseFiles("views/edit-projects.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)

}

func editProjectsForm(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("title")
	desc := c.FormValue("desc")
	StartDate := c.FormValue("startDate")
	EndDate := c.FormValue("endDate")
	duration := calcDuration(StartDate, EndDate)
	javascript := c.FormValue("javascript")
	react := c.FormValue("react")
	php := c.FormValue("php")
	java := c.FormValue("java")
	image := c.FormValue("image")

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET title=$1, start_date=$2, end_date=$3, duration=$4, description=$5, javascript=$6, react=$7, php=$8, java=$9, image=$10 WHERE id=$11", title, StartDate, EndDate, duration, desc, javascript != "", react != "", php != "", java != "", image, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/projects")
}

// Menghitung durasi
func calcDuration(startDate string, endDate string) string {
	StartDate,_ := time.Parse("2006-01-02", startDate)
	EndDate,_ := time.Parse("2006-01-02", endDate)

	durationTime := int(EndDate.Sub(StartDate).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonth := durationWeeks / 4
	durationYears := durationMonth / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " Years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + "Years"
	} else {
		if durationMonth > 1 {
			duration = strconv.Itoa(durationMonth) + " Month"
		} else if durationMonth > 0 {
			duration = strconv.Itoa(durationMonth) + "Month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " Weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " Weeks"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " Days"
				} else {
					duration = strconv.Itoa(durationDays) + " Days"
				}
			}
		}
	}
	return duration
}

// Delete Project
func deleteProject(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index : " , id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusMovedPermanently, map[string]string{"message": err.Error()})
	}
	// dataProjects = append(dataProjects[:id], dataProjects[id+1:] ... )

	return c.Redirect(http.StatusMovedPermanently, "/projects")
}

// Testimonial
func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}


