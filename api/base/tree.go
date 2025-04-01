package base

type BmbpTree interface {
	GetCode() string
	GetParentCode() string
	GetChildren() []BmbpTree
	SetChildren(children []BmbpTree)
}
