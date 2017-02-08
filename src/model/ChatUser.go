package model

import (
	"log"
	. "common"
)

type ChatUser struct {
	Id       int    ` int(10) NOT NULL`
	Username string ` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL`
	Password string ` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL`
	Regtime  int    ` int(10) NOT NULL`
	Status   int    ` tinyint(2) NOT NULL DEFAULT 1`
	Groupid  int    ` int(5) NOT NULL DEFAULT 0`
	Email    string ` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL`
	Roleid   int    ` int(5) NOT NULL DEFAULT 0`
}

//查询所有列表，不建议使用
func (self *ChatUser) SelectAll() ([]ChatUser, error) {

	var chatuser []ChatUser
	err := Orm.Find(&chatuser)
	return chatuser, err
}

func (self *ChatUser) Count() (int64,error) {
	var chatuser ChatUser
	return Orm.Count(chatuser)
}

func (self *ChatUser) GetUser() (*ChatUser,error)  {
	chatuser := &ChatUser{}
	_, err := Orm.Id(self.Id).Get(chatuser)
	return chatuser,err
}

func (self *ChatUser) GetChatUserById(id int) (*ChatUser, error) {
	chatuser := &ChatUser{Id: id}
	_, err := Orm.Get(chatuser)
	return chatuser, err
}

func (self *ChatUser) Exist() (bool, error) {

	return Orm.Get(self)

}

func (self *ChatUser) ExistEmail() (bool, error) {
	log.Print(self.Email)
	return Orm.Get(&ChatUser{Email: self.Email})

}

type ChatUserSession struct {
	Username string ` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL`
}
