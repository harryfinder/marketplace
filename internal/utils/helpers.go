package utils

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	phone_lib "github.com/nyaruka/phonenumbers"
	"github.com/spf13/viper"
	"log"
	"marketplace/internal/database"
	"marketplace/internal/models"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
	salt     = "sjakfslkaf23j213123kjklkjl"
)

var (
	SigningKey = viper.GetString("token_password")
	numberSet  = "0123456789"
)

//type entity struct {
//	database database.Database
//}

type tokenClaims struct {
	jwt.StandardClaims
	UserId     int64 `json:"user_id"`
	Role       int64 `json:"role"`
	Authorized bool  `json:"authorized"`
	// Exp        string `json:"exp"`
}
type entity struct {
	database database.Database
}

// ParseToken ...
func (e *entity) ParseToken(accessToken string) (*models.User, *jwt.Token, error) {

	var (
		err error
		ctx context.Context
	)
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(SigningKey), nil
	})
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, nil, err
	}

	userOld, err := e.database.GetUsersByID(ctx, claims.UserId)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	if userOld.Role != claims.Role {
		log.Println(err)
		return nil, nil, err
	}

	return userOld, token, nil
}

// Validate ...
func Validate(args map[string]interface{}) error {

	for key, value := range args {

		typeValue := fmt.Sprintf("%T", value)

		strValue := fmt.Sprintf("%v", value)
		strValue = strings.TrimSpace(strValue)

		if typeValue == "string" {
			if strValue == "" {
				return errors.New("поля " + key + " пустой")
			}
		} else {
			if value == int64(0) || value == int32(0) || value == int(0) || value == float64(0) || value == float32(0) {
				return errors.New("поля " + key + " пустой")
			}
		}

		if key == "login" {
			if strValue == "" {
				return errors.New("поля " + key + " пустой")
			}
			//res := strings.Split(strValue, "@")
			//if len(res) != 1 {
			//	if !strings.Contains(strValue, "@aliftech.net") && !strings.Contains(strValue, "@team.alif.tj") && !strings.Contains(strValue, "@alif.tj") && !strings.Contains(strValue, "@gmail.com") {
			//		return app.ErrBadEmail // email address is wrong
			//	}
			//}
		}
		//if key == "email" {
		//	if !strings.Contains(strValue, "@aliftech.net") && !strings.Contains(strValue, "@team.alif.tj") && !strings.Contains(strValue, "@alif.tj") && !strings.Contains(strValue, "@gmail.com") {
		//		return app.ErrBadEmail // email address is wrong
		//	}
		//}

		if key == "phone" {
			if len(strValue) > 12 || len(strValue) < 12 {
				return ErrConflictNumber
			}
			numberPhone, err := IsPhone2(strValue)
			if !err {
				return errors.New("номер неправильно")
			}
			err = ProviderDetector(numberPhone)
			if !err {
				return errors.New("номер неправильно")
			}

		}

		if key == "password" || key == "old_password" || key == "new_password" || key == "confirm_new_password" {
			if len(strValue) < 4 {
				return errors.New("пароль неправильно") // password is wrong
			}
		}

		// if key == "phone" {
		// 	if len(strValue) < 9 {
		// 		return errors.New("номер телефона неправильный") // phone is wrong
		// 	}
		// }https://github.com/AlifTech-golang/stock.git

		// if key == "status" {

		// 	intStr, err := strconv.Atoi(strValue)
		// 	if err != nil {
		// 		return errors.New("статус не может быть string")
		// 	}

		// 	if intStr < 1 && intStr >= 10 {
		// 		return errors.New("статус неправильно") // status is wrong
		// 	}
		// }
	}

	return nil
}

// GeneratePasswordHash ...
func GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func ProviderDetector(phone string) bool {
	match, err := regexp.Match(models.TcellTJ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.MegafonTJ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.ZetMobileTJ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.BabilonTJ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.OMobileTJ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.BeelineRU, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.MtsRU, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.MegafonRU, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.Tele2RU, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.Uzmobile, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.UselUZ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.BeelineUZ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.MobiUz, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.HumansUZ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	match, err = regexp.Match(models.PerfectumUZ, []byte(phone))
	if err != nil {
		return false
	}
	if match {
		return true
	}

	return false
}
func IsPhone2(phoneNum string) (phone string, ok bool) {

	phoneNum = GetDigits(phoneNum)

	if StrEmpty(phoneNum) {
		return "", false
	}

	if len(phoneNum) == 12 {
		phoneNum = "+" + phoneNum
	}

	if phoneNum[0] != '+' {
		phoneNum = "+" + phoneNum
	}

	pars, err := phone_lib.Parse(phoneNum, "")
	if err != nil {
		err = errors.New("Cant parse phone: " + phoneNum)
		return
	}
	countryCode := phone_lib.GetRegionCodeForNumber(pars)
	if countryCode == "" {
		ok = false
		return
	}
	ok = true
	phone = strings.ReplaceAll(phoneNum, " ", "")
	phone = strings.ReplaceAll(phone, "+", "")

	return
}

func StrEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func GetDigits(str string) string {
	numRegex := regexp.MustCompile("[0-9]+")
	arrNum := numRegex.FindAllString(str, -1)
	if arrNum == nil {
		return ""
	}

	return strings.Join(arrNum, "")
}

// RandomPasswordInt Функция генерирует рандомный числовой пароль получая длину число
func RandomPasswordInt(lenPassword int) string {
	var s string
	for j := 0; j < lenPassword; j++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		s += strconv.Itoa(r1.Intn(10))
	}
	return s
}

// RandomPasswordIntStr Функция генерирует рандомный числовой пароль получая длину число
func RandomPasswordIntStr(lenPassword int) string {
	var (
		letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		numberSet   = "123456789"
		password    strings.Builder
	)

	for j := 0; j < lenPassword/2; j++ {
		for i := 0; i < 1; i++ {
			random := rand.Intn(len(numberSet))
			randomString := letterRunes[rand.Intn(len(letterRunes))]
			password.WriteString(string(numberSet[random]))
			password.WriteString(string(randomString))
		}
	}
	return password.String()
}

// Worker функция для
func Worker(d time.Duration, f func() (int, error)) {
	var reEntranceFlag int64
	for range time.Tick(d * time.Second) {
		go func() {
			if atomic.CompareAndSwapInt64(&reEntranceFlag, 0, 1) {
				defer atomic.StoreInt64(&reEntranceFlag, 0)
			} else {
				log.Println("Previous worker in process now")
				return
			}
			f()
		}()
	}
}
