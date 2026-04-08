Menu

- [Aws](namespace-aws.md)
- [ApplicationInsights](namespace-aws-applicationinsights.md)
- [Exception](namespace-aws-applicationinsights-exception.md)

## ApplicationInsightsException     extends [AwsException](class-aws-exception-awsexception.md)   in package    - [Aws](package-aws.md)

Represents an error interacting with the **Amazon CloudWatch Application Insights** service.

### Table of Contents  [header link](class-aws-applicationinsights-exception-applicationinsightsexception-toc.md)

#### Methods  [header link](class-aws-applicationinsights-exception-applicationinsightsexception-toc-methods.md)

[\_\_construct()](class-aws-exception-awsexception-method-construct.md)
: mixed [\_\_toString()](class-aws-exception-awsexception-method-tostring.md)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[count()](class-aws-hasdatatrait.md#method_count)
: int [get()](class-aws-exception-awsexception-method-get.md)
: mixed [getAwsErrorCode()](class-aws-exception-awsexception-method-getawserrorcode.md)
: string\|null Get the AWS error code.[getAwsErrorMessage()](class-aws-exception-awsexception-method-getawserrormessage.md)
: string\|null Get the concise error message if any.[getAwsErrorShape()](class-aws-exception-awsexception-method-getawserrorshape.md)
: [Shape](class-aws-api-shape.md) \|null Get the AWS error shape.[getAwsErrorType()](class-aws-exception-awsexception-method-getawserrortype.md)
: string\|null Get the AWS error type.[getAwsRequestId()](class-aws-exception-awsexception-method-getawsrequestid.md)
: string\|null Get the request ID of the error. This value is only present if a
response was received and is not present in the event of a networking
error.[getCommand()](class-aws-exception-awsexception-method-getcommand.md)
: [CommandInterface](class-aws-commandinterface.md)Get the command that was executed.[getIterator()](class-aws-hasdatatrait.md#method_getIterator)
: Traversable[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getRequest()](class-aws-exception-awsexception-method-getrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md) \|null Get the sent HTTP request if any.[getResponse()](class-aws-exception-awsexception-method-getresponse.md)
: [ResponseInterface](class-psr-http-message-responseinterface.md) \|null Get the received HTTP response if any.[getResult()](class-aws-exception-awsexception-method-getresult.md)
: [ResultInterface](class-aws-resultinterface.md) \|null Get the result of the exception if available[getStatusCode()](class-aws-exception-awsexception-method-getstatuscode.md)
: int\|null If available, gets the HTTP status code of the corresponding response[getTransferInfo()](class-aws-exception-awsexception-method-gettransferinfo.md)
: mixed\|null\|array<string\|int, mixed> Get all transfer information as an associative array if no $name
argument is supplied, or gets a specific transfer statistic if
a $name attribute is supplied (e.g., 'retries\_attempted').[hasKey()](class-aws-exception-awsexception-method-haskey.md)
: mixed [isConnectionError()](class-aws-exception-awsexception-method-isconnectionerror.md)
: bool Returns true if this is a connection error.[isMaxRetriesExceeded()](class-aws-exception-awsexception-method-ismaxretriesexceeded.md)
: bool Returns whether the max number of retries is exceeded.[offsetExists()](class-aws-hasdatatrait.md#method_offsetExists)
: bool [offsetGet()](class-aws-hasdatatrait.md#method_offsetGet)
: mixed\|null This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').[offsetSet()](class-aws-hasdatatrait.md#method_offsetSet)
: void [offsetUnset()](class-aws-hasdatatrait.md#method_offsetUnset)
: void [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list[search()](class-aws-exception-awsexception-method-search.md)
: mixed [setMaxRetriesExceeded()](class-aws-exception-awsexception-method-setmaxretriesexceeded.md)
: mixed Sets the flag for max number of retries exceeded.[setTransferInfo()](class-aws-exception-awsexception-method-settransferinfo.md)
: mixed Replace the transfer information associated with an exception.[toArray()](class-aws-hasdatatrait.md#method_toArray)
: mixed

### Methods  [header link](class-aws-applicationinsights-exception-applicationinsightsexception-methods.md)

#### \_\_construct()  [header link](class-aws-exception-awsexception-method-construct.md)

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

#### \_\_toString()  [header link](class-aws-exception-awsexception-method-tostring.md)

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

#### get()  [header link](class-aws-exception-awsexception-method-get.md)

`
    public
                    get(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### getAwsErrorCode()  [header link](class-aws-exception-awsexception-method-getawserrorcode.md)

Get the AWS error code.

`
    public
                    getAwsErrorCode() : string|null`

##### Return values

string\|null
—

Returns null if no response was received

#### getAwsErrorMessage()  [header link](class-aws-exception-awsexception-method-getawserrormessage.md)

Get the concise error message if any.

`
    public
                    getAwsErrorMessage() : string|null`

##### Return values

string\|null

#### getAwsErrorShape()  [header link](class-aws-exception-awsexception-method-getawserrorshape.md)

Get the AWS error shape.

`
    public
                    getAwsErrorShape() : Shape|null`

##### Return values

[Shape](class-aws-api-shape.md) \|null
—

Returns null if no response was received

#### getAwsErrorType()  [header link](class-aws-exception-awsexception-method-getawserrortype.md)

Get the AWS error type.

`
    public
                    getAwsErrorType() : string|null`

##### Return values

string\|null
—

Returns null if no response was received

#### getAwsRequestId()  [header link](class-aws-exception-awsexception-method-getawsrequestid.md)

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

#### getCommand()  [header link](class-aws-exception-awsexception-method-getcommand.md)

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

#### getRequest()  [header link](class-aws-exception-awsexception-method-getrequest.md)

Get the sent HTTP request if any.

`
    public
                    getRequest() : RequestInterface|null`

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md) \|null

#### getResponse()  [header link](class-aws-exception-awsexception-method-getresponse.md)

Get the received HTTP response if any.

`
    public
                    getResponse() : ResponseInterface|null`

##### Return values

[ResponseInterface](class-psr-http-message-responseinterface.md) \|null

#### getResult()  [header link](class-aws-exception-awsexception-method-getresult.md)

Get the result of the exception if available

`
    public
                    getResult() : ResultInterface|null`

##### Return values

[ResultInterface](class-aws-resultinterface.md) \|null

#### getStatusCode()  [header link](class-aws-exception-awsexception-method-getstatuscode.md)

If available, gets the HTTP status code of the corresponding response

`
    public
                    getStatusCode() : int|null`

##### Return values

int\|null

#### getTransferInfo()  [header link](class-aws-exception-awsexception-method-gettransferinfo.md)

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

#### hasKey()  [header link](class-aws-exception-awsexception-method-haskey.md)

`
    public
                    hasKey(mixed $name) : mixed`

##### Parameters

$name
: mixed

#### isConnectionError()  [header link](class-aws-exception-awsexception-method-isconnectionerror.md)

Returns true if this is a connection error.

`
    public
                    isConnectionError() : bool`

##### Return values

bool

#### isMaxRetriesExceeded()  [header link](class-aws-exception-awsexception-method-ismaxretriesexceeded.md)

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

#### search()  [header link](class-aws-exception-awsexception-method-search.md)

`
    public
                    search(mixed $expression) : mixed`

##### Parameters

$expression
: mixed

#### setMaxRetriesExceeded()  [header link](class-aws-exception-awsexception-method-setmaxretriesexceeded.md)

Sets the flag for max number of retries exceeded.

`
    public
                    setMaxRetriesExceeded() : mixed`

#### setTransferInfo()  [header link](class-aws-exception-awsexception-method-settransferinfo.md)

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
  - [Methods](class-aws-applicationinsights-exception-applicationinsightsexception-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-exception-awsexception-method-construct.md)
  - [\_\_toString()](class-aws-exception-awsexception-method-tostring.md)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [count()](class-aws-hasdatatrait.md#method_count)
  - [get()](class-aws-exception-awsexception-method-get.md)
  - [getAwsErrorCode()](class-aws-exception-awsexception-method-getawserrorcode.md)
  - [getAwsErrorMessage()](class-aws-exception-awsexception-method-getawserrormessage.md)
  - [getAwsErrorShape()](class-aws-exception-awsexception-method-getawserrorshape.md)
  - [getAwsErrorType()](class-aws-exception-awsexception-method-getawserrortype.md)
  - [getAwsRequestId()](class-aws-exception-awsexception-method-getawsrequestid.md)
  - [getCommand()](class-aws-exception-awsexception-method-getcommand.md)
  - [getIterator()](class-aws-hasdatatrait.md#method_getIterator)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getRequest()](class-aws-exception-awsexception-method-getrequest.md)
  - [getResponse()](class-aws-exception-awsexception-method-getresponse.md)
  - [getResult()](class-aws-exception-awsexception-method-getresult.md)
  - [getStatusCode()](class-aws-exception-awsexception-method-getstatuscode.md)
  - [getTransferInfo()](class-aws-exception-awsexception-method-gettransferinfo.md)
  - [hasKey()](class-aws-exception-awsexception-method-haskey.md)
  - [isConnectionError()](class-aws-exception-awsexception-method-isconnectionerror.md)
  - [isMaxRetriesExceeded()](class-aws-exception-awsexception-method-ismaxretriesexceeded.md)
  - [offsetExists()](class-aws-hasdatatrait.md#method_offsetExists)
  - [offsetGet()](class-aws-hasdatatrait.md#method_offsetGet)
  - [offsetSet()](class-aws-hasdatatrait.md#method_offsetSet)
  - [offsetUnset()](class-aws-hasdatatrait.md#method_offsetUnset)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
  - [search()](class-aws-exception-awsexception-method-search.md)
  - [setMaxRetriesExceeded()](class-aws-exception-awsexception-method-setmaxretriesexceeded.md)
  - [setTransferInfo()](class-aws-exception-awsexception-method-settransferinfo.md)
  - [toArray()](class-aws-hasdatatrait.md#method_toArray)

[Back To Top](class-aws-applicationinsights-exception-applicationinsightsexception-top.md)
