package followers

import (
	"errors"
	"justCode/lecture9/app/model"
)

type FollowersRepository interface {
	GetAll() []model.Followers
	GetByID(userID int, followerID int64) (*model.Followers, error)
	Create(follower *model.Followers) error
	Update(userID int, followerID int64, follower *model.Followers) error
	Delete(userID int, followerID int64) error
}

type FollowersRepositoryImpl struct {
	followers []model.Followers
}

func NewFollowersRepository() *FollowersRepositoryImpl {
	return &FollowersRepositoryImpl{}
}

func (r *FollowersRepositoryImpl) GetAll() []model.Followers {
	return r.followers
}

func (r *FollowersRepositoryImpl) GetByID(userID int, followerID int64) (*model.Followers, error) {
	for i := range r.followers {
		if r.followers[i].UserID == userID && r.followers[i].Follower_id == followerID {
			return &r.followers[i], nil
		}
	}
	return nil, errors.New("Not found")
}

func (r *FollowersRepositoryImpl) Create(follower *model.Followers) error {
	r.followers = append(r.followers, *follower)
	return nil
}

func (r *FollowersRepositoryImpl) Update(userID int, followerID int64, follower *model.Followers) error {
	for i := range r.followers {
		if r.followers[i].UserID == userID && r.followers[i].Follower_id == followerID {
			r.followers[i] = *follower
			return nil
		}
	}
	return errors.New("Not found")
}

func (r *FollowersRepositoryImpl) Delete(userID int, followerID int64) error {
	for i := range r.followers {
		if r.followers[i].UserID == userID && r.followers[i].Follower_id == followerID {
			r.followers = append(r.followers[:i], r.followers[i+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}
