Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## MaterialsProviderInterface     in    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-crypto-materialsproviderinterface-toc.md)

#### Methods  [header link](class-aws-crypto-materialsproviderinterface-toc-methods.md)

[decryptCek()](class-aws-crypto-materialsproviderinterface-method-decryptcek.md)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.[generateCek()](class-aws-crypto-materialsproviderinterface-method-generatecek.md)
: string [generateIv()](class-aws-crypto-materialsproviderinterface-method-generateiv.md)
: string [getWrapAlgorithmName()](class-aws-crypto-materialsproviderinterface-method-getwrapalgorithmname.md)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](class-aws-crypto-materialsproviderinterface-method-issupportedkeysize.md)
: bool Returns if the requested size is supported by AES.

### Methods  [header link](class-aws-crypto-materialsproviderinterface-methods.md)

#### decryptCek()  [header link](class-aws-crypto-materialsproviderinterface-method-decryptcek.md)

Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.

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

#### generateCek()  [header link](class-aws-crypto-materialsproviderinterface-method-generatecek.md)

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

#### generateIv()  [header link](class-aws-crypto-materialsproviderinterface-method-generateiv.md)

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

#### getWrapAlgorithmName()  [header link](class-aws-crypto-materialsproviderinterface-method-getwrapalgorithmname.md)

Returns the wrap algorithm name for this Provider.

`
    public
                    getWrapAlgorithmName() : string`

##### Return values

string

#### isSupportedKeySize()  [header link](class-aws-crypto-materialsproviderinterface-method-issupportedkeysize.md)

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
  - [Constants](class-aws-crypto-materialsproviderinterface-toc-constants.md)
  - [Methods](class-aws-crypto-materialsproviderinterface-toc-methods.md)
- Methods
  - [decryptCek()](class-aws-crypto-materialsproviderinterface-method-decryptcek.md)
  - [generateCek()](class-aws-crypto-materialsproviderinterface-method-generatecek.md)
  - [generateIv()](class-aws-crypto-materialsproviderinterface-method-generateiv.md)
  - [getWrapAlgorithmName()](class-aws-crypto-materialsproviderinterface-method-getwrapalgorithmname.md)
  - [isSupportedKeySize()](class-aws-crypto-materialsproviderinterface-method-issupportedkeysize.md)

[Back To Top](class-aws-crypto-materialsproviderinterface-top.md)
