Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Utils](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.utils.html)

## StreamDownloadHandler     extends [AbstractDownloadHandler](class-aws-s3-s3transfer-utils-abstractdownloadhandler.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#toc-constants)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method___construct)
: mixed [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_bytesTransferred)
: bool [getHandlerResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_getHandlerResult)
: [StreamInterface](class-psr-http-message-streaminterface.md)Returns the handler result.[isConcurrencySupported()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_isConcurrencySupported)
: bool To control whether the download handler supports
concurrency.[priority()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_priority)
: int [transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_transferComplete)
: void [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_transferFail)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)
: void

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#constants)

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method___construct)

`
    public
                    __construct([StreamInterface|null $stream = null ]) : mixed`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md) \|null
= null

#### bytesTransferred()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_bytesTransferred)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_bytesTransferred\#tags)

inheritDoc

##### Return values

bool

#### getHandlerResult()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_getHandlerResult)

Returns the handler result.

`
    public
                    getHandlerResult() : StreamInterface`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_getHandlerResult\#tags)

inheritDoc

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### isConcurrencySupported()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_isConcurrencySupported)

To control whether the download handler supports
concurrency.

`
    public
                    isConcurrencySupported() : bool`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_isConcurrencySupported\#tags)

inheritDoc

##### Return values

bool

#### priority()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_priority)

`
    public
                    priority() : int`

##### Return values

int

#### transferComplete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_transferComplete)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

#### transferFail()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html\#method_transferFail)

`
    public
                    transferFail(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

#### transferInitiated()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_transferInitiated)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the request initialization.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#toc-methods)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method___construct)
  - [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_bytesTransferred)
  - [getHandlerResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_getHandlerResult)
  - [isConcurrencySupported()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_isConcurrencySupported)
  - [priority()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_priority)
  - [transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_transferComplete)
  - [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#method_transferFail)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Utils.StreamDownloadHandler.html#top)
