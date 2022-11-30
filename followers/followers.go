package followers

import (
	"assignment-activity-reporter/users"
)

type Followers interface {
	UploadActivity(*users.User)
	LikeActivity(*users.User, *users.User)
}
