# Sink
a better way to share files

NEW WORKFLOW:
client1 sends file in POST or wtvr to client2

## TODO
- [ ] impl v1 with server
- [ ] jpg to pdf
- [ ] redis?
- [ ] figure out how to do it p2p
- [ ] better scanning?
- [ ] encryption
- [ ] cmd line flags

## priority
cli
web app

## server thoughts
use redis?
timeout & delete
data structure to hold keys / filenames

client1 up -> server
server -key-> client1
client2 down -key-> server
server -IP-> client1
client1 -POST-> client2

complication: uploading file to mobile site
how will we 'sink down' from mobile

use SSE. server will eventually send ip to send file to.
