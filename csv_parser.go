package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Val interface{}
type DataRow []Val

type CSVParser struct {
	DateFmt      string
	NaNs         []string
	Delim        rune
	BlockText    rune
	IgnoreBlocks bool
}

type DataFrame struct {
	Data    []DataRow
	Columns []string
	n_rows  int32
	n_cols  int32
}

// for printing out to std,
func GetInterfaceValue(value Val) string {
	switch v := value.(type) {
	case int:
		// fmt.Println("VALUEEEEE", v)
		return strconv.Itoa(v)
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return v.(string)
	}
}

func (df DataFrame) head() {
	for _, cname := range df.Columns {
		fmt.Printf("|%*s", 16, cname)
	}
	println("|")

	for i, row := range df.Data {
		if i >= 5 {
			break
		}
		for _, val := range row {
			fmt.Printf("|%*s", 16, GetInterfaceValue(val))
		}
		println("|")
	}
}

func NewCSVParser() *CSVParser {
	return &CSVParser{
		DateFmt:      "2006-01-02",
		NaNs:         []string{"NA", "N/A", "", "nan", "NaN"},
		Delim:        ',',
		BlockText:    '"',
		IgnoreBlocks: false,
	}
}

func ReadCSV(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (parser *CSVParser) splitLineItems(line string) []string {
	var items []string
	var item strings.Builder

	block_text := false // current segment in/out of block text

	for _, char := range line {
		if parser.IgnoreBlocks {
			if char == parser.Delim {
				items = append(items, item.String())
				item.Reset()
			} else {
				item.WriteRune(char)
			}
		} else {
			if char == parser.BlockText {
				block_text = !block_text
			} else if char == parser.Delim && !block_text {
				items = append(items, item.String())
				item.Reset()
			} else {
				item.WriteRune(char)
			}
		}

	}

	items = append(items, item.String())
	return items
}

func (parser *CSVParser) parseValue(item string) Val {
	item = strings.TrimSpace(item)
	item = strings.Trim(item, string(parser.Delim))

	for _, nan_repr := range parser.NaNs {
		if strings.EqualFold(item, nan_repr) {
			return math.NaN()
		}
	}

	int_val, err := strconv.Atoi(item)
	if err == nil {
		return int_val
	}

	float_val, err := strconv.ParseFloat(item, 64)
	if err == nil {
		return float_val
	}

	date_val, err := time.Parse(parser.DateFmt, item)
	if err == nil {
		return date_val
	}

	return item
}

func (parser *CSVParser) parseLine(line string) DataRow {
	items := parser.splitLineItems(line)
	data_row := make(DataRow, len(items))

	for i, item := range items {
		data_row[i] = parser.parseValue(item)
	}
	return data_row
}

func (parser *CSVParser) ParseCSV(csv_file *os.File) (DataFrame, error) {
	f_reader := bufio.NewReader(csv_file)
	var data []DataRow

	defer csv_file.Close()

	for {
		line, err := f_reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return DataFrame{}, err
		}

		line = strings.TrimSpace(line)
		if line != "" {
			parsed_line := parser.parseLine(line)
			data = append(data, parsed_line)
		}

		if err == io.EOF {
			break
		}
	}

	var cols []string
	for _, item := range data[0] {
		cols = append(cols, item.(string))
	}

	parsed_csv := DataFrame{
		Data:    data[1:],
		Columns: cols,
		n_cols:  int32(len(data[0])),
		n_rows:  int32(len(data)),
	}
	return parsed_csv, nil
}

func main() {
	data, err := ReadCSV("test.csv")

	parser := NewCSVParser()
	csv, err := parser.ParseCSV(data)
	if err != nil {
		fmt.Println("Error Parsing CSV", err)
	}

	csv.head()
	// fmt.Println(csv)
}
