Menu

- [Aws](namespace-aws.md)

## HashingStream        in package    - [Aws](package-aws.md)       implements  [StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)  Uses  [StreamDecoratorTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html)

Stream decorator that calculates a rolling hash of the stream as it is read.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html\#toc-interfaces)

[StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)Describes a data stream.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html\#toc-methods)

[\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___call)
: mixed Allow decorators to implement custom methods[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html#method___construct)
: mixed [\_\_get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___get)
: [StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)Magic method used to create a new stream if streams are not added in
the constructor of a decorator (e.g., LazyOpenStream).[\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___toString)
: string [close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_close)
: void [detach()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_detach)
: mixed [eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_eof)
: bool [getContents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getContents)
: string [getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getMetadata)
: mixed [getSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getSize)
: int\|null [isReadable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isReadable)
: bool [isSeekable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isSeekable)
: bool [isWritable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isWritable)
: bool [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html#method_read)
: string Read data from the stream.[rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_rewind)
: void [seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html#method_seek)
: void Seek to a position in the stream.[tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_tell)
: int [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_write)
: int

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html\#methods)

#### \_\_call()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method___call)

Allow decorators to implement custom methods

`
    public
                    __call(string $method, array<string|int, mixed> $args) : mixed`

##### Parameters

$method
: string$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html\#method___construct)

`
    public
                    __construct(StreamInterface $stream, HashInterface $hash[, callable $onComplete = null ]) : mixed`

##### Parameters

$stream
: [StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)

Stream that is being read.

$hash
: [HashInterface](class-aws-hashinterface.md)

Hash used to calculate checksum.

$onComplete
: callable
= null

Optional function invoked when the
hash calculation is completed.

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

[StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method___toString)

`
    public
                    __toString() : string`

##### Return values

string

#### close()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_close)

`
    public
                    close() : void`

#### detach()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_detach)

`
    public
                    detach() : mixed`

#### eof()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_eof)

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

#### getSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_getSize)

`
    public
                    getSize() : int|null`

##### Return values

int\|null

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

#### read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html\#method_read)

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

#### rewind()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_rewind)

`
    public
                    rewind() : void`

#### seek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html\#method_seek)

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

#### write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html\#method_write)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html#toc-methods)
- Methods
  - [\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___call)
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html#method___construct)
  - [\_\_get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___get)
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method___toString)
  - [close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_close)
  - [detach()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_detach)
  - [eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_eof)
  - [getContents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getContents)
  - [getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getMetadata)
  - [getSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_getSize)
  - [isReadable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isReadable)
  - [isSeekable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isSeekable)
  - [isWritable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_isWritable)
  - [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html#method_read)
  - [rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_rewind)
  - [seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html#method_seek)
  - [tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_tell)
  - [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamDecoratorTrait.html#method_write)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HashingStream.html#top)
