Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## DownloadRequest     extends [AbstractTransferRequest](class-aws-s3-s3transfer-models-abstracttransferrequest.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-models-downloadrequest-toc.md)

#### Properties  [header link](class-aws-s3-s3transfer-models-downloadrequest-toc-properties.md)

[$configKeys](class-aws-s3-s3transfer-models-downloadrequest-property-configkeys.md)
: array<string\|int, mixed>

#### Methods  [header link](class-aws-s3-s3transfer-models-downloadrequest-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-downloadrequest-method-construct.md)
: mixed [fromDownloadRequestAndDownloadHandler()](class-aws-s3-s3transfer-models-downloadrequest-method-fromdownloadrequestanddownloadhandler.md)
: self [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
: array<string\|int, mixed> [getDownloadHandler()](class-aws-s3-s3transfer-models-downloadrequest-method-getdownloadhandler.md)
: [AbstractDownloadHandler](class-aws-s3-s3transfer-utils-abstractdownloadhandler.md)[getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
: array<string\|int, mixed> Get current listeners.[getObjectRequestArgs()](class-aws-s3-s3transfer-models-downloadrequest-method-getobjectrequestargs.md)
: array<string\|int, mixed> [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null Get the progress tracker.[getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
: array<string\|int, mixed> Get listeners that should receive single-object events.[getSource()](class-aws-s3-s3transfer-models-downloadrequest-method-getsource.md)
: array<string\|int, mixed>\|string\|null [normalizeSourceAsArray()](class-aws-s3-s3transfer-models-downloadrequest-method-normalizesourceasarray.md)
: array<string\|int, mixed> Helper method to normalize the source as an array with:
\- Bucket
\- Key[updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
: void [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)
: void For validating config. By default, it provides an empty
implementation.

### Properties  [header link](class-aws-s3-s3transfer-models-downloadrequest-properties.md)

#### $configKeys  [header link](class-aws-s3-s3transfer-models-downloadrequest-property-configkeys.md)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['response_checksum_validation' => 'string', 'multipart_download_type' => 'string', 'track_progress' => 'bool', 'concurrency' => 'int', 'resume_enabled' => 'bool', 'resume_file_path' => 'string', 'target_part_size_bytes' => 'int']`

### Methods  [header link](class-aws-s3-s3transfer-models-downloadrequest-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-downloadrequest-method-construct.md)

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
: [AbstractDownloadHandler](class-aws-s3-s3transfer-utils-abstractdownloadhandler.md) \|null
= null$listeners
: array<string\|int, [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) >\|null
= \[\]$progressTracker
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null
= null

#### fromDownloadRequestAndDownloadHandler()  [header link](class-aws-s3-s3transfer-models-downloadrequest-method-fromdownloadrequestanddownloadhandler.md)

`
    public
            static        fromDownloadRequestAndDownloadHandler(DownloadRequest $downloadRequest, FileDownloadHandler $downloadHandler) : self`

##### Parameters

$downloadRequest
: [DownloadRequest](class-aws-s3-s3transfer-models-downloadrequest.md)$downloadHandler
: [FileDownloadHandler](class-aws-s3-s3transfer-utils-filedownloadhandler.md)

##### Return values

self

#### getConfig()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getConfig)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getDownloadHandler()  [header link](class-aws-s3-s3transfer-models-downloadrequest-method-getdownloadhandler.md)

`
    public
                    getDownloadHandler() : AbstractDownloadHandler`

##### Return values

[AbstractDownloadHandler](class-aws-s3-s3transfer-utils-abstractdownloadhandler.md)

#### getListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getListeners)

Get current listeners.

`
    public
                    getListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getObjectRequestArgs()  [header link](class-aws-s3-s3transfer-models-downloadrequest-method-getobjectrequestargs.md)

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

[AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null

#### getSingleObjectListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest.md\#method_getSingleObjectListeners)

Get listeners that should receive single-object events.

`
    public
                    getSingleObjectListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getSource()  [header link](class-aws-s3-s3transfer-models-downloadrequest-method-getsource.md)

`
    public
                    getSource() : array<string|int, mixed>|string|null`

##### Return values

array<string\|int, mixed>\|string\|null

#### normalizeSourceAsArray()  [header link](class-aws-s3-s3transfer-models-downloadrequest-method-normalizesourceasarray.md)

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
  - [Properties](class-aws-s3-s3transfer-models-downloadrequest-toc-properties.md)
  - [Methods](class-aws-s3-s3transfer-models-downloadrequest-toc-methods.md)
- Properties
  - [$configKeys](class-aws-s3-s3transfer-models-downloadrequest-property-configkeys.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-downloadrequest-method-construct.md)
  - [fromDownloadRequestAndDownloadHandler()](class-aws-s3-s3transfer-models-downloadrequest-method-fromdownloadrequestanddownloadhandler.md)
  - [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getConfig)
  - [getDownloadHandler()](class-aws-s3-s3transfer-models-downloadrequest-method-getdownloadhandler.md)
  - [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getListeners)
  - [getObjectRequestArgs()](class-aws-s3-s3transfer-models-downloadrequest-method-getobjectrequestargs.md)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getProgressTracker)
  - [getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_getSingleObjectListeners)
  - [getSource()](class-aws-s3-s3transfer-models-downloadrequest-method-getsource.md)
  - [normalizeSourceAsArray()](class-aws-s3-s3transfer-models-downloadrequest-method-normalizesourceasarray.md)
  - [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_updateConfigWithDefaults)
  - [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest.md#method_validateConfig)

[Back To Top](class-aws-s3-s3transfer-models-downloadrequest-top.md)
