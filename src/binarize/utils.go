package binarize

import (
	"encoding/binary"
	"unsafe"
	"errors"
)

// Id

func IdSize(val int32) int32 {
	return Int32Size(val)
}

func IdPut(buf []byte, offset int32, val int32) (int32) {
	start, end := offset, offset+Int32Size(0)
	binary.BigEndian.PutUint32(buf[start:end], uint32(val))
	offset = end

	return offset
}

func IdRead(buf []byte, offset int32) (int32, int32, error) {
	start, end := offset, offset+Int32Size(0)
	if end > int32(len(buf)) {
		return 0, 0, errors.New("buf contains corrupted data")
	}
	val := binary.BigEndian.Uint32(buf[start:end])
	offset = end

	return int32(val), offset, nil
}

func IdValidate(buf []byte, offset int32, id int32) (int32, error) {
	val, offset, err := IdRead(buf, offset)
	if err != nil {
		return 0, err
	}
	if val != id {
		return 0, errors.New("buf has another struct Id")
	}

	return offset, nil
}

// int32

func Int32Size(val int32) int32 {
	return int32(unsafe.Sizeof(int32(val)))
}

func Int32Put(buf []byte, offset int32, val int32) (int32) {
	start, end := offset, offset+Int32Size(0)
	binary.BigEndian.PutUint32(buf[start:end], uint32(val))
	offset = end

	return offset
}

func Int32Read(buf []byte, offset int32) (int32, int32, error) {
	start, end := offset, offset+Int32Size(0)
	if end > int32(len(buf)) {
		return 0, 0, errors.New("buf contains corrupted data")
	}
	val := binary.BigEndian.Uint32(buf[start:end])
	offset = end

	return int32(val), offset, nil
}

// int64

func Int64Size(val int64) int32 {
	return int32(unsafe.Sizeof(int64(val)))
}

func Int64Put(buf []byte, offset int32, val int64) (int32) {
	start, end := offset, offset+Int64Size(val)
	binary.BigEndian.PutUint64(buf[start:end], uint64(val))
	offset = end

	return offset
}

func Int64Read(buf []byte, offset int32) (int64, int32, error) {
	start, end := offset, offset+Int64Size(0)
	if end > int32(len(buf)) {
		return 0, 0, errors.New("buf contains corrupted data")
	}
	val := binary.BigEndian.Uint64(buf[start:end])
	offset = end

	return int64(val), offset, nil
}

// str

func StrSize(val *string) int32 {
	if val == nil {
		return Int32Size(0)
	}

	return Int32Size(0) + int32(len(*val))
}

func StrPut(buf []byte, offset int32, val *string) int32 {
	if val == nil {
		return Int32Put(buf, offset, -1)
	}

	valLen := int32(len(*val))
	offset = Int32Put(buf, offset, int32(valLen))

	start, end := offset, offset+valLen
	copy(buf[start:end], *val)
	offset = end

	return offset
}

func StrRead(buf []byte, offset int32) (*string, int32, error) {
	valLen, offset, err := Int32Read(buf, offset)
	if err != nil {
		return nil, 0, err
	}
	if valLen == -1 {
		return nil, offset, nil
	}

	val := make([]byte, valLen)
	start, end := offset, offset+valLen
	if end > int32(len(buf)) {
		return nil, 0, errors.New("buf contains corrupted data")
	}

	copy(val, buf[start:end])
	offset = end

	str := string(val)

	return &str, offset, nil
}

// map

func MapSize(val map[string]*string) int32 {
	size := Int32Size(0)
	for k, v := range val {
		size += StrSize(&k)
		size += StrSize(v)
	}

	return size
}

func MapPut(buf []byte, offset int32, val map[string]*string) int32 {
	if val == nil {
		return Int32Put(buf, offset, -1)
	}

	valLen := int32(len(val))
	offset = Int32Put(buf, offset, valLen)

	for k, v := range val {
		offset = StrPut(buf, offset, &k)
		offset = StrPut(buf, offset, v)
	}

	return offset
}

func MapRead(buf []byte, offset int32) (map[string]*string, int32, error) {
	valLen, offset, err := Int32Read(buf, offset)
	if err != nil {
		return nil, 0, err
	}
	if valLen == -1 {
		return nil, offset, nil
	}
	if valLen == 0 {
		return make(map[string]*string), offset, nil
	}

	val := make(map[string]*string, valLen)
	var k, v *string

	for i := int32(0); i < valLen; i++ {
		if k, offset, err = StrRead(buf, offset); err != nil {
			return nil, 0, err
		}
		if v, offset, err = StrRead(buf, offset); err != nil {
			return nil, 0, err
		}

		val[*k] = v
	}

	return val, offset, nil
}

// arr

func ArrSize(val []*string) int32 {
	size := Int32Size(0)
	for _, v := range val {
		size += StrSize(v)
	}

	return size
}

func ArrPut(buf []byte, offset int32, val []*string) int32 {
	if val == nil {
		return Int32Put(buf, offset, -1)
	}

	valLen := int32(len(val))
	offset = Int32Put(buf, offset, valLen)
	for _, v := range val {
		offset = StrPut(buf, offset, v)
	}

	return offset
}

func ArrRead(buf [] byte, offset int32) ([]*string, int32, error) {
	valLen, offset, err := Int32Read(buf, offset)
	if err != nil {
		return nil, 0, err
	}
	if valLen == -1 {
		return nil, offset, nil
	}
	if valLen == 0 {
		return make([]*string, 0), offset, nil
	}

	val := make([]*string, valLen)
	for i := int32(0); i < valLen; i++ {
		val[i], offset, err = StrRead(buf, offset)
		if err != nil {
			return nil, 0, err
		}
	}

	return val, offset, nil
}
