Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## MaterialsProviderInterfaceV2     in    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-crypto-materialsproviderinterfacev2-toc.md)

#### Methods  [header link](class-aws-crypto-materialsproviderinterfacev2-toc-methods.md)

[decryptCek()](class-aws-crypto-materialsproviderinterfacev2-method-decryptcek.md)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.[generateCek()](class-aws-crypto-materialsproviderinterfacev2-method-generatecek.md)
: array<string\|int, mixed> [generateIv()](class-aws-crypto-materialsproviderinterfacev2-method-generateiv.md)
: string [getWrapAlgorithmName()](class-aws-crypto-materialsproviderinterfacev2-method-getwrapalgorithmname.md)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](class-aws-crypto-materialsproviderinterfacev2-method-issupportedkeysize.md)
: bool Returns if the requested size is supported by AES.

### Methods  [header link](class-aws-crypto-materialsproviderinterfacev2-methods.md)

#### decryptCek()  [header link](class-aws-crypto-materialsproviderinterfacev2-method-decryptcek.md)

Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.

`
    public
                    decryptCek(string $encryptedCek, string $materialDescription, array<string|int, mixed> $options) : string`

##### Parameters

$encryptedCek
: string

Encrypted key to be decrypted by the Provider
for use decrypting other data.

$materialDescription
: string

Material Description for use in
decrypting the CEK.

$options
: array<string\|int, mixed>

Options for use in decrypting the CEK.

##### Return values

string

#### generateCek()  [header link](class-aws-crypto-materialsproviderinterfacev2-method-generatecek.md)

`
    public
                    generateCek(string $keySize, array<string|int, mixed> $context, array<string|int, mixed> $options) : array<string|int, mixed>`

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

#### generateIv()  [header link](class-aws-crypto-materialsproviderinterfacev2-method-generateiv.md)

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

#### getWrapAlgorithmName()  [header link](class-aws-crypto-materialsproviderinterfacev2-method-getwrapalgorithmname.md)

Returns the wrap algorithm name for this Provider.

`
    public
                    getWrapAlgorithmName() : string`

##### Return values

string

#### isSupportedKeySize()  [header link](class-aws-crypto-materialsproviderinterfacev2-method-issupportedkeysize.md)

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
  - [Constants](class-aws-crypto-materialsproviderinterfacev2-toc-constants.md)
  - [Methods](class-aws-crypto-materialsproviderinterfacev2-toc-methods.md)
- Methods
  - [decryptCek()](class-aws-crypto-materialsproviderinterfacev2-method-decryptcek.md)
  - [generateCek()](class-aws-crypto-materialsproviderinterfacev2-method-generatecek.md)
  - [generateIv()](class-aws-crypto-materialsproviderinterfacev2-method-generateiv.md)
  - [getWrapAlgorithmName()](class-aws-crypto-materialsproviderinterfacev2-method-getwrapalgorithmname.md)
  - [isSupportedKeySize()](class-aws-crypto-materialsproviderinterfacev2-method-issupportedkeysize.md)

[Back To Top](class-aws-crypto-materialsproviderinterfacev2-top.md)
