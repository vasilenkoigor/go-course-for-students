package file_copier

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type FileCopier interface {
	Copy(args Options) error
}

type UnixDDCopier struct{}

func (receiver UnixDDCopier) Copy(options Options) error {
	var textBuffer []byte

	if options.From != "" {
		buf, err := readFile(options)
		if err != nil {
			return err
		}
		textBuffer = buf
	} else {
		_, err := fmt.Scanln(&textBuffer)
		if err != nil {
			return err
		}
	}

	textBuffer = applyTextTransformations(textBuffer, options.Conv)

	if options.To != "" {
		err := writeFile(textBuffer, options.To)
		if err != nil {
			return err
		}
	} else {
		_, _ = fmt.Fprintln(os.Stderr, string(textBuffer))
	}

	return nil
}

func readFile(options Options) ([]byte, error) {
	file, err := os.Open(options.From)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Seek(options.Offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, options.Limit)
	text, err := file.Read(buf[:cap(buf)])
	buf = buf[:text]
	if err != nil {
		return nil, err
	}

	return buf, err
}

func writeFile(buff []byte, to string) error {
	file, err := os.Create(to)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err2 := file.Write(buff)
	if err2 != nil {
		return err2
	}

	return nil
}

func applyTextTransformations(bytes []byte, transformations []TextTransformation) []byte {
	str := string(bytes)

	for _, transformation := range transformations {
		switch transformation {
		case UpperCase:
			str = strings.ToUpper(str)
		case LowerCase:
			str = strings.ToLower(str)
		case TrimSpaces:
			str = strings.TrimSpace(str)
		}
	}

	return []byte(str)
}
