package tenant

// Context contains all informations related to a well defined tenant
type Context interface {
	ID() string
	UserID() string
}
