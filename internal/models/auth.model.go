package models

type Auth struct {
	User_id  string `db:"user_id" json:"user_id" form:"user_id" valid:"-"`
	Email    string `db:"email" json:"email" form:"email" valid:"email"`
	Password string `db:"password" json:"password" form:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Phone    string `db:"phone" json:"phone" form:"phone" valid:"-"`
	Role     string `db:"role" json:"role" form:"role" valid:"-"`
}

type Auths []Auth