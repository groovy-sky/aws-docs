Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## SingleProgressTracker     extends [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md)   in package    - [Aws](package-aws.md)       implements  [ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

FinalYes

To track single object transfers.

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-toc-interfaces.md)

[ProgressTrackerInterface](class-aws-s3-s3transfer-progress-progresstrackerinterface.md)

#### Constants  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-construct.md)
: mixed [bytesTransferred()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-bytestransferred.md)
: bool [getCurrentSnapshot()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getcurrentsnapshot.md)
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null [getOutput()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getoutput.md)
: mixed [getProgressBar()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getprogressbar.md)
: [ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)[isClear()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-isclear.md)
: bool [isShowProgressOnUpdate()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-isshowprogressonupdate.md)
: bool [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[showProgress()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-showprogress.md)
: void To show the progress being tracked.[transferComplete()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transfercomplete.md)
: void [transferFail()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transferfail.md)
: void [transferInitiated()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transferinitiated.md)
: void

### Constants  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-constants.md)

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

### Methods  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-construct.md)

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
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null
= null$showProgressOnUpdate
: bool
= true

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-bytestransferred.md)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-bytestransferred-tags.md)

inheritDoc

##### Return values

bool
—

true to notify successful handling otherwise false.

#### getCurrentSnapshot()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getcurrentsnapshot.md)

`
    public
                    getCurrentSnapshot() : TransferProgressSnapshot|null`

##### Return values

[TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md) \|null

#### getOutput()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getoutput.md)

`
    public
                    getOutput() : mixed`

#### getProgressBar()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getprogressbar.md)

`
    public
                    getProgressBar() : ProgressBarInterface`

##### Return values

[ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)

#### isClear()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-isclear.md)

`
    public
                    isClear() : bool`

##### Return values

bool

#### isShowProgressOnUpdate()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-isshowprogressonupdate.md)

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

#### showProgress()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-showprogress.md)

To show the progress being tracked.

`
    public
                    showProgress() : void`

##### Tags  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-showprogress-tags.md)

inheritDoc

#### transferComplete()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transfercomplete.md)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the operation that originated the bytes transferred event.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transfercomplete-tags.md)

inheritDoc

#### transferFail()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transferfail.md)

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

##### Tags  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transferfail-tags.md)

inheritDoc

#### transferInitiated()  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transferinitiated.md)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

- request\_args: (array) The request arguments that will be provided
as part of the request initialization.
- progress\_snapshot: (TransferProgressSnapshot) The transfer snapshot holder.

##### Tags  [header link](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transferinitiated-tags.md)

inheritDoc
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3transfer-progress-singleprogresstracker-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-singleprogresstracker-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-construct.md)
  - [bytesTransferred()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-bytestransferred.md)
  - [getCurrentSnapshot()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getcurrentsnapshot.md)
  - [getOutput()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getoutput.md)
  - [getProgressBar()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-getprogressbar.md)
  - [isClear()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-isclear.md)
  - [isShowProgressOnUpdate()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-isshowprogressonupdate.md)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [showProgress()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-showprogress.md)
  - [transferComplete()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transfercomplete.md)
  - [transferFail()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transferfail.md)
  - [transferInitiated()](class-aws-s3-s3transfer-progress-singleprogresstracker-method-transferinitiated.md)

[Back To Top](class-aws-s3-s3transfer-progress-singleprogresstracker-top.md)
