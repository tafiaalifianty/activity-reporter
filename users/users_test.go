package users_test

import (
	"assignment-activity-reporter/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Follow(t *testing.T) {
	t.Run("Should update Bob followers when Alice follows Bob", func(t *testing.T) {
		alice := users.NewUser("alice")
		bob := users.NewUser("bob")

		alice.Follow(bob)

		assert.Contains(t, bob.Followers, alice.Username)
	})

	t.Run("Bob followers should not contain a person that hasn't followed Bob yet", func(t *testing.T) {
		alice := users.NewUser("alice")
		bob := users.NewUser("bob")
		charlie := users.NewUser("charlie")

		alice.Follow(bob)

		assert.NotContains(t, bob.Followers, charlie.Username)
	})

	t.Run("Bob's followers map length should remain the same when followed twice by the same person", func(t *testing.T) {
		alice := users.NewUser("alice")
		bob := users.NewUser("bob")

		alice.Follow(bob)
		alice.Follow(bob)

		assert.Equal(t, len(bob.Followers), 1)
	})
}

func TestUser_Upload(t *testing.T) {
	t.Run("Alice, as Bob's follower, should be notified when Bob uploads a photo", func(t *testing.T) {
		alice := users.NewUser("alice")
		bob := users.NewUser("bob")
		photo := new(users.Photo)

		alice.Follow(bob)
		bob.Upload(photo)

		assert.Equal(t, alice.Activities[0], "bob uploaded photo")
	})
	t.Run("Charlie (not Bob's follower), should not be notified when Bob uploads a photo", func(t *testing.T) {
		charlie := users.NewUser("charlie")
		bob := users.NewUser("bob")
		photo := new(users.Photo)

		bob.Upload(photo)

		assert.Equal(t, len(charlie.Activities), 0)
	})
	t.Run("Bob should be notified when he uploads a photo", func(t *testing.T) {
		bob := users.NewUser("bob")
		photo := new(users.Photo)

		bob.Upload(photo)

		assert.Equal(t, bob.Activities[0], "You uploaded photo")
	})
}

func TestUser_LikeActivity(t *testing.T) {
	t.Run("Alice, Bob, and Charlie should be notified when Alice likes Bob's photo", func(t *testing.T) {
		alice := users.NewUser("alice")
		bob := users.NewUser("bob")
		charlie := users.NewUser("charlie")
		photo := new(users.Photo)

		charlie.Follow(alice)
		bob.Upload(photo)
		alice.LikePhoto(bob)

		assert.Equal(t, alice.Activities[0], "You liked bob's photo")
		assert.Equal(t, bob.Activities[0], "You uploaded photo")
		assert.Equal(t, bob.Activities[1], "alice liked your photo")
		assert.Equal(t, charlie.Activities[0], "alice liked bob's photo")
	})
}

func TestGetTrendingUser(t *testing.T) {
	t.Run("when a user has not upload, should not be included in function return", func(t *testing.T) {
		// Given
		system := users.NewDB()
		username1 := users.NewUser("user1")
		username2 := users.NewUser("user2")
		system.AddUser(username1)
		system.AddUser(username2)
		username1.Upload(&users.Photo{})

		// When
		topUsers := system.GetTrendingUser()

		// Then
		assert.Len(t, topUsers, 1)
		assert.NotContains(t, topUsers, username2)
	})

	t.Run("users with more likes should appear first", func(t *testing.T) {
		// Given
		system := users.NewDB()
		username1 := users.NewUser("user1")
		username2 := users.NewUser("user2")
		username3 := users.NewUser("user3")
		system.AddUser(username1)
		system.AddUser(username2)
		system.AddUser(username3)
		username2.Upload(&users.Photo{})
		username3.Upload(&users.Photo{})
		username1.Upload(&users.Photo{})
		username1.LikePhoto(username2)
		username3.LikePhoto(username1)
		username3.LikePhoto(username2)

		// When
		topUsers := system.GetTrendingUser()

		// Then
		assert.Equal(t, topUsers[0], username2)
		assert.Equal(t, topUsers[1], username1)
	})

	t.Run("should only display top 3 users", func(t *testing.T) {
		// Given
		system := users.NewDB()
		username1 := users.NewUser("user1")
		username2 := users.NewUser("user2")
		username3 := users.NewUser("user3")
		username4 := users.NewUser("user4")
		system.AddUser(username1)
		system.AddUser(username2)
		system.AddUser(username3)
		system.AddUser(username4)
		username1.Upload(&users.Photo{})
		username2.Upload(&users.Photo{})
		username3.Upload(&users.Photo{})
		username4.Upload(&users.Photo{})
		username1.LikePhoto(username2)
		username2.LikePhoto(username1)
		username3.LikePhoto(username1)
		username4.LikePhoto(username1)
		username3.LikePhoto(username2)
		username1.LikePhoto(username3)

		// When
		topUsers := system.GetTrendingUser()

		// Then
		assert.Len(t, topUsers, 3)
		assert.NotContains(t, topUsers, username4)
	})
}
