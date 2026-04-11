Menu

- [Aws](namespace-aws.md)

## DoctrineCacheAdapter        in package    - [Aws](package-aws.md)       implements  [CacheInterface](class-aws-cacheinterface.md), Cache

### Table of Contents  [header link](class-aws-doctrinecacheadapter-toc.md)

#### Interfaces  [header link](class-aws-doctrinecacheadapter-toc-interfaces.md)

[CacheInterface](class-aws-cacheinterface.md)Represents a simple cache interface.Cache

#### Methods  [header link](class-aws-doctrinecacheadapter-toc-methods.md)

[\_\_construct()](class-aws-doctrinecacheadapter-method-construct.md)
: mixed [contains()](class-aws-doctrinecacheadapter-method-contains.md)
: bool [delete()](class-aws-doctrinecacheadapter-method-delete.md)
: bool [fetch()](class-aws-doctrinecacheadapter-method-fetch.md)
: mixed [get()](class-aws-doctrinecacheadapter-method-get.md)
: mixed\|null Get a cache item by key.[getStats()](class-aws-doctrinecacheadapter-method-getstats.md)
: array<string\|int, mixed>\|null [remove()](class-aws-doctrinecacheadapter-method-remove.md)
: mixed Remove a cache key.[save()](class-aws-doctrinecacheadapter-method-save.md)
: bool [set()](class-aws-doctrinecacheadapter-method-set.md)
: mixed Set a cache key value.

### Methods  [header link](class-aws-doctrinecacheadapter-methods.md)

#### \_\_construct()  [header link](class-aws-doctrinecacheadapter-method-construct.md)

`
    public
                    __construct(Cache $cache) : mixed`

##### Parameters

$cache
: Cache

#### contains()  [header link](class-aws-doctrinecacheadapter-method-contains.md)

`
    public
                    contains(mixed $key) : bool`

##### Parameters

$key
: mixed

##### Return values

bool

#### delete()  [header link](class-aws-doctrinecacheadapter-method-delete.md)

`
    public
                    delete(mixed $key) : bool`

##### Parameters

$key
: mixed

##### Return values

bool

#### fetch()  [header link](class-aws-doctrinecacheadapter-method-fetch.md)

`
    public
                    fetch(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### get()  [header link](class-aws-doctrinecacheadapter-method-get.md)

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

#### getStats()  [header link](class-aws-doctrinecacheadapter-method-getstats.md)

`
    public
                    getStats() : array<string|int, mixed>|null`

##### Return values

array<string\|int, mixed>\|null

#### remove()  [header link](class-aws-doctrinecacheadapter-method-remove.md)

Remove a cache key.

`
    public
                    remove(mixed $key) : mixed`

##### Parameters

$key
: mixed

Key to remove.

#### save()  [header link](class-aws-doctrinecacheadapter-method-save.md)

`
    public
                    save(mixed $key, mixed $value[, mixed $ttl = 0 ]) : bool`

##### Parameters

$key
: mixed$value
: mixed$ttl
: mixed
= 0

##### Return values

bool

#### set()  [header link](class-aws-doctrinecacheadapter-method-set.md)

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
  - [Methods](class-aws-doctrinecacheadapter-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-doctrinecacheadapter-method-construct.md)
  - [contains()](class-aws-doctrinecacheadapter-method-contains.md)
  - [delete()](class-aws-doctrinecacheadapter-method-delete.md)
  - [fetch()](class-aws-doctrinecacheadapter-method-fetch.md)
  - [get()](class-aws-doctrinecacheadapter-method-get.md)
  - [getStats()](class-aws-doctrinecacheadapter-method-getstats.md)
  - [remove()](class-aws-doctrinecacheadapter-method-remove.md)
  - [save()](class-aws-doctrinecacheadapter-method-save.md)
  - [set()](class-aws-doctrinecacheadapter-method-set.md)

[Back To Top](class-aws-doctrinecacheadapter-top.md)
