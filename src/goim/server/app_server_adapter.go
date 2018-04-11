package server

import (
	"goim/rpc"
	"log"
)

type appServerAdapter struct {
	authenticated bool
	authenticator Authenticator
	storer        Storer
}

func newAppServerAdapter(authenticator Authenticator, storer Storer) rpc.AppServer {
	server := &appServerAdapter{}
	server.authenticator = authenticator
	server.storer = storer

	return server
}

func (s *appServerAdapter) Authenticate(username *string, password *string) (status int32) {
	authenticated, err := s.authenticator.Authenticate(username, password)
	if err != nil {
		log.Printf("authentication error: %v", err.Error())
		status = rpc.InternalErrorStatus
	}
	if authenticated {
		s.authenticated = true
		status = rpc.SuccessStatus
		return
	}

	status = rpc.WrongCredentialsStatus
	return
}

func (s *appServerAdapter) GetKeys() (res []*string, status int32) {
	return s.storer.GetKeys(), rpc.SuccessStatus
}

func (s *appServerAdapter) GetStr(key *string) (res *string, expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	res, expires = s.storer.GetStr(key)
	return
}

func (s *appServerAdapter) SetStr(key *string, str *string, ttl int64) (expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	expires = s.storer.SetStr(key, str, ttl)
	return
}

func (s *appServerAdapter) GetArr(key *string) (res []*string, expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	res, expires = s.storer.GetArr(key)
	return
}

func (s *appServerAdapter) SetArr(key *string, arr []*string, ttl int64) (expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	expires = s.storer.SetArr(key, arr, ttl)
	return
}

func (s *appServerAdapter) GetArrItem(key *string, index int32) (res *string, expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	res, expires = s.storer.GetArrItem(key, index)
	return
}

func (s *appServerAdapter) SetArrItem(key *string, index int32, str *string, ttl int64) (expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	expires = s.storer.SetArrItem(key, index, str, ttl)
	return
}

func (s *appServerAdapter) GetDict(key *string) (res map[string]*string, expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	res, expires = s.storer.GetDict(key)
	return
}

func (s *appServerAdapter) SetDict(key *string, dict map[string]*string, ttl int64) (expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	expires = s.storer.SetDict(key, dict, ttl)
	return
}

func (s *appServerAdapter) GetDictItem(key *string, subkey *string) (res *string, expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	res, expires = s.storer.GetDictItem(key, subkey)
	return
}

func (s *appServerAdapter) SetDictItem(key *string, subkey *string, str *string, ttl int64) (expires int64, status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	expires = s.storer.SetDictItem(key, subkey, str, ttl)
	return
}

func (s *appServerAdapter) RemoveEntry(key *string) (status int32) {
	if !s.authenticated {
		status = rpc.NotAuthenticatedStatus
		return
	}

	status = rpc.SuccessStatus
	s.storer.RemoveEntry(key)
	return
}
