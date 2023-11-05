# Sink
a better way to share files

## TODO
- [ ] decide on a cli language (ts for web, swift for mobile)
- [ ] impl v1 with server
- [ ] figure out how to do it p2p
- [ ] better scanning?
- [ ] encryption

## priority
cli
web app
mobile app

## server thoughts
super simple
timeout & delete
data structure to hold keys
up -> key
down -> file

( maybe we can pass filename too )
client1 up -> server
server -key-> client1
client2 down -key-> server
server -IP-> client2
client2 -GET-> client1

client1: do you want to remove this file?
