module github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service

go 1.18

replace github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common => ../common

require (
	github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service v0.0.0-20220508215230-00f28c4f6d5f
	github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common v0.0.0-20220505203713-b8b7bb6f5439
	github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service v0.0.0-20220829065427-062f2e45ad55
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/snowzach/rotatefilehook v0.0.0-20220211133110-53752135082d
	github.com/t-tomalak/logrus-easy-formatter v0.0.0-20190827215021-c074f06c5816
	go.mongodb.org/mongo-driver v1.9.1
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3
	google.golang.org/grpc v1.47.0
)

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
