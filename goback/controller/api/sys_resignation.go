package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateResignationProcess(c *gin.Context) {
	var process sysModel.ResignationProcess
	_ = c.ShouldBind(&process)
	process.Status = "pending"
	process.CurrentStep = 1
	err := process.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建离职流程失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "创建离职流程成功", gin.H{"process": process})
}

func GetResignationProcessList(c *gin.Context) {
	var process sysModel.ResignationProcess
	err, list, total := process.GetList(servers.GetPageInfo(c))
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取离职流程列表失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "获取成功", gin.H{
		"list":  list,
		"total": total,
	})
}

func GetResignationProcessByID(c *gin.Context) {
	var process sysModel.ResignationProcess
	_ = c.ShouldBind(&process)
	err, result := process.GetByID()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取离职流程详情失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "获取成功", gin.H{"process": result})
}

func UpdateResignationProcess(c *gin.Context) {
	var process sysModel.ResignationProcess
	_ = c.ShouldBind(&process)
	err := process.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新离职流程失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "更新成功", gin.H{})
}

func DeleteResignationProcess(c *gin.Context) {
	var process sysModel.ResignationProcess
	_ = c.ShouldBind(&process)
	err := process.Delete()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除离职流程失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "删除成功", gin.H{})
}

func ApproveResignation(c *gin.Context) {
	var process sysModel.ResignationProcess
	_ = c.ShouldBind(&process)
	err, result := process.GetByID()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取离职流程失败：%v", err), gin.H{})
		return
	}
	now := time.Now()
	result.ApprovedAt = &now
	result.Status = "approved"
	result.CurrentStep = 2
	err = result.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("审批失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "审批通过", gin.H{"process": result})
}

func RejectResignation(c *gin.Context) {
	var process sysModel.ResignationProcess
	_ = c.ShouldBind(&process)
	err, result := process.GetByID()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取离职流程失败：%v", err), gin.H{})
		return
	}
	result.Status = "rejected"
	err = result.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("拒绝失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "已拒绝", gin.H{"process": result})
}

func AddHandoverItem(c *gin.Context) {
	var item sysModel.HandoverItem
	_ = c.ShouldBind(&item)
	err := item.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加交接项失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "添加成功", gin.H{"item": item})
}

func CompleteHandoverItem(c *gin.Context) {
	var item sysModel.HandoverItem
	_ = c.ShouldBind(&item)
	now := time.Now()
	item.CompletedAt = &now
	item.Status = "completed"
	err := item.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("完成交接项失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "交接项完成", gin.H{})
}

func AddAssetReturn(c *gin.Context) {
	var asset sysModel.AssetReturn
	_ = c.ShouldBind(&asset)
	err := asset.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加资产归还失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "添加成功", gin.H{"asset": asset})
}

func CompleteAssetReturn(c *gin.Context) {
	var asset sysModel.AssetReturn
	_ = c.ShouldBind(&asset)
	now := time.Now()
	asset.ReturnedAt = &now
	asset.Status = "returned"
	err := asset.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("完成资产归还失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "资产归还完成", gin.H{})
}

func CalculateSalarySettlement(c *gin.Context) {
	var settlement sysModel.SalarySettlement
	_ = c.ShouldBind(&settlement)
	settlement.Calculate()
	settlement.SettlementDate = time.Now()
	settlement.Status = "pending"
	err := settlement.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("计算薪资结算失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "薪资结算计算完成", gin.H{"settlement": settlement})
}

func GenerateResignationCertificate(c *gin.Context) {
	var process sysModel.ResignationProcess
	_ = c.ShouldBind(&process)
	err, result := process.GetByID()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取离职流程失败：%v", err), gin.H{})
		return
	}
	certificate := sysModel.ResignationCertificate{
		ResignationProcessID: result.ID,
		CertificateNo:        sysModel.GenerateCertificateNo(),
		EmployeeName:         result.EmployeeName,
		EmployeeNo:           result.EmployeeNo,
		Department:           result.Department,
		Position:             result.Position,
		EntryDate:            result.CreatedAt,
		LeaveDate:            result.LastWorkingDate,
		Reason:               result.ResignationReason,
		IssueDate:            time.Now(),
		GeneratedAt:          time.Now(),
	}
	err = certificate.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("生成离职证明失败：%v", err), gin.H{})
		return
	}
	result.Status = "completed"
	result.CurrentStep = 5
	result.Update()
	servers.ReportFormat(c, true, "离职证明生成成功", gin.H{"certificate": certificate})
}

func AdvanceResignationStep(c *gin.Context) {
	var process sysModel.ResignationProcess
	_ = c.ShouldBind(&process)
	err, result := process.GetByID()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取离职流程失败：%v", err), gin.H{})
		return
	}
	result.CurrentStep++
	if result.CurrentStep > 5 {
		result.Status = "completed"
		result.CurrentStep = 5
	} else {
		result.Status = "processing"
	}
	err = result.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新步骤失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "步骤更新成功", gin.H{"process": result})
}
