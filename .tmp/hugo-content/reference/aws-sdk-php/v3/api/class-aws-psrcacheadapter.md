Menu

- [Aws](namespace-aws.md)

## PsrCacheAdapter        in package    - [Aws](package-aws.md)       implements  [CacheInterface](class-aws-cacheinterface.md)

### Table of Contents  [header link](class-aws-psrcacheadapter-toc.md)

#### Interfaces  [header link](class-aws-psrcacheadapter-toc-interfaces.md)

[CacheInterface](class-aws-cacheinterface.md)Represents a simple cache interface.

#### Methods  [header link](class-aws-psrcacheadapter-toc-methods.md)

[\_\_construct()](class-aws-psrcacheadapter-method-construct.md)
: mixed [get()](class-aws-psrcacheadapter-method-get.md)
: mixed\|null Get a cache item by key.[remove()](class-aws-psrcacheadapter-method-remove.md)
: mixed Remove a cache key.[set()](class-aws-psrcacheadapter-method-set.md)
: mixed Set a cache key value.

### Methods  [header link](class-aws-psrcacheadapter-methods.md)

#### \_\_construct()  [header link](class-aws-psrcacheadapter-method-construct.md)

`
    public
                    __construct(CacheItemPoolInterface $pool) : mixed`

##### Parameters

$pool
: CacheItemPoolInterface

#### get()  [header link](class-aws-psrcacheadapter-method-get.md)

Get a cache item by key.

`
    public
                    get(mixed $key) : mixed|null`

##### Parameters

$key
: mixed

Key to retrieve.

##### Return values

mixed\|null
—

Returns the value or null if not found.

#### remove()  [header link](class-aws-psrcacheadapter-method-remove.md)

Remove a cache key.

`
    public
                    remove(mixed $key) : mixed`

##### Parameters

$key
: mixed

Key to remove.

#### set()  [header link](class-aws-psrcacheadapter-method-set.md)

Set a cache key value.

`
    public
                    set(mixed $key, mixed $value[, mixed $ttl = 0 ]) : mixed`

##### Parameters

$key
: mixed

Key to set

$value
: mixed

Value to set.

$ttl
: mixed
= 0

Number of seconds the item is allowed to live. Set
to 0 to allow an unlimited lifetime.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-psrcacheadapter-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-psrcacheadapter-method-construct.md)
  - [get()](class-aws-psrcacheadapter-method-get.md)
  - [remove()](class-aws-psrcacheadapter-method-remove.md)
  - [set()](class-aws-psrcacheadapter-method-set.md)

[Back To Top](class-aws-psrcacheadapter-top.md)
