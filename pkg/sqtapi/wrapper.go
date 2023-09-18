package sqtapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jamesits/go-meituansqt/pkg/sqtobfuscate"
	"io"
	"net/http"
	urllib "net/url"
	"strings"
	"time"
)

func WrappedApi[Tr, Ts any](httpMethod string, url string, method string) func(context.Context, *http.Client, *Config, *Tr) (*Ts, *http.Response, error) {
	return func(ctx context.Context, client *http.Client, sqt *Config, in *Tr) (*Ts, *http.Response, error) {
		u, err := urllib.ParseRequestURI(sqt.ApiEndpoint)
		if err != nil {
			return nil, nil, fmt.Errorf("sqt: %w", err)
		}
		u = u.JoinPath(url)

		obfs := sqtobfuscate.Obfuscator{SecretKey: sqt.SecretKey}

		i := &Request[Tr]{Data: in}
		i.Method = method
		i.Timestamp = time.Now().Unix()
		i.EnterpriseID = sqt.EnterpriseID
		content, err := json.Marshal(i)
		if err != nil {
			return nil, nil, fmt.Errorf("sqt: %w", err)
		}
		co, err := obfs.Obfuscate(string(content))
		if err != nil {
			return nil, nil, fmt.Errorf("sqt: %w", err)
		}
		body := urllib.Values{}
		body.Set("token", sqt.AccessKey)
		body.Set("version", sqt.Version)
		body.Set("content", co)

		req, err := http.NewRequestWithContext(ctx, httpMethod, u.String(), strings.NewReader(body.Encode()))
		if err != nil {
			return nil, nil, fmt.Errorf("sqt: %w", err)
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		if client == nil {
			client = http.DefaultClient
		}
		resp, err := client.Do(req)
		if err != nil {
			return nil, resp, fmt.Errorf("sqt: %w", err)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, resp, fmt.Errorf("sqt: %w", err)
		}

		ret := &Response[Ts]{}
		err = json.Unmarshal(b, ret)
		return ret.Data, resp, errors.Join(err, ret.Error())
	}
}
