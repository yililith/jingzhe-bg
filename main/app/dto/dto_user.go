package dto

// ResUserLoginDto
// @Description: 用户登录数据类型
type UserLoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserDto struct {
	Username string `json:"username" binding:"required" validate:"alphanum,min=3,max=10"`
	Password string `json:"password" binding:"required" validate:"min=8,max=16,containsany=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.*%!#$"`
	Nickname string `json:"nickname" binding:"required" validate:"min=3,max=10"`
}

type GetUserPagingDto struct {
	Page int `form:"page" binding:"required"`
	Size int `form:"size" binding:"required"`
}

type PutUserAvatarDto struct {
	UID uint64 `json:"uid" binding:"required"`
}
