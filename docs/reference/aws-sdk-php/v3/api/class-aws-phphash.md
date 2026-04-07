Menu

- [Aws](namespace-aws.md)

## PhpHash        in package    - [Aws](package-aws.md)       implements  [HashInterface](class-aws-hashinterface.md)

Incremental hashing using PHP's hash functions.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html\#toc-interfaces)

[HashInterface](class-aws-hashinterface.md)Interface that allows implementing various incremental hashes.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#method___construct)
: mixed [complete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#method_complete)
: string Finalizes the incremental hash and returns the resulting digest.[reset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#method_reset)
: mixed Removes all data from the hash, effectively starting a new hash.[update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#method_update)
: mixed Adds data to the hash.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html\#method___construct)

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

#### complete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html\#method_complete)

Finalizes the incremental hash and returns the resulting digest.

`
    public
                    complete() : string`

##### Return values

string

#### reset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html\#method_reset)

Removes all data from the hash, effectively starting a new hash.

`
    public
                    reset() : mixed`

#### update()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html\#method_update)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#method___construct)
  - [complete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#method_complete)
  - [reset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#method_reset)
  - [update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#method_update)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PhpHash.html#top)
