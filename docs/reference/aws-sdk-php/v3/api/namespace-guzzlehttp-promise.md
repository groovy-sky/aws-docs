Menu

- [GuzzleHttp](namespace-guzzlehttp.md)

## Promise

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html\#toc-interfaces)

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)A promise represents the eventual result of an asynchronous operation.[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.[TaskQueueInterface](class-guzzlehttp-promise-taskqueueinterface.md)

#### Classes  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html\#toc-classes)

[AggregateException](class-guzzlehttp-promise-aggregateexception.md)Exception thrown when too many errors occur in the some() or any() methods.[CancellationException](class-guzzlehttp-promise-cancellationexception.md)Exception that is set as the reason for a promise that has been cancelled.[Coroutine](class-guzzlehttp-promise-coroutine.md)Creates a promise that is resolved using a generator that yields values or
promises (somewhat similar to C#'s async keyword).[Create](class-guzzlehttp-promise-create.md)[Each](class-guzzlehttp-promise-each.md)[EachPromise](class-guzzlehttp-promise-eachpromise.md)Represents a promise that iterates over many promises and invokes
side-effect functions in the process.[FulfilledPromise](class-guzzlehttp-promise-fulfilledpromise.md)A promise that has been fulfilled.[Is](class-guzzlehttp-promise-is.md)[Promise](class-guzzlehttp-promise-promise.md)Promises/A+ implementation that avoids recursion when possible.[RejectedPromise](class-guzzlehttp-promise-rejectedpromise.md)A promise that has been rejected.[RejectionException](class-guzzlehttp-promise-rejectionexception.md)A special exception that is thrown when waiting on a rejected promise.[TaskQueue](class-guzzlehttp-promise-taskqueue.md)A task queue that executes tasks in a FIFO order.[Utils](class-guzzlehttp-promise-utils.md)

```

```

×

**On this page**

- Table Of Contents
  - [Interfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html#toc-interfaces)
  - [Classes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html#toc-classes)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html#top)
