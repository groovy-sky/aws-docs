Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)

## AbstractMultipartDownloader        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

AbstractYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#toc-interfaces)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#toc-constants)

[GET\_OBJECT\_COMMAND](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#constant_GET_OBJECT_COMMAND)
= "GetObject" [PART\_GET\_MULTIPART\_DOWNLOADER](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#constant_PART_GET_MULTIPART_DOWNLOADER)
= "part" [RANGED\_GET\_MULTIPART\_DOWNLOADER](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#constant_RANGED_GET_MULTIPART_DOWNLOADER)
= "ranged"

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method___construct)
: mixed [chooseDownloaderClass()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_chooseDownloaderClass)
: string [computeObjectSizeFromContentRange()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_computeObjectSizeFromContentRange)
: int Calculates the object size from content range.[download()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_download)
: [DownloadResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadResult.html)[getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getConfig)
: array<string\|int, mixed> [getCurrentPartNo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getCurrentPartNo)
: int [getCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getCurrentSnapshot)
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html)[getObjectPartsCount()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getObjectPartsCount)
: int [getObjectSizeInBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getObjectSizeInBytes)
: int [getRangeTo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getRangeTo)
: int [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns that resolves a multipart download operation,
or to a rejection in case of any failures.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#constants)

#### GET\_OBJECT\_COMMAND  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#constant_GET_OBJECT_COMMAND)

`
    public
        mixed
    GET_OBJECT_COMMAND
    = "GetObject"
`

#### PART\_GET\_MULTIPART\_DOWNLOADER  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#constant_PART_GET_MULTIPART_DOWNLOADER)

`
    public
        mixed
    PART_GET_MULTIPART_DOWNLOADER
    = "part"
`

#### RANGED\_GET\_MULTIPART\_DOWNLOADER  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#constant_RANGED_GET_MULTIPART_DOWNLOADER)

`
    public
        mixed
    RANGED_GET_MULTIPART_DOWNLOADER
    = "ranged"
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method___construct)

`
    public
                    __construct(S3ClientInterface $s3Client, array<string|int, mixed> $downloadRequestArgs[, array<string|int, mixed> $config = [] ][, AbstractDownloadHandler|null $downloadHandler = null ][, array<string|int, mixed> $partsCompleted = [] ][, int $objectPartsCount = 0 ][, int $objectSizeInBytes = 0 ][, string|null $eTag = null ][, TransferProgressSnapshot|null $currentSnapshot = null ][, TransferListenerNotifier|null $listenerNotifier = null ][, ResumableDownload|null $resumableDownload = null ]) : mixed`

##### Parameters

$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)$downloadRequestArgs
: array<string\|int, mixed>$config
: array<string\|int, mixed>
= \[\]$downloadHandler
: [AbstractDownloadHandler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.AbstractDownloadHandler.html) \|null
= null$partsCompleted
: array<string\|int, mixed>
= \[\]$objectPartsCount
: int
= 0$objectSizeInBytes
: int
= 0$eTag
: string\|null
= null$currentSnapshot
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null
= null$listenerNotifier
: [TransferListenerNotifier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html) \|null
= null$resumableDownload
: [ResumableDownload](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html) \|null
= null

#### chooseDownloaderClass()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_chooseDownloaderClass)

`
    public
            static        chooseDownloaderClass(mixed $multipartDownloadType) : string`

##### Parameters

$multipartDownloadType
: mixed

##### Return values

string

#### computeObjectSizeFromContentRange()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_computeObjectSizeFromContentRange)

Calculates the object size from content range.

`
    public
            static        computeObjectSizeFromContentRange(string $contentRange) : int`

##### Parameters

$contentRange
: string

##### Return values

int

#### download()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_download)

`
    public
                    download() : DownloadResult`

##### Return values

[DownloadResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadResult.html)

#### getConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_getConfig)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCurrentPartNo()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_getCurrentPartNo)

`
    public
                    getCurrentPartNo() : int`

##### Return values

int

#### getCurrentSnapshot()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_getCurrentSnapshot)

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot`

##### Return values

[TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html)

#### getObjectPartsCount()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_getObjectPartsCount)

`
    public
                    getObjectPartsCount() : int`

##### Return values

int

#### getObjectSizeInBytes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_getObjectSizeInBytes)

`
    public
                    getObjectSizeInBytes() : int`

##### Return values

int

#### getRangeTo()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_getRangeTo)

`
    public
            static        getRangeTo(string $range) : int`

##### Parameters

$range
: string

##### Return values

int

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html\#method_promise)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#toc-methods)
- Constants
  - [GET\_OBJECT\_COMMAND](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#constant_GET_OBJECT_COMMAND)
  - [PART\_GET\_MULTIPART\_DOWNLOADER](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#constant_PART_GET_MULTIPART_DOWNLOADER)
  - [RANGED\_GET\_MULTIPART\_DOWNLOADER](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#constant_RANGED_GET_MULTIPART_DOWNLOADER)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method___construct)
  - [chooseDownloaderClass()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_chooseDownloaderClass)
  - [computeObjectSizeFromContentRange()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_computeObjectSizeFromContentRange)
  - [download()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_download)
  - [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getConfig)
  - [getCurrentPartNo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getCurrentPartNo)
  - [getCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getCurrentSnapshot)
  - [getObjectPartsCount()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getObjectPartsCount)
  - [getObjectSizeInBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getObjectSizeInBytes)
  - [getRangeTo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_getRangeTo)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#method_promise)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.AbstractMultipartDownloader.html#top)
