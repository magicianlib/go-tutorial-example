package main

type UserInfo struct {
	Id      uint64 `sql:"id"`
	Name    string `sql:"name"`
	Deleted uint8  `sql:"deleted"`
}
