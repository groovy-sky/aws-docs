Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## MaterialsProviderV3        in package    - [Aws](package-aws.md)       implements  [MaterialsProviderInterfaceV3](class-aws-crypto-materialsproviderinterfacev3.md)

AbstractYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#toc-interfaces)

[MaterialsProviderInterfaceV3](class-aws-crypto-materialsproviderinterfacev3.md)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#toc-methods)

[decryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_decryptCek)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.[generateCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_generateCek)
: array<string\|int, mixed> [generateIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_generateIv)
: string [getWrapAlgorithmName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_getWrapAlgorithmName)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_isSupportedKeySize)
: bool Returns if the requested size is supported by AES.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#methods)

#### decryptCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#method_decryptCek)

Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.

`
    public
    abstract                decryptCek(string $encryptedCek, array<string|int, mixed> $materialDescription, array<string|int, mixed> $options) : string`

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

##### Return values

string

#### generateCek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#method_generateCek)

`
    public
    abstract                generateCek(string $keySize, array<string|int, mixed> $context, array<string|int, mixed> $options) : array<string|int, mixed>`

##### Parameters

$keySize
: string

Length of a cipher key in bits for generating a
random content encryption key (CEK).

$context
: array<string\|int, mixed>

Context map needed for key encryption

$options
: array<string\|int, mixed>

Additional options to be used in CEK generation

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

#### getWrapAlgorithmName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html\#method_getWrapAlgorithmName)

Returns the wrap algorithm name for this Provider.

`
    public
    abstract                getWrapAlgorithmName() : string`

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#toc-methods)
- Methods
  - [decryptCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_decryptCek)
  - [generateCek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_generateCek)
  - [generateIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_generateIv)
  - [getWrapAlgorithmName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_getWrapAlgorithmName)
  - [isSupportedKeySize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#method_isSupportedKeySize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MaterialsProviderV3.html#top)
