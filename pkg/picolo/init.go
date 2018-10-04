package picolo

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"os"
)

const SERVICE_CREDS_FILE_ENV = "SERVICE_CREDS_FILE"

var fbApp *firebase.App
var ctx = context.Background()

func InitAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile(os.Getenv(SERVICE_CREDS_FILE_ENV))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Println("Error initializing app: ", err)
		return nil
	}
	fbApp = app
	return app
}
