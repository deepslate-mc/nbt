package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const listTypeId listType = 9

type listType int8

type ListTag struct {
	dataType dataType
	Value    []Tag
}

func (_ listType) Read(reader Reader) (Tag, error) {
	dtype, err := reader.readInt8()

	if err != nil {
		return nil, err
	}

	dataType, err := getDataType(dtype)

	if err != nil {
		return nil, err
	}

	length, err := reader.readInt32()

	if err != nil {
		return nil, err
	}

	list := make([]Tag, length)

	for i := range list {
		data, err := dataType.Read(reader)

		if err != nil {
			return nil, fmt.Errorf("unable to read list at index %d. Reason: %w", i, err)
		}

		list[i] = data
	}

	return ListTag{
		dataType: dataType,
		Value:    list,
	}, nil
}

func (_ listType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ListTag)

	if !ok {
		return errors.New("incompatible tag. Expected LIST")
	}

	if err := writer.writeInt8(data.dataType.GetId()); err != nil {
		return fmt.Errorf("unable to write list datatype. Reason: %w", err)
	}

	for i, value := range data.Value {
		if err := data.dataType.Write(writer, value); err != nil {
			return fmt.Errorf("unable to write list at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

func (_ listType) GetId() int8 {
	return int8(listTypeId)
}

func (_ ListTag) getDataType() dataType {
	return listTypeId
}

func (dtype listType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(ListTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	if err := RequireKind(value, reflect.Slice); err != nil {
		return err
	}

	if value.IsNil() {
		value.Set(reflect.MakeSlice(value.Type(), 0, 0))
	}

	value.SetCap(len(data.Value))
	value.SetLen(len(data.Value))

	for i, v := range data.Value {
		if err := v.getDataType().Decode(v, value.Index(i)); err != nil {
			return err
		}
	}

	return nil
}