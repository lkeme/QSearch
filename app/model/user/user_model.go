package user

import (
	"github.com/lkeme/QSearch/app/model"
	"github.com/lkeme/QSearch/pkg/database"
	"github.com/lkeme/QSearch/pkg/hash"
	"github.com/satori/go.uuid"
)

type User struct {
	model.BaseModel
	UUID     uuid.UUID `json:"uuid" gorm:"comment:Uuid"`
	Username string    `json:"userName" gorm:"comment:Username"`
	Password string    `json:"-"  gorm:"comment:Password"`
	NickName string    `json:"nickName" gorm:"default:系统用户;comment:Nickname"`
	Avatar   string    `json:"headerImg" gorm:"avatar:http://q1.qlogo.cn/g?b=qq&nk=100111&s=640;comment:avatar"`
	Email    string    `json:"-" gorm:"comment:email"`
	Phone    string    `json:"-" gorm:"comment:phone"`
	model.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
