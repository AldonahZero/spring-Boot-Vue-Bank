package router

import (
	"gin-vue-admin/controller/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitResignationRouter(Router *gin.RouterGroup) {
	ResignationRouter := Router.Group("resignation").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ResignationRouter.POST("create", api.CreateResignationProcess)
		ResignationRouter.POST("list", api.GetResignationProcessList)
		ResignationRouter.POST("detail", api.GetResignationProcessByID)
		ResignationRouter.POST("update", api.UpdateResignationProcess)
		ResignationRouter.POST("delete", api.DeleteResignationProcess)
		ResignationRouter.POST("approve", api.ApproveResignation)
		ResignationRouter.POST("reject", api.RejectResignation)
		ResignationRouter.POST("addHandoverItem", api.AddHandoverItem)
		ResignationRouter.POST("completeHandoverItem", api.CompleteHandoverItem)
		ResignationRouter.POST("addAssetReturn", api.AddAssetReturn)
		ResignationRouter.POST("completeAssetReturn", api.CompleteAssetReturn)
		ResignationRouter.POST("calculateSalary", api.CalculateSalarySettlement)
		ResignationRouter.POST("generateCertificate", api.GenerateResignationCertificate)
		ResignationRouter.POST("advanceStep", api.AdvanceResignationStep)
	}
}
