package service

import "github.com/brunobandev/reppyapi/internal/domain/bill"

type BillService interface {
	Save(bill bill.Bill) (*bill.Bill, error)
	FindAll() (*[]bill.Bill, error)
	FindById(id uint64) (*bill.Bill, error)
	Delete(id uint64) error
	Update(bill bill.Bill) error
}

type billServiceImpl struct {
}

func NewBillService() BillService {
	return billServiceImpl{}
}

func (billServiceImpl) Save(bill bill.Bill) (*bill.Bill, error) {
	return nil, nil
}

func (billServiceImpl) FindAll() (*[]bill.Bill, error) {
	return nil, nil
}

func (billServiceImpl) FindById(id uint64) (*bill.Bill, error) {
	return nil, nil
}

func (billServiceImpl) Delete(id uint64) error {
	return nil
}

func (billServiceImpl) Update(bill bill.Bill) error {
	return nil
}
