Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## PromiseInterface     in    - [Aws](package-aws.md)

A promise represents the eventual result of an asynchronous operation.

The primary way of interacting with a promise is through its then method,
which registers callbacks to receive either a promise’s eventual value or
the reason why the promise cannot be fulfilled.

##### Tags  [header link](class-guzzlehttp-promise-promiseinterface-tags.md)

see[https://promisesaplus.com/](https://promisesaplus.com/)

### Table of Contents  [header link](class-guzzlehttp-promise-promiseinterface-toc.md)

#### Constants  [header link](class-guzzlehttp-promise-promiseinterface-toc-constants.md)

[FULFILLED](class-guzzlehttp-promise-promiseinterface-constant-fulfilled.md)
= 'fulfilled' [PENDING](class-guzzlehttp-promise-promiseinterface-constant-pending.md)
= 'pending' [REJECTED](class-guzzlehttp-promise-promiseinterface-constant-rejected.md)
= 'rejected'

#### Methods  [header link](class-guzzlehttp-promise-promiseinterface-toc-methods.md)

[cancel()](class-guzzlehttp-promise-promiseinterface-method-cancel.md)
: void Cancels the promise if possible.[getState()](class-guzzlehttp-promise-promiseinterface-method-getstate.md)
: string Get the state of the promise ("pending", "rejected", or "fulfilled").[otherwise()](class-guzzlehttp-promise-promiseinterface-method-otherwise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends a rejection handler callback to the promise, and returns a new
promise resolving to the return value of the callback if it is called,
or to its original fulfillment value if the promise is instead
fulfilled.[reject()](class-guzzlehttp-promise-promiseinterface-method-reject.md)
: void Reject the promise with the given reason.[resolve()](class-guzzlehttp-promise-promiseinterface-method-resolve.md)
: void Resolve the promise with the given value.[then()](class-guzzlehttp-promise-promiseinterface-method-then.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends fulfillment and rejection handlers to the promise, and returns
a new promise resolving to the return value of the called handler.[wait()](class-guzzlehttp-promise-promiseinterface-method-wait.md)
: mixed Waits until the promise completes if possible.

### Constants  [header link](class-guzzlehttp-promise-promiseinterface-constants.md)

#### FULFILLED  [header link](class-guzzlehttp-promise-promiseinterface-constant-fulfilled.md)

`
    public
        mixed
    FULFILLED
    = 'fulfilled'
`

#### PENDING  [header link](class-guzzlehttp-promise-promiseinterface-constant-pending.md)

`
    public
        mixed
    PENDING
    = 'pending'
`

#### REJECTED  [header link](class-guzzlehttp-promise-promiseinterface-constant-rejected.md)

`
    public
        mixed
    REJECTED
    = 'rejected'
`

### Methods  [header link](class-guzzlehttp-promise-promiseinterface-methods.md)

#### cancel()  [header link](class-guzzlehttp-promise-promiseinterface-method-cancel.md)

Cancels the promise if possible.

`
    public
                    cancel() : void`

##### Tags  [header link](class-guzzlehttp-promise-promiseinterface-method-cancel-tags.md)

see[https://github.com/promises-aplus/cancellation-spec/issues/7](https://github.com/promises-aplus/cancellation-spec/issues/7)

#### getState()  [header link](class-guzzlehttp-promise-promiseinterface-method-getstate.md)

Get the state of the promise ("pending", "rejected", or "fulfilled").

`
    public
                    getState() : string`

The three states can be checked against the constants defined on
PromiseInterface: PENDING, FULFILLED, and REJECTED.

##### Return values

string

#### otherwise()  [header link](class-guzzlehttp-promise-promiseinterface-method-otherwise.md)

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

#### reject()  [header link](class-guzzlehttp-promise-promiseinterface-method-reject.md)

Reject the promise with the given reason.

`
    public
                    reject(mixed $reason) : void`

##### Parameters

$reason
: mixed

##### Tags  [header link](class-guzzlehttp-promise-promiseinterface-method-reject-tags.md)

throwsRuntimeException

if the promise is already resolved.

#### resolve()  [header link](class-guzzlehttp-promise-promiseinterface-method-resolve.md)

Resolve the promise with the given value.

`
    public
                    resolve(mixed $value) : void`

##### Parameters

$value
: mixed

##### Tags  [header link](class-guzzlehttp-promise-promiseinterface-method-resolve-tags.md)

throwsRuntimeException

if the promise is already resolved.

#### then()  [header link](class-guzzlehttp-promise-promiseinterface-method-then.md)

Appends fulfillment and rejection handlers to the promise, and returns
a new promise resolving to the return value of the called handler.

`
    public
                    then([callable $onFulfilled = null ][, callable $onRejected = null ]) : PromiseInterface`

##### Parameters

$onFulfilled
: callable
= null

Invoked when the promise fulfills.

$onRejected
: callable
= null

Invoked when the promise is rejected.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### wait()  [header link](class-guzzlehttp-promise-promiseinterface-method-wait.md)

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

##### Tags  [header link](class-guzzlehttp-promise-promiseinterface-method-wait-tags.md)

throwsLogicException

if the promise has no wait function or if the
promise does not settle after waiting.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-guzzlehttp-promise-promiseinterface-toc-constants.md)
  - [Methods](class-guzzlehttp-promise-promiseinterface-toc-methods.md)
- Methods
  - [cancel()](class-guzzlehttp-promise-promiseinterface-method-cancel.md)
  - [getState()](class-guzzlehttp-promise-promiseinterface-method-getstate.md)
  - [otherwise()](class-guzzlehttp-promise-promiseinterface-method-otherwise.md)
  - [reject()](class-guzzlehttp-promise-promiseinterface-method-reject.md)
  - [resolve()](class-guzzlehttp-promise-promiseinterface-method-resolve.md)
  - [then()](class-guzzlehttp-promise-promiseinterface-method-then.md)
  - [wait()](class-guzzlehttp-promise-promiseinterface-method-wait.md)
- Constants
  - [FULFILLED](class-guzzlehttp-promise-promiseinterface-constant-fulfilled.md)
  - [PENDING](class-guzzlehttp-promise-promiseinterface-constant-pending.md)
  - [REJECTED](class-guzzlehttp-promise-promiseinterface-constant-rejected.md)

[Back To Top](class-guzzlehttp-promise-promiseinterface-top.md)
