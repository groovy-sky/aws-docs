Menu

- [Aws](namespace-aws.md)
- [Multipart](namespace-aws-multipart.md)

## UploadState        in package    - [Aws](package-aws.md)

Representation of the multipart upload.

This object keeps track of the state of the upload, including the status and
which parts have been uploaded.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#toc-constants)

[COMPLETED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#constant_COMPLETED)
= 2 [CREATED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#constant_CREATED)
= 0 [INITIATED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#constant_INITIATED)
= 1 [PROGRESS\_THRESHOLD\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#constant_PROGRESS_THRESHOLD_SIZE)
= 8

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method___construct)
: mixed [getDisplayProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_getDisplayProgress)
: void Prints progress of upload.[getId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_getId)
: array<string\|int, mixed> Get the upload's ID, which is a tuple of parameters that can uniquely
identify the upload.[getPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_getPartSize)
: int Get the part size.[getUploadedParts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_getUploadedParts)
: array<string\|int, mixed> Returns a sorted list of all the uploaded parts.[hasPartBeenUploaded()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_hasPartBeenUploaded)
: bool Returns whether a part has been uploaded.[isCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_isCompleted)
: bool Determines whether the upload state is in the COMPLETED status.[isInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_isInitiated)
: bool Determines whether the upload state is in the INITIATED status.[markPartAsUploaded()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_markPartAsUploaded)
: mixed Marks a part as being uploaded.[setPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_setPartSize)
: mixed Set the part size.[setProgressThresholds()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_setProgressThresholds)
: array<string\|int, mixed> Sets the 1/8th thresholds array. $totalSize is only sent if
'track\_upload' is true.[setStatus()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_setStatus)
: mixed Set the status of the upload.[setUploadId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_setUploadId)
: mixed Sets the "upload\_id", or 3rd part of the upload's ID. This typically
only needs to be done after initiating an upload.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#constants)

#### COMPLETED  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#constant_COMPLETED)

`
    public
        mixed
    COMPLETED
    = 2
`

#### CREATED  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#constant_CREATED)

`
    public
        mixed
    CREATED
    = 0
`

#### INITIATED  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#constant_INITIATED)

`
    public
        mixed
    INITIATED
    = 1
`

#### PROGRESS\_THRESHOLD\_SIZE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#constant_PROGRESS_THRESHOLD_SIZE)

`
    public
        mixed
    PROGRESS_THRESHOLD_SIZE
    = 8
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $id[, array<string|int, mixed> $config = [] ]) : mixed`

##### Parameters

$id
: array<string\|int, mixed>

Params used to identity the upload.

$config
: array<string\|int, mixed>
= \[\]

#### getDisplayProgress()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_getDisplayProgress)

Prints progress of upload.

`
    public
                    getDisplayProgress( $totalUploaded) : void`

##### Parameters

$totalUploaded
:

numeric Size of upload so far.

#### getId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_getId)

Get the upload's ID, which is a tuple of parameters that can uniquely
identify the upload.

`
    public
                    getId() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getPartSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_getPartSize)

Get the part size.

`
    public
                    getPartSize() : int`

##### Return values

int

#### getUploadedParts()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_getUploadedParts)

Returns a sorted list of all the uploaded parts.

`
    public
                    getUploadedParts() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### hasPartBeenUploaded()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_hasPartBeenUploaded)

Returns whether a part has been uploaded.

`
    public
                    hasPartBeenUploaded(int $partNumber) : bool`

##### Parameters

$partNumber
: int

The part number.

##### Return values

bool

#### isCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_isCompleted)

Determines whether the upload state is in the COMPLETED status.

`
    public
                    isCompleted() : bool`

##### Return values

bool

#### isInitiated()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_isInitiated)

Determines whether the upload state is in the INITIATED status.

`
    public
                    isInitiated() : bool`

##### Return values

bool

#### markPartAsUploaded()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_markPartAsUploaded)

Marks a part as being uploaded.

`
    public
                    markPartAsUploaded(int $partNumber[, array<string|int, mixed> $partData = [] ]) : mixed`

##### Parameters

$partNumber
: int

The part number.

$partData
: array<string\|int, mixed>
= \[\]

Data from the upload operation that needs to be
recalled during the complete operation.

#### setPartSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_setPartSize)

Set the part size.

`
    public
                    setPartSize( $partSize) : mixed`

##### Parameters

$partSize
:

int Size of upload parts.

#### setProgressThresholds()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_setProgressThresholds)

Sets the 1/8th thresholds array. $totalSize is only sent if
'track\_upload' is true.

`
    public
                    setProgressThresholds( $totalSize) : array<string|int, mixed>`

##### Parameters

$totalSize
:

numeric Size of object to upload.

##### Return values

array<string\|int, mixed>

#### setStatus()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_setStatus)

Set the status of the upload.

`
    public
                    setStatus(int $status) : mixed`

##### Parameters

$status
: int

Status is an integer code defined by the constants
CREATED, INITIATED, and COMPLETED on this class.

#### setUploadId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html\#method_setUploadId)

Sets the "upload\_id", or 3rd part of the upload's ID. This typically
only needs to be done after initiating an upload.

`
    public
                    setUploadId(string $key, string $value) : mixed`

##### Parameters

$key
: string

The param key of the upload\_id.

$value
: string

The param value of the upload\_id.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#toc-methods)
- Constants
  - [COMPLETED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#constant_COMPLETED)
  - [CREATED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#constant_CREATED)
  - [INITIATED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#constant_INITIATED)
  - [PROGRESS\_THRESHOLD\_SIZE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#constant_PROGRESS_THRESHOLD_SIZE)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method___construct)
  - [getDisplayProgress()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_getDisplayProgress)
  - [getId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_getId)
  - [getPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_getPartSize)
  - [getUploadedParts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_getUploadedParts)
  - [hasPartBeenUploaded()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_hasPartBeenUploaded)
  - [isCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_isCompleted)
  - [isInitiated()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_isInitiated)
  - [markPartAsUploaded()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_markPartAsUploaded)
  - [setPartSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_setPartSize)
  - [setProgressThresholds()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_setProgressThresholds)
  - [setStatus()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_setStatus)
  - [setUploadId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#method_setUploadId)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html#top)
