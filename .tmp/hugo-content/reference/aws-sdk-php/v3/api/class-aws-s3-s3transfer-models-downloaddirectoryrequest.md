Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## DownloadDirectoryRequest     extends [AbstractTransferRequest](class-aws-s3-s3transfer-models-abstracttransferrequest.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-toc-constants.md)

[DEFAULT\_MAX\_CONCURRENCY](class-aws-s3-s3transfer-models-downloaddirectoryrequest-constant-default-max-concurrency.md)
= 100

#### Properties  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-toc-properties.md)

[$configKeys](class-aws-s3-s3transfer-models-downloaddirectoryrequest-property-configkeys.md)
: array<string\|int, mixed>

#### Methods  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-construct.md)
: mixed [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
: array<string\|int, mixed> [getDestinationDirectory()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getdestinationdirectory.md)
: string [getDownloadRequestArgs()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getdownloadrequestargs.md)
: array<string\|int, mixed> [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
: array<string\|int, mixed> Get current listeners.[getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null Get the progress tracker.[getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
: array<string\|int, mixed> Get listeners that should receive single-object events.[getSourceBucket()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getsourcebucket.md)
: string [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
: void [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
: void For validating config. By default, it provides an empty
implementation.[validateDestinationDirectory()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-validatedestinationdirectory.md)
: void Helper method to validate the destination directory exists.

### Constants  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-constants.md)

#### DEFAULT\_MAX\_CONCURRENCY  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-constant-default-max-concurrency.md)

`
    public
        mixed
    DEFAULT_MAX_CONCURRENCY
    = 100
`

### Properties  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-properties.md)

#### $configKeys  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-property-configkeys.md)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['s3_prefix' => 'string', 'filter' => 'callable', 'download_object_request_modifier' => 'callable', 'failure_policy' => 'callable', 'max_concurrency' => 'int', 'track_progress' => 'bool', 'target_part_size_bytes' => 'int', 'list_objects_v2_args' => 'array', 'fails_when_destination_exists' => 'bool']`

### Methods  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-construct.md)

`
    public
                    __construct(string $sourceBucket, string $destinationDirectory[, array<string|int, mixed> $downloadRequestArgs = [] ][, array<string|int, mixed> $config = [] ][, array<string|int, AbstractTransferListener> $listeners = [] ][, AbstractTransferListener|null $progressTracker = null ][, array<string|int, mixed> $singleObjectListeners = [] ]) : mixed`

##### Parameters

$sourceBucket
: string

The bucket from where the files are going to be
downloaded from.

$destinationDirectory
: string

The destination path where the downloaded
files will be placed in.

$downloadRequestArgs
: array<string\|int, mixed>
= \[\]$config
: array<string\|int, mixed>
= \[\]

The config options for this download directory operation.

- s3\_prefix: (string, optional) This parameter will be considered just if
not provided as part of the list\_objects\_v2\_args config option.
- filter: (Closure, optional) A callable which will receive an object key as
parameter and should return true or false in order to determine
whether the object should be downloaded.
- download\_object\_request\_modifier: (Closure, optional) A function that will
be invoked right before the download request is performed and that will
receive as parameter the request arguments for each request.
- failure\_policy: (Closure, optional) A function that will be invoked
on a download failure and that will receive as parameters:
  - $requestArgs: (array) The arguments for the request that originated
    the failure.
  - $downloadDirectoryRequestArgs: (array) The arguments for the download
    directory request.
  - $reason: (Throwable) The exception that originated the request failure.
  - $downloadDirectoryResponse: (DownloadDirectoryResult) The download response
    to that point in the upload process.
- track\_progress: (bool, optional) Overrides the config option set
in the transfer manager instantiation to decide whether transfer
progress should be tracked.
- target\_part\_size\_bytes: (int, optional) The part size in bytes
to be used in a range multipart download.
- fails\_when\_destination\_exists: (bool) Whether to fail when a destination
file exists.
- max\_concurrency: (int, optional) The max number of concurrent downloads.
- list\_objects\_v2\_args: (array, optional) The arguments to be included
as part of the listObjectV2 request in order to fetch the objects to
be downloaded. The most common arguments would be:
  - MaxKeys: (int) Sets the maximum number of keys returned in the response.
  - Prefix: (string) To limit the response to keys that begin with the
    specified prefix.

$listeners
: array<string\|int, [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) >
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

#### getDestinationDirectory()  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getdestinationdirectory.md)

`
    public
                    getDestinationDirectory() : string`

##### Return values

string

#### getDownloadRequestArgs()  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getdownloadrequestargs.md)

`
    public
                    getDownloadRequestArgs() : array<string|int, mixed>`

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

#### getSourceBucket()  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getsourcebucket.md)

`
    public
                    getSourceBucket() : string`

##### Return values

string

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

#### validateDestinationDirectory()  [header link](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-validatedestinationdirectory.md)

Helper method to validate the destination directory exists.

`
    public
                    validateDestinationDirectory() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3transfer-models-downloaddirectoryrequest-toc-constants.md)
  - [Properties](class-aws-s3-s3transfer-models-downloaddirectoryrequest-toc-properties.md)
  - [Methods](class-aws-s3-s3transfer-models-downloaddirectoryrequest-toc-methods.md)
- Constants
  - [DEFAULT\_MAX\_CONCURRENCY](class-aws-s3-s3transfer-models-downloaddirectoryrequest-constant-default-max-concurrency.md)
- Properties
  - [$configKeys](class-aws-s3-s3transfer-models-downloaddirectoryrequest-property-configkeys.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-construct.md)
  - [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
  - [getDestinationDirectory()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getdestinationdirectory.md)
  - [getDownloadRequestArgs()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getdownloadrequestargs.md)
  - [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
  - [getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
  - [getSourceBucket()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-getsourcebucket.md)
  - [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
  - [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
  - [validateDestinationDirectory()](class-aws-s3-s3transfer-models-downloaddirectoryrequest-method-validatedestinationdirectory.md)

[Back To Top](class-aws-s3-s3transfer-models-downloaddirectoryrequest-top.md)
