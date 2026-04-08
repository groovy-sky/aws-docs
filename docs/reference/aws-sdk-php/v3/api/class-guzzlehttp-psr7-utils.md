Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## Utils        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-guzzlehttp-psr7-utils-toc.md)

#### Methods  [header link](class-guzzlehttp-psr7-utils-toc-methods.md)

[caselessRemove()](class-guzzlehttp-psr7-utils-method-caselessremove.md)
: array<string\|int, mixed> Remove the items given by the keys, case insensitively from the data.[copyToStream()](class-guzzlehttp-psr7-utils-method-copytostream.md)
: void Copy the contents of a stream into another stream until the given number
of bytes have been read.[copyToString()](class-guzzlehttp-psr7-utils-method-copytostring.md)
: string Copy the contents of a stream into a string until the given number of
bytes have been read.[hash()](class-guzzlehttp-psr7-utils-method-hash.md)
: string Calculate a hash of a stream.[modifyRequest()](class-guzzlehttp-psr7-utils-method-modifyrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Clone and modify a request with the given changes.[readLine()](class-guzzlehttp-psr7-utils-method-readline.md)
: string Read a line from the stream up to the maximum allowed buffer length.[redactUserInfo()](class-guzzlehttp-psr7-utils-method-redactuserinfo.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Redact the password in the user info part of a URI.[streamFor()](class-guzzlehttp-psr7-utils-method-streamfor.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)Create a new stream based on the input type.[tryFopen()](class-guzzlehttp-psr7-utils-method-tryfopen.md)
: resource Safely opens a PHP stream resource using a filename.[tryGetContents()](class-guzzlehttp-psr7-utils-method-trygetcontents.md)
: string Safely gets the contents of a given stream.[uriFor()](class-guzzlehttp-psr7-utils-method-urifor.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Returns a UriInterface for the given value.

### Methods  [header link](class-guzzlehttp-psr7-utils-methods.md)

#### caselessRemove()  [header link](class-guzzlehttp-psr7-utils-method-caselessremove.md)

Remove the items given by the keys, case insensitively from the data.

`
    public
            static        caselessRemove(array<string|int, string|int> $keys, array<string|int, mixed> $data) : array<string|int, mixed>`

##### Parameters

$keys
: array<string\|int, string\|int>$data
: array<string\|int, mixed>

##### Return values

array<string\|int, mixed>

#### copyToStream()  [header link](class-guzzlehttp-psr7-utils-method-copytostream.md)

Copy the contents of a stream into another stream until the given number
of bytes have been read.

`
    public
            static        copyToStream(StreamInterface $source, StreamInterface $dest[, int $maxLen = -1 ]) : void`

##### Parameters

$source
: [StreamInterface](class-psr-http-message-streaminterface.md)

Stream to read from

$dest
: [StreamInterface](class-psr-http-message-streaminterface.md)

Stream to write to

$maxLen
: int
= -1

Maximum number of bytes to read. Pass -1
to read the entire stream.

##### Tags  [header link](class-guzzlehttp-psr7-utils-method-copytostream-tags.md)

throwsRuntimeException

on error.

#### copyToString()  [header link](class-guzzlehttp-psr7-utils-method-copytostring.md)

Copy the contents of a stream into a string until the given number of
bytes have been read.

`
    public
            static        copyToString(StreamInterface $stream[, int $maxLen = -1 ]) : string`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

Stream to read

$maxLen
: int
= -1

Maximum number of bytes to read. Pass -1
to read the entire stream.

##### Tags  [header link](class-guzzlehttp-psr7-utils-method-copytostring-tags.md)

throwsRuntimeException

on error.

##### Return values

string

#### hash()  [header link](class-guzzlehttp-psr7-utils-method-hash.md)

Calculate a hash of a stream.

`
    public
            static        hash(StreamInterface $stream, string $algo[, bool $rawOutput = false ]) : string`

This method reads the entire stream to calculate a rolling hash, based
on PHP's `hash_init` functions.

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

Stream to calculate the hash for

$algo
: string

Hash algorithm (e.g. md5, crc32, etc)

$rawOutput
: bool
= false

Whether or not to use raw output

##### Tags  [header link](class-guzzlehttp-psr7-utils-method-hash-tags.md)

throwsRuntimeException

on error.

##### Return values

string

#### modifyRequest()  [header link](class-guzzlehttp-psr7-utils-method-modifyrequest.md)

Clone and modify a request with the given changes.

`
    public
            static        modifyRequest(RequestInterface $request, array<string|int, mixed> $changes) : RequestInterface`

This method is useful for reducing the number of clones needed to mutate
a message.

The changes can be one of:

- method: (string) Changes the HTTP method.
- set\_headers: (array) Sets the given headers.
- remove\_headers: (array) Remove the given headers.
- body: (mixed) Sets the given body.
- uri: (UriInterface) Set the URI.
- query: (string) Set the query string value of the URI.
- version: (string) Set the protocol version.

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to clone and modify.

$changes
: array<string\|int, mixed>

Changes to apply.

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### readLine()  [header link](class-guzzlehttp-psr7-utils-method-readline.md)

Read a line from the stream up to the maximum allowed buffer length.

`
    public
            static        readLine(StreamInterface $stream[, int|null $maxLength = null ]) : string`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

Stream to read from

$maxLength
: int\|null
= null

Maximum buffer length

##### Return values

string

#### redactUserInfo()  [header link](class-guzzlehttp-psr7-utils-method-redactuserinfo.md)

Redact the password in the user info part of a URI.

`
    public
            static        redactUserInfo(UriInterface $uri) : UriInterface`

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### streamFor()  [header link](class-guzzlehttp-psr7-utils-method-streamfor.md)

Create a new stream based on the input type.

`
    public
            static        streamFor([resource|string|int|float|bool|StreamInterface|callable|Iterator|null $resource = '' ][, array{size?: int, metadata?: array} $options = [] ]) : StreamInterface`

Options is an associative array that can contain the following keys:

- metadata: Array of custom metadata.
- size: Size of the stream.

This method accepts the following `$resource` types:

- `Psr\Http\Message\StreamInterface`: Returns the value as-is.
- `string`: Creates a stream object that uses the given string as the contents.
- `resource`: Creates a stream object that wraps the given PHP stream resource.
- `Iterator`: If the provided value implements `Iterator`, then a read-only
stream object will be created that wraps the given iterable. Each time the
stream is read from, data from the iterator will fill a buffer and will be
continuously called until the buffer is equal to the requested read size.
Subsequent read calls will first read from the buffer and then call `next`
on the underlying iterator until it is exhausted.
- `object` with `__toString()`: If the object has the `__toString()` method,
the object will be cast to a string and then a stream will be returned that
uses the string value.
- `NULL`: When `null` is passed, an empty stream object is returned.
- `callable` When a callable is passed, a read-only stream object will be
created that invokes the given callable. The callable is invoked with the
number of suggested bytes to read. The callable can return any number of
bytes, but MUST return `false` when there is no more data to return. The
stream object that wraps the callable will invoke the callable until the
number of requested bytes are available. Any additional bytes will be
buffered and used in subsequent reads.

##### Parameters

$resource
: resource\|string\|int\|float\|bool\| [StreamInterface](class-psr-http-message-streaminterface.md) \|callable\|Iterator\|null
= ''

Entity body data

$options
: array{size?: int, metadata?: array}
= \[\]

Additional options

##### Tags  [header link](class-guzzlehttp-psr7-utils-method-streamfor-tags.md)

throwsInvalidArgumentException

if the $resource arg is not valid.

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### tryFopen()  [header link](class-guzzlehttp-psr7-utils-method-tryfopen.md)

Safely opens a PHP stream resource using a filename.

`
    public
            static        tryFopen(string $filename, string $mode) : resource`

When fopen fails, PHP normally raises a warning. This function adds an
error handler that checks for errors and throws an exception instead.

##### Parameters

$filename
: string

File to open

$mode
: string

Mode used to open the file

##### Tags  [header link](class-guzzlehttp-psr7-utils-method-tryfopen-tags.md)

throwsRuntimeException

if the file cannot be opened

##### Return values

resource

#### tryGetContents()  [header link](class-guzzlehttp-psr7-utils-method-trygetcontents.md)

Safely gets the contents of a given stream.

`
    public
            static        tryGetContents(resource $stream) : string`

When stream\_get\_contents fails, PHP normally raises a warning. This
function adds an error handler that checks for errors and throws an
exception instead.

##### Parameters

$stream
: resource

##### Tags  [header link](class-guzzlehttp-psr7-utils-method-trygetcontents-tags.md)

throwsRuntimeException

if the stream cannot be read

##### Return values

string

#### uriFor()  [header link](class-guzzlehttp-psr7-utils-method-urifor.md)

Returns a UriInterface for the given value.

`
    public
            static        uriFor(string|UriInterface $uri) : UriInterface`

This function accepts a string or UriInterface and returns a
UriInterface for the given value. If the value is already a
UriInterface, it is returned as-is.

##### Parameters

$uri
: string\| [UriInterface](class-psr-http-message-uriinterface.md)

##### Tags  [header link](class-guzzlehttp-psr7-utils-method-urifor-tags.md)

throwsInvalidArgumentException

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-psr7-utils-toc-methods.md)
- Methods
  - [caselessRemove()](class-guzzlehttp-psr7-utils-method-caselessremove.md)
  - [copyToStream()](class-guzzlehttp-psr7-utils-method-copytostream.md)
  - [copyToString()](class-guzzlehttp-psr7-utils-method-copytostring.md)
  - [hash()](class-guzzlehttp-psr7-utils-method-hash.md)
  - [modifyRequest()](class-guzzlehttp-psr7-utils-method-modifyrequest.md)
  - [readLine()](class-guzzlehttp-psr7-utils-method-readline.md)
  - [redactUserInfo()](class-guzzlehttp-psr7-utils-method-redactuserinfo.md)
  - [streamFor()](class-guzzlehttp-psr7-utils-method-streamfor.md)
  - [tryFopen()](class-guzzlehttp-psr7-utils-method-tryfopen.md)
  - [tryGetContents()](class-guzzlehttp-psr7-utils-method-trygetcontents.md)
  - [uriFor()](class-guzzlehttp-psr7-utils-method-urifor.md)

[Back To Top](class-guzzlehttp-psr7-utils-top.md)
