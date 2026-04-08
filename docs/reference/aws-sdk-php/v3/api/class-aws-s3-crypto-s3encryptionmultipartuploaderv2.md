Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Crypto](namespace-aws-s3-crypto.md)

## S3EncryptionMultipartUploaderV2     extends [MultipartUploader](class-aws-s3-multipartuploader.md)   in package    - [Aws](package-aws.md)       Uses  [CipherBuilderTrait](class-aws-crypto-cipher-cipherbuildertrait.md), [CryptoParamsTraitV2](class-aws-s3-crypto-cryptoparamstraitv2.md), [EncryptionTraitV2](class-aws-crypto-encryptiontraitv2.md), [UserAgentTrait](class-aws-s3-crypto-useragenttrait.md)

Encapsulates the execution of a multipart upload of an encrypted object to S3.

Note that for PHP versions of < 7.1, this class uses an AES-GCM polyfill
for encryption since there is no native PHP support. The performance for large
inputs will be a lot slower than for PHP 7.1+, so upgrading older PHP version
environments may be necessary to use this effectively.

### Table of Contents  [header link](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-toc.md)

#### Constants  [header link](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-toc-constants.md)

[CRYPTO\_VERSION](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-constant-crypto-version.md)
= '2.1' [PART\_MAX\_NUM](class-aws-s3-multipartuploader-constant-part-max-num.md)
= 10000 [PART\_MAX\_SIZE](class-aws-s3-multipartuploader-constant-part-max-size.md)
= 5368709120 [PART\_MIN\_SIZE](class-aws-s3-multipartuploader-constant-part-min-size.md)
= 5242880

#### Methods  [header link](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-toc-methods.md)

[\_\_construct()](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-method-construct.md)
: mixed Creates a multipart upload for an S3 object after encrypting it.[getStateFromService()](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)
: [UploadState](class-aws-multipart-uploadstate.md)Creates an UploadState object for a multipart upload by querying the
service for the specified upload's information.[isSupportedCipher()](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-method-issupportedcipher.md)
: bool Returns if the passed cipher name is supported for encryption by the SDK.

### Constants  [header link](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-constants.md)

#### CRYPTO\_VERSION  [header link](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-constant-crypto-version.md)

`
    public
        mixed
    CRYPTO_VERSION
    = '2.1'
`

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

### Methods  [header link](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-methods.md)

#### \_\_construct()  [header link](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-method-construct.md)

Creates a multipart upload for an S3 object after encrypting it.

`
    public
                    __construct(S3ClientInterface $client, mixed $source[, array<string|int, mixed> $config = [] ]) : mixed`

Note that for PHP versions of < 7.1, this class uses an AES-GCM polyfill
for encryption since there is no native PHP support. The performance for
large inputs will be a lot slower than for PHP 7.1+, so upgrading older
PHP version environments may be necessary to use this effectively.

The required configuration options are as follows:

- @MaterialsProvider: (MaterialsProviderV2) Provides Cek, Iv, and Cek
encrypting/decrypting for encryption metadata.
- @CipherOptions: (array) Cipher options for encrypting data. A Cipher
is required. Accepts the following options:
\- Cipher: (string) gcm
See also: AbstractCryptoClientV2::$supportedCiphers
\- KeySize: (int) 128\|256
See also: MaterialsProvider::$supportedKeySizes
\- Aad: (string) Additional authentication data. This option is
passed directly to OpenSSL when using gcm.
- @KmsEncryptionContext: (array) Only required if using
KmsMaterialsProviderV2. An associative array of key-value
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

#### isSupportedCipher()  [header link](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-method-issupportedcipher.md)

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
  - [Constants](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-toc-constants.md)
  - [Methods](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-toc-methods.md)
- Constants
  - [CRYPTO\_VERSION](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-constant-crypto-version.md)
  - [PART\_MAX\_NUM](class-aws-s3-multipartuploader-constant-part-max-num.md)
  - [PART\_MAX\_SIZE](class-aws-s3-multipartuploader-constant-part-max-size.md)
  - [PART\_MIN\_SIZE](class-aws-s3-multipartuploader-constant-part-min-size.md)
- Methods
  - [\_\_construct()](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-method-construct.md)
  - [getStateFromService()](class-aws-s3-multipartuploadingtrait-method-getstatefromservice.md)
  - [isSupportedCipher()](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-method-issupportedcipher.md)

[Back To Top](class-aws-s3-crypto-s3encryptionmultipartuploaderv2-top.md)
