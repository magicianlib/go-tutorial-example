package main

func main() {
	query()
	queryRow()

	insert(&UserInfo{
		Name:    "Smith",
		Deleted: 1,
	})

	defer DestroyConn()
}
