Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## TaskQueue        in package    - [Aws](package-aws.md)       implements  [TaskQueueInterface](class-guzzlehttp-promise-taskqueueinterface.md)

A task queue that executes tasks in a FIFO order.

This task queue class is used to settle promises asynchronously and
maintains a constant stack size. You can use the task queue asynchronously
by calling the `run()` function of the global task queue in an event loop.

```prettyprint
GuzzleHttp\Promise\Utils::queue()->run();

```

##### Tags  [header link](class-guzzlehttp-promise-taskqueue-tags.md)

final

### Table of Contents  [header link](class-guzzlehttp-promise-taskqueue-toc.md)

#### Interfaces  [header link](class-guzzlehttp-promise-taskqueue-toc-interfaces.md)

[TaskQueueInterface](class-guzzlehttp-promise-taskqueueinterface.md)

#### Methods  [header link](class-guzzlehttp-promise-taskqueue-toc-methods.md)

[\_\_construct()](class-guzzlehttp-promise-taskqueue-method-construct.md)
: mixed [add()](class-guzzlehttp-promise-taskqueue-method-add.md)
: void Adds a task to the queue that will be executed the next time run is
called.[disableShutdown()](class-guzzlehttp-promise-taskqueue-method-disableshutdown.md)
: void The task queue will be run and exhausted by default when the process
exits IFF the exit is not the result of a PHP E\_ERROR error.[isEmpty()](class-guzzlehttp-promise-taskqueue-method-isempty.md)
: bool Returns true if the queue is empty.[run()](class-guzzlehttp-promise-taskqueue-method-run.md)
: void Execute all of the pending task in the queue.

### Methods  [header link](class-guzzlehttp-promise-taskqueue-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-promise-taskqueue-method-construct.md)

`
    public
                    __construct([bool $withShutdown = true ]) : mixed`

##### Parameters

$withShutdown
: bool
= true

#### add()  [header link](class-guzzlehttp-promise-taskqueue-method-add.md)

Adds a task to the queue that will be executed the next time run is
called.

`
    public
                    add(callable $task) : void`

##### Parameters

$task
: callable

#### disableShutdown()  [header link](class-guzzlehttp-promise-taskqueue-method-disableshutdown.md)

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

#### isEmpty()  [header link](class-guzzlehttp-promise-taskqueue-method-isempty.md)

Returns true if the queue is empty.

`
    public
                    isEmpty() : bool`

##### Return values

bool

#### run()  [header link](class-guzzlehttp-promise-taskqueue-method-run.md)

Execute all of the pending task in the queue.

`
    public
                    run() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-promise-taskqueue-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-promise-taskqueue-method-construct.md)
  - [add()](class-guzzlehttp-promise-taskqueue-method-add.md)
  - [disableShutdown()](class-guzzlehttp-promise-taskqueue-method-disableshutdown.md)
  - [isEmpty()](class-guzzlehttp-promise-taskqueue-method-isempty.md)
  - [run()](class-guzzlehttp-promise-taskqueue-method-run.md)

[Back To Top](class-guzzlehttp-promise-taskqueue-top.md)
