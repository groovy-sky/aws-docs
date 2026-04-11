Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## AlgorithmConstants        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-crypto-algorithmconstants-toc.md)

#### Constants  [header link](class-aws-crypto-algorithmconstants-toc-constants.md)

[CBC\_MAX\_CONTENT\_LENGTH\_BYTES](class-aws-crypto-algorithmconstants-constant-cbc-max-content-length-bytes.md)
= 1 << 55 The Maximum length of the content that can be encrypted in CBC mode.[CTR\_MAX\_CONTENT\_LENGTH\_BYTES](class-aws-crypto-algorithmconstants-constant-ctr-max-content-length-bytes.md)
= -1 The maximum number of bytes that can be securely encrypted per a single key using AES/CTR.[GCM\_MAX\_CONTENT\_LENGTH\_BITS](class-aws-crypto-algorithmconstants-constant-gcm-max-content-length-bits.md)
= (1 << 39) - 256 The maximum number of 16-byte blocks that can be encrypted with a
GCM cipher. Note the maximum bit-length of the plaintext is (2^39 - 256),
which translates to a maximum byte-length of (2^36 - 32), which in turn
translates to a maximum block-length of (2^32 - 2).

### Constants  [header link](class-aws-crypto-algorithmconstants-constants.md)

#### CBC\_MAX\_CONTENT\_LENGTH\_BYTES  [header link](class-aws-crypto-algorithmconstants-constant-cbc-max-content-length-bytes.md)

The Maximum length of the content that can be encrypted in CBC mode.

`
    public
        mixed
    CBC_MAX_CONTENT_LENGTH_BYTES
    = 1 << 55
`

#### CTR\_MAX\_CONTENT\_LENGTH\_BYTES  [header link](class-aws-crypto-algorithmconstants-constant-ctr-max-content-length-bytes.md)

The maximum number of bytes that can be securely encrypted per a single key using AES/CTR.

`
    public
        mixed
    CTR_MAX_CONTENT_LENGTH_BYTES
    = -1
`

#### GCM\_MAX\_CONTENT\_LENGTH\_BITS  [header link](class-aws-crypto-algorithmconstants-constant-gcm-max-content-length-bits.md)

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

##### Tags  [header link](class-aws-crypto-algorithmconstants-constant-gcm-max-content-length-bits-tags.md)

link[http://csrc.nist.gov/publications/nistpubs/800-38D/SP-800-38D.pdf](http://csrc.nist.gov/publications/nistpubs/800-38D/SP-800-38D.pdf)

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-crypto-algorithmconstants-toc-constants.md)
- Constants
  - [CBC\_MAX\_CONTENT\_LENGTH\_BYTES](class-aws-crypto-algorithmconstants-constant-cbc-max-content-length-bytes.md)
  - [CTR\_MAX\_CONTENT\_LENGTH\_BYTES](class-aws-crypto-algorithmconstants-constant-ctr-max-content-length-bytes.md)
  - [GCM\_MAX\_CONTENT\_LENGTH\_BITS](class-aws-crypto-algorithmconstants-constant-gcm-max-content-length-bits.md)

[Back To Top](class-aws-crypto-algorithmconstants-top.md)
