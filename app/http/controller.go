package http

import "strings"

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
