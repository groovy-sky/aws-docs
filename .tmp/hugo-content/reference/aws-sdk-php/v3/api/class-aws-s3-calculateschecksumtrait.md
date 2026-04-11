Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## CalculatesChecksumTrait

### Table of Contents  [header link](class-aws-s3-calculateschecksumtrait-toc.md)

#### Properties  [header link](class-aws-s3-calculateschecksumtrait-toc-properties.md)

[$supportedAlgorithms](class-aws-s3-calculateschecksumtrait-property-supportedalgorithms.md)
: array<string\|int, mixed>

#### Methods  [header link](class-aws-s3-calculateschecksumtrait-toc-methods.md)

[filterChecksum()](class-aws-s3-calculateschecksumtrait-method-filterchecksum.md)
: string\|null Returns the first checksum available, if available.[getEncodedValue()](class-aws-s3-calculateschecksumtrait-method-getencodedvalue.md)
: string

### Properties  [header link](class-aws-s3-calculateschecksumtrait-properties.md)

#### $supportedAlgorithms  [header link](class-aws-s3-calculateschecksumtrait-property-supportedalgorithms.md)

`
    public
    static    array<string|int, mixed>
    $supportedAlgorithms
     = ['crc32c' => true, 'crc32' => true, 'sha256' => true, 'sha1' => true]`

### Methods  [header link](class-aws-s3-calculateschecksumtrait-methods.md)

#### filterChecksum()  [header link](class-aws-s3-calculateschecksumtrait-method-filterchecksum.md)

Returns the first checksum available, if available.

`
    public
            static        filterChecksum(array<string|int, mixed> $parameters) : string|null`

##### Parameters

$parameters
: array<string\|int, mixed>

##### Return values

string\|null

#### getEncodedValue()  [header link](class-aws-s3-calculateschecksumtrait-method-getencodedvalue.md)

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
  - [Properties](class-aws-s3-calculateschecksumtrait-toc-properties.md)
  - [Methods](class-aws-s3-calculateschecksumtrait-toc-methods.md)
- Properties
  - [$supportedAlgorithms](class-aws-s3-calculateschecksumtrait-property-supportedalgorithms.md)
- Methods
  - [filterChecksum()](class-aws-s3-calculateschecksumtrait-method-filterchecksum.md)
  - [getEncodedValue()](class-aws-s3-calculateschecksumtrait-method-getencodedvalue.md)

[Back To Top](class-aws-s3-calculateschecksumtrait-top.md)
