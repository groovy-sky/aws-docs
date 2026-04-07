Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## MaterialsProvider        in package    - [Aws](package-aws.md)       implements  [MaterialsProviderInterface](class-aws-crypto-materialsproviderinterface.md)

AbstractYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#toc-interfaces)

[MaterialsProviderInterface](class-aws-crypto-materialsproviderinterface.md)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#toc-methods)

[decryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_decryptCek)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.[encryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_encryptCek)
: string Takes a content encryption key (CEK) and description to return an
encrypted key according to the Provider's specifications.[generateCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_generateCek)
: string [generateIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_generateIv)
: string [getMaterialsDescription()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_getMaterialsDescription)
: string Returns the material description for this Provider so it can be verified
by encryption mechanisms.[getWrapAlgorithmName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_getWrapAlgorithmName)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_isSupportedKeySize)
: bool Returns if the requested size is supported by AES.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#methods)

#### decryptCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#method_decryptCek)

Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.

`
    public
    abstract                decryptCek(string $encryptedCek, string $materialDescription) : string`

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

#### encryptCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#method_encryptCek)

Takes a content encryption key (CEK) and description to return an
encrypted key according to the Provider's specifications.

`
    public
    abstract                encryptCek(string $unencryptedCek, string $materialDescription) : string`

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

#### getMaterialsDescription()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#method_getMaterialsDescription)

Returns the material description for this Provider so it can be verified
by encryption mechanisms.

`
    public
    abstract                getMaterialsDescription() : string`

##### Return values

string

#### getWrapAlgorithmName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html\#method_getWrapAlgorithmName)

Returns the wrap algorithm name for this Provider.

`
    public
    abstract                getWrapAlgorithmName() : string`

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#toc-methods)
- Methods
  - [decryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_decryptCek)
  - [encryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_encryptCek)
  - [generateCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_generateCek)
  - [generateIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_generateIv)
  - [getMaterialsDescription()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_getMaterialsDescription)
  - [getWrapAlgorithmName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_getWrapAlgorithmName)
  - [isSupportedKeySize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#method_isSupportedKeySize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProvider.html#top)
