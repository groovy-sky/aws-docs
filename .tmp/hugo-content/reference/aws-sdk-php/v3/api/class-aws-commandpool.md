Menu

- [Aws](namespace-aws.md)

## CommandPool        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Sends and iterator of commands concurrently using a capped pool size.

The pool will read command objects from an iterator until it is cancelled or
until the iterator is consumed.

### Table of Contents  [header link](class-aws-commandpool-toc.md)

#### Interfaces  [header link](class-aws-commandpool-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](class-aws-commandpool-toc-methods.md)

[\_\_construct()](class-aws-commandpool-method-construct.md)
: mixed The CommandPool constructor accepts a hash of configuration options:[batch()](class-aws-commandpool-method-batch.md)
: array<string\|int, mixed> Executes a pool synchronously and aggregates the results of the pool
into an indexed array in the same order as the passed in array.[promise()](class-aws-commandpool-method-promise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Methods  [header link](class-aws-commandpool-methods.md)

#### \_\_construct()  [header link](class-aws-commandpool-method-construct.md)

The CommandPool constructor accepts a hash of configuration options:

`
    public
                    __construct(AwsClientInterface $client, array<string|int, mixed>|Iterator $commands[, array<string|int, mixed> $config = [] ]) : mixed`

- concurrency: (callable\|int) Maximum number of commands to execute
concurrently. Provide a function to resize the pool dynamically. The
function will be provided the current number of pending requests and
is expected to return an integer representing the new pool size limit.
- before: (callable) function to invoke before sending each command. The
before function accepts the command and the key of the iterator of the
command. You can mutate the command as needed in the before function
before sending the command.
- fulfilled: (callable) Function to invoke when a promise is fulfilled.
The function is provided the result object, id of the iterator that the
result came from, and the aggregate promise that can be resolved/rejected
if you need to short-circuit the pool.
- rejected: (callable) Function to invoke when a promise is rejected.
The function is provided an AwsException object, id of the iterator that
the exception came from, and the aggregate promise that can be
resolved/rejected if you need to short-circuit the pool.
- preserve\_iterator\_keys: (bool) Retain the iterator key when generating
the commands.

##### Parameters

$client
: [AwsClientInterface](class-aws-awsclientinterface.md)

Client used to execute commands.

$commands
: array<string\|int, mixed>\|Iterator

Iterable that yields commands.

$config
: array<string\|int, mixed>
= \[\]

Associative array of options.

#### batch()  [header link](class-aws-commandpool-method-batch.md)

Executes a pool synchronously and aggregates the results of the pool
into an indexed array in the same order as the passed in array.

`
    public
            static        batch(AwsClientInterface $client, mixed $commands[, array<string|int, mixed> $config = [] ]) : array<string|int, mixed>`

##### Parameters

$client
: [AwsClientInterface](class-aws-awsclientinterface.md)

Client used to execute commands.

$commands
: mixed

Iterable that yields commands.

$config
: array<string\|int, mixed>
= \[\]

Configuration options.

##### Tags  [header link](class-aws-commandpool-method-batch-tags.md)

seeCommandPool::\_\_construct

for available configuration options.

##### Return values

array<string\|int, mixed>

#### promise()  [header link](class-aws-commandpool-method-promise.md)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-commandpool-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-commandpool-method-construct.md)
  - [batch()](class-aws-commandpool-method-batch.md)
  - [promise()](class-aws-commandpool-method-promise.md)

[Back To Top](class-aws-commandpool-top.md)
