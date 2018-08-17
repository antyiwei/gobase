package M03_abstract_factory

import "fmt"

/* ======================== golang 中抽象工厂模式 ======================== */
type GirlFriend struct {
	nationality string
	eyesColor   string
	language    string
}

type AbstractFactory interface {
	CreateMyLove() GirlFriend
}

type IndianGirlFriendFactory struct {
}

type KoreanGirlFriendFactory struct {
}

func (a IndianGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Indian", "Black", "Hindi"}
}

func (a KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Korean", "Brown", "Korean"}
}

func GetGirlFriend(typeGf string) GirlFriend {
	var gffact AbstractFactory
	switch typeGf {
	case "Indian":
		gffact = IndianGirlFriendFactory{}
		return gffact.CreateMyLove()
	case "Korean":
		gffact = KoreanGirlFriendFactory{}
		return gffact.CreateMyLove()
	}
	return GirlFriend{}
}

/* ======================== golang 中抽象工厂模式 2 ======================== */

// OrderMainDao 为订单主记录
type OrderMainDao interface {
	SaveOrderMain()
}

// OrderDetailDao 为订单详情记录
type OrderDetailDao interface {
	SaveOrderDetail()
}

// DaoFactory DAO 抽象模式工厂模式
type DaoFactory interface {
	CreateOrderMainDAO() OrderMainDao
	CreateOrderDetailDAO() OrderDetailDao
}

// RDBMainDao 为关系型数据库的OrderMainDAO实现
type RDBMainDao struct {
}

// SaveOrderDetail...
func (*RDBMainDao) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// RDBDetailDao 为关系型数据库的OrderDetailDao实现
type RDBDetailDao struct {
}

// SaveOrderDetail ...
func (*RDBDetailDao) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// RDBDAOFactory 是RDB 抽象工厂实现
type RDBDaoFactory struct {
}

func (*RDBDaoFactory) CreateOrderMainDAO() OrderMainDao {
	return &RDBMainDao{}
}

func (*RDBDaoFactory) CreateOrderDetailDAO() OrderDetailDao {
	return &RDBDetailDao{}
}

//XMLDetailDAO XML存储
type XMLMainDao struct {
}

//SaveOrderDetail...
func (*XMLMainDao) SaveOrderMain() {
	fmt.Print("xml main save")
}

//XMLDetailDao XML存储
type XMLDetailDao struct{}

// SaveOrderDetail ...
func (*XMLDetailDao) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

// XMLDAOFactory 是RDB抽象工厂实现
type XMLDAOFactory struct {
}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDao {
	return &XMLMainDao{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDao {
	return &XMLDetailDao{}
}
