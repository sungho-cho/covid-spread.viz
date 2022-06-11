package utils

import (
	"cloud.google.com/go/storage"
	"context"
	"io/ioutil"
	"log"
)

const bucketName = "covid-spread-viz.appspot.com"

var ctx context.Context
var bkt *storage.BucketHandle

func init() {
	ctx = context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Google Cloud Storage: %s", err)
	}
	bkt = client.Bucket(bucketName)
}

func DoesExist(objName string) bool {
	obj := bkt.Object(objName)
	_, err := obj.NewReader(ctx)
	if err != nil && err != storage.ErrObjectNotExist {
		log.Fatalf("Failed to check if object(%s) exists: %s", objName, err)
		return false
	}
	return err == nil || err != storage.ErrObjectNotExist
}

func ReadObject(objName string) ([]byte, error) {
	obj := bkt.Object(objName)
	r, err := obj.NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to read object(%s): %s", objName, err)
		return []byte{}, err
	}
	defer r.Close()
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalf("Failed to copy object(%s): %s", objName, err)
		return []byte{}, err
	}
	return body, nil
}

func WriteObject(objName string, data []byte) error {
	obj := bkt.Object(objName)
	w := obj.NewWriter(ctx)
	if _, err := w.Write(data); err != nil {
		log.Fatalf("Failed to write to object(%s): %s", objName, err)
		return err
	}
	if err := w.Close(); err != nil {
		log.Fatalf("Failed to close writer to object(%s): %s", objName, err)
		return err
	}
	return nil
}
