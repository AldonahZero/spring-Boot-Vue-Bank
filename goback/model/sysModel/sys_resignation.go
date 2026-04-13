package sysModel

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
	"time"
)

type ResignationProcess struct {
	gorm.Model
	EmployeeName       string    `json:"employeeName" gorm:"comment:员工姓名"`
	EmployeeNo         string    `json:"employeeNo" gorm:"comment:工号"`
	Department         string    `json:"department" gorm:"comment:部门"`
	Position           string    `json:"position" gorm:"comment:职位"`
	ApplyDate          time.Time `json:"applyDate" gorm:"comment:申请日期"`
	LastWorkingDate    time.Time `json:"lastWorkingDate" gorm:"comment:最后工作日"`
	ResignationReason  string    `json:"resignationReason" gorm:"comment:离职原因"`
	Status             string    `json:"status" gorm:"default:'pending';comment:状态(pending/approved/rejected/processing/completed)"`
	CurrentStep        int       `json:"currentStep" gorm:"default:1;comment:当前步骤"`
	ApproverID         uint      `json:"approverId" gorm:"comment:审批人ID"`
	ApprovedAt         *time.Time `json:"approvedAt" gorm:"comment:审批时间"`
	ApprovalRemark     string    `json:"approvalRemark" gorm:"comment:审批备注"`
	HandoverItems      []HandoverItem `json:"handoverItems"`
	AssetReturns       []AssetReturn `json:"assetReturns"`
}

type HandoverItem struct {
	gorm.Model
	ResignationProcessID uint   `json:"resignationProcessId"`
	ItemName             string `json:"itemName" gorm:"comment:交接项目名称"`
	ItemType             string `json:"itemType" gorm:"comment:类型(project/document/permission)"`
	ItemDescription      string `json:"itemDescription" gorm:"comment:详细描述"`
	ReceiverName         string `json:"receiverName" gorm:"comment:接收人"`
	Status               string `json:"status" gorm:"default:'pending';comment:状态(pending/completed)"`
	CompletedAt          *time.Time `json:"completedAt" gorm:"comment:完成时间"`
	Remark               string `json:"remark" gorm:"comment:备注"`
}

type AssetReturn struct {
	gorm.Model
	ResignationProcessID uint   `json:"resignationProcessId"`
	AssetName            string `json:"assetName" gorm:"comment:资产名称"`
	AssetType            string `json:"assetType" gorm:"comment:资产类型(computer/office/card/other)"`
	AssetNo              string `json:"assetNo" gorm:"comment:资产编号"`
	Status               string `json:"status" gorm:"default:'pending';comment:状态(pending/returned/damaged/lost)"`
	ReturnedAt           *time.Time `json:"returnedAt" gorm:"comment:归还时间"`
	ReceiverName         string `json:"receiverName" gorm:"comment:接收人"`
	Remark               string `json:"remark" gorm:"comment:备注"`
}

type SalarySettlement struct {
	gorm.Model
	ResignationProcessID uint      `json:"resignationProcessId"`
	BaseSalary           float64   `json:"baseSalary" gorm:"comment:基本工资"`
	WorkDays             int       `json:"workDays" gorm:"comment:工作天数"`
	PayableAmount       float64   `json:"payableAmount" gorm:"comment:应付金额"`
	Bonus               float64   `json:"bonus" gorm:"comment:奖金"`
	Deduction           float64   `json:"deduction" gorm:"comment:扣款"`
	FinalAmount         float64   `json:"finalAmount" gorm:"comment:最终金额"`
	SettlementDate      time.Time `json:"settlementDate" gorm:"comment:结算日期"`
	Status              string    `json:"status" gorm:"default:'pending';comment:状态(pending/paid)"`
}

type ResignationCertificate struct {
	gorm.Model
	ResignationProcessID uint      `json:"resignationProcessId"`
	CertificateNo        string    `json:"certificateNo" gorm:"comment:证明编号"`
	EmployeeName         string    `json:"employeeName" gorm:"comment:员工姓名"`
	EmployeeNo           string    `json:"employeeNo" gorm:"comment:工号"`
	Department           string    `json:"department" gorm:"comment:部门"`
	Position             string    `json:"position" gorm:"comment:职位"`
	EntryDate            time.Time `json:"entryDate" gorm:"comment:入职日期"`
	LeaveDate            time.Time `json:"leaveDate" gorm:"comment:离职日期"`
	Reason               string    `json:"reason" gorm:"comment:离职原因"`
	IssueDate            time.Time `json:"issueDate" gorm:"comment:开具日期"`
	GeneratedAt          time.Time `json:"generatedAt" gorm:"comment:生成时间"`
}

func (rp *ResignationProcess) Create() error {
	return qmsql.DEFAULTDB.Create(&rp).Error
}

func (rp *ResignationProcess) GetByID() (error, *ResignationProcess) {
	var process ResignationProcess
	err := qmsql.DEFAULTDB.Preload("HandoverItems").Preload("AssetReturns").Where("id = ?", rp.ID).First(&process).Error
	return err, &process
}

func (rp *ResignationProcess) GetList(info modelInterface.PageInfo) (error, interface{}, int) {
	err, db, total := servers.PagingServer(rp, info)
	if err != nil {
		return err, nil, 0
	}
	var list []ResignationProcess
	err = db.Preload("HandoverItems").Preload("AssetReturns").Find(&list).Error
	return err, list, total
}

func (rp *ResignationProcess) Update() error {
	return qmsql.DEFAULTDB.Save(&rp).Error
}

func (rp *ResignationProcess) Delete() error {
	return qmsql.DEFAULTDB.Where("id = ?", rp.ID).Delete(&rp).Error
}

func (hi *HandoverItem) Create() error {
	return qmsql.DEFAULTDB.Create(&hi).Error
}

func (hi *HandoverItem) Update() error {
	return qmsql.DEFAULTDB.Save(&hi).Error
}

func (ar *AssetReturn) Create() error {
	return qmsql.DEFAULTDB.Create(&ar).Error
}

func (ar *AssetReturn) Update() error {
	return qmsql.DEFAULTDB.Save(&ar).Error
}

func (ss *SalarySettlement) Create() error {
	return qmsql.DEFAULTDB.Create(&ss).Error
}

func (ss *SalarySettlement) Update() error {
	return qmsql.DEFAULTDB.Save(&ss).Error
}

func (ss *SalarySettlement) Calculate() {
	ss.FinalAmount = ss.BaseSalary/30.0*float64(ss.WorkDays) + ss.Bonus - ss.Deduction
}

func (rc *ResignationCertificate) Create() error {
	return qmsql.DEFAULTDB.Create(&rc).Error
}

func GenerateCertificateNo() string {
	return "CERT" + time.Now().Format("20060102") + randomString(4)
}
