Menu

- [Aws](namespace-aws.md)
- [Glacier](namespace-aws-glacier.md)

## TreeHash        in package    - [Aws](package-aws.md)       implements  [HashInterface](class-aws-hashinterface.md)

Encapsulates the creation of a tree hash from streamed data

### Table of Contents  [header link](class-aws-glacier-treehash-toc.md)

#### Interfaces  [header link](class-aws-glacier-treehash-toc-interfaces.md)

[HashInterface](class-aws-hashinterface.md)Interface that allows implementing various incremental hashes.

#### Constants  [header link](class-aws-glacier-treehash-toc-constants.md)

[EMPTY\_HASH](class-aws-glacier-treehash-constant-empty-hash.md)
= 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855' [MB](class-aws-glacier-treehash-constant-mb.md)
= 1048576

#### Methods  [header link](class-aws-glacier-treehash-toc-methods.md)

[\_\_construct()](class-aws-glacier-treehash-method-construct.md)
: mixed [addChecksum()](class-aws-glacier-treehash-method-addchecksum.md)
: self Add a checksum to the tree hash directly[complete()](class-aws-glacier-treehash-method-complete.md)
: string Finalizes the incremental hash and returns the resulting digest.[reset()](class-aws-glacier-treehash-method-reset.md)
: mixed Removes all data from the hash, effectively starting a new hash.[update()](class-aws-glacier-treehash-method-update.md)
: mixed Adds data to the hash.

### Constants  [header link](class-aws-glacier-treehash-constants.md)

#### EMPTY\_HASH  [header link](class-aws-glacier-treehash-constant-empty-hash.md)

`
    public
        mixed
    EMPTY_HASH
    = 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'
`

#### MB  [header link](class-aws-glacier-treehash-constant-mb.md)

`
    public
        mixed
    MB
    = 1048576
`

### Methods  [header link](class-aws-glacier-treehash-methods.md)

#### \_\_construct()  [header link](class-aws-glacier-treehash-method-construct.md)

`
    public
                    __construct([mixed $algorithm = 'sha256' ]) : mixed`

##### Parameters

$algorithm
: mixed
= 'sha256'

#### addChecksum()  [header link](class-aws-glacier-treehash-method-addchecksum.md)

Add a checksum to the tree hash directly

`
    public
                    addChecksum(string $checksum[, bool $inBinaryForm = false ]) : self`

##### Parameters

$checksum
: string

The checksum to add

$inBinaryForm
: bool
= false

TRUE if checksum is in binary form

##### Tags  [header link](class-aws-glacier-treehash-method-addchecksum-tags.md)

throwsLogicException

if the root tree hash is already calculated

##### Return values

self

#### complete()  [header link](class-aws-glacier-treehash-method-complete.md)

Finalizes the incremental hash and returns the resulting digest.

`
    public
                    complete() : string`

##### Return values

string

#### reset()  [header link](class-aws-glacier-treehash-method-reset.md)

Removes all data from the hash, effectively starting a new hash.

`
    public
                    reset() : mixed`

#### update()  [header link](class-aws-glacier-treehash-method-update.md)

Adds data to the hash.

`
    public
                    update(mixed $data) : mixed`

##### Parameters

$data
: mixed

Data to add to the hash

##### Tags  [header link](class-aws-glacier-treehash-method-update-tags.md)

throwsLogicException

if the root tree hash is already calculated

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-glacier-treehash-toc-constants.md)
  - [Methods](class-aws-glacier-treehash-toc-methods.md)
- Constants
  - [EMPTY\_HASH](class-aws-glacier-treehash-constant-empty-hash.md)
  - [MB](class-aws-glacier-treehash-constant-mb.md)
- Methods
  - [\_\_construct()](class-aws-glacier-treehash-method-construct.md)
  - [addChecksum()](class-aws-glacier-treehash-method-addchecksum.md)
  - [complete()](class-aws-glacier-treehash-method-complete.md)
  - [reset()](class-aws-glacier-treehash-method-reset.md)
  - [update()](class-aws-glacier-treehash-method-update.md)

[Back To Top](class-aws-glacier-treehash-top.md)
