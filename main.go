package main

import (
	// "fmt"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id int
	Title string
	Desc string
	Tech []string
	StartDate string
	EndDate string
	Duration string
	Author string
}

var dataProjects = []Project {}

func main() {
	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })

	e.Static("/public", "public")
	// Routing
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/projects", projects) // projects
	e.GET("/project/:id", projectDetail) // detail project
	e.GET("/addProjects", addProjects) // tambah project
	e.GET("/testimonial", testimonial)
	e.GET("/editProject/:id", editProject)
	
	e.POST("/formEditProject/:id", formEditProject)
	e.POST("/formAddProjects", formAddProjects)
	e.POST("/addProjects/:id", updateProject)
	e.POST("/deleteProject/:id", deleteProject)
	e.Logger.Fatal(e.Start("localhost:5000"))

}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func projects(c echo.Context) error {

	var tmpl, err = template.ParseFiles("views/projects.html")
	
	

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{} {
		"Projects": dataProjects,
	}

	return tmpl.Execute(c.Response(), projects)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// data := map[string]interface{} {
	// 	"id" : id,
	// 	"title" : "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
	// 	"content" : "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian Manpower Group, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Khusus di sektor teknologi yang berkembang pesat, menurut Kemendikbudristek, Indonesia kekurangan sembilan juta pekerja teknologi hingga tahun 2030. Hal itu berarti Indonesia memerlukan sekitar 600 ribu SDM digital yang memasuki pasar setiap tahunnya.",
	// }

	var projectDetail = Project{}

	for i, data := range dataProjects {
		if id == i {
			projectDetail = Project{
				Title: data.Title,
				Desc: data.Desc,
				Tech: data.Tech,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				Author: data.Author,
			}
		}
	}

	data := map[string]interface{} {
		"Projects" : projectDetail,
	}

	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}


func addProjects(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-projects.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

// nambah project
func formAddProjects(c echo.Context) error {
	title := c.FormValue("title")
	desc := c.FormValue("desc")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := Durasi(startDate, endDate)
	tech := c.Request().Form["technologies"]


	// println("title : " + title)
	// println("desc : " + desc)
	// println("author : " + Author)
	println("Tech : ", strings.Join(tech, ", "))

	var newProject = Project {
		Title: title,
		Desc: desc,
		Duration: duration,
		Author: "Anonymous",
		Tech: tech,
	}

	dataProjects = append(dataProjects, newProject)

	fmt.Println(dataProjects)


	return c.Redirect(http.StatusMovedPermanently, "/projects")
}

// Update Project
func updateProject( c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))

	dataProjects = append(dataProjects[:i], dataProjects[i+ 1])

	return c.Redirect(http.StatusMovedPermanently, "/projects")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range dataProjects{
		if id == i {
			ProjectDetail = Project{
				Id: id,
				Title: data.Title,
				Desc: data.Desc,
				Tech: data.Tech,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}
	var tmpl, err = template.ParseFiles("views/edit-projects.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}

// Delete Project
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index : " , id)

	dataProjects = append(dataProjects[:id], dataProjects[id+1:] ... )

	return c.Redirect(http.StatusMovedPermanently, "/projects")
}



func formEditProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("title")
	desc := c.FormValue("desc")
	tech := c.Request().Form["technologies"]

	var updateProject = Project {
		Title: title,
		Desc: desc,
		Tech: tech,
	}

	dataProjects[id] = updateProject

	return c.Redirect(http.StatusMovedPermanently, "/projects")
}


// func projectDetail(c echo.Context) error {
// 	data := map[string]interface{} {
// 		"login" : true
// 	}

// 	var tmpl,err = template.ParseFiles("views")
// }

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// Menghitung durasi
func Durasi(startDate, endDate string) string {
	startdatestr, _ := time.Parse("2006-01-02", startDate)
	enddatestr, _ := time.Parse("2006-01-02", endDate)

	durasiTime := int(enddatestr.Sub(startdatestr).Hours())
	durasiDays := durasiTime / 24
	durasiWeeks := durasiDays / 7
	durasiMounth := durasiWeeks / 4
	durasiYears := durasiMounth / 12

	var durasi string

	if durasiYears > 1 {
		durasi = strconv.Itoa(durasiYears) + " Tahun"
	} else if durasiYears > 0 {
		durasi = strconv.Itoa(durasiYears) + "Tahun"
	} else {
		if durasiMounth > 1 {
			durasi = strconv.Itoa(durasiMounth) + " Bulan"
		} else if durasiMounth > 0 {
			durasi = strconv.Itoa(durasiMounth) + "Bulan"
		} else {
			if durasiWeeks > 1 {
				durasi = strconv.Itoa(durasiWeeks) + " Minggu"
			} else if durasiWeeks > 0{
				durasi = strconv.Itoa(durasiWeeks) + "Minggu"
			} else {
				if durasiDays > 1 {
					durasi = strconv.Itoa(durasiDays) + " Hari"
				} else {
					durasi = strconv.Itoa(durasiDays) + "Hari"
				}
			}
		}
	}

	return durasi
	
}