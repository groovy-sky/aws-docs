Menu

- [Aws](namespace-aws.md)
- [Sns](namespace-aws-sns.md)

## Message        in package    - [Aws](package-aws.md)       implements  ArrayAccess, IteratorAggregate

Represents an SNS message received over http(s).

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#toc-interfaces)

ArrayAccessIteratorAggregate

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method___construct)
: mixed Creates a Message object from an array of raw message data.[fromJsonString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_fromJsonString)
: [Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html)Creates a Message object from a JSON-decodable string.[fromPsrRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_fromPsrRequest)
: [Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html)Creates a Message object from a PSR-7 Request or ServerRequest object.[fromRawPostData()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_fromRawPostData)
: [Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html)Creates a Message object from the raw POST data[getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_getIterator)
: mixed [offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_offsetExists)
: mixed [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_offsetGet)
: mixed [offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_offsetSet)
: mixed [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_offsetUnset)
: mixed [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_toArray)
: array<string\|int, mixed> Get all the message data as a plain array.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method___construct)

Creates a Message object from an array of raw message data.

`
    public
                    __construct(array<string|int, mixed> $data) : mixed`

##### Parameters

$data
: array<string\|int, mixed>

The message data.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method___construct\#tags)

throwsInvalidArgumentException

If a valid type is not provided or
there are other required keys missing.

#### fromJsonString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_fromJsonString)

Creates a Message object from a JSON-decodable string.

`
    public
            static        fromJsonString(string $requestBody) : Message`

##### Parameters

$requestBody
: string

##### Return values

[Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html)

#### fromPsrRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_fromPsrRequest)

Creates a Message object from a PSR-7 Request or ServerRequest object.

`
    public
            static        fromPsrRequest(RequestInterface $request) : Message`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

##### Return values

[Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html)

#### fromRawPostData()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_fromRawPostData)

Creates a Message object from the raw POST data

`
    public
            static        fromRawPostData() : Message`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_fromRawPostData\#tags)

throwsRuntimeException

If the POST data is absent, or not a valid JSON document

##### Return values

[Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html)

#### getIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_getIterator)

`
    public
                    getIterator() : mixed`

#### offsetExists()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_offsetExists)

`
    public
                    offsetExists(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### offsetGet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_offsetGet)

`
    public
                    offsetGet(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### offsetSet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_offsetSet)

`
    public
                    offsetSet(mixed $key, mixed $value) : mixed`

##### Parameters

$key
: mixed$value
: mixed

#### offsetUnset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html\#method_toArray)

Get all the message data as a plain array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method___construct)
  - [fromJsonString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_fromJsonString)
  - [fromPsrRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_fromPsrRequest)
  - [fromRawPostData()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_fromRawPostData)
  - [getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_getIterator)
  - [offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_offsetExists)
  - [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_offsetGet)
  - [offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_offsetSet)
  - [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_offsetUnset)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.Message.html#top)
