package fs

import (
	"io/fs"
	"log"
	"os"
	"path"
)

// TODO: understand and implement bucket ownership (??)

// This is a bucket
type Bucket struct {
	name string
	path string
	acl  string
}

func (c *S3Controller) getBucketPath(bucket string) string {
	return path.Join(c.root, string(bucket))
}

func (c *S3Controller) CreateBucket(name string, acl string) *FSError {
	if _, exists := c.buckets[name]; !exists {
		return &FSError{
			Code:      ERR_BUCKET_ALREADY_EXISTS_CODE,
			Message:   ERR_BUCKET_ALREADY_EXISTS_MSG,
			Errorcode: 400,
		}
	}

	path := path.Join(c.root, name)

	c.buckets[name] = Bucket{
		name: name,
		acl:  acl,
		path: path,
	}

	if err := os.Mkdir(path, fs.FileMode(FS_FILEMODE)); err != nil {
		log.Fatal(err)
	}

	return nil
}
