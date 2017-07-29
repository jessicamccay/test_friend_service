package domain

type Friendship struct {
	Id             string
	FromUserId     string
	ToUserId       string
	FriendsSince   string  // datetime timestamp in Central Standard
	Status         string
}
