Menu

- [Aws](namespace-aws.md)

## CacheInterface     in    - [Aws](package-aws.md)

Represents a simple cache interface.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html\#toc-methods)

[get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#method_get)
: mixed\|null Get a cache item by key.[remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#method_remove)
: mixed Remove a cache key.[set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#method_set)
: mixed Set a cache key value.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html\#methods)

#### get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html\#method_get)

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

#### remove()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html\#method_remove)

Remove a cache key.

`
    public
                    remove(string $key) : mixed`

##### Parameters

$key
: string

Key to remove.

#### set()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html\#method_set)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#toc-methods)
- Methods
  - [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#method_get)
  - [remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#method_remove)
  - [set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#method_set)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CacheInterface.html#top)
