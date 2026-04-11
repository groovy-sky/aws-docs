Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## ResumeDownloadRequest        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-models-resumedownloadrequest-toc.md)

#### Methods  [header link](class-aws-s3-s3transfer-models-resumedownloadrequest-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-construct.md)
: mixed [getDownloadHandlerClass()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getdownloadhandlerclass.md)
: string [getListeners()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getlisteners.md)
: array<string\|int, mixed> [getProgressTracker()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getprogresstracker.md)
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null [getResumableDownload()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getresumabledownload.md)
: string\| [ResumableDownload](class-aws-s3-s3transfer-models-resumabledownload.md)

### Methods  [header link](class-aws-s3-s3transfer-models-resumedownloadrequest-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-resumedownloadrequest-method-construct.md)

`
    public
                    __construct(ResumableDownload|string $resumableDownload[, string $downloadHandlerClass = FileDownloadHandler::class ][, array<string|int, mixed> $listeners = [] ][, AbstractTransferListener|null $progressTracker = null ]) : mixed`

##### Parameters

$resumableDownload
: [ResumableDownload](class-aws-s3-s3transfer-models-resumabledownload.md) \|string$downloadHandlerClass
: string
= FileDownloadHandler::class$listeners
: array<string\|int, mixed>
= \[\]$progressTracker
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null
= null

#### getDownloadHandlerClass()  [header link](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getdownloadhandlerclass.md)

`
    public
                    getDownloadHandlerClass() : string`

##### Return values

string

#### getListeners()  [header link](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getlisteners.md)

`
    public
                    getListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getProgressTracker()  [header link](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getprogresstracker.md)

`
    public
                    getProgressTracker() : AbstractTransferListener|null`

##### Return values

[AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null

#### getResumableDownload()  [header link](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getresumabledownload.md)

`
    public
                    getResumableDownload() : string|ResumableDownload`

##### Return values

string\| [ResumableDownload](class-aws-s3-s3transfer-models-resumabledownload.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3transfer-models-resumedownloadrequest-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-construct.md)
  - [getDownloadHandlerClass()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getdownloadhandlerclass.md)
  - [getListeners()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getlisteners.md)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getprogresstracker.md)
  - [getResumableDownload()](class-aws-s3-s3transfer-models-resumedownloadrequest-method-getresumabledownload.md)

[Back To Top](class-aws-s3-s3transfer-models-resumedownloadrequest-top.md)
