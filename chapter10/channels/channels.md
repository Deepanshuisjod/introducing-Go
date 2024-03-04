Channels provide a way for two goroutines to communicate with each other and synchronize their execution. 
Using a channel like this synchronizes the two goroutines. When pinger attempts to
send a message on the channel, it will wait until printer is ready to receive the mes‐
sage (this is known as blocking).
