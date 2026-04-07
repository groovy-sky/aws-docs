Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## DirectoryTransferProgressAggregator     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)

FinalYes

Aggregates per-object progress snapshots into a directory-level snapshot.

Acts as an object-level listener and emits directory-level events through an
internal notifier to the provided directory listeners.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#toc-constants)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method___construct)
: mixed [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_bytesTransferred)
: bool [getSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_getSnapshot)
: [DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)[incrementTotals()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_incrementTotals)
: void Update totals, useful when object list is streamed.[notifyDirectoryComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_notifyDirectoryComplete)
: void Notify directory listeners that the directory transfer completed.[notifyDirectoryFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_notifyDirectoryFail)
: void Notify directory listeners that the directory transfer failed.[notifyDirectoryInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_notifyDirectoryInitiated)
: void Notify directory listeners that the directory transfer has been initiated.[priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_transferComplete)
: void [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_transferFail)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)
: void

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#constants)

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method___construct)

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

#### bytesTransferred()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_bytesTransferred)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_bytesTransferred\#tags)

inheritDoc

##### Return values

bool
—

true to notify successful handling otherwise false.

#### getSnapshot()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_getSnapshot)

`
    public
                    getSnapshot() : DirectoryTransferProgressSnapshot`

##### Return values

[DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)

#### incrementTotals()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_incrementTotals)

Update totals, useful when object list is streamed.

`
    public
                    incrementTotals(int $bytes[, int $files = 1 ]) : void`

##### Parameters

$bytes
: int$files
: int
= 1

#### notifyDirectoryComplete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_notifyDirectoryComplete)

Notify directory listeners that the directory transfer completed.

`
    public
                    notifyDirectoryComplete([array<string|int, mixed>|null $response = null ]) : void`

##### Parameters

$response
: array<string\|int, mixed>\|null
= null

#### notifyDirectoryFail()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_notifyDirectoryFail)

Notify directory listeners that the directory transfer failed.

`
    public
                    notifyDirectoryFail(Throwable|string $reason) : void`

##### Parameters

$reason
: Throwable\|string

#### notifyDirectoryInitiated()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_notifyDirectoryInitiated)

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

#### transferComplete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_transferComplete)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_transferComplete\#tags)

inheritDoc

#### transferFail()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_transferFail)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html\#method_transferFail\#tags)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#toc-methods)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method___construct)
  - [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_bytesTransferred)
  - [getSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_getSnapshot)
  - [incrementTotals()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_incrementTotals)
  - [notifyDirectoryComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_notifyDirectoryComplete)
  - [notifyDirectoryFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_notifyDirectoryFail)
  - [notifyDirectoryInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_notifyDirectoryInitiated)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_transferComplete)
  - [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#method_transferFail)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_transferInitiated)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressAggregator.html#top)
