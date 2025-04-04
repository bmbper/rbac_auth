package base

type BaseBean struct {
	DataId     string `json:"dataId" gorm:"primaryKey;not null;type:string;size:36;comment:记录主键" `
	DataFlag   string `json:"dataFlag" gorm:"not null;default:'0';type:string;size:4;comment:记录标识"`
	DataStatus string `json:"dataStatus" gorm:"not null;default:'1';type:string;size:4;comment:记录状态"`
	DataSort   int64  `json:"dataSort" gorm:"not null;default:0;type:bigInt;size:10;comment:记录顺序"`
	DataLevel  string `json:"dataLevel" gorm:"not null;default:'0';type:string;size:4;comment:记录等级"`
	DataOwner  string `json:"dataOwner" gorm:"not null;default:'';type:string;size:36;comment:所属组织"`
	DataSign   string `json:"dataSign" gorm:"not null;default:'';type:string;size:256;comment:记录签名"`
	CreateTime string `json:"createTime" gorm:"not null;default:'';type:string;size:20;comment:创建时间"`
	UpdateTime string `json:"updateTime" gorm:"not null;default:'';type:string;size:20;comment:更新时间"`
	CreateUser string `json:"createUser" gorm:"not null;default:'';type:string;size:36;comment:创建用户"`
	UpdateUser string `json:"updateUser" gorm:"not null;default:'';type:string;size:36;comment:更新用户"`
}

type TreeNode interface {
	GetCode() string
	GetParentCode() string
	GetChildren() []TreeNode
	SetChildren(children []TreeNode)
}

type BaseTree struct {
	NodeCode       string     `json:"nodeCode" gorm:"not null;type:string;size:36;comment:节点编码"`
	NodeTitle      string     `json:"nodeTitle" gorm:"not null;type:string;size:128;comment:节点名称"`
	NodeParentCode string     `json:"nodeParentCode" gorm:"not null;type:string;size:36;comment:父节点编码"`
	NodeCodePath   string     `json:"nodeCodePath" gorm:"not null;type:text;comment:节点编码路径"`
	NodeTitlePath  string     `json:"nodeTitlePath" gorm:"not null;type:text;comment:节点编码名称"`
	NodeType       string     `json:"nodeType" gorm:"not null;type:string;size:10;comment:节点类型"`
	NodeGrade      int        `json:"nodeGrade" gorm:"not null;type:int;size:10;comment:节点层级"`
	NodeLeaf       int        `json:"nodeLeaf" gorm:"-"`
	NodeChecked    int        `json:"nodeChecked" gorm:"-"`
	NodeChildren   []TreeNode `json:"nodeChildren" gorm:"-"`
}

var TREE_NODE_ROOT = "#"
var TREE_NODE_SLIT = "/"
