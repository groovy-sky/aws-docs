Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## TransferProgressSnapshot        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-toc.md)

#### Methods  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-construct.md)
: mixed [fromArray()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-fromarray.md)
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md)[getIdentifier()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getidentifier.md)
: string [getReason()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getreason.md)
: Throwable\|string\|null [getResponse()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getresponse.md)
: array<string\|int, mixed>\|null [getTotalBytes()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-gettotalbytes.md)
: int [getTransferredBytes()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-gettransferredbytes.md)
: int [ratioTransferred()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-ratiotransferred.md)
: float [toArray()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-toarray.md)
: array<string\|int, mixed> [withResponse()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-withresponse.md)
: [TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md)

### Methods  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-construct.md)

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

#### fromArray()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-fromarray.md)

`
    public
            static        fromArray(array<string|int, mixed> $data) : TransferProgressSnapshot`

##### Parameters

$data
: array<string\|int, mixed>

##### Return values

[TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md)

#### getIdentifier()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getidentifier.md)

`
    public
                    getIdentifier() : string`

##### Return values

string

#### getReason()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getreason.md)

`
    public
                    getReason() : Throwable|string|null`

##### Return values

Throwable\|string\|null

#### getResponse()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getresponse.md)

`
    public
                    getResponse() : array<string|int, mixed>|null`

##### Return values

array<string\|int, mixed>\|null

#### getTotalBytes()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-gettotalbytes.md)

`
    public
                    getTotalBytes() : int`

##### Return values

int

#### getTransferredBytes()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-gettransferredbytes.md)

`
    public
                    getTransferredBytes() : int`

##### Return values

int

#### ratioTransferred()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-ratiotransferred.md)

`
    public
                    ratioTransferred() : float`

##### Return values

float

#### toArray()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-toarray.md)

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### withResponse()  [header link](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-withresponse.md)

`
    public
                    withResponse(array<string|int, mixed> $response) : TransferProgressSnapshot`

##### Parameters

$response
: array<string\|int, mixed>

##### Return values

[TransferProgressSnapshot](class-aws-s3-s3transfer-progress-transferprogresssnapshot.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3transfer-progress-transferprogresssnapshot-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-construct.md)
  - [fromArray()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-fromarray.md)
  - [getIdentifier()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getidentifier.md)
  - [getReason()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getreason.md)
  - [getResponse()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-getresponse.md)
  - [getTotalBytes()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-gettotalbytes.md)
  - [getTransferredBytes()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-gettransferredbytes.md)
  - [ratioTransferred()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-ratiotransferred.md)
  - [toArray()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-toarray.md)
  - [withResponse()](class-aws-s3-s3transfer-progress-transferprogresssnapshot-method-withresponse.md)

[Back To Top](class-aws-s3-s3transfer-progress-transferprogresssnapshot-top.md)
