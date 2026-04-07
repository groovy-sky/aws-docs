Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)

## MultipartUploader     extends [AbstractMultipartUploader](class-aws-s3-s3transfer-abstractmultipartuploader.md)   in package    - [Aws](package-aws.md)

FinalYes

Multipart uploader implementation.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html\#toc-constants)

[DEFAULT\_CHECKSUM\_CALCULATION\_ALGORITHM](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#constant_DEFAULT_CHECKSUM_CALCULATION_ALGORITHM)
= 'crc32' [PART\_MAX\_NUM](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MAX_NUM)
= 10000 [PART\_MAX\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MAX_SIZE)
= 5 \* 1024 \* 1024 \* 1024 [PART\_MIN\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MIN_SIZE)
= 5 \* 1024 \* 1024

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#method___construct)
: mixed [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getCurrentSnapshot)
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null Get the current progress snapshot.[getPartsCompleted()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getPartsCompleted)
: array<string\|int, mixed> [getUploadId()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getUploadId)
: string\|null [promise()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.[upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#method_upload)
: [UploadResult](class-aws-s3-s3transfer-models-uploadresult.md)Sync upload method.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html\#constants)

#### DEFAULT\_CHECKSUM\_CALCULATION\_ALGORITHM  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html\#constant_DEFAULT_CHECKSUM_CALCULATION_ALGORITHM)

`
    public
        mixed
    DEFAULT_CHECKSUM_CALCULATION_ALGORITHM
    = 'crc32'
`

#### PART\_MAX\_NUM  [header link](class-aws-s3-s3transfer-abstractmultipartuploader.md\#constant_PART_MAX_NUM)

`
    public
        mixed
    PART_MAX_NUM
    = 10000
`

#### PART\_MAX\_SIZE  [header link](class-aws-s3-s3transfer-abstractmultipartuploader.md\#constant_PART_MAX_SIZE)

`
    public
        mixed
    PART_MAX_SIZE
    = 5 * 1024 * 1024 * 1024
`

#### PART\_MIN\_SIZE  [header link](class-aws-s3-s3transfer-abstractmultipartuploader.md\#constant_PART_MIN_SIZE)

`
    public
        mixed
    PART_MIN_SIZE
    = 5 * 1024 * 1024
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html\#method___construct)

`
    public
                    __construct(S3ClientInterface $s3Client, array<string|int, mixed> $requestArgs, string|StreamInterface $source[, array<string|int, mixed> $config = [] ][, TransferListenerNotifier|null $listenerNotifier = null ][, ResumableUpload|null $resumableUpload = null ]) : mixed`

##### Parameters

$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)$requestArgs
: array<string\|int, mixed>$source
: string\| [StreamInterface](class-psr-http-message-streaminterface.md)$config
: array<string\|int, mixed>
= \[\]

- target\_part\_size\_bytes: (int, optional)
- request\_checksum\_calculation: (string, optional)
- concurrency: (int, optional)

$listenerNotifier
: [TransferListenerNotifier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html) \|null
= null$resumableUpload
: [ResumableUpload](class-aws-s3-s3transfer-models-resumableupload.md) \|null
= null

#### getCurrentSnapshot()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader.md\#method_getCurrentSnapshot)

Get the current progress snapshot.

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot|null`

##### Return values

[TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null

#### getPartsCompleted()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader.md\#method_getPartsCompleted)

`
    public
                    getPartsCompleted() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getUploadId()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader.md\#method_getUploadId)

`
    public
                    getUploadId() : string|null`

##### Return values

string\|null

#### promise()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader.md\#method_promise)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### upload()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html\#method_upload)

Sync upload method.

`
    public
                    upload() : UploadResult`

##### Return values

[UploadResult](class-aws-s3-s3transfer-models-uploadresult.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#toc-methods)
- Constants
  - [DEFAULT\_CHECKSUM\_CALCULATION\_ALGORITHM](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#constant_DEFAULT_CHECKSUM_CALCULATION_ALGORITHM)
  - [PART\_MAX\_NUM](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MAX_NUM)
  - [PART\_MAX\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MAX_SIZE)
  - [PART\_MIN\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MIN_SIZE)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#method___construct)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getCurrentSnapshot)
  - [getPartsCompleted()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getPartsCompleted)
  - [getUploadId()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getUploadId)
  - [promise()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_promise)
  - [upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#method_upload)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.MultipartUploader.html#top)
