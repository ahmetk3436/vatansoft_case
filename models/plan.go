package models

import (
	"github.com/jinzhu/gorm"
)

type Plan struct {
	gorm.Model
	Lesson        string `json:"lesson"`
	LessonContent string `json:"lessonContent"`
	Day           string `json:"day"`
	StartTime     string `json:"startTime"`
	FinishTime    string `json:"finishTime"`
	State         string
}
