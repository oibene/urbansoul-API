package models

type GetCustomerInput struct {
	Customer_id int
}

type GetProductInput struct {
	Product_id int
	Orders_id  int
}

type GetCommentsInput struct {
	Customer_id int
	Product_id  int
}
