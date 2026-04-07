Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)
- [Parser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.api.parser.html)
- [Exception](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.api.parser.exception.html)

## ParserException     extends RuntimeException   in package    - [Aws](package-aws.md)       implements  [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md), [ResponseContainerInterface](class-aws-responsecontainerinterface.md)  Uses  [HasMonitoringEventsTrait](class-aws-hasmonitoringeventstrait.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html\#toc-interfaces)

[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events[ResponseContainerInterface](class-aws-responsecontainerinterface.md)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#method___construct)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[getErrorCode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#method_getErrorCode)
: string\|null Get the error code, if any.[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getRequestId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#method_getRequestId)
: string\|null Get the request ID, if any.[getResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#method_getResponse)
: [ResponseInterface](class-psr-http-message-responseinterface.md) \|null Get the received HTTP response if any.[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html\#method___construct)

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

#### getErrorCode()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html\#method_getErrorCode)

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

#### getRequestId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html\#method_getRequestId)

Get the request ID, if any.

`
    public
                    getRequestId() : string|null`

##### Return values

string\|null

#### getResponse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html\#method_getResponse)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#method___construct)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [getErrorCode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#method_getErrorCode)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getRequestId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#method_getRequestId)
  - [getResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#method_getResponse)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Parser.Exception.ParserException.html#top)
