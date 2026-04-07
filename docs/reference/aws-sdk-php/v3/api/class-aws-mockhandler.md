Menu

- [Aws](namespace-aws.md)

## MockHandler        in package    - [Aws](package-aws.md)       implements  Countable

Returns promises that are rejected or fulfilled using a queue of
Aws\\ResultInterface and Aws\\Exception\\AwsException objects.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#toc-interfaces)

Countable

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method___construct)
: mixed The passed in value must be an array of {@see Aws\\ResultInterface} or
{@see AwsException} objects that acts as a queue of results or
exceptions to return each time the handler is invoked.[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method___invoke)
: mixed [append()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_append)
: mixed Adds one or more variadic ResultInterface or AwsException objects to the
queue.[appendException()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_appendException)
: mixed Adds one or more \\Exception or \\Throwable to the queue[count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_count)
: int Returns the number of remaining items in the queue.[getLastCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_getLastCommand)
: [CommandInterface](class-aws-commandinterface.md)Get the last received command.[getLastRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_getLastRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md) \|null Get the last received request.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#method___construct)

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

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#method___invoke)

`
    public
                    __invoke(CommandInterface $command, RequestInterface $request) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

#### append()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#method_append)

Adds one or more variadic ResultInterface or AwsException objects to the
queue.

`
    public
                    append() : mixed`

#### appendException()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#method_appendException)

Adds one or more \\Exception or \\Throwable to the queue

`
    public
                    appendException() : mixed`

#### count()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#method_count)

Returns the number of remaining items in the queue.

`
    public
                    count() : int`

##### Return values

int

#### getLastCommand()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#method_getLastCommand)

Get the last received command.

`
    public
                    getLastCommand() : CommandInterface`

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getLastRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html\#method_getLastRequest)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method___invoke)
  - [append()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_append)
  - [appendException()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_appendException)
  - [count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_count)
  - [getLastCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_getLastCommand)
  - [getLastRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#method_getLastRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MockHandler.html#top)
