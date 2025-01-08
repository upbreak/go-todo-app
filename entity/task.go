package entity

import (
	"time"
)

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	//ID      TaskID     `json:"id"`
	//Title   string     `json:"title"`
	//Status  TaskStatus `json:"status"`
	//Created time.Time  `json:"created"`
	IDX   TaskID `json:"idx" db:"IDX"`
	SNO   int64  `json:"sno" db:"SNO"`
	TITLE string `json:"title" db:"TITLE" validate:"required"`
	//CONTENT  godror.Lob `json:"content" db:"CONTENT"`
	CONTENT  string    `json:"content" db:"CONTENT" validate:"required"`
	SHOW_YN  string    `json:"show_yn" db:"SHOW_YN"`
	IS_USE   string    `json:"is_use" db:"IS_USE"`
	REG_UNO  int64     `json:"reg_uno" db:"REG_UNO" validate:"required"`
	REG_USER string    `json:"reg_user" db:"REG_USER" validate:"required"`
	REG_DATE time.Time `json:"reg_date" db:"REG_DATE"`
}

type Tasks []*Task

type TaskListRsp struct {
	ID    TaskID `json:"id"`
	Sno   int64  `json:"sno"`
	Title string `json:"title"`
	//Content godror.Lob    `json:"content"`
	Content string    `json:"content"`
	ShowYn  string    `json:"show_yn"`
	IsUse   string    `json:"is_use"`
	RegUno  int64     `json:"reg_uno"`
	RegUser string    `json:"reg_user"`
	RegDate time.Time `json:"reg_date"`
}
