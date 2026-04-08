Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## MultipartStream        in package    - [Aws](package-aws.md)       implements  [StreamInterface](class-psr-http-message-streaminterface.md)  Uses  [StreamDecoratorTrait](class-guzzlehttp-psr7-streamdecoratortrait.md)

FinalYes

Stream that when read returns bytes for a streaming multipart or
multipart/form-data stream.

### Table of Contents  [header link](class-guzzlehttp-psr7-multipartstream-toc.md)

#### Interfaces  [header link](class-guzzlehttp-psr7-multipartstream-toc-interfaces.md)

[StreamInterface](class-psr-http-message-streaminterface.md)Describes a data stream.

#### Methods  [header link](class-guzzlehttp-psr7-multipartstream-toc-methods.md)

[\_\_call()](class-guzzlehttp-psr7-streamdecoratortrait-method-call.md)
: mixed Allow decorators to implement custom methods[\_\_construct()](class-guzzlehttp-psr7-multipartstream-method-construct.md)
: mixed [\_\_get()](class-guzzlehttp-psr7-streamdecoratortrait-method-get.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)Magic method used to create a new stream if streams are not added in
the constructor of a decorator (e.g., LazyOpenStream).[\_\_toString()](class-guzzlehttp-psr7-streamdecoratortrait-method-tostring.md)
: string [close()](class-guzzlehttp-psr7-streamdecoratortrait-method-close.md)
: void [detach()](class-guzzlehttp-psr7-streamdecoratortrait-method-detach.md)
: mixed [eof()](class-guzzlehttp-psr7-streamdecoratortrait-method-eof.md)
: bool [getBoundary()](class-guzzlehttp-psr7-multipartstream-method-getboundary.md)
: string [getContents()](class-guzzlehttp-psr7-streamdecoratortrait-method-getcontents.md)
: string [getMetadata()](class-guzzlehttp-psr7-streamdecoratortrait-method-getmetadata.md)
: mixed [getSize()](class-guzzlehttp-psr7-streamdecoratortrait-method-getsize.md)
: int\|null [isReadable()](class-guzzlehttp-psr7-streamdecoratortrait-method-isreadable.md)
: bool [isSeekable()](class-guzzlehttp-psr7-streamdecoratortrait-method-isseekable.md)
: bool [isWritable()](class-guzzlehttp-psr7-multipartstream-method-iswritable.md)
: bool Returns whether or not the stream is writable.[read()](class-guzzlehttp-psr7-streamdecoratortrait-method-read.md)
: string [rewind()](class-guzzlehttp-psr7-streamdecoratortrait-method-rewind.md)
: void [seek()](class-guzzlehttp-psr7-streamdecoratortrait-method-seek.md)
: void [tell()](class-guzzlehttp-psr7-streamdecoratortrait-method-tell.md)
: int [write()](class-guzzlehttp-psr7-streamdecoratortrait-method-write.md)
: int

### Methods  [header link](class-guzzlehttp-psr7-multipartstream-methods.md)

#### \_\_call()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-call.md)

Allow decorators to implement custom methods

`
    public
                    __call(string $method, array<string|int, mixed> $args) : mixed`

##### Parameters

$method
: string$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-guzzlehttp-psr7-multipartstream-method-construct.md)

`
    public
                    __construct([array<string|int, mixed> $elements = [] ][, string $boundary = null ]) : mixed`

##### Parameters

$elements
: array<string\|int, mixed>
= \[\]

Array of associative arrays, each containing a
required "name" key mapping to the form field,
name, a required "contents" key mapping to any
value accepted by Utils::streamFor() (scalar,
null, resource, StreamInterface, Iterator, or
callable), or an array for nested expansion.
Optional keys include "headers" (associative
array of custom headers) and "filename" (string
to send as the filename in the part).
When "contents" is an array, it is recursively
expanded into multiple fields using bracket notation
(e.g., name\[0\]\[key\]). Empty arrays produce no fields.
The "filename" and "headers" options cannot be used
with array contents.

$boundary
: string
= null

You can optionally provide a specific boundary

##### Tags  [header link](class-guzzlehttp-psr7-multipartstream-method-construct-tags.md)

throwsInvalidArgumentException

#### \_\_get()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-get.md)

Magic method used to create a new stream if streams are not added in
the constructor of a decorator (e.g., LazyOpenStream).

`
    public
                    __get(string $name) : StreamInterface`

##### Parameters

$name
: string

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### \_\_toString()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-tostring.md)

`
    public
                    __toString() : string`

##### Return values

string

#### close()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-close.md)

`
    public
                    close() : void`

#### detach()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-detach.md)

`
    public
                    detach() : mixed`

#### eof()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-eof.md)

`
    public
                    eof() : bool`

##### Return values

bool

#### getBoundary()  [header link](class-guzzlehttp-psr7-multipartstream-method-getboundary.md)

`
    public
                    getBoundary() : string`

##### Return values

string

#### getContents()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-getcontents.md)

`
    public
                    getContents() : string`

##### Return values

string

#### getMetadata()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-getmetadata.md)

`
    public
                    getMetadata([mixed $key = null ]) : mixed`

##### Parameters

$key
: mixed
= null

#### getSize()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-getsize.md)

`
    public
                    getSize() : int|null`

##### Return values

int\|null

#### isReadable()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-isreadable.md)

`
    public
                    isReadable() : bool`

##### Return values

bool

#### isSeekable()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-isseekable.md)

`
    public
                    isSeekable() : bool`

##### Return values

bool

#### isWritable()  [header link](class-guzzlehttp-psr7-multipartstream-method-iswritable.md)

Returns whether or not the stream is writable.

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-read.md)

`
    public
                    read(mixed $length) : string`

##### Parameters

$length
: mixed

##### Return values

string

#### rewind()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-rewind.md)

`
    public
                    rewind() : void`

#### seek()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-seek.md)

`
    public
                    seek(mixed $offset[, mixed $whence = SEEK_SET ]) : void`

##### Parameters

$offset
: mixed$whence
: mixed
= SEEK\_SET

#### tell()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-tell.md)

`
    public
                    tell() : int`

##### Return values

int

#### write()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-write.md)

`
    public
                    write(mixed $string) : int`

##### Parameters

$string
: mixed

##### Return values

int
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-psr7-multipartstream-toc-methods.md)
- Methods
  - [\_\_call()](class-guzzlehttp-psr7-streamdecoratortrait-method-call.md)
  - [\_\_construct()](class-guzzlehttp-psr7-multipartstream-method-construct.md)
  - [\_\_get()](class-guzzlehttp-psr7-streamdecoratortrait-method-get.md)
  - [\_\_toString()](class-guzzlehttp-psr7-streamdecoratortrait-method-tostring.md)
  - [close()](class-guzzlehttp-psr7-streamdecoratortrait-method-close.md)
  - [detach()](class-guzzlehttp-psr7-streamdecoratortrait-method-detach.md)
  - [eof()](class-guzzlehttp-psr7-streamdecoratortrait-method-eof.md)
  - [getBoundary()](class-guzzlehttp-psr7-multipartstream-method-getboundary.md)
  - [getContents()](class-guzzlehttp-psr7-streamdecoratortrait-method-getcontents.md)
  - [getMetadata()](class-guzzlehttp-psr7-streamdecoratortrait-method-getmetadata.md)
  - [getSize()](class-guzzlehttp-psr7-streamdecoratortrait-method-getsize.md)
  - [isReadable()](class-guzzlehttp-psr7-streamdecoratortrait-method-isreadable.md)
  - [isSeekable()](class-guzzlehttp-psr7-streamdecoratortrait-method-isseekable.md)
  - [isWritable()](class-guzzlehttp-psr7-multipartstream-method-iswritable.md)
  - [read()](class-guzzlehttp-psr7-streamdecoratortrait-method-read.md)
  - [rewind()](class-guzzlehttp-psr7-streamdecoratortrait-method-rewind.md)
  - [seek()](class-guzzlehttp-psr7-streamdecoratortrait-method-seek.md)
  - [tell()](class-guzzlehttp-psr7-streamdecoratortrait-method-tell.md)
  - [write()](class-guzzlehttp-psr7-streamdecoratortrait-method-write.md)

[Back To Top](class-guzzlehttp-psr7-multipartstream-top.md)
