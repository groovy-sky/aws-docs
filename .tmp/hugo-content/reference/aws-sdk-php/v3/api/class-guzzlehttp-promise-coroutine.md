Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## Coroutine        in package    - [Aws](package-aws.md)       implements  [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

FinalYes

Creates a promise that is resolved using a generator that yields values or
promises (somewhat similar to C#'s async keyword).

When called, the Coroutine::of method will start an instance of the generator
and returns a promise that is fulfilled with its final yielded value.

Control is returned back to the generator when the yielded promise settles.
This can lead to less verbose code when doing lots of sequential async calls
with minimal processing in between.

```prettyprint
use GuzzleHttp\Promise;

function createPromise($value) {
    return new Promise\FulfilledPromise($value);
}

$promise = Promise\Coroutine::of(function () {
    $value = (yield createPromise('a'));
    try {
        $value = (yield createPromise($value . 'b'));
    } catch (\Throwable $e) {
        // The promise was rejected.
    }
    yield $value . 'c';
});

// Outputs "abc"
$promise->then(function ($v) { echo $v; });

```

##### Tags  [header link](class-guzzlehttp-promise-coroutine-tags.md)

see[https://github.com/petkaantonov/bluebird/blob/master/API.md#generators](https://github.com/petkaantonov/bluebird/blob/master/API.md)

inspiration

### Table of Contents  [header link](class-guzzlehttp-promise-coroutine-toc.md)

#### Interfaces  [header link](class-guzzlehttp-promise-coroutine-toc-interfaces.md)

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)A promise represents the eventual result of an asynchronous operation.

#### Methods  [header link](class-guzzlehttp-promise-coroutine-toc-methods.md)

[\_\_construct()](class-guzzlehttp-promise-coroutine-method-construct.md)
: mixed [cancel()](class-guzzlehttp-promise-coroutine-method-cancel.md)
: void Cancels the promise if possible.[getState()](class-guzzlehttp-promise-coroutine-method-getstate.md)
: string Get the state of the promise ("pending", "rejected", or "fulfilled").[of()](class-guzzlehttp-promise-coroutine-method-of.md)
: self Create a new coroutine.[otherwise()](class-guzzlehttp-promise-coroutine-method-otherwise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends a rejection handler callback to the promise, and returns a new
promise resolving to the return value of the callback if it is called,
or to its original fulfillment value if the promise is instead
fulfilled.[reject()](class-guzzlehttp-promise-coroutine-method-reject.md)
: void Reject the promise with the given reason.[resolve()](class-guzzlehttp-promise-coroutine-method-resolve.md)
: void Resolve the promise with the given value.[then()](class-guzzlehttp-promise-coroutine-method-then.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends fulfillment and rejection handlers to the promise, and returns
a new promise resolving to the return value of the called handler.[wait()](class-guzzlehttp-promise-coroutine-method-wait.md)
: mixed Waits until the promise completes if possible.

### Methods  [header link](class-guzzlehttp-promise-coroutine-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-promise-coroutine-method-construct.md)

`
    public
                    __construct(callable $generatorFn) : mixed`

##### Parameters

$generatorFn
: callable

#### cancel()  [header link](class-guzzlehttp-promise-coroutine-method-cancel.md)

Cancels the promise if possible.

`
    public
                    cancel() : void`

#### getState()  [header link](class-guzzlehttp-promise-coroutine-method-getstate.md)

Get the state of the promise ("pending", "rejected", or "fulfilled").

`
    public
                    getState() : string`

The three states can be checked against the constants defined on
PromiseInterface: PENDING, FULFILLED, and REJECTED.

##### Return values

string

#### of()  [header link](class-guzzlehttp-promise-coroutine-method-of.md)

Create a new coroutine.

`
    public
            static        of(callable $generatorFn) : self`

##### Parameters

$generatorFn
: callable

##### Return values

self

#### otherwise()  [header link](class-guzzlehttp-promise-coroutine-method-otherwise.md)

Appends a rejection handler callback to the promise, and returns a new
promise resolving to the return value of the callback if it is called,
or to its original fulfillment value if the promise is instead
fulfilled.

`
    public
                    otherwise(callable $onRejected) : PromiseInterface`

##### Parameters

$onRejected
: callable

Invoked when the promise is rejected.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### reject()  [header link](class-guzzlehttp-promise-coroutine-method-reject.md)

Reject the promise with the given reason.

`
    public
                    reject(mixed $reason) : void`

##### Parameters

$reason
: mixed

#### resolve()  [header link](class-guzzlehttp-promise-coroutine-method-resolve.md)

Resolve the promise with the given value.

`
    public
                    resolve(mixed $value) : void`

##### Parameters

$value
: mixed

#### then()  [header link](class-guzzlehttp-promise-coroutine-method-then.md)

Appends fulfillment and rejection handlers to the promise, and returns
a new promise resolving to the return value of the called handler.

`
    public
                    then([callable|null $onFulfilled = null ][, callable|null $onRejected = null ]) : PromiseInterface`

##### Parameters

$onFulfilled
: callable\|null
= null

Invoked when the promise fulfills.

$onRejected
: callable\|null
= null

Invoked when the promise is rejected.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### wait()  [header link](class-guzzlehttp-promise-coroutine-method-wait.md)

Waits until the promise completes if possible.

`
    public
                    wait([bool $unwrap = true ]) : mixed`

Pass $unwrap as true to unwrap the result of the promise, either
returning the resolved value or throwing the rejected exception.

If the promise cannot be waited on, then the promise will be rejected.

##### Parameters

$unwrap
: bool
= true
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-promise-coroutine-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-promise-coroutine-method-construct.md)
  - [cancel()](class-guzzlehttp-promise-coroutine-method-cancel.md)
  - [getState()](class-guzzlehttp-promise-coroutine-method-getstate.md)
  - [of()](class-guzzlehttp-promise-coroutine-method-of.md)
  - [otherwise()](class-guzzlehttp-promise-coroutine-method-otherwise.md)
  - [reject()](class-guzzlehttp-promise-coroutine-method-reject.md)
  - [resolve()](class-guzzlehttp-promise-coroutine-method-resolve.md)
  - [then()](class-guzzlehttp-promise-coroutine-method-then.md)
  - [wait()](class-guzzlehttp-promise-coroutine-method-wait.md)

[Back To Top](class-guzzlehttp-promise-coroutine-top.md)
