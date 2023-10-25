package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"lessonProgram/db"
	"lessonProgram/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetTodos(c echo.Context) error {
	var todos []models.Todo
	q := c.QueryParam("q")
	var err error
	now := time.Now()

	if q == "weekly" {
		startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
		endOfWeek := startOfWeek.AddDate(0, 0, 7)
		err = db.Db.Where("DATE(date) >= ? AND DATE(date) <= ?", startOfWeek.Format("2006-01-02"), endOfWeek.Format("2006-01-02")).Preload("Student").Find(&todos).Error
	} else if q == "monthly" {
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
		err = db.Db.Where("DATE(date) >= ? AND DATE(date) <= ?", startOfMonth.Format("2006-01-02"), endOfMonth.Format("2006-01-02")).Preload("Student").Find(&todos).Error
	} else {
		err = db.Db.Model(&models.Todo{}).Preload("Student").Find(&todos).Error
	}

	if err != nil {
		fmt.Println("Todolar bulunamadı:", err)
	}

	return c.Render(http.StatusOK, "list-todos", map[string]interface{}{"Todos": todos})
}

func EditTodo(c echo.Context) error {
	var todo models.Todo
	err2 := db.Db.First(&todo, c.Param("id")).Error
	if err2 != nil {
		fmt.Println("Yapılacak bulunamadı:", err2)
	}

	var students []models.Student
	err := db.Db.Find(&students).Error
	if err != nil {
		fmt.Println("Öğrenciler bulunamadı:", err2)
	}

	startTime := todo.StartTime
	startParts := strings.Split(startTime, " ")
	parsedStartTime := startParts[1]

	endTime := todo.EndTime
	endParts := strings.Split(endTime, " ")
	parsedEndTime := endParts[1]

	return c.Render(http.StatusOK, "edit-todo", map[string]interface{}{"Todo": todo, "Students": students, "StartTime": parsedStartTime, "EndTime": parsedEndTime})
}

func UpdateTodo(c echo.Context) error {
	var todo models.Todo
	err2 := db.Db.First(&todo, c.Param("id")).Error
	if err2 != nil {
		fmt.Println("Todo bulunamadı:", err2)
	} else {
		uintValue, _ := strconv.ParseUint(c.FormValue("student_id"), 10, 0)
		todo.StudentID = uint(uintValue)
		todo.Title = c.FormValue("title")
		todo.StartTime = c.FormValue("date") + " " + c.FormValue("start_time")
		todo.EndTime = c.FormValue("date") + " " + c.FormValue("end_time")
		todo.Date = c.FormValue("date")
		statusStringValue := c.FormValue("status")
		statusIntegerValue, _ := strconv.Atoi(statusStringValue)
		todo.Status = statusIntegerValue

		startTime, err := time.Parse("2006-01-02 15:04", todo.StartTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"mesaj": "Başlangıç saati geçersiz."})
		}
		endTime, err := time.Parse("2006-01-02 15:04", todo.EndTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"mesaj": "Bitiş saati geçersiz."})
		}

		var existingTodo models.Todo
		err = db.Db.Where("date = ? AND ((? >= start_time AND ? <= end_time) AND student_id = ?)", todo.Date, startTime, endTime, todo.StudentID).First(&existingTodo).Error
		if err == nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"mesaj": "Bu tarih ve saat aralığında başka bir plan bulunuyor."})
		}

		db.Db.Save(&todo)
	}
	return c.Redirect(http.StatusSeeOther, "/todos")
}

func DeleteTodo(c echo.Context) error {
	var todo models.Todo
	err := db.Db.Find(&todo, c.Param("id")).Error
	if err != nil {
		fmt.Println("Yapılacak bulunamadı:", err)
	} else {
		db.Db.Delete(&todo)
	}
	return c.Redirect(http.StatusSeeOther, "/todos")
}

func CreateTodo(c echo.Context) error {
	var students []models.Student
	err := db.Db.Find(&students).Error
	if err != nil {
		fmt.Println("Öğrenciler bulunamadı:", err)
	}
	return c.Render(http.StatusOK, "add-todo", map[string]interface{}{"Students": students})
}

func SaveTodo(c echo.Context) error {
	var todo models.Todo
	uintValue, _ := strconv.ParseUint(c.FormValue("student_id"), 10, 0)
	todo.StudentID = uint(uintValue)
	todo.Title = c.FormValue("title")
	todo.StartTime = c.FormValue("date") + " " + c.FormValue("start_time")
	todo.EndTime = c.FormValue("date") + " " + c.FormValue("end_time")
	todo.Date = c.FormValue("date")
	todo.Status = 1

	startTime, err := time.Parse("2006-01-02 15:04", todo.StartTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"mesaj": "Başlangıç saati geçersiz."})
	}
	endTime, _ := time.Parse("2006-01-02 15:04", todo.EndTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"mesaj": "Bitiş saati geçersiz."})
	}

	var existingTodo models.Todo
	err = db.Db.Where("date = ? AND ((? >= start_time AND ? <= end_time) AND student_id = ?)", todo.Date, startTime, endTime, todo.StudentID).First(&existingTodo).Error
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"mesaj": "Bu tarih ve saat aralığında başka bir plan bulunuyor."})
	}

	db.Db.Create(&todo)
	return c.Redirect(http.StatusSeeOther, "/todos")
}
