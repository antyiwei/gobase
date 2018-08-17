package M03_abstract_factory

import (
	"testing"
)

// TestAbstractFactroy 测试1
func TestAbstractFactroy(t *testing.T) {
	a := GetGirlFriend("Indian")

	if &a == nil {
		t.Fatal(" fail ")
	} else {
		t.Log(a)
	}
}

func getMainAndDetail(factory DaoFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()

}

func TestCreateOrderDAO(t *testing.T) {
	var factory DaoFactory
	factory = &RDBDaoFactory{}
	getMainAndDetail(factory)
}

func TestXmlFactory(t *testing.T) {
	var factory DaoFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}
