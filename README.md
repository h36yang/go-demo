# go-demo
Practice materials for Go

### Language Characteristics
* Fast Compilation
* Fully Compiled
* Strongly Typed
* Concurrent by Default
* Garbage Collected
* Simplicity as a Core Value

### What Go is good at?
* **Web Services**
* **Web Applications**
* Task Automation
* GUI / Thick Client
* Machine Learning

### Sample Web Service Project Structure
```
module "demo/webservice"
    package "models"
        "user.go":  User Model
    package "controllers"
        "front.go": Front Controller that routes all HTTP requests to Back Controllers
        "user.go":  User Controller
    "main.go": Program entry point that starts up the web server
```
