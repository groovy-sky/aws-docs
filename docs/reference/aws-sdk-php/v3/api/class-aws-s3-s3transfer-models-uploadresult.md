Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## UploadResult     extends [Result](class-aws-result.md)   in package    - [Aws](package-aws.md)

FinalYes

AWS result.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadResult.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadResult.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadResult.html#method___construct)
: mixed [\_\_toString()](class-aws-result.md#method___toString)
: string Provides debug information about the result object[appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[count()](class-aws-hasdatatrait.md#method_count)
: int [get()](class-aws-result.md#method_get)
: mixed\|null Get a specific key value from the result model.[getIterator()](class-aws-hasdatatrait.md#method_getIterator)
: Traversable[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getPath()](class-aws-result.md#method_getPath)
: mixed [hasKey()](class-aws-result.md#method_hasKey)
: bool Check if the model contains a key by name[offsetExists()](class-aws-hasdatatrait.md#method_offsetExists)
: bool [offsetGet()](class-aws-hasdatatrait.md#method_offsetGet)
: mixed\|null This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').[offsetSet()](class-aws-hasdatatrait.md#method_offsetSet)
: void [offsetUnset()](class-aws-hasdatatrait.md#method_offsetUnset)
: void [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list[search()](class-aws-result.md#method_search)
: mixed Returns the result of executing a JMESPath expression on the contents
of the Result model.[toArray()](class-aws-hasdatatrait.md#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadResult.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadResult.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $data) : mixed`

##### Parameters

$data
: array<string\|int, mixed>

#### \_\_toString()  [header link](class-aws-result.md\#method___toString)

Provides debug information about the result object

`
    public
                    __toString() : string`

##### Return values

string

#### appendMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_appendMonitoringEvent)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### count()  [header link](class-aws-hasdatatrait.md\#method_count)

`
    public
                    count() : int`

##### Return values

int

#### get()  [header link](class-aws-result.md\#method_get)

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

#### getIterator()  [header link](class-aws-hasdatatrait.md\#method_getIterator)

`
    public
                    getIterator() : Traversable`

##### Return values

Traversable

#### getMonitoringEvents()  [header link](class-aws-hasmonitoringeventstrait.md\#method_getMonitoringEvents)

Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.

`
    public
                    getMonitoringEvents() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getPath()  [header link](class-aws-result.md\#method_getPath)

`
    public
                    getPath(mixed $path) : mixed`

##### Parameters

$path
: mixed

##### Tags  [header link](class-aws-result.md\#method_getPath\#tags)

deprecated

#### hasKey()  [header link](class-aws-result.md\#method_hasKey)

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

#### offsetExists()  [header link](class-aws-hasdatatrait.md\#method_offsetExists)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](class-aws-hasdatatrait.md\#method_offsetGet)

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

#### offsetSet()  [header link](class-aws-hasdatatrait.md\#method_offsetSet)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](class-aws-hasdatatrait.md\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### prependMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_prependMonitoringEvent)

Prepend a client-side monitoring event to this object's event list

`
    public
                    prependMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### search()  [header link](class-aws-result.md\#method_search)

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

#### toArray()  [header link](class-aws-hasdatatrait.md\#method_toArray)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadResult.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadResult.html#method___construct)
  - [\_\_toString()](class-aws-result.md#method___toString)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [count()](class-aws-hasdatatrait.md#method_count)
  - [get()](class-aws-result.md#method_get)
  - [getIterator()](class-aws-hasdatatrait.md#method_getIterator)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getPath()](class-aws-result.md#method_getPath)
  - [hasKey()](class-aws-result.md#method_hasKey)
  - [offsetExists()](class-aws-hasdatatrait.md#method_offsetExists)
  - [offsetGet()](class-aws-hasdatatrait.md#method_offsetGet)
  - [offsetSet()](class-aws-hasdatatrait.md#method_offsetSet)
  - [offsetUnset()](class-aws-hasdatatrait.md#method_offsetUnset)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
  - [search()](class-aws-result.md#method_search)
  - [toArray()](class-aws-hasdatatrait.md#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.UploadResult.html#top)
