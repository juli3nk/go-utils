package password

import (
	"unicode"

	"github.com/thoas/go-funk"
	"golang.org/x/crypto/bcrypt"
)

type ValidatePasswordOptions struct {
	MinLen     int
	MaxLen     int
	MinLower   int
	MinUpper   int
	MinNumber  int
	MinSpecial int
}

/*
	200: OK
	300: No password provided
	301: Password length is too short
	302: Password length is too long
	401: Not enough lowercase characters
	402: Not enough uppercase characters
	403: Not enough numbers
	404: Not enough special characters
*/
func ValidatePassword(p string, opts ValidatePasswordOptions) (bool, []int) {
	isSpecial := func(c string) bool {
		validChars := []string{
			"!",
			"@",
			"#",
			"$",
			"%",
			"*",
			"(",
			")",
			"-",
			"_",
		}

		if funk.Contains(validChars, c) {
			return true
		}

		return false
	}

	var (
		lower   int
		upper   int
		number  int
		special int

		codes []int
	)

	pL := len(p)
	if pL == 0 || pL < opts.MinLen || pL > opts.MaxLen {
		return false, []int{300}
	}

	sum := opts.MinLower + opts.MinUpper + opts.MinNumber + opts.MinSpecial
	if sum > opts.MinLen {
		return false, []int{301}
	}

	for _, c := range p {
		switch {
		case unicode.IsLower(c):
			lower++
		case unicode.IsUpper(c):
			upper++
		case unicode.IsNumber(c):
			number++
		case isSpecial(string(c)):
			special++
		}
	}

	if opts.MinLower > 0 && lower < opts.MinLower {
		codes = append(codes, 401)
	}
	if opts.MinUpper > 0 && upper < opts.MinUpper {
		codes = append(codes, 402)
	}
	if opts.MinNumber > 0 && number < opts.MinNumber {
		codes = append(codes, 403)
	}
	if opts.MinSpecial > 0 && special < opts.MinSpecial {
		codes = append(codes, 404)
	}
	if len(codes) > 0 {
		return false, codes
	}

	return true, []int{200}
}

func GeneratePassword(rawpassword string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rawpassword), 10)

	return string(hashedPassword)
}

func ComparePassword(rawpassword, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(password, rawpassword); err != nil {
		return false
	}

	return true
}
