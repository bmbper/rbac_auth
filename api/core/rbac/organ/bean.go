package core

import (
	"api/base"
)

type BmbpRbacOrgan struct {
	base.BaseBean
	OrganCode       string    `json:"orgCode"`         // 组织编码
	OrganName       string    `json:"orgName"`         // 组织名称
	OrganType       OrganType `json:"organType"`       // 组织类型
	OrganCodePath   string    `json:"organCodePath"`   // 组织编码路径
	OrganNamePath   string    `json:"organNamePath"`   // 组织名称路径
	OrganParentCode string    `json:"organParentCode"` // 上级组织编码
	OrganGrade      int       `json:"organGrade"`      // 组织层级
}

type OrganType int

const (
	OrganTypeDivision   OrganType = iota + 1 // 区划
	OrganTypeGroup                           // 分组
	OrganTypeUnit                            // 单位
	OrganTypeDepartment                      // 部门
	OrganTypePosition                        // 岗位
	OrganTypePerson                          // 人员
)
