# Interface in Gin
Gin uses interface pattern in many positions, even though it still has something limitations on interface 
usage.

## Use interface to only export limited functions
Currently, there is a struct that has many functions(let's call them: A,B,C,D,D1,D2). You would like to 
expose this struct for users to call the functions A,B and C and should avoid export D,D1,D2 and the 
data members.

Wrapping A,B and C and return a new type instance seems to be an alternative, but what if there are too 
many functions? 

In such case, returning an interface to **limit the exported functions** is a best solution.

That's what the interface `IRouter` in gin does.

## Encapsulate the general functions for users to choose
If the project wants to provide the feature which validates data format and analyzes it if it's proper
for users, what should it do?

Define a set of `itoa` and export them for users to pass is one solution, but now elegant. It has many shortages:
- Need to switch the value passed by user and set the corresponding handling logic.
- Hard to extend as the handling details fixed inside the library.

In gin validation, such as function `ShouldBindWith` and its interface:
```go
// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
Name() string
Bind(*http.Request, interface{}) error
}

func (c *Context) ShouldBindWith(obj interface{}, b binding.Binding) error
```
The `Binding` interface allows users define their own validation logic once the argument they passed implements the 
interface.



## Allow users to extend logics by themselves
There is an [issue](https://github.com/gin-gonic/gin/issues/1123) in gin which wants to extend the `gin.Context` 
and ask an interface.  

// TODO: consider more about this topic.