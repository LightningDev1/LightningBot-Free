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
