package nbt

import (
	"io"
)


func Read(reader io.Reader) (Tag, error) {
	_, tag, err := NewReader(reader).Read()
	return tag, err
}

func Write(writer io.Writer, tag Tag) error {
	return NewWriter(writer).Write("", tag)
}

//TODO: Unmarshal and Marshal using Reflection
/*
func Unmarshal(reader io.Reader, v interface{}) error {
	for {
		dataType, err := readByte(reader)
		if err != nil {
			return err
		}

		if dataType(dataType) == End {
			return nil
		}

		name, err := readString(reader)
		if err != nil {
			return err
		}

		field, found := reflect.TypeOf(v).
		if !found {
			return fmt.Errorf("field %s not found", name)
		}

		field.Tag.Get("nbt")

		switch dataType(dataType) {
			case End: return nil
			case Byte: {
				readByte(reader)
			}

		}
	}
}

func Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}*/