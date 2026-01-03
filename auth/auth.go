package auth

import (
	"bybit/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func SIGN(message, secret string) string {
	s := hmac.New(sha256.New, []byte(secret))
	s.Write([]byte(message))
	return hex.EncodeToString(s.Sum(nil))
}

func GetAuth(path, query string) ([]byte, error) {
	ApiKey := config.C.ApiKey
	SecretKey := config.C.SecretKey

	if ApiKey == "" || SecretKey == "" {
		return nil, fmt.Errorf("Api/Secret is empty")
	}

	baseURL := config.C.BaseURL
	fullURL := baseURL + path + "?" + query

	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	window := "5000"
	message := timestamp + ApiKey + window + query
	signature := SIGN(message, SecretKey)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-BAPI-SIGN", signature)
	req.Header.Set("X-BAPI-API-KEY", ApiKey)
	req.Header.Set("X-BAPI-TIMESTAMP", timestamp)
	req.Header.Set("X-BAPI-RECV-WINDOW", window)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func PostAuth(path string, bodyData map[string]interface{}) ([]byte, error) {
	ApiKey := config.C.ApiKey
	SecretKey := config.C.SecretKey

	if ApiKey == "" || SecretKey == "" {
		return nil, fmt.Errorf("Api/secret is empty")
	}

	mar, _ := json.Marshal(bodyData)
	bodyString := string(mar)

	baseURL := config.C.BaseURL
	fullURL := baseURL + path

	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	window := "5000"
	message := timestamp + ApiKey + window + bodyString
	signature := SIGN(message, SecretKey)

	req, err := http.NewRequest("POST", fullURL, strings.NewReader(bodyString))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-BAPI-API-KEY", ApiKey)
	req.Header.Set("X-BAPI-TIMESTAMP", timestamp)
	req.Header.Set("X-BAPI-RECV-WINDOW", window)
	req.Header.Set("X-BAPI-SIGN", signature)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
