package json

import "time"

type User struct {
	Username string
	Password string
	Created  time.Time
}
