Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## AlgorithmConstants        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html\#toc-constants)

[CBC\_MAX\_CONTENT\_LENGTH\_BYTES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html#constant_CBC_MAX_CONTENT_LENGTH_BYTES)
= 1 << 55 The Maximum length of the content that can be encrypted in CBC mode.[CTR\_MAX\_CONTENT\_LENGTH\_BYTES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html#constant_CTR_MAX_CONTENT_LENGTH_BYTES)
= -1 The maximum number of bytes that can be securely encrypted per a single key using AES/CTR.[GCM\_MAX\_CONTENT\_LENGTH\_BITS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html#constant_GCM_MAX_CONTENT_LENGTH_BITS)
= (1 << 39) - 256 The maximum number of 16-byte blocks that can be encrypted with a
GCM cipher. Note the maximum bit-length of the plaintext is (2^39 - 256),
which translates to a maximum byte-length of (2^36 - 32), which in turn
translates to a maximum block-length of (2^32 - 2).

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html\#constants)

#### CBC\_MAX\_CONTENT\_LENGTH\_BYTES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html\#constant_CBC_MAX_CONTENT_LENGTH_BYTES)

The Maximum length of the content that can be encrypted in CBC mode.

`
    public
        mixed
    CBC_MAX_CONTENT_LENGTH_BYTES
    = 1 << 55
`

#### CTR\_MAX\_CONTENT\_LENGTH\_BYTES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html\#constant_CTR_MAX_CONTENT_LENGTH_BYTES)

The maximum number of bytes that can be securely encrypted per a single key using AES/CTR.

`
    public
        mixed
    CTR_MAX_CONTENT_LENGTH_BYTES
    = -1
`

#### GCM\_MAX\_CONTENT\_LENGTH\_BITS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html\#constant_GCM_MAX_CONTENT_LENGTH_BITS)

The maximum number of 16-byte blocks that can be encrypted with a
GCM cipher. Note the maximum bit-length of the plaintext is (2^39 - 256),
which translates to a maximum byte-length of (2^36 - 32), which in turn
translates to a maximum block-length of (2^32 - 2).

`
    public
        mixed
    GCM_MAX_CONTENT_LENGTH_BITS
    = (1 << 39) - 256
`

Reference: NIST Special Publication 800-38D.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html\#constant_GCM_MAX_CONTENT_LENGTH_BITS\#tags)

link[http://csrc.nist.gov/publications/nistpubs/800-38D/SP-800-38D.pdf](http://csrc.nist.gov/publications/nistpubs/800-38D/SP-800-38D.pdf)

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html#toc-constants)
- Constants
  - [CBC\_MAX\_CONTENT\_LENGTH\_BYTES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html#constant_CBC_MAX_CONTENT_LENGTH_BYTES)
  - [CTR\_MAX\_CONTENT\_LENGTH\_BYTES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html#constant_CTR_MAX_CONTENT_LENGTH_BYTES)
  - [GCM\_MAX\_CONTENT\_LENGTH\_BITS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html#constant_GCM_MAX_CONTENT_LENGTH_BITS)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AlgorithmConstants.html#top)
