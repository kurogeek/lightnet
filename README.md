# lightnet
### To test:
```
go test ./...
```
To run proxy-server:
```
cd cmd/proxy-server
AUTH_KEY=key go run main.go
```

To run calculator-server:
```
cd cmd/cal-server
AUTH_KEY=key go run main.go
```

Note: AUTH_KEY between cal-server and proxy-server need to be the same. This's because to prevent the traffic to go directly to cal-server.
