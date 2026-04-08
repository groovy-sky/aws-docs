Menu

- [Aws](namespace-aws.md)

## Result        in package    - [Aws](package-aws.md)       implements  [ResultInterface](class-aws-resultinterface.md), [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)  Uses  [HasDataTrait](class-aws-hasdatatrait.md), [HasMonitoringEventsTrait](class-aws-hasmonitoringeventstrait.md)

AWS result.

### Table of Contents  [header link](class-aws-result-toc.md)

#### Interfaces  [header link](class-aws-result-toc-interfaces.md)

[ResultInterface](class-aws-resultinterface.md)Represents an AWS result object that is returned from executing an operation.[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events

#### Methods  [header link](class-aws-result-toc-methods.md)

[\_\_construct()](class-aws-result-method-construct.md)
: mixed [\_\_toString()](class-aws-result-method-tostring.md)
: string Provides debug information about the result object[appendMonitoringEvent()](class-aws-hasmonitoringeventstrait-method-appendmonitoringevent.md)
: mixed Append a client-side monitoring event to this object's event list[count()](class-aws-hasdatatrait-method-count.md)
: int [get()](class-aws-result-method-get.md)
: mixed\|null Get a specific key value from the result model.[getIterator()](class-aws-hasdatatrait-method-getiterator.md)
: Traversable[getMonitoringEvents()](class-aws-hasmonitoringeventstrait-method-getmonitoringevents.md)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getPath()](class-aws-result-method-getpath.md)
: mixed [hasKey()](class-aws-result-method-haskey.md)
: bool Check if the model contains a key by name[offsetExists()](class-aws-hasdatatrait-method-offsetexists.md)
: bool [offsetGet()](class-aws-hasdatatrait-method-offsetget.md)
: mixed\|null This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').[offsetSet()](class-aws-hasdatatrait-method-offsetset.md)
: void [offsetUnset()](class-aws-hasdatatrait-method-offsetunset.md)
: void [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait-method-prependmonitoringevent.md)
: mixed Prepend a client-side monitoring event to this object's event list[search()](class-aws-result-method-search.md)
: mixed Returns the result of executing a JMESPath expression on the contents
of the Result model.[toArray()](class-aws-hasdatatrait-method-toarray.md)
: mixed

### Methods  [header link](class-aws-result-methods.md)

#### \_\_construct()  [header link](class-aws-result-method-construct.md)

`
    public
                    __construct([array<string|int, mixed> $data = [] ]) : mixed`

##### Parameters

$data
: array<string\|int, mixed>
= \[\]

#### \_\_toString()  [header link](class-aws-result-method-tostring.md)

Provides debug information about the result object

`
    public
                    __toString() : string`

##### Return values

string

#### appendMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait-method-appendmonitoringevent.md)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### count()  [header link](class-aws-hasdatatrait-method-count.md)

`
    public
                    count() : int`

##### Return values

int

#### get()  [header link](class-aws-result-method-get.md)

Get a specific key value from the result model.

`
    public
                    get(mixed $key) : mixed|null`

##### Parameters

$key
: mixed

Key to retrieve.

##### Return values

mixed\|null
—

Value of the key or NULL if not found.

#### getIterator()  [header link](class-aws-hasdatatrait-method-getiterator.md)

`
    public
                    getIterator() : Traversable`

##### Return values

Traversable

#### getMonitoringEvents()  [header link](class-aws-hasmonitoringeventstrait-method-getmonitoringevents.md)

Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.

`
    public
                    getMonitoringEvents() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getPath()  [header link](class-aws-result-method-getpath.md)

`
    public
                    getPath(mixed $path) : mixed`

##### Parameters

$path
: mixed

##### Tags  [header link](class-aws-result-method-getpath-tags.md)

deprecated

#### hasKey()  [header link](class-aws-result-method-haskey.md)

Check if the model contains a key by name

`
    public
                    hasKey(mixed $name) : bool`

##### Parameters

$name
: mixed

Name of the key to retrieve

##### Return values

bool

#### offsetExists()  [header link](class-aws-hasdatatrait-method-offsetexists.md)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](class-aws-hasdatatrait-method-offsetget.md)

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

#### offsetSet()  [header link](class-aws-hasdatatrait-method-offsetset.md)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](class-aws-hasdatatrait-method-offsetunset.md)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### prependMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait-method-prependmonitoringevent.md)

Prepend a client-side monitoring event to this object's event list

`
    public
                    prependMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### search()  [header link](class-aws-result-method-search.md)

Returns the result of executing a JMESPath expression on the contents
of the Result model.

`
    public
                    search(mixed $expression) : mixed`

$result = $client->execute($command);
$jpResult = $result->search('foo.\*.bar\[?baz > `10`\]');

##### Parameters

$expression
: mixed

JMESPath expression to execute

##### Return values

mixed
—

Returns the result of the JMESPath expression.

#### toArray()  [header link](class-aws-hasdatatrait-method-toarray.md)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-result-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-result-method-construct.md)
  - [\_\_toString()](class-aws-result-method-tostring.md)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait-method-appendmonitoringevent.md)
  - [count()](class-aws-hasdatatrait-method-count.md)
  - [get()](class-aws-result-method-get.md)
  - [getIterator()](class-aws-hasdatatrait-method-getiterator.md)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait-method-getmonitoringevents.md)
  - [getPath()](class-aws-result-method-getpath.md)
  - [hasKey()](class-aws-result-method-haskey.md)
  - [offsetExists()](class-aws-hasdatatrait-method-offsetexists.md)
  - [offsetGet()](class-aws-hasdatatrait-method-offsetget.md)
  - [offsetSet()](class-aws-hasdatatrait-method-offsetset.md)
  - [offsetUnset()](class-aws-hasdatatrait-method-offsetunset.md)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait-method-prependmonitoringevent.md)
  - [search()](class-aws-result-method-search.md)
  - [toArray()](class-aws-hasdatatrait-method-toarray.md)

[Back To Top](class-aws-result-top.md)
