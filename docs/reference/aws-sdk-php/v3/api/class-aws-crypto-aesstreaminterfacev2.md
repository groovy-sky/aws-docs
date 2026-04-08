Menu

- [Aws](namespace-aws.md)
- [Crypto](namespace-aws-crypto.md)

## AesStreamInterfaceV2    extends  [StreamInterface](class-psr-http-message-streaminterface.md)   in    - [Aws](package-aws.md)

Describes a data stream.

Typically, an instance will wrap a PHP stream; this interface provides
a wrapper around the most common operations, including serialization of
the entire stream to a string.

### Table of Contents  [header link](class-aws-crypto-aesstreaminterfacev2-toc.md)

#### Methods  [header link](class-aws-crypto-aesstreaminterfacev2-toc-methods.md)

[\_\_toString()](class-psr-http-message-streaminterface-method-tostring.md)
: string Reads all data from the stream into a string, from the beginning to end.[close()](class-psr-http-message-streaminterface-method-close.md)
: void Closes the stream and any underlying resources.[detach()](class-psr-http-message-streaminterface-method-detach.md)
: resource\|null Separates any underlying resources from the stream.[eof()](class-psr-http-message-streaminterface-method-eof.md)
: bool Returns true if the stream is at the end of the stream.[getContents()](class-psr-http-message-streaminterface-method-getcontents.md)
: string Returns the remaining contents in a string[getCurrentIv()](class-aws-crypto-aesstreaminterfacev2-method-getcurrentiv.md)
: string Returns the IV that should be used to initialize the next block in
encrypt or decrypt.[getMetadata()](class-psr-http-message-streaminterface-method-getmetadata.md)
: array<string\|int, mixed>\|mixed\|null Get stream metadata as an associative array or retrieve a specific key.[getOpenSslName()](class-aws-crypto-aesstreaminterfacev2-method-getopensslname.md)
: string Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.[getSize()](class-psr-http-message-streaminterface-method-getsize.md)
: int\|null Get the size of the stream if known.[getStaticAesName()](class-aws-crypto-aesstreaminterfacev2-method-getstaticaesname.md)
: string Returns an AES recognizable name, such as 'AES/GCM/NoPadding'. V2
interface is accessible from a static context.[isReadable()](class-psr-http-message-streaminterface-method-isreadable.md)
: bool Returns whether or not the stream is readable.[isSeekable()](class-psr-http-message-streaminterface-method-isseekable.md)
: bool Returns whether or not the stream is seekable.[isWritable()](class-psr-http-message-streaminterface-method-iswritable.md)
: bool Returns whether or not the stream is writable.[read()](class-psr-http-message-streaminterface-method-read.md)
: string Read data from the stream.[rewind()](class-psr-http-message-streaminterface-method-rewind.md)
: void Seek to the beginning of the stream.[seek()](class-psr-http-message-streaminterface-method-seek.md)
: void Seek to a position in the stream.[tell()](class-psr-http-message-streaminterface-method-tell.md)
: int Returns the current position of the file read/write pointer[write()](class-psr-http-message-streaminterface-method-write.md)
: int Write data to the stream.

### Methods  [header link](class-aws-crypto-aesstreaminterfacev2-methods.md)

#### \_\_toString()  [header link](class-psr-http-message-streaminterface-method-tostring.md)

Reads all data from the stream into a string, from the beginning to end.

`
    public
                    __toString() : string`

This method MUST attempt to seek to the beginning of the stream before
reading data and read the stream until the end is reached.

Warning: This could attempt to load a large amount of data into memory.

This method MUST NOT raise an exception in order to conform with PHP's
string casting operations.

##### Tags  [header link](class-psr-http-message-streaminterface-method-tostring-tags.md)

see[http://php.net/manual/en/language.oop5.magic.php#object.tostring](http://php.net/manual/en/language.oop5.magic.php)

##### Return values

string

#### close()  [header link](class-psr-http-message-streaminterface-method-close.md)

Closes the stream and any underlying resources.

`
    public
                    close() : void`

#### detach()  [header link](class-psr-http-message-streaminterface-method-detach.md)

Separates any underlying resources from the stream.

`
    public
                    detach() : resource|null`

After the stream has been detached, the stream is in an unusable state.

##### Return values

resource\|null
—

Underlying PHP stream, if any

#### eof()  [header link](class-psr-http-message-streaminterface-method-eof.md)

Returns true if the stream is at the end of the stream.

`
    public
                    eof() : bool`

##### Return values

bool

#### getContents()  [header link](class-psr-http-message-streaminterface-method-getcontents.md)

Returns the remaining contents in a string

`
    public
                    getContents() : string`

##### Tags  [header link](class-psr-http-message-streaminterface-method-getcontents-tags.md)

throwsRuntimeException

if unable to read or an error occurs while
reading.

##### Return values

string

#### getCurrentIv()  [header link](class-aws-crypto-aesstreaminterfacev2-method-getcurrentiv.md)

Returns the IV that should be used to initialize the next block in
encrypt or decrypt.

`
    public
                    getCurrentIv() : string`

##### Return values

string

#### getMetadata()  [header link](class-psr-http-message-streaminterface-method-getmetadata.md)

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

##### Tags  [header link](class-psr-http-message-streaminterface-method-getmetadata-tags.md)

link[http://php.net/manual/en/function.stream-get-meta-data.php](http://php.net/manual/en/function.stream-get-meta-data.php)

##### Return values

array<string\|int, mixed>\|mixed\|null
—

Returns an associative array if no key is
provided. Returns a specific key value if a key is provided and the
value is found, or null if the key is not found.

#### getOpenSslName()  [header link](class-aws-crypto-aesstreaminterfacev2-method-getopensslname.md)

Returns an identifier recognizable by \`openssl\_\*\` functions, such as
\`aes-256-cbc\` or \`aes-128-ctr\`.

`
    public
                    getOpenSslName() : string`

##### Return values

string

#### getSize()  [header link](class-psr-http-message-streaminterface-method-getsize.md)

Get the size of the stream if known.

`
    public
                    getSize() : int|null`

##### Return values

int\|null
—

Returns the size in bytes if known, or null if unknown.

#### getStaticAesName()  [header link](class-aws-crypto-aesstreaminterfacev2-method-getstaticaesname.md)

Returns an AES recognizable name, such as 'AES/GCM/NoPadding'. V2
interface is accessible from a static context.

`
    public
            static        getStaticAesName() : string`

##### Return values

string

#### isReadable()  [header link](class-psr-http-message-streaminterface-method-isreadable.md)

Returns whether or not the stream is readable.

`
    public
                    isReadable() : bool`

##### Return values

bool

#### isSeekable()  [header link](class-psr-http-message-streaminterface-method-isseekable.md)

Returns whether or not the stream is seekable.

`
    public
                    isSeekable() : bool`

##### Return values

bool

#### isWritable()  [header link](class-psr-http-message-streaminterface-method-iswritable.md)

Returns whether or not the stream is writable.

`
    public
                    isWritable() : bool`

##### Return values

bool

#### read()  [header link](class-psr-http-message-streaminterface-method-read.md)

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

##### Tags  [header link](class-psr-http-message-streaminterface-method-read-tags.md)

throwsRuntimeException

if an error occurs.

##### Return values

string
—

Returns the data read from the stream, or an empty string
if no bytes are available.

#### rewind()  [header link](class-psr-http-message-streaminterface-method-rewind.md)

Seek to the beginning of the stream.

`
    public
                    rewind() : void`

If the stream is not seekable, this method will raise an exception;
otherwise, it will perform a seek(0).

##### Tags  [header link](class-psr-http-message-streaminterface-method-rewind-tags.md)

seeseek()link[http://www.php.net/manual/en/function.fseek.php](http://www.php.net/manual/en/function.fseek.php)throwsRuntimeException

on failure.

#### seek()  [header link](class-psr-http-message-streaminterface-method-seek.md)

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

##### Tags  [header link](class-psr-http-message-streaminterface-method-seek-tags.md)

link[http://www.php.net/manual/en/function.fseek.php](http://www.php.net/manual/en/function.fseek.php)throwsRuntimeException

on failure.

#### tell()  [header link](class-psr-http-message-streaminterface-method-tell.md)

Returns the current position of the file read/write pointer

`
    public
                    tell() : int`

##### Tags  [header link](class-psr-http-message-streaminterface-method-tell-tags.md)

throwsRuntimeException

on error.

##### Return values

int
—

Position of the file pointer

#### write()  [header link](class-psr-http-message-streaminterface-method-write.md)

Write data to the stream.

`
    public
                    write(string $string) : int`

##### Parameters

$string
: string

The string that is to be written.

##### Tags  [header link](class-psr-http-message-streaminterface-method-write-tags.md)

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
  - [Constants](class-aws-crypto-aesstreaminterfacev2-toc-constants.md)
  - [Methods](class-aws-crypto-aesstreaminterfacev2-toc-methods.md)
- Methods
  - [\_\_toString()](class-psr-http-message-streaminterface-method-tostring.md)
  - [close()](class-psr-http-message-streaminterface-method-close.md)
  - [detach()](class-psr-http-message-streaminterface-method-detach.md)
  - [eof()](class-psr-http-message-streaminterface-method-eof.md)
  - [getContents()](class-psr-http-message-streaminterface-method-getcontents.md)
  - [getCurrentIv()](class-aws-crypto-aesstreaminterfacev2-method-getcurrentiv.md)
  - [getMetadata()](class-psr-http-message-streaminterface-method-getmetadata.md)
  - [getOpenSslName()](class-aws-crypto-aesstreaminterfacev2-method-getopensslname.md)
  - [getSize()](class-psr-http-message-streaminterface-method-getsize.md)
  - [getStaticAesName()](class-aws-crypto-aesstreaminterfacev2-method-getstaticaesname.md)
  - [isReadable()](class-psr-http-message-streaminterface-method-isreadable.md)
  - [isSeekable()](class-psr-http-message-streaminterface-method-isseekable.md)
  - [isWritable()](class-psr-http-message-streaminterface-method-iswritable.md)
  - [read()](class-psr-http-message-streaminterface-method-read.md)
  - [rewind()](class-psr-http-message-streaminterface-method-rewind.md)
  - [seek()](class-psr-http-message-streaminterface-method-seek.md)
  - [tell()](class-psr-http-message-streaminterface-method-tell.md)
  - [write()](class-psr-http-message-streaminterface-method-write.md)

[Back To Top](class-aws-crypto-aesstreaminterfacev2-top.md)
