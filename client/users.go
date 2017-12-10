// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": users Resource Client
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowUsersPath computes a request path to the show action of users.
func ShowUsersPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/develop/v1/users/%s", param0)
}

// アカウント情報取得
func (c *Client) ShowUsers(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUsersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUsersRequest create the request corresponding to the show action endpoint of the users resource.
func (c *Client) NewShowUsersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
