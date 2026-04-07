Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## KmsMaterialsProvider     extends [MaterialsProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html)   in package    - [Aws](package-aws.md)       implements  [MaterialsProviderInterface](class-aws-crypto-materialsproviderinterface.md)

Uses KMS to supply materials for encrypting and decrypting data.

Legacy implementation that supports legacy S3EncryptionClient and
S3EncryptionMultipartUploader, which use an older encryption workflow. Use
KmsMaterialsProviderV2 with S3EncryptionClientV2 or
S3EncryptionMultipartUploaderV2 if possible.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#tags)

deprecated

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#toc-interfaces)

[MaterialsProviderInterface](class-aws-crypto-materialsproviderinterface.md)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#toc-constants)

[WRAP\_ALGORITHM\_NAME](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#constant_WRAP_ALGORITHM_NAME)
= 'kms'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method___construct)
: mixed [decryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_decryptCek)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key by using KMS' Decrypt API.[encryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_encryptCek)
: string Takes a content encryption key (CEK) and description to return an encrypted
key by using KMS' Encrypt API.[fromDecryptionEnvelope()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_fromDecryptionEnvelope)
: mixed [generateCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_generateCek)
: string [generateIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_generateIv)
: string [getMaterialsDescription()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_getMaterialsDescription)
: array<string\|int, mixed> The KMS key id for use in matching this Provider to its keys,
consistently with other SDKs as 'kms\_cmk\_id'.[getWrapAlgorithmName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_getWrapAlgorithmName)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_isSupportedKeySize)
: bool Returns if the requested size is supported by AES.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#constants)

#### WRAP\_ALGORITHM\_NAME  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#constant_WRAP_ALGORITHM_NAME)

`
    public
        mixed
    WRAP_ALGORITHM_NAME
    = 'kms'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#method___construct)

`
    public
                    __construct(KmsClient $kmsClient[, string $kmsKeyId = null ]) : mixed`

##### Parameters

$kmsClient
: [KmsClient](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Kms.KmsClient.html)

A KMS Client for use encrypting and
decrypting keys.

$kmsKeyId
: string
= null

The private KMS key id to be used for encrypting
and decrypting keys.

#### decryptCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#method_decryptCek)

Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key by using KMS' Decrypt API.

`
    public
                    decryptCek(string $encryptedCek, string $materialDescription) : string`

##### Parameters

$encryptedCek
: string

Encrypted key to be decrypted by the Provider
for use decrypting other data.

$materialDescription
: string

Material Description for use in
encrypting the $cek.

##### Return values

string

#### encryptCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#method_encryptCek)

Takes a content encryption key (CEK) and description to return an encrypted
key by using KMS' Encrypt API.

`
    public
                    encryptCek(string $unencryptedCek, string $materialDescription) : string`

##### Parameters

$unencryptedCek
: string

Key for use in encrypting other data
that itself needs to be encrypted by the
Provider.

$materialDescription
: string

Material Description for use in
encrypting the $cek.

##### Return values

string

#### fromDecryptionEnvelope()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#method_fromDecryptionEnvelope)

`
    public
                    fromDecryptionEnvelope(MetadataEnvelope $envelope) : mixed`

##### Parameters

$envelope
: MetadataEnvelope

#### generateCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#method_generateCek)

`
    public
                    generateCek(string $keySize) : string`

##### Parameters

$keySize
: string

Length of a cipher key in bits for generating a
random content encryption key (CEK).

##### Return values

string

#### generateIv()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#method_generateIv)

`
    public
                    generateIv(string $openSslName) : string`

##### Parameters

$openSslName
: string

Cipher OpenSSL name to use for generating
an initialization vector.

##### Return values

string

#### getMaterialsDescription()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#method_getMaterialsDescription)

The KMS key id for use in matching this Provider to its keys,
consistently with other SDKs as 'kms\_cmk\_id'.

`
    public
                    getMaterialsDescription() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getWrapAlgorithmName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html\#method_getWrapAlgorithmName)

Returns the wrap algorithm name for this Provider.

`
    public
                    getWrapAlgorithmName() : string`

##### Return values

string

#### isSupportedKeySize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#method_isSupportedKeySize)

Returns if the requested size is supported by AES.

`
    public
            static        isSupportedKeySize(int $keySize) : bool`

##### Parameters

$keySize
: int

Size of the requested key in bits.

##### Return values

bool
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#toc-methods)
- Constants
  - [WRAP\_ALGORITHM\_NAME](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#constant_WRAP_ALGORITHM_NAME)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method___construct)
  - [decryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_decryptCek)
  - [encryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_encryptCek)
  - [fromDecryptionEnvelope()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_fromDecryptionEnvelope)
  - [generateCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_generateCek)
  - [generateIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_generateIv)
  - [getMaterialsDescription()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_getMaterialsDescription)
  - [getWrapAlgorithmName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#method_getWrapAlgorithmName)
  - [isSupportedKeySize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_isSupportedKeySize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProvider.html#top)
