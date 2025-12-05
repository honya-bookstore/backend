package http

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleStaff    UserRole = "staff"
	RoleCustomer UserRole = "customer"
)
