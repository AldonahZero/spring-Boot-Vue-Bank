package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitOnboardingRouter(Router *gin.RouterGroup) {
	OnboardingRouter := Router.Group("onboarding").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		OnboardingRouter.POST("create", api.CreateOnboarding)           // 创建入职流程
		OnboardingRouter.GET("list", api.GetOnboardingList)             // 获取入职流程列表
		OnboardingRouter.GET("detail", api.GetOnboardingDetail)         // 获取入职流程详情
		OnboardingRouter.PUT("update", api.UpdateOnboarding)            // 更新入职流程
		OnboardingRouter.DELETE("delete", api.DeleteOnboarding)         // 删除入职流程
		OnboardingRouter.PUT("task/update", api.UpdateOnboardingTask)   // 更新入职任务状态
		OnboardingRouter.GET("statistics", api.GetOnboardingStatistics) // 获取入职统计
	}
}
