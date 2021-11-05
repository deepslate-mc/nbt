package nbt

import "errors"

const byteTypeId byteType = 1

type byteType int8

type ByteTag struct {
	value int8
}

func (_ byteType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt8()

	if err != nil {
		return nil, err
	}

	return ByteTag{value: data}, nil
}

func (_ byteType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteTag)
	if !ok {
		return errors.New("incompatible tag. Expected END")
	}

	return writer.writeInt8(data.value)
}

func (_ byteType) GetId() int8 {
	return int8(byteTypeId)
}