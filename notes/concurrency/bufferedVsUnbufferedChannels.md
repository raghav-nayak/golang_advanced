
To make a channel an unbuffered channel, we omit the size/capacity.
e.g. 
	unbufferedChannel := make(chan string)
	bufferedChannel := make(chan string, 3)

**Unbuffered channels** are used for synchronous communication
![[unbuffered_channels.png]]


**Buffered channels** are used for asynchronous communication. As the size of the channel is fixed, it does not wait for the receiver to receive the message.
![[buffered_channels.png]]
