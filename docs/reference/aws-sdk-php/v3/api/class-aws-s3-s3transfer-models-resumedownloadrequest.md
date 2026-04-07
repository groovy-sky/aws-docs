Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## ResumeDownloadRequest        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method___construct)
: mixed [getDownloadHandlerClass()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method_getDownloadHandlerClass)
: string [getListeners()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method_getListeners)
: array<string\|int, mixed> [getProgressTracker()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method_getProgressTracker)
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null [getResumableDownload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method_getResumableDownload)
: string\| [ResumableDownload](class-aws-s3-s3transfer-models-resumabledownload.md)

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html\#method___construct)

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
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null
= null

#### getDownloadHandlerClass()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html\#method_getDownloadHandlerClass)

`
    public
                    getDownloadHandlerClass() : string`

##### Return values

string

#### getListeners()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html\#method_getListeners)

`
    public
                    getListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getProgressTracker()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html\#method_getProgressTracker)

`
    public
                    getProgressTracker() : AbstractTransferListener|null`

##### Return values

[AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null

#### getResumableDownload()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html\#method_getResumableDownload)

`
    public
                    getResumableDownload() : string|ResumableDownload`

##### Return values

string\| [ResumableDownload](class-aws-s3-s3transfer-models-resumabledownload.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method___construct)
  - [getDownloadHandlerClass()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method_getDownloadHandlerClass)
  - [getListeners()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method_getListeners)
  - [getProgressTracker()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method_getProgressTracker)
  - [getResumableDownload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#method_getResumableDownload)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.ResumeDownloadRequest.html#top)
