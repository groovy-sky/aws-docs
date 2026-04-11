Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)
- [Cipher](namespace-aws-crypto-cipher.md)

## Cbc        in package    - [Aws](package-aws.md)       implements  [CipherMethod](class-aws-crypto-cipher-ciphermethod.md)

An implementation of the CBC cipher for use with an AesEncryptingStream or
AesDecrypting stream.

This cipher method is deprecated and in maintenance mode - no new updates will be
released. Please see https://docs.aws.amazon.com/general/latest/gr/aws\_sdk\_cryptography.html
for more information.

##### Tags  [header link](class-aws-crypto-cipher-cbc-tags.md)

deprecated

### Table of Contents  [header link](class-aws-crypto-cipher-cbc-toc.md)

#### Interfaces  [header link](class-aws-crypto-cipher-cbc-toc-interfaces.md)

[CipherMethod](class-aws-crypto-cipher-ciphermethod.md)

#### Constants  [header link](class-aws-crypto-cipher-cbc-toc-constants.md)

[BLOCK\_SIZE](class-aws-crypto-cipher-cbc-constant-block-size.md)
= 16

#### Methods  [header link](class-aws-crypto-cipher-cbc-toc-methods.md)

[\_\_construct()](class-aws-crypto-cipher-cbc-method-construct.md)
: mixed [getAesName()](class-aws-crypto-cipher-cbc-method-getaesname.md)
: string Returns an AES recognizable name, such as 'AES/GCM/NoPadding'.[getCurrentIv()](class-aws-crypto-cipher-cbc-method-getcurrentiv.md)
: string Returns the IV that should be used to initialize the next block in
encrypt or decrypt.[getOpenSslName()](class-aws-crypto-cipher-cbc-method-getopensslname.md)
: string Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.[requiresPadding()](class-aws-crypto-cipher-cbc-method-requirespadding.md)
: bool Indicates whether the cipher method used with this IV requires padding
the final block to make sure the plaintext is evenly divisible by the
block size.[seek()](class-aws-crypto-cipher-cbc-method-seek.md)
: mixed Adjust the return of this::getCurrentIv to reflect a seek performed on
the encryption stream using this IV object.[update()](class-aws-crypto-cipher-cbc-method-update.md)
: mixed Take account of the last cipher text block to adjust the return of
this::getCurrentIv

### Constants  [header link](class-aws-crypto-cipher-cbc-constants.md)

#### BLOCK\_SIZE  [header link](class-aws-crypto-cipher-cbc-constant-block-size.md)

`
    public
        mixed
    BLOCK_SIZE
    = 16
`

### Methods  [header link](class-aws-crypto-cipher-cbc-methods.md)

#### \_\_construct()  [header link](class-aws-crypto-cipher-cbc-method-construct.md)

`
    public
                    __construct(string $iv[, int $keySize = 256 ]) : mixed`

##### Parameters

$iv
: string

Base Initialization Vector for the cipher.

$keySize
: int
= 256

Size of the encryption key, in bits, that will be
used.

##### Tags  [header link](class-aws-crypto-cipher-cbc-method-construct-tags.md)

throwsInvalidArgumentException

Thrown if the passed iv does not match
the iv length required by the cipher.

#### getAesName()  [header link](class-aws-crypto-cipher-cbc-method-getaesname.md)

Returns an AES recognizable name, such as 'AES/GCM/NoPadding'.

`
    public
                    getAesName() : string`

##### Return values

string

#### getCurrentIv()  [header link](class-aws-crypto-cipher-cbc-method-getcurrentiv.md)

Returns the IV that should be used to initialize the next block in
encrypt or decrypt.

`
    public
                    getCurrentIv() : string`

##### Return values

string

#### getOpenSslName()  [header link](class-aws-crypto-cipher-cbc-method-getopensslname.md)

Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.

`
    public
                    getOpenSslName() : string`

##### Return values

string

#### requiresPadding()  [header link](class-aws-crypto-cipher-cbc-method-requirespadding.md)

Indicates whether the cipher method used with this IV requires padding
the final block to make sure the plaintext is evenly divisible by the
block size.

`
    public
                    requiresPadding() : bool`

##### Return values

bool

#### seek()  [header link](class-aws-crypto-cipher-cbc-method-seek.md)

Adjust the return of this::getCurrentIv to reflect a seek performed on
the encryption stream using this IV object.

`
    public
                    seek(mixed $offset[, mixed $whence = SEEK_SET ]) : mixed`

##### Parameters

$offset
: mixed$whence
: mixed
= SEEK\_SET

#### update()  [header link](class-aws-crypto-cipher-cbc-method-update.md)

Take account of the last cipher text block to adjust the return of
this::getCurrentIv

`
    public
                    update(mixed $cipherTextBlock) : mixed`

##### Parameters

$cipherTextBlock
: mixed
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-crypto-cipher-cbc-toc-constants.md)
  - [Methods](class-aws-crypto-cipher-cbc-toc-methods.md)
- Constants
  - [BLOCK\_SIZE](class-aws-crypto-cipher-cbc-constant-block-size.md)
- Methods
  - [\_\_construct()](class-aws-crypto-cipher-cbc-method-construct.md)
  - [getAesName()](class-aws-crypto-cipher-cbc-method-getaesname.md)
  - [getCurrentIv()](class-aws-crypto-cipher-cbc-method-getcurrentiv.md)
  - [getOpenSslName()](class-aws-crypto-cipher-cbc-method-getopensslname.md)
  - [requiresPadding()](class-aws-crypto-cipher-cbc-method-requirespadding.md)
  - [seek()](class-aws-crypto-cipher-cbc-method-seek.md)
  - [update()](class-aws-crypto-cipher-cbc-method-update.md)

[Back To Top](class-aws-crypto-cipher-cbc-top.md)
