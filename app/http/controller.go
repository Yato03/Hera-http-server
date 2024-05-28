package http

import (
	"fmt"
	"strings"
)

type Controller interface {
	Handle(request Request) (Response, bool)
}

// Path: /echo/{string}
type HomeController struct {
	Path   string
	Method string
}

func (c *HomeController) Handle(request Request) (Response, bool) {
	if request.Path == c.Path && request.Method == c.Method {
		fmt.Println(request.Path)
		return OK("", request), true
	}
	return NOT_FOUND(), false
}

// Path: /echo/{string}
type EchoController struct {
	Path   string
	Method string
}

func (c *EchoController) Handle(request Request) (Response, bool) {
	if strings.HasPrefix(request.Path, c.Path) && request.Method == c.Method {
		directories := strings.Split(request.Path, "/")
		return OK(directories[2], request), true
	}
	return NOT_FOUND(), false
}

// Path: /user-agent
type UserAgentController struct {
	Path   string
	Method string
}

func (c *UserAgentController) Handle(request Request) (Response, bool) {
	if strings.HasPrefix(request.Path, c.Path) && request.Method == c.Method {
		return OK(request.Headers["User-Agent"], request), true
	}
	return NOT_FOUND(), false
}
