Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Exception](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.exception.html)

## DeleteMultipleObjectsException     extends Exception   in package    - [Aws](package-aws.md)       implements  [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)  Uses  [HasMonitoringEventsTrait](class-aws-hasmonitoringeventstrait.md)

Exception thrown when errors occur while deleting objects using a
{@see S3\\BatchDelete} object.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html\#toc-interfaces)

[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#method___construct)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[createMessageFromErrors()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#method_createMessageFromErrors)
: string Create a single error message from multiple errors.[getDeleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#method_getDeleted)
: array<string\|int, mixed> Get the successfully deleted objects[getErrors()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#method_getErrors)
: array<string\|int, mixed> Get the errored objects[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html\#method___construct)

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

#### createMessageFromErrors()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html\#method_createMessageFromErrors)

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

#### getDeleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html\#method_getDeleted)

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

#### getErrors()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html\#method_getErrors)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#method___construct)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [createMessageFromErrors()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#method_createMessageFromErrors)
  - [getDeleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#method_getDeleted)
  - [getErrors()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#method_getErrors)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html#top)
