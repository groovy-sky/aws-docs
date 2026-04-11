Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## Create        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-guzzlehttp-promise-create-toc.md)

#### Methods  [header link](class-guzzlehttp-promise-create-toc-methods.md)

[exceptionFor()](class-guzzlehttp-promise-create-method-exceptionfor.md)
: ThrowableCreate an exception for a rejected promise value.[iterFor()](class-guzzlehttp-promise-create-method-iterfor.md)
: IteratorReturns an iterator for the given value.[promiseFor()](class-guzzlehttp-promise-create-method-promisefor.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Creates a promise for a value if the value is not a promise.[rejectionFor()](class-guzzlehttp-promise-create-method-rejectionfor.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Creates a rejected promise for a reason if the reason is not a promise.

### Methods  [header link](class-guzzlehttp-promise-create-methods.md)

#### exceptionFor()  [header link](class-guzzlehttp-promise-create-method-exceptionfor.md)

Create an exception for a rejected promise value.

`
    public
            static        exceptionFor(mixed $reason) : Throwable`

##### Parameters

$reason
: mixed

##### Return values

Throwable

#### iterFor()  [header link](class-guzzlehttp-promise-create-method-iterfor.md)

Returns an iterator for the given value.

`
    public
            static        iterFor(mixed $value) : Iterator`

##### Parameters

$value
: mixed

##### Return values

Iterator

#### promiseFor()  [header link](class-guzzlehttp-promise-create-method-promisefor.md)

Creates a promise for a value if the value is not a promise.

`
    public
            static        promiseFor(mixed $value) : PromiseInterface`

##### Parameters

$value
: mixed

Promise or value.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### rejectionFor()  [header link](class-guzzlehttp-promise-create-method-rejectionfor.md)

Creates a rejected promise for a reason if the reason is not a promise.

`
    public
            static        rejectionFor(mixed $reason) : PromiseInterface`

If the provided reason is a promise, then it is returned as-is.

##### Parameters

$reason
: mixed

Promise or reason.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-promise-create-toc-methods.md)
- Methods
  - [exceptionFor()](class-guzzlehttp-promise-create-method-exceptionfor.md)
  - [iterFor()](class-guzzlehttp-promise-create-method-iterfor.md)
  - [promiseFor()](class-guzzlehttp-promise-create-method-promisefor.md)
  - [rejectionFor()](class-guzzlehttp-promise-create-method-rejectionfor.md)

[Back To Top](class-guzzlehttp-promise-create-top.md)
