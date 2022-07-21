package main

func insert(u *UserInfo) error {

	stmt, err := ObtainConn().Prepare("INSERT INTO user_info(name, deleted) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Deleted)
	if err != nil {
		return err
	}

	return nil
}
