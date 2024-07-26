
To make a channel an unbuffered channel, we omit the size/capacity.
e.g. 
	unbufferedChannel := make(chan string)
	bufferedChannel := make(chan string, 3)

#### unbuffered channels
- used for synchronous communication
- communication is blocked till another go routines reads the message
- the receiving go routine is blocked till the sending go routing puts a value in the channel
![[go_unbuffered_channels.png]]


#### buffered channels
**Buffered channels** are used for asynchronous communication. As the size of the channel is fixed, it does not wait for the receiver to receive the message.
![[buffered_channels.png]]
