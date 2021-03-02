package storage

type UserStorage interface {
	All(sessionID string) map[string]string
	Save(sessionID string, username string, value string) error
}

type SessionStorage interface {
	GetVisibility(sessionID string) bool
	SetVisibility(sessionID string, state bool) error
}
