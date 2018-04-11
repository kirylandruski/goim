package rpc

import (
	"github.com/golang/mock/gomock"
	"testing"
	"log"
	"reflect"
)

func ptr(str string) *string {
	return &str
}

func TestResolveLogin(t *testing.T) {
	mock := NewMockAppServer(gomock.NewController(t))
	resolver := NewRequestResolver(mock)

	mock.EXPECT().Authenticate(ptr("Username"), ptr("Password")).Return(int32(0))

	request := &LoginRequest{Username: ptr("Username"), Password: ptr("Password")}
	response, err := resolver.Resolve(request)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := LoginResponse{Status: int32(0)}
	if !reflect.DeepEqual(response, &expected) {
		log.Fatal("response missmatch")
	}
}

func TestResolveGetStr(t *testing.T) {
	mock := NewMockAppServer(gomock.NewController(t))
	resolver := NewRequestResolver(mock)

	mock.EXPECT().GetStr(ptr("Key")).Return(ptr("Str"), int64(5), int32(8))

	request := &GetStrRequest{Key: ptr("Key")}
	response, err := resolver.Resolve(request)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := GetStrResponse{Expires: int64(5), Str: ptr("Str"), Status: int32(8)}
	if !reflect.DeepEqual(response, &expected) {
		log.Fatal("response missmatch")
	}
}

func TestResolveGetArr(t *testing.T) {
	mock := NewMockAppServer(gomock.NewController(t))
	resolver := NewRequestResolver(mock)

	mock.EXPECT().GetArr(ptr("Key")).Return([]*string{ptr("Value1"), ptr("Value2")}, int64(5), int32(7))

	request := &GetArrRequest{Key: ptr("Key")}
	response, err := resolver.Resolve(request)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := GetArrResponse{Expires: int64(5), Arr: []*string{ptr("Value1"), ptr("Value2")}, Status: int32(7)}
	if !reflect.DeepEqual(response, &expected) {
		log.Fatal("response missmatch")
	}
}

func TestResolveSetStr(t *testing.T) {
	mock := NewMockAppServer(gomock.NewController(t))
	resolver := NewRequestResolver(mock)

	mock.EXPECT().SetStr(ptr("Key"), ptr("Str"), int64(12)).Return(int64(112), int32(6))

	request := &SetStrRequest{TTL: int64(12), Key: ptr("Key"), Str: ptr("Str")}
	response, err := resolver.Resolve(request)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := SetStrResponse{Expires: int64(112), Status: int32(6)}
	if !reflect.DeepEqual(response, &expected) {
		log.Fatal("response missmatch")
	}
}
