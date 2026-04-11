Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## DirectoryTransferProgressAggregator     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)

FinalYes

Aggregates per-object progress snapshots into a directory-level snapshot.

Acts as an object-level listener and emits directory-level events through an
internal notifier to the provided directory listeners.

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-construct.md)
: mixed [bytesTransferred()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-bytestransferred.md)
: bool [getSnapshot()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-getsnapshot.md)
: [DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)[incrementTotals()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-incrementtotals.md)
: void Update totals, useful when object list is streamed.[notifyDirectoryComplete()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectorycomplete.md)
: void Notify directory listeners that the directory transfer completed.[notifyDirectoryFail()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectoryfail.md)
: void Notify directory listeners that the directory transfer failed.[notifyDirectoryInitiated()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectoryinitiated.md)
: void Notify directory listeners that the directory transfer has been initiated.[priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[transferComplete()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-transfercomplete.md)
: void [transferFail()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-transferfail.md)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)
: void

### Constants  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-constants.md)

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

### Methods  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-construct.md)

`
    public
                    __construct(string $identifier, int $totalBytes, int $totalFiles[, array<string|int, mixed> $directoryListeners = [] ][, AbstractTransferListener|null $directoryProgressTracker = null ]) : mixed`

##### Parameters

$identifier
: string$totalBytes
: int$totalFiles
: int$directoryListeners
: array<string\|int, mixed>
= \[\]$directoryProgressTracker
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null
= null

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-bytestransferred.md)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-bytestransferred-tags.md)

inheritDoc

##### Return values

bool
—

true to notify successful handling otherwise false.

#### getSnapshot()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-getsnapshot.md)

`
    public
                    getSnapshot() : DirectoryTransferProgressSnapshot`

##### Return values

[DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)

#### incrementTotals()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-incrementtotals.md)

Update totals, useful when object list is streamed.

`
    public
                    incrementTotals(int $bytes[, int $files = 1 ]) : void`

##### Parameters

$bytes
: int$files
: int
= 1

#### notifyDirectoryComplete()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectorycomplete.md)

Notify directory listeners that the directory transfer completed.

`
    public
                    notifyDirectoryComplete([array<string|int, mixed>|null $response = null ]) : void`

##### Parameters

$response
: array<string\|int, mixed>\|null
= null

#### notifyDirectoryFail()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectoryfail.md)

Notify directory listeners that the directory transfer failed.

`
    public
                    notifyDirectoryFail(Throwable|string $reason) : void`

##### Parameters

$reason
: Throwable\|string

#### notifyDirectoryInitiated()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectoryinitiated.md)

Notify directory listeners that the directory transfer has been initiated.

`
    public
                    notifyDirectoryInitiated(array<string|int, mixed> $requestArgs) : void`

##### Parameters

$requestArgs
: array<string\|int, mixed>

#### priority()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_priority)

To provide an order on which listener is notified first.

`
    public
                    priority() : int`

By default, it will provide a neutral value.

##### Return values

int

#### transferComplete()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-transfercomplete.md)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-transfercomplete-tags.md)

inheritDoc

#### transferFail()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-transferfail.md)

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

##### Tags  [header link](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-transferfail-tags.md)

inheritDoc

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
  - [Constants](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-construct.md)
  - [bytesTransferred()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-bytestransferred.md)
  - [getSnapshot()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-getsnapshot.md)
  - [incrementTotals()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-incrementtotals.md)
  - [notifyDirectoryComplete()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectorycomplete.md)
  - [notifyDirectoryFail()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectoryfail.md)
  - [notifyDirectoryInitiated()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-notifydirectoryinitiated.md)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [transferComplete()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-transfercomplete.md)
  - [transferFail()](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-method-transferfail.md)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)

[Back To Top](class-aws-s3-s3transfer-progress-directorytransferprogressaggregator-top.md)
