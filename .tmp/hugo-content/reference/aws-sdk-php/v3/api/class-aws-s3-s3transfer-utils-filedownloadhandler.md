Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Utils](namespace-aws-s3-s3transfer-utils.md)

## FileDownloadHandler     extends [AbstractDownloadHandler](class-aws-s3-s3transfer-utils-abstractdownloadhandler.md)   in package    - [Aws](package-aws.md)       implements  [ResumableDownloadHandlerInterface](class-aws-s3-s3transfer-utils-resumabledownloadhandlerinterface.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-toc-interfaces.md)

[ResumableDownloadHandlerInterface](class-aws-s3-s3transfer-utils-resumabledownloadhandlerinterface.md)

#### Constants  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-toc-constants.md)

[PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
= 'progress\_snapshot' [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
= 'reason' [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
= 'request\_args'

#### Methods  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-construct.md)
: mixed [bytesTransferred()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-bytestransferred.md)
: bool [getDestination()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getdestination.md)
: string [getFixedPartSize()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getfixedpartsize.md)
: int [getHandlerResult()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-gethandlerresult.md)
: string Returns the handler result.[getResumeFilePath()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getresumefilepath.md)
: string [getTemporaryFilePath()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-gettemporaryfilepath.md)
: string [initializeDestination()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-initializedestination.md)
: void [isConcurrencySupported()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-isconcurrencysupported.md)
: bool To control whether the download handler supports
concurrency.[isFailsWhenDestinationExists()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-isfailswhendestinationexists.md)
: bool [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
: int To provide an order on which listener is notified first.[transferComplete()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transfercomplete.md)
: void [transferFail()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transferfail.md)
: void [transferInitiated()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transferinitiated.md)
: void

### Constants  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-constants.md)

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

### Methods  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-construct.md)

`
    public
                    __construct(string $destination, bool $failsWhenDestinationExists[, bool $resumeEnabled = false ][, string|null $temporaryFilePath = null ][, int|null $fixedPartSize = null ]) : mixed`

##### Parameters

$destination
: string$failsWhenDestinationExists
: bool$resumeEnabled
: bool
= false$temporaryFilePath
: string\|null
= null$fixedPartSize
: int\|null
= null

#### bytesTransferred()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-bytestransferred.md)

`
    public
                    bytesTransferred(array<string|int, mixed> $context) : bool`

##### Parameters

$context
: array<string\|int, mixed>

##### Return values

bool

#### getDestination()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getdestination.md)

`
    public
                    getDestination() : string`

##### Return values

string

#### getFixedPartSize()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getfixedpartsize.md)

`
    public
                    getFixedPartSize() : int`

##### Return values

int

#### getHandlerResult()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-gethandlerresult.md)

Returns the handler result.

`
    public
                    getHandlerResult() : string`

##### Return values

string

#### getResumeFilePath()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getresumefilepath.md)

`
    public
                    getResumeFilePath() : string`

##### Return values

string

#### getTemporaryFilePath()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-gettemporaryfilepath.md)

`
    public
                    getTemporaryFilePath() : string`

##### Return values

string

#### initializeDestination()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-initializedestination.md)

`
    public
                    initializeDestination(array<string|int, mixed> $response) : void`

##### Parameters

$response
: array<string\|int, mixed>

#### isConcurrencySupported()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-isconcurrencysupported.md)

To control whether the download handler supports
concurrency.

`
    public
                    isConcurrencySupported() : bool`

##### Tags  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-isconcurrencysupported-tags.md)

inheritDoc

##### Return values

bool

#### isFailsWhenDestinationExists()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-isfailswhendestinationexists.md)

`
    public
                    isFailsWhenDestinationExists() : bool`

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

#### transferComplete()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transfercomplete.md)

`
    public
                    transferComplete(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

#### transferFail()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transferfail.md)

`
    public
                    transferFail(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>

#### transferInitiated()  [header link](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transferinitiated.md)

`
    public
                    transferInitiated(array<string|int, mixed> $context) : void`

##### Parameters

$context
: array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3transfer-utils-filedownloadhandler-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-utils-filedownloadhandler-toc-methods.md)
- Constants
  - [PROGRESS\_SNAPSHOT\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_PROGRESS_SNAPSHOT_KEY)
  - [REASON\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REASON_KEY)
  - [REQUEST\_ARGS\_KEY](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#constant_REQUEST_ARGS_KEY)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-construct.md)
  - [bytesTransferred()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-bytestransferred.md)
  - [getDestination()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getdestination.md)
  - [getFixedPartSize()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getfixedpartsize.md)
  - [getHandlerResult()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-gethandlerresult.md)
  - [getResumeFilePath()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-getresumefilepath.md)
  - [getTemporaryFilePath()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-gettemporaryfilepath.md)
  - [initializeDestination()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-initializedestination.md)
  - [isConcurrencySupported()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-isconcurrencysupported.md)
  - [isFailsWhenDestinationExists()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-isfailswhendestinationexists.md)
  - [priority()](class-aws-s3-s3transfer-progress-abstracttransferlistener.md#method_priority)
  - [transferComplete()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transfercomplete.md)
  - [transferFail()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transferfail.md)
  - [transferInitiated()](class-aws-s3-s3transfer-utils-filedownloadhandler-method-transferinitiated.md)

[Back To Top](class-aws-s3-s3transfer-utils-filedownloadhandler-top.md)
