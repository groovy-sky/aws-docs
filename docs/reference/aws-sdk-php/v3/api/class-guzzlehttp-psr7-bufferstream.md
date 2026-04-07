Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## BufferStream        in package    - [Aws](package-aws.md)       implements  [StreamInterface](class-psr-http-message-streaminterface.md)

FinalYes

Provides a buffer stream that can be written to to fill a buffer, and read
from to remove bytes from the buffer.

This stream returns a "hwm" metadata value that tells upstream consumers
what the configured high water mark of the stream is, or the maximum
preferred size of the buffer.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#toc-interfaces)

[StreamInterface](class-psr-http-message-streaminterface.md)Describes a data stream.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method___construct)
: mixed [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method___toString)
: string Reads all data from the stream into a string, from the beginning to end.[close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_close)
: void Closes the stream and any underlying resources.[detach()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_detach)
: resource\|null Separates any underlying resources from the stream.[eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_eof)
: bool Returns true if the stream is at the end of the stream.[getContents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_getContents)
: string Returns the remaining contents in a string[getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_getMetadata)
: mixed Get stream metadata as an associative array or retrieve a specific key.[getSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_getSize)
: int\|null Get the size of the stream if known.[isReadable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_isReadable)
: bool Returns whether or not the stream is readable.[isSeekable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_isSeekable)
: bool Returns whether or not the stream is seekable.[isWritable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_isWritable)
: bool Returns whether or not the stream is writable.[read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_read)
: string Reads data from the buffer.[rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_rewind)
: void Seek to the beginning of the stream.[seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_seek)
: void Seek to a position in the stream.[tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_tell)
: int Returns the current position of the file read/write pointer[write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_write)
: int Writes data to the buffer.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method___construct)

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

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method___toString)

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

#### close()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_close)

Closes the stream and any underlying resources.

`
    public
                    close() : void`

#### detach()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_detach)

Separates any underlying resources from the stream.

`
    public
                    detach() : resource|null`

After the stream has been detached, the stream is in an unusable state.

##### Return values

resource\|null
—

Underlying PHP stream, if any

#### eof()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_eof)

Returns true if the stream is at the end of the stream.

`
    public
                    eof() : bool`

##### Return values

bool

#### getContents()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_getContents)

Returns the remaining contents in a string

`
    public
                    getContents() : string`

##### Return values

string

#### getMetadata()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_getMetadata)

Get stream metadata as an associative array or retrieve a specific key.

`
    public
                    getMetadata([mixed $key = null ]) : mixed`

##### Parameters

$key
: mixed
= null

Specific metadata to retrieve.

#### getSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_getSize)

Get the size of the stream if known.

`
    public
                    getSize() : int|null`

##### Return values

int\|null
—

Returns the size in bytes if known, or null if unknown.

#### isReadable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_isReadable)

Returns whether or not the stream is readable.

`
    public
                    isReadable() : bool`

##### Return values

bool

#### isSeekable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_isSeekable)

Returns whether or not the stream is seekable.

`
    public
                    isSeekable() : bool`

##### Return values

bool

#### isWritable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_isWritable)

Returns whether or not the stream is writable.

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_read)

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

#### rewind()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_rewind)

Seek to the beginning of the stream.

`
    public
                    rewind() : void`

If the stream is not seekable, this method will raise an exception;
otherwise, it will perform a seek(0).

#### seek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_seek)

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

#### tell()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_tell)

Returns the current position of the file read/write pointer

`
    public
                    tell() : int`

##### Return values

int
—

Position of the file pointer

#### write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html\#method_write)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method___construct)
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method___toString)
  - [close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_close)
  - [detach()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_detach)
  - [eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_eof)
  - [getContents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_getContents)
  - [getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_getMetadata)
  - [getSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_getSize)
  - [isReadable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_isReadable)
  - [isSeekable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_isSeekable)
  - [isWritable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_isWritable)
  - [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_read)
  - [rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_rewind)
  - [seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_seek)
  - [tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_tell)
  - [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#method_write)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.BufferStream.html#top)
