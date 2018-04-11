package rpc

import (
	"errors"
)

var UnknownRequestError = errors.New("passed request does not match any known request")

type requestResolver struct {
	server AppServer
}

func NewRequestResolver(server AppServer) Resolver {
	return &requestResolver{server: server}
}

func (r *requestResolver) Resolve(req Binarizer) (Binarizer, error) {
	switch req := req.(type) {
	case *LoginRequest:
		return &LoginResponse{Status: r.server.Authenticate(req.Username, req.Password)}, nil
	case *GetKeysRequest:
		arr, status := r.server.GetKeys()
		return &GetKeysResponse{Arr: arr, Status: status}, nil
	case *GetStrRequest:
		res, expires, status := r.server.GetStr(req.Key)
		return &GetStrResponse{Expires: expires, Str: res, Status: status}, nil
	case *GetArrRequest:
		res, expires, status := r.server.GetArr(req.Key)
		return &GetArrResponse{Expires: expires, Arr: res, Status: status}, nil
	case *GetArrItemRequest:
		res, expires, status := r.server.GetArrItem(req.Key, req.Index)
		return &GetArrItemResponse{Expires: expires, Str: res, Status: status}, nil
	case *GetDictRequest:
		res, expires, status := r.server.GetDict(req.Key)
		return &GetDictResponse{Expires: expires, Dict: res, Status: status}, nil
	case *GetDictItemRequest:
		res, expires, status := r.server.GetDictItem(req.Key, req.Subkey)
		return &GetDictItemResponse{Expires: expires, Str: res, Status: status}, nil
	case *SetStrRequest:
		expires, status := r.server.SetStr(req.Key, req.Str, req.TTL)
		return &SetStrResponse{Expires: expires, Status: status}, nil
	case *SetArrRequest:
		expires, status := r.server.SetArr(req.Key, req.Arr, req.TTL)
		return &SetArrResponse{Expires: expires, Status: status}, nil
	case *SetArrItemRequest:
		expires, status := r.server.SetArrItem(req.Key, req.Index, req.Str, req.TTL)
		return &SetArrItemResponse{Expires: expires, Status: status}, nil
	case *SetDictRequest:
		expires, status := r.server.SetDict(req.Key, req.Dict, req.TTL)
		return &SetDictResponse{Expires: expires, Status: status}, nil
	case *SetDictItemRequest:
		expires, status := r.server.SetDictItem(req.Key, req.Subkey, req.Str, req.TTL)
		return &SetDictItemResponse{Expires: expires, Status: status}, nil
	default:
		return nil, UnknownRequestError
	}
}
