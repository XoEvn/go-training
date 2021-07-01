package proxy

import (
	"log"
	"time"
)







// IUser IUser
type IUser interface {
	Login(username,password string) error
}
// User 用户
type User struct {

}

func ( u *User) Login(username,password string) error {
	return nil
}

type UserProxy struct {
	user *User
}
func NewUserProxy(user *User) *UserProxy{
	return &UserProxy{
		user:user,
	}
}


func ( p *UserProxy)Login(username,password string) error {
	start := time.Now()

	if err := p.user.Login(username,password);err!=nil{
		return err
	}

	log.Printf("user login cost time: %s",time.Now().Sub(start))
	return nil
}



