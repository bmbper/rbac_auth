package base

type BaseBean struct {
	DataId     string `json:"dataId" gorm:"primaryKey;not null"`
	DataFlag   string `json:"dataFlag" gorm:"not null;default:'0'"`
	DataStatus string `json:"dataStatus" gorm:"not null;default:'1'"`
	DataSort   int64  `json:"dataSort" gorm:"not null;default:0"`
	DataLevel  string `json:"dataLevel" gorm:"not null;default:'0'"`
	DataOwner  string `json:"dataOwner" gorm:"not null;default:''"`
	DataSign   string `json:"dataSign" gorm:"not null;default:''"`
	CreateTime string `json:"createTime" gorm:"not null;default:''"`
	UpdateTime string `json:"updateTime" gorm:"not null;default:''"`
	CreateUser string `json:"createUser" gorm:"not null;default:''"`
	UpdateUser string `json:"updateUser" gorm:"not null;default:''"`
}
