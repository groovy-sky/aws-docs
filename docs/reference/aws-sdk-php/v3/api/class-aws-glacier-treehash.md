Menu

- [Aws](namespace-aws.md)
- [Glacier](namespace-aws-glacier.md)

## TreeHash        in package    - [Aws](package-aws.md)       implements  [HashInterface](class-aws-hashinterface.md)

Encapsulates the creation of a tree hash from streamed data

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#toc-interfaces)

[HashInterface](class-aws-hashinterface.md)Interface that allows implementing various incremental hashes.

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#toc-constants)

[EMPTY\_HASH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#constant_EMPTY_HASH)
= 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855' [MB](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#constant_MB)
= 1048576

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method___construct)
: mixed [addChecksum()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method_addChecksum)
: self Add a checksum to the tree hash directly[complete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method_complete)
: string Finalizes the incremental hash and returns the resulting digest.[reset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method_reset)
: mixed Removes all data from the hash, effectively starting a new hash.[update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method_update)
: mixed Adds data to the hash.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#constants)

#### EMPTY\_HASH  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#constant_EMPTY_HASH)

`
    public
        mixed
    EMPTY_HASH
    = 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'
`

#### MB  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#constant_MB)

`
    public
        mixed
    MB
    = 1048576
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#method___construct)

`
    public
                    __construct([mixed $algorithm = 'sha256' ]) : mixed`

##### Parameters

$algorithm
: mixed
= 'sha256'

#### addChecksum()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#method_addChecksum)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#method_addChecksum\#tags)

throwsLogicException

if the root tree hash is already calculated

##### Return values

self

#### complete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#method_complete)

Finalizes the incremental hash and returns the resulting digest.

`
    public
                    complete() : string`

##### Return values

string

#### reset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#method_reset)

Removes all data from the hash, effectively starting a new hash.

`
    public
                    reset() : mixed`

#### update()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#method_update)

Adds data to the hash.

`
    public
                    update(mixed $data) : mixed`

##### Parameters

$data
: mixed

Data to add to the hash

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html\#method_update\#tags)

throwsLogicException

if the root tree hash is already calculated

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#toc-methods)
- Constants
  - [EMPTY\_HASH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#constant_EMPTY_HASH)
  - [MB](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#constant_MB)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method___construct)
  - [addChecksum()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method_addChecksum)
  - [complete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method_complete)
  - [reset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method_reset)
  - [update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#method_update)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Glacier.TreeHash.html#top)
