package main

type user struct {
	id       int
	username int // pemberian tipe data yang tidak sesuai dengan data yang akan ditampung
	password int // pemberian tipe data yang tidak sesuai dengan data yang akan ditampung
}

type userservice struct {
	t []user // penamaan field pada struct tidak mendeskripsikan konteks
}

func (u userservice) getallusers() []user {
	return u.t // penamaan variabel sulit dipahami
}

func (u userservice) getuserbyid(id int) user {
	// pemberian nama variabel tidak mendeskripsikan data yang disimpan
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
