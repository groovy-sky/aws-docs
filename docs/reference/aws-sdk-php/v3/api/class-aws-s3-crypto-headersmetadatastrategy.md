Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Crypto](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.crypto.html)

## HeadersMetadataStrategy        in package    - [Aws](package-aws.md)       implements  [MetadataStrategyInterface](class-aws-crypto-metadatastrategyinterface.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html\#toc-interfaces)

[MetadataStrategyInterface](class-aws-crypto-metadatastrategyinterface.md)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html\#toc-methods)

[load()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html#method_load)
: MetadataEnvelopeGenerates a MetadataEnvelope according to the metadata headers from the
GetObject result.[save()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html#method_save)
: array<string\|int, mixed> Places the information in the MetadataEnvelope in to the metadata for
the PutObject request of the encrypted object.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html\#methods)

#### load()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html\#method_load)

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

#### save()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html\#method_save)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html#toc-methods)
- Methods
  - [load()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html#method_load)
  - [save()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html#method_save)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Crypto.HeadersMetadataStrategy.html#top)
