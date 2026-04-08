Menu

- [Aws](namespace-aws.md)

## HashInterface     in    - [Aws](package-aws.md)

Interface that allows implementing various incremental hashes.

### Table of Contents  [header link](class-aws-hashinterface-toc.md)

#### Methods  [header link](class-aws-hashinterface-toc-methods.md)

[complete()](class-aws-hashinterface-method-complete.md)
: string Finalizes the incremental hash and returns the resulting digest.[reset()](class-aws-hashinterface-method-reset.md)
: mixed Removes all data from the hash, effectively starting a new hash.[update()](class-aws-hashinterface-method-update.md)
: mixed Adds data to the hash.

### Methods  [header link](class-aws-hashinterface-methods.md)

#### complete()  [header link](class-aws-hashinterface-method-complete.md)

Finalizes the incremental hash and returns the resulting digest.

`
    public
                    complete() : string`

##### Return values

string

#### reset()  [header link](class-aws-hashinterface-method-reset.md)

Removes all data from the hash, effectively starting a new hash.

`
    public
                    reset() : mixed`

#### update()  [header link](class-aws-hashinterface-method-update.md)

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
  - [Constants](class-aws-hashinterface-toc-constants.md)
  - [Methods](class-aws-hashinterface-toc-methods.md)
- Methods
  - [complete()](class-aws-hashinterface-method-complete.md)
  - [reset()](class-aws-hashinterface-method-reset.md)
  - [update()](class-aws-hashinterface-method-update.md)

[Back To Top](class-aws-hashinterface-top.md)
