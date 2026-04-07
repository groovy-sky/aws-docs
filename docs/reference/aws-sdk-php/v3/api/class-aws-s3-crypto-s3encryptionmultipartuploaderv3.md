Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Crypto](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.crypto.html)

## S3EncryptionMultipartUploaderV3     extends [MultipartUploader](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html)   in package    - [Aws](package-aws.md)       Uses  [CipherBuilderTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherBuilderTrait.html), [CryptoParamsTraitV3](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.CryptoParamsTraitV3.html), [EncryptionTraitV3](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.EncryptionTraitV3.html), [UserAgentTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.UserAgentTrait.html)

Encapsulates the execution of a multipart upload of an encrypted object to S3.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html\#toc-constants)

[CRYPTO\_VERSION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#constant_CRYPTO_VERSION)
= '3.0' [PART\_MAX\_NUM](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MAX_NUM)
= 10000 [PART\_MAX\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MAX_SIZE)
= 5368709120 [PART\_MIN\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MIN_SIZE)
= 5242880

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#method___construct)
: mixed Creates a multipart upload for an S3 object after encrypting it.[getStateFromService()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploadingTrait.html#method_getStateFromService)
: [UploadState](class-aws-multipart-uploadstate.md)Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.[isSupportedCipher()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#method_isSupportedCipher)
: bool Returns if the passed cipher name is supported for encryption by the SDK.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html\#constants)

#### CRYPTO\_VERSION  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html\#constant_CRYPTO_VERSION)

`
    public
        mixed
    CRYPTO_VERSION
    = '3.0'
`

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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html\#method___construct)

Creates a multipart upload for an S3 object after encrypting it.

`
    public
                    __construct(S3ClientInterface $client, mixed $source[, array<string|int, mixed> $config = [] ]) : mixed`

The required configuration options are as follows:

- @MaterialsProvider: (MaterialsProviderV3) Provides Cek, Iv, and Cek
encrypting/decrypting for encryption metadata.
- @CommitmentPolicy: (string) Must be set to 'FORBID\_ENCRYPT\_ALLOW\_DECRYPT',
'REQUIRE\_ENCRYPT\_ALLOW\_DECRYPT', or 'REQUIRE\_ENCRYPT\_REQUIRE\_DECRYPT'.
  - 'FORBID\_ENCRYPT\_ALLOW\_DECRYPT' indicates that the client is configured
    to write messages without key commitment and read messages encrypted
    with key commitment or without key commitment.
  - 'REQUIRE\_ENCRYPT\_ALLOW\_DECRYPT' indicates that the client is configured
    to write messages with key commitment and read messages encrypted
    with key commitment or without key commitment.
  - 'REQUIRE\_ENCRYPT\_REQUIRE\_DECRYPT' indicates that the client is configured
    to write messages with key commitment and read messages encrypted
    with key commitment.
- @CipherOptions: (array) Cipher options for encrypting data. A Cipher
is required. Accepts the following options:
\- Cipher: (string) gcm
See also: AbstractCryptoClientV3::$supportedCiphers
\- KeySize: (int) 256
See also: MaterialsProviderV3::$supportedKeySizes
\- Aad: (string) Additional authentication data. This option is
passed directly to OpenSSL when using gcm.
- @KmsEncryptionContext: (array) Only required if using
KmsMaterialsProviderV3. An associative array of key-value
pairs to be added to the encryption context for KMS key encryption. An
empty array may be passed if no additional context is desired.
- bucket: (string) Name of the bucket to which the object is
being uploaded.
- key: (string) Key to use for the object being uploaded.

The optional configuration arguments are as follows:

- @MetadataStrategy: (MetadataStrategy\|string\|null) Strategy for storing
MetadataEnvelope information. Defaults to using a
HeadersMetadataStrategy. Can either be a class implementing
MetadataStrategy, a class name of a predefined strategy, or empty/null
to default.
- @InstructionFileSuffix: (string\|null) Suffix used when writing to an
instruction file if an using an InstructionFileMetadataHandler was
determined.
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
- concurrency: (int, default=int(5)) Maximum number of concurrent
`UploadPart` operations allowed during the multipart upload.
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

#### isSupportedCipher()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html\#method_isSupportedCipher)

Returns if the passed cipher name is supported for encryption by the SDK.

`
    public
            static        isSupportedCipher(string $cipherName) : bool`

##### Parameters

$cipherName
: string

The name of a cipher to verify is registered.

##### Return values

bool
—

If the cipher passed is in our supported list.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#toc-methods)
- Constants
  - [CRYPTO\_VERSION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#constant_CRYPTO_VERSION)
  - [PART\_MAX\_NUM](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MAX_NUM)
  - [PART\_MAX\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MAX_SIZE)
  - [PART\_MIN\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html#constant_PART_MIN_SIZE)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#method___construct)
  - [getStateFromService()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploadingTrait.html#method_getStateFromService)
  - [isSupportedCipher()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#method_isSupportedCipher)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.S3EncryptionMultipartUploaderV3.html#top)
