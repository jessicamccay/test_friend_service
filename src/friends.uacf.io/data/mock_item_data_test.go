package data

import (
	"os"
	"testing"
	//"friends.uacf.io/domain"
	//. "github.com/smartystreets/goconvey/convey"
)

var funcTest functional

type functional struct{}

// This 'setup' function is a pattern we came up with to do the common pattern
// of skipping functional tests when some pre-conditions are not met.
// And then tests can simply start with
//    funcTest.setUp()
func (f *functional) setUp(t *testing.T) FriendData {
	if testing.Short() || os.Getenv("SKIP_FUNCTIONAL_TESTS") != "" {
		t.SkipNow()
	}

	// You could do other setup actions here.
	db, _ := NewFriendData("root", "password", "tcp(mysql:3306)/datapath")
	return db
}

///*
//TestCreate is just an example of adding a functional test class
//and using convey for assertions
//
//This would normally be used if you needed to talk to an external api
//or database linked in functional.yml
//
//These are automatically ran on gerrit review and can also be ran with go test
//*/
//func TestCreate(t *testing.T) {
//	db := funcTest.setUp(t)
//
//	Convey("With a backing store", t, func() {
//
//		Convey("simple create succeeds", func() {
//			domainItem := &domain.Item{Id: 0, Name: "name"}
//			item, err := db.Create(domainItem)
//			//TODO Check all functionality with a lot more tests
//			So(err, ShouldBeNil)
//			So(item.Id, ShouldNotEqual, 0)
//			So(domainItem.Id, ShouldEqual, item.Id)
//		})
//
//		Convey("create ignores any ID's you pass in", func() {
//			domainItem := &domain.Item{Id: 4, Name: "name"}
//			item, err := db.Create(domainItem)
//			So(err, ShouldBeNil)
//			So(item.Id, ShouldNotEqual, 4)
//
//		})
//	})
//}
