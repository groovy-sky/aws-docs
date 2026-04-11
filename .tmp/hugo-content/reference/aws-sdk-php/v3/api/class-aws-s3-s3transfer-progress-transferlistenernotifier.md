Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## TransferListenerNotifier     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-construct.md)
: mixed [addListener()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-addlistener.md)
: void [bytesTransferred()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-bytestransferred.md)
: bool [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[transferComplete()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transfercomplete.md)
: void [transferFail()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transferfail.md)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transferinitiated.md)
: void

### Constants  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-constants.md)

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

### Methods  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-construct.md)

`
    public
                    __construct([array<string|int, mixed> $listeners = [] ]) : mixed`

##### Parameters

$listeners
: array<string\|int, mixed>
= \[\]

#### addListener()  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-addlistener.md)

`
    public
                    addListener(AbstractTransferListener $listener) : void`

##### Parameters

$listener
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-bytestransferred.md)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-bytestransferred-tags.md)

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

#### transferComplete()  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transfercomplete.md)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transfercomplete-tags.md)

inheritDoc

#### transferFail()  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transferfail.md)

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

##### Tags  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transferfail-tags.md)

inheritDoc

#### transferInitiated()  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transferinitiated.md)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the request initialization.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transferinitiated-tags.md)

inheritDoc
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3transfer-progress-transferlistenernotifier-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-transferlistenernotifier-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-construct.md)
  - [addListener()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-addlistener.md)
  - [bytesTransferred()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-bytestransferred.md)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [transferComplete()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transfercomplete.md)
  - [transferFail()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transferfail.md)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-transferlistenernotifier-method-transferinitiated.md)

[Back To Top](class-aws-s3-s3transfer-progress-transferlistenernotifier-top.md)
