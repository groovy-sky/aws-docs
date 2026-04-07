Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## DownloadDirectoryRequest     extends [AbstractTransferRequest](class-aws-s3-s3transfer-models-abstracttransferrequest.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#toc-constants)

[DEFAULT\_MAX\_CONCURRENCY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#constant_DEFAULT_MAX_CONCURRENCY)
= 100

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#toc-properties)

[$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#property_configKeys)
: array<string\|int, mixed>

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method___construct)
: mixed [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
: array<string\|int, mixed> [getDestinationDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method_getDestinationDirectory)
: string [getDownloadRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method_getDownloadRequestArgs)
: array<string\|int, mixed> [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
: array<string\|int, mixed> Get current listeners.[getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null Get the progress tracker.[getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
: array<string\|int, mixed> Get listeners that should receive single-object events.[getSourceBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method_getSourceBucket)
: string [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
: void [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
: void For validating config. By default, it provides an empty
implementation.[validateDestinationDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method_validateDestinationDirectory)
: void Helper method to validate the destination directory exists.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#constants)

#### DEFAULT\_MAX\_CONCURRENCY  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#constant_DEFAULT_MAX_CONCURRENCY)

`
    public
        mixed
    DEFAULT_MAX_CONCURRENCY
    = 100
`

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#properties)

#### $configKeys  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#property_configKeys)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['s3_prefix' => 'string', 'filter' => 'callable', 'download_object_request_modifier' => 'callable', 'failure_policy' => 'callable', 'max_concurrency' => 'int', 'track_progress' => 'bool', 'target_part_size_bytes' => 'int', 'list_objects_v2_args' => 'array', 'fails_when_destination_exists' => 'bool']`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#method___construct)

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
: array<string\|int, [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) >
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

#### getDestinationDirectory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#method_getDestinationDirectory)

`
    public
                    getDestinationDirectory() : string`

##### Return values

string

#### getDownloadRequestArgs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#method_getDownloadRequestArgs)

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

[AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null

#### getSingleObjectListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getSingleObjectListeners)

Get listeners that should receive single-object events.

`
    public
                    getSingleObjectListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getSourceBucket()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#method_getSourceBucket)

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

#### validateDestinationDirectory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html\#method_validateDestinationDirectory)

Helper method to validate the destination directory exists.

`
    public
                    validateDestinationDirectory() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#toc-constants)
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#toc-methods)
- Constants
  - [DEFAULT\_MAX\_CONCURRENCY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#constant_DEFAULT_MAX_CONCURRENCY)
- Properties
  - [$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#property_configKeys)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method___construct)
  - [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
  - [getDestinationDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method_getDestinationDirectory)
  - [getDownloadRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method_getDownloadRequestArgs)
  - [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
  - [getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
  - [getSourceBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method_getSourceBucket)
  - [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
  - [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
  - [validateDestinationDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#method_validateDestinationDirectory)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadDirectoryRequest.html#top)
