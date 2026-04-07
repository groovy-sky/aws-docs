Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## CachingStream        in package    - [Aws](package-aws.md)       implements  [StreamInterface](class-psr-http-message-streaminterface.md)  Uses  [StreamDecoratorTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html)

FinalYes

Stream decorator that can cache previously read bytes from a sequentially
read stream.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#toc-interfaces)

[StreamInterface](class-psr-http-message-streaminterface.md)Describes a data stream.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#toc-methods)

[\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___call)
: mixed Allow decorators to implement custom methods[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method___construct)
: mixed We will treat the buffer object as the body of the stream[\_\_get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___get)
: [StreamInterface](class-psr-http-message-streaminterface.md)Magic method used to create a new stream if streams are not added in
the constructor of a decorator (e.g., LazyOpenStream).[\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___toString)
: string [close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_close)
: void Close both the remote stream and buffer stream[detach()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_detach)
: mixed [eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_eof)
: bool Returns true if the stream is at the end of the stream.[getContents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getContents)
: string [getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getMetadata)
: mixed [getSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_getSize)
: int\|null Get the size of the stream if known.[isReadable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isReadable)
: bool [isSeekable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isSeekable)
: bool [isWritable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isWritable)
: bool [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_read)
: string Read data from the stream.[rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_rewind)
: void Seek to the beginning of the stream.[seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_seek)
: void Seek to a position in the stream.[tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_tell)
: int [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_write)
: int Write data to the stream.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#methods)

#### \_\_call()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method___call)

Allow decorators to implement custom methods

`
    public
                    __call(string $method, array<string|int, mixed> $args) : mixed`

##### Parameters

$method
: string$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#method___construct)

We will treat the buffer object as the body of the stream

`
    public
                    __construct(StreamInterface $stream[, StreamInterface $target = null ]) : mixed`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

Stream to cache. The cursor is assumed to be at the beginning of the stream.

$target
: [StreamInterface](class-psr-http-message-streaminterface.md)
= null

Optionally specify where data is cached

#### \_\_get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method___get)

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

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method___toString)

`
    public
                    __toString() : string`

##### Return values

string

#### close()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#method_close)

Close both the remote stream and buffer stream

`
    public
                    close() : void`

#### detach()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_detach)

`
    public
                    detach() : mixed`

#### eof()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#method_eof)

Returns true if the stream is at the end of the stream.

`
    public
                    eof() : bool`

##### Return values

bool

#### getContents()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_getContents)

`
    public
                    getContents() : string`

##### Return values

string

#### getMetadata()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_getMetadata)

`
    public
                    getMetadata([mixed $key = null ]) : mixed`

##### Parameters

$key
: mixed
= null

#### getSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#method_getSize)

Get the size of the stream if known.

`
    public
                    getSize() : int|null`

##### Return values

int\|null
—

Returns the size in bytes if known, or null if unknown.

#### isReadable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_isReadable)

`
    public
                    isReadable() : bool`

##### Return values

bool

#### isSeekable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_isSeekable)

`
    public
                    isSeekable() : bool`

##### Return values

bool

#### isWritable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_isWritable)

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#method_read)

Read data from the stream.

`
    public
                    read(mixed $length) : string`

##### Parameters

$length
: mixed

Read up to $length bytes from the object and return
them. Fewer than $length bytes may be returned if underlying stream
call returns fewer bytes.

##### Return values

string
—

Returns the data read from the stream, or an empty string
if no bytes are available.

#### rewind()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#method_rewind)

Seek to the beginning of the stream.

`
    public
                    rewind() : void`

If the stream is not seekable, this method will raise an exception;
otherwise, it will perform a seek(0).

#### seek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#method_seek)

Seek to a position in the stream.

`
    public
                    seek(mixed $offset[, mixed $whence = SEEK_SET ]) : void`

##### Parameters

$offset
: mixed

Stream offset

$whence
: mixed
= SEEK\_SET

Specifies how the cursor position will be calculated
based on the seek offset. Valid values are identical to the built-in
PHP $whence values for `fseek()`. SEEK\_SET: Set position equal to
offset bytes SEEK\_CUR: Set position to current location plus offset
SEEK\_END: Set position to end-of-stream plus offset.

#### tell()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_tell)

`
    public
                    tell() : int`

##### Return values

int

#### write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html\#method_write)

Write data to the stream.

`
    public
                    write(mixed $string) : int`

##### Parameters

$string
: mixed

The string that is to be written.

##### Return values

int
—

Returns the number of bytes written to the stream.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#toc-methods)
- Methods
  - [\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___call)
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method___construct)
  - [\_\_get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___get)
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___toString)
  - [close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_close)
  - [detach()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_detach)
  - [eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_eof)
  - [getContents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getContents)
  - [getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getMetadata)
  - [getSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_getSize)
  - [isReadable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isReadable)
  - [isSeekable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isSeekable)
  - [isWritable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isWritable)
  - [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_read)
  - [rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_rewind)
  - [seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_seek)
  - [tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_tell)
  - [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#method_write)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.CachingStream.html#top)
