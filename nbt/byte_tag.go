package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const byteTypeId byteType = 1

type byteType int8

type ByteTag struct {
	Value int8
}

func (_ byteType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt8()

	if err != nil {
		return nil, err
	}

	return ByteTag{Value: data}, nil
}

func (_ byteType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteTag)
	if !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return writer.writeInt8(data.Value)
}

func (_ byteType) GetId() int8 {
	return int8(byteTypeId)
}

func (_ ByteTag) getDataType() dataType {
	return byteTypeId
}

func (dtype byteType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(ByteTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	if value.Kind() == reflect.Bool {
		value.SetBool(data.Value != 0)
		return nil
	}

	err := RequireKind(value, reflect.Int8)
	if err != nil {
		return err
	}

	value.SetInt(int64(data.Value))

	return nil
}