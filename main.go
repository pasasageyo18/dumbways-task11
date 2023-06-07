package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.GET("/hehe", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)

	e.GET("/project-detail/:id", projectDetail)
	e.GET("/form-project", formAddProject)
	e.POST("/add-project", addProject)

	e.Logger.Fatal(e.Start("localhost:7000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contactnew.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":        id,
		"Title":     "Mahmud",
		"StartDate": "12 Januari 2021",
		"EndDate":   "30 Januari 2022",
		"Content":   "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian Manpower Group, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Khusus di sektor teknologi yang berkembang pesat, menurut Kemendikbudristek, Indonesia kekurangan sembilan juta pekerja teknologi hingga tahun 2030. Hal itu berarti Indonesia memerlukan sekitar 600 ribu SDM digital yang memasuki pasar setiap tahunnya.",
	}

	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/addpro.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	title := c.FormValue("inputTitle")
	startDate := c.FormValue("inputStartDate")
	endDate := c.FormValue("inputEndDate")
	content := c.FormValue("inputProContent")
	tech1 := c.FormValue("tech1")
	tech2 := c.FormValue("tech2")
	tech3 := c.FormValue("tech3")
	tech4 := c.FormValue("tech4")

	println("Title : " + title)
	println("startDate : " + startDate)
	println("endDate : " + endDate)
	println("Content : " + content)
	println("tech1 : " + tech1)
	println("tech2 : " + tech2)
	println("tech3 : " + tech3)
	println("tech4 : " + tech4)

	return c.Redirect(http.StatusMovedPermanently, "/add-project")
}
