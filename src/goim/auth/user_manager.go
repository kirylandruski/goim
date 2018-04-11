package auth

import (
	"golang.org/x/crypto/bcrypt"
	"encoding/base64"
	"encoding/json"
	"sync"
	"io"
	"io/ioutil"
)

type UserManager struct {
	sync.RWMutex
	// map storing username and it's hashed password
	usernames map[string]string
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (s *UserManager) Read(reader io.Reader) error {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	if len(buf) == 0 {
		s.usernames = make(map[string]string)
	} else {
		err := json.Unmarshal(buf, &s.usernames);
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *UserManager) Write(writer io.Writer) error {
	buf, err := json.Marshal(&s.usernames);
	if err != nil {
		return err
	}

	_, err = writer.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserManager) GetUsernames() []string {
	s.RLock()
	defer s.RUnlock()

	keys := make([]string, 0, len(s.usernames))
	for key := range s.usernames {
		keys = append(keys, key)
	}

	return keys
}

func (s *UserManager) SetUser(username string, password string) error {
	s.Lock()
	defer s.Unlock()

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.usernames[username] = base64.StdEncoding.EncodeToString(hashed)

	return nil
}

func (s *UserManager) RemoveUser(username string) {
	s.Lock()
	defer s.Unlock()

	delete(s.usernames, username)
}

func (s *UserManager) Authenticate(username *string, password *string) (bool, error) {
	s.RLock()
	b64hashed, ok := s.usernames[*username]
	s.RUnlock()

	if !ok {
		return false, nil
	}

	hashed, err := base64.StdEncoding.DecodeString(b64hashed)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(hashed, []byte(*password))
	if err != nil {
		return false, err
	}

	return true, nil
}
