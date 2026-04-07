Menu

- [Aws](namespace-aws.md)

## PsrCacheAdapter        in package    - [Aws](package-aws.md)       implements  [CacheInterface](class-aws-cacheinterface.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html\#toc-interfaces)

[CacheInterface](class-aws-cacheinterface.md)Represents a simple cache interface.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#method___construct)
: mixed [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#method_get)
: mixed\|null Get a cache item by key.[remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#method_remove)
: mixed Remove a cache key.[set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#method_set)
: mixed Set a cache key value.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html\#method___construct)

`
    public
                    __construct(CacheItemPoolInterface $pool) : mixed`

##### Parameters

$pool
: CacheItemPoolInterface

#### get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html\#method_get)

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

#### remove()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html\#method_remove)

Remove a cache key.

`
    public
                    remove(mixed $key) : mixed`

##### Parameters

$key
: mixed

Key to remove.

#### set()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html\#method_set)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#method___construct)
  - [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#method_get)
  - [remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#method_remove)
  - [set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#method_set)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.PsrCacheAdapter.html#top)
