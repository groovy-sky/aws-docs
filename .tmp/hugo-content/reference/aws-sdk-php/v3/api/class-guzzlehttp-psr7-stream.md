Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## Stream        in package    - [Aws](package-aws.md)       implements  [StreamInterface](class-psr-http-message-streaminterface.md)

PHP stream implementation.

### Table of Contents  [header link](class-guzzlehttp-psr7-stream-toc.md)

#### Interfaces  [header link](class-guzzlehttp-psr7-stream-toc-interfaces.md)

[StreamInterface](class-psr-http-message-streaminterface.md)Describes a data stream.

#### Methods  [header link](class-guzzlehttp-psr7-stream-toc-methods.md)

[\_\_construct()](class-guzzlehttp-psr7-stream-method-construct.md)
: mixed This constructor accepts an associative array of options.[\_\_destruct()](class-guzzlehttp-psr7-stream-method-destruct.md)
: mixed Closes the stream when the destructed[\_\_toString()](class-guzzlehttp-psr7-stream-method-tostring.md)
: string Reads all data from the stream into a string, from the beginning to end.[close()](class-guzzlehttp-psr7-stream-method-close.md)
: void Closes the stream and any underlying resources.[detach()](class-guzzlehttp-psr7-stream-method-detach.md)
: resource\|null Separates any underlying resources from the stream.[eof()](class-guzzlehttp-psr7-stream-method-eof.md)
: bool Returns true if the stream is at the end of the stream.[getContents()](class-guzzlehttp-psr7-stream-method-getcontents.md)
: string Returns the remaining contents in a string[getMetadata()](class-guzzlehttp-psr7-stream-method-getmetadata.md)
: mixed Get stream metadata as an associative array or retrieve a specific key.[getSize()](class-guzzlehttp-psr7-stream-method-getsize.md)
: int\|null Get the size of the stream if known.[isReadable()](class-guzzlehttp-psr7-stream-method-isreadable.md)
: bool Returns whether or not the stream is readable.[isSeekable()](class-guzzlehttp-psr7-stream-method-isseekable.md)
: bool Returns whether or not the stream is seekable.[isWritable()](class-guzzlehttp-psr7-stream-method-iswritable.md)
: bool Returns whether or not the stream is writable.[read()](class-guzzlehttp-psr7-stream-method-read.md)
: string Read data from the stream.[rewind()](class-guzzlehttp-psr7-stream-method-rewind.md)
: void Seek to the beginning of the stream.[seek()](class-guzzlehttp-psr7-stream-method-seek.md)
: void Seek to a position in the stream.[tell()](class-guzzlehttp-psr7-stream-method-tell.md)
: int Returns the current position of the file read/write pointer[write()](class-guzzlehttp-psr7-stream-method-write.md)
: int Write data to the stream.

### Methods  [header link](class-guzzlehttp-psr7-stream-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-psr7-stream-method-construct.md)

This constructor accepts an associative array of options.

`
    public
                    __construct(resource $stream[, array{size?: int, metadata?: array} $options = [] ]) : mixed`

- size: (int) If a read stream would otherwise have an indeterminate
size, but the size is known due to foreknowledge, then you can
provide that size, in bytes.
- metadata: (array) Any additional metadata to return when the metadata
of the stream is accessed.

##### Parameters

$stream
: resource

Stream resource to wrap.

$options
: array{size?: int, metadata?: array}
= \[\]

Associative array of options.

##### Tags  [header link](class-guzzlehttp-psr7-stream-method-construct-tags.md)

throwsInvalidArgumentException

if the stream is not a stream resource

#### \_\_destruct()  [header link](class-guzzlehttp-psr7-stream-method-destruct.md)

Closes the stream when the destructed

`
    public
                    __destruct() : mixed`

#### \_\_toString()  [header link](class-guzzlehttp-psr7-stream-method-tostring.md)

Reads all data from the stream into a string, from the beginning to end.

`
    public
                    __toString() : string`

This method MUST attempt to seek to the beginning of the stream before
reading data and read the stream until the end is reached.

Warning: This could attempt to load a large amount of data into memory.

This method MUST NOT raise an exception in order to conform with PHP's
string casting operations.

##### Return values

string

#### close()  [header link](class-guzzlehttp-psr7-stream-method-close.md)

Closes the stream and any underlying resources.

`
    public
                    close() : void`

#### detach()  [header link](class-guzzlehttp-psr7-stream-method-detach.md)

Separates any underlying resources from the stream.

`
    public
                    detach() : resource|null`

After the stream has been detached, the stream is in an unusable state.

##### Return values

resource\|null
—

Underlying PHP stream, if any

#### eof()  [header link](class-guzzlehttp-psr7-stream-method-eof.md)

Returns true if the stream is at the end of the stream.

`
    public
                    eof() : bool`

##### Return values

bool

#### getContents()  [header link](class-guzzlehttp-psr7-stream-method-getcontents.md)

Returns the remaining contents in a string

`
    public
                    getContents() : string`

##### Return values

string

#### getMetadata()  [header link](class-guzzlehttp-psr7-stream-method-getmetadata.md)

Get stream metadata as an associative array or retrieve a specific key.

`
    public
                    getMetadata([mixed $key = null ]) : mixed`

##### Parameters

$key
: mixed
= null

Specific metadata to retrieve.

#### getSize()  [header link](class-guzzlehttp-psr7-stream-method-getsize.md)

Get the size of the stream if known.

`
    public
                    getSize() : int|null`

##### Return values

int\|null
—

Returns the size in bytes if known, or null if unknown.

#### isReadable()  [header link](class-guzzlehttp-psr7-stream-method-isreadable.md)

Returns whether or not the stream is readable.

`
    public
                    isReadable() : bool`

##### Return values

bool

#### isSeekable()  [header link](class-guzzlehttp-psr7-stream-method-isseekable.md)

Returns whether or not the stream is seekable.

`
    public
                    isSeekable() : bool`

##### Return values

bool

#### isWritable()  [header link](class-guzzlehttp-psr7-stream-method-iswritable.md)

Returns whether or not the stream is writable.

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](class-guzzlehttp-psr7-stream-method-read.md)

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

#### rewind()  [header link](class-guzzlehttp-psr7-stream-method-rewind.md)

Seek to the beginning of the stream.

`
    public
                    rewind() : void`

If the stream is not seekable, this method will raise an exception;
otherwise, it will perform a seek(0).

#### seek()  [header link](class-guzzlehttp-psr7-stream-method-seek.md)

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

#### tell()  [header link](class-guzzlehttp-psr7-stream-method-tell.md)

Returns the current position of the file read/write pointer

`
    public
                    tell() : int`

##### Return values

int
—

Position of the file pointer

#### write()  [header link](class-guzzlehttp-psr7-stream-method-write.md)

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
  - [Methods](class-guzzlehttp-psr7-stream-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-psr7-stream-method-construct.md)
  - [\_\_destruct()](class-guzzlehttp-psr7-stream-method-destruct.md)
  - [\_\_toString()](class-guzzlehttp-psr7-stream-method-tostring.md)
  - [close()](class-guzzlehttp-psr7-stream-method-close.md)
  - [detach()](class-guzzlehttp-psr7-stream-method-detach.md)
  - [eof()](class-guzzlehttp-psr7-stream-method-eof.md)
  - [getContents()](class-guzzlehttp-psr7-stream-method-getcontents.md)
  - [getMetadata()](class-guzzlehttp-psr7-stream-method-getmetadata.md)
  - [getSize()](class-guzzlehttp-psr7-stream-method-getsize.md)
  - [isReadable()](class-guzzlehttp-psr7-stream-method-isreadable.md)
  - [isSeekable()](class-guzzlehttp-psr7-stream-method-isseekable.md)
  - [isWritable()](class-guzzlehttp-psr7-stream-method-iswritable.md)
  - [read()](class-guzzlehttp-psr7-stream-method-read.md)
  - [rewind()](class-guzzlehttp-psr7-stream-method-rewind.md)
  - [seek()](class-guzzlehttp-psr7-stream-method-seek.md)
  - [tell()](class-guzzlehttp-psr7-stream-method-tell.md)
  - [write()](class-guzzlehttp-psr7-stream-method-write.md)

[Back To Top](class-guzzlehttp-psr7-stream-top.md)
