package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClient struct {
	BaseURL    string
	httpClient *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 15 * time.Second, // Таймаут, чтобы приложение не зависало, если Python упал
		},
	}
}

// GetJSON делает GET-запрос и парсит ответ в target
func (c *APIClient) GetJSON(endpoint string, target interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка сети при запросе к %s: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("ошибка API (статус %d): %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// PostJSON делает POST-запрос с JSON-телом и парсит ответ в target
func (c *APIClient) PostJSON(endpoint string, payload interface{}, target interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("ошибка сериализации запроса: %w", err)
	}

	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ошибка сети при запросе к %s: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("ошибка API (статус %d): %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
