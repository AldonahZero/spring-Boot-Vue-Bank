package registTable

import (
	"gin-vue-admin/model/dbModel"
	"gin-vue-admin/model/sysModel"
	"github.com/jinzhu/gorm"
)

//注册数据库表专用
func RegistTable(db *gorm.DB) {
	db.AutoMigrate(sysModel.SysUser{},
		sysModel.SysAuthority{},
		sysModel.SysMenu{},
		sysModel.SysApi{},
		sysModel.SysBaseMenu{},
		sysModel.JwtBlacklist{},
		dbModel.ExaFileUploadAndDownload{},
		sysModel.SysWorkflow{},
		sysModel.SysWorkflowStepInfo{},
		sysModel.OnboardingProcess{},
		sysModel.OnboardingChecklist{},
		sysModel.OnboardingTask{},
		sysModel.ResignationProcess{},
		sysModel.HandoverItem{},
		sysModel.AssetReturn{},
		sysModel.SalarySettlement{},
		sysModel.ResignationCertificate{},
	)
}
