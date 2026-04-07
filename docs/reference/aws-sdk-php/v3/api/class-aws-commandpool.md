Menu

- [Aws](namespace-aws.md)

## CommandPool        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromisorInterface.html)

Sends and iterator of commands concurrently using a capped pool size.

The pool will read command objects from an iterator until it is cancelled or
until the iterator is consumed.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html\#toc-interfaces)

[PromisorInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromisorInterface.html)Interface used with classes that return a promise.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html#method___construct)
: mixed The CommandPool constructor accepts a hash of configuration options:[batch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html#method_batch)
: array<string\|int, mixed> Executes a pool synchronously and aggregates the results of the pool
into an indexed array in the same order as the passed in array.[promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html\#method___construct)

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

#### batch()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html\#method_batch)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html\#method_batch\#tags)

seeCommandPool::\_\_construct

for available configuration options.

##### Return values

array<string\|int, mixed>

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html\#method_promise)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html#method___construct)
  - [batch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html#method_batch)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html#method_promise)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandPool.html#top)
