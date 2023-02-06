package directory

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	_ "github.com/lib/pq"
)

var minioClientMap map[NamespaceID]*minio.Client

func init() {
	minioClientMap = map[NamespaceID]*minio.Client{}
}

type AlreadyPresentError struct {
	err error
}

func (e AlreadyPresentError) Error() string {
	return e.err.Error()
}

type FileManager interface {
	UploadFile(context.Context, string, []byte, interface{}) (UploadResult, error)
	DownloadFile(context.Context, string, string, interface{}) error
	ExposeFile(context.Context, string) (string, error)
	Client() interface{}
	Bucket() string
}

type MinioFileManager struct {
	client    *minio.Client
	namespace string
}

type UploadResult struct {
	Bucket       string
	Key          string
	ETag         string
	Size         int64
	LastModified time.Time
	Location     string
	VersionID    string
}

func checkIfFileExists(ctx context.Context, client *minio.Client, bucket, filename string) (string, bool) {
	info, err := client.StatObject(ctx, bucket, filename, minio.StatObjectOptions{})
	if err != nil {
		return "", false
	}
	return info.Key, true
}

func (mfm *MinioFileManager) UploadFile(ctx context.Context, filename string, data []byte, extra interface{}) (UploadResult, error) {
	err := mfm.createBucketIfNeeded(ctx)
	if err != nil {
		return UploadResult{}, err
	}

	if key, has := checkIfFileExists(ctx, mfm.client, mfm.namespace, path.Join(mfm.namespace, filename)); has {
		return UploadResult{}, AlreadyPresentError{err: errors.New(key)}
	}

	info, err := mfm.client.PutObject(ctx, mfm.namespace, path.Join(mfm.namespace, filename),
		bytes.NewReader(data), int64(len(data)),
		extra.(minio.PutObjectOptions))

	if err != nil {
		return UploadResult{}, err
	}

	return UploadResult{
		Location:     info.Location,
		Bucket:       info.Bucket,
		Key:          info.Key,
		ETag:         info.ETag,
		Size:         info.Size,
		LastModified: info.LastModified,
		VersionID:    info.VersionID,
	}, nil
}

func (mfm *MinioFileManager) DownloadFile(ctx context.Context, remoteFile string, localFile string, extra interface{}) error {
	return mfm.client.FGetObject(ctx, mfm.namespace, path.Join(mfm.namespace, remoteFile), localFile, extra.(minio.GetObjectOptions))
}

func (mfm *MinioFileManager) ExposeFile(ctx context.Context, filepath string) (string, error) {

	console_ip, err := GetManagementHost(NewGlobalContext())
	if err != nil {
		return "", err
	}

	headers := http.Header{"Host": []string{console_ip}}

	url, err := mfm.client.PresignHeader(context.Background(),
		"GET",
		mfm.namespace,
		filepath,
		time.Hour*10,
		url.Values{},
		headers)

	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (mfm *MinioFileManager) Client() interface{} {
	return mfm.client
}

func (mfm *MinioFileManager) Bucket() string {
	return mfm.namespace
}

func (mfm *MinioFileManager) createBucketIfNeeded(ctx context.Context) error {

	exists, err := mfm.client.BucketExists(ctx, mfm.namespace)

	if err != nil {
		return err
	}

	if !exists {
		err = mfm.client.MakeBucket(ctx, mfm.namespace,
			minio.MakeBucketOptions{ObjectLocking: false})

	}
	return err
}

func newMinioClient(endpoints DBConfigs) (*minio.Client, error) {
	if endpoints.Minio == nil {
		return nil, errors.New("no defined minio config")
	}
	minioClient, err := minio.New(endpoints.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(endpoints.Minio.Username, endpoints.Minio.Password, ""),
		Secure: endpoints.Minio.Secure,
	})
	return minioClient, err
}

func MinioClient(ctx context.Context) (FileManager, error) {
	client, err := getClient(NewGlobalContext(), minioClientMap, newMinioClient)
	if err != nil {
		return nil, err
	}

	ns, err := ExtractNamespace(ctx)
	if err != nil {
		return nil, err
	}

	return &MinioFileManager{
		client:    client,
		namespace: string(ns),
	}, err
}
