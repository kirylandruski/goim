package rpc

import (
	"testing"
	"reflect"
)

var lorem = []string{
	"Lorem ipsum dolor sit amet",
	"consectetur adipiscing elit",
	"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
	"Ut enim ad minim veniam",
	"quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat",
	"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur",
	"",
}

func testSerializer(src Binarizer, dst Binarizer, t *testing.T) {
	buf := src.Serialize()
	err := dst.Deserialize(buf)
	if err != nil {
		t.Errorf("deserialization error %v", err.Error())
	}

	if !reflect.DeepEqual(src, dst) {
		t.Errorf("deserialized model missmatch initial model\n%v\n%v", src, dst)
	}
}

func TestLoginRequest(t *testing.T) {
	src := LoginRequest{Username: &lorem[0], Password: &lorem[1]}
	dst := LoginRequest{}

	testSerializer(&src, &dst, t)
}

func TestLoginRequestEmpty(t *testing.T) {
	src := LoginRequest{Username: &lorem[6], Password: nil}
	dst := LoginRequest{}

	testSerializer(&src, &dst, t)
}

func TestLoginResponse(t *testing.T) {
	src := LoginResponse{Status: int32(5)}
	dst := LoginResponse{}
	testSerializer(&src, &dst, t)
}

func TestLoginResponseEmpty(t *testing.T) {
	src := LoginResponse{Status: int32(0)}
	dst := LoginResponse{}
	testSerializer(&src, &dst, t)
}

func TestSetDictItemRequest(t *testing.T) {
	src := SetDictItemRequest{
		TTL: 5000000000,
		Key: &lorem[1],
		Str: &lorem[2],
	}
	dst := SetDictItemRequest{}
	testSerializer(&src, &dst, t)
}

func BenchmarkActionSetDict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := SetDictItemRequest{
			TTL: 5000,
			Key: &lorem[1],
			Str: &lorem[2],
		}
		dst := SetDictItemRequest{}

		buf := src.Serialize()
		err := dst.Deserialize(buf)
		if err != nil {
			b.Errorf("deserialization error %v", err.Error())
		}
	}
}

func BenchmarkGenericDeserialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := SetDictItemRequest{
			TTL: 6746252,
			Key: &lorem[1],
			Str: &lorem[2],
		}
		buf := src.Serialize()

		res, err := DynamicDeserialize(buf)
		if err != nil {
			b.Fatal(err.Error())
		}

		switch res.(type) {
		case *SetDictItemRequest:
			// do nothing
		default:
			b.Fatal("wrong deserialization result")
		}

	}
}
