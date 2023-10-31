package cmd

import (
	"github.com/andrelmm/imersao-fs-fc/codepix-go/application/grpc"
	"github.com/andrelmm/imersao-fs-fc/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
