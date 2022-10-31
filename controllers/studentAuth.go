package controllers

import (
	"net/http"

	"ahmetk3436.github.com/models"
	"github.com/labstack/echo"
)

func CreateStudent(c echo.Context) error {
	student := &models.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	newStudent := &models.Student{}
	db := models.GetDB()
	if db.Table("students").Where("email = ?", student.Email).Find(&newStudent); newStudent.Email == "" {
		db.Create(&student)
	} else {
		return c.String(http.StatusBadRequest, "Bu e-posta ile kayıt mevcut !")
	}
	return c.JSON(http.StatusCreated, student)
}
func UpdateStudent(c echo.Context) error {
	student := &models.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	temporaryStudent := &models.Student{}
	temporaryStudent.Email = student.Email
	temporaryStudent.Password = student.Password
	temporaryStudent.Surname = student.Surname
	temporaryStudent.Name = student.Name
	db := models.GetDB()
	db.First(&student)
	student.Email = temporaryStudent.Email
	student.Name = temporaryStudent.Name
	student.Surname = temporaryStudent.Surname
	student.Password = temporaryStudent.Password
	db.Save(&student)
	return c.JSON(http.StatusOK, student)
}
func DeleteStudent(c echo.Context) error {
	id := c.QueryParam("id")
	models.GetDB().Delete(&models.Student{}, id)
	return c.String(http.StatusOK, "Öğrenci Başarıyla Silindi")

}
func GetStudents(c echo.Context) error {
	db := models.GetDB()
	students := []models.Student{}

	err := db.Find(&students).Error
	if err != nil {
		return c.String(http.StatusBadRequest, "Hata mevcut öğrenciler getirilemedi !")
	}
	return c.JSON(http.StatusOK, students)
}
