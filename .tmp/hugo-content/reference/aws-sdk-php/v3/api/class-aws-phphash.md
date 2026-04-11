Menu

- [Aws](namespace-aws.md)

## PhpHash        in package    - [Aws](package-aws.md)       implements  [HashInterface](class-aws-hashinterface.md)

Incremental hashing using PHP's hash functions.

### Table of Contents  [header link](class-aws-phphash-toc.md)

#### Interfaces  [header link](class-aws-phphash-toc-interfaces.md)

[HashInterface](class-aws-hashinterface.md)Interface that allows implementing various incremental hashes.

#### Methods  [header link](class-aws-phphash-toc-methods.md)

[\_\_construct()](class-aws-phphash-method-construct.md)
: mixed [complete()](class-aws-phphash-method-complete.md)
: string Finalizes the incremental hash and returns the resulting digest.[reset()](class-aws-phphash-method-reset.md)
: mixed Removes all data from the hash, effectively starting a new hash.[update()](class-aws-phphash-method-update.md)
: mixed Adds data to the hash.

### Methods  [header link](class-aws-phphash-methods.md)

#### \_\_construct()  [header link](class-aws-phphash-method-construct.md)

`
    public
                    __construct(string $algo[, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$algo
: string

Hashing algorithm. One of PHP's hash\_algos()
return values (e.g. md5, sha1, etc...).

$options
: array<string\|int, mixed>
= \[\]

Associative array of hashing options:

- key: Secret key used with the hashing algorithm.
- base64: Set to true to base64 encode the value when complete.

#### complete()  [header link](class-aws-phphash-method-complete.md)

Finalizes the incremental hash and returns the resulting digest.

`
    public
                    complete() : string`

##### Return values

string

#### reset()  [header link](class-aws-phphash-method-reset.md)

Removes all data from the hash, effectively starting a new hash.

`
    public
                    reset() : mixed`

#### update()  [header link](class-aws-phphash-method-update.md)

Adds data to the hash.

`
    public
                    update(mixed $data) : mixed`

##### Parameters

$data
: mixed

Data to add to the hash

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-phphash-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-phphash-method-construct.md)
  - [complete()](class-aws-phphash-method-complete.md)
  - [reset()](class-aws-phphash-method-reset.md)
  - [update()](class-aws-phphash-method-update.md)

[Back To Top](class-aws-phphash-top.md)
