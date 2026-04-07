Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## S3TransferManagerConfig        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#toc-constants)

[DEFAULT\_CONCURRENCY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_CONCURRENCY)
= 5 [DEFAULT\_MULTIPART\_DOWNLOAD\_TYPE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_MULTIPART_DOWNLOAD_TYPE)
= 'part' [DEFAULT\_MULTIPART\_UPLOAD\_THRESHOLD\_BYTES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_MULTIPART_UPLOAD_THRESHOLD_BYTES)
= 16777216 [DEFAULT\_REQUEST\_CHECKSUM\_CALCULATION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_REQUEST_CHECKSUM_CALCULATION)
= 'when\_supported' [DEFAULT\_RESPONSE\_CHECKSUM\_VALIDATION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_RESPONSE_CHECKSUM_VALIDATION)
= 'when\_supported' [DEFAULT\_TARGET\_PART\_SIZE\_BYTES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_TARGET_PART_SIZE_BYTES)
= 8388608

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method___construct)
: mixed [fromArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_fromArray)
: self $config:
\- target\_part\_size\_bytes: (int, default=(8388608 \`8MB\`))
The minimum part size to be used in a multipart upload/download.[getConcurrency()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getConcurrency)
: int [getDefaultRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getDefaultRegion)
: string\|null [getMultipartDownloadType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getMultipartDownloadType)
: string [getMultipartUploadThresholdBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getMultipartUploadThresholdBytes)
: int [getRequestChecksumCalculation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getRequestChecksumCalculation)
: string [getResponseChecksumValidation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getResponseChecksumValidation)
: string [getTargetPartSizeBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getTargetPartSizeBytes)
: int [isTrackProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_isTrackProgress)
: bool [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_toArray)
: array<string\|int, mixed>

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#constants)

#### DEFAULT\_CONCURRENCY  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#constant_DEFAULT_CONCURRENCY)

`
    public
        mixed
    DEFAULT_CONCURRENCY
    = 5
`

#### DEFAULT\_MULTIPART\_DOWNLOAD\_TYPE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#constant_DEFAULT_MULTIPART_DOWNLOAD_TYPE)

`
    public
        mixed
    DEFAULT_MULTIPART_DOWNLOAD_TYPE
    = 'part'
`

#### DEFAULT\_MULTIPART\_UPLOAD\_THRESHOLD\_BYTES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#constant_DEFAULT_MULTIPART_UPLOAD_THRESHOLD_BYTES)

`
    public
        mixed
    DEFAULT_MULTIPART_UPLOAD_THRESHOLD_BYTES
    = 16777216
`

#### DEFAULT\_REQUEST\_CHECKSUM\_CALCULATION  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#constant_DEFAULT_REQUEST_CHECKSUM_CALCULATION)

`
    public
        mixed
    DEFAULT_REQUEST_CHECKSUM_CALCULATION
    = 'when_supported'
`

#### DEFAULT\_RESPONSE\_CHECKSUM\_VALIDATION  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#constant_DEFAULT_RESPONSE_CHECKSUM_VALIDATION)

`
    public
        mixed
    DEFAULT_RESPONSE_CHECKSUM_VALIDATION
    = 'when_supported'
`

#### DEFAULT\_TARGET\_PART\_SIZE\_BYTES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#constant_DEFAULT_TARGET_PART_SIZE_BYTES)

`
    public
        mixed
    DEFAULT_TARGET_PART_SIZE_BYTES
    = 8388608
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method___construct)

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

#### fromArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_fromArray)

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

#### getConcurrency()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_getConcurrency)

`
    public
                    getConcurrency() : int`

##### Return values

int

#### getDefaultRegion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_getDefaultRegion)

`
    public
                    getDefaultRegion() : string|null`

##### Return values

string\|null

#### getMultipartDownloadType()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_getMultipartDownloadType)

`
    public
                    getMultipartDownloadType() : string`

##### Return values

string

#### getMultipartUploadThresholdBytes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_getMultipartUploadThresholdBytes)

`
    public
                    getMultipartUploadThresholdBytes() : int`

##### Return values

int

#### getRequestChecksumCalculation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_getRequestChecksumCalculation)

`
    public
                    getRequestChecksumCalculation() : string`

##### Return values

string

#### getResponseChecksumValidation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_getResponseChecksumValidation)

`
    public
                    getResponseChecksumValidation() : string`

##### Return values

string

#### getTargetPartSizeBytes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_getTargetPartSizeBytes)

`
    public
                    getTargetPartSizeBytes() : int`

##### Return values

int

#### isTrackProgress()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_isTrackProgress)

`
    public
                    isTrackProgress() : bool`

##### Return values

bool

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html\#method_toArray)

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#toc-methods)
- Constants
  - [DEFAULT\_CONCURRENCY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_CONCURRENCY)
  - [DEFAULT\_MULTIPART\_DOWNLOAD\_TYPE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_MULTIPART_DOWNLOAD_TYPE)
  - [DEFAULT\_MULTIPART\_UPLOAD\_THRESHOLD\_BYTES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_MULTIPART_UPLOAD_THRESHOLD_BYTES)
  - [DEFAULT\_REQUEST\_CHECKSUM\_CALCULATION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_REQUEST_CHECKSUM_CALCULATION)
  - [DEFAULT\_RESPONSE\_CHECKSUM\_VALIDATION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_RESPONSE_CHECKSUM_VALIDATION)
  - [DEFAULT\_TARGET\_PART\_SIZE\_BYTES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#constant_DEFAULT_TARGET_PART_SIZE_BYTES)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method___construct)
  - [fromArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_fromArray)
  - [getConcurrency()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getConcurrency)
  - [getDefaultRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getDefaultRegion)
  - [getMultipartDownloadType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getMultipartDownloadType)
  - [getMultipartUploadThresholdBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getMultipartUploadThresholdBytes)
  - [getRequestChecksumCalculation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getRequestChecksumCalculation)
  - [getResponseChecksumValidation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getResponseChecksumValidation)
  - [getTargetPartSizeBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_getTargetPartSizeBytes)
  - [isTrackProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_isTrackProgress)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.S3TransferManagerConfig.html#top)
