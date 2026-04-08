Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## RejectedPromise        in package    - [Aws](package-aws.md)       implements  [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

A promise that has been rejected.

Thenning off of this promise will invoke the onRejected callback
immediately and ignore other callbacks.

##### Tags  [header link](class-guzzlehttp-promise-rejectedpromise-tags.md)

final

### Table of Contents  [header link](class-guzzlehttp-promise-rejectedpromise-toc.md)

#### Interfaces  [header link](class-guzzlehttp-promise-rejectedpromise-toc-interfaces.md)

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)A promise represents the eventual result of an asynchronous operation.

#### Methods  [header link](class-guzzlehttp-promise-rejectedpromise-toc-methods.md)

[\_\_construct()](class-guzzlehttp-promise-rejectedpromise-method-construct.md)
: mixed [cancel()](class-guzzlehttp-promise-rejectedpromise-method-cancel.md)
: void Cancels the promise if possible.[getState()](class-guzzlehttp-promise-rejectedpromise-method-getstate.md)
: string Get the state of the promise ("pending", "rejected", or "fulfilled").[otherwise()](class-guzzlehttp-promise-rejectedpromise-method-otherwise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends a rejection handler callback to the promise, and returns a new
promise resolving to the return value of the callback if it is called,
or to its original fulfillment value if the promise is instead
fulfilled.[reject()](class-guzzlehttp-promise-rejectedpromise-method-reject.md)
: void Reject the promise with the given reason.[resolve()](class-guzzlehttp-promise-rejectedpromise-method-resolve.md)
: void Resolve the promise with the given value.[then()](class-guzzlehttp-promise-rejectedpromise-method-then.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends fulfillment and rejection handlers to the promise, and returns
a new promise resolving to the return value of the called handler.[wait()](class-guzzlehttp-promise-rejectedpromise-method-wait.md)
: mixed Waits until the promise completes if possible.

### Methods  [header link](class-guzzlehttp-promise-rejectedpromise-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-promise-rejectedpromise-method-construct.md)

`
    public
                    __construct(mixed $reason) : mixed`

##### Parameters

$reason
: mixed

#### cancel()  [header link](class-guzzlehttp-promise-rejectedpromise-method-cancel.md)

Cancels the promise if possible.

`
    public
                    cancel() : void`

#### getState()  [header link](class-guzzlehttp-promise-rejectedpromise-method-getstate.md)

Get the state of the promise ("pending", "rejected", or "fulfilled").

`
    public
                    getState() : string`

The three states can be checked against the constants defined on
PromiseInterface: PENDING, FULFILLED, and REJECTED.

##### Return values

string

#### otherwise()  [header link](class-guzzlehttp-promise-rejectedpromise-method-otherwise.md)

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

#### reject()  [header link](class-guzzlehttp-promise-rejectedpromise-method-reject.md)

Reject the promise with the given reason.

`
    public
                    reject(mixed $reason) : void`

##### Parameters

$reason
: mixed

#### resolve()  [header link](class-guzzlehttp-promise-rejectedpromise-method-resolve.md)

Resolve the promise with the given value.

`
    public
                    resolve(mixed $value) : void`

##### Parameters

$value
: mixed

#### then()  [header link](class-guzzlehttp-promise-rejectedpromise-method-then.md)

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

#### wait()  [header link](class-guzzlehttp-promise-rejectedpromise-method-wait.md)

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
  - [Methods](class-guzzlehttp-promise-rejectedpromise-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-promise-rejectedpromise-method-construct.md)
  - [cancel()](class-guzzlehttp-promise-rejectedpromise-method-cancel.md)
  - [getState()](class-guzzlehttp-promise-rejectedpromise-method-getstate.md)
  - [otherwise()](class-guzzlehttp-promise-rejectedpromise-method-otherwise.md)
  - [reject()](class-guzzlehttp-promise-rejectedpromise-method-reject.md)
  - [resolve()](class-guzzlehttp-promise-rejectedpromise-method-resolve.md)
  - [then()](class-guzzlehttp-promise-rejectedpromise-method-then.md)
  - [wait()](class-guzzlehttp-promise-rejectedpromise-method-wait.md)

[Back To Top](class-guzzlehttp-promise-rejectedpromise-top.md)
