Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## TransferProgressSnapshot        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method___construct)
: mixed [fromArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_fromArray)
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html)[getIdentifier()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getIdentifier)
: string [getReason()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getReason)
: Throwable\|string\|null [getResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getResponse)
: array<string\|int, mixed>\|null [getTotalBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getTotalBytes)
: int [getTransferredBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getTransferredBytes)
: int [ratioTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_ratioTransferred)
: float [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_toArray)
: array<string\|int, mixed> [withResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_withResponse)
: [TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html)

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method___construct)

`
    public
                    __construct(string $identifier, int $transferredBytes, int $totalBytes[, array<string|int, mixed>|null $response = null ][, Throwable|string|null $reason = null ]) : mixed`

##### Parameters

$identifier
: string$transferredBytes
: int$totalBytes
: int$response
: array<string\|int, mixed>\|null
= null$reason
: Throwable\|string\|null
= null

#### fromArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_fromArray)

`
    public
            static        fromArray(array<string|int, mixed> $data) : TransferProgressSnapshot`

##### Parameters

$data
: array<string\|int, mixed>

##### Return values

[TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html)

#### getIdentifier()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_getIdentifier)

`
    public
                    getIdentifier() : string`

##### Return values

string

#### getReason()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_getReason)

`
    public
                    getReason() : Throwable|string|null`

##### Return values

Throwable\|string\|null

#### getResponse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_getResponse)

`
    public
                    getResponse() : array<string|int, mixed>|null`

##### Return values

array<string\|int, mixed>\|null

#### getTotalBytes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_getTotalBytes)

`
    public
                    getTotalBytes() : int`

##### Return values

int

#### getTransferredBytes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_getTransferredBytes)

`
    public
                    getTransferredBytes() : int`

##### Return values

int

#### ratioTransferred()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_ratioTransferred)

`
    public
                    ratioTransferred() : float`

##### Return values

float

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_toArray)

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### withResponse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html\#method_withResponse)

`
    public
                    withResponse(array<string|int, mixed> $response) : TransferProgressSnapshot`

##### Parameters

$response
: array<string\|int, mixed>

##### Return values

[TransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method___construct)
  - [fromArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_fromArray)
  - [getIdentifier()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getIdentifier)
  - [getReason()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getReason)
  - [getResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getResponse)
  - [getTotalBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getTotalBytes)
  - [getTransferredBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_getTransferredBytes)
  - [ratioTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_ratioTransferred)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_toArray)
  - [withResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#method_withResponse)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.TransferProgressSnapshot.html#top)
