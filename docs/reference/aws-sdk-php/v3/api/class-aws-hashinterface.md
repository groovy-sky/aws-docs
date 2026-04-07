Menu

- [Aws](namespace-aws.md)

## HashInterface     in    - [Aws](package-aws.md)

Interface that allows implementing various incremental hashes.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html\#toc-methods)

[complete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#method_complete)
: string Finalizes the incremental hash and returns the resulting digest.[reset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#method_reset)
: mixed Removes all data from the hash, effectively starting a new hash.[update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#method_update)
: mixed Adds data to the hash.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html\#methods)

#### complete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html\#method_complete)

Finalizes the incremental hash and returns the resulting digest.

`
    public
                    complete() : string`

##### Return values

string

#### reset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html\#method_reset)

Removes all data from the hash, effectively starting a new hash.

`
    public
                    reset() : mixed`

#### update()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html\#method_update)

Adds data to the hash.

`
    public
                    update(string $data) : mixed`

##### Parameters

$data
: string

Data to add to the hash

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#toc-methods)
- Methods
  - [complete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#method_complete)
  - [reset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#method_reset)
  - [update()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#method_update)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashInterface.html#top)
