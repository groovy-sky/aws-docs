Menu

- [Aws](namespace-aws.md)

## DoctrineCacheAdapter        in package    - [Aws](package-aws.md)       implements  [CacheInterface](class-aws-cacheinterface.md), Cache

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#toc-interfaces)

[CacheInterface](class-aws-cacheinterface.md)Represents a simple cache interface.Cache

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method___construct)
: mixed [contains()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_contains)
: bool [delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_delete)
: bool [fetch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_fetch)
: mixed [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_get)
: mixed\|null Get a cache item by key.[getStats()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_getStats)
: array<string\|int, mixed>\|null [remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_remove)
: mixed Remove a cache key.[save()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_save)
: bool [set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_set)
: mixed Set a cache key value.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method___construct)

`
    public
                    __construct(Cache $cache) : mixed`

##### Parameters

$cache
: Cache

#### contains()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method_contains)

`
    public
                    contains(mixed $key) : bool`

##### Parameters

$key
: mixed

##### Return values

bool

#### delete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method_delete)

`
    public
                    delete(mixed $key) : bool`

##### Parameters

$key
: mixed

##### Return values

bool

#### fetch()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method_fetch)

`
    public
                    fetch(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method_get)

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

#### getStats()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method_getStats)

`
    public
                    getStats() : array<string|int, mixed>|null`

##### Return values

array<string\|int, mixed>\|null

#### remove()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method_remove)

Remove a cache key.

`
    public
                    remove(mixed $key) : mixed`

##### Parameters

$key
: mixed

Key to remove.

#### save()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method_save)

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

#### set()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html\#method_set)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method___construct)
  - [contains()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_contains)
  - [delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_delete)
  - [fetch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_fetch)
  - [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_get)
  - [getStats()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_getStats)
  - [remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_remove)
  - [save()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_save)
  - [set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#method_set)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DoctrineCacheAdapter.html#top)
