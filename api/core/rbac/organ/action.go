package organ

import (
	"api/base"

	"github.com/gin-gonic/gin"
)

func OrganTree(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: nil})
}
func OrganPage(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: base.PageData{
		Page:  1,
		Size:  10,
		Total: 10,
		Data: []interface{}{
			gin.H{"id": 1, "name": "组织1"},
			gin.H{"id": 2, "name": "组织2"},
		},
	}})
}
func OrganList(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: nil})
}
func OrganInfo(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: nil})
}
func OrganSave(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: nil})
}
func OrganDel(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: nil})
}
func OrganDisable(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: nil})
}
func OrganEnable(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: nil})
}

func OrganSaveParent(ctx *gin.Context) {
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: nil})
}
