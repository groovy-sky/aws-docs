Menu

- [Aws](namespace-aws.md)
- [Multipart](namespace-aws-multipart.md)

## UploadState        in package    - [Aws](package-aws.md)

Representation of the multipart upload.

This object keeps track of the state of the upload, including the status and
which parts have been uploaded.

### Table of Contents  [header link](class-aws-multipart-uploadstate-toc.md)

#### Constants  [header link](class-aws-multipart-uploadstate-toc-constants.md)

[COMPLETED](class-aws-multipart-uploadstate-constant-completed.md)
= 2 [CREATED](class-aws-multipart-uploadstate-constant-created.md)
= 0 [INITIATED](class-aws-multipart-uploadstate-constant-initiated.md)
= 1 [PROGRESS\_THRESHOLD\_SIZE](class-aws-multipart-uploadstate-constant-progress-threshold-size.md)
= 8

#### Methods  [header link](class-aws-multipart-uploadstate-toc-methods.md)

[\_\_construct()](class-aws-multipart-uploadstate-method-construct.md)
: mixed [getDisplayProgress()](class-aws-multipart-uploadstate-method-getdisplayprogress.md)
: void Prints progress of upload.[getId()](class-aws-multipart-uploadstate-method-getid.md)
: array<string\|int, mixed> Get the upload's ID, which is a tuple of parameters that can uniquely
identify the upload.[getPartSize()](class-aws-multipart-uploadstate-method-getpartsize.md)
: int Get the part size.[getUploadedParts()](class-aws-multipart-uploadstate-method-getuploadedparts.md)
: array<string\|int, mixed> Returns a sorted list of all the uploaded parts.[hasPartBeenUploaded()](class-aws-multipart-uploadstate-method-haspartbeenuploaded.md)
: bool Returns whether a part has been uploaded.[isCompleted()](class-aws-multipart-uploadstate-method-iscompleted.md)
: bool Determines whether the upload state is in the COMPLETED status.[isInitiated()](class-aws-multipart-uploadstate-method-isinitiated.md)
: bool Determines whether the upload state is in the INITIATED status.[markPartAsUploaded()](class-aws-multipart-uploadstate-method-markpartasuploaded.md)
: mixed Marks a part as being uploaded.[setPartSize()](class-aws-multipart-uploadstate-method-setpartsize.md)
: mixed Set the part size.[setProgressThresholds()](class-aws-multipart-uploadstate-method-setprogressthresholds.md)
: array<string\|int, mixed> Sets the 1/8th thresholds array. $totalSize is only sent if
'track\_upload' is true.[setStatus()](class-aws-multipart-uploadstate-method-setstatus.md)
: mixed Set the status of the upload.[setUploadId()](class-aws-multipart-uploadstate-method-setuploadid.md)
: mixed Sets the "upload\_id", or 3rd part of the upload's ID. This typically
only needs to be done after initiating an upload.

### Constants  [header link](class-aws-multipart-uploadstate-constants.md)

#### COMPLETED  [header link](class-aws-multipart-uploadstate-constant-completed.md)

`
    public
        mixed
    COMPLETED
    = 2
`

#### CREATED  [header link](class-aws-multipart-uploadstate-constant-created.md)

`
    public
        mixed
    CREATED
    = 0
`

#### INITIATED  [header link](class-aws-multipart-uploadstate-constant-initiated.md)

`
    public
        mixed
    INITIATED
    = 1
`

#### PROGRESS\_THRESHOLD\_SIZE  [header link](class-aws-multipart-uploadstate-constant-progress-threshold-size.md)

`
    public
        mixed
    PROGRESS_THRESHOLD_SIZE
    = 8
`

### Methods  [header link](class-aws-multipart-uploadstate-methods.md)

#### \_\_construct()  [header link](class-aws-multipart-uploadstate-method-construct.md)

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

#### getDisplayProgress()  [header link](class-aws-multipart-uploadstate-method-getdisplayprogress.md)

Prints progress of upload.

`
    public
                    getDisplayProgress( $totalUploaded) : void`

##### Parameters

$totalUploaded
:

numeric Size of upload so far.

#### getId()  [header link](class-aws-multipart-uploadstate-method-getid.md)

Get the upload's ID, which is a tuple of parameters that can uniquely
identify the upload.

`
    public
                    getId() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getPartSize()  [header link](class-aws-multipart-uploadstate-method-getpartsize.md)

Get the part size.

`
    public
                    getPartSize() : int`

##### Return values

int

#### getUploadedParts()  [header link](class-aws-multipart-uploadstate-method-getuploadedparts.md)

Returns a sorted list of all the uploaded parts.

`
    public
                    getUploadedParts() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### hasPartBeenUploaded()  [header link](class-aws-multipart-uploadstate-method-haspartbeenuploaded.md)

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

#### isCompleted()  [header link](class-aws-multipart-uploadstate-method-iscompleted.md)

Determines whether the upload state is in the COMPLETED status.

`
    public
                    isCompleted() : bool`

##### Return values

bool

#### isInitiated()  [header link](class-aws-multipart-uploadstate-method-isinitiated.md)

Determines whether the upload state is in the INITIATED status.

`
    public
                    isInitiated() : bool`

##### Return values

bool

#### markPartAsUploaded()  [header link](class-aws-multipart-uploadstate-method-markpartasuploaded.md)

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

#### setPartSize()  [header link](class-aws-multipart-uploadstate-method-setpartsize.md)

Set the part size.

`
    public
                    setPartSize( $partSize) : mixed`

##### Parameters

$partSize
:

int Size of upload parts.

#### setProgressThresholds()  [header link](class-aws-multipart-uploadstate-method-setprogressthresholds.md)

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

#### setStatus()  [header link](class-aws-multipart-uploadstate-method-setstatus.md)

Set the status of the upload.

`
    public
                    setStatus(int $status) : mixed`

##### Parameters

$status
: int

Status is an integer code defined by the constants
CREATED, INITIATED, and COMPLETED on this class.

#### setUploadId()  [header link](class-aws-multipart-uploadstate-method-setuploadid.md)

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
  - [Constants](class-aws-multipart-uploadstate-toc-constants.md)
  - [Methods](class-aws-multipart-uploadstate-toc-methods.md)
- Constants
  - [COMPLETED](class-aws-multipart-uploadstate-constant-completed.md)
  - [CREATED](class-aws-multipart-uploadstate-constant-created.md)
  - [INITIATED](class-aws-multipart-uploadstate-constant-initiated.md)
  - [PROGRESS\_THRESHOLD\_SIZE](class-aws-multipart-uploadstate-constant-progress-threshold-size.md)
- Methods
  - [\_\_construct()](class-aws-multipart-uploadstate-method-construct.md)
  - [getDisplayProgress()](class-aws-multipart-uploadstate-method-getdisplayprogress.md)
  - [getId()](class-aws-multipart-uploadstate-method-getid.md)
  - [getPartSize()](class-aws-multipart-uploadstate-method-getpartsize.md)
  - [getUploadedParts()](class-aws-multipart-uploadstate-method-getuploadedparts.md)
  - [hasPartBeenUploaded()](class-aws-multipart-uploadstate-method-haspartbeenuploaded.md)
  - [isCompleted()](class-aws-multipart-uploadstate-method-iscompleted.md)
  - [isInitiated()](class-aws-multipart-uploadstate-method-isinitiated.md)
  - [markPartAsUploaded()](class-aws-multipart-uploadstate-method-markpartasuploaded.md)
  - [setPartSize()](class-aws-multipart-uploadstate-method-setpartsize.md)
  - [setProgressThresholds()](class-aws-multipart-uploadstate-method-setprogressthresholds.md)
  - [setStatus()](class-aws-multipart-uploadstate-method-setstatus.md)
  - [setUploadId()](class-aws-multipart-uploadstate-method-setuploadid.md)

[Back To Top](class-aws-multipart-uploadstate-top.md)
