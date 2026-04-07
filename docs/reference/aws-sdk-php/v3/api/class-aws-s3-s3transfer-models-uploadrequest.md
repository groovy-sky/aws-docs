Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## UploadRequest     extends [AbstractTransferRequest](class-aws-s3-s3transfer-models-abstracttransferrequest.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#toc)

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#toc-properties)

[$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#property_configKeys)
: array<string\|int, mixed>

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method___construct)
: mixed [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
: array<string\|int, mixed> [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
: array<string\|int, mixed> Get current listeners.[getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null Get the progress tracker.[getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
: array<string\|int, mixed> Get listeners that should receive single-object events.[getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method_getSource)
: [StreamInterface](class-psr-http-message-streaminterface.md) \|string Get the source.[getUploadRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method_getUploadRequestArgs)
: array<string\|int, mixed> Get the put object request.[updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
: void [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
: void For validating config. By default, it provides an empty
implementation.[validateRequiredParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method_validateRequiredParameters)
: void Helper method for validating required parameters.[validateSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method_validateSource)
: void Helper method for validating the given source.

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#properties)

#### $configKeys  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#property_configKeys)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['multipart_upload_threshold_bytes' => 'int', 'target_part_size_bytes' => 'int', 'track_progress' => 'bool', 'concurrency' => 'int', 'request_checksum_calculation' => 'string', 'resume_enabled' => 'bool', 'resume_file_path' => 'string']`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#method___construct)

`
    public
                    __construct(string|StreamInterface $source, array<string|int, mixed> $uploadRequestArgs[, array<string|int, mixed> $config = [] ][, array<string|int, AbstractTransferListener>|null $listeners = [] ][, AbstractTransferListener|null $progressTracker = null ]) : mixed`

##### Parameters

$source
: string\| [StreamInterface](class-psr-http-message-streaminterface.md)$uploadRequestArgs
: array<string\|int, mixed>

The putObject request arguments.
Required parameters would be:

- Bucket: (string, required)
- Key: (string, required)

$config
: array<string\|int, mixed>
= \[\]

The config options for this upload operation.

- multipart\_upload\_threshold\_bytes: (int, optional)
To override the default threshold for when to use multipart upload.
- target\_part\_size\_bytes: (int, optional) To override the default
target part size in bytes.
- track\_progress: (bool, optional) To override the default option for
enabling progress tracking. If this option is resolved as true and
a progressTracker parameter is not provided then, a default implementation
will be resolved. This option is intended to make the operation to use
a default progress tracker implementation when $progressTracker is null.
- concurrency: (int, optional) To override default value for concurrency.
- request\_checksum\_calculation: (string, optional, defaulted to `when_supported`)
- resume\_enabled: (bool): To enable resuming a multipart download when a
failure occurs.
- resume\_file\_path (string, optional): To override the default resume file
location to be generated. If specified the file name must end in `.resume`
otherwise it will be added automatically.

$listeners
: array<string\|int, [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) >\|null
= \[\]$progressTracker
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null
= null

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

#### getSource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#method_getSource)

Get the source.

`
    public
                    getSource() : StreamInterface|string`

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md) \|string

#### getUploadRequestArgs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#method_getUploadRequestArgs)

Get the put object request.

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

#### validateRequiredParameters()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#method_validateRequiredParameters)

Helper method for validating required parameters.

`
    public
                    validateRequiredParameters([string|null $customMessage = null ]) : void`

##### Parameters

$customMessage
: string\|null
= null

#### validateSource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html\#method_validateSource)

Helper method for validating the given source.

`
    public
                    validateSource() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#toc-methods)
- Properties
  - [$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#property_configKeys)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method___construct)
  - [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
  - [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
  - [getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
  - [getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method_getSource)
  - [getUploadRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method_getUploadRequestArgs)
  - [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
  - [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
  - [validateRequiredParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method_validateRequiredParameters)
  - [validateSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#method_validateSource)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadRequest.html#top)
