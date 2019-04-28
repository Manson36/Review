package service

import "github.com/review/project08/model"

type CustomerService struct {
	customers []model.Customer
	customerNum int
}