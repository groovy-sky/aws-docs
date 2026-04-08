Menu

- [Aws](namespace-aws.md)

## History        in package    - [Aws](package-aws.md)       implements  Countable, IteratorAggregate

Represents a history container that is required when using the history
middleware.

### Table of Contents  [header link](class-aws-history-toc.md)

#### Interfaces  [header link](class-aws-history-toc-interfaces.md)

CountableIteratorAggregate

#### Methods  [header link](class-aws-history-toc-methods.md)

[\_\_construct()](class-aws-history-method-construct.md)
: mixed [clear()](class-aws-history-method-clear.md)
: mixed Flush the history[count()](class-aws-history-method-count.md)
: int [finish()](class-aws-history-method-finish.md)
: mixed Finish adding an entry to the history container.[getIterator()](class-aws-history-method-getiterator.md)
: mixed [getLastCommand()](class-aws-history-method-getlastcommand.md)
: [CommandInterface](class-aws-commandinterface.md)Get the last finished command seen by the history container.[getLastRequest()](class-aws-history-method-getlastrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Get the last finished request seen by the history container.[getLastReturn()](class-aws-history-method-getlastreturn.md)
: [ResultInterface](class-aws-resultinterface.md) \| [AwsException](class-aws-exception-awsexception.md)Get the last received result or exception.[start()](class-aws-history-method-start.md)
: string Initiate an entry being added to the history.[toArray()](class-aws-history-method-toarray.md)
: array<string\|int, mixed> Converts the history to an array.

### Methods  [header link](class-aws-history-methods.md)

#### \_\_construct()  [header link](class-aws-history-method-construct.md)

`
    public
                    __construct([int $maxEntries = 10 ]) : mixed`

##### Parameters

$maxEntries
: int
= 10

Maximum number of entries to store.

#### clear()  [header link](class-aws-history-method-clear.md)

Flush the history

`
    public
                    clear() : mixed`

#### count()  [header link](class-aws-history-method-count.md)

`
    public
                    count() : int`

##### Return values

int

#### finish()  [header link](class-aws-history-method-finish.md)

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

#### getIterator()  [header link](class-aws-history-method-getiterator.md)

`
    public
                    getIterator() : mixed`

#### getLastCommand()  [header link](class-aws-history-method-getlastcommand.md)

Get the last finished command seen by the history container.

`
    public
                    getLastCommand() : CommandInterface`

##### Tags  [header link](class-aws-history-method-getlastcommand-tags.md)

throwsLogicException

if no commands have been seen.

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getLastRequest()  [header link](class-aws-history-method-getlastrequest.md)

Get the last finished request seen by the history container.

`
    public
                    getLastRequest() : RequestInterface`

##### Tags  [header link](class-aws-history-method-getlastrequest-tags.md)

throwsLogicException

if no requests have been seen.

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### getLastReturn()  [header link](class-aws-history-method-getlastreturn.md)

Get the last received result or exception.

`
    public
                    getLastReturn() : ResultInterface|AwsException`

##### Tags  [header link](class-aws-history-method-getlastreturn-tags.md)

throwsLogicException

if no return values have been received.

##### Return values

[ResultInterface](class-aws-resultinterface.md) \| [AwsException](class-aws-exception-awsexception.md)

#### start()  [header link](class-aws-history-method-start.md)

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

#### toArray()  [header link](class-aws-history-method-toarray.md)

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
  - [Methods](class-aws-history-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-history-method-construct.md)
  - [clear()](class-aws-history-method-clear.md)
  - [count()](class-aws-history-method-count.md)
  - [finish()](class-aws-history-method-finish.md)
  - [getIterator()](class-aws-history-method-getiterator.md)
  - [getLastCommand()](class-aws-history-method-getlastcommand.md)
  - [getLastRequest()](class-aws-history-method-getlastrequest.md)
  - [getLastReturn()](class-aws-history-method-getlastreturn.md)
  - [start()](class-aws-history-method-start.md)
  - [toArray()](class-aws-history-method-toarray.md)

[Back To Top](class-aws-history-top.md)
