Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html)

## Create        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html\#toc-methods)

[exceptionFor()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#method_exceptionFor)
: ThrowableCreate an exception for a rejected promise value.[iterFor()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#method_iterFor)
: IteratorReturns an iterator for the given value.[promiseFor()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#method_promiseFor)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Creates a promise for a value if the value is not a promise.[rejectionFor()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#method_rejectionFor)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Creates a rejected promise for a reason if the reason is not a promise.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html\#methods)

#### exceptionFor()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html\#method_exceptionFor)

Create an exception for a rejected promise value.

`
    public
            static        exceptionFor(mixed $reason) : Throwable`

##### Parameters

$reason
: mixed

##### Return values

Throwable

#### iterFor()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html\#method_iterFor)

Returns an iterator for the given value.

`
    public
            static        iterFor(mixed $value) : Iterator`

##### Parameters

$value
: mixed

##### Return values

Iterator

#### promiseFor()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html\#method_promiseFor)

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

#### rejectionFor()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html\#method_rejectionFor)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#toc-methods)
- Methods
  - [exceptionFor()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#method_exceptionFor)
  - [iterFor()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#method_iterFor)
  - [promiseFor()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#method_promiseFor)
  - [rejectionFor()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#method_rejectionFor)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Create.html#top)
