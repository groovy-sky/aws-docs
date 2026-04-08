Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## MaterialsProvider        in package    - [Aws](package-aws.md)       implements  [MaterialsProviderInterface](class-aws-crypto-materialsproviderinterface.md)

AbstractYes

### Table of Contents  [header link](class-aws-crypto-materialsprovider-toc.md)

#### Interfaces  [header link](class-aws-crypto-materialsprovider-toc-interfaces.md)

[MaterialsProviderInterface](class-aws-crypto-materialsproviderinterface.md)

#### Methods  [header link](class-aws-crypto-materialsprovider-toc-methods.md)

[decryptCek()](class-aws-crypto-materialsprovider-method-decryptcek.md)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key according to the Provider's specifications.[encryptCek()](class-aws-crypto-materialsprovider-method-encryptcek.md)
: string Takes a content encryption key (CEK) and description to return an
encrypted key according to the Provider's specifications.[generateCek()](class-aws-crypto-materialsprovider-method-generatecek.md)
: string [generateIv()](class-aws-crypto-materialsprovider-method-generateiv.md)
: string [getMaterialsDescription()](class-aws-crypto-materialsprovider-method-getmaterialsdescription.md)
: string Returns the material description for this Provider so it can be verified
by encryption mechanisms.[getWrapAlgorithmName()](class-aws-crypto-materialsprovider-method-getwrapalgorithmname.md)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](class-aws-crypto-materialsprovider-method-issupportedkeysize.md)
: bool Returns if the requested size is supported by AES.

### Methods  [header link](class-aws-crypto-materialsprovider-methods.md)

#### decryptCek()  [header link](class-aws-crypto-materialsprovider-method-decryptcek.md)

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

#### encryptCek()  [header link](class-aws-crypto-materialsprovider-method-encryptcek.md)

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

#### generateCek()  [header link](class-aws-crypto-materialsprovider-method-generatecek.md)

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

#### generateIv()  [header link](class-aws-crypto-materialsprovider-method-generateiv.md)

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

#### getMaterialsDescription()  [header link](class-aws-crypto-materialsprovider-method-getmaterialsdescription.md)

Returns the material description for this Provider so it can be verified
by encryption mechanisms.

`
    public
    abstract                getMaterialsDescription() : string`

##### Return values

string

#### getWrapAlgorithmName()  [header link](class-aws-crypto-materialsprovider-method-getwrapalgorithmname.md)

Returns the wrap algorithm name for this Provider.

`
    public
    abstract                getWrapAlgorithmName() : string`

##### Return values

string

#### isSupportedKeySize()  [header link](class-aws-crypto-materialsprovider-method-issupportedkeysize.md)

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
  - [Methods](class-aws-crypto-materialsprovider-toc-methods.md)
- Methods
  - [decryptCek()](class-aws-crypto-materialsprovider-method-decryptcek.md)
  - [encryptCek()](class-aws-crypto-materialsprovider-method-encryptcek.md)
  - [generateCek()](class-aws-crypto-materialsprovider-method-generatecek.md)
  - [generateIv()](class-aws-crypto-materialsprovider-method-generateiv.md)
  - [getMaterialsDescription()](class-aws-crypto-materialsprovider-method-getmaterialsdescription.md)
  - [getWrapAlgorithmName()](class-aws-crypto-materialsprovider-method-getwrapalgorithmname.md)
  - [isSupportedKeySize()](class-aws-crypto-materialsprovider-method-issupportedkeysize.md)

[Back To Top](class-aws-crypto-materialsprovider-top.md)
