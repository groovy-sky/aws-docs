Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## MultiProgressTracker     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)       implements  [ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-toc-interfaces.md)

[ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

#### Constants  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-construct.md)
: mixed [bytesTransferred()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-bytestransferred.md)
: bool [getCompleted()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getcompleted.md)
: int [getFailed()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getfailed.md)
: int [getOutput()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getoutput.md)
: mixed [getProgressBarFactory()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getprogressbarfactory.md)
: [ProgressBarFactoryInterface](class-aws-s3-s3transfer-progress-progressbarfactoryinterface.md) \|Closure\|null [getSingleProgressTrackers()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getsingleprogresstrackers.md)
: array<string\|int, mixed> [getTransferCount()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-gettransfercount.md)
: int [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[showProgress()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-showprogress.md)
: void To show the progress being tracked.[transferComplete()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transfercomplete.md)
: void [transferFail()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transferfail.md)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transferinitiated.md)
: void

### Constants  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-constants.md)

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

### Methods  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-construct.md)

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

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-bytestransferred.md)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-bytestransferred-tags.md)

inheritDoc

##### Return values

bool
—

true to notify successful handling otherwise false.

#### getCompleted()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getcompleted.md)

`
    public
                    getCompleted() : int`

##### Return values

int

#### getFailed()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getfailed.md)

`
    public
                    getFailed() : int`

##### Return values

int

#### getOutput()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getoutput.md)

`
    public
                    getOutput() : mixed`

#### getProgressBarFactory()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getprogressbarfactory.md)

`
    public
                    getProgressBarFactory() : ProgressBarFactoryInterface|Closure|null`

##### Return values

[ProgressBarFactoryInterface](class-aws-s3-s3transfer-progress-progressbarfactoryinterface.md) \|Closure\|null

#### getSingleProgressTrackers()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getsingleprogresstrackers.md)

`
    public
                    getSingleProgressTrackers() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getTransferCount()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-gettransfercount.md)

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

#### showProgress()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-showprogress.md)

To show the progress being tracked.

`
    public
                    showProgress() : void`

##### Tags  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-showprogress-tags.md)

inheritDoc

#### transferComplete()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transfercomplete.md)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transfercomplete-tags.md)

inheritDoc

#### transferFail()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transferfail.md)

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

##### Tags  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transferfail-tags.md)

inheritDoc

#### transferInitiated()  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transferinitiated.md)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the request initialization.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transferinitiated-tags.md)

inheritDoc
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3transfer-progress-multiprogresstracker-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-multiprogresstracker-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-construct.md)
  - [bytesTransferred()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-bytestransferred.md)
  - [getCompleted()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getcompleted.md)
  - [getFailed()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getfailed.md)
  - [getOutput()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getoutput.md)
  - [getProgressBarFactory()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getprogressbarfactory.md)
  - [getSingleProgressTrackers()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-getsingleprogresstrackers.md)
  - [getTransferCount()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-gettransfercount.md)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [showProgress()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-showprogress.md)
  - [transferComplete()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transfercomplete.md)
  - [transferFail()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transferfail.md)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-multiprogresstracker-method-transferinitiated.md)

[Back To Top](class-aws-s3-s3transfer-progress-multiprogresstracker-top.md)
