Menu

- [Aws](namespace-aws.md)
- [Batch](namespace-aws-batch.md)
- [Exception](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.batch.exception.html)

## BatchException     extends [AwsException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html)   in package    - [Aws](package-aws.md)

Represents an error interacting with the **AWS Batch** service.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Batch.Exception.BatchException.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Batch.Exception.BatchException.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method___construct)
: mixed [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method___toString)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[count()](class-aws-hasdatatrait.md#method_count)
: int [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_get)
: mixed [getAwsErrorCode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsErrorCode)
: string\|null Get the AWS error code.[getAwsErrorMessage()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsErrorMessage)
: string\|null Get the concise error message if any.[getAwsErrorShape()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsErrorShape)
: [Shape](class-aws-api-shape.md) \|null Get the AWS error shape.[getAwsErrorType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsErrorType)
: string\|null Get the AWS error type.[getAwsRequestId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsRequestId)
: string\|null Get the request ID of the error. This value is only present if a
response was received and is not present in the event of a networking
error.[getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)Get the command that was executed.[getIterator()](class-aws-hasdatatrait.md#method_getIterator)
: Traversable[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md) \|null Get the sent HTTP request if any.[getResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getResponse)
: [ResponseInterface](class-psr-http-message-responseinterface.md) \|null Get the received HTTP response if any.[getResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getResult)
: [ResultInterface](class-aws-resultinterface.md) \|null Get the result of the exception if available[getStatusCode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getStatusCode)
: int\|null If available, gets the HTTP status code of the corresponding response[getTransferInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getTransferInfo)
: mixed\|null\|array<string\|int, mixed> Get all transfer information as an associative array if no $name
argument is supplied, or gets a specific transfer statistic if
a $name attribute is supplied (e.g., 'retries\_attempted').[hasKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_hasKey)
: mixed [isConnectionError()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_isConnectionError)
: bool Returns true if this is a connection error.[isMaxRetriesExceeded()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_isMaxRetriesExceeded)
: bool Returns whether the max number of retries is exceeded.[offsetExists()](class-aws-hasdatatrait.md#method_offsetExists)
: bool [offsetGet()](class-aws-hasdatatrait.md#method_offsetGet)
: mixed\|null This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').[offsetSet()](class-aws-hasdatatrait.md#method_offsetSet)
: void [offsetUnset()](class-aws-hasdatatrait.md#method_offsetUnset)
: void [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list[search()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_search)
: mixed [setMaxRetriesExceeded()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_setMaxRetriesExceeded)
: mixed Sets the flag for max number of retries exceeded.[setTransferInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_setTransferInfo)
: mixed Replace the transfer information associated with an exception.[toArray()](class-aws-hasdatatrait.md#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Batch.Exception.BatchException.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method___construct)

`
    public
                    __construct(string $message, CommandInterface $command[, array<string|int, mixed> $context = [] ][, Throwable|null $previous = null ]) : mixed`

##### Parameters

$message
: string

Exception message

$command
: [CommandInterface](class-aws-commandinterface.md)$context
: array<string\|int, mixed>
= \[\]

Exception context

$previous
: Throwable\|null
= null

Previous exception (if any)

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method___toString)

`
    public
                    __toString() : mixed`

#### appendMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_appendMonitoringEvent)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### count()  [header link](class-aws-hasdatatrait.md\#method_count)

`
    public
                    count() : int`

##### Return values

int

#### get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_get)

`
    public
                    get(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### getAwsErrorCode()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getAwsErrorCode)

Get the AWS error code.

`
    public
                    getAwsErrorCode() : string|null`

##### Return values

string\|null
—

Returns null if no response was received

#### getAwsErrorMessage()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getAwsErrorMessage)

Get the concise error message if any.

`
    public
                    getAwsErrorMessage() : string|null`

##### Return values

string\|null

#### getAwsErrorShape()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getAwsErrorShape)

Get the AWS error shape.

`
    public
                    getAwsErrorShape() : Shape|null`

##### Return values

[Shape](class-aws-api-shape.md) \|null
—

Returns null if no response was received

#### getAwsErrorType()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getAwsErrorType)

Get the AWS error type.

`
    public
                    getAwsErrorType() : string|null`

##### Return values

string\|null
—

Returns null if no response was received

#### getAwsRequestId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getAwsRequestId)

Get the request ID of the error. This value is only present if a
response was received and is not present in the event of a networking
error.

`
    public
                    getAwsRequestId() : string|null`

##### Return values

string\|null
—

Returns null if no response was received

#### getCommand()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getCommand)

Get the command that was executed.

`
    public
                    getCommand() : CommandInterface`

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getIterator()  [header link](class-aws-hasdatatrait.md\#method_getIterator)

`
    public
                    getIterator() : Traversable`

##### Return values

Traversable

#### getMonitoringEvents()  [header link](class-aws-hasmonitoringeventstrait.md\#method_getMonitoringEvents)

Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.

`
    public
                    getMonitoringEvents() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getRequest)

Get the sent HTTP request if any.

`
    public
                    getRequest() : RequestInterface|null`

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md) \|null

#### getResponse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getResponse)

Get the received HTTP response if any.

`
    public
                    getResponse() : ResponseInterface|null`

##### Return values

[ResponseInterface](class-psr-http-message-responseinterface.md) \|null

#### getResult()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getResult)

Get the result of the exception if available

`
    public
                    getResult() : ResultInterface|null`

##### Return values

[ResultInterface](class-aws-resultinterface.md) \|null

#### getStatusCode()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getStatusCode)

If available, gets the HTTP status code of the corresponding response

`
    public
                    getStatusCode() : int|null`

##### Return values

int\|null

#### getTransferInfo()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_getTransferInfo)

Get all transfer information as an associative array if no $name
argument is supplied, or gets a specific transfer statistic if
a $name attribute is supplied (e.g., 'retries\_attempted').

`
    public
                    getTransferInfo([string $name = null ]) : mixed|null|array<string|int, mixed>`

##### Parameters

$name
: string
= null

Name of the transfer stat to retrieve

##### Return values

mixed\|null\|array<string\|int, mixed>

#### hasKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_hasKey)

`
    public
                    hasKey(mixed $name) : mixed`

##### Parameters

$name
: mixed

#### isConnectionError()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_isConnectionError)

Returns true if this is a connection error.

`
    public
                    isConnectionError() : bool`

##### Return values

bool

#### isMaxRetriesExceeded()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_isMaxRetriesExceeded)

Returns whether the max number of retries is exceeded.

`
    public
                    isMaxRetriesExceeded() : bool`

##### Return values

bool

#### offsetExists()  [header link](class-aws-hasdatatrait.md\#method_offsetExists)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](class-aws-hasdatatrait.md\#method_offsetGet)

This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').

`
    public
                &    offsetGet( $offset) : mixed|null`

##### Parameters

$offset
:

##### Return values

mixed\|null

#### offsetSet()  [header link](class-aws-hasdatatrait.md\#method_offsetSet)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](class-aws-hasdatatrait.md\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### prependMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_prependMonitoringEvent)

Prepend a client-side monitoring event to this object's event list

`
    public
                    prependMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### search()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_search)

`
    public
                    search(mixed $expression) : mixed`

##### Parameters

$expression
: mixed

#### setMaxRetriesExceeded()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_setMaxRetriesExceeded)

Sets the flag for max number of retries exceeded.

`
    public
                    setMaxRetriesExceeded() : mixed`

#### setTransferInfo()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html\#method_setTransferInfo)

Replace the transfer information associated with an exception.

`
    public
                    setTransferInfo(array<string|int, mixed> $info) : mixed`

##### Parameters

$info
: array<string\|int, mixed>

#### toArray()  [header link](class-aws-hasdatatrait.md\#method_toArray)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Batch.Exception.BatchException.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method___construct)
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method___toString)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [count()](class-aws-hasdatatrait.md#method_count)
  - [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_get)
  - [getAwsErrorCode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsErrorCode)
  - [getAwsErrorMessage()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsErrorMessage)
  - [getAwsErrorShape()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsErrorShape)
  - [getAwsErrorType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsErrorType)
  - [getAwsRequestId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getAwsRequestId)
  - [getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getCommand)
  - [getIterator()](class-aws-hasdatatrait.md#method_getIterator)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getRequest)
  - [getResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getResponse)
  - [getResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getResult)
  - [getStatusCode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getStatusCode)
  - [getTransferInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_getTransferInfo)
  - [hasKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_hasKey)
  - [isConnectionError()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_isConnectionError)
  - [isMaxRetriesExceeded()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_isMaxRetriesExceeded)
  - [offsetExists()](class-aws-hasdatatrait.md#method_offsetExists)
  - [offsetGet()](class-aws-hasdatatrait.md#method_offsetGet)
  - [offsetSet()](class-aws-hasdatatrait.md#method_offsetSet)
  - [offsetUnset()](class-aws-hasdatatrait.md#method_offsetUnset)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
  - [search()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_search)
  - [setMaxRetriesExceeded()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_setMaxRetriesExceeded)
  - [setTransferInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.AwsException.html#method_setTransferInfo)
  - [toArray()](class-aws-hasdatatrait.md#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Batch.Exception.BatchException.html#top)
