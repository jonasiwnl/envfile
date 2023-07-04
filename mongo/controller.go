package mongo

import (
	"bufio"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func Execute(ctx context.Context, client *mongo.Client) error {
	file, err := os.Open("temp")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	for reader.Scan() {
		line := reader.Text()

		err = client.Database("TODO").RunCommand(ctx, line).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
