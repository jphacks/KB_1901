package util

import (
	//"fmt"
	"os"
	"os/exec"
	"io/ioutil"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)


func FileUpload( uploadname string ) error {
	sess, err := S3Connect()

	if err != nil {
		return err
	}

	file, err := os.Open( "send.json" )
	defer file.Close()
	
	if err != nil {
		return err
	}

	bucketname := "jphacksstorage"
	
	uploader := s3manager.NewUploader( sess )
	_, err = uploader.Upload( &s3manager.UploadInput {
		Bucket: aws.String( bucketname ),
		Key: aws.String( uploadname ),
		Body: file,
	})

	exec.Command( "rm", "-rf", "send.json" ).Run()

	if err != nil {
		return err
	}

	return nil
}

func FileDownload( downloadname string ) ( []byte, error ) {
	var jsonBytes []byte
	sess, err := S3Connect()

	if err != nil {
		return jsonBytes, err
	}
	
	filename := "send.json"
	bucketname := "jphacksstorage"
	
	file, err := os.Create( filename )
	defer file.Close()
	
	if err != nil {
		return jsonBytes, err
	}

	downloader := s3manager.NewDownloader( sess )

	_, err = downloader.Download( file,
		&s3.GetObjectInput{
			Bucket: aws.String( bucketname ),
			Key:    aws.String( downloadname ),
		})

	if err != nil {
		return jsonBytes, err
	}

	jsonBytes, err = ioutil.ReadAll( file )
	exec.Command( "rm", "-rf", filename ).Run()

	if err != nil {
		return jsonBytes, err
	}

	return jsonBytes, nil
}

func S3Connect() ( *session.Session, error ) {
	creds := credentials.NewStaticCredentials("AKIA3F5VIQI6XEOHZT2U", "BdGBlLOD6kmwGznB5SMFL3+v7PbTBwD7Qw9WjRDC", "" )
	
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region: aws.String("us-east-2")},
	)

	if err != nil {
		return nil, err
	}

	return sess, nil
} 
