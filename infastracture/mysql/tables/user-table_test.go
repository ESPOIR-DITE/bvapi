package tables

import (
	"bvapi/domain"
	"fmt"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	user := domain.User{"", "0617825205", "espoir", "ditekemena", time.Now(), []byte{123}}
	res := CreateUser(user)
	fmt.Println(res)
}
