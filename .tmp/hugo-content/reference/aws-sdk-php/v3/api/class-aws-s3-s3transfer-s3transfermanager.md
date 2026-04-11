Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)

## S3TransferManager        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-s3transfermanager-toc.md)

#### Methods  [header link](class-aws-s3-s3transfer-s3transfermanager-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-s3transfermanager-method-construct.md)
: mixed [download()](class-aws-s3-s3transfer-s3transfermanager-method-download.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[downloadDirectory()](class-aws-s3-s3transfer-s3transfermanager-method-downloaddirectory.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[downloadFile()](class-aws-s3-s3transfer-s3transfermanager-method-downloadfile.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[getConfig()](class-aws-s3-s3transfer-s3transfermanager-method-getconfig.md)
: [S3TransferManagerConfig](class-aws-s3-s3transfer-models-s3transfermanagerconfig.md)[getS3Client()](class-aws-s3-s3transfer-s3transfermanager-method-gets3client.md)
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)[isValidS3URI()](class-aws-s3-s3transfer-s3transfermanager-method-isvalids3uri.md)
: bool Validates a string value is a valid S3 URI.[resumeDownload()](class-aws-s3-s3transfer-s3transfermanager-method-resumedownload.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[resumeUpload()](class-aws-s3-s3transfer-s3transfermanager-method-resumeupload.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[s3UriAsBucketAndKey()](class-aws-s3-s3transfer-s3transfermanager-method-s3uriasbucketandkey.md)
: array<string\|int, mixed> Converts a S3 URI into an array with a Bucket and Key
properties set.[upload()](class-aws-s3-s3transfer-s3transfermanager-method-upload.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[uploadDirectory()](class-aws-s3-s3transfer-s3transfermanager-method-uploaddirectory.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

### Methods  [header link](class-aws-s3-s3transfer-s3transfermanager-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-construct.md)

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

#### download()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-download.md)

`
    public
                    download(DownloadRequest $downloadRequest) : PromiseInterface`

##### Parameters

$downloadRequest
: [DownloadRequest](class-aws-s3-s3transfer-models-downloadrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### downloadDirectory()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-downloaddirectory.md)

`
    public
                    downloadDirectory(DownloadDirectoryRequest $downloadDirectoryRequest) : PromiseInterface`

##### Parameters

$downloadDirectoryRequest
: [DownloadDirectoryRequest](class-aws-s3-s3transfer-models-downloaddirectoryrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### downloadFile()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-downloadfile.md)

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

#### getConfig()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-getconfig.md)

`
    public
                    getConfig() : S3TransferManagerConfig`

##### Return values

[S3TransferManagerConfig](class-aws-s3-s3transfer-models-s3transfermanagerconfig.md)

#### getS3Client()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-gets3client.md)

`
    public
                    getS3Client() : S3ClientInterface`

##### Return values

[S3ClientInterface](class-aws-s3-s3clientinterface.md)

#### isValidS3URI()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-isvalids3uri.md)

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

#### resumeDownload()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-resumedownload.md)

`
    public
                    resumeDownload(ResumeDownloadRequest $resumeDownloadRequest) : PromiseInterface`

##### Parameters

$resumeDownloadRequest
: [ResumeDownloadRequest](class-aws-s3-s3transfer-models-resumedownloadrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### resumeUpload()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-resumeupload.md)

`
    public
                    resumeUpload(ResumeUploadRequest $resumeUploadRequest) : PromiseInterface`

##### Parameters

$resumeUploadRequest
: [ResumeUploadRequest](class-aws-s3-s3transfer-models-resumeuploadrequest.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### s3UriAsBucketAndKey()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-s3uriasbucketandkey.md)

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

#### upload()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-upload.md)

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

#### uploadDirectory()  [header link](class-aws-s3-s3transfer-s3transfermanager-method-uploaddirectory.md)

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
  - [Methods](class-aws-s3-s3transfer-s3transfermanager-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-s3transfermanager-method-construct.md)
  - [download()](class-aws-s3-s3transfer-s3transfermanager-method-download.md)
  - [downloadDirectory()](class-aws-s3-s3transfer-s3transfermanager-method-downloaddirectory.md)
  - [downloadFile()](class-aws-s3-s3transfer-s3transfermanager-method-downloadfile.md)
  - [getConfig()](class-aws-s3-s3transfer-s3transfermanager-method-getconfig.md)
  - [getS3Client()](class-aws-s3-s3transfer-s3transfermanager-method-gets3client.md)
  - [isValidS3URI()](class-aws-s3-s3transfer-s3transfermanager-method-isvalids3uri.md)
  - [resumeDownload()](class-aws-s3-s3transfer-s3transfermanager-method-resumedownload.md)
  - [resumeUpload()](class-aws-s3-s3transfer-s3transfermanager-method-resumeupload.md)
  - [s3UriAsBucketAndKey()](class-aws-s3-s3transfer-s3transfermanager-method-s3uriasbucketandkey.md)
  - [upload()](class-aws-s3-s3transfer-s3transfermanager-method-upload.md)
  - [uploadDirectory()](class-aws-s3-s3transfer-s3transfermanager-method-uploaddirectory.md)

[Back To Top](class-aws-s3-s3transfer-s3transfermanager-top.md)
