package sysModel

import (
	"gin-vue-admin/init/qmsql"
	"github.com/jinzhu/gorm"
)

// 离职流程主表
type SysOffboarding struct {
	gorm.Model
	EmployeeId       uint                   `json:"employeeId" gorm:"comment:'员工ID'"`
	EmployeeName     string                 `json:"employeeName" gorm:"comment:'员工姓名'"`
	EmployeeNo       string                 `json:"employeeNo" gorm:"comment:'员工工号'"`
	DepartmentId     string                 `json:"departmentId" gorm:"comment:'部门ID'"`
	Position         string                 `json:"position" gorm:"comment:'职位'"`
	ResignationDate  string                 `json:"resignationDate" gorm:"comment:'申请离职日期'"`
	LastWorkingDate  string                 `json:"lastWorkingDate" gorm:"comment:'最后工作日期'"`
	ResignationReason string                `json:"resignationReason" gorm:"comment:'离职原因'"`
	Status           int                    `json:"status" gorm:"default:0;comment:'状态:0-待审批,1-审批中,2-交接中,3-已完成,4-已拒绝'"`
	Progress         float64                `json:"progress" gorm:"default:0;comment:'进度百分比'"`
	CurrentStep      int                    `json:"currentStep" gorm:"default:0;comment:'当前步骤'"`
	TotalSteps       int                    `json:"totalSteps" gorm:"default:5;comment:'总步骤数'"`
	ApproverId       uint                   `json:"approverId" gorm:"comment:'审批人ID'"`
	ApprovalRemarks  string                 `json:"approvalRemarks" gorm:"comment:'审批备注'"`
	SalarySettlement string                 `json:"salarySettlement" gorm:"comment:'薪资结算情况'"`
	OffboardingTasks []SysOffboardingTask   `json:"offboardingTasks" gorm:"foreignKey:OffboardingId"`
	HandoverItems    []SysHandoverItem      `json:"handoverItems" gorm:"foreignKey:OffboardingId"`
	AssetReturns     []SysAssetReturn       `json:"assetReturns" gorm:"foreignKey:OffboardingId"`
}

// 离职任务表
type SysOffboardingTask struct {
	gorm.Model
	OffboardingId uint   `json:"offboardingId" gorm:"comment:'离职流程ID'"`
	TaskType      string `json:"taskType" gorm:"comment:'任务类型:HANDOVER-工作交接,ASSET-资产归还,PERMISSION-权限回收,FINANCE-财务结算,HR-HR手续'"`
	TaskName      string `json:"taskName" gorm:"comment:'任务名称'"`
	Description   string `json:"description" gorm:"comment:'任务描述'"`
	AssigneeId    uint   `json:"assigneeId" gorm:"comment:'负责人ID'"`
	Status        int    `json:"status" gorm:"default:0;comment:'状态:0-待处理,1-进行中,2-已完成'"`
	DueDate       string `json:"dueDate" gorm:"comment:'截止日期'"`
	CompletedAt   string `json:"completedAt" gorm:"comment:'完成时间'"`
	CompletedBy   uint   `json:"completedBy" gorm:"comment:'完成人ID'"`
	Remarks       string `json:"remarks" gorm:"comment:'备注'"`
}

// 工作交接清单表
type SysHandoverItem struct {
	gorm.Model
	OffboardingId uint   `json:"offboardingId" gorm:"comment:'离职流程ID'"`
	ItemType      string `json:"itemType" gorm:"comment:'交接类型:PROJECT-项目,DOCUMENT-文档,ACCOUNT-账号,OTHER-其他'"`
	ItemName      string `json:"itemName" gorm:"comment:'交接事项名称'"`
	Description   string `json:"description" gorm:"comment:'描述'"`
	ReceiverId    uint   `json:"receiverId" gorm:"comment:'接收人ID'"`
	ReceiverName  string `json:"receiverName" gorm:"comment:'接收人姓名'"`
	Status        int    `json:"status" gorm:"default:0;comment:'状态:0-待交接,1-已交接'"`
	HandoverDate  string `json:"handoverDate" gorm:"comment:'交接日期'"`
	Remarks       string `json:"remarks" gorm:"comment:'备注'"`
}

// 资产归还记录表
type SysAssetReturn struct {
	gorm.Model
	OffboardingId uint   `json:"offboardingId" gorm:"comment:'离职流程ID'"`
	AssetType     string `json:"assetType" gorm:"comment:'资产类型:COMPUTER-电脑,ACCESS_CARD-门禁卡,OTHER-其他'"`
	AssetName     string `json:"assetName" gorm:"comment:'资产名称'"`
	AssetNo       string `json:"assetNo" gorm:"comment:'资产编号'"`
	Status        int    `json:"status" gorm:"default:0;comment:'状态:0-待归还,1-已归还,2-已损坏/丢失'"`
	ReturnDate    string `json:"returnDate" gorm:"comment:'归还日期'"`
	CheckedBy     uint   `json:"checkedBy" gorm:"comment:'验收人ID'"`
	Remarks       string `json:"remarks" gorm:"comment:'备注'"`
}

// 创建离职流程
func (o *SysOffboarding) Create() error {
	err := qmsql.DEFAULTDB.Create(&o).Error
	return err
}

// 更新离职流程
func (o *SysOffboarding) Update() error {
	err := qmsql.DEFAULTDB.Save(&o).Error
	return err
}

// 删除离职流程
func (o *SysOffboarding) Delete() error {
	err := qmsql.DEFAULTDB.Delete(&o).Error
	return err
}

// 根据ID获取离职流程
func (o *SysOffboarding) GetById(id uint) (err error) {
	err = qmsql.DEFAULTDB.Preload("OffboardingTasks").Preload("HandoverItems").Preload("AssetReturns").First(&o, id).Error
	return err
}

// 获取离职流程列表
func (o *SysOffboarding) GetList(page, pageSize int, status int, departmentId string) (list []SysOffboarding, total int, err error) {
	db := qmsql.DEFAULTDB.Model(&SysOffboarding{})
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

// 更新离职任务状态
func (t *SysOffboardingTask) UpdateStatus(status int, completedBy uint, remarks string) error {
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

// 添加交接事项
func (h *SysHandoverItem) Create() error {
	err := qmsql.DEFAULTDB.Create(&h).Error
	return err
}

// 更新交接事项状态
func (h *SysHandoverItem) UpdateStatus(status int, handoverDate string) error {
	err := qmsql.DEFAULTDB.Model(&h).Updates(map[string]interface{}{
		"status":        status,
		"handover_date": handoverDate,
	}).Error
	return err
}

// 添加资产归还记录
func (a *SysAssetReturn) Create() error {
	err := qmsql.DEFAULTDB.Create(&a).Error
	return err
}

// 更新资产归还状态
func (a *SysAssetReturn) UpdateStatus(status int, returnDate string, checkedBy uint, remarks string) error {
	err := qmsql.DEFAULTDB.Model(&a).Updates(map[string]interface{}{
		"status":      status,
		"return_date": returnDate,
		"checked_by":  checkedBy,
		"remarks":     remarks,
	}).Error
	return err
}
