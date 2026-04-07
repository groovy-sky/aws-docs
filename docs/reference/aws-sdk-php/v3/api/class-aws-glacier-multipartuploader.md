Menu

- [Aws](namespace-aws.md)
- [Glacier](namespace-aws-glacier.md)

## MultipartUploader     extends [AbstractUploader](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.AbstractUploader.html)   in package    - [Aws](package-aws.md)

Encapsulates the execution of a multipart upload to Glacier.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html\#toc-constants)

[PART\_MIN\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#constant_PART_MIN_SIZE)
= 1048576

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#method___construct)
: mixed Creates a multipart upload for a Glacier archive.[getStateFromService()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#method_getStateFromService)
: [UploadState](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html)Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html\#constants)

#### PART\_MIN\_SIZE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html\#constant_PART_MIN_SIZE)

`
    public
        mixed
    PART_MIN_SIZE
    = 1048576
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html\#method___construct)

Creates a multipart upload for a Glacier archive.

`
    public
                    __construct(GlacierClient $client, mixed $source[, array<string|int, mixed> $config = [] ]) : mixed`

The valid configuration options are as follows:

- account\_id: (string, default=string('-')) Account ID for the archive
being uploaded, if different from the account making the request.
- archive\_description: (string) Description of the archive.
- before\_complete: (callable) Callback to invoke before the
`CompleteMultipartUpload` operation. The callback should have a
function signature like `function (Aws\Command $command) {...}`.
- before\_initiate: (callable) Callback to invoke before the
`InitiateMultipartUpload` operation. The callback should have a
function signature like `function (Aws\Command $command) {...}`.
- before\_upload: (callable) Callback to invoke before any
`UploadMultipartPart` operations. The callback should have a function
signature like `function (Aws\Command $command) {...}`.
- concurrency: (int, default=int(3)) Maximum number of concurrent
`UploadMultipartPart` operations allowed during the multipart upload.
- part\_size: (int, default=int(1048576)) Part size, in bytes, to use when
doing a multipart upload. This must between 1 MB and 4 GB, and must be
a power of 2 (in megabytes).
- prepare\_data\_source: (callable) Callback to invoke before starting the
multipart upload workflow. The callback should have a function
signature like `function () {...}`.
- state: (Aws\\Multipart\\UploadState) An object that represents the state
of the multipart upload and that is used to resume a previous upload.
When this options is provided, the `account_id`, `key`, and `part_size`
options are ignored.
- vault\_name: (string, required) Vault name to use for the archive being
uploaded.

##### Parameters

$client
: [GlacierClient](class-aws-glacier-glacierclient.md)

Client used for the upload.

$source
: mixed

Source of the data to upload.

$config
: array<string\|int, mixed>
= \[\]

Configuration used to perform the upload.

#### getStateFromService()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html\#method_getStateFromService)

Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

`
    public
            static        getStateFromService(GlacierClient $client, string $vaultName, string $uploadId[, string $accountId = '-' ]) : UploadState`

##### Parameters

$client
: [GlacierClient](class-aws-glacier-glacierclient.md)

GlacierClient object to use.

$vaultName
: string

Vault name for the multipart upload.

$uploadId
: string

Upload ID for the multipart upload.

$accountId
: string
= '-'

Account ID for the multipart upload.

##### Return values

[UploadState](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#toc-methods)
- Constants
  - [PART\_MIN\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#constant_PART_MIN_SIZE)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#method___construct)
  - [getStateFromService()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#method_getStateFromService)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.MultipartUploader.html#top)
