Generics in Go allow you to write flexible and reusable code by defining functions, data structures, or types that can work with any data type. Introduced in Go 1.18, generics enable type parameters, which are a way to parameterize types within your code.

### Basics of Generics in Go

#### 1. **Type Parameters**
Type parameters allow you to write functions and data structures that work with any type. They are defined using square brackets `[]` and can be applied to functions, structs, and methods.
#### 2. **Type Constraints**
Type constraints specify what types can be used with your generic type. The most basic constraint is `any`, which allows any type. However, you can also use interfaces to define more specific constraints.

### Example 1: Generic Function

Let's start with a simple example of a generic function that can find the minimum value in a slice of any comparable type:

```go
package main

import "fmt"

// Generic function to find the minimum value
func Min[T comparable](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Min(3, 5))          // Works with integers
    fmt.Println(Min(2.3, 1.7))      // Works with floats
    fmt.Println(Min("apple", "bat")) // Works with strings
}
```


### Explanation:

- **Type Parameter `[T comparable]`**: The `[T comparable]` part defines a type parameter `T` that is constrained by the `comparable` constraint, meaning that the type `T` can be compared using the `<` operator.
- **Function Implementation**: The function `Min` works for any type that can be compared, making it flexible for use with various data types.

### Example 2: Generic Data Structure

You can also use generics to define data structures like a stack:

```go
package main

import "fmt"

// Generic Stack
type Stack[T any] struct {
    elements []T
}

// Push adds an element to the stack
func (s *Stack[T]) Push(element T) {
    s.elements = append(s.elements, element)
}

// Pop removes and returns the top element from the stack
func (s *Stack[T]) Pop() T {
    if len(s.elements) == 0 {
        var zero T
        return zero
    }
    element := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return element
}

func main() {
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    fmt.Println(intStack.Pop()) // 2

    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")
    fmt.Println(stringStack.Pop()) // world
}
```

### Explanation:

- **Generic Struct `Stack[T any]`**: The `Stack` struct is defined with a type parameter `T`, which allows it to hold elements of any type.
- **Method Definitions**: The `Push` and `Pop` methods work with the generic type `T`, making the stack implementation flexible for any data type.

### Example 3: Type Constraints

Suppose you want to restrict the generic function to only numeric types (e.g., integers, floats). You can use a custom constraint:

```go
package main

import "fmt"

// Constraint that allows only numeric types
type Numeric interface {
    int | int64 | float64 | float32
}

// Generic function with numeric constraint
func Add[T Numeric](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Add(10, 20))         // Works with int
    fmt.Println(Add(10.5, 20.3))     // Works with float64
    // fmt.Println(Add("a", "b"))    // This will cause a compile-time error
}
```

### Explanation:

- **Custom Constraint `Numeric`**: The `Numeric` interface restricts the types that can be passed to the `Add` function to only numeric types.
- **Compile-Time Safety**: Attempting to use a non-numeric type, like a string, with `Add` will result in a compile-time error, ensuring type safety.

### Summary of Generics in Go

- **Type Parameters**: Use type parameters to write functions and structs that work with any data type.
- **Type Constraints**: Use constraints to limit the types that can be used with generics. This ensures that the operations you perform on the generic type are valid.
- **Flexibility and Reusability**: Generics make your code more flexible and reusable, reducing the need to write multiple versions of the same function or data structure for different types.

Generics are a powerful feature in Go that allow for more abstract, reusable code without sacrificing type safety.

<hr>


Let's build a generic notification system in Go that can handle different types of notifications, such as email, SMS, and push notifications. We’ll use Go's generics to create a flexible system where different notification methods can be easily added and managed.

### Step 1: Define a Generic Notification Interface

We'll start by defining a `Notifier` interface that all notification types must implement:

```go
package main

import "fmt"

type Notifier interface {
    SendNotification() string
}
```

### Step 2: Implement Different Notification Types

Now let's create concrete implementations for email, SMS, and push notifications:

#### Email Notification

```go
type EmailNotification struct {
    EmailAddress string
    Message      string
}

func (e EmailNotification) SendNotification() string {
    return fmt.Sprintf("Sending Email to %s: %s", e.EmailAddress, e.Message)
}
```

#### SMS Notification

```go
type SMSNotification struct {
    PhoneNumber string
    Message     string
}

func (s SMSNotification) SendNotification() string {
    return fmt.Sprintf("Sending SMS to %s: %s", s.PhoneNumber, s.Message)
}
```

#### Push Notification

```go
type PushNotification struct {
    DeviceID string
    Message  string
}

func (p PushNotification) SendNotification() string {
    return fmt.Sprintf("Sending Push Notification to %s: %s", p.DeviceID, p.Message)
}
```

### Step 3: Create a Generic Notification Manager

Now we’ll create a generic `NotificationManager` that can handle sending notifications regardless of their type:

```go
type NotificationManager[T Notifier] struct {
    Notifications []T
}

func (nm *NotificationManager[T]) AddNotification(notification T) {
    nm.Notifications = append(nm.Notifications, notification)
}

func (nm NotificationManager[T]) SendAll() {
    for _, notification := range nm.Notifications {
        fmt.Println(notification.SendNotification())
    }
}
```

### Explanation:

- **Generic `NotificationManager`**: The `NotificationManager` is generic, meaning it can manage notifications of any type that implements the `Notifier` interface.
- **`AddNotification` Method**: This method adds a new notification to the manager's list.
- **`SendAll` Method**: This method iterates over all notifications and sends them using their respective `SendNotification` method.

### Step 4: Use the Notification System

Now let’s put everything together:

```go
func main() {
    email := EmailNotification{
        EmailAddress: "user@example.com",
        Message:      "You've got a new email!",
    }

    sms := SMSNotification{
        PhoneNumber: "+1234567890",
        Message:     "Your OTP is 123456",
    }

    push := PushNotification{
        DeviceID: "device_123",
        Message:  "You have a new message!",
    }

    // Create a notification manager for email notifications
    emailManager := NotificationManager[EmailNotification]{}
    emailManager.AddNotification(email)

    // Create a notification manager for SMS notifications
    smsManager := NotificationManager[SMSNotification]{}
    smsManager.AddNotification(sms)

    // Create a notification manager for push notifications
    pushManager := NotificationManager[PushNotification]{}
    pushManager.AddNotification(push)

    // Send all notifications
    emailManager.SendAll()
    smsManager.SendAll()
    pushManager.SendAll()
}
```

### Output:

```sh
Sending Email to user@example.com: You've got a new email!
Sending SMS to +1234567890: Your OTP is 123456
Sending Push Notification to device_123: You have a new message!
```

### Summary

- **Generic Notification Manager**: The `NotificationManager` is a generic structure that can manage and send notifications of any type that implements the `Notifier` interface.
- **Flexibility**: You can easily add new notification types (e.g., in-app notifications) by implementing the `Notifier` interface, without changing the `NotificationManager`.
- **Type Safety**: Generics ensure that only valid notification types are added to the manager, providing type safety while maintaining flexibility.

This example illustrates how generics in Go can be used to build a flexible, type-safe notification system that can easily accommodate different types of notifications and grow as new requirements emerge.
