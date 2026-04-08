Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Exception](namespace-aws-s3-exception.md)

## DeleteMultipleObjectsException     extends Exception   in package    - [Aws](package-aws.md)       implements  [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)  Uses  [HasMonitoringEventsTrait](class-aws-hasmonitoringeventstrait.md)

Exception thrown when errors occur while deleting objects using a
{@see S3\\BatchDelete} object.

### Table of Contents  [header link](class-aws-s3-exception-deletemultipleobjectsexception-toc.md)

#### Interfaces  [header link](class-aws-s3-exception-deletemultipleobjectsexception-toc-interfaces.md)

[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events

#### Methods  [header link](class-aws-s3-exception-deletemultipleobjectsexception-toc-methods.md)

[\_\_construct()](class-aws-s3-exception-deletemultipleobjectsexception-method-construct.md)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[createMessageFromErrors()](class-aws-s3-exception-deletemultipleobjectsexception-method-createmessagefromerrors.md)
: string Create a single error message from multiple errors.[getDeleted()](class-aws-s3-exception-deletemultipleobjectsexception-method-getdeleted.md)
: array<string\|int, mixed> Get the successfully deleted objects[getErrors()](class-aws-s3-exception-deletemultipleobjectsexception-method-geterrors.md)
: array<string\|int, mixed> Get the errored objects[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](class-aws-s3-exception-deletemultipleobjectsexception-methods.md)

#### \_\_construct()  [header link](class-aws-s3-exception-deletemultipleobjectsexception-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $deleted, array<string|int, mixed> $errors) : mixed`

##### Parameters

$deleted
: array<string\|int, mixed>

Array of successfully deleted keys

$errors
: array<string\|int, mixed>

Array of errors that were encountered

#### appendMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_appendMonitoringEvent)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### createMessageFromErrors()  [header link](class-aws-s3-exception-deletemultipleobjectsexception-method-createmessagefromerrors.md)

Create a single error message from multiple errors.

`
    public
            static        createMessageFromErrors(array<string|int, mixed> $errors) : string`

##### Parameters

$errors
: array<string\|int, mixed>

Errors encountered

##### Return values

string

#### getDeleted()  [header link](class-aws-s3-exception-deletemultipleobjectsexception-method-getdeleted.md)

Get the successfully deleted objects

`
    public
                    getDeleted() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
—

Returns an array of associative arrays, each containing
a 'Key' and optionally 'DeleteMarker' and
'DeleterMarkerVersionId'

#### getErrors()  [header link](class-aws-s3-exception-deletemultipleobjectsexception-method-geterrors.md)

Get the errored objects

`
    public
                    getErrors() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
—

Returns an array of associative arrays, each containing
a 'Code', 'Message', and 'Key' key.

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
  - [Methods](class-aws-s3-exception-deletemultipleobjectsexception-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-exception-deletemultipleobjectsexception-method-construct.md)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [createMessageFromErrors()](class-aws-s3-exception-deletemultipleobjectsexception-method-createmessagefromerrors.md)
  - [getDeleted()](class-aws-s3-exception-deletemultipleobjectsexception-method-getdeleted.md)
  - [getErrors()](class-aws-s3-exception-deletemultipleobjectsexception-method-geterrors.md)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](class-aws-s3-exception-deletemultipleobjectsexception-top.md)
