Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html)

## RejectedPromise        in package    - [Aws](package-aws.md)       implements  [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

A promise that has been rejected.

Thenning off of this promise will invoke the onRejected callback
immediately and ignore other callbacks.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#tags)

final

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#toc-interfaces)

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)A promise represents the eventual result of an asynchronous operation.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method___construct)
: mixed [cancel()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_cancel)
: void Cancels the promise if possible.[getState()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_getState)
: string Get the state of the promise ("pending", "rejected", or "fulfilled").[otherwise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_otherwise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends a rejection handler callback to the promise, and returns a new
promise resolving to the return value of the callback if it is called,
or to its original fulfillment value if the promise is instead
fulfilled.[reject()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_reject)
: void Reject the promise with the given reason.[resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_resolve)
: void Resolve the promise with the given value.[then()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_then)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends fulfillment and rejection handlers to the promise, and returns
a new promise resolving to the return value of the called handler.[wait()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_wait)
: mixed Waits until the promise completes if possible.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#method___construct)

`
    public
                    __construct(mixed $reason) : mixed`

##### Parameters

$reason
: mixed

#### cancel()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#method_cancel)

Cancels the promise if possible.

`
    public
                    cancel() : void`

#### getState()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#method_getState)

Get the state of the promise ("pending", "rejected", or "fulfilled").

`
    public
                    getState() : string`

The three states can be checked against the constants defined on
PromiseInterface: PENDING, FULFILLED, and REJECTED.

##### Return values

string

#### otherwise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#method_otherwise)

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

#### reject()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#method_reject)

Reject the promise with the given reason.

`
    public
                    reject(mixed $reason) : void`

##### Parameters

$reason
: mixed

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#method_resolve)

Resolve the promise with the given value.

`
    public
                    resolve(mixed $value) : void`

##### Parameters

$value
: mixed

#### then()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#method_then)

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

#### wait()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html\#method_wait)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method___construct)
  - [cancel()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_cancel)
  - [getState()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_getState)
  - [otherwise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_otherwise)
  - [reject()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_reject)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_resolve)
  - [then()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_then)
  - [wait()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#method_wait)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.RejectedPromise.html#top)
