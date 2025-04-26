package organ

import (
	"bmbp/base"
)

type BmbpRbacOrgan struct {
	base.BaseBean
	base.BaseTree
}

func (o BmbpRbacOrgan) GetCode() string {
	return o.NodeCode
}
func (o BmbpRbacOrgan) GetParentCode() string {
	return o.NodeParentCode
}
func (o BmbpRbacOrgan) GetChildren() []base.TreeNode {
	return o.NodeChildren
}
func (o BmbpRbacOrgan) SetChildren(children []base.TreeNode) {
	o.NodeChildren = children
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

type BmbpRbacOrganArea struct {
	base.BaseBean
	NodeCode string `json:"nodeCode" gorm:"not null;type:string;size:36;comment:组织编码"` // 组织编码
}
type BmbpRbacOrganGroup struct {
	base.BaseBean
	NodeCode string `json:"nodeCode" gorm:"not null;type:string;size:36;comment:组织编码"` // 组织编码
}
type BmbpRbacOrganUnit struct {
	base.BaseBean
	NodeCode string `json:"nodeCode" gorm:"not null;type:string;size:36;comment:组织编码"` // 组织编码
}
type BmbpRbacOrganDept struct {
	base.BaseBean
	NodeCode string `json:"nodeCode" gorm:"not null;type:string;size:36;comment:组织编码"` // 组织编码
}
type BmbpRbacOrganPost struct {
	base.BaseBean
	NodeCode string `json:"nodeCode" gorm:"not null;type:string;size:36;comment:组织编码"` // 组织编码
}
type BmbpRbacOrganPerson struct {
	base.BaseBean
	NodeCode string `json:"nodeCode" gorm:"not null;type:string;size:36;comment:组织编码"` // 组织编码
}
