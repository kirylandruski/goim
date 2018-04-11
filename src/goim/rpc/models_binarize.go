
// THIS IS FILE IS GENERATE BY BINARIZE CODEGEN TOOL
// DO NOT EDIT!
		
package rpc
		
import "binarize"
import "errors"

func (s *LoginRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.StrSize(s.Username)+binarize.StrSize(s.Password))
	offset = binarize.IdPut(buf, offset, 1)
	
	offset = binarize.StrPut(buf, offset, s.Username)
	offset = binarize.StrPut(buf, offset, s.Password)

	return buf
}

func (s *LoginRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 1)
	if err != nil {
		return err
	}
	
	s.Username, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Password, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *LoginResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 2)
	
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *LoginResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 2)
	if err != nil {
		return err
	}
	
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetStrRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.StrSize(s.Key))
	offset = binarize.IdPut(buf, offset, 3)
	
	offset = binarize.StrPut(buf, offset, s.Key)

	return buf
}

func (s *GetStrRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 3)
	if err != nil {
		return err
	}
	
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetStrResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.StrSize(s.Str)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 4)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.StrPut(buf, offset, s.Str)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *GetStrResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 4)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Str, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetArrRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.StrSize(s.Key))
	offset = binarize.IdPut(buf, offset, 5)
	
	offset = binarize.StrPut(buf, offset, s.Key)

	return buf
}

func (s *GetArrRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 5)
	if err != nil {
		return err
	}
	
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetArrResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.ArrSize(s.Arr)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 6)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.ArrPut(buf, offset, s.Arr)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *GetArrResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 6)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Arr, offset, err = binarize.ArrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetArrItemRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.StrSize(s.Key)+binarize.Int32Size(s.Index))
	offset = binarize.IdPut(buf, offset, 7)
	
	offset = binarize.StrPut(buf, offset, s.Key)
	offset = binarize.Int32Put(buf, offset, s.Index)

	return buf
}

func (s *GetArrItemRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 7)
	if err != nil {
		return err
	}
	
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Index, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetArrItemResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.StrSize(s.Str)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 8)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.StrPut(buf, offset, s.Str)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *GetArrItemResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 8)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Str, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetDictRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.StrSize(s.Key))
	offset = binarize.IdPut(buf, offset, 9)
	
	offset = binarize.StrPut(buf, offset, s.Key)

	return buf
}

func (s *GetDictRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 9)
	if err != nil {
		return err
	}
	
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetDictResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.MapSize(s.Dict)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 10)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.MapPut(buf, offset, s.Dict)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *GetDictResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 10)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Dict, offset, err = binarize.MapRead(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetDictItemRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.StrSize(s.Key)+binarize.StrSize(s.Subkey))
	offset = binarize.IdPut(buf, offset, 11)
	
	offset = binarize.StrPut(buf, offset, s.Key)
	offset = binarize.StrPut(buf, offset, s.Subkey)

	return buf
}

func (s *GetDictItemRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 11)
	if err != nil {
		return err
	}
	
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Subkey, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetDictItemResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.StrSize(s.Str)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 12)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.StrPut(buf, offset, s.Str)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *GetDictItemResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 12)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Str, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetStrRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.TTL)+binarize.StrSize(s.Key)+binarize.StrSize(s.Str))
	offset = binarize.IdPut(buf, offset, 13)
	
	offset = binarize.Int64Put(buf, offset, s.TTL)
	offset = binarize.StrPut(buf, offset, s.Key)
	offset = binarize.StrPut(buf, offset, s.Str)

	return buf
}

func (s *SetStrRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 13)
	if err != nil {
		return err
	}
	
	s.TTL, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Str, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetStrResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 14)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *SetStrResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 14)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetArrRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.TTL)+binarize.StrSize(s.Key)+binarize.ArrSize(s.Arr))
	offset = binarize.IdPut(buf, offset, 15)
	
	offset = binarize.Int64Put(buf, offset, s.TTL)
	offset = binarize.StrPut(buf, offset, s.Key)
	offset = binarize.ArrPut(buf, offset, s.Arr)

	return buf
}

func (s *SetArrRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 15)
	if err != nil {
		return err
	}
	
	s.TTL, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Arr, offset, err = binarize.ArrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetArrResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 16)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *SetArrResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 16)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetArrItemRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.TTL)+binarize.StrSize(s.Key)+binarize.Int32Size(s.Index)+binarize.StrSize(s.Str))
	offset = binarize.IdPut(buf, offset, 17)
	
	offset = binarize.Int64Put(buf, offset, s.TTL)
	offset = binarize.StrPut(buf, offset, s.Key)
	offset = binarize.Int32Put(buf, offset, s.Index)
	offset = binarize.StrPut(buf, offset, s.Str)

	return buf
}

func (s *SetArrItemRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 17)
	if err != nil {
		return err
	}
	
	s.TTL, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Index, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}
	s.Str, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetArrItemResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 18)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *SetArrItemResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 18)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetDictRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.TTL)+binarize.StrSize(s.Key)+binarize.MapSize(s.Dict)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 19)
	
	offset = binarize.Int64Put(buf, offset, s.TTL)
	offset = binarize.StrPut(buf, offset, s.Key)
	offset = binarize.MapPut(buf, offset, s.Dict)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *SetDictRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 19)
	if err != nil {
		return err
	}
	
	s.TTL, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Dict, offset, err = binarize.MapRead(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetDictResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 20)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *SetDictResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 20)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetDictItemRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.TTL)+binarize.StrSize(s.Key)+binarize.StrSize(s.Subkey)+binarize.StrSize(s.Str))
	offset = binarize.IdPut(buf, offset, 21)
	
	offset = binarize.Int64Put(buf, offset, s.TTL)
	offset = binarize.StrPut(buf, offset, s.Key)
	offset = binarize.StrPut(buf, offset, s.Subkey)
	offset = binarize.StrPut(buf, offset, s.Str)

	return buf
}

func (s *SetDictItemRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 21)
	if err != nil {
		return err
	}
	
	s.TTL, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Subkey, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Str, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetDictItemResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int64Size(s.Expires)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 22)
	
	offset = binarize.Int64Put(buf, offset, s.Expires)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *SetDictItemResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 22)
	if err != nil {
		return err
	}
	
	s.Expires, offset, err = binarize.Int64Read(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *RemoveEntryRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.StrSize(s.Key))
	offset = binarize.IdPut(buf, offset, 23)
	
	offset = binarize.StrPut(buf, offset, s.Key)

	return buf
}

func (s *RemoveEntryRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 23)
	if err != nil {
		return err
	}
	
	s.Key, offset, err = binarize.StrRead(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *RemoveEntryResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 24)
	
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *RemoveEntryResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 24)
	if err != nil {
		return err
	}
	
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}

func (s *GetKeysRequest) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0))
	offset = binarize.IdPut(buf, offset, 25)
	

	return buf
}

func (s *GetKeysRequest) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 25)
	if err != nil {
		return err
	}
	

	return nil
}

func (s *GetKeysResponse) Serialize() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0)+binarize.ArrSize(s.Arr)+binarize.Int32Size(s.Status))
	offset = binarize.IdPut(buf, offset, 26)
	
	offset = binarize.ArrPut(buf, offset, s.Arr)
	offset = binarize.Int32Put(buf, offset, s.Status)

	return buf
}

func (s *GetKeysResponse) Deserialize(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, 26)
	if err != nil {
		return err
	}
	
	s.Arr, offset, err = binarize.ArrRead(buf, offset)
	if err != nil {
		return err
	}
	s.Status, offset, err = binarize.Int32Read(buf, offset)
	if err != nil {
		return err
	}

	return nil
}


func DynamicDeserialize(buf []byte) (Binarizer, error) {
	var err error
	offset := int32(0)
	id, offset, err := binarize.IdRead(buf, offset)
	if err != nil {
		return nil, err
	}

	switch id {
 	
	case 1:
		res := &LoginRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 2:
		res := &LoginResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 3:
		res := &GetStrRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 4:
		res := &GetStrResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 5:
		res := &GetArrRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 6:
		res := &GetArrResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 7:
		res := &GetArrItemRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 8:
		res := &GetArrItemResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 9:
		res := &GetDictRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 10:
		res := &GetDictResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 11:
		res := &GetDictItemRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 12:
		res := &GetDictItemResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 13:
		res := &SetStrRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 14:
		res := &SetStrResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 15:
		res := &SetArrRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 16:
		res := &SetArrResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 17:
		res := &SetArrItemRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 18:
		res := &SetArrItemResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 19:
		res := &SetDictRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 20:
		res := &SetDictResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 21:
		res := &SetDictItemRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 22:
		res := &SetDictItemResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 23:
		res := &RemoveEntryRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 24:
		res := &RemoveEntryResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 25:
		res := &GetKeysRequest{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	case 26:
		res := &GetKeysResponse{}
		err := res.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	
	default:
		return nil, errors.New("could not parse struct - unexpected struct id")
	}
}


type Binarizer interface {
	Serialize() []byte
	Deserialize(buf []byte) error
}
