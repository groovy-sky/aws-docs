Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html)

## EachPromise        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Represents a promise that iterates over many promises and invokes
side-effect functions in the process.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html\#tags)

final

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html\#toc-interfaces)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html#method___construct)
: mixed Configuration hash can include the following key value pairs:[promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html\#method___construct)

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

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html\#method_promise)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html\#method_promise\#tags)

psalm-suppress

InvalidNullableReturnType

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html#method___construct)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html#method_promise)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.EachPromise.html#top)
