Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## MultipartUploader     extends [AbstractUploader](class-aws-multipart-abstractuploader.md)   in package    - [Aws](package-aws.md)       Uses  [MultipartUploadingTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploadingTrait.html)

Encapsulates the execution of a multipart upload to S3 or Glacier.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#toc-constants)

[PART\_MAX\_NUM](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MAX_NUM)
= 10000 [PART\_MAX\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MAX_SIZE)
= 5368709120 [PART\_MIN\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MIN_SIZE)
= 5242880

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#method___construct)
: mixed Creates a multipart upload for an S3 object.[getStateFromService()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploadingTrait.html#method_getStateFromService)
: [UploadState](class-aws-multipart-uploadstate.md)Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#constants)

#### PART\_MAX\_NUM  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#constant_PART_MAX_NUM)

`
    public
        mixed
    PART_MAX_NUM
    = 10000
`

#### PART\_MAX\_SIZE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#constant_PART_MAX_SIZE)

`
    public
        mixed
    PART_MAX_SIZE
    = 5368709120
`

#### PART\_MIN\_SIZE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#constant_PART_MIN_SIZE)

`
    public
        mixed
    PART_MIN_SIZE
    = 5242880
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html\#method___construct)

Creates a multipart upload for an S3 object.

`
    public
                    __construct(S3ClientInterface $client, mixed $source[, array<string|int, mixed> $config = [] ]) : mixed`

The valid configuration options are as follows:

- acl: (string) ACL to set on the object being upload. Objects are
private by default.
- before\_complete: (callable) Callback to invoke before the
`CompleteMultipartUpload` operation. The callback should have a
function signature like `function (Aws\Command $command) {...}`.
- before\_initiate: (callable) Callback to invoke before the
`CreateMultipartUpload` operation. The callback should have a function
signature like `function (Aws\Command $command) {...}`.
- before\_upload: (callable) Callback to invoke before any `UploadPart`
operations. The callback should have a function signature like
`function (Aws\Command $command) {...}`.
- bucket: (string, required) Name of the bucket to which the object is
being uploaded, or an S3 access point ARN.
- concurrency: (int, default=int(5)) Maximum number of concurrent
`UploadPart` operations allowed during the multipart upload.
- key: (string, required) Key to use for the object being uploaded.
- params: (array) An array of key/value parameters that will be applied
to each of the sub-commands run by the uploader as a base.
Auto-calculated options will override these parameters. If you need
more granularity over parameters to each sub-command, use the before\_\*
options detailed above to update the commands directly.
- part\_size: (int, default=int(5242880)) Part size, in bytes, to use when
doing a multipart upload. This must between 5 MB and 5 GB, inclusive.
- prepare\_data\_source: (callable) Callback to invoke before starting the
multipart upload workflow. The callback should have a function
signature like `function () {...}`.
- state: (Aws\\Multipart\\UploadState) An object that represents the state
of the multipart upload and that is used to resume a previous upload.
When this option is provided, the `bucket`, `key`, and `part_size`
options are ignored.
- track\_upload: (boolean) Set true to track status in 1/8th increments
for upload.

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

Client used for the upload.

$source
: mixed

Source of the data to upload.

$config
: array<string\|int, mixed>
= \[\]

Configuration used to perform the upload.

#### getStateFromService()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploadingTrait.html\#method_getStateFromService)

Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

`
    public
            static        getStateFromService(S3ClientInterface $client, string $bucket, string $key, string $uploadId) : UploadState`

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

S3Client used for the upload.

$bucket
: string

Bucket for the multipart upload.

$key
: string

Object key for the multipart upload.

$uploadId
: string

Upload ID for the multipart upload.

##### Return values

[UploadState](class-aws-multipart-uploadstate.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#toc-methods)
- Constants
  - [PART\_MAX\_NUM](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MAX_NUM)
  - [PART\_MAX\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MAX_SIZE)
  - [PART\_MIN\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MIN_SIZE)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#method___construct)
  - [getStateFromService()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploadingTrait.html#method_getStateFromService)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#top)
