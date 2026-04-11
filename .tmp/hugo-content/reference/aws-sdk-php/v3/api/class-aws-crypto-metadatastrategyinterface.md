Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## MetadataStrategyInterface     in    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-crypto-metadatastrategyinterface-toc.md)

#### Methods  [header link](class-aws-crypto-metadatastrategyinterface-toc-methods.md)

[load()](class-aws-crypto-metadatastrategyinterface-method-load.md)
: MetadataEnvelopeGenerates a MetadataEnvelope according to the specific strategy using the
passed arguments.[save()](class-aws-crypto-metadatastrategyinterface-method-save.md)
: array<string\|int, mixed> Places the information in the MetadataEnvelope to the strategy specific
location. Populates the PutObject arguments with any information
necessary for loading.

### Methods  [header link](class-aws-crypto-metadatastrategyinterface-methods.md)

#### load()  [header link](class-aws-crypto-metadatastrategyinterface-method-load.md)

Generates a MetadataEnvelope according to the specific strategy using the
passed arguments.

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

#### save()  [header link](class-aws-crypto-metadatastrategyinterface-method-save.md)

Places the information in the MetadataEnvelope to the strategy specific
location. Populates the PutObject arguments with any information
necessary for loading.

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

Starting arguments for PutObject.

##### Return values

array<string\|int, mixed>
—

Updated arguments for PutObject.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-crypto-metadatastrategyinterface-toc-constants.md)
  - [Methods](class-aws-crypto-metadatastrategyinterface-toc-methods.md)
- Methods
  - [load()](class-aws-crypto-metadatastrategyinterface-method-load.md)
  - [save()](class-aws-crypto-metadatastrategyinterface-method-save.md)

[Back To Top](class-aws-crypto-metadatastrategyinterface-top.md)
