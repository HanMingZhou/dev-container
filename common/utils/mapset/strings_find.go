package mapset

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Split 把字符串s按照给定的分隔符sep进行分割返回字符串切片
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func SplitLeft(s, sep string) (result string) {
	strArr := strings.Split(s, sep)
	result = strArr[len(strArr)-1]

	return
}

// InSlice 判断字符串是否在 slice 中。
func InSlice(items []string, item string) (bool, string) {
	for _, eachItem := range items {
		if strings.Contains(eachItem, item) {
			pass := SplitLeft(eachItem, "=")
			return true, pass
		}
	}
	return false, ""
}

/**
 * @Author Flamingo
 * @Description //查询字符串切片中否包含某一个字符串
 * @Date 2023/2/1 10:48
 **/
func IsHave(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	//index的取值：0 ~ (len(str_array)-1)
	return index < len(str_array) && str_array[index] == target
}

/**
 * @Author Flamingo
 * @Description //生成随机字符串
 * @Date 2023/2/1 10:55
 **/

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitBytes  = "0123456789"
)

func GenerateRandomString(length int) (string, error) {
	charSet := letterBytes + digitBytes
	charSetLength := len(charSet)
	result := strings.Builder{}

	// 第一位是字母
	randomIndex, err := crand.Int(crand.Reader, big.NewInt(int64(len(letterBytes))))
	if err != nil {
		return "", err
	}
	result.WriteByte(letterBytes[randomIndex.Int64()])

	// 至少一个数字和一个字母
	randomIndex, err = crand.Int(crand.Reader, big.NewInt(int64(len(digitBytes))))
	if err != nil {
		return "", err
	}
	result.WriteByte(digitBytes[randomIndex.Int64()])

	// 生成剩余的字符
	for i := 2; i < length; i++ {
		randomIndex, err := crand.Int(crand.Reader, big.NewInt(int64(charSetLength)))
		if err != nil {
			return "", err
		}
		result.WriteByte(charSet[randomIndex.Int64()])
	}

	return result.String(), nil
}
