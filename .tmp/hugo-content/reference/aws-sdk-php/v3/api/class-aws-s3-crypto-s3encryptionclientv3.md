Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Crypto](namespace-aws-s3-crypto.md)

## S3EncryptionClientV3     extends AbstractCryptoClientV3   in package    - [Aws](package-aws.md)       Uses  [CipherBuilderTrait](class-aws-crypto-cipher-cipherbuildertrait.md), [CryptoParamsTraitV3](class-aws-s3-crypto-cryptoparamstraitv3.md), [DecryptionTraitV3](class-aws-crypto-decryptiontraitv3.md), [EncryptionTraitV3](class-aws-crypto-encryptiontraitv3.md), [UserAgentTrait](class-aws-s3-crypto-useragenttrait.md)

Provides a wrapper for an S3Client that supplies functionality to encrypt
data on putObject\[Async\] calls and decrypt data on getObject\[Async\] calls.

AWS strongly recommends the upgrade to the S3EncryptionClientV3 (over the
S3EncryptionClientV2 or S3EncryptionClient), as it offers updated
data security best practices to our customers who upgrade.
S3EncryptionClientV3 contains breaking changes, so this
will require planning by engineering teams to migrate. New workflows should
just start with S3EncryptionClientV3.

Example write path:

`
use Aws\Crypto\KmsMaterialsProviderV3;
use Aws\S3\Crypto\S3EncryptionClientV3;
use Aws\S3\S3Client;
$encryptionClient = new S3EncryptionClientV3(
new S3Client([
'region' => 'us-west-2',
'version' => 'latest'
])
);
$materialsProvider = new KmsMaterialsProviderV3(
new KmsClient([
'profile' => 'default',
'region' => 'us-east-1',
'version' => 'latest',
],
'your-kms-key-id'
);
$encryptionClient->putObject([
'@MaterialsProvider' => $materialsProvider,
'@CipherOptions' => [
'Cipher' => 'gcm',
'KeySize' => 256,
],
'@CommitmentPolicy' => 'REQUIRE_ENCRYPT_REQUIRE_DECRYPT',
'@KmsEncryptionContext' => ['foo' => 'bar'],
'Bucket' => 'your-bucket',
'Key' => 'your-key',
'Body' => 'your-encrypted-data',
]);
`

Example read call (using objects from previous example):

`
$encryptionClient->getObject([
    '@MaterialsProvider' => $materialsProvider,
    '@CipherOptions' => [
        'Cipher' => 'gcm',
        'KeySize' => 256,
    ],
    '@CommitmentPolicy' => 'REQUIRE_ENCRYPT_REQUIRE_DECRYPT',
    'Bucket' => 'your-bucket',
    'Key' => 'your-key',
]);
`

### Table of Contents  [header link](class-aws-s3-crypto-s3encryptionclientv3-toc.md)

#### Constants  [header link](class-aws-s3-crypto-s3encryptionclientv3-toc-constants.md)

[CRYPTO\_VERSION](class-aws-s3-crypto-s3encryptionclientv3-constant-crypto-version.md)
= '3.0'

#### Methods  [header link](class-aws-s3-crypto-s3encryptionclientv3-toc-methods.md)

[\_\_construct()](class-aws-s3-crypto-s3encryptionclientv3-method-construct.md)
: mixed [getObject()](class-aws-s3-crypto-s3encryptionclientv3-method-getobject.md)
: [Result](class-aws-result.md)Retrieves an object from S3 and decrypts the data in the 'Body' field.[getObjectAsync()](class-aws-s3-crypto-s3encryptionclientv3-method-getobjectasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Promises to retrieve an object from S3 and decrypt the data in the
'Body' field.[putObject()](class-aws-s3-crypto-s3encryptionclientv3-method-putobject.md)
: [Result](class-aws-result.md)Encrypts the data in the 'Body' field of $args and uploads it to the
specified location on S3.[putObjectAsync()](class-aws-s3-crypto-s3encryptionclientv3-method-putobjectasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Encrypts the data in the 'Body' field of $args and promises to upload it
to the specified location on S3.

### Constants  [header link](class-aws-s3-crypto-s3encryptionclientv3-constants.md)

#### CRYPTO\_VERSION  [header link](class-aws-s3-crypto-s3encryptionclientv3-constant-crypto-version.md)

`
    public
        mixed
    CRYPTO_VERSION
    = '3.0'
`

### Methods  [header link](class-aws-s3-crypto-s3encryptionclientv3-methods.md)

#### \_\_construct()  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-construct.md)

`
    public
                    __construct(S3Client $client[, string|null $instructionFileSuffix = null ]) : mixed`

##### Parameters

$client
: [S3Client](class-aws-s3-s3client.md)

The S3Client to be used for true uploading and
retrieving objects from S3 when using the
encryption client.

$instructionFileSuffix
: string\|null
= null

Suffix for a client wide
default when using instruction
files for metadata storage.

#### getObject()  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-getobject.md)

Retrieves an object from S3 and decrypts the data in the 'Body' field.

`
    public
                    getObject(array<string|int, mixed> $args) : Result`

##### Parameters

$args
: array<string\|int, mixed>

Arguments for retrieving an object from S3 via
GetObject and decrypting it.

The required configuration argument is as follows:

- @MaterialsProvider: (MaterialsProviderInterface) Provides Cek, Iv, and Cek
encrypting/decrypting for decryption metadata. May have data loaded
from the MetadataEnvelope upon decryption.
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
- @SecurityProfile: (string) Must be set to 'V3' or 'V3\_AND\_LEGACY'.
  - 'V3' indicates that only objects encrypted with S3EncryptionClientV3
    content encryption and key wrap schemas are able to be decrypted.
  - 'V3\_AND\_LEGACY' indicates that objects encrypted with both
    S3EncryptionClientV3 and older legacy encryption clients are able
    to be decrypted.

The optional configuration arguments are as follows:

- SaveAs: (string) The path to a file on disk to save the decrypted
object data. This will be handled by file\_put\_contents instead of the
Guzzle sink.
- @InstructionFileSuffix: (string\|null) Suffix used when looking for an
instruction file if an InstructionFileMetadataHandler was detected.
- @CipherOptions: (array) Cipher options for encrypting data. A Cipher
is required. Accepts the following options:
\- Aad: (string) Additional authentication data. This option is
passed directly to OpenSSL when using gcm. It is ignored when
using cbc.
- @KmsAllowDecryptWithAnyCmk: (bool) This allows decryption with
KMS materials for any KMS key ID, instead of needing the KMS key ID to
be specified and provided to the decrypt operation. Ignored for non-KMS
materials providers. Defaults to false.

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-getobject-tags.md)

throwsInvalidArgumentException

Thrown when arguments above are not
passed or are passed incorrectly.

##### Return values

[Result](class-aws-result.md)
—

GetObject call result with the 'Body' field
wrapped in a decryption stream with its metadata
information.

#### getObjectAsync()  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-getobjectasync.md)

Promises to retrieve an object from S3 and decrypt the data in the
'Body' field.

`
    public
                    getObjectAsync(array<string|int, mixed> $args) : PromiseInterface`

##### Parameters

$args
: array<string\|int, mixed>

Arguments for retrieving an object from S3 via
GetObject and decrypting it.

The required configuration argument is as follows:

- @MaterialsProvider: (MaterialsProviderInterface) Provides Cek, Iv, and Cek
encrypting/decrypting for decryption metadata. May have data loaded
from the MetadataEnvelope upon decryption.
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
- @SecurityProfile: (string) Must be set to 'V3' or 'V3\_AND\_LEGACY'.
  - 'V3' indicates that only objects encrypted with S3EncryptionClientV3
    content encryption and key wrap schemas are able to be decrypted.
  - 'V3\_AND\_LEGACY' indicates that objects encrypted with both
    S3EncryptionClientV3 and older legacy encryption clients are able
    to be decrypted.

The optional configuration arguments are as follows:

- SaveAs: (string) The path to a file on disk to save the decrypted
object data. This will be handled by file\_put\_contents instead of the
Guzzle sink.

- @MetadataStrategy: (MetadataStrategy\|string\|null) Strategy for reading
MetadataEnvelope information. Defaults to determining based on object
response headers. Can either be a class implementing MetadataStrategy,
a class name of a predefined strategy, or empty/null to default.

- @InstructionFileSuffix: (string) Suffix used when looking for an
instruction file if an InstructionFileMetadataHandler is being used.

- @CipherOptions: (array) Cipher options for decrypting data. A Cipher
is required. Accepts the following options:
\- Aad: (string) Additional authentication data. This option is
passed directly to OpenSSL when using gcm. It is ignored when
using cbc.

- @KmsAllowDecryptWithAnyCmk: (bool) This allows decryption with
KMS materials for any KMS key ID, instead of needing the KMS key ID to
be specified and provided to the decrypt operation. Ignored for non-KMS
materials providers. Defaults to false.

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-getobjectasync-tags.md)

throwsInvalidArgumentException

Thrown when required arguments are not
passed or are passed incorrectly.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### putObject()  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-putobject.md)

Encrypts the data in the 'Body' field of $args and uploads it to the
specified location on S3.

`
    public
                    putObject(array<string|int, mixed> $args) : Result`

##### Parameters

$args
: array<string\|int, mixed>

Arguments for encrypting an object and uploading it
to S3 via PutObject.

The required configuration arguments are as follows:

- @MaterialsProvider: (MaterialsProvider) Provides Cek, Iv, and Cek
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
\- KeySize: (int) 128\|256
See also: MaterialsProvider::$supportedKeySizes
\- Aad: (string) Additional authentication data. This option is
passed directly to OpenSSL when using gcm. Note if you pass in
Aad, the PHP SDK will be able to decrypt the resulting object,
but other AWS SDKs may not be able to do so.
- @KmsEncryptionContext: (array) Only required if using
KmsMaterialsProviderV3. An associative array of key-value
pairs to be added to the encryption context for KMS key encryption. An
empty array may be passed if no additional context is desired.

The optional configuration arguments are as follows:

- @MetadataStrategy: (MetadataStrategy\|string\|null) Strategy for storing
MetadataEnvelope information. Defaults to using a
HeadersMetadataStrategy. Can either be a class implementing
MetadataStrategy, a class name of a predefined strategy, or empty/null
to default.
- @InstructionFileSuffix: (string\|null) Suffix used when writing to an
instruction file if an using an InstructionFileMetadataHandler was
determined.

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-putobject-tags.md)

throwsInvalidArgumentException

Thrown when arguments above are not
passed or are passed incorrectly.

##### Return values

[Result](class-aws-result.md)
—

PutObject call result with the details of uploading
the encrypted file.

#### putObjectAsync()  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-putobjectasync.md)

Encrypts the data in the 'Body' field of $args and promises to upload it
to the specified location on S3.

`
    public
                    putObjectAsync(array<string|int, mixed> $args) : PromiseInterface`

##### Parameters

$args
: array<string\|int, mixed>

Arguments for encrypting an object and uploading it
to S3 via PutObject.

The required configuration arguments are as follows:

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
- @CipherOptions: (array) Cipher options for encrypting data. Only the
Cipher option is required. Accepts the following:
\- Cipher: (string) gcm
See also: AbstractCryptoClientV3::$supportedCiphers
\- KeySize: (int) 256
See also: MaterialsProvider::$supportedKeySizes
\- Aad: (string) Additional authentication data. This option is
passed directly to OpenSSL when using gcm. Note if you pass in
Aad, the PHP SDK will be able to decrypt the resulting object,
but other AWS SDKs may not be able to do so.
- @KmsEncryptionContext: (array) Only required if using
KmsMaterialsProviderV3. An associative array of key-value
pairs to be added to the encryption context for KMS key encryption. An
empty array may be passed if no additional context is desired.

The optional configuration arguments are as follows:

- @MetadataStrategy: (MetadataStrategy\|string\|null) Strategy for storing
MetadataEnvelope information. Defaults to using a
HeadersMetadataStrategy. Can either be a class implementing
MetadataStrategy, a class name of a predefined strategy, or empty/null
to default.
- @InstructionFileSuffix: (string\|null) Suffix used when writing to an
instruction file if using an InstructionFileMetadataHandler.

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclientv3-method-putobjectasync-tags.md)

throwsInvalidArgumentException

Thrown when arguments above are not
passed or are passed incorrectly.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-crypto-s3encryptionclientv3-toc-constants.md)
  - [Methods](class-aws-s3-crypto-s3encryptionclientv3-toc-methods.md)
- Constants
  - [CRYPTO\_VERSION](class-aws-s3-crypto-s3encryptionclientv3-constant-crypto-version.md)
- Methods
  - [\_\_construct()](class-aws-s3-crypto-s3encryptionclientv3-method-construct.md)
  - [getObject()](class-aws-s3-crypto-s3encryptionclientv3-method-getobject.md)
  - [getObjectAsync()](class-aws-s3-crypto-s3encryptionclientv3-method-getobjectasync.md)
  - [putObject()](class-aws-s3-crypto-s3encryptionclientv3-method-putobject.md)
  - [putObjectAsync()](class-aws-s3-crypto-s3encryptionclientv3-method-putobjectasync.md)

[Back To Top](class-aws-s3-crypto-s3encryptionclientv3-top.md)
