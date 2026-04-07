Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## KmsMaterialsProviderV3     extends [MaterialsProviderV3](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html)   in package    - [Aws](package-aws.md)       implements  [MaterialsProviderInterfaceV3](class-aws-crypto-materialsproviderinterfacev3.md)

Uses KMS to supply materials for encrypting and decrypting data. This
V2 implementation should be used with the V2 encryption clients (i.e.

S3EncryptionClientV2).

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#toc-interfaces)

[MaterialsProviderInterfaceV3](class-aws-crypto-materialsproviderinterfacev3.md)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#toc-constants)

[WRAP\_ALGORITHM\_NAME](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#constant_WRAP_ALGORITHM_NAME)
= 'kms+context'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#method___construct)
: mixed [decryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#method_decryptCek)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.[generateCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#method_generateCek)
: array<string\|int, mixed> [generateIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_generateIv)
: string [getWrapAlgorithmName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#method_getWrapAlgorithmName)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_isSupportedKeySize)
: bool Returns if the requested size is supported by AES.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#constants)

#### WRAP\_ALGORITHM\_NAME  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#constant_WRAP_ALGORITHM_NAME)

`
    public
        mixed
    WRAP_ALGORITHM_NAME
    = 'kms+context'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#method___construct)

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

#### decryptCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#method_decryptCek)

Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.

`
    public
                    decryptCek(string $encryptedCek, array<string|int, mixed> $materialDescription, array<string|int, mixed> $options) : string`

##### Parameters

$encryptedCek
: string

Encrypted key to be decrypted by the Provider
for use decrypting other data.

$materialDescription
: array<string\|int, mixed>

Material Description for use in
decrypting the CEK.

$options
: array<string\|int, mixed>

Options for use in decrypting the CEK.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#method_decryptCek\#tags)

inheritDoc

##### Return values

string

#### generateCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#method_generateCek)

`
    public
                    generateCek(mixed $keySize, mixed $context, mixed $options) : array<string|int, mixed>`

##### Parameters

$keySize
: mixed

Length of a cipher key in bits for generating a
random content encryption key (CEK).

$context
: mixed

Context map needed for key encryption

$options
: mixed

Additional options to be used in CEK generation

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#method_generateCek\#tags)

inheritDoc

##### Return values

array<string\|int, mixed>

#### generateIv()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#method_generateIv)

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

#### getWrapAlgorithmName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#method_getWrapAlgorithmName)

Returns the wrap algorithm name for this Provider.

`
    public
                    getWrapAlgorithmName() : string`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html\#method_getWrapAlgorithmName\#tags)

inheritDoc

##### Return values

string

#### isSupportedKeySize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#method_isSupportedKeySize)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#toc-methods)
- Constants
  - [WRAP\_ALGORITHM\_NAME](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#constant_WRAP_ALGORITHM_NAME)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#method___construct)
  - [decryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#method_decryptCek)
  - [generateCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#method_generateCek)
  - [generateIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_generateIv)
  - [getWrapAlgorithmName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#method_getWrapAlgorithmName)
  - [isSupportedKeySize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_isSupportedKeySize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.KmsMaterialsProviderV3.html#top)
