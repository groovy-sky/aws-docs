Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html)

## PromiseInterface     in    - [Aws](package-aws.md)

A promise represents the eventual result of an asynchronous operation.

The primary way of interacting with a promise is through its then method,
which registers callbacks to receive either a promise’s eventual value or
the reason why the promise cannot be fulfilled.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#tags)

see[https://promisesaplus.com/](https://promisesaplus.com/)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#toc-constants)

[FULFILLED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#constant_FULFILLED)
= 'fulfilled' [PENDING](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#constant_PENDING)
= 'pending' [REJECTED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#constant_REJECTED)
= 'rejected'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#toc-methods)

[cancel()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_cancel)
: void Cancels the promise if possible.[getState()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_getState)
: string Get the state of the promise ("pending", "rejected", or "fulfilled").[otherwise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_otherwise)
: [PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)Appends a rejection handler callback to the promise, and returns a new
promise resolving to the return value of the callback if it is called,
or to its original fulfillment value if the promise is instead
fulfilled.[reject()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_reject)
: void Reject the promise with the given reason.[resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_resolve)
: void Resolve the promise with the given value.[then()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_then)
: [PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)Appends fulfillment and rejection handlers to the promise, and returns
a new promise resolving to the return value of the called handler.[wait()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_wait)
: mixed Waits until the promise completes if possible.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#constants)

#### FULFILLED  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#constant_FULFILLED)

`
    public
        mixed
    FULFILLED
    = 'fulfilled'
`

#### PENDING  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#constant_PENDING)

`
    public
        mixed
    PENDING
    = 'pending'
`

#### REJECTED  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#constant_REJECTED)

`
    public
        mixed
    REJECTED
    = 'rejected'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#methods)

#### cancel()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_cancel)

Cancels the promise if possible.

`
    public
                    cancel() : void`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_cancel\#tags)

see[https://github.com/promises-aplus/cancellation-spec/issues/7](https://github.com/promises-aplus/cancellation-spec/issues/7)

#### getState()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_getState)

Get the state of the promise ("pending", "rejected", or "fulfilled").

`
    public
                    getState() : string`

The three states can be checked against the constants defined on
PromiseInterface: PENDING, FULFILLED, and REJECTED.

##### Return values

string

#### otherwise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_otherwise)

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

[PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)

#### reject()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_reject)

Reject the promise with the given reason.

`
    public
                    reject(mixed $reason) : void`

##### Parameters

$reason
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_reject\#tags)

throwsRuntimeException

if the promise is already resolved.

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_resolve)

Resolve the promise with the given value.

`
    public
                    resolve(mixed $value) : void`

##### Parameters

$value
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_resolve\#tags)

throwsRuntimeException

if the promise is already resolved.

#### then()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_then)

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

[PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)

#### wait()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_wait)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html\#method_wait\#tags)

throwsLogicException

if the promise has no wait function or if the
promise does not settle after waiting.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#toc-methods)
- Methods
  - [cancel()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_cancel)
  - [getState()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_getState)
  - [otherwise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_otherwise)
  - [reject()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_reject)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_resolve)
  - [then()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_then)
  - [wait()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#method_wait)
- Constants
  - [FULFILLED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#constant_FULFILLED)
  - [PENDING](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#constant_PENDING)
  - [REJECTED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#constant_REJECTED)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html#top)
