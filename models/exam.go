package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Subject struct {
	gorm.Model
	SubjectName string `json:"subject_name"`
}

type QuestionBase struct {
	Type      int    `json:"type"`
	Score     int    `json:"score"`
	Desc      string `json:"desc"`
	Answer    string `json:"answer"`
	SubjectId uint   `json:"subject_id"` //
	UserId    uint   `json:"user_id"`
}

type Question struct {
	gorm.Model
	QuestionBase
	UserId uint
}

type Paper struct {
	gorm.Model  `json:"extra"`
	QuestionIds []uint `json:"question_ids"`
	PaperName   string `json:"paper_name"`
	UserId      uint   `json:"user_id"`
	Desc        string `json:"desc"`
}

type Exam struct {
	gorm.Model  `json:"extra"`
	PaperIds    []uint `json:"paper_ids"`
	ExamName    string `json:"exam_name"`
	Desc        string `json:"desc"`
	ExamineeIds []uint `json:"examinee_ids"`
}

type Examinee struct {
	UserId     uint      `json:"user_id"`
	Status     int       `json:"status"`
	AttendTime time.Time `json:"attend_time"`
	Score      int       `json:"score"`
	ExamId     uint      `json:"exam_id"`
}

