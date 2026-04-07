Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## ResumableUpload     extends [AbstractResumableTransfer](class-aws-s3-s3transfer-models-abstractresumabletransfer.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method___construct)
: mixed [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_deleteResumeFile)
: void [fromFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_fromFile)
: self Load a resumable state from a file.[fromJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_fromJson)
: self Deserialize a resumable state from JSON format.[getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getBucket)
: string [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getConfig)
: array<string\|int, mixed> [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getCurrentSnapshot)
: array<string\|int, mixed> [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getKey)
: string [getObjectSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getObjectSize)
: int [getPartsCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getPartsCompleted)
: array<string\|int, mixed> [getPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getPartSize)
: int [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getRequestArgs)
: array<string\|int, mixed> [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getResumeFilePath)
: string [getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getSource)
: string [getUploadId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getUploadId)
: string [isFullObjectChecksum()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_isFullObjectChecksum)
: bool [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_isResumeFile)
: bool Check if a file path is a valid resume file.[markPartCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_markPartCompleted)
: void Mark a part as completed.[toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_toFile)
: void Save the resumable state to a file.[toJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_toJson)
: string Serialize the resumable state to JSON format.[updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_updateCurrentSnapshot)
: void Update the current snapshot.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method___construct)

`
    public
                    __construct(string $resumeFilePath, array<string|int, mixed> $requestArgs, array<string|int, mixed> $config, array<string|int, mixed> $currentSnapshot, string $uploadId, array<string|int, mixed> $partsCompleted, string $source, int $objectSize, int $partSize, bool $isFullObjectChecksum) : mixed`

##### Parameters

$resumeFilePath
: string$requestArgs
: array<string\|int, mixed>$config
: array<string\|int, mixed>$currentSnapshot
: array<string\|int, mixed>$uploadId
: string$partsCompleted
: array<string\|int, mixed>$source
: string$objectSize
: int$partSize
: int$isFullObjectChecksum
: bool

#### deleteResumeFile()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_deleteResumeFile)

`
    public
                    deleteResumeFile([string|null $filePath = null ]) : void`

##### Parameters

$filePath
: string\|null
= null

#### fromFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_fromFile)

Load a resumable state from a file.

`
    public
            static        fromFile(string $filePath) : self`

##### Parameters

$filePath
: string

##### Return values

self

#### fromJson()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_fromJson)

Deserialize a resumable state from JSON format.

`
    public
            static        fromJson(string $json) : self`

##### Parameters

$json
: string

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

#### getKey()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer.md\#method_getKey)

`
    public
                    getKey() : string`

##### Return values

string

#### getObjectSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_getObjectSize)

`
    public
                    getObjectSize() : int`

##### Return values

int

#### getPartsCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_getPartsCompleted)

`
    public
                    getPartsCompleted() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getPartSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_getPartSize)

`
    public
                    getPartSize() : int`

##### Return values

int

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

#### getSource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_getSource)

`
    public
                    getSource() : string`

##### Return values

string

#### getUploadId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_getUploadId)

`
    public
                    getUploadId() : string`

##### Return values

string

#### isFullObjectChecksum()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_isFullObjectChecksum)

`
    public
                    isFullObjectChecksum() : bool`

##### Return values

bool

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

#### markPartCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_markPartCompleted)

Mark a part as completed.

`
    public
                    markPartCompleted(int $partNumber, array<string|int, mixed> $part) : void`

##### Parameters

$partNumber
: int

The part number to mark as completed

$part
: array<string\|int, mixed>

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

#### toJson()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html\#method_toJson)

Serialize the resumable state to JSON format.

`
    public
                    toJson() : string`

##### Return values

string

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method___construct)
  - [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_deleteResumeFile)
  - [fromFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_fromFile)
  - [fromJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_fromJson)
  - [getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getBucket)
  - [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getConfig)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getCurrentSnapshot)
  - [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getKey)
  - [getObjectSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getObjectSize)
  - [getPartsCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getPartsCompleted)
  - [getPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getPartSize)
  - [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getRequestArgs)
  - [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getResumeFilePath)
  - [getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getSource)
  - [getUploadId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_getUploadId)
  - [isFullObjectChecksum()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_isFullObjectChecksum)
  - [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_isResumeFile)
  - [markPartCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_markPartCompleted)
  - [toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_toFile)
  - [toJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#method_toJson)
  - [updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_updateCurrentSnapshot)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumableUpload.html#top)
