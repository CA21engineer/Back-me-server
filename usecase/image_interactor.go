package usecase

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/joho/godotenv"
	"ca-zoooom/entity"
	"time"
	"os"
)

type ImageInteractor struct {
	ImageRepository ImageRepository
	StatusCode      int
}

func (interactor *ImageInteractor) ListImages(limit int, offset int) (i entity.Images, totalPages int, err error) {
	i, err = interactor.ImageRepository.Get(limit, offset)
	totalPages, err = interactor.ImageRepository.Count()
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	interactor.StatusCode = 200
	return
}

func (interactor *ImageInteractor) GetByID(id int) (i entity.Image, err error) {
	i, err = interactor.ImageRepository.GetById(id)
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	interactor.StatusCode = 200
	return
}

func (interactor *ImageInteractor) Add(image *entity.Image) (i entity.Image, err error) {
	err = interactor.ImageRepository.Insert(image)
	if err != nil {
		interactor.StatusCode = 500
		return
	}
	i, err = interactor.ImageRepository.GetById(image.Id)
	if err != nil {
		interactor.StatusCode = 500
		return
	}
	interactor.StatusCode = 201
	return
}

func (interactor *ImageInteractor) GetSignedUrl(rowImgName string) (url string, fileKey string, err error) {

	//環境変数読み込み
	err = godotenv.Load()

	//ファイル名作成
	fileKey = time.Now().Format("2006-01-02T15:04:05Z07:00") + rowImgName

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     os.Getenv("ACCESS_KEY_ID"),
				SecretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
			}),
		},
	)

	svc := s3.New(sess)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("ca-home-hackathon"),
		Key:    aws.String(fileKey),
	})
	url, err = req.Presign(15 * time.Minute)

	if err != nil {
		interactor.StatusCode = 500
		return
	}
	return
}
