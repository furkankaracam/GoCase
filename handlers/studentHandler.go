package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"lessonProgram/db"
	"lessonProgram/models"
	"net/http"
)

func GetStudents(c echo.Context) error {
	var students []models.Student
	err2 := db.Db.Find(&students).Error
	if err2 != nil {
		fmt.Println("Öğrenciler bulunamadı:", err2)
	}
	return c.Render(http.StatusOK, "list-students", map[string]interface{}{"Students": students})
}

func EditStudent(c echo.Context) error {
	var student models.Student
	err2 := db.Db.First(&student, c.Param("id")).Error
	if err2 != nil {
		fmt.Println("Öğrenci bulunamadı:", err2)
	}
	return c.Render(http.StatusOK, "edit-student", map[string]interface{}{"Student": student})
}

func UpdateStudent(c echo.Context) error {
	var student models.Student
	err2 := db.Db.First(&student, c.Param("id")).Error
	if err2 != nil {
		fmt.Println("Öğrenci bulunamadı:", err2)
	} else {
		student.Name = c.FormValue("name")
		student.Surname = c.FormValue("surname")
		student.SchoolNumber = c.FormValue("school_number")
		db.Db.Save(&student)
	}
	return c.Redirect(http.StatusSeeOther, "/students")
}

func DeleteStudent(c echo.Context) error {
	var student models.Student
	err := db.Db.Find(&student, c.Param("id")).Error
	if err != nil {
		fmt.Println("Öğrenci bulunamadı:", err)
	} else {
		var todos []models.Todo
		db.Db.Where("student_id = ?", student.ID).Find(&todos)

		for _, todo := range todos {
			db.Db.Delete(&todo)
		}
		db.Db.Delete(&student)
	}
	return c.Redirect(http.StatusSeeOther, "/students")
}

func CreateStudent(c echo.Context) error {
	return c.Render(http.StatusOK, "add-student", "")
}

func SaveStudent(c echo.Context) error {
	var student models.Student
	student.Name = c.FormValue("name")
	student.Surname = c.FormValue("surname")
	student.SchoolNumber = c.FormValue("school_number")
	db.Db.Create(&student)
	return c.Redirect(http.StatusSeeOther, "/students")
}
