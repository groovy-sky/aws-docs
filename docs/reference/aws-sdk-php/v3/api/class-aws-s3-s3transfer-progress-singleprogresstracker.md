Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## SingleProgressTracker     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)       implements  [ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

FinalYes

To track single object transfers.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#toc-interfaces)

[ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#toc-constants)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method___construct)
: mixed [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_bytesTransferred)
: bool [getCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_getCurrentSnapshot)
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null [getOutput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_getOutput)
: mixed [getProgressBar()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_getProgressBar)
: [ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)[isClear()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_isClear)
: bool [isShowProgressOnUpdate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_isShowProgressOnUpdate)
: bool [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[showProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_showProgress)
: void To show the progress being tracked.[transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_transferComplete)
: void [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_transferFail)
: void [transferInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_transferInitiated)
: void

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#constants)

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method___construct)

`
    public
                    __construct([ProgressBarInterface $progressBar = new ConsoleProgressBar() ][, mixed|false|resource $output = STDOUT ][, bool $clear = true ][, TransferProgressSnapshot|null $currentSnapshot = null ][, bool $showProgressOnUpdate = true ]) : mixed`

##### Parameters

$progressBar
: [ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)
= new ConsoleProgressBar()$output
: mixed\|false\|resource
= STDOUT$clear
: bool
= true$currentSnapshot
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null
= null$showProgressOnUpdate
: bool
= true

#### bytesTransferred()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_bytesTransferred)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_bytesTransferred\#tags)

inheritDoc

##### Return values

bool
—

true to notify successful handling otherwise false.

#### getCurrentSnapshot()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_getCurrentSnapshot)

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot|null`

##### Return values

[TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html) \|null

#### getOutput()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_getOutput)

`
    public
                    getOutput() : mixed`

#### getProgressBar()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_getProgressBar)

`
    public
                    getProgressBar() : ProgressBarInterface`

##### Return values

[ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)

#### isClear()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_isClear)

`
    public
                    isClear() : bool`

##### Return values

bool

#### isShowProgressOnUpdate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_isShowProgressOnUpdate)

`
    public
                    isShowProgressOnUpdate() : bool`

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

#### showProgress()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_showProgress)

To show the progress being tracked.

`
    public
                    showProgress() : void`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_showProgress\#tags)

inheritDoc

#### transferComplete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_transferComplete)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_transferComplete\#tags)

inheritDoc

#### transferFail()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_transferFail)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_transferFail\#tags)

inheritDoc

#### transferInitiated()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_transferInitiated)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the request initialization.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html\#method_transferInitiated\#tags)

inheritDoc
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#toc-methods)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method___construct)
  - [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_bytesTransferred)
  - [getCurrentSnapshot()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_getCurrentSnapshot)
  - [getOutput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_getOutput)
  - [getProgressBar()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_getProgressBar)
  - [isClear()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_isClear)
  - [isShowProgressOnUpdate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_isShowProgressOnUpdate)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [showProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_showProgress)
  - [transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_transferComplete)
  - [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_transferFail)
  - [transferInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#method_transferInitiated)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.SingleProgressTracker.html#top)
