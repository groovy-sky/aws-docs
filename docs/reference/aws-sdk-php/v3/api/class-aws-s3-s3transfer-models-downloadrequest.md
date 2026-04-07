Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## DownloadRequest     extends [AbstractTransferRequest](class-aws-s3-s3transfer-models-abstracttransferrequest.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#toc)

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#toc-properties)

[$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#property_configKeys)
: array<string\|int, mixed>

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method___construct)
: mixed [fromDownloadRequestAndDownloadHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_fromDownloadRequestAndDownloadHandler)
: self [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
: array<string\|int, mixed> [getDownloadHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_getDownloadHandler)
: [AbstractDownloadHandler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.AbstractDownloadHandler.html)[getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
: array<string\|int, mixed> Get current listeners.[getObjectRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_getObjectRequestArgs)
: array<string\|int, mixed> [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null Get the progress tracker.[getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
: array<string\|int, mixed> Get listeners that should receive single-object events.[getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_getSource)
: array<string\|int, mixed>\|string\|null [normalizeSourceAsArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_normalizeSourceAsArray)
: array<string\|int, mixed> Helper method to normalize the source as an array with:
\- Bucket
\- Key[updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
: void [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
: void For validating config. By default, it provides an empty
implementation.

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#properties)

#### $configKeys  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#property_configKeys)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['response_checksum_validation' => 'string', 'multipart_download_type' => 'string', 'track_progress' => 'bool', 'concurrency' => 'int', 'resume_enabled' => 'bool', 'resume_file_path' => 'string', 'target_part_size_bytes' => 'int']`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#method___construct)

`
    public
                    __construct(string|array<string|int, mixed>|null $source[, array<string|int, mixed> $downloadRequestArgs = [] ][, array<string|int, mixed> $config = [] ][, AbstractDownloadHandler|null $downloadHandler = null ][, array<string|int, AbstractTransferListener>|null $listeners = [] ][, AbstractTransferListener|null $progressTracker = null ]) : mixed`

##### Parameters

$source
: string\|array<string\|int, mixed>\|null

The object to be downloaded from S3.
It can be either a string with a S3 URI or an array with a Bucket and Key
properties set.

$downloadRequestArgs
: array<string\|int, mixed>
= \[\]$config
: array<string\|int, mixed>
= \[\]

The configuration to be used for this operation:

- multipart\_download\_type: (string, optional)
Overrides the resolved value from the transfer manager config.
- response\_checksum\_validation: (string, optional) Overrides the resolved
value from transfer manager config for whether checksum validation
should be done. This option will be considered just if ChecksumMode
is not present in the request args.
- track\_progress: (bool) Overrides the config option set in the transfer
manager instantiation to decide whether transfer progress should be
tracked.
- target\_part\_size\_bytes: (int) The part size in bytes to be used
in a range multipart download. If this parameter is not provided
then it fallbacks to the transfer manager `target_part_size_bytes`
config value.
- resume\_enabled: (bool): To enable resuming a multipart download when a
failure occurs.
- resume\_file\_path (string, optional): To override the default resume file
location to be generated. If specified the file name must end in `.resume`
otherwise it will be added automatically.

$downloadHandler
: [AbstractDownloadHandler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.AbstractDownloadHandler.html) \|null
= null$listeners
: array<string\|int, [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) >\|null
= \[\]$progressTracker
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null
= null

#### fromDownloadRequestAndDownloadHandler()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#method_fromDownloadRequestAndDownloadHandler)

`
    public
            static        fromDownloadRequestAndDownloadHandler(DownloadRequest $downloadRequest, FileDownloadHandler $downloadHandler) : self`

##### Parameters

$downloadRequest
: [DownloadRequest](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html)$downloadHandler
: [FileDownloadHandler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html)

##### Return values

self

#### getConfig()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getConfig)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getDownloadHandler()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#method_getDownloadHandler)

`
    public
                    getDownloadHandler() : AbstractDownloadHandler`

##### Return values

[AbstractDownloadHandler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.AbstractDownloadHandler.html)

#### getListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getListeners)

Get current listeners.

`
    public
                    getListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getObjectRequestArgs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#method_getObjectRequestArgs)

`
    public
                    getObjectRequestArgs() : array<string|int, mixed>`

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

#### getSource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#method_getSource)

`
    public
                    getSource() : array<string|int, mixed>|string|null`

##### Return values

array<string\|int, mixed>\|string\|null

#### normalizeSourceAsArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html\#method_normalizeSourceAsArray)

Helper method to normalize the source as an array with:
\- Bucket
\- Key

`
    public
                    normalizeSourceAsArray() : array<string|int, mixed>`

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

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#toc-methods)
- Properties
  - [$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#property_configKeys)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method___construct)
  - [fromDownloadRequestAndDownloadHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_fromDownloadRequestAndDownloadHandler)
  - [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
  - [getDownloadHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_getDownloadHandler)
  - [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
  - [getObjectRequestArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_getObjectRequestArgs)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
  - [getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
  - [getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_getSource)
  - [normalizeSourceAsArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#method_normalizeSourceAsArray)
  - [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
  - [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.DownloadRequest.html#top)
