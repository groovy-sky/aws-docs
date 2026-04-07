Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## AbstractResumableTransfer        in package    - [Aws](package-aws.md)

AbstractYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method___construct)
: mixed [deleteResumeFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_deleteResumeFile)
: void [fromFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_fromFile)
: self Load a resumable state from a file.[fromJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_fromJson)
: self Deserialize a resumable state from JSON format.[getBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getBucket)
: string [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getConfig)
: array<string\|int, mixed> [getCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getCurrentSnapshot)
: array<string\|int, mixed> [getKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getKey)
: string [getRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getRequestArgs)
: array<string\|int, mixed> [getResumeFilePath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getResumeFilePath)
: string [isResumeFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_isResumeFile)
: bool Check if a file path is a valid resume file.[toFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_toFile)
: void Save the resumable state to a file.[toJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_toJson)
: string Serialize the resumable state to JSON format.[updateCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_updateCurrentSnapshot)
: void Update the current snapshot.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method___construct)

`
    public
                    __construct(string $resumeFilePath, array<string|int, mixed> $requestArgs, array<string|int, mixed> $config, array<string|int, mixed> $currentSnapshot) : mixed`

##### Parameters

$resumeFilePath
: string$requestArgs
: array<string\|int, mixed>

The request arguments used for the transfer

$config
: array<string\|int, mixed>

The config used in the request

$currentSnapshot
: array<string\|int, mixed>

The current progress snapshot

#### deleteResumeFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_deleteResumeFile)

`
    public
                    deleteResumeFile([string|null $filePath = null ]) : void`

##### Parameters

$filePath
: string\|null
= null

#### fromFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_fromFile)

Load a resumable state from a file.

`
    public
    abstract        static        fromFile(string $filePath) : self`

##### Parameters

$filePath
: string

Path to the resume file

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_fromFile\#tags)

throws[S3TransferException](class-aws-s3-s3transfer-exception-s3transferexception.md)

If the file cannot be read or is invalid

##### Return values

self

#### fromJson()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_fromJson)

Deserialize a resumable state from JSON format.

`
    public
    abstract        static        fromJson(string $json) : self`

##### Parameters

$json
: string

JSON-encoded state

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_fromJson\#tags)

throws[S3TransferException](class-aws-s3-s3transfer-exception-s3transferexception.md)

If the JSON is invalid or missing required fields

##### Return values

self

#### getBucket()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_getBucket)

`
    public
                    getBucket() : string`

##### Return values

string

#### getConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_getConfig)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCurrentSnapshot()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_getCurrentSnapshot)

`
    public
                    getCurrentSnapshot() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_getKey)

`
    public
                    getKey() : string`

##### Return values

string

#### getRequestArgs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_getRequestArgs)

`
    public
                    getRequestArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getResumeFilePath()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_getResumeFilePath)

`
    public
                    getResumeFilePath() : string`

##### Return values

string

#### isResumeFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_isResumeFile)

Check if a file path is a valid resume file.

`
    public
            static        isResumeFile(string $filePath) : bool`

##### Parameters

$filePath
: string

##### Return values

bool

#### toFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_toFile)

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

#### toJson()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_toJson)

Serialize the resumable state to JSON format.

`
    public
    abstract                toJson() : string`

##### Return values

string
—

JSON-encoded state

#### updateCurrentSnapshot()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html\#method_updateCurrentSnapshot)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method___construct)
  - [deleteResumeFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_deleteResumeFile)
  - [fromFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_fromFile)
  - [fromJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_fromJson)
  - [getBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getBucket)
  - [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getConfig)
  - [getCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getCurrentSnapshot)
  - [getKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getKey)
  - [getRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getRequestArgs)
  - [getResumeFilePath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_getResumeFilePath)
  - [isResumeFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_isResumeFile)
  - [toFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_toFile)
  - [toJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_toJson)
  - [updateCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#method_updateCurrentSnapshot)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractResumableTransfer.html#top)
