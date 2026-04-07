Menu

- [Aws](namespace-aws.md)

## Result        in package    - [Aws](package-aws.md)       implements  [ResultInterface](class-aws-resultinterface.md), [MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)  Uses  [HasDataTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html), [HasMonitoringEventsTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html)

AWS result.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#toc-interfaces)

[ResultInterface](class-aws-resultinterface.md)Represents an AWS result object that is returned from executing an operation.[MonitoringEventsInterface](class-aws-monitoringeventsinterface.md)Interface for adding and retrieving client-side monitoring events

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method___construct)
: mixed [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method___toString)
: string Provides debug information about the result object[appendMonitoringEvent()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_count)
: int [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method_get)
: mixed\|null Get a specific key value from the result model.[getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_getIterator)
: Traversable[getMonitoringEvents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method_getPath)
: mixed [hasKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method_hasKey)
: bool Check if the model contains a key by name[offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetExists)
: bool [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetGet)
: mixed\|null This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').[offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetSet)
: void [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetUnset)
: void [prependMonitoringEvent()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list[search()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method_search)
: mixed Returns the result of executing a JMESPath expression on the contents
of the Result model.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#method___construct)

`
    public
                    __construct([array<string|int, mixed> $data = [] ]) : mixed`

##### Parameters

$data
: array<string\|int, mixed>
= \[\]

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#method___toString)

Provides debug information about the result object

`
    public
                    __toString() : string`

##### Return values

string

#### appendMonitoringEvent()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html\#method_appendMonitoringEvent)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### count()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_count)

`
    public
                    count() : int`

##### Return values

int

#### get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#method_get)

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

#### getIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_getIterator)

`
    public
                    getIterator() : Traversable`

##### Return values

Traversable

#### getMonitoringEvents()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html\#method_getMonitoringEvents)

Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.

`
    public
                    getMonitoringEvents() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getPath()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#method_getPath)

`
    public
                    getPath(mixed $path) : mixed`

##### Parameters

$path
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#method_getPath\#tags)

deprecated

#### hasKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#method_hasKey)

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

#### offsetExists()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_offsetExists)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_offsetGet)

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

#### offsetSet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_offsetSet)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### prependMonitoringEvent()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html\#method_prependMonitoringEvent)

Prepend a client-side monitoring event to this object's event list

`
    public
                    prependMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### search()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html\#method_search)

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

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_toArray)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method___construct)
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method___toString)
  - [appendMonitoringEvent()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html#method_appendMonitoringEvent)
  - [count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_count)
  - [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method_get)
  - [getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_getIterator)
  - [getMonitoringEvents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html#method_getMonitoringEvents)
  - [getPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method_getPath)
  - [hasKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method_hasKey)
  - [offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetExists)
  - [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetGet)
  - [offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetSet)
  - [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetUnset)
  - [prependMonitoringEvent()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasMonitoringEventsTrait.html#method_prependMonitoringEvent)
  - [search()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#method_search)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Result.html#top)
