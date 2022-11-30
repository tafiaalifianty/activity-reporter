package users

type Database struct {
	User   map[string]*User
	Photos []*Photo
}

type UserAlreadyExistError struct {
}

type NoUserExistError struct {
}

func (u *UserAlreadyExistError) Error() string {
	return "User Already Exists"
}

func (nu *NoUserExistError) Error() string {
	return "unknown user"
}

func NewDB() *Database {
	return &Database{
		User: make(map[string]*User, 0),
	}
}

func (d *Database) AddUser(u *User) error {
	key := u.Username
	if _, ok := d.User[key]; ok {
		return &UserAlreadyExistError{}
	}
	d.User[u.Username] = u
	return nil
}

func (d *Database) GetUser(username string) (*User, error) {
	if _, ok := d.User[username]; ok {
		return d.User[username], nil
	}
	return nil, &NoUserExistError{}
}

func (d *Database) IsExist(username string) bool {
	_, ok := d.User[username]
	return ok
}
