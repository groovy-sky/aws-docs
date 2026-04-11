Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## UploadDirectoryRequest     extends [AbstractTransferRequest](class-aws-s3-s3transfer-models-abstracttransferrequest.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-toc-constants.md)

[DEFAULT\_MAX\_CONCURRENCY](class-aws-s3-s3transfer-models-uploaddirectoryrequest-constant-default-max-concurrency.md)
= 100

#### Properties  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-toc-properties.md)

[$configKeys](class-aws-s3-s3transfer-models-uploaddirectoryrequest-property-configkeys.md)
: array<string\|int, mixed>

#### Methods  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-construct.md)
: mixed [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
: array<string\|int, mixed> [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
: array<string\|int, mixed> Get current listeners.[getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null Get the progress tracker.[getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
: array<string\|int, mixed> Get listeners that should receive single-object events.[getSourceDirectory()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-getsourcedirectory.md)
: string [getTargetBucket()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-gettargetbucket.md)
: string [getUploadRequestArgs()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-getuploadrequestargs.md)
: array<string\|int, mixed> [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
: void [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
: void For validating config. By default, it provides an empty
implementation.[validateSourceDirectory()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-validatesourcedirectory.md)
: void Helper method to validate source directory

### Constants  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-constants.md)

#### DEFAULT\_MAX\_CONCURRENCY  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-constant-default-max-concurrency.md)

`
    public
        mixed
    DEFAULT_MAX_CONCURRENCY
    = 100
`

### Properties  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-properties.md)

#### $configKeys  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-property-configkeys.md)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['follow_symbolic_links' => 'bool', 'recursive' => 'bool', 's3_prefix' => 'string', 'filter' => 'callable', 's3_delimiter' => 'string', 'upload_object_request_modifier' => 'callable', 'failure_policy' => 'callable', 'max_concurrency' => 'int', 'max_depth' => 'int', 'track_progress' => 'bool']`

### Methods  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-construct.md)

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
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null
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

[AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null

#### getSingleObjectListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getSingleObjectListeners)

Get listeners that should receive single-object events.

`
    public
                    getSingleObjectListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getSourceDirectory()  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-getsourcedirectory.md)

`
    public
                    getSourceDirectory() : string`

##### Return values

string

#### getTargetBucket()  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-gettargetbucket.md)

`
    public
                    getTargetBucket() : string`

##### Return values

string

#### getUploadRequestArgs()  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-getuploadrequestargs.md)

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

#### validateSourceDirectory()  [header link](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-validatesourcedirectory.md)

Helper method to validate source directory

`
    public
                    validateSourceDirectory() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3transfer-models-uploaddirectoryrequest-toc-constants.md)
  - [Properties](class-aws-s3-s3transfer-models-uploaddirectoryrequest-toc-properties.md)
  - [Methods](class-aws-s3-s3transfer-models-uploaddirectoryrequest-toc-methods.md)
- Constants
  - [DEFAULT\_MAX\_CONCURRENCY](class-aws-s3-s3transfer-models-uploaddirectoryrequest-constant-default-max-concurrency.md)
- Properties
  - [$configKeys](class-aws-s3-s3transfer-models-uploaddirectoryrequest-property-configkeys.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-construct.md)
  - [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
  - [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
  - [getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
  - [getSourceDirectory()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-getsourcedirectory.md)
  - [getTargetBucket()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-gettargetbucket.md)
  - [getUploadRequestArgs()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-getuploadrequestargs.md)
  - [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
  - [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
  - [validateSourceDirectory()](class-aws-s3-s3transfer-models-uploaddirectoryrequest-method-validatesourcedirectory.md)

[Back To Top](class-aws-s3-s3transfer-models-uploaddirectoryrequest-top.md)
