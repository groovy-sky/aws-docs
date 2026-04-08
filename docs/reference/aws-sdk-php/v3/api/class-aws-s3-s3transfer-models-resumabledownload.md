Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## ResumableDownload     extends [AbstractResumableTransfer](class-aws-s3-s3transfer-models-abstractresumabletransfer.md)   in package    - [Aws](package-aws.md)

FinalYes

Represents the state of a resumable multipart download.

This class can be serialized to/from JSON to persist download progress.

### Table of Contents  [header link](class-aws-s3-s3transfer-models-resumabledownload-toc.md)

#### Methods  [header link](class-aws-s3-s3transfer-models-resumabledownload-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-resumabledownload-method-construct.md)
: mixed [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_deleteResumeFile)
: void [fromFile()](class-aws-s3-s3transfer-models-resumabledownload-method-fromfile.md)
: self Load a resumable state from a file.[fromJson()](class-aws-s3-s3transfer-models-resumabledownload-method-fromjson.md)
: self Deserialize a resumable download state from JSON format.[getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getBucket)
: string [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getConfig)
: array<string\|int, mixed> [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getCurrentSnapshot)
: array<string\|int, mixed> [getDestination()](class-aws-s3-s3transfer-models-resumabledownload-method-getdestination.md)
: string [getETag()](class-aws-s3-s3transfer-models-resumabledownload-method-getetag.md)
: string [getFixedPartSize()](class-aws-s3-s3transfer-models-resumabledownload-method-getfixedpartsize.md)
: int [getInitialRequestResult()](class-aws-s3-s3transfer-models-resumabledownload-method-getinitialrequestresult.md)
: array<string\|int, mixed> [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getKey)
: string [getObjectSizeInBytes()](class-aws-s3-s3transfer-models-resumabledownload-method-getobjectsizeinbytes.md)
: int [getPartsCompleted()](class-aws-s3-s3transfer-models-resumabledownload-method-getpartscompleted.md)
: array<string\|int, mixed> [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getRequestArgs)
: array<string\|int, mixed> [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getResumeFilePath)
: string [getTemporaryFile()](class-aws-s3-s3transfer-models-resumabledownload-method-gettemporaryfile.md)
: string\|null [getTotalNumberOfParts()](class-aws-s3-s3transfer-models-resumabledownload-method-gettotalnumberofparts.md)
: int [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_isResumeFile)
: bool Check if a file path is a valid resume file.[markPartCompleted()](class-aws-s3-s3transfer-models-resumabledownload-method-markpartcompleted.md)
: void Mark a part as completed.[toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_toFile)
: void Save the resumable state to a file.[toJson()](class-aws-s3-s3transfer-models-resumabledownload-method-tojson.md)
: string Serialize the resumable download state to JSON format.[updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_updateCurrentSnapshot)
: void Update the current snapshot.

### Methods  [header link](class-aws-s3-s3transfer-models-resumabledownload-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-construct.md)

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

#### fromFile()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-fromfile.md)

Load a resumable state from a file.

`
    public
            static        fromFile(string $filePath) : self`

##### Parameters

$filePath
: string

##### Return values

self

#### fromJson()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-fromjson.md)

Deserialize a resumable download state from JSON format.

`
    public
            static        fromJson(string $json) : self`

##### Parameters

$json
: string

JSON-encoded state

##### Tags  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-fromjson-tags.md)

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

#### getDestination()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-getdestination.md)

`
    public
                    getDestination() : string`

##### Return values

string

#### getETag()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-getetag.md)

`
    public
                    getETag() : string`

##### Return values

string

#### getFixedPartSize()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-getfixedpartsize.md)

`
    public
                    getFixedPartSize() : int`

##### Return values

int

#### getInitialRequestResult()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-getinitialrequestresult.md)

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

#### getObjectSizeInBytes()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-getobjectsizeinbytes.md)

`
    public
                    getObjectSizeInBytes() : int`

##### Return values

int

#### getPartsCompleted()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-getpartscompleted.md)

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

#### getTemporaryFile()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-gettemporaryfile.md)

`
    public
                    getTemporaryFile() : string|null`

##### Return values

string\|null

#### getTotalNumberOfParts()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-gettotalnumberofparts.md)

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

#### markPartCompleted()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-markpartcompleted.md)

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

#### toJson()  [header link](class-aws-s3-s3transfer-models-resumabledownload-method-tojson.md)

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
  - [Methods](class-aws-s3-s3transfer-models-resumabledownload-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-resumabledownload-method-construct.md)
  - [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_deleteResumeFile)
  - [fromFile()](class-aws-s3-s3transfer-models-resumabledownload-method-fromfile.md)
  - [fromJson()](class-aws-s3-s3transfer-models-resumabledownload-method-fromjson.md)
  - [getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getBucket)
  - [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getConfig)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getCurrentSnapshot)
  - [getDestination()](class-aws-s3-s3transfer-models-resumabledownload-method-getdestination.md)
  - [getETag()](class-aws-s3-s3transfer-models-resumabledownload-method-getetag.md)
  - [getFixedPartSize()](class-aws-s3-s3transfer-models-resumabledownload-method-getfixedpartsize.md)
  - [getInitialRequestResult()](class-aws-s3-s3transfer-models-resumabledownload-method-getinitialrequestresult.md)
  - [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getKey)
  - [getObjectSizeInBytes()](class-aws-s3-s3transfer-models-resumabledownload-method-getobjectsizeinbytes.md)
  - [getPartsCompleted()](class-aws-s3-s3transfer-models-resumabledownload-method-getpartscompleted.md)
  - [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getRequestArgs)
  - [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getResumeFilePath)
  - [getTemporaryFile()](class-aws-s3-s3transfer-models-resumabledownload-method-gettemporaryfile.md)
  - [getTotalNumberOfParts()](class-aws-s3-s3transfer-models-resumabledownload-method-gettotalnumberofparts.md)
  - [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_isResumeFile)
  - [markPartCompleted()](class-aws-s3-s3transfer-models-resumabledownload-method-markpartcompleted.md)
  - [toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_toFile)
  - [toJson()](class-aws-s3-s3transfer-models-resumabledownload-method-tojson.md)
  - [updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_updateCurrentSnapshot)

[Back To Top](class-aws-s3-s3transfer-models-resumabledownload-top.md)
