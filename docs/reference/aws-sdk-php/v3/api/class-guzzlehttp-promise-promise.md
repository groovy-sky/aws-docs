Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## Promise        in package    - [Aws](package-aws.md)       implements  [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

Promises/A+ implementation that avoids recursion when possible.

##### Tags  [header link](class-guzzlehttp-promise-promise-tags.md)

see[https://promisesaplus.com/](https://promisesaplus.com/)final

### Table of Contents  [header link](class-guzzlehttp-promise-promise-toc.md)

#### Interfaces  [header link](class-guzzlehttp-promise-promise-toc-interfaces.md)

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)A promise represents the eventual result of an asynchronous operation.

#### Methods  [header link](class-guzzlehttp-promise-promise-toc-methods.md)

[\_\_construct()](class-guzzlehttp-promise-promise-method-construct.md)
: mixed [cancel()](class-guzzlehttp-promise-promise-method-cancel.md)
: void Cancels the promise if possible.[getState()](class-guzzlehttp-promise-promise-method-getstate.md)
: string Get the state of the promise ("pending", "rejected", or "fulfilled").[otherwise()](class-guzzlehttp-promise-promise-method-otherwise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends a rejection handler callback to the promise, and returns a new
promise resolving to the return value of the callback if it is called,
or to its original fulfillment value if the promise is instead
fulfilled.[reject()](class-guzzlehttp-promise-promise-method-reject.md)
: void Reject the promise with the given reason.[resolve()](class-guzzlehttp-promise-promise-method-resolve.md)
: void Resolve the promise with the given value.[then()](class-guzzlehttp-promise-promise-method-then.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Appends fulfillment and rejection handlers to the promise, and returns
a new promise resolving to the return value of the called handler.[wait()](class-guzzlehttp-promise-promise-method-wait.md)
: mixed Waits until the promise completes if possible.

### Methods  [header link](class-guzzlehttp-promise-promise-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-promise-promise-method-construct.md)

`
    public
                    __construct([callable $waitFn = null ][, callable $cancelFn = null ]) : mixed`

##### Parameters

$waitFn
: callable
= null

Fn that when invoked resolves the promise.

$cancelFn
: callable
= null

Fn that when invoked cancels the promise.

#### cancel()  [header link](class-guzzlehttp-promise-promise-method-cancel.md)

Cancels the promise if possible.

`
    public
                    cancel() : void`

#### getState()  [header link](class-guzzlehttp-promise-promise-method-getstate.md)

Get the state of the promise ("pending", "rejected", or "fulfilled").

`
    public
                    getState() : string`

The three states can be checked against the constants defined on
PromiseInterface: PENDING, FULFILLED, and REJECTED.

##### Return values

string

#### otherwise()  [header link](class-guzzlehttp-promise-promise-method-otherwise.md)

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

#### reject()  [header link](class-guzzlehttp-promise-promise-method-reject.md)

Reject the promise with the given reason.

`
    public
                    reject(mixed $reason) : void`

##### Parameters

$reason
: mixed

#### resolve()  [header link](class-guzzlehttp-promise-promise-method-resolve.md)

Resolve the promise with the given value.

`
    public
                    resolve(mixed $value) : void`

##### Parameters

$value
: mixed

#### then()  [header link](class-guzzlehttp-promise-promise-method-then.md)

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

#### wait()  [header link](class-guzzlehttp-promise-promise-method-wait.md)

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
  - [Methods](class-guzzlehttp-promise-promise-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-promise-promise-method-construct.md)
  - [cancel()](class-guzzlehttp-promise-promise-method-cancel.md)
  - [getState()](class-guzzlehttp-promise-promise-method-getstate.md)
  - [otherwise()](class-guzzlehttp-promise-promise-method-otherwise.md)
  - [reject()](class-guzzlehttp-promise-promise-method-reject.md)
  - [resolve()](class-guzzlehttp-promise-promise-method-resolve.md)
  - [then()](class-guzzlehttp-promise-promise-method-then.md)
  - [wait()](class-guzzlehttp-promise-promise-method-wait.md)

[Back To Top](class-guzzlehttp-promise-promise-top.md)
