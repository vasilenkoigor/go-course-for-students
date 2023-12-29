package file_copier

import (
	"errors"
	"flag"
	"strings"
)

type TextTransformation string

const (
	UpperCase  TextTransformation = "upper_case"
	LowerCase                     = "lower_case"
	TrimSpaces                    = "trim_spaces"
)

func (receiver TextTransformation) String() string {
	return string(receiver)
}

type Options struct {
	From      string
	To        string
	Offset    int64
	Limit     int64
	BlockSize int64
	Conv      []TextTransformation
}

type FlagsParser interface {
	Parse() (*Options, error)
}

//////////////////////////////////////////////////////////////

type UnixCmdFlagsParser struct{}

func (receiver UnixCmdFlagsParser) Parse() (*Options, error) {
	var from = flag.String("from", "", "file to read. by default - stdin")
	var to = flag.String("to", "", "file to write. by default - stdout")
	var offset = flag.Int64("offset", 0, "the number of bytes inside the input that must be skipped when copying. by default = 0")
	var limit = flag.Int64("limit", 0, "the maximum number of bytes to read. By default, we copy all contents starting with -offset")
	var blockSize = flag.Int64("block_size", 0, "the size of one block in bytes when reading and writing. That is, you can neither read nor write more bytes than -block-size at a time;")
	var conv = flag.String("conv", "", "one or more of the possible transformations over the text, separated by a comma. possible transformations: upper_case, lower_case, trim_spaces")
	flag.Parse()

	var transformations []TextTransformation
	if *conv != "" {
		var separatedConv = strings.Split(*conv, ",")
		for _, s := range separatedConv {
			transformation := TextTransformation(s)
			switch transformation {
			case UpperCase, LowerCase, TrimSpaces:
				transformations = append(transformations, transformation)
			default:
				var formattedError = "unavailable transformation " + s + " in 'conv' parameter"
				return nil, errors.New(formattedError)
			}
		}
	}

	var opts = Options{
		From:      *from,
		To:        *to,
		Offset:    *offset,
		Limit:     *limit,
		BlockSize: *blockSize,
		Conv:      transformations,
	}

	return &opts, nil
}
