package serialize

import "github.com/kadycui/gin-rank/model"

type UserSerialize struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserSerialize(user model.User) UserSerialize {
	return UserSerialize{
		Name:      user.Name,
		Telephone: user.Telephone,
	}

}
