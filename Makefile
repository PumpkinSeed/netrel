dev:
	go build -o netrel main.go
	sudo setcap cap_net_raw=+ep ./netrel
