package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateOnboardingProcess(c *gin.Context) {
	var process sysModel.OnboardingProcess
	_ = c.ShouldBind(&process)
	process.EmployeeNo = sysModel.GenerateEmployeeNo()
	process.Email = sysModel.GenerateEmail(process.EmployeeName)
	process.Status = "pending"
	process.CurrentStep = 1
	err := process.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建入职流程失败：%v", err), gin.H{})
		return
	}
	itTask := sysModel.OnboardingTask{
		OnboardingProcessID: process.ID,
		Department:          "it",
		TaskType:            "设备准备",
		TaskContent:         "准备电脑、开通账号",
		Status:              "pending",
		DueDate:             process.OnboardingDate.AddDate(0, 0, -1),
	}
	itTask.Create()
	adminTask := sysModel.OnboardingTask{
		OnboardingProcessID: process.ID,
		Department:          "admin",
		TaskType:            "工位准备",
		TaskContent:         "准备工位、门禁卡",
		Status:              "pending",
		DueDate:             process.OnboardingDate.AddDate(0, 0, -1),
	}
	adminTask.Create()
	mentorTask := sysModel.OnboardingTask{
		OnboardingProcessID: process.ID,
		Department:          "mentor",
		TaskType:            "入职引导",
		TaskContent:         "发送欢迎邮件、安排入职引导",
		Status:              "pending",
		DueDate:             process.OnboardingDate,
	}
	mentorTask.Create()
	checklistItems := []sysModel.OnboardingChecklist{
		{OnboardingProcessID: process.ID, ItemName: "电脑配置完成", ItemCategory: "it"},
		{OnboardingProcessID: process.ID, ItemName: "账号开通完成", ItemCategory: "it"},
		{OnboardingProcessID: process.ID, ItemName: "工位分配完成", ItemCategory: "admin"},
		{OnboardingProcessID: process.ID, ItemName: "门禁卡发放完成", ItemCategory: "admin"},
		{OnboardingProcessID: process.ID, ItemName: "欢迎邮件发送", ItemCategory: "mentor"},
		{OnboardingProcessID: process.ID, ItemName: "入职引导完成", ItemCategory: "mentor"},
	}
	for _, item := range checklistItems {
		item.Create()
	}
	servers.ReportFormat(c, true, "创建入职流程成功", gin.H{"process": process})
}

func GetOnboardingProcessList(c *gin.Context) {
	var process sysModel.OnboardingProcess
	err, list, total := process.GetList(servers.GetPageInfo(c))
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取入职流程列表失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "获取成功", gin.H{
		"list":  list,
		"total": total,
	})
}

func GetOnboardingProcessByID(c *gin.Context) {
	var process sysModel.OnboardingProcess
	_ = c.ShouldBind(&process)
	err, result := process.GetByID()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取入职流程详情失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "获取成功", gin.H{"process": result})
}

func UpdateOnboardingProcess(c *gin.Context) {
	var process sysModel.OnboardingProcess
	_ = c.ShouldBind(&process)
	err := process.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新入职流程失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "更新成功", gin.H{})
}

func DeleteOnboardingProcess(c *gin.Context) {
	var process sysModel.OnboardingProcess
	_ = c.ShouldBind(&process)
	err := process.Delete()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除入职流程失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "删除成功", gin.H{})
}

func CompleteOnboardingTask(c *gin.Context) {
	var task sysModel.OnboardingTask
	_ = c.ShouldBind(&task)
	now := time.Now()
	task.CompletedAt = &now
	task.Status = "completed"
	err := task.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("完成任务失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "任务完成", gin.H{})
}

func CompleteOnboardingChecklist(c *gin.Context) {
	var checklist sysModel.OnboardingChecklist
	_ = c.ShouldBind(&checklist)
	now := time.Now()
	checklist.CompletedAt = &now
	checklist.IsCompleted = true
	err := checklist.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("完成检查项失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "检查项完成", gin.H{})
}

func GetOnboardingTasks(c *gin.Context) {
	var req struct {
		ProcessID uint `json:"processId"`
	}
	_ = c.ShouldBind(&req)
	var task sysModel.OnboardingTask
	err, tasks := task.GetByProcessID(req.ProcessID)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取任务列表失败：%v", err), gin.H{})
		return
	}
	servers.ReportFormat(c, true, "获取成功", gin.H{"tasks": tasks})
}

func AdvanceOnboardingStep(c *gin.Context) {
	var process sysModel.OnboardingProcess
	_ = c.ShouldBind(&process)
	err, result := process.GetByID()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取入职流程失败：%v", err), gin.H{})
		return
	}
	result.CurrentStep++
	if result.CurrentStep > 4 {
		result.Status = "completed"
		result.CurrentStep = 4
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
