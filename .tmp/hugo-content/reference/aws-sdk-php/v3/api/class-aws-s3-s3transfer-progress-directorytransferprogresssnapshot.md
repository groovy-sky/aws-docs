Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## DirectoryTransferProgressSnapshot        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-toc.md)

#### Methods  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-construct.md)
: mixed [fromArray()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-fromarray.md)
: [DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)[getIdentifier()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getidentifier.md)
: string [getReason()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getreason.md)
: Throwable\|string\|null [getResponse()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getresponse.md)
: array<string\|int, mixed>\|null [getTotalBytes()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettotalbytes.md)
: int [getTotalFiles()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettotalfiles.md)
: int [getTransferredBytes()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettransferredbytes.md)
: int [getTransferredFiles()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettransferredfiles.md)
: int [ratioTransferred()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-ratiotransferred.md)
: float [toArray()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-toarray.md)
: array<string\|int, mixed> [withProgress()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withprogress.md)
: [DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)[withResponse()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withresponse.md)
: [DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)[withTotals()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withtotals.md)
: [DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)

### Methods  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-construct.md)

`
    public
                    __construct(string $identifier, int $transferredBytes, int $totalBytes, int $transferredFiles, int $totalFiles[, array<string|int, mixed>|null $response = null ][, Throwable|string|null $reason = null ]) : mixed`

##### Parameters

$identifier
: string$transferredBytes
: int$totalBytes
: int$transferredFiles
: int$totalFiles
: int$response
: array<string\|int, mixed>\|null
= null$reason
: Throwable\|string\|null
= null

#### fromArray()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-fromarray.md)

`
    public
            static        fromArray(array<string|int, mixed> $data) : DirectoryTransferProgressSnapshot`

##### Parameters

$data
: array<string\|int, mixed>

##### Return values

[DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)

#### getIdentifier()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getidentifier.md)

`
    public
                    getIdentifier() : string`

##### Return values

string

#### getReason()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getreason.md)

`
    public
                    getReason() : Throwable|string|null`

##### Return values

Throwable\|string\|null

#### getResponse()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getresponse.md)

`
    public
                    getResponse() : array<string|int, mixed>|null`

##### Return values

array<string\|int, mixed>\|null

#### getTotalBytes()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettotalbytes.md)

`
    public
                    getTotalBytes() : int`

##### Return values

int

#### getTotalFiles()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettotalfiles.md)

`
    public
                    getTotalFiles() : int`

##### Return values

int

#### getTransferredBytes()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettransferredbytes.md)

`
    public
                    getTransferredBytes() : int`

##### Return values

int

#### getTransferredFiles()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettransferredfiles.md)

`
    public
                    getTransferredFiles() : int`

##### Return values

int

#### ratioTransferred()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-ratiotransferred.md)

`
    public
                    ratioTransferred() : float`

##### Return values

float

#### toArray()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-toarray.md)

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### withProgress()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withprogress.md)

`
    public
                    withProgress(int $transferredBytes, int $transferredFiles) : DirectoryTransferProgressSnapshot`

##### Parameters

$transferredBytes
: int$transferredFiles
: int

##### Return values

[DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)

#### withResponse()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withresponse.md)

`
    public
                    withResponse(array<string|int, mixed> $response) : DirectoryTransferProgressSnapshot`

##### Parameters

$response
: array<string\|int, mixed>

##### Return values

[DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)

#### withTotals()  [header link](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withtotals.md)

`
    public
                    withTotals(int $totalBytes, int $totalFiles) : DirectoryTransferProgressSnapshot`

##### Parameters

$totalBytes
: int$totalFiles
: int

##### Return values

[DirectoryTransferProgressSnapshot](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-construct.md)
  - [fromArray()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-fromarray.md)
  - [getIdentifier()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getidentifier.md)
  - [getReason()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getreason.md)
  - [getResponse()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-getresponse.md)
  - [getTotalBytes()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettotalbytes.md)
  - [getTotalFiles()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettotalfiles.md)
  - [getTransferredBytes()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettransferredbytes.md)
  - [getTransferredFiles()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-gettransferredfiles.md)
  - [ratioTransferred()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-ratiotransferred.md)
  - [toArray()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-toarray.md)
  - [withProgress()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withprogress.md)
  - [withResponse()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withresponse.md)
  - [withTotals()](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-method-withtotals.md)

[Back To Top](class-aws-s3-s3transfer-progress-directorytransferprogresssnapshot-top.md)
