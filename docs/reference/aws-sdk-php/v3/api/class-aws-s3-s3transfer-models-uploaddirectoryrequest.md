Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## UploadDirectoryRequest     extends [AbstractTransferRequest](class-aws-s3-s3transfer-models-abstracttransferrequest.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#toc-constants)

[DEFAULT\_MAX\_CONCURRENCY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#constant_DEFAULT_MAX_CONCURRENCY)
= 100

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#toc-properties)

[$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#property_configKeys)
: array<string\|int, mixed>

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method___construct)
: mixed [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
: array<string\|int, mixed> [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
: array<string\|int, mixed> Get current listeners.[getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null Get the progress tracker.[getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
: array<string\|int, mixed> Get listeners that should receive single-object events.[getSourceDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method_getSourceDirectory)
: string [getTargetBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method_getTargetBucket)
: string [getUploadRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method_getUploadRequestArgs)
: array<string\|int, mixed> [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
: void [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
: void For validating config. By default, it provides an empty
implementation.[validateSourceDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method_validateSourceDirectory)
: void Helper method to validate source directory

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#constants)

#### DEFAULT\_MAX\_CONCURRENCY  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#constant_DEFAULT_MAX_CONCURRENCY)

`
    public
        mixed
    DEFAULT_MAX_CONCURRENCY
    = 100
`

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#properties)

#### $configKeys  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#property_configKeys)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['follow_symbolic_links' => 'bool', 'recursive' => 'bool', 's3_prefix' => 'string', 'filter' => 'callable', 's3_delimiter' => 'string', 'upload_object_request_modifier' => 'callable', 'failure_policy' => 'callable', 'max_concurrency' => 'int', 'max_depth' => 'int', 'track_progress' => 'bool']`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#method___construct)

`
    public
                    __construct(string $sourceDirectory, string $targetBucket[, array<string|int, mixed> $uploadRequestArgs = [] ][, array<string|int, mixed> $config = [] ][, array<string|int, mixed> $listeners = [] ][, AbstractTransferListener|null $progressTracker = null ][, array<string|int, mixed> $singleObjectListeners = [] ]) : mixed`

##### Parameters

$sourceDirectory
: string

The source directory to upload.

$targetBucket
: string

The name of the bucket to upload objects to.

$uploadRequestArgs
: array<string\|int, mixed>
= \[\]

The extract arguments to be passed in
each upload request.

$config
: array<string\|int, mixed>
= \[\]

- follow\_symbolic\_links: (boolean, optional) Whether to follow symbolic links when
traversing the file tree.
- recursive: (boolean, optional) Whether to upload directories recursively.
- s3\_prefix: (string, optional) The S3 key prefix to use for each object.
If not provided, files will be uploaded to the root of the bucket.
- filter: (callable, optional) A callback to allow users to filter out unwanted files.
It is invoked for each file. An example implementation is a predicate
that takes a file and returns a boolean indicating whether this file
should be uploaded.
- s3\_delimiter: The S3 delimiter. A delimiter causes a list operation
to roll up all the keys that share a common prefix into a single summary list result.
- upload\_object\_request\_modifier: (callable, optional) A callback mechanism
to allow customers to update individual putObjectRequest that the S3 Transfer Manager generates.
- failure\_policy: (callable, optional) The failure policy to handle failed requests.
- max\_concurrency: (int, optional) The max number of concurrent uploads.
- max\_depth: (int, optional) To indicate the maximum depth of the recursive
file tree walk. By default, it will use the built-in default value which is -1.

$listeners
: array<string\|int, mixed>
= \[\]

Directory-level listeners that receive directory snapshots.

$progressTracker
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null
= null

Directory-level progress tracker.

$singleObjectListeners
: array<string\|int, mixed>
= \[\]

Per-object listeners that receive single-object snapshots.

#### getConfig()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getConfig)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getListeners)

Get current listeners.

`
    public
                    getListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getProgressTracker()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getProgressTracker)

Get the progress tracker.

`
    public
                    getProgressTracker() : AbstractTransferListener|null`

##### Return values

[AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null

#### getSingleObjectListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getSingleObjectListeners)

Get listeners that should receive single-object events.

`
    public
                    getSingleObjectListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getSourceDirectory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#method_getSourceDirectory)

`
    public
                    getSourceDirectory() : string`

##### Return values

string

#### getTargetBucket()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#method_getTargetBucket)

`
    public
                    getTargetBucket() : string`

##### Return values

string

#### getUploadRequestArgs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#method_getUploadRequestArgs)

`
    public
                    getUploadRequestArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### updateConfigWithDefaults()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_updateConfigWithDefaults)

`
    public
                    updateConfigWithDefaults(array<string|int, mixed> $defaultConfig) : void`

##### Parameters

$defaultConfig
: array<string\|int, mixed>

#### validateConfig()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_validateConfig)

For validating config. By default, it provides an empty
implementation.

`
    public
                    validateConfig() : void`

#### validateSourceDirectory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html\#method_validateSourceDirectory)

Helper method to validate source directory

`
    public
                    validateSourceDirectory() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#toc-constants)
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#toc-methods)
- Constants
  - [DEFAULT\_MAX\_CONCURRENCY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#constant_DEFAULT_MAX_CONCURRENCY)
- Properties
  - [$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#property_configKeys)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method___construct)
  - [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
  - [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
  - [getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
  - [getSourceDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method_getSourceDirectory)
  - [getTargetBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method_getTargetBucket)
  - [getUploadRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method_getUploadRequestArgs)
  - [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
  - [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
  - [validateSourceDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#method_validateSourceDirectory)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadDirectoryRequest.html#top)
