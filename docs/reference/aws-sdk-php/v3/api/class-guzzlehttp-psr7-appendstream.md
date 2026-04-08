Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## AppendStream        in package    - [Aws](package-aws.md)       implements  [StreamInterface](class-psr-http-message-streaminterface.md)

FinalYes

Reads from multiple streams, one after the other.

This is a read-only stream decorator.

### Table of Contents  [header link](class-guzzlehttp-psr7-appendstream-toc.md)

#### Interfaces  [header link](class-guzzlehttp-psr7-appendstream-toc-interfaces.md)

[StreamInterface](class-psr-http-message-streaminterface.md)Describes a data stream.

#### Methods  [header link](class-guzzlehttp-psr7-appendstream-toc-methods.md)

[\_\_construct()](class-guzzlehttp-psr7-appendstream-method-construct.md)
: mixed [\_\_toString()](class-guzzlehttp-psr7-appendstream-method-tostring.md)
: string Reads all data from the stream into a string, from the beginning to end.[addStream()](class-guzzlehttp-psr7-appendstream-method-addstream.md)
: void Add a stream to the AppendStream[close()](class-guzzlehttp-psr7-appendstream-method-close.md)
: void Closes each attached stream.[detach()](class-guzzlehttp-psr7-appendstream-method-detach.md)
: resource\|null Detaches each attached stream.[eof()](class-guzzlehttp-psr7-appendstream-method-eof.md)
: bool Returns true if the stream is at the end of the stream.[getContents()](class-guzzlehttp-psr7-appendstream-method-getcontents.md)
: string Returns the remaining contents in a string[getMetadata()](class-guzzlehttp-psr7-appendstream-method-getmetadata.md)
: mixed Get stream metadata as an associative array or retrieve a specific key.[getSize()](class-guzzlehttp-psr7-appendstream-method-getsize.md)
: int\|null Tries to calculate the size by adding the size of each stream.[isReadable()](class-guzzlehttp-psr7-appendstream-method-isreadable.md)
: bool Returns whether or not the stream is readable.[isSeekable()](class-guzzlehttp-psr7-appendstream-method-isseekable.md)
: bool Returns whether or not the stream is seekable.[isWritable()](class-guzzlehttp-psr7-appendstream-method-iswritable.md)
: bool Returns whether or not the stream is writable.[read()](class-guzzlehttp-psr7-appendstream-method-read.md)
: string Reads from all of the appended streams until the length is met or EOF.[rewind()](class-guzzlehttp-psr7-appendstream-method-rewind.md)
: void Seek to the beginning of the stream.[seek()](class-guzzlehttp-psr7-appendstream-method-seek.md)
: void Attempts to seek to the given position. Only supports SEEK\_SET.[tell()](class-guzzlehttp-psr7-appendstream-method-tell.md)
: int Returns the current position of the file read/write pointer[write()](class-guzzlehttp-psr7-appendstream-method-write.md)
: int Write data to the stream.

### Methods  [header link](class-guzzlehttp-psr7-appendstream-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-psr7-appendstream-method-construct.md)

`
    public
                    __construct([array<string|int, StreamInterface> $streams = [] ]) : mixed`

##### Parameters

$streams
: array<string\|int, [StreamInterface](class-psr-http-message-streaminterface.md) >
= \[\]

Streams to decorate. Each stream must
be readable.

#### \_\_toString()  [header link](class-guzzlehttp-psr7-appendstream-method-tostring.md)

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

#### addStream()  [header link](class-guzzlehttp-psr7-appendstream-method-addstream.md)

Add a stream to the AppendStream

`
    public
                    addStream(StreamInterface $stream) : void`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

Stream to append. Must be readable.

##### Tags  [header link](class-guzzlehttp-psr7-appendstream-method-addstream-tags.md)

throwsInvalidArgumentException

if the stream is not readable

#### close()  [header link](class-guzzlehttp-psr7-appendstream-method-close.md)

Closes each attached stream.

`
    public
                    close() : void`

#### detach()  [header link](class-guzzlehttp-psr7-appendstream-method-detach.md)

Detaches each attached stream.

`
    public
                    detach() : resource|null`

Returns null as it's not clear which underlying stream resource to return.

##### Return values

resource\|null
—

Underlying PHP stream, if any

#### eof()  [header link](class-guzzlehttp-psr7-appendstream-method-eof.md)

Returns true if the stream is at the end of the stream.

`
    public
                    eof() : bool`

##### Return values

bool

#### getContents()  [header link](class-guzzlehttp-psr7-appendstream-method-getcontents.md)

Returns the remaining contents in a string

`
    public
                    getContents() : string`

##### Return values

string

#### getMetadata()  [header link](class-guzzlehttp-psr7-appendstream-method-getmetadata.md)

Get stream metadata as an associative array or retrieve a specific key.

`
    public
                    getMetadata([mixed $key = null ]) : mixed`

##### Parameters

$key
: mixed
= null

Specific metadata to retrieve.

#### getSize()  [header link](class-guzzlehttp-psr7-appendstream-method-getsize.md)

Tries to calculate the size by adding the size of each stream.

`
    public
                    getSize() : int|null`

If any of the streams do not return a valid number, then the size of the
append stream cannot be determined and null is returned.

##### Return values

int\|null
—

Returns the size in bytes if known, or null if unknown.

#### isReadable()  [header link](class-guzzlehttp-psr7-appendstream-method-isreadable.md)

Returns whether or not the stream is readable.

`
    public
                    isReadable() : bool`

##### Return values

bool

#### isSeekable()  [header link](class-guzzlehttp-psr7-appendstream-method-isseekable.md)

Returns whether or not the stream is seekable.

`
    public
                    isSeekable() : bool`

##### Return values

bool

#### isWritable()  [header link](class-guzzlehttp-psr7-appendstream-method-iswritable.md)

Returns whether or not the stream is writable.

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](class-guzzlehttp-psr7-appendstream-method-read.md)

Reads from all of the appended streams until the length is met or EOF.

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

#### rewind()  [header link](class-guzzlehttp-psr7-appendstream-method-rewind.md)

Seek to the beginning of the stream.

`
    public
                    rewind() : void`

If the stream is not seekable, this method will raise an exception;
otherwise, it will perform a seek(0).

#### seek()  [header link](class-guzzlehttp-psr7-appendstream-method-seek.md)

Attempts to seek to the given position. Only supports SEEK\_SET.

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

#### tell()  [header link](class-guzzlehttp-psr7-appendstream-method-tell.md)

Returns the current position of the file read/write pointer

`
    public
                    tell() : int`

##### Return values

int
—

Position of the file pointer

#### write()  [header link](class-guzzlehttp-psr7-appendstream-method-write.md)

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
  - [Methods](class-guzzlehttp-psr7-appendstream-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-psr7-appendstream-method-construct.md)
  - [\_\_toString()](class-guzzlehttp-psr7-appendstream-method-tostring.md)
  - [addStream()](class-guzzlehttp-psr7-appendstream-method-addstream.md)
  - [close()](class-guzzlehttp-psr7-appendstream-method-close.md)
  - [detach()](class-guzzlehttp-psr7-appendstream-method-detach.md)
  - [eof()](class-guzzlehttp-psr7-appendstream-method-eof.md)
  - [getContents()](class-guzzlehttp-psr7-appendstream-method-getcontents.md)
  - [getMetadata()](class-guzzlehttp-psr7-appendstream-method-getmetadata.md)
  - [getSize()](class-guzzlehttp-psr7-appendstream-method-getsize.md)
  - [isReadable()](class-guzzlehttp-psr7-appendstream-method-isreadable.md)
  - [isSeekable()](class-guzzlehttp-psr7-appendstream-method-isseekable.md)
  - [isWritable()](class-guzzlehttp-psr7-appendstream-method-iswritable.md)
  - [read()](class-guzzlehttp-psr7-appendstream-method-read.md)
  - [rewind()](class-guzzlehttp-psr7-appendstream-method-rewind.md)
  - [seek()](class-guzzlehttp-psr7-appendstream-method-seek.md)
  - [tell()](class-guzzlehttp-psr7-appendstream-method-tell.md)
  - [write()](class-guzzlehttp-psr7-appendstream-method-write.md)

[Back To Top](class-guzzlehttp-psr7-appendstream-top.md)
