package services
//
//import (
//	"testing"
//
//	"go.uacf.io/apierrors"
//
//	"friends.uacf.io/data"
//	"friends.uacf.io/domain"
//	"friends.uacf.io/services/auth"
//
//	. "github.com/smartystreets/goconvey/convey"
//	"golang.org/x/net/context"
//)
//
///*
//This is an example of a unit test suite against the service layer by using
//a mock data layer which allows you to test the service in isolation.
//*/
//func TestItemService(t *testing.T) {
//	Convey("With a service layer using an in-memory data layer", t, func() {
//		db := data.NewMemoryData()
//		service := NewItemService(db)
//		bareCtx := context.Background()
//		ctx := auth.NewContext(bareCtx, 7)
//
//		Convey("#List", func() {
//			items, err := service.List(ctx)
//			So(err, ShouldBeNil)
//			So(len(items), ShouldBeGreaterThanOrEqualTo, 10)
//		})
//
//		Convey("#Create", func() {
//			item := &domain.Item{
//				Name:     "foobar",
//				AuthorId: 7,
//			}
//
//			Convey("Fails when we aren't logged in", func() {
//				_, err := service.Create(bareCtx, item)
//				So(err, ShouldResemble, apierrors.Unauthenticated)
//			})
//
//			Convey("Fails when author ID doesn't match", func() {
//				item.AuthorId = 4
//				_, err := service.Create(ctx, item)
//				So(err, ShouldResemble, apierrors.Forbidden)
//			})
//
//			Convey("Can successfully create an item", func() {
//				newItem, err := service.Create(ctx, item)
//				So(err, ShouldBeNil)
//				So(newItem, ShouldNotBeNil)
//				So(newItem.Id, ShouldNotEqual, 0)
//				So(newItem.AuthorId, ShouldEqual, 7)
//			})
//		})
//
//		// etc, more unit tests
//	})
//}
