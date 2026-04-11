Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)

## MultipartUploader     extends [AbstractMultipartUploader](class-aws-s3-s3transfer-abstractmultipartuploader.md)   in package    - [Aws](package-aws.md)

FinalYes

Multipart uploader implementation.

### Table of Contents  [header link](class-aws-s3-s3transfer-multipartuploader-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-multipartuploader-toc-constants.md)

[DEFAULT\_CHECKSUM\_CALCULATION\_ALGORITHM](class-aws-s3-s3transfer-multipartuploader-constant-default-checksum-calculation-algorithm.md)
= 'crc32' [PART\_MAX\_NUM](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MAX_NUM)
= 10000 [PART\_MAX\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MAX_SIZE)
= 5 \* 1024 \* 1024 \* 1024 [PART\_MIN\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MIN_SIZE)
= 5 \* 1024 \* 1024

#### Methods  [header link](class-aws-s3-s3transfer-multipartuploader-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-multipartuploader-method-construct.md)
: mixed [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getCurrentSnapshot)
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null Get the current progress snapshot.[getPartsCompleted()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getPartsCompleted)
: array<string\|int, mixed> [getUploadId()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getUploadId)
: string\|null [promise()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.[upload()](class-aws-s3-s3transfer-multipartuploader-method-upload.md)
: [UploadResult](class-aws-s3-s3transfer-models-uploadresult.md)Sync upload method.

### Constants  [header link](class-aws-s3-s3transfer-multipartuploader-constants.md)

#### DEFAULT\_CHECKSUM\_CALCULATION\_ALGORITHM  [header link](class-aws-s3-s3transfer-multipartuploader-constant-default-checksum-calculation-algorithm.md)

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

### Methods  [header link](class-aws-s3-s3transfer-multipartuploader-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-multipartuploader-method-construct.md)

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
: [TransferListenerNotifier](class-aws-s3-s3transfer-progress-transferlistenernotifier.md) \|null
= null$resumableUpload
: [ResumableUpload](class-aws-s3-s3transfer-models-resumableupload.md) \|null
= null

#### getCurrentSnapshot()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader.md\#method_getCurrentSnapshot)

Get the current progress snapshot.

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot|null`

##### Return values

[TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null

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

#### upload()  [header link](class-aws-s3-s3transfer-multipartuploader-method-upload.md)

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
  - [Constants](class-aws-s3-s3transfer-multipartuploader-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-multipartuploader-toc-methods.md)
- Constants
  - [DEFAULT\_CHECKSUM\_CALCULATION\_ALGORITHM](class-aws-s3-s3transfer-multipartuploader-constant-default-checksum-calculation-algorithm.md)
  - [PART\_MAX\_NUM](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MAX_NUM)
  - [PART\_MAX\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MAX_SIZE)
  - [PART\_MIN\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader.md#constant_PART_MIN_SIZE)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-multipartuploader-method-construct.md)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getCurrentSnapshot)
  - [getPartsCompleted()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getPartsCompleted)
  - [getUploadId()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_getUploadId)
  - [promise()](class-aws-s3-s3transfer-abstractmultipartuploader.md#method_promise)
  - [upload()](class-aws-s3-s3transfer-multipartuploader-method-upload.md)

[Back To Top](class-aws-s3-s3transfer-multipartuploader-top.md)
