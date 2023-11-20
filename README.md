# Sink
a better way to share files

NEW WORKFLOW:
client1 sends file in POST or wtvr to client2

sink expect -> sink send?

sink expect -> server
server has { key: ip }
sink send -key-> server
send <-ip- server
send -file-> expect

## TODO
- [ ] impl v1 with server
- [ ] jpg to pdf
- [ ] redis?
- [ ] figure out how to do it p2p
- [ ] better scanning?
- [ ] cmd line flags

## priority
cli
web app

## server thoughts
use redis?
timeout & delete
data structure to hold keys / filenames

