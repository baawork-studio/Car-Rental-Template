package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const (
	lineAPI     = "https://api.line.me/v2/bot"
	lineDataAPI = "https://api-data.line.me/v2/bot"
)

type RichMenuService struct {
	token  string
	client *http.Client
}
type richMenu struct {
	Size        richMenuSize   `json:"size"`
	Selected    bool           `json:"selected"`
	Name        string         `json:"name"`
	ChatBarText string         `json:"chatBarText"`
	Areas       []richMenuArea `json:"areas"`
}
type richMenuSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type richMenuArea struct {
	Bounds richMenuBounds `json:"bounds"`
	Action richMenuAction `json:"action"`
}
type richMenuBounds struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
type richMenuAction struct {
	Type string `json:"type"`
	URI  string `json:"uri"`
}

func NewRichMenuService(token string) *RichMenuService {
	return &RichMenuService{token: token, client: &http.Client{}}
}
func (s *RichMenuService) CreateMenu(name, chatBarText, imagePath, liffURL string) (string, error) {
	menu := richMenu{Size: richMenuSize{Width: 2500, Height: 843}, Name: name, ChatBarText: chatBarText, Areas: []richMenuArea{{Bounds: richMenuBounds{Width: 2500, Height: 843}, Action: richMenuAction{Type: "uri", URI: liffURL}}}}
	body, _ := json.Marshal(menu)
	response, err := s.request(http.MethodPost, "/richmenu", "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode >= 300 {
		return "", lineError(response)
	}
	var created struct {
		RichMenuID string `json:"richMenuId"`
	}
	if err := json.NewDecoder(response.Body).Decode(&created); err != nil {
		return "", err
	}
	if err := s.uploadImage(created.RichMenuID, imagePath); err != nil {
		return "", err
	}
	return created.RichMenuID, nil
}
func (s *RichMenuService) SetDefault(menuID string) error {
	response, err := s.request(http.MethodPost, "/user/all/richmenu/"+menuID, "", nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode >= 300 {
		return lineError(response)
	}
	return nil
}
func (s *RichMenuService) uploadImage(menuID, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	response, err := s.dataRequest(http.MethodPost, "/richmenu/"+menuID+"/content", mime.TypeByExtension(filepath.Ext(path)), file)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode >= 300 {
		return lineError(response)
	}
	return nil
}
func (s *RichMenuService) request(method, path, contentType string, body io.Reader) (*http.Response, error) {
	return s.requestTo(lineAPI, method, path, contentType, body)
}
func (s *RichMenuService) dataRequest(method, path, contentType string, body io.Reader) (*http.Response, error) {
	return s.requestTo(lineDataAPI, method, path, contentType, body)
}
func (s *RichMenuService) requestTo(baseURL, method, path, contentType string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, baseURL+path, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+s.token)
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	return s.client.Do(request)
}
func lineError(response *http.Response) error {
	body, _ := io.ReadAll(response.Body)
	return fmt.Errorf("LINE API error (%s): %s", response.Status, string(body))
}
