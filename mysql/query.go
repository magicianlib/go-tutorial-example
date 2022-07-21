package main

import (
	"fmt"
)

func query() []*UserInfo {

	rows, err := ObtainConn().Query("SELECT id, name, deleted FROM user_info")
	if err != nil {
		panic(err)
	}

	// columns, _ := rows.Columns()
	//
	// columnTypes, _ := rows.ColumnTypes()
	// for _, columnType := range columnTypes {
	// 	fmt.Println(columnType.Name(), columnType.ScanType(), columnType.DatabaseTypeName())
	// }

	// values := make([]sql.RawBytes, len(columns))
	// scanArgs := make([]interface{}, len(columns))
	//
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	var users = make([]*UserInfo, 1)

	for rows.Next() {

		// err := rows.Scan(scanArgs...)
		// if err != nil {
		// 	panic(err)
		// }

		var u = &UserInfo{}
		err := rows.Scan(&u.Id, &u.Name, &u.Deleted)
		if err != nil {
			panic(err)
		}

		users = append(users, u)
	}

	fmt.Printf("%#v\n", users)
	return users
}

func queryRow() (*UserInfo, error) {

	row := ObtainConn().QueryRow("SELECT id, name, deleted FROM user_info WHERE id = 2")
	if row.Err() != nil {
		return nil, row.Err()
	}

	var u = &UserInfo{}
	err := row.Scan(&u.Id, &u.Name, &u.Deleted)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v\n", u)

	return u, nil
}
