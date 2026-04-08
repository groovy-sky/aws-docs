Menu

- [Aws](namespace-aws.md)

## MockHandler        in package    - [Aws](package-aws.md)       implements  Countable

Returns promises that are rejected or fulfilled using a queue of
Aws\\ResultInterface and Aws\\Exception\\AwsException objects.

### Table of Contents  [header link](class-aws-mockhandler-toc.md)

#### Interfaces  [header link](class-aws-mockhandler-toc-interfaces.md)

Countable

#### Methods  [header link](class-aws-mockhandler-toc-methods.md)

[\_\_construct()](class-aws-mockhandler-method-construct.md)
: mixed The passed in value must be an array of {@see Aws\\ResultInterface} or
{@see AwsException} objects that acts as a queue of results or
exceptions to return each time the handler is invoked.[\_\_invoke()](class-aws-mockhandler-method-invoke.md)
: mixed [append()](class-aws-mockhandler-method-append.md)
: mixed Adds one or more variadic ResultInterface or AwsException objects to the
queue.[appendException()](class-aws-mockhandler-method-appendexception.md)
: mixed Adds one or more \\Exception or \\Throwable to the queue[count()](class-aws-mockhandler-method-count.md)
: int Returns the number of remaining items in the queue.[getLastCommand()](class-aws-mockhandler-method-getlastcommand.md)
: [CommandInterface](class-aws-commandinterface.md)Get the last received command.[getLastRequest()](class-aws-mockhandler-method-getlastrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md) \|null Get the last received request.

### Methods  [header link](class-aws-mockhandler-methods.md)

#### \_\_construct()  [header link](class-aws-mockhandler-method-construct.md)

The passed in value must be an array of {@see Aws\\ResultInterface} or
{@see AwsException} objects that acts as a queue of results or
exceptions to return each time the handler is invoked.

`
    public
                    __construct([array<string|int, mixed> $resultOrQueue = [] ][, callable $onFulfilled = null ][, callable $onRejected = null ]) : mixed`

##### Parameters

$resultOrQueue
: array<string\|int, mixed>
= \[\]$onFulfilled
: callable
= null

Callback to invoke when the return value is fulfilled.

$onRejected
: callable
= null

Callback to invoke when the return value is rejected.

#### \_\_invoke()  [header link](class-aws-mockhandler-method-invoke.md)

`
    public
                    __invoke(CommandInterface $command, RequestInterface $request) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

#### append()  [header link](class-aws-mockhandler-method-append.md)

Adds one or more variadic ResultInterface or AwsException objects to the
queue.

`
    public
                    append() : mixed`

#### appendException()  [header link](class-aws-mockhandler-method-appendexception.md)

Adds one or more \\Exception or \\Throwable to the queue

`
    public
                    appendException() : mixed`

#### count()  [header link](class-aws-mockhandler-method-count.md)

Returns the number of remaining items in the queue.

`
    public
                    count() : int`

##### Return values

int

#### getLastCommand()  [header link](class-aws-mockhandler-method-getlastcommand.md)

Get the last received command.

`
    public
                    getLastCommand() : CommandInterface`

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getLastRequest()  [header link](class-aws-mockhandler-method-getlastrequest.md)

Get the last received request.

`
    public
                    getLastRequest() : RequestInterface|null`

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md) \|null
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-mockhandler-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-mockhandler-method-construct.md)
  - [\_\_invoke()](class-aws-mockhandler-method-invoke.md)
  - [append()](class-aws-mockhandler-method-append.md)
  - [appendException()](class-aws-mockhandler-method-appendexception.md)
  - [count()](class-aws-mockhandler-method-count.md)
  - [getLastCommand()](class-aws-mockhandler-method-getlastcommand.md)
  - [getLastRequest()](class-aws-mockhandler-method-getlastrequest.md)

[Back To Top](class-aws-mockhandler-top.md)
