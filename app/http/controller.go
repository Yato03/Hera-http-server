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
		return OK("", request)
	} else if len(directories) >= 1 {
		if directories[0] == "echo" && request.Method == "GET" {
			return echo(directories[1:], request)
		}
		if directories[0] == "user-agent" && request.Method == "GET" {
			return userAgent(request)
		}
		if directories[0] == "files" && request.Method == "GET" {
			return getFiles(directories[1:])
		}
		if directories[0] == "files" && request.Method == "POST" {
			return postFiles(directories[1:], request)
		}
	}
	return NOT_FOUND()

}

func echo(directories []string, request Request) Response {
	if len(directories) == 0 {
		return OK("ECHO!!", request)
	} else if len(directories) >= 1 {
		return OK(directories[0], request)
	}
	return NOT_FOUND()
}

func userAgent(request Request) Response {
	if request.Headers["User-Agent"] != "" {
		return OK(request.Headers["User-Agent"], request)
	}
	return BAD_REQUEST("User-Agent not found")
}

func getFiles(directories []string) Response {
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

func postFiles(directories []string, request Request) Response {
	if len(directories) >= 1 {
		err := fileUtils.WriteFile(directories[0], request.Body)
		if err != nil {
			fmt.Println(err)
			return BAD_REQUEST("")
		}
		return CREATED()
	}
	return BAD_REQUEST("")
}
