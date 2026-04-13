package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Tags Offboarding
// @Summary 创建离职流程
// @Produce application/json
// @Param data body sysModel.SysOffboarding true "创建离职流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /offboarding/create [post]
func CreateOffboarding(c *gin.Context) {
	var offboarding sysModel.SysOffboarding
	_ = c.ShouldBind(&offboarding)

	offboarding.Status = 0
	offboarding.Progress = 0
	offboarding.CurrentStep = 0
	offboarding.TotalSteps = 5

	err := offboarding.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
		return
	}

	// 自动创建离职任务
	tasks := []sysModel.SysOffboardingTask{
		{
			OffboardingId: offboarding.ID,
			TaskType:      "HR",
			TaskName:      "离职申请审批",
			Description:   "审批员工离职申请",
			Status:        0,
		},
		{
			OffboardingId: offboarding.ID,
			TaskType:      "HANDOVER",
			TaskName:      "工作交接",
			Description:   "项目、文档、工作交接",
			Status:        0,
		},
		{
			OffboardingId: offboarding.ID,
			TaskType:      "ASSET",
			TaskName:      "资产归还",
			Description:   "归还电脑、办公用品等",
			Status:        0,
		},
		{
			OffboardingId: offboarding.ID,
			TaskType:      "FINANCE",
			TaskName:      "薪资结算",
			Description:   "核算未结薪资、补偿金等",
			Status:        0,
		},
		{
			OffboardingId: offboarding.ID,
			TaskType:      "HR",
			TaskName:      "离职手续办理",
			Description:   "开具离职证明、档案转移",
			Status:        0,
		},
	}

	for _, task := range tasks {
		qmsql.DEFAULTDB.Create(&task)
	}

	servers.ReportFormat(c, true, "创建成功", gin.H{
		"offboarding": offboarding,
	})
}

// @Tags Offboarding
// @Summary 获取离职流程列表
// @Produce application/json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param status query int false "状态"
// @Param departmentId query string false "部门ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /offboarding/list [get]
func GetOffboardingList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	status, _ := strconv.Atoi(c.DefaultQuery("status", "-1"))
	departmentId := c.Query("departmentId")

	var offboarding sysModel.SysOffboarding
	list, total, err := offboarding.GetList(page, pageSize, status, departmentId)
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

// @Tags Offboarding
// @Summary 获取离职流程详情
// @Produce application/json
// @Param id query int true "离职流程ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /offboarding/detail [get]
func GetOffboardingDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	var offboarding sysModel.SysOffboarding
	err := offboarding.GetById(uint(id))
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "获取成功", gin.H{
		"offboarding": offboarding,
	})
}

// @Tags Offboarding
// @Summary 更新离职流程
// @Produce application/json
// @Param data body sysModel.SysOffboarding true "更新离职流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /offboarding/update [put]
func UpdateOffboarding(c *gin.Context) {
	var offboarding sysModel.SysOffboarding
	_ = c.ShouldBind(&offboarding)

	err := offboarding.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "更新成功", gin.H{})
}

// @Tags Offboarding
// @Summary 删除离职流程
// @Produce application/json
// @Param id query int true "离职流程ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /offboarding/delete [delete]
func DeleteOffboarding(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	offboarding := sysModel.SysOffboarding{}
	offboarding.ID = uint(id)
	err := offboarding.Delete()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败：%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "删除成功", gin.H{})
}

// @Tags Offboarding
// @Summary 审批离职申请
// @Produce application/json
// @Param id query int true "离职流程ID"
// @Param approved query bool true "是否批准"
// @Param remarks query string false "审批备注"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"审批成功"}"
// @Router /offboarding/approve [put]
func ApproveOffboarding(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	approved, _ := strconv.ParseBool(c.Query("approved"))
	remarks := c.Query("remarks")
	userId, _ := c.Get("userId")

	var offboarding sysModel.SysOffboarding
	err := qmsql.DEFAULTDB.First(&offboarding, id).Error
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
		return
	}

	if approved {
		offboarding.Status = 2 // 进入交接阶段
	} else {
		offboarding.Status = 4 // 已拒绝
	}
	offboarding.ApproverId = userId.(uint)
	offboarding.ApprovalRemarks = remarks

	err = offboarding.Update()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("审批失败：%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "审批成功", gin.H{})
}

// @Tags Offboarding
// @Summary 更新离职任务状态
// @Produce application/json
// @Param taskId query int true "任务ID"
// @Param status query int true "状态"
// @Param remarks query string false "备注"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /offboarding/task/update [put]
func UpdateOffboardingTask(c *gin.Context) {
	taskId, _ := strconv.Atoi(c.Query("taskId"))
	status, _ := strconv.Atoi(c.Query("status"))
	remarks := c.Query("remarks")
	userId, _ := c.Get("userId")

	var task sysModel.SysOffboardingTask
	task.ID = uint(taskId)
	err := task.UpdateStatus(status, userId.(uint), remarks)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return
	}

	// 更新离职流程进度
	var offboarding sysModel.SysOffboarding
	qmsql.DEFAULTDB.First(&offboarding, task.OffboardingId)

	var completedTasks int64
	qmsql.DEFAULTDB.Model(&sysModel.SysOffboardingTask{}).Where("offboarding_id = ? AND status = 2", task.OffboardingId).Count(&completedTasks)

	offboarding.Progress = float64(completedTasks) / float64(offboarding.TotalSteps) * 100
	offboarding.CurrentStep = int(completedTasks)

	if completedTasks == int64(offboarding.TotalSteps) {
		offboarding.Status = 3 // 已完成
	}

	offboarding.Update()

	servers.ReportFormat(c, true, "更新成功", gin.H{
		"progress": offboarding.Progress,
	})
}

// @Tags Offboarding
// @Summary 添加交接事项
// @Produce application/json
// @Param data body sysModel.SysHandoverItem true "交接事项"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /offboarding/handover/add [post]
func AddHandoverItem(c *gin.Context) {
	var item sysModel.SysHandoverItem
	_ = c.ShouldBind(&item)

	err := item.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败：%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "添加成功", gin.H{})
}

// @Tags Offboarding
// @Summary 更新交接事项状态
// @Produce application/json
// @Param id query int true "交接事项ID"
// @Param status query int true "状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /offboarding/handover/update [put]
func UpdateHandoverItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	status, _ := strconv.Atoi(c.Query("status"))

	var item sysModel.SysHandoverItem
	item.ID = uint(id)
	handoverDate := ""
	if status == 1 {
		handoverDate = "now()"
	}
	err := item.UpdateStatus(status, handoverDate)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "更新成功", gin.H{})
}

// @Tags Offboarding
// @Summary 添加资产归还记录
// @Produce application/json
// @Param data body sysModel.SysAssetReturn true "资产归还记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /offboarding/asset/add [post]
func AddAssetReturn(c *gin.Context) {
	var asset sysModel.SysAssetReturn
	_ = c.ShouldBind(&asset)

	err := asset.Create()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("添加失败：%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "添加成功", gin.H{})
}

// @Tags Offboarding
// @Summary 更新资产归还状态
// @Produce application/json
// @Param id query int true "资产ID"
// @Param status query int true "状态"
// @Param remarks query string false "备注"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /offboarding/asset/update [put]
func UpdateAssetReturn(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	status, _ := strconv.Atoi(c.Query("status"))
	remarks := c.Query("remarks")
	userId, _ := c.Get("userId")

	var asset sysModel.SysAssetReturn
	asset.ID = uint(id)
	returnDate := ""
	if status == 1 {
		returnDate = "now()"
	}
	err := asset.UpdateStatus(status, returnDate, userId.(uint), remarks)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("更新失败：%v", err), gin.H{})
		return
	}

	servers.ReportFormat(c, true, "更新成功", gin.H{})
}

// @Tags Offboarding
// @Summary 生成离职证明
// @Produce application/json
// @Param id query int true "离职流程ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"生成成功"}"
// @Router /offboarding/certificate [get]
func GenerateCertificate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	var offboarding sysModel.SysOffboarding
	err := qmsql.DEFAULTDB.First(&offboarding, id).Error
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取失败：%v", err), gin.H{})
		return
	}

	// 生成离职证明内容
	certificate := fmt.Sprintf("离职证明\n\n兹证明 %s（工号：%s）于本公司 %s 部门担任 %s 职位，\n已于 %s 正式离职，所有离职手续已办理完毕。\n\n特此证明。\n\n公司盖章\n日期：%s",
		offboarding.EmployeeName,
		offboarding.EmployeeNo,
		offboarding.DepartmentId,
		offboarding.Position,
		offboarding.LastWorkingDate,
		"now()",
	)

	servers.ReportFormat(c, true, "生成成功", gin.H{
		"certificate": certificate,
	})
}

// @Tags Offboarding
// @Summary 获取离职统计
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /offboarding/statistics [get]
func GetOffboardingStatistics(c *gin.Context) {
	var total, pending, inProgress, completed, rejected int64

	qmsql.DEFAULTDB.Model(&sysModel.SysOffboarding{}).Count(&total)
	qmsql.DEFAULTDB.Model(&sysModel.SysOffboarding{}).Where("status = 0").Count(&pending)
	qmsql.DEFAULTDB.Model(&sysModel.SysOffboarding{}).Where("status = 1 OR status = 2").Count(&inProgress)
	qmsql.DEFAULTDB.Model(&sysModel.SysOffboarding{}).Where("status = 3").Count(&completed)
	qmsql.DEFAULTDB.Model(&sysModel.SysOffboarding{}).Where("status = 4").Count(&rejected)

	servers.ReportFormat(c, true, "获取成功", gin.H{
		"total":      total,
		"pending":    pending,
		"inProgress": inProgress,
		"completed":  completed,
		"rejected":   rejected,
	})
}
