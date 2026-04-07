Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Crypto](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.crypto.html)

## InstructionFileMetadataStrategy        in package    - [Aws](package-aws.md)       implements  [MetadataStrategyInterface](class-aws-crypto-metadatastrategyinterface.md)

Stores and reads encryption MetadataEnvelope information in a file on Amazon
S3.

A file with the contents of a MetadataEnvelope will be created or read from
alongside the base file on Amazon S3. The provided client will be used for
reading or writing this object. A specified suffix (default of '.instruction'
will be applied to each of the operations involved with the instruction file.

If there is a failure after an instruction file has been uploaded, it will
not be automatically deleted.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#toc-interfaces)

[MetadataStrategyInterface](class-aws-crypto-metadatastrategyinterface.md)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#toc-constants)

[DEFAULT\_FILE\_SUFFIX](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#constant_DEFAULT_FILE_SUFFIX)
= '.instruction'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#method___construct)
: mixed [load()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#method_load)
: MetadataEnvelopeUses the strategy's client to retrieve the instruction file from S3 and generates
a MetadataEnvelope from its contents.[save()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#method_save)
: array<string\|int, mixed> Places the information in the MetadataEnvelope to a location on S3.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#constants)

#### DEFAULT\_FILE\_SUFFIX  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#constant_DEFAULT_FILE_SUFFIX)

`
    public
        mixed
    DEFAULT_FILE_SUFFIX
    = '.instruction'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#method___construct)

`
    public
                    __construct(S3Client $client[, string|null $suffix = null ]) : mixed`

##### Parameters

$client
: [S3Client](class-aws-s3-s3client.md)

Client for use in uploading the instruction file.

$suffix
: string\|null
= null

Optional override suffix for instruction file
object keys.

#### load()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#method_load)

Uses the strategy's client to retrieve the instruction file from S3 and generates
a MetadataEnvelope from its contents.

`
    public
                    load(array<string|int, mixed> $args) : MetadataEnvelope`

##### Parameters

$args
: array<string\|int, mixed>

Arguments from Command and Result that contains
S3 Object information, relevant headers, and command
configuration.

##### Return values

MetadataEnvelope

#### save()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html\#method_save)

Places the information in the MetadataEnvelope to a location on S3.

`
    public
                    save(MetadataEnvelope $envelope, array<string|int, mixed> $args) : array<string|int, mixed>`

##### Parameters

$envelope
: MetadataEnvelope

Encryption data to save according to
the strategy.

$args
: array<string\|int, mixed>

Starting arguments for PutObject, used for saving
extra the instruction file.

##### Return values

array<string\|int, mixed>
—

Updated arguments for PutObject.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#toc-methods)
- Constants
  - [DEFAULT\_FILE\_SUFFIX](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#constant_DEFAULT_FILE_SUFFIX)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#method___construct)
  - [load()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#method_load)
  - [save()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#method_save)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.InstructionFileMetadataStrategy.html#top)
