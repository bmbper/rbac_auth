package organ

import (
	"api/base"
	"api/db"
	"errors"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

func OrganTree(ctx *gin.Context) {
	var req BmbpRbacOrgan
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var msg string
		contentType := ctx.Request.Header.Get("Content-Type")
		mediaType := strings.Split(contentType, ";")[0]
		if !strings.HasPrefix(mediaType, "application/json") {
			msg = "请传递json格式的参数,当前传递的是: " + mediaType
		} else {
			msg = err.Error()
		}
		ctx.JSON(200, base.RespFail("组织机构参数接收失败: "+msg))
		return
	}
	query := db.DbUtil.Model(&BmbpRbacOrgan{})
	if req.NodeTitle != "" {
		query.Where("node_title LIKE ?", "%"+req.NodeTitle+"%")
	}
	var organs []BmbpRbacOrgan
	if err := query.Find(&organs).Error; err != nil {
		ctx.JSON(200, base.RespFail("查询组织机构树失败: "+err.Error()))
		return
	}
	tree := base.BuildTree(organs)
	ctx.JSON(200, base.RespOk(tree))
}
func OrganPage(ctx *gin.Context) {
	var req base.PageReq[BmbpRbacOrgan]
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var msg string
		contentType := ctx.Request.Header.Get("Content-Type")
		mediaType := strings.Split(contentType, ";")[0]
		if !strings.HasPrefix(mediaType, "application/json") {
			msg = "请传递json格式的参数,当前传递的是: " + mediaType
		} else {
			msg = err.Error()
		}
		ctx.JSON(200, base.RespFail("组织机构参数接收失败: "+msg))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}
	var organPageData = base.PageData{
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
		Data:  []BmbpRbacOrgan{},
	}
	query := db.DbUtil.Model(&BmbpRbacOrgan{})
	if req.Params != nil {
		params := req.Params
		if params.NodeTitle != "" {
			query.Where("NodeTitle LIKE ?", "%"+params.NodeTitle+"%")
		}
	}
	// 计算总记录数
	if err := query.Count(&organPageData.Total).Error; err != nil {
		ctx.JSON(200, base.RespFail("查询组织机构列表失败: "+err.Error()))
		return
	}
	if err := query.Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&organPageData.Data).Error; err != nil {
		ctx.JSON(200, base.RespFail("查询组织机构列表失败: "+err.Error()))
		return
	}
	ctx.JSON(200, base.Resp{Code: "0", Msg: "success", Data: organPageData})
}
func OrganList(ctx *gin.Context) {
	var req BmbpRbacOrgan
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var msg string
		contentType := ctx.Request.Header.Get("Content-Type")
		mediaType := strings.Split(contentType, ";")[0]
		if !strings.HasPrefix(mediaType, "application/json") {
			msg = "请传递json格式的参数,当前传递的是: " + mediaType
		} else {
			msg = err.Error()
		}
		ctx.JSON(200, base.RespFail("组织机构参数接收失败: "+msg))
		return
	}
	query := db.DbUtil.Model(&BmbpRbacOrgan{})
	if req.NodeTitle != "" {
		query.Where("node_title LIKE ?", "%"+req.NodeTitle+"%")
	}
	var organs []BmbpRbacOrgan
	if err := query.Find(&organs).Error; err != nil {
		ctx.JSON(200, base.RespFail("查询组织机构树失败: "+err.Error()))
		return
	}
	ctx.JSON(200, base.RespOk(organs))
}
func OrganInfo(ctx *gin.Context) {
	data_id := ctx.Query("dataId")
	if data_id == "" {
		ctx.JSON(200, base.RespFail("组织机构ID不能为空"))
		return
	}
	var organ BmbpRbacOrgan
	if err := db.DbUtil.Model(&organ).Find(&organ, "data_id = ?", data_id).Error; err != nil {
		ctx.JSON(200, base.RespFail("查询组织机构详情失败: "+err.Error()))
		return
	}
	ctx.JSON(200, base.RespOkWithMsg(organ, "查询成功!"))
}
func OrganSave(ctx *gin.Context) {
	var organ BmbpRbacOrgan
	if err := ctx.ShouldBindJSON(&organ); err != nil {
		var msg string
		contentType := ctx.Request.Header.Get("Content-Type")
		mediaType := strings.Split(contentType, ";")[0]
		if !strings.HasPrefix(mediaType, "application/json") {
			msg = "请传递json格式的参数,当前传递的是: " + mediaType
		} else {
			msg = err.Error()
		}
		ctx.JSON(200, base.RespFail("组织机构参数接收失败: "+msg))
		return
	}
	if organ.DataId == "" {
		if err := insertOrgan(ctx, organ); err != nil {
			ctx.JSON(200, base.RespFail("新增组织机构失败: "+err.Error()))
			return
		}
	} else {
		if err := updateOrgan(ctx, organ); err != nil {
			ctx.JSON(200, base.RespFail("更新组织机构失败: "+err.Error()))
			return
		}
	}
	ctx.JSON(200, base.RespOkWithMsg(organ, "保存成功!"))
}

func insertOrgan(ctx *gin.Context, organ BmbpRbacOrgan) error {
	slog.Info("新增组织机构")
	slog.Info("校验组织信息....")
	if organ.NodeTitle == "" {
		return errors.New("组织名称不能为空")
	}
	organ.NodeCode = base.SimpleGuid()
	if organ.NodeParentCode == "" {
		organ.NodeParentCode = base.TREE_NODE_ROOT
	}
	// 校验组织编码是否重复
	checkSameCode, err := queryOrganByCode(organ.NodeCode)
	if err != nil {
		return err
	}
	if checkSameCode.DataId != "" {
		slog.Info("组织编码重复", "nodeCode", organ.NodeCode)
		return errors.New("组织编码重复")
	}
	// 校验组织名称是否重复
	checkSameName, err := queryOrganByName(organ.NodeTitle, organ.NodeParentCode)
	if err != nil {
		return err
	}
	if checkSameName.DataId != "" {
		return errors.New("组织名称重复")
	}
	pOrgan, err := queryOrganByCode(organ.NodeParentCode)
	if err != nil {
		return err
	}
	if pOrgan == nil {
		return errors.New("父级组织不存在")
	}
	// 计算路径和层级
	organ.NodeCodePath = pOrgan.NodeCodePath + organ.NodeCode + ","
	organ.NodeTitlePath = pOrgan.NodeTitlePath + organ.NodeTitle + ","
	organ.NodeGrade = organ.NodeGrade + 1

	organ.CreateTime = base.NowTimeStr()
	organ.CreateUser = "_"
	organ.UpdateTime = base.NowTimeStr()
	organ.UpdateUser = "_"
	organ.DataStatus = "1"
	organ.DataFlag = "0"
	organ.DataId = base.Guid()
	if err := db.DbUtil.Create(&organ).Error; err != nil {
		return err
	}
	return nil
}

func queryOrganByName(nodeTitle, nodeParentCode string) (*BmbpRbacOrgan, error) {
	var organ BmbpRbacOrgan
	if err := db.DbUtil.Model(&organ).Find(&organ, "node_title = ? and node_parent_code = ?", nodeTitle, nodeParentCode).Error; err != nil {
		return nil, err
	}
	return &organ, nil
}

func queryOrganByCode(nodeCode string) (*BmbpRbacOrgan, error) {
	if nodeCode == "" {
		return nil, errors.New("组织编码不能为空")
	}
	if nodeCode == base.TREE_NODE_ROOT {
		return &BmbpRbacOrgan{
			BaseBean: base.BaseBean{
				DataId: base.TREE_NODE_ROOT,
			},
			BaseTree: base.BaseTree{
				NodeCode:      base.TREE_NODE_ROOT,
				NodeTitle:     base.TREE_NODE_ROOT,
				NodeCodePath:  base.TREE_NODE_ROOT + ",",
				NodeTitlePath: base.TREE_NODE_ROOT + ",",
				NodeGrade:     0,
			},
		}, nil
	}
	var organ BmbpRbacOrgan
	if err := db.DbUtil.Model(&organ).Find(&organ, "node_code =?", nodeCode).Error; err != nil {
		return nil, err
	}
	return &organ, nil
}
func updateOrgan(ctx *gin.Context, organ BmbpRbacOrgan) error {
	slog.Info("更新组织机构")
	if err := db.DbUtil.Model(&organ).Updates(&organ).Error; err != nil {

	}
	return nil
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
