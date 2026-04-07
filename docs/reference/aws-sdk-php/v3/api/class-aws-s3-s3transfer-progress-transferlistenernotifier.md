Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## TransferListenerNotifier     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#toc-constants)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method___construct)
: mixed [addListener()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_addListener)
: void [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_bytesTransferred)
: bool [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_transferComplete)
: void [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_transferFail)
: void [transferInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_transferInitiated)
: void

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#constants)

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method___construct)

`
    public
                    __construct([array<string|int, mixed> $listeners = [] ]) : mixed`

##### Parameters

$listeners
: array<string\|int, mixed>
= \[\]

#### addListener()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_addListener)

`
    public
                    addListener(AbstractTransferListener $listener) : void`

##### Parameters

$listener
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)

#### bytesTransferred()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_bytesTransferred)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_bytesTransferred\#tags)

inheritDoc

##### Return values

bool
—

true to notify successful handling otherwise false.

#### priority()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_priority)

To provide an order on which listener is notified first.

`
    public
                    priority() : int`

By default, it will provide a neutral value.

##### Return values

int

#### transferComplete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_transferComplete)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_transferComplete\#tags)

inheritDoc

#### transferFail()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_transferFail)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_transferFail\#tags)

inheritDoc

#### transferInitiated()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_transferInitiated)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the request initialization.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html\#method_transferInitiated\#tags)

inheritDoc
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#toc-methods)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method___construct)
  - [addListener()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_addListener)
  - [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_bytesTransferred)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_transferComplete)
  - [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_transferFail)
  - [transferInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#method_transferInitiated)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferListenerNotifier.html#top)
