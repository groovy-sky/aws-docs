Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Utils](namespace-aws-s3-s3transfer-utils.md)

## AbstractDownloadHandler     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)

AbstractYes

### Table of Contents  [header link](class-aws-s3-s3transfer-utils-abstractdownloadhandler-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-utils-abstractdownloadhandler-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-utils-abstractdownloadhandler-toc-methods.md)

[bytesTransferred()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_bytesTransferred)
: bool [getHandlerResult()](class-aws-s3-s3transfer-utils-abstractdownloadhandler-method-gethandlerresult.md)
: mixed Returns the handler result.[isConcurrencySupported()](class-aws-s3-s3transfer-utils-abstractdownloadhandler-method-isconcurrencysupported.md)
: bool To control whether the download handler supports
concurrency.[priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[transferComplete()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferComplete)
: void [transferFail()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferFail)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)
: void

### Constants  [header link](class-aws-s3-s3transfer-utils-abstractdownloadhandler-constants.md)

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

### Methods  [header link](class-aws-s3-s3transfer-utils-abstractdownloadhandler-methods.md)

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_bytesTransferred)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Return values

bool
—

true to notify successful handling otherwise false.

#### getHandlerResult()  [header link](class-aws-s3-s3transfer-utils-abstractdownloadhandler-method-gethandlerresult.md)

Returns the handler result.

`
    public
    abstract                getHandlerResult() : mixed`

- For FileDownloadHandler it may return the file destination.
- For StreamDownloadHandler it may return an instance of StreamInterface
containing the content of the object.

#### isConcurrencySupported()  [header link](class-aws-s3-s3transfer-utils-abstractdownloadhandler-method-isconcurrencysupported.md)

To control whether the download handler supports
concurrency.

`
    public
    abstract                isConcurrencySupported() : bool`

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

#### transferComplete()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_transferComplete)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

#### transferFail()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_transferFail)

`
    public
                    transferFail(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.
- reason: (Throwable) The exception originated by the transfer failure.

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
  - [Constants](class-aws-s3-s3transfer-utils-abstractdownloadhandler-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-utils-abstractdownloadhandler-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [bytesTransferred()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_bytesTransferred)
  - [getHandlerResult()](class-aws-s3-s3transfer-utils-abstractdownloadhandler-method-gethandlerresult.md)
  - [isConcurrencySupported()](class-aws-s3-s3transfer-utils-abstractdownloadhandler-method-isconcurrencysupported.md)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [transferComplete()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferComplete)
  - [transferFail()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferFail)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)

[Back To Top](class-aws-s3-s3transfer-utils-abstractdownloadhandler-top.md)
