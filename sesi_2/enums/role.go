package enums

type RoleUser string

const (
	Admin RoleUser = "admin"
	User  RoleUser = "user"
)

// func (e *RoleUser) Scan(value interface{}) error {
// 	*e = RoleUser(value.([]byte))
// 	return nil
// }

// func (e RoleUser) Value() (driver.Value, error) {
// 	return string(e), nil
// }
