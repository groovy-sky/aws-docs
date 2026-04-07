Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)
- [Cipher](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.crypto.cipher.html)

## Cbc        in package    - [Aws](package-aws.md)       implements  [CipherMethod](class-aws-crypto-cipher-ciphermethod.md)

An implementation of the CBC cipher for use with an AesEncryptingStream or
AesDecrypting stream.

This cipher method is deprecated and in maintenance mode - no new updates will be
released. Please see https://docs.aws.amazon.com/general/latest/gr/aws\_sdk\_cryptography.html
for more information.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#tags)

deprecated

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#toc-interfaces)

[CipherMethod](class-aws-crypto-cipher-ciphermethod.md)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#toc-constants)

[BLOCK\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#constant_BLOCK_SIZE)
= 16

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method___construct)
: mixed [getAesName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_getAesName)
: string Returns an AES recognizable name, such as 'AES/GCM/NoPadding'.[getCurrentIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_getCurrentIv)
: string Returns the IV that should be used to initialize the next block in
encrypt or decrypt.[getOpenSslName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_getOpenSslName)
: string Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.[requiresPadding()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_requiresPadding)
: bool Indicates whether the cipher method used with this IV requires padding
the final block to make sure the plaintext is evenly divisible by the
block size.[seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_seek)
: mixed Adjust the return of this::getCurrentIv to reflect a seek performed on
the encryption stream using this IV object.[update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_update)
: mixed Take account of the last cipher text block to adjust the return of
this::getCurrentIv

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#constants)

#### BLOCK\_SIZE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#constant_BLOCK_SIZE)

`
    public
        mixed
    BLOCK_SIZE
    = 16
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#method___construct)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#method___construct\#tags)

throwsInvalidArgumentException

Thrown if the passed iv does not match
the iv length required by the cipher.

#### getAesName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#method_getAesName)

Returns an AES recognizable name, such as 'AES/GCM/NoPadding'.

`
    public
                    getAesName() : string`

##### Return values

string

#### getCurrentIv()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#method_getCurrentIv)

Returns the IV that should be used to initialize the next block in
encrypt or decrypt.

`
    public
                    getCurrentIv() : string`

##### Return values

string

#### getOpenSslName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#method_getOpenSslName)

Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.

`
    public
                    getOpenSslName() : string`

##### Return values

string

#### requiresPadding()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#method_requiresPadding)

Indicates whether the cipher method used with this IV requires padding
the final block to make sure the plaintext is evenly divisible by the
block size.

`
    public
                    requiresPadding() : bool`

##### Return values

bool

#### seek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#method_seek)

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

#### update()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html\#method_update)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#toc-methods)
- Constants
  - [BLOCK\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#constant_BLOCK_SIZE)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method___construct)
  - [getAesName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_getAesName)
  - [getCurrentIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_getCurrentIv)
  - [getOpenSslName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_getOpenSslName)
  - [requiresPadding()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_requiresPadding)
  - [seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_seek)
  - [update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#method_update)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.Cbc.html#top)
