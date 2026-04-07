Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## ResumableDownload     extends [AbstractResumableTransfer](class-aws-s3-s3transfer-models-abstractresumabletransfer.md)   in package    - [Aws](package-aws.md)

FinalYes

Represents the state of a resumable multipart download.

This class can be serialized to/from JSON to persist download progress.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method___construct)
: mixed [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_deleteResumeFile)
: void [fromFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_fromFile)
: self Load a resumable state from a file.[fromJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_fromJson)
: self Deserialize a resumable download state from JSON format.[getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getBucket)
: string [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getConfig)
: array<string\|int, mixed> [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getCurrentSnapshot)
: array<string\|int, mixed> [getDestination()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getDestination)
: string [getETag()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getETag)
: string [getFixedPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getFixedPartSize)
: int [getInitialRequestResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getInitialRequestResult)
: array<string\|int, mixed> [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getKey)
: string [getObjectSizeInBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getObjectSizeInBytes)
: int [getPartsCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getPartsCompleted)
: array<string\|int, mixed> [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getRequestArgs)
: array<string\|int, mixed> [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getResumeFilePath)
: string [getTemporaryFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getTemporaryFile)
: string\|null [getTotalNumberOfParts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getTotalNumberOfParts)
: int [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_isResumeFile)
: bool Check if a file path is a valid resume file.[markPartCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_markPartCompleted)
: void Mark a part as completed.[toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_toFile)
: void Save the resumable state to a file.[toJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_toJson)
: string Serialize the resumable download state to JSON format.[updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_updateCurrentSnapshot)
: void Update the current snapshot.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method___construct)

`
    public
                    __construct(string $resumeFilePath, array<string|int, mixed> $requestArgs, array<string|int, mixed> $config, array<string|int, mixed> $currentSnapshot, array<string|int, mixed> $initialRequestResult, array<string|int, mixed> $partsCompleted, int $totalNumberOfParts, string|null $temporaryFile, string $eTag, int $objectSizeInBytes, int $fixedPartSize, string $destination) : mixed`

##### Parameters

$resumeFilePath
: string$requestArgs
: array<string\|int, mixed>

The request arguments used for the download

$config
: array<string\|int, mixed>

The config used in the request

$currentSnapshot
: array<string\|int, mixed>

The current progress snapshot

$initialRequestResult
: array<string\|int, mixed>

The response from the initial request

$partsCompleted
: array<string\|int, mixed>

Map of completed part numbers (partNo => true)

$totalNumberOfParts
: int

Total number of parts in the download

$temporaryFile
: string\|null

Path to the temporary file being downloaded to

$eTag
: string

ETag of the S3 object for consistency verification

$objectSizeInBytes
: int

Total size of the object in bytes

$fixedPartSize
: int

Size of each part in bytes

$destination
: string

Final destination path for the downloaded file

#### deleteResumeFile()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_deleteResumeFile)

`
    public
                    deleteResumeFile([string|null $filePath = null ]) : void`

##### Parameters

$filePath
: string\|null
= null

#### fromFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_fromFile)

Load a resumable state from a file.

`
    public
            static        fromFile(string $filePath) : self`

##### Parameters

$filePath
: string

##### Return values

self

#### fromJson()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_fromJson)

Deserialize a resumable download state from JSON format.

`
    public
            static        fromJson(string $json) : self`

##### Parameters

$json
: string

JSON-encoded state

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_fromJson\#tags)

throws[S3TransferException](class-aws-s3-s3transfer-exception-s3transferexception.md)

If the JSON is invalid or missing required fields

##### Return values

self

#### getBucket()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_getBucket)

`
    public
                    getBucket() : string`

##### Return values

string

#### getConfig()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_getConfig)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCurrentSnapshot()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_getCurrentSnapshot)

`
    public
                    getCurrentSnapshot() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getDestination()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_getDestination)

`
    public
                    getDestination() : string`

##### Return values

string

#### getETag()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_getETag)

`
    public
                    getETag() : string`

##### Return values

string

#### getFixedPartSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_getFixedPartSize)

`
    public
                    getFixedPartSize() : int`

##### Return values

int

#### getInitialRequestResult()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_getInitialRequestResult)

`
    public
                    getInitialRequestResult() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getKey()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_getKey)

`
    public
                    getKey() : string`

##### Return values

string

#### getObjectSizeInBytes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_getObjectSizeInBytes)

`
    public
                    getObjectSizeInBytes() : int`

##### Return values

int

#### getPartsCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_getPartsCompleted)

`
    public
                    getPartsCompleted() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getRequestArgs()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_getRequestArgs)

`
    public
                    getRequestArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getResumeFilePath()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_getResumeFilePath)

`
    public
                    getResumeFilePath() : string`

##### Return values

string

#### getTemporaryFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_getTemporaryFile)

`
    public
                    getTemporaryFile() : string|null`

##### Return values

string\|null

#### getTotalNumberOfParts()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_getTotalNumberOfParts)

`
    public
                    getTotalNumberOfParts() : int`

##### Return values

int

#### isResumeFile()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_isResumeFile)

Check if a file path is a valid resume file.

`
    public
            static        isResumeFile(string $filePath) : bool`

##### Parameters

$filePath
: string

##### Return values

bool

#### markPartCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_markPartCompleted)

Mark a part as completed.

`
    public
                    markPartCompleted(int $partNumber) : void`

##### Parameters

$partNumber
: int

The part number to mark as completed

#### toFile()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_toFile)

Save the resumable state to a file.

`
    public
                    toFile([string|null $filePath = null ]) : void`

When a file path is not provided by default it will use
the `resumeFilePath` property.

##### Parameters

$filePath
: string\|null
= null

Path where the resume file should be saved

#### toJson()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html\#method_toJson)

Serialize the resumable download state to JSON format.

`
    public
                    toJson() : string`

##### Return values

string
—

JSON-encoded state

#### updateCurrentSnapshot()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_updateCurrentSnapshot)

Update the current snapshot.

`
    public
                    updateCurrentSnapshot(array<string|int, mixed> $snapshot) : void`

##### Parameters

$snapshot
: array<string\|int, mixed>

The new snapshot data

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method___construct)
  - [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_deleteResumeFile)
  - [fromFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_fromFile)
  - [fromJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_fromJson)
  - [getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getBucket)
  - [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getConfig)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getCurrentSnapshot)
  - [getDestination()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getDestination)
  - [getETag()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getETag)
  - [getFixedPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getFixedPartSize)
  - [getInitialRequestResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getInitialRequestResult)
  - [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getKey)
  - [getObjectSizeInBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getObjectSizeInBytes)
  - [getPartsCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getPartsCompleted)
  - [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getRequestArgs)
  - [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getResumeFilePath)
  - [getTemporaryFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getTemporaryFile)
  - [getTotalNumberOfParts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_getTotalNumberOfParts)
  - [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_isResumeFile)
  - [markPartCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_markPartCompleted)
  - [toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_toFile)
  - [toJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#method_toJson)
  - [updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_updateCurrentSnapshot)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableDownload.html#top)
