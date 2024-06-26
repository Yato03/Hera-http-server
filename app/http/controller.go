package http

import (
	"fmt"
	"strings"
)

type Controller interface {
	Handle(request Request) (Response, bool)
}

// Path: GET /echo/{string}
type HomeController struct {
	Path   string
	Method string
}

func (c *HomeController) Handle(request Request) (Response, bool) {
	if request.Path == c.Path && request.Method == c.Method {
		return NoBody(OK), true
	}
	return NoBody(NOT_FOUND), false
}

// Path: GET /echo/{string}
type EchoController struct {
	Path   string
	Method string
}

func (c *EchoController) Handle(request Request) (Response, bool) {
	if strings.HasPrefix(request.Path, c.Path) && request.Method == c.Method {
		directories := strings.Split(request.Path, "/")
		return TextPlain(directories[2], request, OK), true
	}
	return NoBody(NOT_FOUND), false
}

// Path: GET /user-agent
type UserAgentController struct {
	Path   string
	Method string
}

func (c *UserAgentController) Handle(request Request) (Response, bool) {
	if strings.HasPrefix(request.Path, c.Path) && request.Method == c.Method {
		return TextPlain(request.Headers["User-Agent"], request, OK), true
	}
	return NoBody(NOT_FOUND), false
}

// Path: GET /files/{string}
type GetFilesController struct {
	Path   string
	Method string
}

func (c *GetFilesController) Handle(request Request) (Response, bool) {
	if strings.HasPrefix(request.Path, c.Path) && request.Method == c.Method {
		directories := strings.Split(request.Path, "/")
		return GetFile(directories[2]), true
	}
	return NoBody(NOT_FOUND), false
}

// Path: POST /files/{string}
type UploadFileController struct {
	Path   string
	Method string
}

func (c *UploadFileController) Handle(request Request) (Response, bool) {
	if strings.HasPrefix(request.Path, c.Path) && request.Method == c.Method {
		directories := strings.Split(request.Path, "/")
		fmt.Println(directories[2])
		return UploadFile(directories[2], request.Body, CREATED), true
	}
	return NoBody(NOT_FOUND), false
}

// Path: GET /files/{string}
type GetHTMLController struct {
	Path   string
	Method string
}

func (c *GetHTMLController) Handle(request Request) (Response, bool) {
	if strings.HasPrefix(request.Path, c.Path) && request.Method == c.Method {
		directories := strings.Split(request.Path, "/")
		return GetHTML(directories[2]), true
	}
	return NoBody(NOT_FOUND), false
}
