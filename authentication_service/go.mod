module github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service

go 1.18

replace github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common => ../common

require (
	github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common v0.0.0-20220505203713-b8b7bb6f5439
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/sirupsen/logrus v1.4.2
	github.com/snowzach/rotatefilehook v0.0.0-20220211133110-53752135082d
	github.com/t-tomalak/logrus-easy-formatter v0.0.0-20190827215021-c074f06c5816
	go.mongodb.org/mongo-driver v1.9.1
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
	google.golang.org/grpc v1.46.0
	gopkg.in/mail.v2 v2.3.1
	gorm.io/driver/postgres v1.3.5
	gorm.io/gorm v1.23.5
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/klauspost/compress v1.15.2 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.2.0 // indirect
	github.com/nats-io/nats.go v1.14.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	google.golang.org/genproto v0.0.0-20220429170224-98d788798c3e // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
