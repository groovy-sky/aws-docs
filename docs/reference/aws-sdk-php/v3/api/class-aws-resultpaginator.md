Menu

- [Aws](namespace-aws.md)

## ResultPaginator        in package    - [Aws](package-aws.md)       implements  Iterator

Iterator that yields each page of results of a pageable operation.

### Table of Contents  [header link](class-aws-resultpaginator-toc.md)

#### Interfaces  [header link](class-aws-resultpaginator-toc-interfaces.md)

Iterator

#### Methods  [header link](class-aws-resultpaginator-toc-methods.md)

[\_\_construct()](class-aws-resultpaginator-method-construct.md)
: mixed [current()](class-aws-resultpaginator-method-current.md)
: [Result](class-aws-result.md)[each()](class-aws-resultpaginator-method-each.md)
: [Promise](class-guzzlehttp-promise-promise.md)Runs a paginator asynchronously and uses a callback to handle results.[key()](class-aws-resultpaginator-method-key.md)
: mixed [next()](class-aws-resultpaginator-method-next.md)
: void [rewind()](class-aws-resultpaginator-method-rewind.md)
: void [search()](class-aws-resultpaginator-method-search.md)
: IteratorReturns an iterator that iterates over the values of applying a JMESPath
search to each result yielded by the iterator as a flat sequence.[valid()](class-aws-resultpaginator-method-valid.md)
: bool

### Methods  [header link](class-aws-resultpaginator-methods.md)

#### \_\_construct()  [header link](class-aws-resultpaginator-method-construct.md)

`
    public
                    __construct(AwsClientInterface $client, string $operation, array<string|int, mixed> $args, array<string|int, mixed> $config) : mixed`

##### Parameters

$client
: [AwsClientInterface](class-aws-awsclientinterface.md)$operation
: string$args
: array<string\|int, mixed>$config
: array<string\|int, mixed>

#### current()  [header link](class-aws-resultpaginator-method-current.md)

`
    public
                    current() : Result`

##### Return values

[Result](class-aws-result.md)

#### each()  [header link](class-aws-resultpaginator-method-each.md)

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

[Promise](class-guzzlehttp-promise-promise.md)

#### key()  [header link](class-aws-resultpaginator-method-key.md)

`
    public
                    key() : mixed`

#### next()  [header link](class-aws-resultpaginator-method-next.md)

`
    public
                    next() : void`

#### rewind()  [header link](class-aws-resultpaginator-method-rewind.md)

`
    public
                    rewind() : void`

#### search()  [header link](class-aws-resultpaginator-method-search.md)

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

#### valid()  [header link](class-aws-resultpaginator-method-valid.md)

`
    public
                    valid() : bool`

##### Return values

bool
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-resultpaginator-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-resultpaginator-method-construct.md)
  - [current()](class-aws-resultpaginator-method-current.md)
  - [each()](class-aws-resultpaginator-method-each.md)
  - [key()](class-aws-resultpaginator-method-key.md)
  - [next()](class-aws-resultpaginator-method-next.md)
  - [rewind()](class-aws-resultpaginator-method-rewind.md)
  - [search()](class-aws-resultpaginator-method-search.md)
  - [valid()](class-aws-resultpaginator-method-valid.md)

[Back To Top](class-aws-resultpaginator-top.md)
