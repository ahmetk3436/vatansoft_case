package main

import (
	"net/http"

	"ahmetk3436.github.com/controllers"
	"ahmetk3436.github.com/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	models.GetDB()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/getStudents", controllers.GetStudents)
	e.GET("/getPlans", controllers.GetPlans)
	e.GET("/getPlansWeekly", controllers.GetPlansWeekly)
	e.GET("/getPlansMonthly", controllers.GetPlansMonthly)
	e.POST("/createStudent", controllers.CreateStudent)
	e.POST("/updateStudent", controllers.UpdateStudent)
	e.POST("/deleteStudent", controllers.DeleteStudent)
	e.POST("/addPlan", controllers.AddPlan)
	e.POST("/updatePlan", controllers.UpdatePlan)
	e.POST("/deletePlan", controllers.DeletePlan)
	e.POST("/cancelPlan", controllers.CancelPlan)
	e.POST("/finishPlan", controllers.FinishPlan)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Merhaba , Vatan Soft Case !")
}
