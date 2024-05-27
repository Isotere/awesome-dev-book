package main

import (
	"encoding/json"
	"time"

	"math/rand"
	"regexp"

	"strconv"
	"testing"
)

type Data struct {
	Rank int    `json:"rank"`
	Body string `json:"body"`
}

type DataArray []Data

var jsonData string = ""

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())

	dataArray := make(DataArray, 100)
	for i := 0; i < len(dataArray); i++ {
		dataArray[i] = Data{
			Rank: i,
			Body: RandStringRunes(rand.Intn(1024)),
		}
	}

	for i := len(dataArray) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		dataArray[i], dataArray[j] = dataArray[j], dataArray[i]
	}

	jsonDataBytes, err := json.Marshal(dataArray)
	if err != nil {
		panic(err)
	}
	jsonData = string(jsonDataBytes)
}

func BenchmarkRegexpJSON(b *testing.B) {
	b.ResetTimer()

	r := regexp.MustCompile("\"rank\":\\s*(\\d)")

	for i := 0; i < b.N; i++ {
		matchs := r.FindAllStringSubmatch(jsonData, -1)
		values := make([]int, len(matchs))

		for z, value := range matchs {
			var err error
			values[z], err = strconv.Atoi(value[1])
			if err != nil {
				panic("value error")
			}
		}

		if len(matchs) != len(values) {
			panic("value len error")
		}
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var dataArray DataArray
		if err := json.Unmarshal([]byte(jsonData), &dataArray); err != nil {
			panic(err)
		}

		ranks := make([]int, len(dataArray))
		for i, v := range dataArray {
			ranks[i] = v.Rank
		}
	}
}
