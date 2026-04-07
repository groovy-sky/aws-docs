Menu

- [Aws](namespace-aws.md)

## MonitoringEventsInterface     in    - [Aws](package-aws.md)

Interface for adding and retrieving client-side monitoring events

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html\#toc-methods)

[appendMonitoringEvent()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[getMonitoringEvents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[prependMonitoringEvent()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html\#methods)

#### appendMonitoringEvent()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html\#method_appendMonitoringEvent)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### getMonitoringEvents()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html\#method_getMonitoringEvents)

Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.

`
    public
                    getMonitoringEvents() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### prependMonitoringEvent()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html\#method_prependMonitoringEvent)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#toc-methods)
- Methods
  - [appendMonitoringEvent()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#method_appendMonitoringEvent)
  - [getMonitoringEvents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#method_getMonitoringEvents)
  - [prependMonitoringEvent()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#method_prependMonitoringEvent)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MonitoringEventsInterface.html#top)
