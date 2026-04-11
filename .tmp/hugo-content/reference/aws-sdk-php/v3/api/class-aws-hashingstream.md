Menu

- [Aws](namespace-aws.md)

## HashingStream        in package    - [Aws](package-aws.md)       implements  [StreamInterface](class-psr-http-message-streaminterface.md)  Uses  [StreamDecoratorTrait](class-guzzlehttp-psr7-streamdecoratortrait.md)

Stream decorator that calculates a rolling hash of the stream as it is read.

### Table of Contents  [header link](class-aws-hashingstream-toc.md)

#### Interfaces  [header link](class-aws-hashingstream-toc-interfaces.md)

[StreamInterface](class-psr-http-message-streaminterface.md)Describes a data stream.

#### Methods  [header link](class-aws-hashingstream-toc-methods.md)

[\_\_call()](class-guzzlehttp-psr7-streamdecoratortrait-method-call.md)
: mixed Allow decorators to implement custom methods[\_\_construct()](class-aws-hashingstream-method-construct.md)
: mixed [\_\_get()](class-guzzlehttp-psr7-streamdecoratortrait-method-get.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)Magic method used to create a new stream if streams are not added in
the constructor of a decorator (e.g., LazyOpenStream).[\_\_toString()](class-guzzlehttp-psr7-streamdecoratortrait-method-tostring.md)
: string [close()](class-guzzlehttp-psr7-streamdecoratortrait-method-close.md)
: void [detach()](class-guzzlehttp-psr7-streamdecoratortrait-method-detach.md)
: mixed [eof()](class-guzzlehttp-psr7-streamdecoratortrait-method-eof.md)
: bool [getContents()](class-guzzlehttp-psr7-streamdecoratortrait-method-getcontents.md)
: string [getMetadata()](class-guzzlehttp-psr7-streamdecoratortrait-method-getmetadata.md)
: mixed [getSize()](class-guzzlehttp-psr7-streamdecoratortrait-method-getsize.md)
: int\|null [isReadable()](class-guzzlehttp-psr7-streamdecoratortrait-method-isreadable.md)
: bool [isSeekable()](class-guzzlehttp-psr7-streamdecoratortrait-method-isseekable.md)
: bool [isWritable()](class-guzzlehttp-psr7-streamdecoratortrait-method-iswritable.md)
: bool [read()](class-aws-hashingstream-method-read.md)
: string Read data from the stream.[rewind()](class-guzzlehttp-psr7-streamdecoratortrait-method-rewind.md)
: void [seek()](class-aws-hashingstream-method-seek.md)
: void Seek to a position in the stream.[tell()](class-guzzlehttp-psr7-streamdecoratortrait-method-tell.md)
: int [write()](class-guzzlehttp-psr7-streamdecoratortrait-method-write.md)
: int

### Methods  [header link](class-aws-hashingstream-methods.md)

#### \_\_call()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-call.md)

Allow decorators to implement custom methods

`
    public
                    __call(string $method, array<string|int, mixed> $args) : mixed`

##### Parameters

$method
: string$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-hashingstream-method-construct.md)

`
    public
                    __construct(StreamInterface $stream, HashInterface $hash[, callable $onComplete = null ]) : mixed`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

Stream that is being read.

$hash
: [HashInterface](class-aws-hashinterface.md)

Hash used to calculate checksum.

$onComplete
: callable
= null

Optional function invoked when the
hash calculation is completed.

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

#### isWritable()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-iswritable.md)

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](class-aws-hashingstream-method-read.md)

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

#### rewind()  [header link](class-guzzlehttp-psr7-streamdecoratortrait-method-rewind.md)

`
    public
                    rewind() : void`

#### seek()  [header link](class-aws-hashingstream-method-seek.md)

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
  - [Methods](class-aws-hashingstream-toc-methods.md)
- Methods
  - [\_\_call()](class-guzzlehttp-psr7-streamdecoratortrait-method-call.md)
  - [\_\_construct()](class-aws-hashingstream-method-construct.md)
  - [\_\_get()](class-guzzlehttp-psr7-streamdecoratortrait-method-get.md)
  - [\_\_toString()](class-guzzlehttp-psr7-streamdecoratortrait-method-tostring.md)
  - [close()](class-guzzlehttp-psr7-streamdecoratortrait-method-close.md)
  - [detach()](class-guzzlehttp-psr7-streamdecoratortrait-method-detach.md)
  - [eof()](class-guzzlehttp-psr7-streamdecoratortrait-method-eof.md)
  - [getContents()](class-guzzlehttp-psr7-streamdecoratortrait-method-getcontents.md)
  - [getMetadata()](class-guzzlehttp-psr7-streamdecoratortrait-method-getmetadata.md)
  - [getSize()](class-guzzlehttp-psr7-streamdecoratortrait-method-getsize.md)
  - [isReadable()](class-guzzlehttp-psr7-streamdecoratortrait-method-isreadable.md)
  - [isSeekable()](class-guzzlehttp-psr7-streamdecoratortrait-method-isseekable.md)
  - [isWritable()](class-guzzlehttp-psr7-streamdecoratortrait-method-iswritable.md)
  - [read()](class-aws-hashingstream-method-read.md)
  - [rewind()](class-guzzlehttp-psr7-streamdecoratortrait-method-rewind.md)
  - [seek()](class-aws-hashingstream-method-seek.md)
  - [tell()](class-guzzlehttp-psr7-streamdecoratortrait-method-tell.md)
  - [write()](class-guzzlehttp-psr7-streamdecoratortrait-method-write.md)

[Back To Top](class-aws-hashingstream-top.md)
