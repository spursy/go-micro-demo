**1. After clone code, install inspired package**
```
govendor sync
```

**2. Install Protobuf (Protocol Buffers - Google's data interchange format)**

for ubuntu
```
apt-get install protobuf
```
for mac
```
brew install protobuf
```

**3. Install consul**
for ubunut
```
apt-get install consul
```

for mac
```
brew install consul
```

**4. Init consul**
All services will register to consul
```
consul agent -dev
```

check service status
```
http://localhost:8500/
```

**5. Generate micro and go file for proto**
```
cd proto
protoc --micro_out=. --go_out=. greeter.proto
```

**6. Process hello service**
```
go run services/hello.go
```

** 7. Check services
```
micro list services
```

**7. Process client service**
```
go run clients/helloclient.go
```
