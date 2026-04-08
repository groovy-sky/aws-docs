Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)

## AbstractMultipartDownloader        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

AbstractYes

### Table of Contents  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Constants  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-toc-constants.md)

[GET\_OBJECT\_COMMAND](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-get-object-command.md)
= "GetObject" [PART\_GET\_MULTIPART\_DOWNLOADER](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-part-get-multipart-downloader.md)
= "part" [RANGED\_GET\_MULTIPART\_DOWNLOADER](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-ranged-get-multipart-downloader.md)
= "ranged"

#### Methods  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-construct.md)
: mixed [chooseDownloaderClass()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-choosedownloaderclass.md)
: string [computeObjectSizeFromContentRange()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-computeobjectsizefromcontentrange.md)
: int Calculates the object size from content range.[download()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-download.md)
: [DownloadResult](class-aws-s3-s3transfer-models-downloadresult.md)[getConfig()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getconfig.md)
: array<string\|int, mixed> [getCurrentPartNo()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getcurrentpartno.md)
: int [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getcurrentsnapshot.md)
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md)[getObjectPartsCount()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getobjectpartscount.md)
: int [getObjectSizeInBytes()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getobjectsizeinbytes.md)
: int [getRangeTo()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getrangeto.md)
: int [promise()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-promise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns that resolves a multipart download operation,
or to a rejection in case of any failures.

### Constants  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-constants.md)

#### GET\_OBJECT\_COMMAND  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-get-object-command.md)

`
    public
        mixed
    GET_OBJECT_COMMAND
    = "GetObject"
`

#### PART\_GET\_MULTIPART\_DOWNLOADER  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-part-get-multipart-downloader.md)

`
    public
        mixed
    PART_GET_MULTIPART_DOWNLOADER
    = "part"
`

#### RANGED\_GET\_MULTIPART\_DOWNLOADER  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-ranged-get-multipart-downloader.md)

`
    public
        mixed
    RANGED_GET_MULTIPART_DOWNLOADER
    = "ranged"
`

### Methods  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-construct.md)

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

#### chooseDownloaderClass()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-choosedownloaderclass.md)

`
    public
            static        chooseDownloaderClass(mixed $multipartDownloadType) : string`

##### Parameters

$multipartDownloadType
: mixed

##### Return values

string

#### computeObjectSizeFromContentRange()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-computeobjectsizefromcontentrange.md)

Calculates the object size from content range.

`
    public
            static        computeObjectSizeFromContentRange(string $contentRange) : int`

##### Parameters

$contentRange
: string

##### Return values

int

#### download()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-download.md)

`
    public
                    download() : DownloadResult`

##### Return values

[DownloadResult](class-aws-s3-s3transfer-models-downloadresult.md)

#### getConfig()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getconfig.md)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCurrentPartNo()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getcurrentpartno.md)

`
    public
                    getCurrentPartNo() : int`

##### Return values

int

#### getCurrentSnapshot()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getcurrentsnapshot.md)

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot`

##### Return values

[TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md)

#### getObjectPartsCount()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getobjectpartscount.md)

`
    public
                    getObjectPartsCount() : int`

##### Return values

int

#### getObjectSizeInBytes()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getobjectsizeinbytes.md)

`
    public
                    getObjectSizeInBytes() : int`

##### Return values

int

#### getRangeTo()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getrangeto.md)

`
    public
            static        getRangeTo(string $range) : int`

##### Parameters

$range
: string

##### Return values

int

#### promise()  [header link](class-aws-s3-s3transfer-abstractmultipartdownloader-method-promise.md)

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
  - [Constants](class-aws-s3-s3transfer-abstractmultipartdownloader-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-abstractmultipartdownloader-toc-methods.md)
- Constants
  - [GET\_OBJECT\_COMMAND](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-get-object-command.md)
  - [PART\_GET\_MULTIPART\_DOWNLOADER](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-part-get-multipart-downloader.md)
  - [RANGED\_GET\_MULTIPART\_DOWNLOADER](class-aws-s3-s3transfer-abstractmultipartdownloader-constant-ranged-get-multipart-downloader.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-construct.md)
  - [chooseDownloaderClass()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-choosedownloaderclass.md)
  - [computeObjectSizeFromContentRange()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-computeobjectsizefromcontentrange.md)
  - [download()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-download.md)
  - [getConfig()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getconfig.md)
  - [getCurrentPartNo()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getcurrentpartno.md)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getcurrentsnapshot.md)
  - [getObjectPartsCount()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getobjectpartscount.md)
  - [getObjectSizeInBytes()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getobjectsizeinbytes.md)
  - [getRangeTo()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-getrangeto.md)
  - [promise()](class-aws-s3-s3transfer-abstractmultipartdownloader-method-promise.md)

[Back To Top](class-aws-s3-s3transfer-abstractmultipartdownloader-top.md)
