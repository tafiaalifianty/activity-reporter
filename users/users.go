package users

import (
	"fmt"
	"sort"
)

type User struct {
	Username   string
	LikePhotos map[string]*User
	Followers  map[string]*User
	Activities []string
	Photos     *Photo
}

type Photo struct {
}

func NewUser(username string) *User {
	newUser := User{
		Username:   username,
		Followers:  make(map[string]*User, 0),
		Activities: make([]string, 0),
		Photos:     nil,
		LikePhotos: make(map[string]*User, 0),
	}
	return &newUser
}

func (u *User) Follow(target *User) {
	target.AddFollower(u)
}

func (u *User) AddFollower(follower *User) {
	u.Followers[follower.Username] = follower
}

func (u *User) Upload(Photos *Photo) {
	if u.Photos == nil {
		u.Photos = Photos
		u.AddActivity("You uploaded photo")
		for _, value := range u.Followers {
			(*value).UploadActivity(u)
		}
	}
}

func (u *User) LikePhoto(target *User) {
	if _, ok := u.LikePhotos[target.Username]; !ok {
		target.LikePhotos[u.Username] = u
		if u.Username != target.Username {
			u.AddActivity(fmt.Sprintf("You liked %v's photo", target.Username))
		}
		if _, ok := u.Followers[target.Username]; !ok {
			target.LikeActivity(u, target)
		}
		for _, value := range u.Followers {
			(*value).LikeActivity(u, target)
		}
	}
}

func (u *User) AddActivity(message string) {
	u.Activities = append(u.Activities, message)
}

func (u *User) UploadActivity(uploader *User) {
	activity := fmt.Sprintf("%v uploaded photo", uploader.Username)
	u.AddActivity(activity)
}

func (u *User) LikeActivity(likers *User, liked *User) {
	likersName := likers.Username
	likedName := liked.Username

	if likedName != u.Username {
		likedName = likedName + "'s"
	}
	if likedName == u.Username {
		likedName = "your"
	}
	if likersName == u.Username {
		likersName = "You"
	}

	activity := fmt.Sprintf("%v liked %v photo", likersName, likedName)
	u.AddActivity(activity)
}

func (sys *Database) GetTrendingUser() []*User {
	usersAsSlice := []*User{}
	for _, user := range sys.User {
		if user.Photos != nil {
			usersAsSlice = append(usersAsSlice, user)
		}
	}

	sort.SliceStable(usersAsSlice, func(i, j int) bool {
		return usersAsSlice[i].GetLikersCount() > usersAsSlice[j].GetLikersCount()
	})

	var lastIdx int
	if len(usersAsSlice) < 3 {
		lastIdx = len(usersAsSlice)
	} else {
		lastIdx = 2
		for i := lastIdx + 1; i < len(usersAsSlice); i++ {
			if usersAsSlice[i].GetLikersCount() >= usersAsSlice[lastIdx].GetLikersCount() {
				lastIdx++
			} else {
				break
			}
		}
	}

	usersAsSlice = usersAsSlice[:lastIdx]

	return usersAsSlice
}

func (p *User) GetLikersCount() int {
	return len(p.LikePhotos)
}

func (u *User) GetName() string {
	return u.Username

}
