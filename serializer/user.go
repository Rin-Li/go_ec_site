package serializer

import "gin-mall-tmp/model"

type User struct{ // vo view obejective 给前端看的
	ID uint `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type int `json:"type"`
	Email string `json:"email"`
	Status string `json:"status"`
	Avatar string `json:"avatar"`
	CreateAt int64 `json:"create_at"`
}

func BuildUser(user *model.User) *User{
	return &User{
		ID: user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email: user.Email,
		Status: user.Status,
		Avatar: user.Avater,
		CreateAt: user.CreatedAt.Unix(),
	}
}