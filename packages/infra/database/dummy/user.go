package dummy

import "github.com/aokuyama/go-study-layered/packages/domain/model/user"

type UserRepositoryDummy struct {
}

func (r UserRepositoryDummy) LoadByID(id user.UserID) (*user.User, error) {
	if id == "123456789" {
		return user.New("太郎", "山田", "ok@example.com")
	}
	return nil, nil
}
