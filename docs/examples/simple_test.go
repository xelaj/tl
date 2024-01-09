package main_test

import (
	"fmt"

	"github.com/xelaj/tl"
)

type RepoUser struct {
	_       struct{} `tl:"flag,bitflag"`
	Nick    string
	Creator bool `tl:",omitempty:flag:0,implicit"`
	Editor  bool `tl:",omitempty:flag:7,implicit"`
}

func (c *RepoUser) CRC() uint32 { return 0x12345678 }

func ExampleSimplest() {
	tl.RegisterObjectDefault[*RepoUser]()

	response := &RepoUser{
		Nick:    "Hello user!",
		Creator: true,
		Editor:  true,
	}

	b, _ := tl.Marshal(response)
	fmt.Printf("%x\n", b)
}
