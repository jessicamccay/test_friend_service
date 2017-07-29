package handlers

import (
	"friends.uacf.io/apps/friendsapi/rpc"
	"friends.uacf.io/domain"
)

func toRpcFriendship(domainItem *domain.Friendship) *rpc.Friendship {
	return &rpc.Friendship{
		Id:   domainItem.Id,
		FromUserId: domainItem.FromUserId,
		ToUserId: domainItem.ToUserId,
		FriendsSince: domainItem.FriendsSince,
	}
}

func toRpcFriendships(domainItems []*domain.Friendship) []*rpc.Friendship {
	var rpcItems []*rpc.Friendship
	for _, model := range domainItems {
		rpcItems = append(rpcItems, toRpcFriendship(model))
	}
	return rpcItems
}

func toDomainFriendship(rpcItem *rpc.Friendship) *domain.Friendship {
	return &domain.Friendship{
		Id:   rpcItem.Id,
		FromUserId: rpcItem.FromUserId,
		ToUserId: rpcItem.ToUserId,
		FriendsSince: rpcItem.FriendsSince,
	}
}

func toDomainFriendships(rpcItems []*rpc.Friendship) []*domain.Friendship {
	var domainItems []*domain.Friendship
	for _, model := range rpcItems {
		domainItems = append(domainItems, toDomainFriendship(model))
	}
	return domainItems
}
