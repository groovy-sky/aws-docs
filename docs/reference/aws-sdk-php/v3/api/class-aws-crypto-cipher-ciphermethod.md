Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)
- [Cipher](namespace-aws-crypto-cipher.md)

## CipherMethod     in    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-crypto-cipher-ciphermethod-toc.md)

#### Methods  [header link](class-aws-crypto-cipher-ciphermethod-toc-methods.md)

[getAesName()](class-aws-crypto-cipher-ciphermethod-method-getaesname.md)
: string Returns an AES recognizable name, such as 'AES/GCM/NoPadding'.[getCurrentIv()](class-aws-crypto-cipher-ciphermethod-method-getcurrentiv.md)
: string Returns the IV that should be used to initialize the next block in
encrypt or decrypt.[getOpenSslName()](class-aws-crypto-cipher-ciphermethod-method-getopensslname.md)
: string Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.[requiresPadding()](class-aws-crypto-cipher-ciphermethod-method-requirespadding.md)
: bool Indicates whether the cipher method used with this IV requires padding
the final block to make sure the plaintext is evenly divisible by the
block size.[seek()](class-aws-crypto-cipher-ciphermethod-method-seek.md)
: mixed Adjust the return of this::getCurrentIv to reflect a seek performed on
the encryption stream using this IV object.[update()](class-aws-crypto-cipher-ciphermethod-method-update.md)
: mixed Take account of the last cipher text block to adjust the return of
this::getCurrentIv

### Methods  [header link](class-aws-crypto-cipher-ciphermethod-methods.md)

#### getAesName()  [header link](class-aws-crypto-cipher-ciphermethod-method-getaesname.md)

Returns an AES recognizable name, such as 'AES/GCM/NoPadding'.

`
    public
                    getAesName() : string`

##### Return values

string

#### getCurrentIv()  [header link](class-aws-crypto-cipher-ciphermethod-method-getcurrentiv.md)

Returns the IV that should be used to initialize the next block in
encrypt or decrypt.

`
    public
                    getCurrentIv() : string`

##### Return values

string

#### getOpenSslName()  [header link](class-aws-crypto-cipher-ciphermethod-method-getopensslname.md)

Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.

`
    public
                    getOpenSslName() : string`

##### Return values

string

#### requiresPadding()  [header link](class-aws-crypto-cipher-ciphermethod-method-requirespadding.md)

Indicates whether the cipher method used with this IV requires padding
the final block to make sure the plaintext is evenly divisible by the
block size.

`
    public
                    requiresPadding() : bool`

##### Return values

bool

#### seek()  [header link](class-aws-crypto-cipher-ciphermethod-method-seek.md)

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

##### Tags  [header link](class-aws-crypto-cipher-ciphermethod-method-seek-tags.md)

throwsLogicException

Thrown if the requested seek is not supported by
this IV implementation. For example, a CBC IV
only supports a full rewind ($offset === 0 &&
$whence === SEEK\_SET)

#### update()  [header link](class-aws-crypto-cipher-ciphermethod-method-update.md)

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
  - [Constants](class-aws-crypto-cipher-ciphermethod-toc-constants.md)
  - [Methods](class-aws-crypto-cipher-ciphermethod-toc-methods.md)
- Methods
  - [getAesName()](class-aws-crypto-cipher-ciphermethod-method-getaesname.md)
  - [getCurrentIv()](class-aws-crypto-cipher-ciphermethod-method-getcurrentiv.md)
  - [getOpenSslName()](class-aws-crypto-cipher-ciphermethod-method-getopensslname.md)
  - [requiresPadding()](class-aws-crypto-cipher-ciphermethod-method-requirespadding.md)
  - [seek()](class-aws-crypto-cipher-ciphermethod-method-seek.md)
  - [update()](class-aws-crypto-cipher-ciphermethod-method-update.md)

[Back To Top](class-aws-crypto-cipher-ciphermethod-top.md)
