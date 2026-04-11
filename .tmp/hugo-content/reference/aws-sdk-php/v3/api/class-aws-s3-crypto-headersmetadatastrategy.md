Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Crypto](namespace-aws-s3-crypto.md)

## HeadersMetadataStrategy        in package    - [Aws](package-aws.md)       implements  [MetadataStrategyInterface](class-aws-crypto-metadatastrategyinterface.md)

### Table of Contents  [header link](class-aws-s3-crypto-headersmetadatastrategy-toc.md)

#### Interfaces  [header link](class-aws-s3-crypto-headersmetadatastrategy-toc-interfaces.md)

[MetadataStrategyInterface](class-aws-crypto-metadatastrategyinterface.md)

#### Methods  [header link](class-aws-s3-crypto-headersmetadatastrategy-toc-methods.md)

[load()](class-aws-s3-crypto-headersmetadatastrategy-method-load.md)
: MetadataEnvelopeGenerates a MetadataEnvelope according to the metadata headers from the
GetObject result.[save()](class-aws-s3-crypto-headersmetadatastrategy-method-save.md)
: array<string\|int, mixed> Places the information in the MetadataEnvelope in to the metadata for
the PutObject request of the encrypted object.

### Methods  [header link](class-aws-s3-crypto-headersmetadatastrategy-methods.md)

#### load()  [header link](class-aws-s3-crypto-headersmetadatastrategy-method-load.md)

Generates a MetadataEnvelope according to the metadata headers from the
GetObject result.

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

#### save()  [header link](class-aws-s3-crypto-headersmetadatastrategy-method-save.md)

Places the information in the MetadataEnvelope in to the metadata for
the PutObject request of the encrypted object.

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

Arguments for PutObject that can be manipulated to
store strategy related information.

##### Return values

array<string\|int, mixed>
—

Updated arguments for PutObject.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-crypto-headersmetadatastrategy-toc-methods.md)
- Methods
  - [load()](class-aws-s3-crypto-headersmetadatastrategy-method-load.md)
  - [save()](class-aws-s3-crypto-headersmetadatastrategy-method-save.md)

[Back To Top](class-aws-s3-crypto-headersmetadatastrategy-top.md)
