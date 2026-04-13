package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitOffboardingRouter(Router *gin.RouterGroup) {
	OffboardingRouter := Router.Group("offboarding").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		OffboardingRouter.POST("create", api.CreateOffboarding)             // 创建离职流程
		OffboardingRouter.GET("list", api.GetOffboardingList)               // 获取离职流程列表
		OffboardingRouter.GET("detail", api.GetOffboardingDetail)           // 获取离职流程详情
		OffboardingRouter.PUT("update", api.UpdateOffboarding)              // 更新离职流程
		OffboardingRouter.DELETE("delete", api.DeleteOffboarding)           // 删除离职流程
		OffboardingRouter.PUT("approve", api.ApproveOffboarding)            // 审批离职申请
		OffboardingRouter.PUT("task/update", api.UpdateOffboardingTask)     // 更新离职任务状态
		OffboardingRouter.POST("handover/add", api.AddHandoverItem)         // 添加交接事项
		OffboardingRouter.PUT("handover/update", api.UpdateHandoverItem)    // 更新交接事项状态
		OffboardingRouter.POST("asset/add", api.AddAssetReturn)             // 添加资产归还记录
		OffboardingRouter.PUT("asset/update", api.UpdateAssetReturn)        // 更新资产归还状态
		OffboardingRouter.GET("certificate", api.GenerateCertificate)       // 生成离职证明
		OffboardingRouter.GET("statistics", api.GetOffboardingStatistics)   // 获取离职统计
	}
}
