Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Utils](namespace-aws-s3-s3transfer-utils.md)

## StreamDownloadHandler     extends [AbstractDownloadHandler](class-aws-s3-s3transfer-utils-abstractdownloadhandler.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-construct.md)
: mixed [bytesTransferred()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-bytestransferred.md)
: bool [getHandlerResult()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-gethandlerresult.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)Returns the handler result.[isConcurrencySupported()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-isconcurrencysupported.md)
: bool To control whether the download handler supports
concurrency.[priority()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-priority.md)
: int [transferComplete()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-transfercomplete.md)
: void [transferFail()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-transferfail.md)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)
: void

### Constants  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-constants.md)

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

### Methods  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-construct.md)

`
    public
                    __construct([StreamInterface|null $stream = null ]) : mixed`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md) \|null
= null

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-bytestransferred.md)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

##### Tags  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-bytestransferred-tags.md)

inheritDoc

##### Return values

bool

#### getHandlerResult()  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-gethandlerresult.md)

Returns the handler result.

`
    public
                    getHandlerResult() : StreamInterface`

##### Tags  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-gethandlerresult-tags.md)

inheritDoc

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### isConcurrencySupported()  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-isconcurrencysupported.md)

To control whether the download handler supports
concurrency.

`
    public
                    isConcurrencySupported() : bool`

##### Tags  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-isconcurrencysupported-tags.md)

inheritDoc

##### Return values

bool

#### priority()  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-priority.md)

`
    public
                    priority() : int`

##### Return values

int

#### transferComplete()  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-transfercomplete.md)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

#### transferFail()  [header link](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-transferfail.md)

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
  - [Constants](class-aws-s3-s3transfer-utils-streamdownloadhandler-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-utils-streamdownloadhandler-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-construct.md)
  - [bytesTransferred()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-bytestransferred.md)
  - [getHandlerResult()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-gethandlerresult.md)
  - [isConcurrencySupported()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-isconcurrencysupported.md)
  - [priority()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-priority.md)
  - [transferComplete()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-transfercomplete.md)
  - [transferFail()](class-aws-s3-s3transfer-utils-streamdownloadhandler-method-transferfail.md)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)

[Back To Top](class-aws-s3-s3transfer-utils-streamdownloadhandler-top.md)
