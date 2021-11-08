package nbt

import "errors"

const endTypeId endType = 0

type endType int8

type endTag struct{}

func (end endType) Read(reader Reader) (Tag, error) {
	return endTag{}, nil
}

func (_ endType) Write(writer Writer, tag Tag) error {
	if _, ok := tag.(endTag); !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return nil
}

func (_ endType) GetId() int8 {
	return int8(endTypeId)
}

func (_ endTag) getDataType() dataType {
	return endTypeId
}