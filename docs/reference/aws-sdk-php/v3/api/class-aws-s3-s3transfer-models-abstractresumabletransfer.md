Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## AbstractResumableTransfer        in package    - [Aws](package-aws.md)

AbstractYes

### Table of Contents  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-toc.md)

#### Methods  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-construct.md)
: mixed [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-deleteresumefile.md)
: void [fromFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-fromfile.md)
: self Load a resumable state from a file.[fromJson()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-fromjson.md)
: self Deserialize a resumable state from JSON format.[getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getbucket.md)
: string [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getconfig.md)
: array<string\|int, mixed> [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getcurrentsnapshot.md)
: array<string\|int, mixed> [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getkey.md)
: string [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getrequestargs.md)
: array<string\|int, mixed> [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getresumefilepath.md)
: string [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-isresumefile.md)
: bool Check if a file path is a valid resume file.[toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-tofile.md)
: void Save the resumable state to a file.[toJson()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-tojson.md)
: string Serialize the resumable state to JSON format.[updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-updatecurrentsnapshot.md)
: void Update the current snapshot.

### Methods  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-construct.md)

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

#### deleteResumeFile()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-deleteresumefile.md)

`
    public
                    deleteResumeFile([string|null $filePath = null ]) : void`

##### Parameters

$filePath
: string\|null
= null

#### fromFile()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-fromfile.md)

Load a resumable state from a file.

`
    public
    abstract        static        fromFile(string $filePath) : self`

##### Parameters

$filePath
: string

Path to the resume file

##### Tags  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-fromfile-tags.md)

throws[S3TransferException](class-aws-s3-s3transfer-exception-s3transferexception.md)

If the file cannot be read or is invalid

##### Return values

self

#### fromJson()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-fromjson.md)

Deserialize a resumable state from JSON format.

`
    public
    abstract        static        fromJson(string $json) : self`

##### Parameters

$json
: string

JSON-encoded state

##### Tags  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-fromjson-tags.md)

throws[S3TransferException](class-aws-s3-s3transfer-exception-s3transferexception.md)

If the JSON is invalid or missing required fields

##### Return values

self

#### getBucket()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getbucket.md)

`
    public
                    getBucket() : string`

##### Return values

string

#### getConfig()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getconfig.md)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCurrentSnapshot()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getcurrentsnapshot.md)

`
    public
                    getCurrentSnapshot() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getKey()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getkey.md)

`
    public
                    getKey() : string`

##### Return values

string

#### getRequestArgs()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getrequestargs.md)

`
    public
                    getRequestArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getResumeFilePath()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getresumefilepath.md)

`
    public
                    getResumeFilePath() : string`

##### Return values

string

#### isResumeFile()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-isresumefile.md)

Check if a file path is a valid resume file.

`
    public
            static        isResumeFile(string $filePath) : bool`

##### Parameters

$filePath
: string

##### Return values

bool

#### toFile()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-tofile.md)

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

#### toJson()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-tojson.md)

Serialize the resumable state to JSON format.

`
    public
    abstract                toJson() : string`

##### Return values

string
—

JSON-encoded state

#### updateCurrentSnapshot()  [header link](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-updatecurrentsnapshot.md)

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
  - [Methods](class-aws-s3-s3transfer-models-abstractresumabletransfer-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-construct.md)
  - [deleteResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-deleteresumefile.md)
  - [fromFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-fromfile.md)
  - [fromJson()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-fromjson.md)
  - [getBucket()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getbucket.md)
  - [getConfig()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getconfig.md)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getcurrentsnapshot.md)
  - [getKey()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getkey.md)
  - [getRequestArgs()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getrequestargs.md)
  - [getResumeFilePath()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-getresumefilepath.md)
  - [isResumeFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-isresumefile.md)
  - [toFile()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-tofile.md)
  - [toJson()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-tojson.md)
  - [updateCurrentSnapshot()](class-aws-s3-s3transfer-models-abstractresumabletransfer-method-updatecurrentsnapshot.md)

[Back To Top](class-aws-s3-s3transfer-models-abstractresumabletransfer-top.md)
