Asynchronous call graph
	- In done channel
		from main, done channel is passed to all the go routines called from main


Context
- like done channel, it is passed to the go routines called from main
- immutable
- cannot be cancelled by the children
- helps to remove leaks
- provides API to cancel the call graphs i.e. cancel()
- you can pass data throughout the call graph. We can pass anything want(kind of abuse as go is a strict typed language)
- 

ctx, cancel = newContext()
defer cancel

only ctx is passed to the children from the main.

children can pass the context from main to their children or you can create a separate context to their children.

newCtx, cancel = withCancel(ctx)

![[go_context.png]]

withCancel() 
cancel() -> closes the Done channel
