package instagram

import "github.com/Davincible/goinsta"

type InstaUser struct {
	User *goinsta.Instagram
}

func NewUser(login, password string) *InstaUser {
	user := goinsta.New(login, password)
	return &InstaUser{
		User: user,
	}
}

func (i *InstaUser) GetUserStories(userName string) ([]*goinsta.Item, error) {
	profile, err := i.User.VisitProfile(userName)
	if err != nil {
		return nil, err
	}

	stories := profile.Stories.Reel

	return stories.Items, nil
}
