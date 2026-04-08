Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)
- [Parser](namespace-aws-api-parser.md)
- [Exception](namespace-aws-api-parser-exception.md)

## ParserException     extends RuntimeException   in package    - [Aws](package-aws.md)       implements  [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md), [ResponseContainerInterface](class-aws-responsecontainerinterface.md)  Uses  [HasMonitoringEventsTrait](class-aws-hasmonitoringeventstrait.md)

### Table of Contents  [header link](class-aws-api-parser-exception-parserexception-toc.md)

#### Interfaces  [header link](class-aws-api-parser-exception-parserexception-toc-interfaces.md)

[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events[ResponseContainerInterface](class-aws-responsecontainerinterface.md)

#### Methods  [header link](class-aws-api-parser-exception-parserexception-toc-methods.md)

[\_\_construct()](class-aws-api-parser-exception-parserexception-method-construct.md)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[getErrorCode()](class-aws-api-parser-exception-parserexception-method-geterrorcode.md)
: string\|null Get the error code, if any.[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getRequestId()](class-aws-api-parser-exception-parserexception-method-getrequestid.md)
: string\|null Get the request ID, if any.[getResponse()](class-aws-api-parser-exception-parserexception-method-getresponse.md)
: [ResponseInterface](class-psr-http-message-responseinterface.md) \|null Get the received HTTP response if any.[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](class-aws-api-parser-exception-parserexception-methods.md)

#### \_\_construct()  [header link](class-aws-api-parser-exception-parserexception-method-construct.md)

`
    public
                    __construct([mixed $message = '' ][, mixed $code = 0 ][, mixed $previous = null ][, array<string|int, mixed> $context = [] ]) : mixed`

##### Parameters

$message
: mixed
= ''$code
: mixed
= 0$previous
: mixed
= null$context
: array<string\|int, mixed>
= \[\]

#### appendMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_appendMonitoringEvent)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### getErrorCode()  [header link](class-aws-api-parser-exception-parserexception-method-geterrorcode.md)

Get the error code, if any.

`
    public
                    getErrorCode() : string|null`

##### Return values

string\|null

#### getMonitoringEvents()  [header link](class-aws-hasmonitoringeventstrait.md\#method_getMonitoringEvents)

Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.

`
    public
                    getMonitoringEvents() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getRequestId()  [header link](class-aws-api-parser-exception-parserexception-method-getrequestid.md)

Get the request ID, if any.

`
    public
                    getRequestId() : string|null`

##### Return values

string\|null

#### getResponse()  [header link](class-aws-api-parser-exception-parserexception-method-getresponse.md)

Get the received HTTP response if any.

`
    public
                    getResponse() : ResponseInterface|null`

##### Return values

[ResponseInterface](class-psr-http-message-responseinterface.md) \|null

#### prependMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_prependMonitoringEvent)

Prepend a client-side monitoring event to this object's event list

`
    public
                    prependMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-api-parser-exception-parserexception-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-parser-exception-parserexception-method-construct.md)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [getErrorCode()](class-aws-api-parser-exception-parserexception-method-geterrorcode.md)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getRequestId()](class-aws-api-parser-exception-parserexception-method-getrequestid.md)
  - [getResponse()](class-aws-api-parser-exception-parserexception-method-getresponse.md)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](class-aws-api-parser-exception-parserexception-top.md)
