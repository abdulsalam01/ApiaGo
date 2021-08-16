package utils

type Crud interface {
	GetAll()
	GetById()
	Create()
	Update()
	Delete()
}
