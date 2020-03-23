# PC-CONTROL

Http service enable to turning on/off PC by GPIO.

**Platform**: Raspberry PI B+  
**Interface**: GPIO  
**GPIO port**: 17  
**HTTP port**: 80

## Install
```sh
go get -u github.com/andrzejd-pl/pc-control
go build github.com/andrzejd-pl/pc-control
$GOPATH/bin/pc-control
```

### Systemd

You can use `systemd` and create this program as service
```sh
cp $GOPATH/github.com/andrzejd-pl/pc-control/http-power-switch.service /etc/systemd/system/
```

## TODO

- custom configuration
- dockerize
- auto shutdown
