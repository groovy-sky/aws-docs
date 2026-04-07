Menu

- [Aws](namespace-aws.md)

## ResultPaginator        in package    - [Aws](package-aws.md)       implements  Iterator

Iterator that yields each page of results of a pageable operation.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#toc-interfaces)

Iterator

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method___construct)
: mixed [current()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_current)
: [Result](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html)[each()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_each)
: [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Promise.html)Runs a paginator asynchronously and uses a callback to handle results.[key()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_key)
: mixed [next()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_next)
: void [rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_rewind)
: void [search()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_search)
: IteratorReturns an iterator that iterates over the values of applying a JMESPath
search to each result yielded by the iterator as a flat sequence.[valid()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_valid)
: bool

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#method___construct)

`
    public
                    __construct(AwsClientInterface $client, string $operation, array<string|int, mixed> $args, array<string|int, mixed> $config) : mixed`

##### Parameters

$client
: [AwsClientInterface](class-aws-awsclientinterface.md)$operation
: string$args
: array<string\|int, mixed>$config
: array<string\|int, mixed>

#### current()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#method_current)

`
    public
                    current() : Result`

##### Return values

[Result](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html)

#### each()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#method_each)

Runs a paginator asynchronously and uses a callback to handle results.

`
    public
                    each(callable $handleResult) : Promise`

The callback should have the signature: function (Aws\\Result $result).
A non-null return value from the callback will be yielded by the
promise. This means that you can return promises from the callback that
will need to be resolved before continuing iteration over the remaining
items, essentially merging in other promises to the iteration. The last
non-null value returned by the callback will be the result that fulfills
the promise to any downstream promises.

##### Parameters

$handleResult
: callable

Callback for handling each page of results.
The callback accepts the result that was
yielded as a single argument. If the
callback returns a promise, the promise
will be merged into the coroutine.

##### Return values

[Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Promise.html)

#### key()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#method_key)

`
    public
                    key() : mixed`

#### next()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#method_next)

`
    public
                    next() : void`

#### rewind()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#method_rewind)

`
    public
                    rewind() : void`

#### search()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#method_search)

Returns an iterator that iterates over the values of applying a JMESPath
search to each result yielded by the iterator as a flat sequence.

`
    public
                    search(string $expression) : Iterator`

##### Parameters

$expression
: string

JMESPath expression to apply to each result.

##### Return values

Iterator

#### valid()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html\#method_valid)

`
    public
                    valid() : bool`

##### Return values

bool
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method___construct)
  - [current()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_current)
  - [each()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_each)
  - [key()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_key)
  - [next()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_next)
  - [rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_rewind)
  - [search()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_search)
  - [valid()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#method_valid)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html#top)
