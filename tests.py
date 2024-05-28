import requests
from pwn import log

# Types
class Status:
    def __init__(self, counter, total, p):
        self.counter = counter
        self.total = total
        self.p = p
    
    def status(self, message):
        self.p.status(f"{message} [{self.counter}/{self.total}]")

# Tests
def test_user_agent(status):
    status.status("user-agent")
    user_agent = "Mozilla/5.0"
    response = requests.get("http://localhost:4221/user-agent", headers={"User-Agent": user_agent})
    check_condition(response.status_code == 200, "Status code is not 200")
    check_condition(response.text == user_agent, "The response is not the user-agent sent in the request")

def test_echo(status):
    status.status("echo")
    response = requests.get("http://localhost:4221/echo/Hello")
    check_condition(response.status_code == 200, "Status code is not 200")
    check_condition(response.text == "Hello", "The response is not Hello")

def test_not_found(status):
    status.status("not-found")
    response = requests.get("http://localhost:4221/not-found")
    check_condition(response.status_code == 404, "Status code is not 404")

def get_file(status):
    status.status("get-file")
    response = requests.get("http://localhost:4221/files/file")
    check_condition(response.status_code == 200, "Status code is not 200")
    check_condition("This is a file" in response.text, "The response is not 'This is a file'")

def post_file(status):
    status.status("post-file")
    response = requests.post("http://localhost:4221/files/prueba", data="Hello")
    check_condition(response.status_code == 201, "Status code is not 201")

# Utils
def check_server():
    try:
        response = requests.get("http://localhost:4221")
        return response.status_code == 200
    except:
        return False

def check_condition(condition, message):
    if not condition:
        log.failure("Test failed: " + message)
        exit()

# Run tests
def run_tests():
    tests = [test_user_agent, test_echo, test_not_found, get_file, post_file]
    counter = 1
    total = len(tests)
    p = log.progress("Testing:")

    for test in tests:
        test(Status(counter, total, p))
        counter += 1
    p.success("All tests passed")


if __name__ == "__main__":
    if not check_server():
        log.failure("Server is not running")
        exit()
    log.info("Server is running")
    log.info("Executing tests")
    run_tests()

    exit()