Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## MultiProgressTracker     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)       implements  [ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#toc-interfaces)

[ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#toc-constants)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method___construct)
: mixed [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_bytesTransferred)
: bool [getCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getCompleted)
: int [getFailed()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getFailed)
: int [getOutput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getOutput)
: mixed [getProgressBarFactory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getProgressBarFactory)
: [ProgressBarFactoryInterface](class-aws-s3-s3transfer-progress-progressbarfactoryinterface.md) \|Closure\|null [getSingleProgressTrackers()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getSingleProgressTrackers)
: array<string\|int, mixed> [getTransferCount()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getTransferCount)
: int [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[showProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_showProgress)
: void To show the progress being tracked.[transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_transferComplete)
: void [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_transferFail)
: void [transferInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_transferInitiated)
: void

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#constants)

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method___construct)

`
    public
                    __construct([array<string|int, mixed> $singleProgressTrackers = [] ][, mixed|false|resource $output = STDOUT ][, int $transferCount = 0 ][, int $completed = 0 ][, int $failed = 0 ][, ProgressBarFactoryInterface|Closure|null $progressBarFactory = null ]) : mixed`

##### Parameters

$singleProgressTrackers
: array<string\|int, mixed>
= \[\]$output
: mixed\|false\|resource
= STDOUT$transferCount
: int
= 0$completed
: int
= 0$failed
: int
= 0$progressBarFactory
: [ProgressBarFactoryInterface](class-aws-s3-s3transfer-progress-progressbarfactoryinterface.md) \|Closure\|null
= null

#### bytesTransferred()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_bytesTransferred)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_bytesTransferred\#tags)

inheritDoc

##### Return values

bool
—

true to notify successful handling otherwise false.

#### getCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_getCompleted)

`
    public
                    getCompleted() : int`

##### Return values

int

#### getFailed()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_getFailed)

`
    public
                    getFailed() : int`

##### Return values

int

#### getOutput()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_getOutput)

`
    public
                    getOutput() : mixed`

#### getProgressBarFactory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_getProgressBarFactory)

`
    public
                    getProgressBarFactory() : ProgressBarFactoryInterface|Closure|null`

##### Return values

[ProgressBarFactoryInterface](class-aws-s3-s3transfer-progress-progressbarfactoryinterface.md) \|Closure\|null

#### getSingleProgressTrackers()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_getSingleProgressTrackers)

`
    public
                    getSingleProgressTrackers() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getTransferCount()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_getTransferCount)

`
    public
                    getTransferCount() : int`

##### Return values

int

#### priority()  [header link](class-aws-s3-s3transfer-progress-abstracttransferlistener.md\#method_priority)

To provide an order on which listener is notified first.

`
    public
                    priority() : int`

By default, it will provide a neutral value.

##### Return values

int

#### showProgress()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_showProgress)

To show the progress being tracked.

`
    public
                    showProgress() : void`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_showProgress\#tags)

inheritDoc

#### transferComplete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_transferComplete)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_transferComplete\#tags)

inheritDoc

#### transferFail()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_transferFail)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_transferFail\#tags)

inheritDoc

#### transferInitiated()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_transferInitiated)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the request initialization.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html\#method_transferInitiated\#tags)

inheritDoc
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#toc-methods)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method___construct)
  - [bytesTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_bytesTransferred)
  - [getCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getCompleted)
  - [getFailed()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getFailed)
  - [getOutput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getOutput)
  - [getProgressBarFactory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getProgressBarFactory)
  - [getSingleProgressTrackers()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getSingleProgressTrackers)
  - [getTransferCount()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_getTransferCount)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [showProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_showProgress)
  - [transferComplete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_transferComplete)
  - [transferFail()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_transferFail)
  - [transferInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#method_transferInitiated)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.MultiProgressTracker.html#top)
