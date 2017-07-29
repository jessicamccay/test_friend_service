package services

import (
	//"fmt"
	//"time"

	"friends.uacf.io/data"
	"friends.uacf.io/domain"
	//"friends.uacf.io/services/auth"

	//"go.uacf.io/apierrors"
	"golang.org/x/net/context"
	"go.uacf.io/logging"
)

type FriendService interface {
	Get(ctx context.Context, id int64) (*domain.Friendship, error)
	List(ctx context.Context, status string, from_user_id string, to_user_id string) ([]*domain.Friendship, error)
	//StreamingList(ctx context.Context) (<-chan *domain.Friendship, error)
}

type friendService struct {
	data    data.FriendData
}

func NewFriendService(data data.FriendData) FriendService {
	return &friendService{data}
}

func (s *friendService) Get(ctx context.Context, id int64) (*domain.Friendship, error) {
	item, err := s.data.Get(id)
	//if err == nil && !auth.FromContext(ctx).MatchUserId(item.AuthorId) {
	//	return nil, apierrors.Forbidden
	//}
	return item, err
}

func (s *friendService) List(ctx context.Context, status string, from_user_id string, to_user_id string)(items []*domain.Friendship, err error) {
	logging.Debugf("SERVICE CONTEXT: %s status %s from_user_id %s to_user_id %s", ctx, status,
		from_user_id, to_user_id)
	return s.data.List(status, from_user_id, to_user_id)
}
//
//// We're doing this just for the fun of it, providing an intentionally slow stream method.
//// This shows all the patterns like cancellation, deadlines, net-context, and streaming.
//func (s *friendService) StreamingList(ctx context.Context) (<-chan *domain.Friendship, error) {
//	items, _ := s.data.List()
//	c := make(chan *domain.Friendship)
//	go func() {
//		defer close(c)F
//		for _, item := range items {
//			select {
//			case <-ctx.Done():
//				break
//			case c <- item:
//				time.Sleep(50 * time.Millisecond)
//			}
//		}
//	}()
//	return c, nil
//}
