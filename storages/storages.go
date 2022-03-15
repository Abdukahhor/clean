package storages

type DB interface {
	// check error, if pgx.ErrNoRows returns true
	IsNotFound(error) bool
	//
	IsNotAffected(error) bool

	IsExists(err error) bool

	// close database
	Close()
}
