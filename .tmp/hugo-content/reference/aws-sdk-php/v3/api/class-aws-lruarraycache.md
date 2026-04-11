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

### Table of Contents  [header link](class-aws-lruarraycache-toc.md)

#### Interfaces  [header link](class-aws-lruarraycache-toc-interfaces.md)

[CacheInterface](class-aws-cacheinterface.md)Represents a simple cache interface.Countable

#### Methods  [header link](class-aws-lruarraycache-toc-methods.md)

[\_\_construct()](class-aws-lruarraycache-method-construct.md)
: mixed [count()](class-aws-lruarraycache-method-count.md)
: int [get()](class-aws-lruarraycache-method-get.md)
: mixed\|null Get a cache item by key.[remove()](class-aws-lruarraycache-method-remove.md)
: mixed Remove a cache key.[set()](class-aws-lruarraycache-method-set.md)
: mixed Set a cache key value.

### Methods  [header link](class-aws-lruarraycache-methods.md)

#### \_\_construct()  [header link](class-aws-lruarraycache-method-construct.md)

`
    public
                    __construct([int $maxItems = 1000 ]) : mixed`

##### Parameters

$maxItems
: int
= 1000

Maximum number of allowed cache items.

#### count()  [header link](class-aws-lruarraycache-method-count.md)

`
    public
                    count() : int`

##### Return values

int

#### get()  [header link](class-aws-lruarraycache-method-get.md)

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

#### remove()  [header link](class-aws-lruarraycache-method-remove.md)

Remove a cache key.

`
    public
                    remove(mixed $key) : mixed`

##### Parameters

$key
: mixed

Key to remove.

#### set()  [header link](class-aws-lruarraycache-method-set.md)

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
  - [Methods](class-aws-lruarraycache-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-lruarraycache-method-construct.md)
  - [count()](class-aws-lruarraycache-method-count.md)
  - [get()](class-aws-lruarraycache-method-get.md)
  - [remove()](class-aws-lruarraycache-method-remove.md)
  - [set()](class-aws-lruarraycache-method-set.md)

[Back To Top](class-aws-lruarraycache-top.md)
