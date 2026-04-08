Menu

- [Aws](namespace-aws.md)
- [Exception](namespace-aws-exception.md)

## MultipartUploadException     extends RuntimeException   in package    - [Aws](package-aws.md)       implements  [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)  Uses  [HasMonitoringEventsTrait](class-aws-hasmonitoringeventstrait.md)

### Table of Contents  [header link](class-aws-exception-multipartuploadexception-toc.md)

#### Interfaces  [header link](class-aws-exception-multipartuploadexception-toc-interfaces.md)

[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events

#### Methods  [header link](class-aws-exception-multipartuploadexception-toc-methods.md)

[\_\_construct()](class-aws-exception-multipartuploadexception-method-construct.md)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getState()](class-aws-exception-multipartuploadexception-method-getstate.md)
: [UploadState](class-aws-multipart-uploadstate.md)Get the state of the transfer[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](class-aws-exception-multipartuploadexception-methods.md)

#### \_\_construct()  [header link](class-aws-exception-multipartuploadexception-method-construct.md)

`
    public
                    __construct(UploadState $state[, Exception|array<string|int, mixed> $prev = null ]) : mixed`

##### Parameters

$state
: [UploadState](class-aws-multipart-uploadstate.md)

Upload state at time of the exception.

$prev
: Exception\|array<string\|int, mixed>
= null

Exception being thrown.

#### appendMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_appendMonitoringEvent)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### getMonitoringEvents()  [header link](class-aws-hasmonitoringeventstrait.md\#method_getMonitoringEvents)

Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.

`
    public
                    getMonitoringEvents() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getState()  [header link](class-aws-exception-multipartuploadexception-method-getstate.md)

Get the state of the transfer

`
    public
                    getState() : UploadState`

##### Return values

[UploadState](class-aws-multipart-uploadstate.md)

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
  - [Methods](class-aws-exception-multipartuploadexception-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-exception-multipartuploadexception-method-construct.md)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getState()](class-aws-exception-multipartuploadexception-method-getstate.md)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](class-aws-exception-multipartuploadexception-top.md)
