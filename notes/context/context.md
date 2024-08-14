The `context` package in Go is a powerful tool for controlling the lifecycle of operations, managing timeouts, deadlines, and cancellation signals across goroutines. It is particularly useful in concurrent programming, where you need to coordinate multiple goroutines or handle requests with time limits.

### Key Concepts

1. **`Context` Interface**: The `Context` interface carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.
2. **Creating a Context**:
    - `context.Background()`: Returns an empty `Context`, often used at the start of a request or goroutine.
    - `context.TODO()`: Returns an empty `Context` that can be used when you're unsure about what `Context` to use.
3. **Derived Contexts**:
    - `context.WithCancel(parent Context)`: Returns a copy of the parent context that can be canceled.
    - `context.WithDeadline(parent Context, deadline time.Time)`: Returns a copy with a specified deadline.
    - `context.WithTimeout(parent Context, timeout time.Duration)`: Returns a copy with a specified timeout.
    - `context.WithValue(parent Context, key, val interface{})`: Returns a copy that carries a specific key-value pair.
4. **Canceling a Context**: When a context is canceled, all contexts derived from it are also canceled.
