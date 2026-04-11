Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Promise](namespace-guzzlehttp-promise.md)

## TaskQueueInterface     in    - [Aws](package-aws.md)

### Table of Contents  [header link](class-guzzlehttp-promise-taskqueueinterface-toc.md)

#### Methods  [header link](class-guzzlehttp-promise-taskqueueinterface-toc-methods.md)

[add()](class-guzzlehttp-promise-taskqueueinterface-method-add.md)
: void Adds a task to the queue that will be executed the next time run is
called.[isEmpty()](class-guzzlehttp-promise-taskqueueinterface-method-isempty.md)
: bool Returns true if the queue is empty.[run()](class-guzzlehttp-promise-taskqueueinterface-method-run.md)
: void Execute all of the pending task in the queue.

### Methods  [header link](class-guzzlehttp-promise-taskqueueinterface-methods.md)

#### add()  [header link](class-guzzlehttp-promise-taskqueueinterface-method-add.md)

Adds a task to the queue that will be executed the next time run is
called.

`
    public
                    add(callable $task) : void`

##### Parameters

$task
: callable

#### isEmpty()  [header link](class-guzzlehttp-promise-taskqueueinterface-method-isempty.md)

Returns true if the queue is empty.

`
    public
                    isEmpty() : bool`

##### Return values

bool

#### run()  [header link](class-guzzlehttp-promise-taskqueueinterface-method-run.md)

Execute all of the pending task in the queue.

`
    public
                    run() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-guzzlehttp-promise-taskqueueinterface-toc-constants.md)
  - [Methods](class-guzzlehttp-promise-taskqueueinterface-toc-methods.md)
- Methods
  - [add()](class-guzzlehttp-promise-taskqueueinterface-method-add.md)
  - [isEmpty()](class-guzzlehttp-promise-taskqueueinterface-method-isempty.md)
  - [run()](class-guzzlehttp-promise-taskqueueinterface-method-run.md)

[Back To Top](class-guzzlehttp-promise-taskqueueinterface-top.md)
