Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)

## RangeGetMultipartDownloader     extends [AbstractMultipartDownloader](class-aws-s3-s3transfer-abstractmultipartdownloader.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-rangegetmultipartdownloader-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-rangegetmultipartdownloader-toc-constants.md)

[GET\_OBJECT\_COMMAND](class-aws-s3-s3transfer-abstractmultipartdownloader.md#constant_GET_OBJECT_COMMAND)
= "GetObject" [PART\_GET\_MULTIPART\_DOWNLOADER](class-aws-s3-s3transfer-abstractmultipartdownloader.md#constant_PART_GET_MULTIPART_DOWNLOADER)
= "part" [RANGED\_GET\_MULTIPART\_DOWNLOADER](class-aws-s3-s3transfer-abstractmultipartdownloader.md#constant_RANGED_GET_MULTIPART_DOWNLOADER)
= "ranged"

#### Methods  [header link](class-aws-s3-s3transfer-rangegetmultipartdownloader-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method___construct)
: mixed [chooseDownloaderClass()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_chooseDownloaderClass)
: string [computeObjectSizeFromContentRange()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_computeObjectSizeFromContentRange)
: int Calculates the object size from content range.[download()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_download)
: [DownloadResult](class-aws-s3-s3transfer-models-downloadresult.md)[getConfig()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getConfig)
: array<string\|int, mixed> [getCurrentPartNo()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getCurrentPartNo)
: int [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getCurrentSnapshot)
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md)[getObjectPartsCount()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getObjectPartsCount)
: int [getObjectSizeInBytes()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getObjectSizeInBytes)
: int [getRangeTo()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getRangeTo)
: int [promise()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns that resolves a multipart download operation,
or to a rejection in case of any failures.

### Constants  [header link](class-aws-s3-s3transfer-rangegetmultipartdownloader-constants.md)

#### GET\_OBJECT\_COMMAND  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#constant_GET_OBJECT_COMMAND)

`
    public
        mixed
    GET_OBJECT_COMMAND
    = "GetObject"
`

#### PART\_GET\_MULTIPART\_DOWNLOADER  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#constant_PART_GET_MULTIPART_DOWNLOADER)

`
    public
        mixed
    PART_GET_MULTIPART_DOWNLOADER
    = "part"
`

#### RANGED\_GET\_MULTIPART\_DOWNLOADER  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#constant_RANGED_GET_MULTIPART_DOWNLOADER)

`
    public
        mixed
    RANGED_GET_MULTIPART_DOWNLOADER
    = "ranged"
`

### Methods  [header link](class-aws-s3-s3transfer-rangegetmultipartdownloader-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method___construct)

`
    public
                    __construct(S3ClientInterface $s3Client, array<string|int, mixed> $downloadRequestArgs[, array<string|int, mixed> $config = [] ][, AbstractDownloadHandler|null $downloadHandler = null ][, array<string|int, mixed> $partsCompleted = [] ][, int $objectPartsCount = 0 ][, int $objectSizeInBytes = 0 ][, string|null $eTag = null ][, TransferProgressSnapshot|null $currentSnapshot = null ][, TransferListenerNotifier|null $listenerNotifier = null ][, ResumableDownload|null $resumableDownload = null ]) : mixed`

##### Parameters

$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)$downloadRequestArgs
: array<string\|int, mixed>$config
: array<string\|int, mixed>
= \[\]$downloadHandler
: [AbstractDownloadHandler](class-aws-s3-s3transfer-utils-abstractdownloadhandler.md) \|null
= null$partsCompleted
: array<string\|int, mixed>
= \[\]$objectPartsCount
: int
= 0$objectSizeInBytes
: int
= 0$eTag
: string\|null
= null$currentSnapshot
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null
= null$listenerNotifier
: [TransferListenerNotifier](class-aws-s3-s3transfer-progress-transferlistenernotifier.md) \|null
= null$resumableDownload
: [ResumableDownload](class-aws-s3-s3transfer-models-resumabledownload.md) \|null
= null

#### chooseDownloaderClass()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_chooseDownloaderClass)

`
    public
            static        chooseDownloaderClass(mixed $multipartDownloadType) : string`

##### Parameters

$multipartDownloadType
: mixed

##### Return values

string

#### computeObjectSizeFromContentRange()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_computeObjectSizeFromContentRange)

Calculates the object size from content range.

`
    public
            static        computeObjectSizeFromContentRange(string $contentRange) : int`

##### Parameters

$contentRange
: string

##### Return values

int

#### download()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_download)

`
    public
                    download() : DownloadResult`

##### Return values

[DownloadResult](class-aws-s3-s3transfer-models-downloadresult.md)

#### getConfig()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_getConfig)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCurrentPartNo()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_getCurrentPartNo)

`
    public
                    getCurrentPartNo() : int`

##### Return values

int

#### getCurrentSnapshot()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_getCurrentSnapshot)

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot`

##### Return values

[TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md)

#### getObjectPartsCount()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_getObjectPartsCount)

`
    public
                    getObjectPartsCount() : int`

##### Return values

int

#### getObjectSizeInBytes()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_getObjectSizeInBytes)

`
    public
                    getObjectSizeInBytes() : int`

##### Return values

int

#### getRangeTo()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_getRangeTo)

`
    public
            static        getRangeTo(string $range) : int`

##### Parameters

$range
: string

##### Return values

int

#### promise()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader.md\#method_promise)

Returns that resolves a multipart download operation,
or to a rejection in case of any failures.

`
    public
                    promise() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3transfer-rangegetmultipartdownloader-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-rangegetmultipartdownloader-toc-methods.md)
- Constants
  - [GET\_OBJECT\_COMMAND](class-aws-s3-s3transfer-abstractmultipartdownloader.md#constant_GET_OBJECT_COMMAND)
  - [PART\_GET\_MULTIPART\_DOWNLOADER](class-aws-s3-s3transfer-abstractmultipartdownloader.md#constant_PART_GET_MULTIPART_DOWNLOADER)
  - [RANGED\_GET\_MULTIPART\_DOWNLOADER](class-aws-s3-s3transfer-abstractmultipartdownloader.md#constant_RANGED_GET_MULTIPART_DOWNLOADER)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method___construct)
  - [chooseDownloaderClass()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_chooseDownloaderClass)
  - [computeObjectSizeFromContentRange()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_computeObjectSizeFromContentRange)
  - [download()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_download)
  - [getConfig()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getConfig)
  - [getCurrentPartNo()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getCurrentPartNo)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getCurrentSnapshot)
  - [getObjectPartsCount()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getObjectPartsCount)
  - [getObjectSizeInBytes()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getObjectSizeInBytes)
  - [getRangeTo()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_getRangeTo)
  - [promise()](class-aws-s3-s3transfer-abstractmultipartdownloader.md#method_promise)

[Back To Top](class-aws-s3-s3transfer-rangegetmultipartdownloader-top.md)
