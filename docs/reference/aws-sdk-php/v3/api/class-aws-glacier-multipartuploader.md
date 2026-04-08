Menu

- [Aws](namespace-aws.md)
- [Glacier](namespace-aws-glacier.md)

## MultipartUploader     extends [AbstractUploader](class-aws-multipart-abstractuploader.md)   in package    - [Aws](package-aws.md)

Encapsulates the execution of a multipart upload to Glacier.

### Table of Contents  [header link](class-aws-glacier-multipartuploader-toc.md)

#### Constants  [header link](class-aws-glacier-multipartuploader-toc-constants.md)

[PART\_MIN\_SIZE](class-aws-glacier-multipartuploader-constant-part-min-size.md)
= 1048576

#### Methods  [header link](class-aws-glacier-multipartuploader-toc-methods.md)

[\_\_construct()](class-aws-glacier-multipartuploader-method-construct.md)
: mixed Creates a multipart upload for a Glacier archive.[getStateFromService()](class-aws-glacier-multipartuploader-method-getstatefromservice.md)
: [UploadState](class-aws-multipart-uploadstate.md)Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.

### Constants  [header link](class-aws-glacier-multipartuploader-constants.md)

#### PART\_MIN\_SIZE  [header link](class-aws-glacier-multipartuploader-constant-part-min-size.md)

`
    public
        mixed
    PART_MIN_SIZE
    = 1048576
`

### Methods  [header link](class-aws-glacier-multipartuploader-methods.md)

#### \_\_construct()  [header link](class-aws-glacier-multipartuploader-method-construct.md)

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

#### getStateFromService()  [header link](class-aws-glacier-multipartuploader-method-getstatefromservice.md)

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

[UploadState](class-aws-multipart-uploadstate.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-glacier-multipartuploader-toc-constants.md)
  - [Methods](class-aws-glacier-multipartuploader-toc-methods.md)
- Constants
  - [PART\_MIN\_SIZE](class-aws-glacier-multipartuploader-constant-part-min-size.md)
- Methods
  - [\_\_construct()](class-aws-glacier-multipartuploader-method-construct.md)
  - [getStateFromService()](class-aws-glacier-multipartuploader-method-getstatefromservice.md)

[Back To Top](class-aws-glacier-multipartuploader-top.md)
