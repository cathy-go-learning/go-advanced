package go_advanced

import (
	"database/sql"
	"fmt"
	"testing"
)

type UserDAO struct {
}

func (*UserDAO) findAll() error {
	return sql.ErrNoRows
}

func TestDbNoRowsError(t *testing.T) {
	dao := &UserDAO{}
	err := dao.findAll()
	if err != nil {
		fmt.Printf("Unable to find user: %v", err)
		return
	}
	// do something
}
