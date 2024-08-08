package models

type Auth struct {
	User_id  string `db:"user_id" json:"user_id" valid:"-"`
	Email    string `db:"email" json:"email" valid:"email"`
	Password string `db:"password" json:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Phone    string `db:"phone" json:"phone" valid:"-"`
	Role     string `db:"role" json:"role" valid:"-"`
}

type Auths []Auth