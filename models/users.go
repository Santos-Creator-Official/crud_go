package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Users struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string        `bson:"username" json:"username" binding:"required"`
	Password  string        `bson:"password" json:"password" binding:"required"`
	Name      string        `bson:"name" json:"name" binding:"required"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}

//type Users struct {
//	ID        int64     `json:"id" gorm:"column:id"`
//	Username  string    `json:"username" gorm:"column:username"`
//	Password  string    `json:"password" gorm:"column:password"`
//	Name      string    `json:"name" gorm:"column:name"`
//	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
//}

//func (m *Users) TableName() string {
//	return "users"
//}
//
//type UsersDao struct {
//	sourceDB  *gorm.DB
//	replicaDB []*gorm.DB
//	m         *Users
//}
//
//func NewUsersDao(ctx context.Context, dbs ...*gorm.DB) *UsersDao {
//	dao := new(UsersDao)
//	switch len(dbs) {
//	case 0:
//		panic("database connection required")
//	case 1:
//		dao.sourceDB = dbs[0]
//		dao.replicaDB = []*gorm.DB{dbs[0]}
//	default:
//		dao.sourceDB = dbs[0]
//		dao.replicaDB = dbs[1:]
//	}
//	return dao
//}
//
//
//func (d *UsersDao) Create(ctx context.Context, obj *Users) error {
//	err := d.sourceDB.Model(d.m).Create(&obj).Error
//	if err != nil {
//		return fmt.Errorf("UsersDao: %w", err)
//	}
//	return nil
//}
//
//func (d *UsersDao) Get(ctx context.Context, fields, where string) (*Users, error) {
//	items, err := d.List(ctx, fields, where, 0, 1)
//	if err != nil {
//		return nil, fmt.Errorf("UsersDao: Get where=%s: %w", where, err)
//	}
//	if len(items) == 0 {
//		return nil, gorm.ErrRecordNotFound
//	}
//	return &items[0], nil
//}
//
//func (d *UsersDao) List(ctx context.Context, fields, where string, offset, limit interface{}) ([]Users, error) {
//	var results []Users
//	err := d.replicaDB[rand.Intn(len(d.replicaDB))].Model(d.m).
//		Select(fields).Where(where).Offset(offset).Limit(limit).Find(&results).Error
//	if err != nil {
//		return nil, fmt.Errorf("UsersDao: List where=%s: %w", where, err)
//	}
//	return results, nil
//}
//
//func (d *UsersDao) Update(ctx context.Context, where string, update map[string]interface{}, args ...interface{}) error {
//	err := d.sourceDB.Model(d.m).Where(where, args...).
//		Updates(update).Error
//	if err != nil {
//		return fmt.Errorf("UsersDao:Update where=%s: %w", where, err)
//	}
//	return nil
//}
//
//func (d *UsersDao) Delete(ctx context.Context, where string, args ...interface{}) error {
//	if len(where) == 0 {
//		return gorm.ErrInvalidSQL
//	}
//	if err := d.sourceDB.Where(where, args...).Delete(d.m).Error; err != nil {
//		return fmt.Errorf("UsersDao: Delete where=%s: %w", where, err)
//	}
//	return nil
//}
