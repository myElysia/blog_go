package model

import (
	"gorm.io/gorm"
)

type User struct {
	Model      gorm.Model  `gorm:"embedded"`
	Username   string      `gorm:"comment:用户名称;Not Null;"`
	Password   string      `gorm:"comment:密码(注意加密);"`
	Email      string      `gorm:"comment:邮箱;Not Null;"`
	Github     string      `gorm:"comment:Github;"`
	LastLogin  string      `gorm:"autoUpdateTime;comment:上次登录时间;"`
	LoginIP    string      `gorm:"comment:登录IP;"`
	UserGroups []UserGroup `gorm:"many2many:user_group_users;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Roles      []Role      `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserGroup struct {
	Model gorm.Model `gorm:"embedded"`
	Name  string     `gorm:"comment:用户组名称;Not Null;"`
	Roles []Role     `gorm:"many2many:user_group_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Role struct {
	Model       gorm.Model   `gorm:"embedded"`
	Name        string       `gorm:"comment:角色名称;Not Null;"`
	Remark      string       `gorm:"comment:描述;"`
	Permissions []Permission `gorm:"many2many:user_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Permission struct {
	Model    gorm.Model `gorm:"embedded"`
	ParentID uint       `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name     string     `gorm:"comment:权限名称;"`
	Remark   string     `gorm:"comment:描述;"`
}

func (user *User) GetPermissions() []Permission {
	var permissions []Permission
	for _, group := range user.UserGroups {
		for _, role := range group.Roles {
			if role.Permissions == nil {
				continue
			}
			permissions = append(permissions, role.Permissions...)
		}
	}

	for _, role := range user.Roles {
		if role.Permissions == nil {
			continue
		}
		permissions = append(permissions, role.Permissions...)
	}

	return permissions
}
