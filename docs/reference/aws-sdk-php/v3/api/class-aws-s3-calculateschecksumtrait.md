Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## CalculatesChecksumTrait

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html\#toc)

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html\#toc-properties)

[$supportedAlgorithms](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#property_supportedAlgorithms)
: array<string\|int, mixed>

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html\#toc-methods)

[filterChecksum()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#method_filterChecksum)
: string\|null Returns the first checksum available, if available.[getEncodedValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#method_getEncodedValue)
: string

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html\#properties)

#### $supportedAlgorithms  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html\#property_supportedAlgorithms)

`
    public
    static    array<string|int, mixed>
    $supportedAlgorithms
     = ['crc32c' => true, 'crc32' => true, 'sha256' => true, 'sha1' => true]`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html\#methods)

#### filterChecksum()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html\#method_filterChecksum)

Returns the first checksum available, if available.

`
    public
            static        filterChecksum(array<string|int, mixed> $parameters) : string|null`

##### Parameters

$parameters
: array<string\|int, mixed>

##### Return values

string\|null

#### getEncodedValue()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html\#method_getEncodedValue)

`
    public
            static        getEncodedValue(string $requestedAlgorithm, string $value) : string`

##### Parameters

$requestedAlgorithm
: string

the algorithm to encode with

$value
: string

the value to be encoded

##### Return values

string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#toc-methods)
- Properties
  - [$supportedAlgorithms](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#property_supportedAlgorithms)
- Methods
  - [filterChecksum()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#method_filterChecksum)
  - [getEncodedValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#method_getEncodedValue)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html#top)
