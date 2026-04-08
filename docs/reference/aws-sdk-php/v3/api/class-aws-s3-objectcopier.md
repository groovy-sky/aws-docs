Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## ObjectCopier        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Copies objects from one S3 location to another, utilizing a multipart copy
when appropriate.

### Table of Contents  [header link](class-aws-s3-objectcopier-toc.md)

#### Interfaces  [header link](class-aws-s3-objectcopier-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Constants  [header link](class-aws-s3-objectcopier-toc-constants.md)

[DEFAULT\_MULTIPART\_THRESHOLD](class-aws-s3-objectcopier-constant-default-multipart-threshold.md)
= \\Aws\\S3\\MultipartUploader::PART\_MAX\_SIZE

#### Methods  [header link](class-aws-s3-objectcopier-toc-methods.md)

[\_\_construct()](class-aws-s3-objectcopier-method-construct.md)
: mixed [copy()](class-aws-s3-objectcopier-method-copy.md)
: [Result](class-aws-result.md)Perform the configured copy synchronously. Returns the result of the
CompleteMultipartUpload or CopyObject operation.[promise()](class-aws-s3-objectcopier-method-promise.md)
: [Coroutine](class-guzzlehttp-promise-coroutine.md)Perform the configured copy asynchronously. Returns a promise that is
fulfilled with the result of the CompleteMultipartUpload or CopyObject
operation or rejected with an exception.

### Constants  [header link](class-aws-s3-objectcopier-constants.md)

#### DEFAULT\_MULTIPART\_THRESHOLD  [header link](class-aws-s3-objectcopier-constant-default-multipart-threshold.md)

`
    public
        mixed
    DEFAULT_MULTIPART_THRESHOLD
    = \Aws\S3\MultipartUploader::PART_MAX_SIZE
`

### Methods  [header link](class-aws-s3-objectcopier-methods.md)

#### \_\_construct()  [header link](class-aws-s3-objectcopier-method-construct.md)

`
    public
                    __construct(S3ClientInterface $client, array<string|int, mixed> $source, array<string|int, mixed> $destination[, string $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

The S3 Client used to execute
the copy command(s).

$source
: array<string\|int, mixed>

The object to copy, specified as
an array with a 'Bucket' and
'Key' keys. Provide a
'VersionID' key to copy a
specified version of an object.

$destination
: array<string\|int, mixed>

The bucket and key to which to
copy the $source, specified as
an array with a 'Bucket' and
'Key' keys.

$acl
: string
= 'private'

ACL to apply to the copy
(default: private).

$options
: array<string\|int, mixed>
= \[\]

Options used to configure the
copy process. Options passed in
through 'params' are added to
the sub commands.

##### Tags  [header link](class-aws-s3-objectcopier-method-construct-tags.md)

throwsInvalidArgumentException

#### copy()  [header link](class-aws-s3-objectcopier-method-copy.md)

Perform the configured copy synchronously. Returns the result of the
CompleteMultipartUpload or CopyObject operation.

`
    public
                    copy() : Result`

##### Tags  [header link](class-aws-s3-objectcopier-method-copy-tags.md)

throws[S3Exception](class-aws-s3-exception-s3exception.md)throws[MultipartUploadException](class-aws-exception-multipartuploadexception.md)

##### Return values

[Result](class-aws-result.md)

#### promise()  [header link](class-aws-s3-objectcopier-method-promise.md)

Perform the configured copy asynchronously. Returns a promise that is
fulfilled with the result of the CompleteMultipartUpload or CopyObject
operation or rejected with an exception.

`
    public
                    promise() : Coroutine`

##### Return values

[Coroutine](class-guzzlehttp-promise-coroutine.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-objectcopier-toc-constants.md)
  - [Methods](class-aws-s3-objectcopier-toc-methods.md)
- Constants
  - [DEFAULT\_MULTIPART\_THRESHOLD](class-aws-s3-objectcopier-constant-default-multipart-threshold.md)
- Methods
  - [\_\_construct()](class-aws-s3-objectcopier-method-construct.md)
  - [copy()](class-aws-s3-objectcopier-method-copy.md)
  - [promise()](class-aws-s3-objectcopier-method-promise.md)

[Back To Top](class-aws-s3-objectcopier-top.md)
