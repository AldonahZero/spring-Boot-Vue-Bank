package sysModel

import (
	"fmt"
	"gin-vue-admin/init/qmsql"
	"github.com/jinzhu/gorm"
)

// 生成企业邮箱
func GenerateEmail(employeeNo string) string {
	return fmt.Sprintf("%s@company.com", employeeNo)
}

// 入职流程主表
type SysOnboarding struct {
	gorm.Model
	EmployeeName     string                 `json:"employeeName" gorm:"comment:'员工姓名'"`
	EmployeeNo       string                 `json:"employeeNo" gorm:"unique;comment:'员工工号'"`
	Email            string                 `json:"email" gorm:"comment:'企业邮箱'"`
	Phone            string                 `json:"phone" gorm:"comment:'联系电话'"`
	DepartmentId     string                 `json:"departmentId" gorm:"comment:'部门ID'"`
	Position         string                 `json:"position" gorm:"comment:'职位'"`
	EntryDate        string                 `json:"entryDate" gorm:"comment:'入职日期'"`
	Status           int                    `json:"status" gorm:"default:0;comment:'状态:0-待处理,1-进行中,2-已完成,3-已取消'"`
	Progress         float64                `json:"progress" gorm:"default:0;comment:'进度百分比'"`
	CurrentStep      int                    `json:"currentStep" gorm:"default:0;comment:'当前步骤'"`
	TotalSteps       int                    `json:"totalSteps" gorm:"default:5;comment:'总步骤数'"`
	MentorId         uint                   `json:"mentorId" gorm:"comment:'导师ID'"`
	CreatedBy        uint                   `json:"createdBy" gorm:"comment:'创建人ID'"`
	OnboardingTasks  []SysOnboardingTask    `json:"onboardingTasks" gorm:"foreignKey:OnboardingId"`
}

// 入职任务表
type SysOnboardingTask struct {
	gorm.Model
	OnboardingId  uint   `json:"onboardingId" gorm:"comment:'入职流程ID'"`
	TaskType      string `json:"taskType" gorm:"comment:'任务类型:IT-IT部门,ADMIN-行政部门,MENTOR-导师,HR-HR部门'"`
	TaskName      string `json:"taskName" gorm:"comment:'任务名称'"`
	Description   string `json:"description" gorm:"comment:'任务描述'"`
	AssigneeId    uint   `json:"assigneeId" gorm:"comment:'负责人ID'"`
	Status        int    `json:"status" gorm:"default:0;comment:'状态:0-待处理,1-进行中,2-已完成'"`
	DueDate       string `json:"dueDate" gorm:"comment:'截止日期'"`
	CompletedAt   string `json:"completedAt" gorm:"comment:'完成时间'"`
	CompletedBy   uint   `json:"completedBy" gorm:"comment:'完成人ID'"`
	Remarks       string `json:"remarks" gorm:"comment:'备注'"`
}

// 创建入职流程
func (o *SysOnboarding) Create() error {
	err := qmsql.DEFAULTDB.Create(&o).Error
	return err
}

// 更新入职流程
func (o *SysOnboarding) Update() error {
	err := qmsql.DEFAULTDB.Save(&o).Error
	return err
}

// 删除入职流程
func (o *SysOnboarding) Delete() error {
	err := qmsql.DEFAULTDB.Delete(&o).Error
	return err
}

// 根据ID获取入职流程
func (o *SysOnboarding) GetById(id uint) (err error) {
	err = qmsql.DEFAULTDB.Preload("OnboardingTasks").First(&o, id).Error
	return err
}

// 获取入职流程列表
func (o *SysOnboarding) GetList(page, pageSize int, status int, departmentId string) (list []SysOnboarding, total int, err error) {
	db := qmsql.DEFAULTDB.Model(&SysOnboarding{})
	if status >= 0 {
		db = db.Where("status = ?", status)
	}
	if departmentId != "" {
		db = db.Where("department_id = ?", departmentId)
	}
	db.Count(&total)
	err = db.Order("created_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, total, err
}

// 更新任务状态
func (t *SysOnboardingTask) UpdateStatus(status int, completedBy uint, remarks string) error {
	updates := map[string]interface{}{
		"status":       status,
		"completed_by": completedBy,
		"remarks":      remarks,
	}
	if status == 2 {
		updates["completed_at"] = "now()"
	}
	err := qmsql.DEFAULTDB.Model(&t).Updates(updates).Error
	return err
}

// 生成工号
func GenerateEmployeeNo() string {
	var count int64
	qmsql.DEFAULTDB.Model(&SysOnboarding{}).Count(&count)
	return fmt.Sprintf("EMP%06d", count+1)
}
