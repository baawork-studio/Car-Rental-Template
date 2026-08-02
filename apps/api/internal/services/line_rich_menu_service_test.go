package services

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (function roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func TestCreateMenuUploadsImageToLineDataAPI(t *testing.T) {
	imagePath := filepath.Join(t.TempDir(), "menu.jpg")
	if err := os.WriteFile(imagePath, []byte("image"), 0o600); err != nil {
		t.Fatalf("create test image: %v", err)
	}
	var requests []*http.Request
	service := NewRichMenuService("token")
	service.client = &http.Client{Transport: roundTripperFunc(func(request *http.Request) (*http.Response, error) {
		requests = append(requests, request)
		body := "{}"
		if request.URL.Host == "api.line.me" {
			body = `{"richMenuId":"richmenu-id"}`
		}
		return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}, nil
	})}

	menuID, err := service.CreateMenu("Rental", "Rent", imagePath, "https://liff.line.me/id")
	if err != nil || menuID != "richmenu-id" {
		t.Fatalf("unexpected menu result: id=%q error=%v", menuID, err)
	}
	if len(requests) != 2 || requests[0].URL.Host != "api.line.me" || requests[1].URL.Host != "api-data.line.me" {
		t.Fatalf("unexpected request targets: %#v", requests)
	}
	if requests[0].Header.Get("Authorization") != "Bearer token" || requests[1].Header.Get("Content-Type") != "image/jpeg" {
		t.Fatalf("unexpected request headers: %#v %#v", requests[0].Header, requests[1].Header)
	}
}

func TestSetDefaultUsesMessagingAPI(t *testing.T) {
	service := NewRichMenuService("token")
	service.client = &http.Client{Transport: roundTripperFunc(func(request *http.Request) (*http.Response, error) {
		if request.URL.Host != "api.line.me" || request.URL.Path != "/v2/bot/user/all/richmenu/menu-id" {
			t.Fatalf("unexpected default-menu request: %s", request.URL)
		}
		return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: make(http.Header), Body: io.NopCloser(strings.NewReader(""))}, nil
	})}
	if err := service.SetDefault("menu-id"); err != nil {
		t.Fatalf("set default menu: %v", err)
	}
}
