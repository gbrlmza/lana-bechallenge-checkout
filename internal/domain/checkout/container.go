package checkout

/*
	The container has all the external functionality required by the business logic(domain).
	Provides an abstraction to the domain, because the business logic doesn't(and shouldn't)
	known where the data is stored to or retrieved from or witch external services is accessed.
	This allow us to change databases, services or any external provider without affecting
	the business model.

	Also allows us to test the domain logic regardless of the specific repository implementations.
*/

type Container struct {
	Storage Storage
	Locker  Locker
}

type Storage interface {
}

type Locker interface {
}
