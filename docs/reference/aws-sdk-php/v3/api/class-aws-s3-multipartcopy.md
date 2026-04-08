Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## MultipartCopy     extends AbstractUploadManager   in package    - [Aws](package-aws.md)       Uses  [MultipartUploadingTrait](class-aws-s3-multipartuploadingtrait.md)

### Table of Contents  [header link](class-aws-s3-multipartcopy-toc.md)

#### Methods  [header link](class-aws-s3-multipartcopy-toc-methods.md)

[\_\_construct()](class-aws-s3-multipartcopy-method-construct.md)
: mixed Creates a multipart upload for copying an S3 object.[copy()](class-aws-s3-multipartcopy-method-copy.md)
: mixed An alias of the self::upload method.[getStateFromService()](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)
: [UploadState](class-aws-multipart-uploadstate.md)Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

### Methods  [header link](class-aws-s3-multipartcopy-methods.md)

#### \_\_construct()  [header link](class-aws-s3-multipartcopy-method-construct.md)

Creates a multipart upload for copying an S3 object.

`
    public
                    __construct(S3ClientInterface $client, string|array<string|int, mixed> $source[, array<string|int, mixed> $config = [] ]) : mixed`

The valid configuration options are as follows:

- acl: (string) ACL to set on the object being upload. Objects are
private by default.
- before\_complete: (callable) Callback to invoke before the
`CompleteMultipartUpload` operation. The callback should have a
function signature like `function (Aws\Command $command) {...}`.
- before\_initiate: (callable) Callback to invoke before the
`CreateMultipartUpload` operation. The callback should have a function
signature like `function (Aws\Command $command) {...}`.
- before\_upload: (callable) Callback to invoke before `UploadPartCopy`
operations. The callback should have a function signature like
`function (Aws\Command $command) {...}`.
- bucket: (string, required) Name of the bucket to which the object is
being uploaded.
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
- state: (Aws\\Multipart\\UploadState) An object that represents the state
of the multipart upload and that is used to resume a previous upload.
When this option is provided, the `bucket`, `key`, and `part_size`
options are ignored.
- source\_metadata: (Aws\\ResultInterface) An object that represents the
result of executing a HeadObject command on the copy source.
- display\_progress: (boolean) Set true to track status in 1/8th increments
for upload.

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

Client used for the upload.

$source
: string\|array<string\|int, mixed>

Location of the data to be copied (in the
form //). If the key contains a '?'
character, instead pass an array of source\_key,
source\_bucket, and source\_version\_id.

$config
: array<string\|int, mixed>
= \[\]

Configuration used to perform the upload.

#### copy()  [header link](class-aws-s3-multipartcopy-method-copy.md)

An alias of the self::upload method.

`
    public
                    copy() : mixed`

##### Tags  [header link](class-aws-s3-multipartcopy-method-copy-tags.md)

seeself::upload

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
  - [Methods](class-aws-s3-multipartcopy-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-multipartcopy-method-construct.md)
  - [copy()](class-aws-s3-multipartcopy-method-copy.md)
  - [getStateFromService()](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)

[Back To Top](class-aws-s3-multipartcopy-top.md)
