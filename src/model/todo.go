package model

import "gorm.io/gorm"

type TodoInfo struct {
	gorm.Model `gorm:"embedded" swaggerignore:"true"`
	Title      string `gorm:"comment:title;Not Null;index:idx_todoTitle;"`
	Content    string `gorm:"comment:content;"`
	Status     int    `gorm:"comment:todoStatus;"`
	CreatedBy  string `gorm:"comment:createdBy;"`
	UpdatedBy  string `gorm:"comment:updatedBy;"`
}

type TodoMessage struct {
	gorm.Model `gorm:"embedded" swaggerignore:"true"`
	MailTo     string   `gorm:"comment:to;Not Null;"`
	Todo       TodoInfo `gorm:"embedded"`
	Status     bool     `gorm:"comment:status;default false;"`
}

type TodoStatus int

const (
	Done TodoStatus = iota
	Waiting
	Approved
)

var TodoStatusMap = []string{"Done", "Waiting", "Approved"}

func (s TodoStatus) String() string {
	return TodoStatusMap[s]
}

func GetTodoStatus(status string) TodoStatus {
	switch status {
	case "Done":
		return Done
	case "Waiting":
		return Waiting
	case "Approved":
		return Approved
	}
	return -1
}
