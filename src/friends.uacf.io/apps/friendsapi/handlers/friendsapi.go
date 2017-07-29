package handlers

import (
	"time"

	"friends.uacf.io/apps/friendsapi/rpc"
	"friends.uacf.io/services"

	log "go.uacf.io/logging"
	"go.uacf.io/metrics"
	"golang.org/x/net/context"
)

type FriendsApiServer struct {
	service services.FriendService
}

func NewFriendsApiServer(service services.FriendService) rpc.FriendsApiServiceServer {
	return &FriendsApiServer{service}
}

func (t *FriendsApiServer) Ping(ctx context.Context, request *rpc.PingRequest) (*rpc.PongResponse, error) {
	return &rpc.PongResponse{Pong: true}, nil
}

// This is an example of a simple handler so that the full set of operations is de-mystified.
// With this you can see every action that's happening without any abstraction hiding it
func (t *FriendsApiServer) Get(ctx context.Context, request *rpc.GetRequest) (*rpc.GetResponse, error) {
	// log a debug message on request handler begin
	log.With(log.Fields{"context": ctx, "MYDEBUGGGG   request": request}).Debug("get")
	// instrument metrics which count number of requests and average request time.
	metrics.Inc("get.requested", 1)
	defer metrics.TimeElapsed("get.response_time", time.Now())

	// Extract authentication from GRPC metadata and annotate it into our context
	ctx = AuthContext(ctx)

	// Call into the service layer and get the item
	item, err := t.service.Get(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	// Convert the response back to the RPC format and return
	resp := &rpc.GetResponse{
		Friendship: toRpcFriendship(item),
	}
	return resp, nil
}

//// Another simple example without all the commenting to see it a bit more densely.
//func (t *FriendsApiServer) Create(ctx context.Context, request *rpc.CreateRequest) (*rpc.CreateResponse, error) {
//	log.With(log.Fields{"context": ctx, "request": request}).Debug("create")
//	metrics.Inc("create.requested", 1)
//	defer metrics.TimeElapsed("create.response_time", time.Now())
//
//	item, err := t.service.Create(ctx, toDomainItem(request.Item))
//	if err != nil {
//		return nil, grpcError(err)
//	}
//
//	resp := &rpc.CreateResponse{
//		Item: toRpcItem(item),
//	}
//	return resp, nil
//}
//
//// Now that we've shown you the Get example, we have another example which
//// shows a potential pattern you can write to do your "repetetive work" for all requests.
//func (t *FriendsApiServer) Update(ctx context.Context, request *rpc.UpdateRequest) (resp *rpc.UpdateResponse, err error) {
//	err = Middleware(ctx, "update", func(ctx context.Context) error {
//		item, err := t.service.Update(ctx, toDomainItem(request.Item))
//		if err != nil {
//			return err
//		}
//
//		resp = &rpc.UpdateResponse{
//			Item: toRpcItem(item),
//		}
//		return nil
//	})
//	return
//}
//
//func (t *FriendsApiServer) Delete(ctx context.Context, request *rpc.DeleteRequest) (resp *rpc.DeleteResponse, err error) {
//	err = Middleware(ctx, "delete", func(ctx context.Context) error {
//		err := t.service.Delete(ctx, request.Id)
//		if err != nil {
//			return err
//		}
//
//		resp = &rpc.DeleteResponse{}
//		return nil
//	})
//	return
//}

func (t *FriendsApiServer) List(ctx context.Context, request *rpc.ListRequest) (resp *rpc.ListResponse, err error) {
	log.Debugf("in handler MCCAY request: %s", request)
	err = Middleware(ctx, "list", func(ctx context.Context) error {
		items, err := t.service.List(ctx, request.Status, request.FromUserId, request.ToUserId)
		if err != nil {
			return err
		}

		resp = &rpc.ListResponse{
			Friendships: toRpcFriendships(items),
		}
		return nil
	})
	return
}
//
//// Here is an example of a handler for an imperative method.
////
//// From the perspective of GRPC itself, there's really no difference, all
//// methods are the same, the method is only imperative in that the REST mapper
//// cannot map this to a RESTful call in the REST+JSON layer.
//func (t *FriendsApiServer) Method(ctx context.Context, request *rpc.MethodRequest) (resp *rpc.MethodResponse, err error) {
//	err = Middleware(ctx, "method", func(ctx context.Context) error {
//		value, err := t.service.Method(ctx, request.Parameter1, request.Parameter2)
//		if err != nil {
//			return err
//		}
//
//		resp = &rpc.MethodResponse{
//			Result: value,
//		}
//		return nil
//	})
//	return
//}
//
//// Show how to do a streaming response method.
////
//// The code-gen creates a "Server" object which allows you to send (or sometimes receive) request objects
//// that you can keep sending data to. In order to implement this concept at the service level, we have the
//// service layer returning a go channel which we then can send items as they come out of the channel.
////
//// This way, the service layer can simply close the channel when it's done sending results allowing the
//// streaming response method to complete.
//func (t *FriendsApiServer) StreamingMethod(request *rpc.MethodRequest, stream rpc.FriendsApiService_StreamingMethodServer) error {
//	// Set up a cancellation context so we cancel the underlying RPC whenever we end for any reason.
//	// A likely scenario is we hit EOF or timeout or lost connection sending to our requester, for example.
//	ctx, cancel := context.WithCancel(AuthContext(stream.Context()))
//	defer cancel()
//	entry := log.With(log.Fields{"context": ctx, "request": request})
//	entry.Debug("streamingMethod")
//	// call our streaming method which gives us a channel of results.
//	ch, err := t.service.StreamingList(ctx)
//	if err != nil {
//		return err
//	}
//	// loop the channel, sending and marshaling each result we get.
//	for item := range ch {
//		err = stream.Send(toRpcFriendship(item))
//		if err != nil {
//			entry.Warnf("Got error %v while doing streaming", err)
//			break
//		}
//	}
//	return err
//}
