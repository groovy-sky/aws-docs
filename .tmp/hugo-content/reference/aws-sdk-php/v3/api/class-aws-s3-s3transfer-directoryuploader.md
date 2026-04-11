Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)

## DirectoryUploader        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-directoryuploader-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-directoryuploader-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](class-aws-s3-s3transfer-directoryuploader-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-directoryuploader-method-construct.md)
: mixed [promise()](class-aws-s3-s3transfer-directoryuploader-method-promise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Methods  [header link](class-aws-s3-s3transfer-directoryuploader-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-directoryuploader-method-construct.md)

`
    public
                    __construct(S3ClientInterface $s3Client, array<string|int, mixed> $config, Closure $uploadObject, UploadDirectoryRequest $uploadDirectoryRequest) : mixed`

##### Parameters

$s3Client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)$config
: array<string\|int, mixed>$uploadObject
: Closure

A closure that receives
(S3ClientInterface, UploadRequest) and returns PromiseInterface

$uploadDirectoryRequest
: [UploadDirectoryRequest](class-aws-s3-s3transfer-models-uploaddirectoryrequest.md)

#### promise()  [header link](class-aws-s3-s3transfer-directoryuploader-method-promise.md)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Tags  [header link](class-aws-s3-s3transfer-directoryuploader-method-promise-tags.md)

throwsThrowable

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3transfer-directoryuploader-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-directoryuploader-method-construct.md)
  - [promise()](class-aws-s3-s3transfer-directoryuploader-method-promise.md)

[Back To Top](class-aws-s3-s3transfer-directoryuploader-top.md)
