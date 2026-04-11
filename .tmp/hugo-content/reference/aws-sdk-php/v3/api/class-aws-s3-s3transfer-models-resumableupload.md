Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## ResumableUpload     extends [AbstractResumableTransfer](class-aws-s3-s3transfer-models-abstractresumabletransfer.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-models-resumableupload-toc.md)

#### Methods  [header link](class-aws-s3-s3transfer-models-resumableupload-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-resumableupload-method-construct.md)
: mixed [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_deleteResumeFile)
: void [fromFile()](class-aws-s3-s3transfer-models-resumableupload-method-fromfile.md)
: self Load a resumable state from a file.[fromJson()](class-aws-s3-s3transfer-models-resumableupload-method-fromjson.md)
: self Deserialize a resumable state from JSON format.[getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getBucket)
: string [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getConfig)
: array<string\|int, mixed> [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getCurrentSnapshot)
: array<string\|int, mixed> [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getKey)
: string [getObjectSize()](class-aws-s3-s3transfer-models-resumableupload-method-getobjectsize.md)
: int [getPartsCompleted()](class-aws-s3-s3transfer-models-resumableupload-method-getpartscompleted.md)
: array<string\|int, mixed> [getPartSize()](class-aws-s3-s3transfer-models-resumableupload-method-getpartsize.md)
: int [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getRequestArgs)
: array<string\|int, mixed> [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getResumeFilePath)
: string [getSource()](class-aws-s3-s3transfer-models-resumableupload-method-getsource.md)
: string [getUploadId()](class-aws-s3-s3transfer-models-resumableupload-method-getuploadid.md)
: string [isFullObjectChecksum()](class-aws-s3-s3transfer-models-resumableupload-method-isfullobjectchecksum.md)
: bool [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_isResumeFile)
: bool Check if a file path is a valid resume file.[markPartCompleted()](class-aws-s3-s3transfer-models-resumableupload-method-markpartcompleted.md)
: void Mark a part as completed.[toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_toFile)
: void Save the resumable state to a file.[toJson()](class-aws-s3-s3transfer-models-resumableupload-method-tojson.md)
: string Serialize the resumable state to JSON format.[updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_updateCurrentSnapshot)
: void Update the current snapshot.

### Methods  [header link](class-aws-s3-s3transfer-models-resumableupload-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-construct.md)

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

#### fromFile()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-fromfile.md)

Load a resumable state from a file.

`
    public
            static        fromFile(string $filePath) : self`

##### Parameters

$filePath
: string

##### Return values

self

#### fromJson()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-fromjson.md)

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

#### getObjectSize()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-getobjectsize.md)

`
    public
                    getObjectSize() : int`

##### Return values

int

#### getPartsCompleted()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-getpartscompleted.md)

`
    public
                    getPartsCompleted() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getPartSize()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-getpartsize.md)

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

#### getSource()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-getsource.md)

`
    public
                    getSource() : string`

##### Return values

string

#### getUploadId()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-getuploadid.md)

`
    public
                    getUploadId() : string`

##### Return values

string

#### isFullObjectChecksum()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-isfullobjectchecksum.md)

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

#### markPartCompleted()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-markpartcompleted.md)

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

#### toJson()  [header link](class-aws-s3-s3transfer-models-resumableupload-method-tojson.md)

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
  - [Methods](class-aws-s3-s3transfer-models-resumableupload-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-resumableupload-method-construct.md)
  - [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_deleteResumeFile)
  - [fromFile()](class-aws-s3-s3transfer-models-resumableupload-method-fromfile.md)
  - [fromJson()](class-aws-s3-s3transfer-models-resumableupload-method-fromjson.md)
  - [getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getBucket)
  - [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getConfig)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getCurrentSnapshot)
  - [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getKey)
  - [getObjectSize()](class-aws-s3-s3transfer-models-resumableupload-method-getobjectsize.md)
  - [getPartsCompleted()](class-aws-s3-s3transfer-models-resumableupload-method-getpartscompleted.md)
  - [getPartSize()](class-aws-s3-s3transfer-models-resumableupload-method-getpartsize.md)
  - [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getRequestArgs)
  - [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_getResumeFilePath)
  - [getSource()](class-aws-s3-s3transfer-models-resumableupload-method-getsource.md)
  - [getUploadId()](class-aws-s3-s3transfer-models-resumableupload-method-getuploadid.md)
  - [isFullObjectChecksum()](class-aws-s3-s3transfer-models-resumableupload-method-isfullobjectchecksum.md)
  - [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_isResumeFile)
  - [markPartCompleted()](class-aws-s3-s3transfer-models-resumableupload-method-markpartcompleted.md)
  - [toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_toFile)
  - [toJson()](class-aws-s3-s3transfer-models-resumableupload-method-tojson.md)
  - [updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer.md#method_updateCurrentSnapshot)

[Back To Top](class-aws-s3-s3transfer-models-resumableupload-top.md)
