Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## ObjectUploader        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Uploads an object to S3, using a PutObject command or a multipart upload as
appropriate.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#toc-interfaces)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#toc-constants)

[DEFAULT\_MULTIPART\_THRESHOLD](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#constant_DEFAULT_MULTIPART_THRESHOLD)
= 16777216

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#method___construct)
: mixed [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.[upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#method_upload)
: mixed

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#constants)

#### DEFAULT\_MULTIPART\_THRESHOLD  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#constant_DEFAULT_MULTIPART_THRESHOLD)

`
    public
        mixed
    DEFAULT_MULTIPART_THRESHOLD
    = 16777216
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#method___construct)

`
    public
                    __construct(S3ClientInterface $client, string $bucket, string $key, mixed $body[, string $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

The S3 Client used to execute
the upload command(s).

$bucket
: string

Bucket to upload the object, or
an S3 access point ARN.

$key
: string

Key of the object.

$body
: mixed

Object data to upload. Can be a
StreamInterface, PHP stream
resource, or a string of data to
upload.

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
the sub command(s).

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#method_promise)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### upload()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html\#method_upload)

`
    public
                    upload() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#toc-methods)
- Constants
  - [DEFAULT\_MULTIPART\_THRESHOLD](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#constant_DEFAULT_MULTIPART_THRESHOLD)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#method___construct)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#method_promise)
  - [upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#method_upload)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html#top)
