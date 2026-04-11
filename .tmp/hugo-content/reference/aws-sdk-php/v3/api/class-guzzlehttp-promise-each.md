Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## Each        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-guzzlehttp-promise-each-toc.md)

#### Methods  [header link](class-guzzlehttp-promise-each-toc-methods.md)

[of()](class-guzzlehttp-promise-each-method-of.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Given an iterator that yields promises or values, returns a promise that
is fulfilled with a null value when the iterator has been consumed or
the aggregate promise has been fulfilled or rejected.[ofLimit()](class-guzzlehttp-promise-each-method-oflimit.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Like of, but only allows a certain number of outstanding promises at any
given time.[ofLimitAll()](class-guzzlehttp-promise-each-method-oflimitall.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Like limit, but ensures that no promise in the given $iterable argument
is rejected. If any promise is rejected, then the aggregate promise is
rejected with the encountered rejection.

### Methods  [header link](class-guzzlehttp-promise-each-methods.md)

#### of()  [header link](class-guzzlehttp-promise-each-method-of.md)

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

#### ofLimit()  [header link](class-guzzlehttp-promise-each-method-oflimit.md)

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

#### ofLimitAll()  [header link](class-guzzlehttp-promise-each-method-oflimitall.md)

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
  - [Methods](class-guzzlehttp-promise-each-toc-methods.md)
- Methods
  - [of()](class-guzzlehttp-promise-each-method-of.md)
  - [ofLimit()](class-guzzlehttp-promise-each-method-oflimit.md)
  - [ofLimitAll()](class-guzzlehttp-promise-each-method-oflimitall.md)

[Back To Top](class-guzzlehttp-promise-each-top.md)
