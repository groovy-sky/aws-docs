Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## S3TransferManagerConfig        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-toc-constants.md)

[DEFAULT\_CONCURRENCY](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-concurrency.md)
= 5 [DEFAULT\_MULTIPART\_DOWNLOAD\_TYPE](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-multipart-download-type.md)
= 'part' [DEFAULT\_MULTIPART\_UPLOAD\_THRESHOLD\_BYTES](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-multipart-upload-threshold-bytes.md)
= 16777216 [DEFAULT\_REQUEST\_CHECKSUM\_CALCULATION](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-request-checksum-calculation.md)
= 'when\_supported' [DEFAULT\_RESPONSE\_CHECKSUM\_VALIDATION](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-response-checksum-validation.md)
= 'when\_supported' [DEFAULT\_TARGET\_PART\_SIZE\_BYTES](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-target-part-size-bytes.md)
= 8388608

#### Methods  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-construct.md)
: mixed [fromArray()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-fromarray.md)
: self $config:
\- target\_part\_size\_bytes: (int, default=(8388608 \`8MB\`))
The minimum part size to be used in a multipart upload/download.[getConcurrency()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getconcurrency.md)
: int [getDefaultRegion()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getdefaultregion.md)
: string\|null [getMultipartDownloadType()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getmultipartdownloadtype.md)
: string [getMultipartUploadThresholdBytes()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getmultipartuploadthresholdbytes.md)
: int [getRequestChecksumCalculation()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getrequestchecksumcalculation.md)
: string [getResponseChecksumValidation()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getresponsechecksumvalidation.md)
: string [getTargetPartSizeBytes()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-gettargetpartsizebytes.md)
: int [isTrackProgress()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-istrackprogress.md)
: bool [toArray()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-toarray.md)
: array<string\|int, mixed>

### Constants  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constants.md)

#### DEFAULT\_CONCURRENCY  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-concurrency.md)

`
    public
        mixed
    DEFAULT_CONCURRENCY
    = 5
`

#### DEFAULT\_MULTIPART\_DOWNLOAD\_TYPE  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-multipart-download-type.md)

`
    public
        mixed
    DEFAULT_MULTIPART_DOWNLOAD_TYPE
    = 'part'
`

#### DEFAULT\_MULTIPART\_UPLOAD\_THRESHOLD\_BYTES  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-multipart-upload-threshold-bytes.md)

`
    public
        mixed
    DEFAULT_MULTIPART_UPLOAD_THRESHOLD_BYTES
    = 16777216
`

#### DEFAULT\_REQUEST\_CHECKSUM\_CALCULATION  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-request-checksum-calculation.md)

`
    public
        mixed
    DEFAULT_REQUEST_CHECKSUM_CALCULATION
    = 'when_supported'
`

#### DEFAULT\_RESPONSE\_CHECKSUM\_VALIDATION  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-response-checksum-validation.md)

`
    public
        mixed
    DEFAULT_RESPONSE_CHECKSUM_VALIDATION
    = 'when_supported'
`

#### DEFAULT\_TARGET\_PART\_SIZE\_BYTES  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-target-part-size-bytes.md)

`
    public
        mixed
    DEFAULT_TARGET_PART_SIZE_BYTES
    = 8388608
`

### Methods  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-construct.md)

`
    public
                    __construct(int $targetPartSizeBytes, int $multipartUploadThresholdBytes, string $requestChecksumCalculation, string $responseChecksumValidation, string $multipartDownloadType, int $concurrency, bool $trackProgress, string|null $defaultRegion) : mixed`

##### Parameters

$targetPartSizeBytes
: int$multipartUploadThresholdBytes
: int$requestChecksumCalculation
: string$responseChecksumValidation
: string$multipartDownloadType
: string$concurrency
: int$trackProgress
: bool$defaultRegion
: string\|null

#### fromArray()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-fromarray.md)

$config:
\- target\_part\_size\_bytes: (int, default=(8388608 \`8MB\`))
The minimum part size to be used in a multipart upload/download.

`
    public
            static        fromArray(array<string|int, mixed> $config) : self`

- multipart\_upload\_threshold\_bytes: (int, default=(16777216 `16 MB`))
The threshold to decided whether a multipart upload is needed.
- request\_checksum\_calculation: (string, default= `when_supported`)
To decide whether a checksum validation will be applied to the response.
- response\_checksum\_validation: (string, default= `when_supported`)
- multipart\_download\_type: (string, default='part')
The download type to be used in a multipart download.
- concurrency: (int, default=5)
Maximum number of concurrent operations allowed during a multipart
upload/download.
- track\_progress: (bool, default=false)
To enable progress tracker in a multipart upload/download, and or
a directory upload/download operation.
- default\_region: (string, default="us-east-2")

##### Parameters

$config
: array<string\|int, mixed>

##### Return values

self

#### getConcurrency()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getconcurrency.md)

`
    public
                    getConcurrency() : int`

##### Return values

int

#### getDefaultRegion()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getdefaultregion.md)

`
    public
                    getDefaultRegion() : string|null`

##### Return values

string\|null

#### getMultipartDownloadType()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getmultipartdownloadtype.md)

`
    public
                    getMultipartDownloadType() : string`

##### Return values

string

#### getMultipartUploadThresholdBytes()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getmultipartuploadthresholdbytes.md)

`
    public
                    getMultipartUploadThresholdBytes() : int`

##### Return values

int

#### getRequestChecksumCalculation()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getrequestchecksumcalculation.md)

`
    public
                    getRequestChecksumCalculation() : string`

##### Return values

string

#### getResponseChecksumValidation()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getresponsechecksumvalidation.md)

`
    public
                    getResponseChecksumValidation() : string`

##### Return values

string

#### getTargetPartSizeBytes()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-gettargetpartsizebytes.md)

`
    public
                    getTargetPartSizeBytes() : int`

##### Return values

int

#### isTrackProgress()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-istrackprogress.md)

`
    public
                    isTrackProgress() : bool`

##### Return values

bool

#### toArray()  [header link](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-toarray.md)

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3transfer-models-s3transfermanagerconfig-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-models-s3transfermanagerconfig-toc-methods.md)
- Constants
  - [DEFAULT\_CONCURRENCY](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-concurrency.md)
  - [DEFAULT\_MULTIPART\_DOWNLOAD\_TYPE](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-multipart-download-type.md)
  - [DEFAULT\_MULTIPART\_UPLOAD\_THRESHOLD\_BYTES](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-multipart-upload-threshold-bytes.md)
  - [DEFAULT\_REQUEST\_CHECKSUM\_CALCULATION](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-request-checksum-calculation.md)
  - [DEFAULT\_RESPONSE\_CHECKSUM\_VALIDATION](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-response-checksum-validation.md)
  - [DEFAULT\_TARGET\_PART\_SIZE\_BYTES](class-aws-s3-s3transfer-models-s3transfermanagerconfig-constant-default-target-part-size-bytes.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-construct.md)
  - [fromArray()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-fromarray.md)
  - [getConcurrency()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getconcurrency.md)
  - [getDefaultRegion()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getdefaultregion.md)
  - [getMultipartDownloadType()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getmultipartdownloadtype.md)
  - [getMultipartUploadThresholdBytes()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getmultipartuploadthresholdbytes.md)
  - [getRequestChecksumCalculation()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getrequestchecksumcalculation.md)
  - [getResponseChecksumValidation()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-getresponsechecksumvalidation.md)
  - [getTargetPartSizeBytes()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-gettargetpartsizebytes.md)
  - [isTrackProgress()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-istrackprogress.md)
  - [toArray()](class-aws-s3-s3transfer-models-s3transfermanagerconfig-method-toarray.md)

[Back To Top](class-aws-s3-s3transfer-models-s3transfermanagerconfig-top.md)
