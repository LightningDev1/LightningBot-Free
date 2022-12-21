package http

import (
	"github.com/valyala/fasthttp"
)

type HttpResponse struct {
	Success       bool
	Error         error
	Body          string
	BodyBytes     []byte
	StatusCode    int
	Headers       string
	InnerResponse *fasthttp.Response
}

func errorResponse(err error) *HttpResponse {
	return &HttpResponse{
		Success:       false,
		Error:         err,
		Body:          "",
		BodyBytes:     nil,
		StatusCode:    0,
		Headers:       "",
		InnerResponse: nil,
	}
}

func Get(url string, headersList ...map[string]string) *HttpResponse {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)

	for _, headers := range headersList {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	err := fasthttp.Do(req, resp)
	if err != nil {
		return errorResponse(err)
	}

	return &HttpResponse{
		Success:       true,
		Error:         nil,
		Body:          string(resp.Body()),
		BodyBytes:     resp.Body(),
		StatusCode:    resp.StatusCode(),
		Headers:       resp.Header.String(),
		InnerResponse: resp,
	}
}

func CustomMethodWithData(url, method, data string, headersList ...map[string]string) *HttpResponse {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.SetBody([]byte(data))

	for _, headers := range headersList {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	err := fasthttp.Do(req, resp)
	if err != nil {
		return errorResponse(err)
	}

	return &HttpResponse{
		Success:       true,
		Error:         nil,
		Body:          string(resp.Body()),
		BodyBytes:     resp.Body(),
		StatusCode:    resp.StatusCode(),
		Headers:       resp.Header.String(),
		InnerResponse: resp,
	}
}

func Post(url string, data string, headersList ...map[string]string) *HttpResponse {
	return CustomMethodWithData(url, "POST", data, headersList...)
}

func Patch(url string, data string, headersList ...map[string]string) *HttpResponse {
	return CustomMethodWithData(url, "PATCH", data, headersList...)
}

func GetDiscordHeaders(token string) map[string]string {
	return map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "",
		"accept-language":    "en-GB,q=0.9",
		"cookie":             "__cfduid=d7e8d2784592da39fb3f621664b9aede51620414171; __dcfduid=24a543339247480f9b0bb95c710ce1e6",
		"referer":            "https://discord.com/",
		"origin":             "https://discord.com",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9003 Chrome/91.0.4472.164 Electron/13.4.0 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDAzIiwib3NfdmVyc2lvbiI6IjEwLjAuMjIwMDAiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLUdCIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTA5MTkwLCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==",
		"content-type":       "application/json",
		"authorization":      token,
	}
}
