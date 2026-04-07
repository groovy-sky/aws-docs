Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## AesStreamInterfaceV2    extends  [StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)   in    - [Aws](package-aws.md)

Describes a data stream.

Typically, an instance will wrap a PHP stream; this interface provides
a wrapper around the most common operations, including serialization of
the entire stream to a string.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html\#toc-methods)

[\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method___toString)
: string Reads all data from the stream into a string, from the beginning to end.[close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_close)
: void Closes the stream and any underlying resources.[detach()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_detach)
: resource\|null Separates any underlying resources from the stream.[eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_eof)
: bool Returns true if the stream is at the end of the stream.[getContents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_getContents)
: string Returns the remaining contents in a string[getCurrentIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#method_getCurrentIv)
: string Returns the IV that should be used to initialize the next block in
encrypt or decrypt.[getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_getMetadata)
: array<string\|int, mixed>\|mixed\|null Get stream metadata as an associative array or retrieve a specific key.[getOpenSslName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#method_getOpenSslName)
: string Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.[getSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_getSize)
: int\|null Get the size of the stream if known.[getStaticAesName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#method_getStaticAesName)
: string Returns an AES recognizable name, such as 'AES/GCM/NoPadding'. V2
interface is accessible from a static context.[isReadable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_isReadable)
: bool Returns whether or not the stream is readable.[isSeekable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_isSeekable)
: bool Returns whether or not the stream is seekable.[isWritable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_isWritable)
: bool Returns whether or not the stream is writable.[read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_read)
: string Read data from the stream.[rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_rewind)
: void Seek to the beginning of the stream.[seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_seek)
: void Seek to a position in the stream.[tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_tell)
: int Returns the current position of the file read/write pointer[write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_write)
: int Write data to the stream.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html\#methods)

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method___toString)

Reads all data from the stream into a string, from the beginning to end.

`
    public
                    __toString() : string`

This method MUST attempt to seek to the beginning of the stream before
reading data and read the stream until the end is reached.

Warning: This could attempt to load a large amount of data into memory.

This method MUST NOT raise an exception in order to conform with PHP's
string casting operations.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method___toString\#tags)

see[http://php.net/manual/en/language.oop5.magic.php#object.tostring](http://php.net/manual/en/language.oop5.magic.php)

##### Return values

string

#### close()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_close)

Closes the stream and any underlying resources.

`
    public
                    close() : void`

#### detach()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_detach)

Separates any underlying resources from the stream.

`
    public
                    detach() : resource|null`

After the stream has been detached, the stream is in an unusable state.

##### Return values

resource\|null
—

Underlying PHP stream, if any

#### eof()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_eof)

Returns true if the stream is at the end of the stream.

`
    public
                    eof() : bool`

##### Return values

bool

#### getContents()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_getContents)

Returns the remaining contents in a string

`
    public
                    getContents() : string`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_getContents\#tags)

throwsRuntimeException

if unable to read or an error occurs while
reading.

##### Return values

string

#### getCurrentIv()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html\#method_getCurrentIv)

Returns the IV that should be used to initialize the next block in
encrypt or decrypt.

`
    public
                    getCurrentIv() : string`

##### Return values

string

#### getMetadata()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_getMetadata)

Get stream metadata as an associative array or retrieve a specific key.

`
    public
                    getMetadata([string|null $key = null ]) : array<string|int, mixed>|mixed|null`

The keys returned are identical to the keys returned from PHP's
stream\_get\_meta\_data() function.

##### Parameters

$key
: string\|null
= null

Specific metadata to retrieve.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_getMetadata\#tags)

link[http://php.net/manual/en/function.stream-get-meta-data.php](http://php.net/manual/en/function.stream-get-meta-data.php)

##### Return values

array<string\|int, mixed>\|mixed\|null
—

Returns an associative array if no key is
provided. Returns a specific key value if a key is provided and the
value is found, or null if the key is not found.

#### getOpenSslName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html\#method_getOpenSslName)

Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.

`
    public
                    getOpenSslName() : string`

##### Return values

string

#### getSize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_getSize)

Get the size of the stream if known.

`
    public
                    getSize() : int|null`

##### Return values

int\|null
—

Returns the size in bytes if known, or null if unknown.

#### getStaticAesName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html\#method_getStaticAesName)

Returns an AES recognizable name, such as 'AES/GCM/NoPadding'. V2
interface is accessible from a static context.

`
    public
            static        getStaticAesName() : string`

##### Return values

string

#### isReadable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_isReadable)

Returns whether or not the stream is readable.

`
    public
                    isReadable() : bool`

##### Return values

bool

#### isSeekable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_isSeekable)

Returns whether or not the stream is seekable.

`
    public
                    isSeekable() : bool`

##### Return values

bool

#### isWritable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_isWritable)

Returns whether or not the stream is writable.

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_read)

Read data from the stream.

`
    public
                    read(int $length) : string`

##### Parameters

$length
: int

Read up to $length bytes from the object and return
them. Fewer than $length bytes may be returned if underlying stream
call returns fewer bytes.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_read\#tags)

throwsRuntimeException

if an error occurs.

##### Return values

string
—

Returns the data read from the stream, or an empty string
if no bytes are available.

#### rewind()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_rewind)

Seek to the beginning of the stream.

`
    public
                    rewind() : void`

If the stream is not seekable, this method will raise an exception;
otherwise, it will perform a seek(0).

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_rewind\#tags)

seeseek()link[http://www.php.net/manual/en/function.fseek.php](http://www.php.net/manual/en/function.fseek.php)throwsRuntimeException

on failure.

#### seek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_seek)

Seek to a position in the stream.

`
    public
                    seek(int $offset[, int $whence = SEEK_SET ]) : void`

##### Parameters

$offset
: int

Stream offset

$whence
: int
= SEEK\_SET

Specifies how the cursor position will be calculated
based on the seek offset. Valid values are identical to the built-in
PHP $whence values for `fseek()`. SEEK\_SET: Set position equal to
offset bytes SEEK\_CUR: Set position to current location plus offset
SEEK\_END: Set position to end-of-stream plus offset.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_seek\#tags)

link[http://www.php.net/manual/en/function.fseek.php](http://www.php.net/manual/en/function.fseek.php)throwsRuntimeException

on failure.

#### tell()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_tell)

Returns the current position of the file read/write pointer

`
    public
                    tell() : int`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_tell\#tags)

throwsRuntimeException

on error.

##### Return values

int
—

Position of the file pointer

#### write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_write)

Write data to the stream.

`
    public
                    write(string $string) : int`

##### Parameters

$string
: string

The string that is to be written.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html\#method_write\#tags)

throwsRuntimeException

on failure.

##### Return values

int
—

Returns the number of bytes written to the stream.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#toc-methods)
- Methods
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method___toString)
  - [close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_close)
  - [detach()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_detach)
  - [eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_eof)
  - [getContents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_getContents)
  - [getCurrentIv()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#method_getCurrentIv)
  - [getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_getMetadata)
  - [getOpenSslName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#method_getOpenSslName)
  - [getSize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_getSize)
  - [getStaticAesName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#method_getStaticAesName)
  - [isReadable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_isReadable)
  - [isSeekable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_isSeekable)
  - [isWritable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_isWritable)
  - [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_read)
  - [rewind()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_rewind)
  - [seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_seek)
  - [tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_tell)
  - [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html#method_write)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Crypto.AesStreamInterfaceV2.html#top)
