Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.promise.html)

## TaskQueueInterface     in    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html\#toc-methods)

[add()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#method_add)
: void Adds a task to the queue that will be executed the next time run is
called.[isEmpty()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#method_isEmpty)
: bool Returns true if the queue is empty.[run()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#method_run)
: void Execute all of the pending task in the queue.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html\#methods)

#### add()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html\#method_add)

Adds a task to the queue that will be executed the next time run is
called.

`
    public
                    add(callable $task) : void`

##### Parameters

$task
: callable

#### isEmpty()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html\#method_isEmpty)

Returns true if the queue is empty.

`
    public
                    isEmpty() : bool`

##### Return values

bool

#### run()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html\#method_run)

Execute all of the pending task in the queue.

`
    public
                    run() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#toc-methods)
- Methods
  - [add()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#method_add)
  - [isEmpty()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#method_isEmpty)
  - [run()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#method_run)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.TaskQueueInterface.html#top)
