Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## ObjectUploader        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Uploads an object to S3, using a PutObject command or a multipart upload as
appropriate.

### Table of Contents  [header link](class-aws-s3-objectuploader-toc.md)

#### Interfaces  [header link](class-aws-s3-objectuploader-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Constants  [header link](class-aws-s3-objectuploader-toc-constants.md)

[DEFAULT\_MULTIPART\_THRESHOLD](class-aws-s3-objectuploader-constant-default-multipart-threshold.md)
= 16777216

#### Methods  [header link](class-aws-s3-objectuploader-toc-methods.md)

[\_\_construct()](class-aws-s3-objectuploader-method-construct.md)
: mixed [promise()](class-aws-s3-objectuploader-method-promise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.[upload()](class-aws-s3-objectuploader-method-upload.md)
: mixed

### Constants  [header link](class-aws-s3-objectuploader-constants.md)

#### DEFAULT\_MULTIPART\_THRESHOLD  [header link](class-aws-s3-objectuploader-constant-default-multipart-threshold.md)

`
    public
        mixed
    DEFAULT_MULTIPART_THRESHOLD
    = 16777216
`

### Methods  [header link](class-aws-s3-objectuploader-methods.md)

#### \_\_construct()  [header link](class-aws-s3-objectuploader-method-construct.md)

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

#### promise()  [header link](class-aws-s3-objectuploader-method-promise.md)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### upload()  [header link](class-aws-s3-objectuploader-method-upload.md)

`
    public
                    upload() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-objectuploader-toc-constants.md)
  - [Methods](class-aws-s3-objectuploader-toc-methods.md)
- Constants
  - [DEFAULT\_MULTIPART\_THRESHOLD](class-aws-s3-objectuploader-constant-default-multipart-threshold.md)
- Methods
  - [\_\_construct()](class-aws-s3-objectuploader-method-construct.md)
  - [promise()](class-aws-s3-objectuploader-method-promise.md)
  - [upload()](class-aws-s3-objectuploader-method-upload.md)

[Back To Top](class-aws-s3-objectuploader-top.md)
