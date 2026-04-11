Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## MultipartUploader     extends [AbstractUploader](class-aws-multipart-abstractuploader.md)   in package    - [Aws](package-aws.md)       Uses  [MultipartUploadingTrait](class-aws-s3-multipartuploadingtrait.md)

Encapsulates the execution of a multipart upload to S3 or Glacier.

### Table of Contents  [header link](class-aws-s3-multipartuploader-toc.md)

#### Constants  [header link](class-aws-s3-multipartuploader-toc-constants.md)

[PART\_MAX\_NUM](class-aws-s3-multipartuploader-constant-part-max-num.md)
= 10000 [PART\_MAX\_SIZE](class-aws-s3-multipartuploader-constant-part-max-size.md)
= 5368709120 [PART\_MIN\_SIZE](class-aws-s3-multipartuploader-constant-part-min-size.md)
= 5242880

#### Methods  [header link](class-aws-s3-multipartuploader-toc-methods.md)

[\_\_construct()](class-aws-s3-multipartuploader-method-construct.md)
: mixed Creates a multipart upload for an S3 object.[getStateFromService()](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)
: [UploadState](class-aws-multipart-uploadstate.md)Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

### Constants  [header link](class-aws-s3-multipartuploader-constants.md)

#### PART\_MAX\_NUM  [header link](class-aws-s3-multipartuploader-constant-part-max-num.md)

`
    public
        mixed
    PART_MAX_NUM
    = 10000
`

#### PART\_MAX\_SIZE  [header link](class-aws-s3-multipartuploader-constant-part-max-size.md)

`
    public
        mixed
    PART_MAX_SIZE
    = 5368709120
`

#### PART\_MIN\_SIZE  [header link](class-aws-s3-multipartuploader-constant-part-min-size.md)

`
    public
        mixed
    PART_MIN_SIZE
    = 5242880
`

### Methods  [header link](class-aws-s3-multipartuploader-methods.md)

#### \_\_construct()  [header link](class-aws-s3-multipartuploader-method-construct.md)

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

#### getStateFromService()  [header link](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)

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
  - [Constants](class-aws-s3-multipartuploader-toc-constants.md)
  - [Methods](class-aws-s3-multipartuploader-toc-methods.md)
- Constants
  - [PART\_MAX\_NUM](class-aws-s3-multipartuploader-constant-part-max-num.md)
  - [PART\_MAX\_SIZE](class-aws-s3-multipartuploader-constant-part-max-size.md)
  - [PART\_MIN\_SIZE](class-aws-s3-multipartuploader-constant-part-min-size.md)
- Methods
  - [\_\_construct()](class-aws-s3-multipartuploader-method-construct.md)
  - [getStateFromService()](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)

[Back To Top](class-aws-s3-multipartuploader-top.md)
