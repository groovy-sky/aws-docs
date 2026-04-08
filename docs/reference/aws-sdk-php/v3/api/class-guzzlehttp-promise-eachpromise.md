Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## EachPromise        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Represents a promise that iterates over many promises and invokes
side-effect functions in the process.

##### Tags  [header link](class-guzzlehttp-promise-eachpromise-tags.md)

final

### Table of Contents  [header link](class-guzzlehttp-promise-eachpromise-toc.md)

#### Interfaces  [header link](class-guzzlehttp-promise-eachpromise-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](class-guzzlehttp-promise-eachpromise-toc-methods.md)

[\_\_construct()](class-guzzlehttp-promise-eachpromise-method-construct.md)
: mixed Configuration hash can include the following key value pairs:[promise()](class-guzzlehttp-promise-eachpromise-method-promise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Methods  [header link](class-guzzlehttp-promise-eachpromise-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-promise-eachpromise-method-construct.md)

Configuration hash can include the following key value pairs:

`
    public
                    __construct(mixed $iterable[, array<string|int, mixed> $config = [] ]) : mixed`

- fulfilled: (callable) Invoked when a promise fulfills. The function
is invoked with three arguments: the fulfillment value, the index
position from the iterable list of the promise, and the aggregate
promise that manages all of the promises. The aggregate promise may
be resolved from within the callback to short-circuit the promise.
- rejected: (callable) Invoked when a promise is rejected. The
function is invoked with three arguments: the rejection reason, the
index position from the iterable list of the promise, and the
aggregate promise that manages all of the promises. The aggregate
promise may be resolved from within the callback to short-circuit
the promise.
- concurrency: (integer) Pass this configuration option to limit the
allowed number of outstanding concurrently executing promises,
creating a capped pool of promises. There is no limit by default.

##### Parameters

$iterable
: mixed

Promises or values to iterate.

$config
: array<string\|int, mixed>
= \[\]

Configuration options

#### promise()  [header link](class-guzzlehttp-promise-eachpromise-method-promise.md)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Tags  [header link](class-guzzlehttp-promise-eachpromise-method-promise-tags.md)

psalm-suppress

InvalidNullableReturnType

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-promise-eachpromise-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-promise-eachpromise-method-construct.md)
  - [promise()](class-guzzlehttp-promise-eachpromise-method-promise.md)

[Back To Top](class-guzzlehttp-promise-eachpromise-top.md)
