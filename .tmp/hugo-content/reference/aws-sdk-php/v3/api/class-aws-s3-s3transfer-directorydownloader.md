Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)

## DirectoryDownloader        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-directorydownloader-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-directorydownloader-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](class-aws-s3-s3transfer-directorydownloader-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-directorydownloader-method-construct.md)
: mixed [promise()](class-aws-s3-s3transfer-directorydownloader-method-promise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Methods  [header link](class-aws-s3-s3transfer-directorydownloader-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-directorydownloader-method-construct.md)

`
    public
                    __construct(S3ClientInterface $s3Client, array<string|int, mixed> $config, Closure $downloadFile, DownloadDirectoryRequest $downloadDirectoryRequest) : mixed`

##### Parameters

$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)$config
: array<string\|int, mixed>$downloadFile
: Closure

A closure that receives (S3ClientInterface, DownloadFileRequest) and returns PromiseInterface

$downloadDirectoryRequest
: [DownloadDirectoryRequest](class-aws-s3-s3transfer-models-downloaddirectoryrequest.md)

#### promise()  [header link](class-aws-s3-s3transfer-directorydownloader-method-promise.md)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Tags  [header link](class-aws-s3-s3transfer-directorydownloader-method-promise-tags.md)

throwsThrowable

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3transfer-directorydownloader-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-directorydownloader-method-construct.md)
  - [promise()](class-aws-s3-s3transfer-directorydownloader-method-promise.md)

[Back To Top](class-aws-s3-s3transfer-directorydownloader-top.md)
