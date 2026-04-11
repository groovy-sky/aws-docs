Menu

- [Aws](namespace-aws.md)
- [Sns](namespace-aws-sns.md)

## Message        in package    - [Aws](package-aws.md)       implements  ArrayAccess, IteratorAggregate

Represents an SNS message received over http(s).

### Table of Contents  [header link](class-aws-sns-message-toc.md)

#### Interfaces  [header link](class-aws-sns-message-toc-interfaces.md)

ArrayAccessIteratorAggregate

#### Methods  [header link](class-aws-sns-message-toc-methods.md)

[\_\_construct()](class-aws-sns-message-method-construct.md)
: mixed Creates a Message object from an array of raw message data.[fromJsonString()](class-aws-sns-message-method-fromjsonstring.md)
: [Message](class-aws-sns-message.md)Creates a Message object from a JSON-decodable string.[fromPsrRequest()](class-aws-sns-message-method-frompsrrequest.md)
: [Message](class-aws-sns-message.md)Creates a Message object from a PSR-7 Request or ServerRequest object.[fromRawPostData()](class-aws-sns-message-method-fromrawpostdata.md)
: [Message](class-aws-sns-message.md)Creates a Message object from the raw POST data[getIterator()](class-aws-sns-message-method-getiterator.md)
: mixed [offsetExists()](class-aws-sns-message-method-offsetexists.md)
: mixed [offsetGet()](class-aws-sns-message-method-offsetget.md)
: mixed [offsetSet()](class-aws-sns-message-method-offsetset.md)
: mixed [offsetUnset()](class-aws-sns-message-method-offsetunset.md)
: mixed [toArray()](class-aws-sns-message-method-toarray.md)
: array<string\|int, mixed> Get all the message data as a plain array.

### Methods  [header link](class-aws-sns-message-methods.md)

#### \_\_construct()  [header link](class-aws-sns-message-method-construct.md)

Creates a Message object from an array of raw message data.

`
    public
                    __construct(array<string|int, mixed> $data) : mixed`

##### Parameters

$data
: array<string\|int, mixed>

The message data.

##### Tags  [header link](class-aws-sns-message-method-construct-tags.md)

throwsInvalidArgumentException

If a valid type is not provided or
there are other required keys missing.

#### fromJsonString()  [header link](class-aws-sns-message-method-fromjsonstring.md)

Creates a Message object from a JSON-decodable string.

`
    public
            static        fromJsonString(string $requestBody) : Message`

##### Parameters

$requestBody
: string

##### Return values

[Message](class-aws-sns-message.md)

#### fromPsrRequest()  [header link](class-aws-sns-message-method-frompsrrequest.md)

Creates a Message object from a PSR-7 Request or ServerRequest object.

`
    public
            static        fromPsrRequest(RequestInterface $request) : Message`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

##### Return values

[Message](class-aws-sns-message.md)

#### fromRawPostData()  [header link](class-aws-sns-message-method-fromrawpostdata.md)

Creates a Message object from the raw POST data

`
    public
            static        fromRawPostData() : Message`

##### Tags  [header link](class-aws-sns-message-method-fromrawpostdata-tags.md)

throwsRuntimeException

If the POST data is absent, or not a valid JSON document

##### Return values

[Message](class-aws-sns-message.md)

#### getIterator()  [header link](class-aws-sns-message-method-getiterator.md)

`
    public
                    getIterator() : mixed`

#### offsetExists()  [header link](class-aws-sns-message-method-offsetexists.md)

`
    public
                    offsetExists(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### offsetGet()  [header link](class-aws-sns-message-method-offsetget.md)

`
    public
                    offsetGet(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### offsetSet()  [header link](class-aws-sns-message-method-offsetset.md)

`
    public
                    offsetSet(mixed $key, mixed $value) : mixed`

##### Parameters

$key
: mixed$value
: mixed

#### offsetUnset()  [header link](class-aws-sns-message-method-offsetunset.md)

`
    public
                    offsetUnset(mixed $key) : mixed`

##### Parameters

$key
: mixed

#### toArray()  [header link](class-aws-sns-message-method-toarray.md)

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
  - [Methods](class-aws-sns-message-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-sns-message-method-construct.md)
  - [fromJsonString()](class-aws-sns-message-method-fromjsonstring.md)
  - [fromPsrRequest()](class-aws-sns-message-method-frompsrrequest.md)
  - [fromRawPostData()](class-aws-sns-message-method-fromrawpostdata.md)
  - [getIterator()](class-aws-sns-message-method-getiterator.md)
  - [offsetExists()](class-aws-sns-message-method-offsetexists.md)
  - [offsetGet()](class-aws-sns-message-method-offsetget.md)
  - [offsetSet()](class-aws-sns-message-method-offsetset.md)
  - [offsetUnset()](class-aws-sns-message-method-offsetunset.md)
  - [toArray()](class-aws-sns-message-method-toarray.md)

[Back To Top](class-aws-sns-message-top.md)
