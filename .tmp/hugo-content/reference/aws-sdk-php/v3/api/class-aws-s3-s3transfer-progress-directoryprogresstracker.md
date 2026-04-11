Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## DirectoryProgressTracker     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)       implements  [ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

FinalYes

Progress tracker for directory-level transfers using directory snapshots.

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-toc-interfaces.md)

[ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

#### Constants  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-construct.md)
: mixed [bytesTransferred()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-bytestransferred.md)
: bool [getProgressBar()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-getprogressbar.md)
: [ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)[priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[showProgress()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-showprogress.md)
: void To show the progress being tracked.[transferComplete()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transfercomplete.md)
: void [transferFail()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transferfail.md)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transferinitiated.md)
: void

### Constants  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-constants.md)

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

### Methods  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-construct.md)

`
    public
                    __construct([ProgressBarInterface $progressBar = new ConsoleProgressBar() ][, mixed $output = STDOUT ][, bool $clear = true ][, DirectoryTransferProgressSnapshot|null $currentSnapshot = null ][, bool $showProgressOnUpdate = true ]) : mixed`

##### Parameters

$progressBar
: [ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)
= new ConsoleProgressBar()$output
: mixed
= STDOUT$clear
: bool
= true$currentSnapshot
: [DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md) \|null
= null$showProgressOnUpdate
: bool
= true

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-bytestransferred.md)

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

#### getProgressBar()  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-getprogressbar.md)

`
    public
                    getProgressBar() : ProgressBarInterface`

##### Return values

[ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)

#### priority()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_priority)

To provide an order on which listener is notified first.

`
    public
                    priority() : int`

By default, it will provide a neutral value.

##### Return values

int

#### showProgress()  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-showprogress.md)

To show the progress being tracked.

`
    public
                    showProgress() : void`

#### transferComplete()  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transfercomplete.md)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

#### transferFail()  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transferfail.md)

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

#### transferInitiated()  [header link](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transferinitiated.md)

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
  - [Constants](class-aws-s3-s3transfer-progress-directoryprogresstracker-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-directoryprogresstracker-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-construct.md)
  - [bytesTransferred()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-bytestransferred.md)
  - [getProgressBar()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-getprogressbar.md)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [showProgress()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-showprogress.md)
  - [transferComplete()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transfercomplete.md)
  - [transferFail()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transferfail.md)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-directoryprogresstracker-method-transferinitiated.md)

[Back To Top](class-aws-s3-s3transfer-progress-directoryprogresstracker-top.md)
