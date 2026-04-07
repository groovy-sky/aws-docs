Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html)

## TaskQueue        in package    - [Aws](package-aws.md)       implements  [TaskQueueInterface](class-guzzlehttp-promise-taskqueueinterface.md)

A task queue that executes tasks in a FIFO order.

This task queue class is used to settle promises asynchronously and
maintains a constant stack size. You can use the task queue asynchronously
by calling the `run()` function of the global task queue in an event loop.

```prettyprint
GuzzleHttp\Promise\Utils::queue()->run();

```

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#tags)

final

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#toc-interfaces)

[TaskQueueInterface](class-guzzlehttp-promise-taskqueueinterface.md)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method___construct)
: mixed [add()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method_add)
: void Adds a task to the queue that will be executed the next time run is
called.[disableShutdown()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method_disableShutdown)
: void The task queue will be run and exhausted by default when the process
exits IFF the exit is not the result of a PHP E\_ERROR error.[isEmpty()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method_isEmpty)
: bool Returns true if the queue is empty.[run()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method_run)
: void Execute all of the pending task in the queue.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#method___construct)

`
    public
                    __construct([bool $withShutdown = true ]) : mixed`

##### Parameters

$withShutdown
: bool
= true

#### add()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#method_add)

Adds a task to the queue that will be executed the next time run is
called.

`
    public
                    add(callable $task) : void`

##### Parameters

$task
: callable

#### disableShutdown()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#method_disableShutdown)

The task queue will be run and exhausted by default when the process
exits IFF the exit is not the result of a PHP E\_ERROR error.

`
    public
                    disableShutdown() : void`

You can disable running the automatic shutdown of the queue by calling
this function. If you disable the task queue shutdown process, then you
MUST either run the task queue (as a result of running your event loop
or manually using the run() method) or wait on each outstanding promise.

Note: This shutdown will occur before any destructors are triggered.

#### isEmpty()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#method_isEmpty)

Returns true if the queue is empty.

`
    public
                    isEmpty() : bool`

##### Return values

bool

#### run()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html\#method_run)

Execute all of the pending task in the queue.

`
    public
                    run() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method___construct)
  - [add()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method_add)
  - [disableShutdown()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method_disableShutdown)
  - [isEmpty()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method_isEmpty)
  - [run()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#method_run)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueue.html#top)
