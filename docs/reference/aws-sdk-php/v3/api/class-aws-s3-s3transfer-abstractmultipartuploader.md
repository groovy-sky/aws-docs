Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)

## AbstractMultipartUploader        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

AbstractYes

Abstract base class for multipart operations

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#toc-interfaces)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#toc-constants)

[PART\_MAX\_NUM](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#constant_PART_MAX_NUM)
= 10000 [PART\_MAX\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#constant_PART_MAX_SIZE)
= 5 \* 1024 \* 1024 \* 1024 [PART\_MIN\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#constant_PART_MIN_SIZE)
= 5 \* 1024 \* 1024

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method___construct)
: mixed [getCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method_getCurrentSnapshot)
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null Get the current progress snapshot.[getPartsCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method_getPartsCompleted)
: array<string\|int, mixed> [getUploadId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method_getUploadId)
: string\|null [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#constants)

#### PART\_MAX\_NUM  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#constant_PART_MAX_NUM)

`
    public
        mixed
    PART_MAX_NUM
    = 10000
`

#### PART\_MAX\_SIZE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#constant_PART_MAX_SIZE)

`
    public
        mixed
    PART_MAX_SIZE
    = 5 * 1024 * 1024 * 1024
`

#### PART\_MIN\_SIZE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#constant_PART_MIN_SIZE)

`
    public
        mixed
    PART_MIN_SIZE
    = 5 * 1024 * 1024
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#method___construct)

`
    public
                    __construct(S3ClientInterface $s3Client, array<string|int, mixed> $requestArgs[, array<string|int, mixed> $config = [] ][, string|null $uploadId = null ][, array<string|int, mixed> $partsCompleted = [] ][, TransferProgressSnapshot|null $currentSnapshot = null ][, TransferListenerNotifier|null $listenerNotifier = null ]) : mixed`

##### Parameters

$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)$requestArgs
: array<string\|int, mixed>$config
: array<string\|int, mixed>
= \[\]

- target\_part\_size\_bytes: (int, optional)
- concurrency: (int, optional)

$uploadId
: string\|null
= null$partsCompleted
: array<string\|int, mixed>
= \[\]$currentSnapshot
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null
= null$listenerNotifier
: [TransferListenerNotifier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html) \|null
= null

#### getCurrentSnapshot()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#method_getCurrentSnapshot)

Get the current progress snapshot.

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot|null`

##### Return values

[TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null

#### getPartsCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#method_getPartsCompleted)

`
    public
                    getPartsCompleted() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getUploadId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#method_getUploadId)

`
    public
                    getUploadId() : string|null`

##### Return values

string\|null

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html\#method_promise)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#toc-methods)
- Constants
  - [PART\_MAX\_NUM](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#constant_PART_MAX_NUM)
  - [PART\_MAX\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#constant_PART_MAX_SIZE)
  - [PART\_MIN\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#constant_PART_MIN_SIZE)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method___construct)
  - [getCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method_getCurrentSnapshot)
  - [getPartsCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method_getPartsCompleted)
  - [getUploadId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method_getUploadId)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#method_promise)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartUploader.html#top)
