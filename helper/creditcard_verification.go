package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/nehaal10/CardGaurd/utils"
)

func Cvv(num string) bool {
	if len(num) == 3 {
		for _, i := range num {
			ch := fmt.Sprintf("%c", i)
			_, err := strconv.Atoi(ch)
			isErr := utils.CheckErrReturn(err)
			if !isErr {
				return false
			}
		}
		return true
	}
	return false
}

func DateVerify(date string) bool {
	if len(date) < 5 || len(date) > 5 {
		return false
	}

	expiry := strings.Split(date, "/")

	expiryMonth, err := strconv.Atoi(expiry[0])
	utils.CheckError(err)
	if expiryMonth > 12 {
		return false
	}
	expiryYear, err := strconv.Atoi(expiry[1])
	utils.CheckError(err)

	currTime := time.Now()
	currYear := currTime.Year() % 100
	currMonth := int(currTime.Month())

	if expiryYear > currYear+3 {
		return false
	}

	//fmt.Println(currMonth, expiryMonth)

	//check month and year
	if currYear == expiryYear {
		return currMonth <= expiryMonth
	}
	return currYear < expiryYear
}

func Number(cn string) bool {
	var ss string
	cNumber := strings.Replace(cn, "-", " ", -1)
	for _, r := range cNumber {
		if !unicode.IsSpace(r) {
			ss += string(r)
		}
	}
	var sum int64 = 0
	parity := len(ss) % 2

	cardNumWithoutChecksum := ss[:len(ss)-1]

	for i, v := range cardNumWithoutChecksum {

		item, err := strconv.Atoi(string(v))

		if err != nil {
			fmt.Println(err)
			return false
		}
		if int64(i)%2 != int64(parity) {
			sum += int64(item)
		} else if item > 4 {
			sum += int64(2*item - 9)
		} else {
			sum += int64(2 * item)
		}

	}

	checkDigit, err := strconv.Atoi(ss[len(ss)-1:])

	if err != nil {
		fmt.Println(err)
		return false
	}

	SumMod := sum % 10

	if SumMod == int64(0) {
		return SumMod == int64(checkDigit)
	}
	return int64(10)-SumMod == int64(checkDigit)
}
