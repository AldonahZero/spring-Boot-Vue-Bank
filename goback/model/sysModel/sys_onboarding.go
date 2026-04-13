package sysModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	"time"
)

type OnboardingProcess struct {
	gorm.Model
	EmployeeName      string    `json:"employeeName" gorm:"comment:员工姓名"`
	EmployeeNo        string    `json:"employeeNo" gorm:"comment:工号"`
	Email             string    `json:"email" gorm:"comment:邮箱"`
	Department        string    `json:"department" gorm:"comment:部门"`
	Position          string    `json:"position" gorm:"comment:职位"`
	OnboardingDate    time.Time `json:"onboardingDate" gorm:"comment:入职日期"`
	Phone             string    `json:"phone" gorm:"comment:联系电话"`
	EmergencyContact  string    `json:"emergencyContact" gorm:"comment:紧急联系人"`
	EmergencyPhone    string    `json:"emergencyPhone" gorm:"comment:紧急联系电话"`
	Status            string    `json:"status" gorm:"default:'pending';comment:状态(pending/processing/completed)"`
	CurrentStep       int       `json:"currentStep" gorm:"default:1;comment:当前步骤"`
	ITTaskID          uint      `json:"itTaskId" gorm:"comment:IT部门任务ID"`
	AdminTaskID       uint      `json:"adminTaskId" gorm:"comment:行政部门任务ID"`
	MentorTaskID      uint      `json:"mentorTaskId" gorm:"comment:导师任务ID"`
	ChecklistItems    []OnboardingChecklist `json:"checklistItems"`
}

type OnboardingChecklist struct {
	gorm.Model
	OnboardingProcessID uint   `json:"onboardingProcessId"`
	ItemName            string `json:"itemName" gorm:"comment:检查项名称"`
	ItemCategory        string `json:"itemCategory" gorm:"comment:分类(it/admin/mentor)"`
	IsCompleted         bool   `json:"isCompleted" gorm:"default:false;comment:是否完成"`
	CompletedBy         string `json:"completedBy" gorm:"comment:完成人"`
	CompletedAt         *time.Time `json:"completedAt" gorm:"comment:完成时间"`
	Remark              string `json:"remark" gorm:"comment:备注"`
}

type OnboardingTask struct {
	gorm.Model
	OnboardingProcessID uint      `json:"onboardingProcessId"`
	Department          string    `json:"department" gorm:"comment:负责部门(it/admin/mentor)"`
	TaskType            string    `json:"taskType" gorm:"comment:任务类型"`
	TaskContent         string    `json:"taskContent" gorm:"comment:任务内容"`
	Status              string    `json:"status" gorm:"default:'pending';comment:状态(pending/completed)"`
	Assignee            string    `json:"assignee" gorm:"comment:负责人"`
	DueDate             time.Time `json:"dueDate" gorm:"comment:截止日期"`
	CompletedAt         *time.Time `json:"completedAt" gorm:"comment:完成时间"`
	Remark              string    `json:"remark" gorm:"comment:备注"`
}

func (op *OnboardingProcess) Create() error {
	return qmsql.DEFAULTDB.Create(&op).Error
}

func (op *OnboardingProcess) GetByID() (error, *OnboardingProcess) {
	var process OnboardingProcess
	err := qmsql.DEFAULTDB.Preload("ChecklistItems").Where("id = ?", op.ID).First(&process).Error
	return err, &process
}

func (op *OnboardingProcess) GetList(info modelInterface.PageInfo) (error, interface{}, int) {
	err, db, total := servers.PagingServer(op, info)
	if err != nil {
		return err, nil, 0
	}
	var list []OnboardingProcess
	err = db.Preload("ChecklistItems").Find(&list).Error
	return err, list, total
}

func (op *OnboardingProcess) Update() error {
	return qmsql.DEFAULTDB.Save(&op).Error
}

func (op *OnboardingProcess) Delete() error {
	return qmsql.DEFAULTDB.Where("id = ?", op.ID).Delete(&op).Error
}

func (ot *OnboardingTask) Create() error {
	return qmsql.DEFAULTDB.Create(&ot).Error
}

func (ot *OnboardingTask) Update() error {
	return qmsql.DEFAULTDB.Save(&ot).Error
}

func (ot *OnboardingTask) GetByProcessID(processID uint) (error, []OnboardingTask) {
	var tasks []OnboardingTask
	err := qmsql.DEFAULTDB.Where("onboarding_process_id = ?", processID).Find(&tasks).Error
	return err, tasks
}

func (oc *OnboardingChecklist) Create() error {
	return qmsql.DEFAULTDB.Create(&oc).Error
}

func (oc *OnboardingChecklist) Update() error {
	return qmsql.DEFAULTDB.Save(&oc).Error
}

func GenerateEmployeeNo() string {
	return "EMP" + time.Now().Format("20060102") + randomString(4)
}

func GenerateEmail(name string) string {
	return name + "@company.com"
}

func randomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().Nanosecond()%len(charset)]
	}
	return string(b)
}
