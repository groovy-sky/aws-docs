Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## MultipartUploadingTrait

### Table of Contents  [header link](class-aws-s3-multipartuploadingtrait-toc.md)

#### Methods  [header link](class-aws-s3-multipartuploadingtrait-toc-methods.md)

[getStateFromService()](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)
: [UploadState](class-aws-multipart-uploadstate.md)Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

### Methods  [header link](class-aws-s3-multipartuploadingtrait-methods.md)

#### getStateFromService()  [header link](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)

Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

`
    public
            static        getStateFromService(S3ClientInterface $client, string $bucket, string $key, string $uploadId) : UploadState`

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

S3Client used for the upload.

$bucket
: string

Bucket for the multipart upload.

$key
: string

Object key for the multipart upload.

$uploadId
: string

Upload ID for the multipart upload.

##### Return values

[UploadState](class-aws-multipart-uploadstate.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-multipartuploadingtrait-toc-methods.md)
- Methods
  - [getStateFromService()](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)

[Back To Top](class-aws-s3-multipartuploadingtrait-top.md)
