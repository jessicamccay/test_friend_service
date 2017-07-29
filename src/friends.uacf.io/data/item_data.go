package data

import (
	"database/sql"

	"friends.uacf.io/domain"

	"go.uacf.io/apierrors"
	"fmt"
	log "go.uacf.io/logging"
	//"go.uacf.io/metrics"
	//"encoding/json"
	//"golang.org/x/net/context"
	//"friends.uacf.io/services/auth"
	_ "github.com/go-sql-driver/mysql"
)
const (
	SelectActiveFriendshipById = "SELECT f.id, f.from_user_id, f.to_user_id, f.friends_since from friend f where f.id = ?"
	SelectActiveFriendshipsFromUser = "SELECT f.id, f.from_user_id, f.to_user_id, f.friends_since from friend f where f.from_user_id = ?"
	SelectActiveFriendshipsToUser = "SELECT f.id, f.from_user_id, f.to_user_id, f.friends_since from friend f where f.to_user_id = ?"
	SelectPendingFriendshipsFromUser = "SELECT f.id, f.requesting_user_id, f.requested_friend_id, f.date_requested from mmf_friend_request f where f.requesting_user_id = ?"
	SelectPendingFriendshipsToUser = "SELECT f.id, f.requesting_user_id, f.requested_friend_id, f.date_requested from mmf_friend_request f where f.requested_friend_id = ?"
)

// ItemData is the interface given by your data layer.
// We use an interface instead of a concrete type so that it's easy for you
// to override/mock for isolated tests.
type FriendData interface {
	Get(int64) (*domain.Friendship, error)
	List(status string, from_user_id string, to_user_id string) ([]*domain.Friendship, error)
}
type friendData struct {
	db *sql.DB
}

// This factory function is your "constructor" for your data layer.
// This will not work unmodified
//func NewItemData(dataSourceName string) (ItemData, error) {
//	conn, err := sql.Open("FIXME", dataSourceName)
//	if err != nil {
//		return nil, err
//	}
//	return &itemData{
//		conn: conn,
//	}, nil
//}
func NewFriendData(user string, password string, serverConn string) (*friendData, error) {
	connectionString := fmt.Sprintf("%v:%v@%v", user, password, serverConn)
	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	log.Debugf("NNNEEEWWW connectionstring: %s  conn: %s err: %s", connectionString, conn, err)

	return &friendData{
		db: conn,
	}, nil
}

func (d *friendData) Get(friendId int64) (*domain.Friendship, error) {

	var friendshipId, fromUserId, toUserId, friendsSince string
	err := d.db.QueryRow(SelectActiveFriendshipById, friendId).Scan(&friendshipId, &fromUserId, &toUserId, &friendsSince)
	if err != nil {
		log.Errorf("Placeholder error handling %v", err)
		return nil, apierrors.NewNotFound("Friendship not found")
	}
	friendDomain := domain.Friendship{
		Id: friendshipId,
		FromUserId: fromUserId,
		ToUserId: toUserId,
		FriendsSince: friendsSince,
		Status: "active",
	}

	log.Debugf("d.DB ==== %s", d.db)
	if err := d.db.Ping(); err != nil {
		log.Debugf("DATABASE ping error: %s", err)
	}

	//friendDomain := domain.Friendship{
	//	Id: "1",
	//	FromUserId: "444",
	//	ToUserId: "555",
	//	FriendsSince: "2016-05-15 00:21:45",
	//}
	//log.With(log.Fields{"friendID", friendId,).Info("in data getter **********")
	log.Debugf("in data GETter ********** friendID %s friendDomain %s", friendId, friendDomain)
	return &friendDomain, nil
}

func (d *friendData) List(status string, from_user_id string, to_user_id string) ([]*domain.Friendship, error) {
	log.Debugf("in data LISTer ********** status %s from_user_id %s to_user_id %s", status, from_user_id, to_user_id)

	var friendshipId, fromUserId, toUserId, friendsSince string
	query := ""
	if status == "active"{
		if from_user_id != "" {
			query = SelectActiveFriendshipsFromUser
		} else {
			query = SelectActiveFriendshipsToUser
		}
	} else if status == "pending" {
		if from_user_id != "" {
			query = SelectPendingFriendshipsFromUser
		} else {
			query = SelectPendingFriendshipsToUser
		}
	} else {
		log.Errorf("How do we do error handing??")
	}
	statement, _ := d.db.Prepare(query)
	rows, _ := statement.Query(from_user_id)
	//if err != nil {
	//	log.Fatal(err)
	//}
	friendDomainSlice := make([]*domain.Friendship, 0)

	for rows.Next() {
		rows.Scan(&friendshipId, &fromUserId, &toUserId, &friendsSince)
		//if err := rows.Scan(&name); err != nil {
		//	log.Fatal(err)
		//}
		friendDomain := domain.Friendship{
			Id: friendshipId,
			FromUserId: fromUserId,
			ToUserId: toUserId,
			FriendsSince: friendsSince,
			Status: "active",
		}
		friendDomainSlice = append(friendDomainSlice, &friendDomain)
	}

	return friendDomainSlice, nil
}
//
//func (d *mysqlMapper) Read(ctx context.Context, datapathId string) (datapath.Path, error) {
//
//	log.Debugf("MysqlMapper.Read %s", datapathId)
//	authUser := auth.FromContext(ctx)
//
//	if len(authUser.AccountLinks()) == 0 {
//		metrics.Inc("data.get_datapath.no_user", 1)
//		return datapath.Path{}, data.ErrNoValidUserProvided
//	}
//
//	datapathUserMatch := false
//	accountLinks := authUser.AccountLinks()
//	// Checking every account link for this user
//	for _, accountLink := range accountLinks {
//		// Check that the user is paired with this datapathId
//		var count int
//		log.Debugf("MysqlMapper.Read looking up match for %s - %s:%s", datapathId, accountLink.Domain, accountLink.DomainUserId)
//		_ = d.db.QueryRow(SelectDatapathUserSqlStr, datapathId, accountLink.Domain, accountLink.DomainUserId).Scan(&count)
//
//		if count > 0 {
//			datapathUserMatch = true
//			break
//		}
//	}
//
//	if !datapathUserMatch {
//		// User is not associated with this datapath
//		metrics.Inc("data.get_datapath.no_match", 1)
//		return datapath.Path{}, data.ErrDatapathUserComboDoesntExist
//	}
//	log.Debugf("Found datapath-user match for %s", datapathId)
//
//	datapathObj, err := d.selectDatapath(datapathId)
//	if err != nil {
//		return datapath.Path{}, err
//	}
//
//	metrics.Inc("data.get_datapath.success", 1)
//	return datapathObj, nil
//}
//
//func (d *mysqlMapper) selectDatapath(datapathId string) (datapath.Path, error) {
//	var flattened_datapath string
//	err := d.db.QueryRow(SelectStepSqlStr, datapathId).Scan(&flattened_datapath)
//	switch {
//	case err == sql.ErrNoRows:
//		return datapath.Path{}, apierrors.NotFound
//	case err != nil:
//		log.Errorf("Error Querying for Step: %v", err)
//		metrics.Inc("data.get_datapath.failure", 1)
//		return datapath.Path{}, err
//	}
//
//	log.Debugf("Successfully queried for Datapath, returned %v", flattened_datapath)
//
//	var flattenedPath datapath.FlattenedPath
//	json.Unmarshal([]byte(flattened_datapath), &flattenedPath)
//
//	path, err := flattenedPath.AsPath()
//	if err != nil {
//		log.Errorf("Error Converting FlattenedPath to Path: %v", err)
//		metrics.Inc("data.get_datapath.convert_failure", 1)
//		return datapath.Path{}, err
//	}
//
//	log.Debugf("Unmarshaled JSON into datapath.Path %v", path)
//	return path, nil
//}

//// itemData provides a non-working stub you can fill in.
//// For a working example of itemData implemented in-memory, look at mock_item_data.go
//type itemData struct {
//	//placeholder for examples sake, we don't really care if you use an SQL db
//	conn *sql.DB
//}

