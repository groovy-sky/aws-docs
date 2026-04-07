Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)
- [Cipher](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.crypto.cipher.html)

## CipherMethod     in    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#toc-methods)

[getAesName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_getAesName)
: string Returns an AES recognizable name, such as 'AES/GCM/NoPadding'.[getCurrentIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_getCurrentIv)
: string Returns the IV that should be used to initialize the next block in
encrypt or decrypt.[getOpenSslName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_getOpenSslName)
: string Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.[requiresPadding()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_requiresPadding)
: bool Indicates whether the cipher method used with this IV requires padding
the final block to make sure the plaintext is evenly divisible by the
block size.[seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_seek)
: mixed Adjust the return of this::getCurrentIv to reflect a seek performed on
the encryption stream using this IV object.[update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_update)
: mixed Take account of the last cipher text block to adjust the return of
this::getCurrentIv

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#methods)

#### getAesName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#method_getAesName)

Returns an AES recognizable name, such as 'AES/GCM/NoPadding'.

`
    public
                    getAesName() : string`

##### Return values

string

#### getCurrentIv()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#method_getCurrentIv)

Returns the IV that should be used to initialize the next block in
encrypt or decrypt.

`
    public
                    getCurrentIv() : string`

##### Return values

string

#### getOpenSslName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#method_getOpenSslName)

Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.

`
    public
                    getOpenSslName() : string`

##### Return values

string

#### requiresPadding()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#method_requiresPadding)

Indicates whether the cipher method used with this IV requires padding
the final block to make sure the plaintext is evenly divisible by the
block size.

`
    public
                    requiresPadding() : bool`

##### Return values

bool

#### seek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#method_seek)

Adjust the return of this::getCurrentIv to reflect a seek performed on
the encryption stream using this IV object.

`
    public
                    seek(int $offset[, int $whence = SEEK_SET ]) : mixed`

##### Parameters

$offset
: int$whence
: int
= SEEK\_SET

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#method_seek\#tags)

throwsLogicException

Thrown if the requested seek is not supported by
this IV implementation. For example, a CBC IV
only supports a full rewind ($offset === 0 &&
$whence === SEEK\_SET)

#### update()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html\#method_update)

Take account of the last cipher text block to adjust the return of
this::getCurrentIv

`
    public
                    update(string $cipherTextBlock) : mixed`

##### Parameters

$cipherTextBlock
: string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#toc-methods)
- Methods
  - [getAesName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_getAesName)
  - [getCurrentIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_getCurrentIv)
  - [getOpenSslName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_getOpenSslName)
  - [requiresPadding()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_requiresPadding)
  - [seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_seek)
  - [update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#method_update)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.Cipher.CipherMethod.html#top)
