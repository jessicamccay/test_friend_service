package handlers

import (
	"testing"

	"friends.uacf.io/apps/friendsapi/rpc"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRpcToDomain(t *testing.T) {
	/*
	TestRpcToDomain is an example of adding a unit test class
	and using convey for assertions

	These are automatically ran on build and can also be ran with go test -short
	*/
	Convey("When translating Rpc to Domain", t, func() {
		Convey("going to domain succeeds", func() {
			rpcItem := &rpc.Friendship{
				Id: "1",
				ToUserId: "123",
				FromUserId: "321",
				FriendsSince: "2012-12-31 00:00"}
			domainItem := toDomainFriendship(rpcItem)
			//TODO Check all functionality
			So(domainItem, ShouldNotBeNil)
			So(domainItem.Id, ShouldEqual, rpcItem.Id)
		})
	})
}
