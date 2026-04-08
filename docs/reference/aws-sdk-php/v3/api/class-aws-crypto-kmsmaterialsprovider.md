Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## KmsMaterialsProvider     extends [MaterialsProvider](class-aws-crypto-materialsprovider.md)   in package    - [Aws](package-aws.md)       implements  [MaterialsProviderInterface](class-aws-crypto-materialsproviderinterface.md)

Uses KMS to supply materials for encrypting and decrypting data.

Legacy implementation that supports legacy S3EncryptionClient and
S3EncryptionMultipartUploader, which use an older encryption workflow. Use
KmsMaterialsProviderV2 with S3EncryptionClientV2 or
S3EncryptionMultipartUploaderV2 if possible.

##### Tags  [header link](class-aws-crypto-kmsmaterialsprovider-tags.md)

deprecated

### Table of Contents  [header link](class-aws-crypto-kmsmaterialsprovider-toc.md)

#### Interfaces  [header link](class-aws-crypto-kmsmaterialsprovider-toc-interfaces.md)

[MaterialsProviderInterface](class-aws-crypto-materialsproviderinterface.md)

#### Constants  [header link](class-aws-crypto-kmsmaterialsprovider-toc-constants.md)

[WRAP\_ALGORITHM\_NAME](class-aws-crypto-kmsmaterialsprovider-constant-wrap-algorithm-name.md)
= 'kms'

#### Methods  [header link](class-aws-crypto-kmsmaterialsprovider-toc-methods.md)

[\_\_construct()](class-aws-crypto-kmsmaterialsprovider-method-construct.md)
: mixed [decryptCek()](class-aws-crypto-kmsmaterialsprovider-method-decryptcek.md)
: string Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key by using KMS' Decrypt API.[encryptCek()](class-aws-crypto-kmsmaterialsprovider-method-encryptcek.md)
: string Takes a content encryption key (CEK) and description to return an encrypted
key by using KMS' Encrypt API.[fromDecryptionEnvelope()](class-aws-crypto-kmsmaterialsprovider-method-fromdecryptionenvelope.md)
: mixed [generateCek()](class-aws-crypto-materialsprovider-method-generatecek.md)
: string [generateIv()](class-aws-crypto-materialsprovider-method-generateiv.md)
: string [getMaterialsDescription()](class-aws-crypto-kmsmaterialsprovider-method-getmaterialsdescription.md)
: array<string\|int, mixed> The KMS key id for use in matching this Provider to its keys,
consistently with other SDKs as 'kms\_cmk\_id'.[getWrapAlgorithmName()](class-aws-crypto-kmsmaterialsprovider-method-getwrapalgorithmname.md)
: string Returns the wrap algorithm name for this Provider.[isSupportedKeySize()](class-aws-crypto-materialsprovider-method-issupportedkeysize.md)
: bool Returns if the requested size is supported by AES.

### Constants  [header link](class-aws-crypto-kmsmaterialsprovider-constants.md)

#### WRAP\_ALGORITHM\_NAME  [header link](class-aws-crypto-kmsmaterialsprovider-constant-wrap-algorithm-name.md)

`
    public
        mixed
    WRAP_ALGORITHM_NAME
    = 'kms'
`

### Methods  [header link](class-aws-crypto-kmsmaterialsprovider-methods.md)

#### \_\_construct()  [header link](class-aws-crypto-kmsmaterialsprovider-method-construct.md)

`
    public
                    __construct(KmsClient $kmsClient[, string $kmsKeyId = null ]) : mixed`

##### Parameters

$kmsClient
: [KmsClient](class-aws-kms-kmsclient.md)

A KMS Client for use encrypting and
decrypting keys.

$kmsKeyId
: string
= null

The private KMS key id to be used for encrypting
and decrypting keys.

#### decryptCek()  [header link](class-aws-crypto-kmsmaterialsprovider-method-decryptcek.md)

Takes an encrypted content encryption key (CEK) and material description
for use decrypting the key by using KMS' Decrypt API.

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

#### encryptCek()  [header link](class-aws-crypto-kmsmaterialsprovider-method-encryptcek.md)

Takes a content encryption key (CEK) and description to return an encrypted
key by using KMS' Encrypt API.

`
    public
                    encryptCek(string $unencryptedCek, string $materialDescription) : string`

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

#### fromDecryptionEnvelope()  [header link](class-aws-crypto-kmsmaterialsprovider-method-fromdecryptionenvelope.md)

`
    public
                    fromDecryptionEnvelope(MetadataEnvelope $envelope) : mixed`

##### Parameters

$envelope
: MetadataEnvelope

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

#### getMaterialsDescription()  [header link](class-aws-crypto-kmsmaterialsprovider-method-getmaterialsdescription.md)

The KMS key id for use in matching this Provider to its keys,
consistently with other SDKs as 'kms\_cmk\_id'.

`
    public
                    getMaterialsDescription() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getWrapAlgorithmName()  [header link](class-aws-crypto-kmsmaterialsprovider-method-getwrapalgorithmname.md)

Returns the wrap algorithm name for this Provider.

`
    public
                    getWrapAlgorithmName() : string`

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
  - [Constants](class-aws-crypto-kmsmaterialsprovider-toc-constants.md)
  - [Methods](class-aws-crypto-kmsmaterialsprovider-toc-methods.md)
- Constants
  - [WRAP\_ALGORITHM\_NAME](class-aws-crypto-kmsmaterialsprovider-constant-wrap-algorithm-name.md)
- Methods
  - [\_\_construct()](class-aws-crypto-kmsmaterialsprovider-method-construct.md)
  - [decryptCek()](class-aws-crypto-kmsmaterialsprovider-method-decryptcek.md)
  - [encryptCek()](class-aws-crypto-kmsmaterialsprovider-method-encryptcek.md)
  - [fromDecryptionEnvelope()](class-aws-crypto-kmsmaterialsprovider-method-fromdecryptionenvelope.md)
  - [generateCek()](class-aws-crypto-materialsprovider-method-generatecek.md)
  - [generateIv()](class-aws-crypto-materialsprovider-method-generateiv.md)
  - [getMaterialsDescription()](class-aws-crypto-kmsmaterialsprovider-method-getmaterialsdescription.md)
  - [getWrapAlgorithmName()](class-aws-crypto-kmsmaterialsprovider-method-getwrapalgorithmname.md)
  - [isSupportedKeySize()](class-aws-crypto-materialsprovider-method-issupportedkeysize.md)

[Back To Top](class-aws-crypto-kmsmaterialsprovider-top.md)
