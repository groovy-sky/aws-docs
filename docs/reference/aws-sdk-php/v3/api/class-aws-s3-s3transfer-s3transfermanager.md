Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)

## S3TransferManager        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method___construct)
: mixed [download()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_download)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[downloadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_downloadDirectory)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[downloadFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_downloadFile)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_getConfig)
: [S3TransferManagerConfig](class-aws-s3-s3transfer-models-s3transfermanagerconfig.md)[getS3Client()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_getS3Client)
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)[isValidS3URI()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_isValidS3URI)
: bool Validates a string value is a valid S3 URI.[resumeDownload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_resumeDownload)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[resumeUpload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_resumeUpload)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[s3UriAsBucketAndKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_s3UriAsBucketAndKey)
: array<string\|int, mixed> Converts a S3 URI into an array with a Bucket and Key
properties set.[upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_upload)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[uploadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_uploadDirectory)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method___construct)

`
    public
                    __construct([S3ClientInterface|null $s3Client = null ][, array<string|int, mixed>|S3TransferManagerConfig|null $config = null ]) : mixed`

##### Parameters

$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md) \|null
= null

If provided as null then,
a default client will be created where its region will be the one
resolved from either the default from the config or the provided.

$config
: array<string\|int, mixed>\| [S3TransferManagerConfig](class-aws-s3-s3transfer-models-s3transfermanagerconfig.md) \|null
= null

#### download()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_download)

`
    public
                    download(DownloadRequest $downloadRequest) : PromiseInterface`

##### Parameters

$downloadRequest
: [DownloadRequest](class-aws-s3-s3transfer-models-downloadrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### downloadDirectory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_downloadDirectory)

`
    public
                    downloadDirectory(DownloadDirectoryRequest $downloadDirectoryRequest) : PromiseInterface`

##### Parameters

$downloadDirectoryRequest
: [DownloadDirectoryRequest](class-aws-s3-s3transfer-models-downloaddirectoryrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### downloadFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_downloadFile)

`
    public
                    downloadFile(DownloadFileRequest $downloadFileRequest[, S3ClientInterface|null $s3Client = null ]) : PromiseInterface`

##### Parameters

$downloadFileRequest
: [DownloadFileRequest](class-aws-s3-s3transfer-models-downloadfilerequest.md)$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md) \|null
= null

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_getConfig)

`
    public
                    getConfig() : S3TransferManagerConfig`

##### Return values

[S3TransferManagerConfig](class-aws-s3-s3transfer-models-s3transfermanagerconfig.md)

#### getS3Client()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_getS3Client)

`
    public
                    getS3Client() : S3ClientInterface`

##### Return values

[S3ClientInterface](class-aws-s3-s3clientinterface.md)

#### isValidS3URI()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_isValidS3URI)

Validates a string value is a valid S3 URI.

`
    public
            static        isValidS3URI(string $uri) : bool`

Valid S3 URI Example: S3://mybucket.dev/myobject.txt

##### Parameters

$uri
: string

##### Return values

bool

#### resumeDownload()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_resumeDownload)

`
    public
                    resumeDownload(ResumeDownloadRequest $resumeDownloadRequest) : PromiseInterface`

##### Parameters

$resumeDownloadRequest
: [ResumeDownloadRequest](class-aws-s3-s3transfer-models-resumedownloadrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### resumeUpload()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_resumeUpload)

`
    public
                    resumeUpload(ResumeUploadRequest $resumeUploadRequest) : PromiseInterface`

##### Parameters

$resumeUploadRequest
: [ResumeUploadRequest](class-aws-s3-s3transfer-models-resumeuploadrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### s3UriAsBucketAndKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_s3UriAsBucketAndKey)

Converts a S3 URI into an array with a Bucket and Key
properties set.

`
    public
            static        s3UriAsBucketAndKey(string $uri) : array<string|int, mixed>`

##### Parameters

$uri
: string

##### Return values

array<string\|int, mixed>

#### upload()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_upload)

`
    public
                    upload(UploadRequest $uploadRequest[, S3ClientInterface|null $s3Client = null ]) : PromiseInterface`

##### Parameters

$uploadRequest
: [UploadRequest](class-aws-s3-s3transfer-models-uploadrequest.md)$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md) \|null
= null

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### uploadDirectory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html\#method_uploadDirectory)

`
    public
                    uploadDirectory(UploadDirectoryRequest $uploadDirectoryRequest) : PromiseInterface`

##### Parameters

$uploadDirectoryRequest
: [UploadDirectoryRequest](class-aws-s3-s3transfer-models-uploaddirectoryrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method___construct)
  - [download()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_download)
  - [downloadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_downloadDirectory)
  - [downloadFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_downloadFile)
  - [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_getConfig)
  - [getS3Client()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_getS3Client)
  - [isValidS3URI()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_isValidS3URI)
  - [resumeDownload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_resumeDownload)
  - [resumeUpload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_resumeUpload)
  - [s3UriAsBucketAndKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_s3UriAsBucketAndKey)
  - [upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_upload)
  - [uploadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#method_uploadDirectory)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.S3TransferManager.html#top)
