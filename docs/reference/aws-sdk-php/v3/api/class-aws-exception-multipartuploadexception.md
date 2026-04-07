Menu

- [Aws](namespace-aws.md)
- [Exception](namespace-aws-exception.md)

## MultipartUploadException     extends RuntimeException   in package    - [Aws](package-aws.md)       implements  [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)  Uses  [HasMonitoringEventsTrait](class-aws-hasmonitoringeventstrait.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html\#toc-interfaces)

[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html#method___construct)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getState()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html#method_getState)
: [UploadState](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html)Get the state of the transfer[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html\#method___construct)

`
    public
                    __construct(UploadState $state[, Exception|array<string|int, mixed> $prev = null ]) : mixed`

##### Parameters

$state
: [UploadState](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html)

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

#### getState()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html\#method_getState)

Get the state of the transfer

`
    public
                    getState() : UploadState`

##### Return values

[UploadState](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Multipart.UploadState.html)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html#method___construct)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getState()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html#method_getState)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.MultipartUploadException.html#top)
