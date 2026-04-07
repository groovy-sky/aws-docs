Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Utils](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.utils.html)

## FileDownloadHandler     extends [AbstractDownloadHandler](class-aws-s3-s3transfer-utils-abstractdownloadhandler.md)   in package    - [Aws](package-aws.md)       implements  [ResumableDownloadHandlerInterface](class-aws-s3-s3transfer-utils-resumabledownloadhandlerinterface.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#toc-interfaces)

[ResumableDownloadHandlerInterface](class-aws-s3-s3transfer-utils-resumabledownloadhandlerinterface.md)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#toc-constants)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method___construct)
: mixed [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_bytesTransferred)
: bool [getDestination()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getDestination)
: string [getFixedPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getFixedPartSize)
: int [getHandlerResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getHandlerResult)
: string Returns the handler result.[getResumeFilePath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getResumeFilePath)
: string [getTemporaryFilePath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getTemporaryFilePath)
: string [initializeDestination()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_initializeDestination)
: void [isConcurrencySupported()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_isConcurrencySupported)
: bool To control whether the download handler supports
concurrency.[isFailsWhenDestinationExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_isFailsWhenDestinationExists)
: bool [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_transferComplete)
: void [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_transferFail)
: void [transferInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_transferInitiated)
: void

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#constants)

#### PROGRESS\_SNAPSHOT\_KEY  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#constant_PROGRESS_SNAPSHOT_KEY)

`
    public
        mixed
    PROGRESS_SNAPSHOT_KEY
    = 'progress_snapshot'
`

#### REASON\_KEY  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#constant_REASON_KEY)

`
    public
        mixed
    REASON_KEY
    = 'reason'
`

#### REQUEST\_ARGS\_KEY  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#constant_REQUEST_ARGS_KEY)

`
    public
        mixed
    REQUEST_ARGS_KEY
    = 'request_args'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method___construct)

`
    public
                    __construct(string $destination, bool $failsWhenDestinationExists[, bool $resumeEnabled = false ][, string|null $temporaryFilePath = null ][, int|null $fixedPartSize = null ]) : mixed`

##### Parameters

$destination
: string$failsWhenDestinationExists
: bool$resumeEnabled
: bool
= false$temporaryFilePath
: string\|null
= null$fixedPartSize
: int\|null
= null

#### bytesTransferred()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_bytesTransferred)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

##### Return values

bool

#### getDestination()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_getDestination)

`
    public
                    getDestination() : string`

##### Return values

string

#### getFixedPartSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_getFixedPartSize)

`
    public
                    getFixedPartSize() : int`

##### Return values

int

#### getHandlerResult()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_getHandlerResult)

Returns the handler result.

`
    public
                    getHandlerResult() : string`

##### Return values

string

#### getResumeFilePath()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_getResumeFilePath)

`
    public
                    getResumeFilePath() : string`

##### Return values

string

#### getTemporaryFilePath()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_getTemporaryFilePath)

`
    public
                    getTemporaryFilePath() : string`

##### Return values

string

#### initializeDestination()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_initializeDestination)

`
    public
                    initializeDestination(array<string|int, mixed> $response) : void`

##### Parameters

$response
: array<string\|int, mixed>

#### isConcurrencySupported()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_isConcurrencySupported)

To control whether the download handler supports
concurrency.

`
    public
                    isConcurrencySupported() : bool`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_isConcurrencySupported\#tags)

inheritDoc

##### Return values

bool

#### isFailsWhenDestinationExists()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_isFailsWhenDestinationExists)

`
    public
                    isFailsWhenDestinationExists() : bool`

##### Return values

bool

#### priority()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_priority)

To provide an order on which listener is notified first.

`
    public
                    priority() : int`

By default, it will provide a neutral value.

##### Return values

int

#### transferComplete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_transferComplete)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

#### transferFail()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_transferFail)

`
    public
                    transferFail(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

#### transferInitiated()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html\#method_transferInitiated)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#toc-methods)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method___construct)
  - [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_bytesTransferred)
  - [getDestination()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getDestination)
  - [getFixedPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getFixedPartSize)
  - [getHandlerResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getHandlerResult)
  - [getResumeFilePath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getResumeFilePath)
  - [getTemporaryFilePath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_getTemporaryFilePath)
  - [initializeDestination()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_initializeDestination)
  - [isConcurrencySupported()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_isConcurrencySupported)
  - [isFailsWhenDestinationExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_isFailsWhenDestinationExists)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_transferComplete)
  - [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_transferFail)
  - [transferInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#method_transferInitiated)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.FileDownloadHandler.html#top)
