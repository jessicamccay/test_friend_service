package data
//
//import (
//	"friends.uacf.io/domain"
//
//	"go.uacf.io/apierrors"
//)
//
//func NewMemoryData() ItemData {
//	data := &memoryData{
//		lastId: 1,
//		store:  map[int64]*domain.Item{},
//	}
//	data.initialize()
//	return data
//}
//
//type memoryData struct {
//	lastId int64
//	store  map[int64]*domain.Item
//}
//
//func (mock *memoryData) NextId() int64 {
//	mock.lastId++
//	return mock.lastId
//}
//
//func (mock *memoryData) Create(item *domain.Item) (*domain.Item, error) {
//	item.Id = mock.NextId()
//	mock.store[item.Id] = item
//	return item, nil
//}
//
//func (mock *memoryData) Get(id int64) (*domain.Item, error) {
//	if item, ok := mock.store[id]; ok {
//		return item, nil
//	}
//	return nil, apierrors.NotFound
//}
//
//func (mock *memoryData) Put(id int64, item *domain.Item) error {
//	mock.store[item.Id] = item
//	return nil
//}
//
//func (mock *memoryData) List() (items []*domain.Item, err error) {
//	for _, item := range mock.store {
//		items = append(items, item)
//	}
//	return items, nil
//}
//
//// add some example data so we aren't forced to create it all.
//func (mock *memoryData) initialize() {
//	for i := 1; i <= 10; i++ {
//		mock.Create(&domain.Item{AuthorId: int64(i), Name: "Foo"})
//	}
//}
