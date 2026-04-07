Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## MetadataStrategyInterface     in    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html\#toc-methods)

[load()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html#method_load)
: MetadataEnvelopeGenerates a MetadataEnvelope according to the specific strategy using the
passed arguments.[save()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html#method_save)
: array<string\|int, mixed> Places the information in the MetadataEnvelope to the strategy specific
location. Populates the PutObject arguments with any information
necessary for loading.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html\#methods)

#### load()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html\#method_load)

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

#### save()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html\#method_save)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html#toc-methods)
- Methods
  - [load()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html#method_load)
  - [save()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html#method_save)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.MetadataStrategyInterface.html#top)
