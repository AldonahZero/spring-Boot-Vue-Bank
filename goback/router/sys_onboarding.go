package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitOnboardingRouter(Router *gin.RouterGroup) {
	OnboardingRouter := Router.Group("onboarding").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		OnboardingRouter.POST("create", api.CreateOnboardingProcess)
		OnboardingRouter.POST("list", api.GetOnboardingProcessList)
		OnboardingRouter.POST("detail", api.GetOnboardingProcessByID)
		OnboardingRouter.POST("update", api.UpdateOnboardingProcess)
		OnboardingRouter.POST("delete", api.DeleteOnboardingProcess)
		OnboardingRouter.POST("completeTask", api.CompleteOnboardingTask)
		OnboardingRouter.POST("completeChecklist", api.CompleteOnboardingChecklist)
		OnboardingRouter.POST("getTasks", api.GetOnboardingTasks)
		OnboardingRouter.POST("advanceStep", api.AdvanceOnboardingStep)
	}
}
