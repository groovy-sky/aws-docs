Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Crypto](namespace-aws-s3-crypto.md)

## S3EncryptionClient     extends AbstractCryptoClient   in package    - [Aws](package-aws.md)       Uses  [CipherBuilderTrait](class-aws-crypto-cipher-cipherbuildertrait.md), [CryptoParamsTrait](class-aws-s3-crypto-cryptoparamstrait.md), [DecryptionTrait](class-aws-crypto-decryptiontrait.md), [EncryptionTrait](class-aws-crypto-encryptiontrait.md), [UserAgentTrait](class-aws-s3-crypto-useragenttrait.md)

Provides a wrapper for an S3Client that supplies functionality to encrypt
data on putObject\[Async\] calls and decrypt data on getObject\[Async\] calls.

Legacy implementation using older encryption workflow.

AWS strongly recommends the upgrade to the S3EncryptionClientV2 (over the
S3EncryptionClient), as it offers updated data security best practices to our
customers who upgrade. S3EncryptionClientV2 contains breaking changes, so this
will require planning by engineering teams to migrate. New workflows should
just start with S3EncryptionClientV2.

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclient-tags.md)

deprecated

### Table of Contents  [header link](class-aws-s3-crypto-s3encryptionclient-toc.md)

#### Constants  [header link](class-aws-s3-crypto-s3encryptionclient-toc-constants.md)

[CRYPTO\_VERSION](class-aws-s3-crypto-s3encryptionclient-constant-crypto-version.md)
= '1n'

#### Methods  [header link](class-aws-s3-crypto-s3encryptionclient-toc-methods.md)

[\_\_construct()](class-aws-s3-crypto-s3encryptionclient-method-construct.md)
: mixed [getObject()](class-aws-s3-crypto-s3encryptionclient-method-getobject.md)
: [Result](class-aws-result.md)Retrieves an object from S3 and decrypts the data in the 'Body' field.[getObjectAsync()](class-aws-s3-crypto-s3encryptionclient-method-getobjectasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Promises to retrieve an object from S3 and decrypt the data in the
'Body' field.[putObject()](class-aws-s3-crypto-s3encryptionclient-method-putobject.md)
: [Result](class-aws-result.md)Encrypts the data in the 'Body' field of $args and uploads it to the
specified location on S3.[putObjectAsync()](class-aws-s3-crypto-s3encryptionclient-method-putobjectasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Encrypts the data in the 'Body' field of $args and promises to upload it
to the specified location on S3.

### Constants  [header link](class-aws-s3-crypto-s3encryptionclient-constants.md)

#### CRYPTO\_VERSION  [header link](class-aws-s3-crypto-s3encryptionclient-constant-crypto-version.md)

`
    public
        mixed
    CRYPTO_VERSION
    = '1n'
`

### Methods  [header link](class-aws-s3-crypto-s3encryptionclient-methods.md)

#### \_\_construct()  [header link](class-aws-s3-crypto-s3encryptionclient-method-construct.md)

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

#### getObject()  [header link](class-aws-s3-crypto-s3encryptionclient-method-getobject.md)

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

- @MaterialsProvider: (MaterialsProvider) Provides Cek, Iv, and Cek
encrypting/decrypting for decryption metadata. May have data loaded
from the MetadataEnvelope upon decryption.

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

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclient-method-getobject-tags.md)

throwsInvalidArgumentException

Thrown when arguments above are not
passed or are passed incorrectly.

##### Return values

[Result](class-aws-result.md)
—

GetObject call result with the 'Body' field
wrapped in a decryption stream with its metadata
information.

#### getObjectAsync()  [header link](class-aws-s3-crypto-s3encryptionclient-method-getobjectasync.md)

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

- @MaterialsProvider: (MaterialsProvider) Provides Cek, Iv, and Cek
encrypting/decrypting for decryption metadata. May have data loaded
from the MetadataEnvelope upon decryption.

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

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclient-method-getobjectasync-tags.md)

throwsInvalidArgumentException

Thrown when required arguments are not
passed or are passed incorrectly.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### putObject()  [header link](class-aws-s3-crypto-s3encryptionclient-method-putobject.md)

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
- @CipherOptions: (array) Cipher options for encrypting data. A Cipher
is required. Accepts the following options:
\- Cipher: (string) cbc\|gcm
See also: AbstractCryptoClient::$supportedCiphers. Note that
cbc is deprecated and gcm should be used when possible.
\- KeySize: (int) 128\|192\|256
See also: MaterialsProvider::$supportedKeySizes
\- Aad: (string) Additional authentication data. This option is
passed directly to OpenSSL when using gcm. It is ignored when
using cbc. Note if you pass in Aad for gcm encryption, the
PHP SDK will be able to decrypt the resulting object, but other
AWS SDKs may not be able to do so.

The optional configuration arguments are as follows:

- @MetadataStrategy: (MetadataStrategy\|string\|null) Strategy for storing
MetadataEnvelope information. Defaults to using a
HeadersMetadataStrategy. Can either be a class implementing
MetadataStrategy, a class name of a predefined strategy, or empty/null
to default.
- @InstructionFileSuffix: (string\|null) Suffix used when writing to an
instruction file if an using an InstructionFileMetadataHandler was
determined.

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclient-method-putobject-tags.md)

throwsInvalidArgumentException

Thrown when arguments above are not
passed or are passed incorrectly.

##### Return values

[Result](class-aws-result.md)
—

PutObject call result with the details of uploading
the encrypted file.

#### putObjectAsync()  [header link](class-aws-s3-crypto-s3encryptionclient-method-putobjectasync.md)

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

- @MaterialsProvider: (MaterialsProvider) Provides Cek, Iv, and Cek
encrypting/decrypting for encryption metadata.
- @CipherOptions: (array) Cipher options for encrypting data. Only the
Cipher option is required. Accepts the following:
\- Cipher: (string) cbc\|gcm
See also: AbstractCryptoClient::$supportedCiphers. Note that
cbc is deprecated and gcm should be used when possible.
\- KeySize: (int) 128\|192\|256
See also: MaterialsProvider::$supportedKeySizes
\- Aad: (string) Additional authentication data. This option is
passed directly to OpenSSL when using gcm. It is ignored when
using cbc. Note if you pass in Aad for gcm encryption, the
PHP SDK will be able to decrypt the resulting object, but other
AWS SDKs may not be able to do so.

The optional configuration arguments are as follows:

- @MetadataStrategy: (MetadataStrategy\|string\|null) Strategy for storing
MetadataEnvelope information. Defaults to using a
HeadersMetadataStrategy. Can either be a class implementing
MetadataStrategy, a class name of a predefined strategy, or empty/null
to default.
- @InstructionFileSuffix: (string\|null) Suffix used when writing to an
instruction file if using an InstructionFileMetadataHandler.

##### Tags  [header link](class-aws-s3-crypto-s3encryptionclient-method-putobjectasync-tags.md)

throwsInvalidArgumentException

Thrown when arguments above are not
passed or are passed incorrectly.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-crypto-s3encryptionclient-toc-constants.md)
  - [Methods](class-aws-s3-crypto-s3encryptionclient-toc-methods.md)
- Constants
  - [CRYPTO\_VERSION](class-aws-s3-crypto-s3encryptionclient-constant-crypto-version.md)
- Methods
  - [\_\_construct()](class-aws-s3-crypto-s3encryptionclient-method-construct.md)
  - [getObject()](class-aws-s3-crypto-s3encryptionclient-method-getobject.md)
  - [getObjectAsync()](class-aws-s3-crypto-s3encryptionclient-method-getobjectasync.md)
  - [putObject()](class-aws-s3-crypto-s3encryptionclient-method-putobject.md)
  - [putObjectAsync()](class-aws-s3-crypto-s3encryptionclient-method-putobjectasync.md)

[Back To Top](class-aws-s3-crypto-s3encryptionclient-top.md)
