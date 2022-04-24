package tmptst

import (
	"fmt"
	"regexp"
	"testing"
)

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validate("2020-03-04")
	}
}

func validate(datestr string) error {
	var year, month, day int
	if _, err := fmt.Sscanf(datestr, "%d-%d-%d", &year, &month, &day); err != nil {
		return err
	}
	if year < 1970 || year > 3000 || month < 0 || month > 12 || day < 0 || day > 31 {
		return fmt.Errorf("invalid date string %s", datestr)
	}
	return nil
}

func BenchmarkValidate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validate2("2020-03-04")
	}
}

const dateRegexFormat = `\d{4}-\d{2}-\d{2}`

func validate2(datestr string) error {
	matches, err := regexp.MatchString(dateRegexFormat, datestr)
	if err != nil {
		return err
	}
	if !matches {
		return fmt.Errorf("incorrect date: %s, does not match required format of YYYY-MM-dd", datestr)
	}
	return nil
}
