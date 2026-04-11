Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## BufferStream        in package    - [Aws](package-aws.md)       implements  [StreamInterface](class-psr-http-message-streaminterface.md)

FinalYes

Provides a buffer stream that can be written to to fill a buffer, and read
from to remove bytes from the buffer.

This stream returns a "hwm" metadata value that tells upstream consumers
what the configured high water mark of the stream is, or the maximum
preferred size of the buffer.

### Table of Contents  [header link](class-guzzlehttp-psr7-bufferstream-toc.md)

#### Interfaces  [header link](class-guzzlehttp-psr7-bufferstream-toc-interfaces.md)

[StreamInterface](class-psr-http-message-streaminterface.md)Describes a data stream.

#### Methods  [header link](class-guzzlehttp-psr7-bufferstream-toc-methods.md)

[\_\_construct()](class-guzzlehttp-psr7-bufferstream-method-construct.md)
: mixed [\_\_toString()](class-guzzlehttp-psr7-bufferstream-method-tostring.md)
: string Reads all data from the stream into a string, from the beginning to end.[close()](class-guzzlehttp-psr7-bufferstream-method-close.md)
: void Closes the stream and any underlying resources.[detach()](class-guzzlehttp-psr7-bufferstream-method-detach.md)
: resource\|null Separates any underlying resources from the stream.[eof()](class-guzzlehttp-psr7-bufferstream-method-eof.md)
: bool Returns true if the stream is at the end of the stream.[getContents()](class-guzzlehttp-psr7-bufferstream-method-getcontents.md)
: string Returns the remaining contents in a string[getMetadata()](class-guzzlehttp-psr7-bufferstream-method-getmetadata.md)
: mixed Get stream metadata as an associative array or retrieve a specific key.[getSize()](class-guzzlehttp-psr7-bufferstream-method-getsize.md)
: int\|null Get the size of the stream if known.[isReadable()](class-guzzlehttp-psr7-bufferstream-method-isreadable.md)
: bool Returns whether or not the stream is readable.[isSeekable()](class-guzzlehttp-psr7-bufferstream-method-isseekable.md)
: bool Returns whether or not the stream is seekable.[isWritable()](class-guzzlehttp-psr7-bufferstream-method-iswritable.md)
: bool Returns whether or not the stream is writable.[read()](class-guzzlehttp-psr7-bufferstream-method-read.md)
: string Reads data from the buffer.[rewind()](class-guzzlehttp-psr7-bufferstream-method-rewind.md)
: void Seek to the beginning of the stream.[seek()](class-guzzlehttp-psr7-bufferstream-method-seek.md)
: void Seek to a position in the stream.[tell()](class-guzzlehttp-psr7-bufferstream-method-tell.md)
: int Returns the current position of the file read/write pointer[write()](class-guzzlehttp-psr7-bufferstream-method-write.md)
: int Writes data to the buffer.

### Methods  [header link](class-guzzlehttp-psr7-bufferstream-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-psr7-bufferstream-method-construct.md)

`
    public
                    __construct([int $hwm = 16384 ]) : mixed`

##### Parameters

$hwm
: int
= 16384

High water mark, representing the preferred maximum
buffer size. If the size of the buffer exceeds the high
water mark, then calls to write will continue to succeed
but will return 0 to inform writers to slow down
until the buffer has been drained by reading from it.

#### \_\_toString()  [header link](class-guzzlehttp-psr7-bufferstream-method-tostring.md)

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

#### close()  [header link](class-guzzlehttp-psr7-bufferstream-method-close.md)

Closes the stream and any underlying resources.

`
    public
                    close() : void`

#### detach()  [header link](class-guzzlehttp-psr7-bufferstream-method-detach.md)

Separates any underlying resources from the stream.

`
    public
                    detach() : resource|null`

After the stream has been detached, the stream is in an unusable state.

##### Return values

resource\|null
—

Underlying PHP stream, if any

#### eof()  [header link](class-guzzlehttp-psr7-bufferstream-method-eof.md)

Returns true if the stream is at the end of the stream.

`
    public
                    eof() : bool`

##### Return values

bool

#### getContents()  [header link](class-guzzlehttp-psr7-bufferstream-method-getcontents.md)

Returns the remaining contents in a string

`
    public
                    getContents() : string`

##### Return values

string

#### getMetadata()  [header link](class-guzzlehttp-psr7-bufferstream-method-getmetadata.md)

Get stream metadata as an associative array or retrieve a specific key.

`
    public
                    getMetadata([mixed $key = null ]) : mixed`

##### Parameters

$key
: mixed
= null

Specific metadata to retrieve.

#### getSize()  [header link](class-guzzlehttp-psr7-bufferstream-method-getsize.md)

Get the size of the stream if known.

`
    public
                    getSize() : int|null`

##### Return values

int\|null
—

Returns the size in bytes if known, or null if unknown.

#### isReadable()  [header link](class-guzzlehttp-psr7-bufferstream-method-isreadable.md)

Returns whether or not the stream is readable.

`
    public
                    isReadable() : bool`

##### Return values

bool

#### isSeekable()  [header link](class-guzzlehttp-psr7-bufferstream-method-isseekable.md)

Returns whether or not the stream is seekable.

`
    public
                    isSeekable() : bool`

##### Return values

bool

#### isWritable()  [header link](class-guzzlehttp-psr7-bufferstream-method-iswritable.md)

Returns whether or not the stream is writable.

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](class-guzzlehttp-psr7-bufferstream-method-read.md)

Reads data from the buffer.

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

#### rewind()  [header link](class-guzzlehttp-psr7-bufferstream-method-rewind.md)

Seek to the beginning of the stream.

`
    public
                    rewind() : void`

If the stream is not seekable, this method will raise an exception;
otherwise, it will perform a seek(0).

#### seek()  [header link](class-guzzlehttp-psr7-bufferstream-method-seek.md)

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

#### tell()  [header link](class-guzzlehttp-psr7-bufferstream-method-tell.md)

Returns the current position of the file read/write pointer

`
    public
                    tell() : int`

##### Return values

int
—

Position of the file pointer

#### write()  [header link](class-guzzlehttp-psr7-bufferstream-method-write.md)

Writes data to the buffer.

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
  - [Methods](class-guzzlehttp-psr7-bufferstream-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-psr7-bufferstream-method-construct.md)
  - [\_\_toString()](class-guzzlehttp-psr7-bufferstream-method-tostring.md)
  - [close()](class-guzzlehttp-psr7-bufferstream-method-close.md)
  - [detach()](class-guzzlehttp-psr7-bufferstream-method-detach.md)
  - [eof()](class-guzzlehttp-psr7-bufferstream-method-eof.md)
  - [getContents()](class-guzzlehttp-psr7-bufferstream-method-getcontents.md)
  - [getMetadata()](class-guzzlehttp-psr7-bufferstream-method-getmetadata.md)
  - [getSize()](class-guzzlehttp-psr7-bufferstream-method-getsize.md)
  - [isReadable()](class-guzzlehttp-psr7-bufferstream-method-isreadable.md)
  - [isSeekable()](class-guzzlehttp-psr7-bufferstream-method-isseekable.md)
  - [isWritable()](class-guzzlehttp-psr7-bufferstream-method-iswritable.md)
  - [read()](class-guzzlehttp-psr7-bufferstream-method-read.md)
  - [rewind()](class-guzzlehttp-psr7-bufferstream-method-rewind.md)
  - [seek()](class-guzzlehttp-psr7-bufferstream-method-seek.md)
  - [tell()](class-guzzlehttp-psr7-bufferstream-method-tell.md)
  - [write()](class-guzzlehttp-psr7-bufferstream-method-write.md)

[Back To Top](class-guzzlehttp-psr7-bufferstream-top.md)
