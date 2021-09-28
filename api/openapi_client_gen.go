// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	serverURL string
	http      HTTPClient
}

func NewClient(serverURL string) *Client {
	return &Client{
		serverURL: serverURL,
		http: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (c *Client) FoobarGet(ctx context.Context, params FoobarGetParams) (FoobarGetResponse, error) {
	path := c.serverURL
	path += "/foobar"

	r, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	q := r.URL.Query()
	{
		s := conv.Int64ToString(params.Query.InlinedParam)
		q.Set("inlinedParam", s)
	}
	{
		s := conv.Int32ToString(params.Query.Skip)
		q.Set("skip", s)
	}
	r.URL.RawQuery = q.Encode()

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := DecodeFoobarGetResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) FoobarPut(ctx context.Context) (*FoobarPutDefault, error) {
	path := c.serverURL
	path += "/foobar"

	r, err := http.NewRequestWithContext(ctx, "PUT", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := DecodeFoobarPutResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) FoobarPost(ctx context.Context, req *Pet) (FoobarPostResponse, error) {
	body, contentType, err := EncodeFoobarPostRequest(req)
	if err != nil {
		return nil, err
	}

	path := c.serverURL
	path += "/foobar"

	r, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	r.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := DecodeFoobarPostResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetGet(ctx context.Context, params PetGetParams) (PetGetResponse, error) {
	path := c.serverURL
	path += "/pet"

	r, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	{
		value := conv.StringArrayToString(params.Header.XScope)
		for _, v := range value {
			r.Header.Add("x-scope", v)
		}
	}

	q := r.URL.Query()
	{
		s := conv.Int64ToString(params.Query.PetID)
		q.Set("petID", s)
	}
	r.URL.RawQuery = q.Encode()

	{
		value := conv.StringToString(params.Cookie.Token)
		r.AddCookie(&http.Cookie{
			Name:   "token",
			Value:  value,
			MaxAge: 1337,
		})
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := DecodePetGetResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetCreate(ctx context.Context, req PetCreateRequest) (*Pet, error) {
	body, contentType, err := EncodePetCreateRequest(req)
	if err != nil {
		return nil, err
	}

	path := c.serverURL
	path += "/pet"

	r, err := http.NewRequestWithContext(ctx, "POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	r.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := DecodePetCreateResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *Client) PetGetByName(ctx context.Context, params PetGetByNameParams) (*Pet, error) {
	path := c.serverURL
	path += "/pet"
	{
		value := conv.StringToString(params.Path.Name)
		path += "/" + value
	}

	r, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.http.Do(r)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	result, err := DecodePetGetByNameResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}