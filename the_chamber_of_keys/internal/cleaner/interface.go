package ttl

// Cleaner: operations for the background cleaner
type Cleaner interface {
	Start()
	Stop()
}
