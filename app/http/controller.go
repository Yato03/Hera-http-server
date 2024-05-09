package http

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/fileUtils"
)

func Controller(request Request) Response {
	var response Response

	directories := strings.Split(request.Path, "/")

	response = root(directories[1:], request)

	return response
}

func root(directories []string, request Request) Response {
	if len(directories) == 1 && directories[0] == "" {
		return OK("")
	} else if len(directories) >= 1 {
		if directories[0] == "echo" {
			return echo(directories[1:])
		}
		if directories[0] == "user-agent" {
			return userAgent(request)
		}
		if directories[0] == "files" {
			return files(directories[1:])
		}
	}
	return NOT_FOUND()

}

func echo(directories []string) Response {
	if len(directories) == 0 {
		return OK("ECHO!!")
	} else if len(directories) >= 1 {
		return OK(directories[0])
	}
	return NOT_FOUND()
}

func userAgent(request Request) Response {
	if request.Headers["User-Agent"] != "" {
		return OK(request.Headers["User-Agent"])
	}
	return BAD_REQUEST("User-Agent not found")
}

func files(directories []string) Response {
	if len(directories) >= 1 {
		content, err := fileUtils.ReadFile(directories[0])
		if err != nil {
			fmt.Println(err)
			return NOT_FOUND()
		}
		return OK_FILE(content)
	}
	return NOT_FOUND()
}
