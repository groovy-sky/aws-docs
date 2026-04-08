Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## AbstractTransferListener        in package    - [Aws](package-aws.md)

AbstractYes

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-progress-snapshot-key.md)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-reason-key.md)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-request-args-key.md)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-toc-methods.md)

[bytesTransferred()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-bytestransferred.md)
: bool [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-priority.md)
: int To provide an order on which listener is notified first.[transferComplete()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transfercomplete.md)
: void [transferFail()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transferfail.md)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transferinitiated.md)
: void

### Constants  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-constants.md)

#### PROGRESS\_SNAPSHOT\_KEY  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-progress-snapshot-key.md)

`
    public
        mixed
    PROGRESS_SNAPSHOT_KEY
    = 'progress_snapshot'
`

#### REASON\_KEY  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-reason-key.md)

`
    public
        mixed
    REASON_KEY
    = 'reason'
`

#### REQUEST\_ARGS\_KEY  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-request-args-key.md)

`
    public
        mixed
    REQUEST_ARGS_KEY
    = 'request_args'
`

### Methods  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-methods.md)

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-bytestransferred.md)

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

#### priority()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-priority.md)

To provide an order on which listener is notified first.

`
    public
                    priority() : int`

By default, it will provide a neutral value.

##### Return values

int

#### transferComplete()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transfercomplete.md)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

#### transferFail()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transferfail.md)

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

#### transferInitiated()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transferinitiated.md)

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
  - [Constants](class-aws-s3-s3transfer-progress-abstracttransferlistener-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-abstracttransferlistener-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-progress-snapshot-key.md)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-reason-key.md)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener-constant-request-args-key.md)
- Methods
  - [bytesTransferred()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-bytestransferred.md)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-priority.md)
  - [transferComplete()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transfercomplete.md)
  - [transferFail()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transferfail.md)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener-method-transferinitiated.md)

[Back To Top](class-aws-s3-s3transfer-progress-abstracttransferlistener-top.md)
