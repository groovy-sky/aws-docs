Menu

- [Aws](namespace-aws.md)

## History        in package    - [Aws](package-aws.md)       implements  Countable, IteratorAggregate

Represents a history container that is required when using the history
middleware.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#toc-interfaces)

CountableIteratorAggregate

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method___construct)
: mixed [clear()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_clear)
: mixed Flush the history[count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_count)
: int [finish()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_finish)
: mixed Finish adding an entry to the history container.[getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_getIterator)
: mixed [getLastCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_getLastCommand)
: [CommandInterface](class-aws-commandinterface.md)Get the last finished command seen by the history container.[getLastRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_getLastRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)Get the last finished request seen by the history container.[getLastReturn()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_getLastReturn)
: [ResultInterface](class-aws-resultinterface.md) \| [AwsException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html)Get the last received result or exception.[start()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_start)
: string Initiate an entry being added to the history.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_toArray)
: array<string\|int, mixed> Converts the history to an array.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method___construct)

`
    public
                    __construct([int $maxEntries = 10 ]) : mixed`

##### Parameters

$maxEntries
: int
= 10

Maximum number of entries to store.

#### clear()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_clear)

Flush the history

`
    public
                    clear() : mixed`

#### count()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_count)

`
    public
                    count() : int`

##### Return values

int

#### finish()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_finish)

Finish adding an entry to the history container.

`
    public
                    finish(string $ticket, mixed $result) : mixed`

##### Parameters

$ticket
: string

Ticket returned from the start call.

$result
: mixed

The result (an exception or AwsResult).

#### getIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_getIterator)

`
    public
                    getIterator() : mixed`

#### getLastCommand()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_getLastCommand)

Get the last finished command seen by the history container.

`
    public
                    getLastCommand() : CommandInterface`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_getLastCommand\#tags)

throwsLogicException

if no commands have been seen.

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getLastRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_getLastRequest)

Get the last finished request seen by the history container.

`
    public
                    getLastRequest() : RequestInterface`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_getLastRequest\#tags)

throwsLogicException

if no requests have been seen.

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### getLastReturn()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_getLastReturn)

Get the last received result or exception.

`
    public
                    getLastReturn() : ResultInterface|AwsException`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_getLastReturn\#tags)

throwsLogicException

if no return values have been received.

##### Return values

[ResultInterface](class-aws-resultinterface.md) \| [AwsException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html)

#### start()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_start)

Initiate an entry being added to the history.

`
    public
                    start(CommandInterface $cmd, RequestInterface $req) : string`

##### Parameters

$cmd
: [CommandInterface](class-aws-commandinterface.md)

Command be executed.

$req
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request being sent.

##### Return values

string
—

Returns the ticket used to finish the entry.

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html\#method_toArray)

Converts the history to an array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method___construct)
  - [clear()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_clear)
  - [count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_count)
  - [finish()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_finish)
  - [getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_getIterator)
  - [getLastCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_getLastCommand)
  - [getLastRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_getLastRequest)
  - [getLastReturn()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_getLastReturn)
  - [start()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_start)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.History.html#top)
