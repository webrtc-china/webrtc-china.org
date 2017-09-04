package views

import "webrtc-china.org/models"

type UserView struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
}

func BuildUserView(user *models.User) UserView {
	userView := UserView{
		Id:        user.Id,
		Email:     user.Email,
		Username:  user.Username,
		FullName:  user.FullName,
		AvatarURL: user.AvatarURL,
	}
	return userView
}
