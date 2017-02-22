package session

import(
	"net/http"

	"github.com/gorilla/sessions"
)

type Session struct {
	Name string
	AuthKey string
	Store *sessions.CookieStore

	User *map[string]string
}

func (s *Session) Init() {
	s.Name = "session-name"
	s.AuthKey = "something-very-secret"
	s.Store = sessions.NewCookieStore([]byte(s.AuthKey))
}

func (s *Session) SetUser(w http.ResponseWriter, r *http.Request, u *map[string]string) error {
	return nil
}

func (s *Session) GetUser(r *http.Request) (*map[string]string, error) {
	return nil, nil
}

func (s *Session) DeleteUser(w http.ResponseWriter, r *http.Request) (*map[string]string, error) {
	return nil, nil
}