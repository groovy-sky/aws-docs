Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## KmsMaterialsProviderV3     extends [MaterialsProviderV3](class-aws-crypto-materialsproviderv3.md)   in package    - [Aws](package-aws.md)       implements  [MaterialsProviderInterfaceV3](class-aws-crypto-materialsproviderinterfacev3.md)

Uses KMS to supply materials for encrypting and decrypting data. This
V2 implementation should be used with the V2 encryption clients (i.e.

S3EncryptionClientV2).

### Table of Contents  [header link](class-aws-crypto-kmsmaterialsproviderv3-toc.md)

#### Interfaces  [header link](class-aws-crypto-kmsmaterialsproviderv3-toc-interfaces.md)

[MaterialsProviderInterfaceV3](class-aws-crypto-materialsproviderinterfacev3.md)

#### Constants  [header link](class-aws-crypto-kmsmaterialsproviderv3-toc-constants.md)

[WRAP\_ALGORITHM\_NAME](class-aws-crypto-kmsmaterialsproviderv3-constant-wrap-algorithm-name.md)
= 'kms+context'

#### Methods  [header link](class-aws-crypto-kmsmaterialsproviderv3-toc-methods.md)

[\_\_construct()](class-aws-crypto-kmsmaterialsproviderv3-method-construct.md)
: mixed [decryptCek()](class-aws-crypto-kmsmaterialsproviderv3-method-decryptcek.md)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.[generateCek()](class-aws-crypto-kmsmaterialsproviderv3-method-generatecek.md)
: array<string\|int, mixed> [generateIv()](class-aws-crypto-materialsproviderv3-method-generateiv.md)
: string [getWrapAlgorithmName()](class-aws-crypto-kmsmaterialsproviderv3-method-getwrapalgorithmname.md)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](class-aws-crypto-materialsproviderv3-method-issupportedkeysize.md)
: bool Returns if the requested size is supported by AES.

### Constants  [header link](class-aws-crypto-kmsmaterialsproviderv3-constants.md)

#### WRAP\_ALGORITHM\_NAME  [header link](class-aws-crypto-kmsmaterialsproviderv3-constant-wrap-algorithm-name.md)

`
    public
        mixed
    WRAP_ALGORITHM_NAME
    = 'kms+context'
`

### Methods  [header link](class-aws-crypto-kmsmaterialsproviderv3-methods.md)

#### \_\_construct()  [header link](class-aws-crypto-kmsmaterialsproviderv3-method-construct.md)

`
    public
                    __construct(KmsClient $kmsClient[, string $kmsKeyId = null ]) : mixed`

##### Parameters

$kmsClient
: [KmsClient](class-aws-kms-kmsclient.md)

A KMS Client for use encrypting and
decrypting keys.

$kmsKeyId
: string
= null

The private KMS key id to be used for encrypting
and decrypting keys.

#### decryptCek()  [header link](class-aws-crypto-kmsmaterialsproviderv3-method-decryptcek.md)

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

##### Tags  [header link](class-aws-crypto-kmsmaterialsproviderv3-method-decryptcek-tags.md)

inheritDoc

##### Return values

string

#### generateCek()  [header link](class-aws-crypto-kmsmaterialsproviderv3-method-generatecek.md)

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

##### Tags  [header link](class-aws-crypto-kmsmaterialsproviderv3-method-generatecek-tags.md)

inheritDoc

##### Return values

array<string\|int, mixed>

#### generateIv()  [header link](class-aws-crypto-materialsproviderv3-method-generateiv.md)

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

#### getWrapAlgorithmName()  [header link](class-aws-crypto-kmsmaterialsproviderv3-method-getwrapalgorithmname.md)

Returns the wrap algorithm name for this Provider.

`
    public
                    getWrapAlgorithmName() : string`

##### Tags  [header link](class-aws-crypto-kmsmaterialsproviderv3-method-getwrapalgorithmname-tags.md)

inheritDoc

##### Return values

string

#### isSupportedKeySize()  [header link](class-aws-crypto-materialsproviderv3-method-issupportedkeysize.md)

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
  - [Constants](class-aws-crypto-kmsmaterialsproviderv3-toc-constants.md)
  - [Methods](class-aws-crypto-kmsmaterialsproviderv3-toc-methods.md)
- Constants
  - [WRAP\_ALGORITHM\_NAME](class-aws-crypto-kmsmaterialsproviderv3-constant-wrap-algorithm-name.md)
- Methods
  - [\_\_construct()](class-aws-crypto-kmsmaterialsproviderv3-method-construct.md)
  - [decryptCek()](class-aws-crypto-kmsmaterialsproviderv3-method-decryptcek.md)
  - [generateCek()](class-aws-crypto-kmsmaterialsproviderv3-method-generatecek.md)
  - [generateIv()](class-aws-crypto-materialsproviderv3-method-generateiv.md)
  - [getWrapAlgorithmName()](class-aws-crypto-kmsmaterialsproviderv3-method-getwrapalgorithmname.md)
  - [isSupportedKeySize()](class-aws-crypto-materialsproviderv3-method-issupportedkeysize.md)

[Back To Top](class-aws-crypto-kmsmaterialsproviderv3-top.md)
