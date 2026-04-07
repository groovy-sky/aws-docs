Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## ObjectCopier        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Copies objects from one S3 location to another, utilizing a multipart copy
when appropriate.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#toc-interfaces)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#toc-constants)

[DEFAULT\_MULTIPART\_THRESHOLD](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#constant_DEFAULT_MULTIPART_THRESHOLD)
= \\Aws\\S3\\MultipartUploader::PART\_MAX\_SIZE

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#method___construct)
: mixed [copy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#method_copy)
: [Result](class-aws-result.md)Perform the configured copy synchronously. Returns the result of the
CompleteMultipartUpload or CopyObject operation.[promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#method_promise)
: [Coroutine](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Coroutine.html)Perform the configured copy asynchronously. Returns a promise that is
fulfilled with the result of the CompleteMultipartUpload or CopyObject
operation or rejected with an exception.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#constants)

#### DEFAULT\_MULTIPART\_THRESHOLD  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#constant_DEFAULT_MULTIPART_THRESHOLD)

`
    public
        mixed
    DEFAULT_MULTIPART_THRESHOLD
    = \Aws\S3\MultipartUploader::PART_MAX_SIZE
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#method___construct)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#method___construct\#tags)

throwsInvalidArgumentException

#### copy()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#method_copy)

Perform the configured copy synchronously. Returns the result of the
CompleteMultipartUpload or CopyObject operation.

`
    public
                    copy() : Result`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#method_copy\#tags)

throws[S3Exception](class-aws-s3-exception-s3exception.md)throws[MultipartUploadException](class-aws-exception-multipartuploadexception.md)

##### Return values

[Result](class-aws-result.md)

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html\#method_promise)

Perform the configured copy asynchronously. Returns a promise that is
fulfilled with the result of the CompleteMultipartUpload or CopyObject
operation or rejected with an exception.

`
    public
                    promise() : Coroutine`

##### Return values

[Coroutine](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Coroutine.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#toc-methods)
- Constants
  - [DEFAULT\_MULTIPART\_THRESHOLD](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#constant_DEFAULT_MULTIPART_THRESHOLD)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#method___construct)
  - [copy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#method_copy)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#method_promise)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html#top)
