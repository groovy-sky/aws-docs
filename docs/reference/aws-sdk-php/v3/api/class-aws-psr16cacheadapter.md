Menu

- [Aws](namespace-aws.md)

## Psr16CacheAdapter        in package    - [Aws](package-aws.md)       implements  [CacheInterface](class-aws-cacheinterface.md)

### Table of Contents  [header link](class-aws-psr16cacheadapter-toc.md)

#### Interfaces  [header link](class-aws-psr16cacheadapter-toc-interfaces.md)

[CacheInterface](class-aws-cacheinterface.md)Represents a simple cache interface.

#### Methods  [header link](class-aws-psr16cacheadapter-toc-methods.md)

[\_\_construct()](class-aws-psr16cacheadapter-method-construct.md)
: mixed [get()](class-aws-psr16cacheadapter-method-get.md)
: mixed\|null Get a cache item by key.[remove()](class-aws-psr16cacheadapter-method-remove.md)
: mixed Remove a cache key.[set()](class-aws-psr16cacheadapter-method-set.md)
: mixed Set a cache key value.

### Methods  [header link](class-aws-psr16cacheadapter-methods.md)

#### \_\_construct()  [header link](class-aws-psr16cacheadapter-method-construct.md)

`
    public
                    __construct(CacheInterface $cache) : mixed`

##### Parameters

$cache
: CacheInterface

#### get()  [header link](class-aws-psr16cacheadapter-method-get.md)

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

#### remove()  [header link](class-aws-psr16cacheadapter-method-remove.md)

Remove a cache key.

`
    public
                    remove(mixed $key) : mixed`

##### Parameters

$key
: mixed

Key to remove.

#### set()  [header link](class-aws-psr16cacheadapter-method-set.md)

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
  - [Methods](class-aws-psr16cacheadapter-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-psr16cacheadapter-method-construct.md)
  - [get()](class-aws-psr16cacheadapter-method-get.md)
  - [remove()](class-aws-psr16cacheadapter-method-remove.md)
  - [set()](class-aws-psr16cacheadapter-method-set.md)

[Back To Top](class-aws-psr16cacheadapter-top.md)
