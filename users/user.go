package users

import (
	"fmt"
	"time"

	"github.com/og11423074s/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)

}
