package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Tags Onboarding
// @Summary 创建入职流程
// @Produce application/json
// @Param data body sysModel.SysOnboarding true "创建入职流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /onboarding/create [post]
func CreateOnboarding(c *gin.Context) {
	var onboarding sysModel.SysOnboarding
	_ = c.ShouldBind(&onboarding)
	
	// 自动生成工号和邮箱
	onboarding.EmployeeNo = sysModel.GenerateEmployeeNo()
	onboarding.Email = sysModel.GenerateEmail(onboarding.EmployeeNo)
	onboarding.Status = 0
	onboarding.Progress = 0
	onboarding.CurrentStep = 0
	onboarding.TotalSteps = 5
	
	err := onboarding.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
		return
	}
	
	// 自动创建入职任务
	tasks := []sysModel.SysOnboardingTask{
		{
			OnboardingId: onboarding.ID,
			TaskType:     "HR",
			TaskName:     "入职资料审核",
			Description:  "审核新员工入职资料完整性",
			Status:       0,
		},
		{
			OnboardingId: onboarding.ID,
			TaskType:     "IT",
			TaskName:     "IT设备准备",
			Description:  "准备电脑、开通账号权限",
			Status:       0,
		},
		{
			OnboardingId: onboarding.ID,
			TaskType:     "ADMIN",
			TaskName:     "行政准备",
			Description:  "安排工位、准备门禁卡",
			Status:       0,
		},
		{
			OnboardingId: onboarding.ID,
			TaskType:     "MENTOR",
			TaskName:     "导师准备",
			Description:  "发送欢迎邮件、准备入职引导",
			Status:       0,
		},
		{
			OnboardingId: onboarding.ID,
			TaskType:     "HR",
			TaskName:     "入职完成确认",
			Description:  "确认所有入职事项已完成",
			Status:       0,
		},
	}
	
	for _, task := range tasks {
		qmsql.DEFAULTDB.Create(&task)
	}
	
	servers.ReportFormat(c, true, "创建成功", gin.H{
		"onboarding": onboarding,
		"employeeNo": onboarding.EmployeeNo,
		"email":      onboarding.Email,
	})
}

// @Tags Onboarding
// @Summary 获取入职流程列表
// @Produce application/json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param status query int false "状态"
// @Param departmentId query string false "部门ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /onboarding/list [get]
func GetOnboardingList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	status, _ := strconv.Atoi(c.DefaultQuery("status", "-1"))
	departmentId := c.Query("departmentId")
	
	var onboarding sysModel.SysOnboarding
	list, total, err := onboarding.GetList(page, pageSize, status, departmentId)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
		return
	}
	
	servers.ReportFormat(c, true, "获取成功", gin.H{
		"list":     list,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// @Tags Onboarding
// @Summary 获取入职流程详情
// @Produce application/json
// @Param id query int true "入职流程ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /onboarding/detail [get]
func GetOnboardingDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	
	var onboarding sysModel.SysOnboarding
	err := onboarding.GetById(uint(id))
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
		return
	}
	
	servers.ReportFormat(c, true, "获取成功", gin.H{
		"onboarding": onboarding,
	})
}

// @Tags Onboarding
// @Summary 更新入职流程
// @Produce application/json
// @Param data body sysModel.SysOnboarding true "更新入职流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /onboarding/update [put]
func UpdateOnboarding(c *gin.Context) {
	var onboarding sysModel.SysOnboarding
	_ = c.ShouldBind(&onboarding)
	
	err := onboarding.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return
	}
	
	servers.ReportFormat(c, true, "更新成功", gin.H{})
}

// @Tags Onboarding
// @Summary 删除入职流程
// @Produce application/json
// @Param id query int true "入职流程ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /onboarding/delete [delete]
func DeleteOnboarding(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	
	onboarding := sysModel.SysOnboarding{}
	onboarding.ID = uint(id)
	err := onboarding.Delete()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
		return
	}
	
	servers.ReportFormat(c, true, "删除成功", gin.H{})
}

// @Tags Onboarding
// @Summary 更新入职任务状态
// @Produce application/json
// @Param taskId query int true "任务ID"
// @Param status query int true "状态"
// @Param remarks query string false "备注"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /onboarding/task/update [put]
func UpdateOnboardingTask(c *gin.Context) {
	taskId, _ := strconv.Atoi(c.Query("taskId"))
	status, _ := strconv.Atoi(c.Query("status"))
	remarks := c.Query("remarks")
	userId, _ := c.Get("userId")
	
	var task sysModel.SysOnboardingTask
	task.ID = uint(taskId)
	err := task.UpdateStatus(status, userId.(uint), remarks)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return
	}
	
	// 更新入职流程进度
	var onboarding sysModel.SysOnboarding
	qmsql.DEFAULTDB.First(&onboarding, task.OnboardingId)
	
	var completedTasks int64
	qmsql.DEFAULTDB.Model(&sysModel.SysOnboardingTask{}).Where("onboarding_id = ? AND status = 2", task.OnboardingId).Count(&completedTasks)
	
	onboarding.Progress = float64(completedTasks) / float64(onboarding.TotalSteps) * 100
	onboarding.CurrentStep = int(completedTasks)
	
	if completedTasks == int64(onboarding.TotalSteps) {
		onboarding.Status = 2 // 已完成
	} else if completedTasks > 0 {
		onboarding.Status = 1 // 进行中
	}
	
	onboarding.Update()
	
	servers.ReportFormat(c, true, "更新成功", gin.H{
		"progress": onboarding.Progress,
	})
}

// @Tags Onboarding
// @Summary 获取入职统计
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /onboarding/statistics [get]
func GetOnboardingStatistics(c *gin.Context) {
	var total, pending, inProgress, completed int64
	
	qmsql.DEFAULTDB.Model(&sysModel.SysOnboarding{}).Count(&total)
	qmsql.DEFAULTDB.Model(&sysModel.SysOnboarding{}).Where("status = 0").Count(&pending)
	qmsql.DEFAULTDB.Model(&sysModel.SysOnboarding{}).Where("status = 1").Count(&inProgress)
	qmsql.DEFAULTDB.Model(&sysModel.SysOnboarding{}).Where("status = 2").Count(&completed)
	
	servers.ReportFormat(c, true, "获取成功", gin.H{
		"total":      total,
		"pending":    pending,
		"inProgress": inProgress,
		"completed":  completed,
	})
}
