Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html)

## Each        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html\#toc-methods)

[of()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html#method_of)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Given an iterator that yields promises or values, returns a promise that
is fulfilled with a null value when the iterator has been consumed or
the aggregate promise has been fulfilled or rejected.[ofLimit()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html#method_ofLimit)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Like of, but only allows a certain number of outstanding promises at any
given time.[ofLimitAll()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html#method_ofLimitAll)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Like limit, but ensures that no promise in the given $iterable argument
is rejected. If any promise is rejected, then the aggregate promise is
rejected with the encountered rejection.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html\#methods)

#### of()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html\#method_of)

Given an iterator that yields promises or values, returns a promise that
is fulfilled with a null value when the iterator has been consumed or
the aggregate promise has been fulfilled or rejected.

`
    public
            static        of(mixed $iterable[, callable|null $onFulfilled = null ][, callable|null $onRejected = null ]) : PromiseInterface`

$onFulfilled is a function that accepts the fulfilled value, iterator
index, and the aggregate promise. The callback can invoke any necessary
side effects and choose to resolve or reject the aggregate if needed.

$onRejected is a function that accepts the rejection reason, iterator
index, and the aggregate promise. The callback can invoke any necessary
side effects and choose to resolve or reject the aggregate if needed.

##### Parameters

$iterable
: mixed

Iterator or array to iterate over.

$onFulfilled
: callable\|null
= null$onRejected
: callable\|null
= null

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### ofLimit()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html\#method_ofLimit)

Like of, but only allows a certain number of outstanding promises at any
given time.

`
    public
            static        ofLimit(mixed $iterable, int|callable $concurrency[, callable|null $onFulfilled = null ][, callable|null $onRejected = null ]) : PromiseInterface`

$concurrency may be an integer or a function that accepts the number of
pending promises and returns a numeric concurrency limit value to allow
for dynamic a concurrency size.

##### Parameters

$iterable
: mixed$concurrency
: int\|callable$onFulfilled
: callable\|null
= null$onRejected
: callable\|null
= null

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### ofLimitAll()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html\#method_ofLimitAll)

Like limit, but ensures that no promise in the given $iterable argument
is rejected. If any promise is rejected, then the aggregate promise is
rejected with the encountered rejection.

`
    public
            static        ofLimitAll(mixed $iterable, int|callable $concurrency[, callable|null $onFulfilled = null ]) : PromiseInterface`

##### Parameters

$iterable
: mixed$concurrency
: int\|callable$onFulfilled
: callable\|null
= null

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html#toc-methods)
- Methods
  - [of()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html#method_of)
  - [ofLimit()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html#method_ofLimit)
  - [ofLimitAll()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html#method_ofLimitAll)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Each.html#top)
