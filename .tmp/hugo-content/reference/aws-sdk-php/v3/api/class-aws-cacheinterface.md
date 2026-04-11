Menu

- [Aws](namespace-aws.md)

## CacheInterface     in    - [Aws](package-aws.md)

Represents a simple cache interface.

### Table of Contents  [header link](class-aws-cacheinterface-toc.md)

#### Methods  [header link](class-aws-cacheinterface-toc-methods.md)

[get()](class-aws-cacheinterface-method-get.md)
: mixed\|null Get a cache item by key.[remove()](class-aws-cacheinterface-method-remove.md)
: mixed Remove a cache key.[set()](class-aws-cacheinterface-method-set.md)
: mixed Set a cache key value.

### Methods  [header link](class-aws-cacheinterface-methods.md)

#### get()  [header link](class-aws-cacheinterface-method-get.md)

Get a cache item by key.

`
    public
                    get(string $key) : mixed|null`

##### Parameters

$key
: string

Key to retrieve.

##### Return values

mixed\|null
—

Returns the value or null if not found.

#### remove()  [header link](class-aws-cacheinterface-method-remove.md)

Remove a cache key.

`
    public
                    remove(string $key) : mixed`

##### Parameters

$key
: string

Key to remove.

#### set()  [header link](class-aws-cacheinterface-method-set.md)

Set a cache key value.

`
    public
                    set(string $key, mixed $value[, int $ttl = 0 ]) : mixed`

##### Parameters

$key
: string

Key to set

$value
: mixed

Value to set.

$ttl
: int
= 0

Number of seconds the item is allowed to live. Set
to 0 to allow an unlimited lifetime.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-cacheinterface-toc-constants.md)
  - [Methods](class-aws-cacheinterface-toc-methods.md)
- Methods
  - [get()](class-aws-cacheinterface-method-get.md)
  - [remove()](class-aws-cacheinterface-method-remove.md)
  - [set()](class-aws-cacheinterface-method-set.md)

[Back To Top](class-aws-cacheinterface-top.md)
