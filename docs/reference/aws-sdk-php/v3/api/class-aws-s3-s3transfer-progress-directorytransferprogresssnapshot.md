Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## DirectoryTransferProgressSnapshot        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method___construct)
: mixed [fromArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_fromArray)
: [DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)[getIdentifier()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getIdentifier)
: string [getReason()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getReason)
: Throwable\|string\|null [getResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getResponse)
: array<string\|int, mixed>\|null [getTotalBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getTotalBytes)
: int [getTotalFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getTotalFiles)
: int [getTransferredBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getTransferredBytes)
: int [getTransferredFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getTransferredFiles)
: int [ratioTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_ratioTransferred)
: float [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_toArray)
: array<string\|int, mixed> [withProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_withProgress)
: [DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)[withResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_withResponse)
: [DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)[withTotals()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_withTotals)
: [DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method___construct)

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

#### fromArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_fromArray)

`
    public
            static        fromArray(array<string|int, mixed> $data) : DirectoryTransferProgressSnapshot`

##### Parameters

$data
: array<string\|int, mixed>

##### Return values

[DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)

#### getIdentifier()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_getIdentifier)

`
    public
                    getIdentifier() : string`

##### Return values

string

#### getReason()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_getReason)

`
    public
                    getReason() : Throwable|string|null`

##### Return values

Throwable\|string\|null

#### getResponse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_getResponse)

`
    public
                    getResponse() : array<string|int, mixed>|null`

##### Return values

array<string\|int, mixed>\|null

#### getTotalBytes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_getTotalBytes)

`
    public
                    getTotalBytes() : int`

##### Return values

int

#### getTotalFiles()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_getTotalFiles)

`
    public
                    getTotalFiles() : int`

##### Return values

int

#### getTransferredBytes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_getTransferredBytes)

`
    public
                    getTransferredBytes() : int`

##### Return values

int

#### getTransferredFiles()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_getTransferredFiles)

`
    public
                    getTransferredFiles() : int`

##### Return values

int

#### ratioTransferred()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_ratioTransferred)

`
    public
                    ratioTransferred() : float`

##### Return values

float

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_toArray)

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### withProgress()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_withProgress)

`
    public
                    withProgress(int $transferredBytes, int $transferredFiles) : DirectoryTransferProgressSnapshot`

##### Parameters

$transferredBytes
: int$transferredFiles
: int

##### Return values

[DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)

#### withResponse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_withResponse)

`
    public
                    withResponse(array<string|int, mixed> $response) : DirectoryTransferProgressSnapshot`

##### Parameters

$response
: array<string\|int, mixed>

##### Return values

[DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)

#### withTotals()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html\#method_withTotals)

`
    public
                    withTotals(int $totalBytes, int $totalFiles) : DirectoryTransferProgressSnapshot`

##### Parameters

$totalBytes
: int$totalFiles
: int

##### Return values

[DirectoryTransferProgressSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method___construct)
  - [fromArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_fromArray)
  - [getIdentifier()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getIdentifier)
  - [getReason()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getReason)
  - [getResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getResponse)
  - [getTotalBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getTotalBytes)
  - [getTotalFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getTotalFiles)
  - [getTransferredBytes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getTransferredBytes)
  - [getTransferredFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_getTransferredFiles)
  - [ratioTransferred()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_ratioTransferred)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_toArray)
  - [withProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_withProgress)
  - [withResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_withResponse)
  - [withTotals()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#method_withTotals)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.DirectoryTransferProgressSnapshot.html#top)
