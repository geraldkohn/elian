package jwt

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var hmacSecret = []byte("elian-sharing-secrectKey")

// 用uid, 当前时间签名token
func SignedToken(uid string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   uid,
		"year":  strconv.Itoa(now.Year()),
		"month": strconv.Itoa(int(now.Month())),
		"day":   strconv.Itoa(now.Day()),
		"hour":  strconv.Itoa(now.Hour()),
	})
	tokenString, err := token.SignedString(hmacSecret)

	if err != nil {
		log.Printf("error occurs when signing token: %v", err)
		return "", err
	}

	return tokenString, nil
}

// 解析token, 返回uid, err
func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("签名方式错误")
		}
		return hmacSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uid := claims["uid"].(string)
		year := claims["year"].(string)
		month := claims["month"].(string)
		day := claims["day"].(string)
		hour := claims["hour"].(string)
		if judgeTimeHasExpired(year, month, day, hour) {
			return "", errors.New("token 已过期")
		}
		return uid, err
	} else {
		return "", errors.New("token 解析错误")
	}
}

func judgeTimeHasExpired(year string, month string, day string, hour string) bool {
	yearInt, _ := strconv.ParseInt(year, 10, 64)
	monthInt, _ := strconv.ParseInt(month, 10, 64)
	dayInt, _ := strconv.ParseInt(day, 10, 64)
	hourInt, _ := strconv.ParseInt(hour, 10, 64)

	//期限48小时
	now := time.Now()
	if int(yearInt) != now.Year() || int(monthInt) != int(now.Month()) {
		return true //已过期
	}
	if (now.Day()*24+now.Hour())-int(dayInt*24+hourInt) >= 48 {
		return true //已过期
	} else {
		return false //未过期
	}
}

func FormatTime(t time.Time) string {
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	return strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(day) + " " + strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)
}
