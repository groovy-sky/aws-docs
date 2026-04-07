Menu

- [Aws](namespace-aws.md)

## LruArrayCache        in package    - [Aws](package-aws.md)       implements  [CacheInterface](class-aws-cacheinterface.md), Countable

Simple in-memory LRU cache that limits the number of cached entries.

The LRU cache is implemented using PHP's ordered associative array. When
accessing an element, the element is removed from the hash and re-added to
ensure that recently used items are always at the end of the list while
least recently used are at the beginning. When a value is added to the
cache, if the number of cached items exceeds the allowed number, the first
N number of items are removed from the array.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#toc-interfaces)

[CacheInterface](class-aws-cacheinterface.md)Represents a simple cache interface.Countable

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method___construct)
: mixed [count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method_count)
: int [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method_get)
: mixed\|null Get a cache item by key.[remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method_remove)
: mixed Remove a cache key.[set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method_set)
: mixed Set a cache key value.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#method___construct)

`
    public
                    __construct([int $maxItems = 1000 ]) : mixed`

##### Parameters

$maxItems
: int
= 1000

Maximum number of allowed cache items.

#### count()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#method_count)

`
    public
                    count() : int`

##### Return values

int

#### get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#method_get)

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

#### remove()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#method_remove)

Remove a cache key.

`
    public
                    remove(mixed $key) : mixed`

##### Parameters

$key
: mixed

Key to remove.

#### set()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html\#method_set)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method___construct)
  - [count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method_count)
  - [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method_get)
  - [remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method_remove)
  - [set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#method_set)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.LruArrayCache.html#top)
