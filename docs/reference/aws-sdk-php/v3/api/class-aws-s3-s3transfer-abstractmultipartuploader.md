Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)

## AbstractMultipartUploader        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

AbstractYes

Abstract base class for multipart operations

### Table of Contents  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Constants  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-toc-constants.md)

[PART\_MAX\_NUM](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-max-num.md)
= 10000 [PART\_MAX\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-max-size.md)
= 5 \* 1024 \* 1024 \* 1024 [PART\_MIN\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-min-size.md)
= 5 \* 1024 \* 1024

#### Methods  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-abstractmultipartuploader-method-construct.md)
: mixed [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartuploader-method-getcurrentsnapshot.md)
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null Get the current progress snapshot.[getPartsCompleted()](class-aws-s3-s3transfer-abstractmultipartuploader-method-getpartscompleted.md)
: array<string\|int, mixed> [getUploadId()](class-aws-s3-s3transfer-abstractmultipartuploader-method-getuploadid.md)
: string\|null [promise()](class-aws-s3-s3transfer-abstractmultipartuploader-method-promise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Constants  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-constants.md)

#### PART\_MAX\_NUM  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-max-num.md)

`
    public
        mixed
    PART_MAX_NUM
    = 10000
`

#### PART\_MAX\_SIZE  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-max-size.md)

`
    public
        mixed
    PART_MAX_SIZE
    = 5 * 1024 * 1024 * 1024
`

#### PART\_MIN\_SIZE  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-min-size.md)

`
    public
        mixed
    PART_MIN_SIZE
    = 5 * 1024 * 1024
`

### Methods  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-method-construct.md)

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
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null
= null$listenerNotifier
: [TransferListenerNotifier](class-aws-s3-s3transfer-progress-transferlistenernotifier.md) \|null
= null

#### getCurrentSnapshot()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-method-getcurrentsnapshot.md)

Get the current progress snapshot.

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot|null`

##### Return values

[TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null

#### getPartsCompleted()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-method-getpartscompleted.md)

`
    public
                    getPartsCompleted() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getUploadId()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-method-getuploadid.md)

`
    public
                    getUploadId() : string|null`

##### Return values

string\|null

#### promise()  [header link](class-aws-s3-s3transfer-abstractmultipartuploader-method-promise.md)

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
  - [Constants](class-aws-s3-s3transfer-abstractmultipartuploader-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-abstractmultipartuploader-toc-methods.md)
- Constants
  - [PART\_MAX\_NUM](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-max-num.md)
  - [PART\_MAX\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-max-size.md)
  - [PART\_MIN\_SIZE](class-aws-s3-s3transfer-abstractmultipartuploader-constant-part-min-size.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-abstractmultipartuploader-method-construct.md)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartuploader-method-getcurrentsnapshot.md)
  - [getPartsCompleted()](class-aws-s3-s3transfer-abstractmultipartuploader-method-getpartscompleted.md)
  - [getUploadId()](class-aws-s3-s3transfer-abstractmultipartuploader-method-getuploadid.md)
  - [promise()](class-aws-s3-s3transfer-abstractmultipartuploader-method-promise.md)

[Back To Top](class-aws-s3-s3transfer-abstractmultipartuploader-top.md)
