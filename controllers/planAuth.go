package controllers

import (
	"net/http"

	"ahmetk3436.github.com/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func AddPlan(c echo.Context) error {
	plan := &models.Plan{}
	if err := c.Bind(&plan); err != nil {
		return err
	}
	newPlan := &models.Plan{}
	db := models.GetDB()
	if db.Table("plans").Where("day = ? AND start_time BETWEEN ? AND ? AND finish_time BETWEEN ? AND ?", plan.Day, plan.StartTime, plan.FinishTime, plan.StartTime, plan.FinishTime).Find(&newPlan); newPlan.ID != 0 {
		return c.String(http.StatusForbidden, "Belirtilen gün ve saat Aralığında başka planınız bulunduğu için plan eklenemiyor !")
	}
	plan.State = "Yapılıyor"
	models.GetDB().Create(&plan)
	return c.JSON(http.StatusCreated, plan)
}
func UpdatePlan(c echo.Context) error {
	plan := &models.Plan{}
	if err := c.Bind(&plan); err != nil {
		return err
	}
	temporaryPlan := &models.Plan{}
	temporaryPlan.Lesson = plan.Lesson
	temporaryPlan.LessonContent = plan.LessonContent
	temporaryPlan.Day = plan.Day
	temporaryPlan.StartTime = plan.StartTime
	temporaryPlan.FinishTime = plan.FinishTime
	db := models.GetDB()
	db.First(&plan)
	plan.Lesson = temporaryPlan.Lesson
	plan.LessonContent = temporaryPlan.LessonContent
	plan.Day = temporaryPlan.Day
	plan.StartTime = temporaryPlan.StartTime
	plan.FinishTime = temporaryPlan.FinishTime
	db.Table("plans").Where("id = ?", plan.ID).Save(&plan)
	return c.JSON(http.StatusOK, plan)
}
func DeletePlan(c echo.Context) error {
	id := c.QueryParam("id")
	plan := &models.Plan{}
	err := models.GetDB().Table("plans").Where("id = ?", id).First(plan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusBadRequest, "Belirtilen ID değerinde plan bulunamadı")
		}
		return c.String(http.StatusBadRequest, "Bağlantı hatası oluştu. Lütfen tekrar deneyiniz!")
	}
	models.GetDB().Delete(plan)
	return c.String(http.StatusOK, "Plan silindi !")
}
func FinishPlan(c echo.Context) error {
	id := c.QueryParam("id")
	plan := &models.Plan{}
	err := models.GetDB().Table("plans").Where("id = ?", id).First(plan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusBadRequest, "Belirtilen ID değerinde plan bulunamadı")
		}
		return c.String(http.StatusBadRequest, "Bağlantı hatası oluştu. Lütfen tekrar deneyiniz!")
	}
	plan.State = "Bitti"
	models.GetDB().Save(plan)
	return c.JSON(http.StatusOK, plan)
}

func CancelPlan(c echo.Context) error {
	id := c.QueryParam("id")
	plan := &models.Plan{}
	err := models.GetDB().Table("plans").Where("id = ?", id).First(plan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusBadRequest, "Belirtilen ID değerinde plan bulunamadı")
		}
		return c.String(http.StatusBadRequest, "Bağlantı hatası oluştu. Lütfen tekrar deneyiniz!")
	}
	plan.State = "İptal"
	models.GetDB().Save(plan)
	return c.JSON(http.StatusOK, plan)
}

func GetPlans(c echo.Context) error {
	db := models.GetDB()
	plans := []models.Plan{}

	if err := db.Find(&plans).Error; err != nil {
		return c.String(int(http.StateClosed), "Hata var bağlantı kapandı")
	}
	return c.JSON(http.StatusOK, plans)
}
func GetPlansWeekly(c echo.Context) error {
	db := models.GetDB()
	plans := []models.Plan{}
	//golangDateTime := time.Now().Format("2006-01-02")
	//golangDateTimeNextWeek := time.Now().Add(time.Hour * 24 * 7 * time.Duration(1)).Format("2006-01-02")

	if err := db.Raw("select * from plans where day < now() + interval 1 week and day > now()").Find(&plans).Error; err != nil {
		return c.String(int(http.StateClosed), "Hata var bağlantı kapandı")
	}
	return c.JSON(http.StatusOK, plans)
}
func GetPlansMonthly(c echo.Context) error {
	db := models.GetDB()
	plans := []models.Plan{}
	//golangDateTime := time.Now().Format("2006-01-02")
	//golangDateTimeNextMonth := time.Now().Add(time.Hour * 24 * 7 * time.Duration(4)).Format("2006-01-02")
	if err := db.Table("plans").Raw("select * from plans where day < now() + interval 4 week and day > now()").Find(&plans).Error; err != nil {
		return c.String(int(http.StateClosed), "Hata var bağlantı kapandı")
	}
	return c.JSON(http.StatusOK, plans)
}
