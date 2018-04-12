//go:generate binarize -source models.go -destination models_binarize.go -import_binarize goim/binarize -serialize_fname Serialize -deserialize_fname Deserialize

package rpc

// binarize id 1
type LoginRequest struct {
	// binarize order 0
	Username *string
	// binarize order 1
	Password *string
}

// binarize id 2
type LoginResponse struct {
	// binarize order 0
	Status int32
}

// binarize id 3
type GetStrRequest struct {
	// binarize order 0
	Key *string
}

// binarize id 4
type GetStrResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Str *string
	// binarize order 2
	Status int32
}

// binarize id 5
type GetArrRequest struct {
	// binarize order 0
	Key *string
}

// binarize id 6
type GetArrResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Arr []*string
	// binarize order 2
	Status int32
}

// binarize id 7
type GetArrItemRequest struct {
	// binarize order 0
	Key *string
	// binarize order 1
	Index int32
}

// binarize id 8
type GetArrItemResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Str *string
	// binarize order 2
	Status int32
}

// binarize id 9
type GetDictRequest struct {
	// binarize order 0
	Key *string
}

// binarize id 10
type GetDictResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Dict map[string]*string
	// binarize order 2
	Status int32
}

// binarize id 11
type GetDictItemRequest struct {
	// binarize order 0
	Key *string
	// binarize order 1
	Subkey *string
}

// binarize id 12
type GetDictItemResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Str *string
	// binarize order 2
	Status int32
}

// binarize id 13
type SetStrRequest struct {
	// binarize order 0
	TTL int64
	// binarize order 1
	Key *string
	// binarize order 2
	Str *string
}

// binarize id 14
type SetStrResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Status int32
}

// binarize id 15
type SetArrRequest struct {
	// binarize order 0
	TTL int64
	// binarize order 1
	Key *string
	// binarize order 2
	Arr []*string
}

// binarize id 16
type SetArrResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Status int32
}

// binarize id 17
type SetArrItemRequest struct {
	// binarize order 0
	TTL int64
	// binarize order 1
	Key *string
	// binarize order 2
	Index int32
	// binarize order 3
	Str *string
}

// binarize id 18
type SetArrItemResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Status int32
}

// binarize id 19
type SetDictRequest struct {
	// binarize order 0
	TTL int64
	// binarize order 1
	Key *string
	// binarize order 2
	Dict map[string]*string
	// binarize order 3
	Status int32
}

// binarize id 20
type SetDictResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Status int32
}

// binarize id 21
type SetDictItemRequest struct {
	// binarize order 0
	TTL int64
	// binarize order 1
	Key *string
	// binarize order 2
	Subkey *string
	// binarize order 3
	Str *string
}

// binarize id 22
type SetDictItemResponse struct {
	// binarize order 0
	Expires int64
	// binarize order 1
	Status int32
}

// binarize id 23
type RemoveEntryRequest struct {
	// binarize order 0
	Key *string
}

// binarize id 24
type RemoveEntryResponse struct {
	// binarize order 0
	Status int32
}

// binarize id 25
type GetKeysRequest struct {
}

// binarize id 26
type GetKeysResponse struct {
	// binarize order 0
	Arr []*string
	// binarize order 1
	Status int32
}
