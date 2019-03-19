package entity

import "github.com/goosmesh/goos/core/support-db/orm"

func Init() {
	orm.Register(new(Account))
	orm.Register(new(User))
}