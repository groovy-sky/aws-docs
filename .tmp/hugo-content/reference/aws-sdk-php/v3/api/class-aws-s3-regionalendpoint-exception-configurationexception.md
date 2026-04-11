Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [RegionalEndpoint](namespace-aws-s3-regionalendpoint.md)
- [Exception](namespace-aws-s3-regionalendpoint-exception.md)

## ConfigurationException     extends RuntimeException   in package    - [Aws](package-aws.md)       implements  [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)  Uses  [HasMonitoringEventsTrait](class-aws-hasmonitoringeventstrait.md)

Represents an error interacting with configuration for sts regional endpoints

### Table of Contents  [header link](class-aws-s3-regionalendpoint-exception-configurationexception-toc.md)

#### Interfaces  [header link](class-aws-s3-regionalendpoint-exception-configurationexception-toc-interfaces.md)

[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events

#### Methods  [header link](class-aws-s3-regionalendpoint-exception-configurationexception-toc-methods.md)

[appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](class-aws-s3-regionalendpoint-exception-configurationexception-methods.md)

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
  - [Methods](class-aws-s3-regionalendpoint-exception-configurationexception-toc-methods.md)
- Methods
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](class-aws-s3-regionalendpoint-exception-configurationexception-top.md)
