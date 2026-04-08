Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Crypto](namespace-aws-s3-crypto.md)

## InstructionFileMetadataStrategy        in package    - [Aws](package-aws.md)       implements  [MetadataStrategyInterface](class-aws-crypto-metadatastrategyinterface.md)

Stores and reads encryption MetadataEnvelope information in a file on Amazon
S3.

A file with the contents of a MetadataEnvelope will be created or read from
alongside the base file on Amazon S3. The provided client will be used for
reading or writing this object. A specified suffix (default of '.instruction'
will be applied to each of the operations involved with the instruction file.

If there is a failure after an instruction file has been uploaded, it will
not be automatically deleted.

### Table of Contents  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-toc.md)

#### Interfaces  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-toc-interfaces.md)

[MetadataStrategyInterface](class-aws-crypto-metadatastrategyinterface.md)

#### Constants  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-toc-constants.md)

[DEFAULT\_FILE\_SUFFIX](class-aws-s3-crypto-instructionfilemetadatastrategy-constant-default-file-suffix.md)
= '.instruction'

#### Methods  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-toc-methods.md)

[\_\_construct()](class-aws-s3-crypto-instructionfilemetadatastrategy-method-construct.md)
: mixed [load()](class-aws-s3-crypto-instructionfilemetadatastrategy-method-load.md)
: MetadataEnvelopeUses the strategy's client to retrieve the instruction file from S3 and generates
a MetadataEnvelope from its contents.[save()](class-aws-s3-crypto-instructionfilemetadatastrategy-method-save.md)
: array<string\|int, mixed> Places the information in the MetadataEnvelope to a location on S3.

### Constants  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-constants.md)

#### DEFAULT\_FILE\_SUFFIX  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-constant-default-file-suffix.md)

`
    public
        mixed
    DEFAULT_FILE_SUFFIX
    = '.instruction'
`

### Methods  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-methods.md)

#### \_\_construct()  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-method-construct.md)

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

#### load()  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-method-load.md)

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

#### save()  [header link](class-aws-s3-crypto-instructionfilemetadatastrategy-method-save.md)

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
  - [Constants](class-aws-s3-crypto-instructionfilemetadatastrategy-toc-constants.md)
  - [Methods](class-aws-s3-crypto-instructionfilemetadatastrategy-toc-methods.md)
- Constants
  - [DEFAULT\_FILE\_SUFFIX](class-aws-s3-crypto-instructionfilemetadatastrategy-constant-default-file-suffix.md)
- Methods
  - [\_\_construct()](class-aws-s3-crypto-instructionfilemetadatastrategy-method-construct.md)
  - [load()](class-aws-s3-crypto-instructionfilemetadatastrategy-method-load.md)
  - [save()](class-aws-s3-crypto-instructionfilemetadatastrategy-method-save.md)

[Back To Top](class-aws-s3-crypto-instructionfilemetadatastrategy-top.md)
