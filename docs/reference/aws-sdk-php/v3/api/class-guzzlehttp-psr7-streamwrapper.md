Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## StreamWrapper        in package    - [Aws](package-aws.md)

FinalYes

Converts Guzzle streams into PHP stream resources.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#tags)

see[https://www.php.net/streamwrapper](https://www.php.net/streamwrapper)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#toc)

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#toc-properties)

[$context](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#property_context)
: resource

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#toc-methods)

[createStreamContext()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_createStreamContext)
: resource Creates a stream context that can be used to open a stream as a php stream resource.[getResource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_getResource)
: resource Returns a resource representing the stream.[register()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_register)
: void Registers the stream wrapper if needed[stream\_cast()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_cast)
: resource\|false [stream\_eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_eof)
: bool [stream\_open()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_open)
: bool [stream\_read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_read)
: string [stream\_seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_seek)
: bool [stream\_stat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_stat)
: array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}\|false [stream\_tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_tell)
: int [stream\_write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_write)
: int [url\_stat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_url_stat)
: array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#properties)

#### $context  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#property_context)

`
    public
        resource
    $context
    `

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#methods)

#### createStreamContext()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_createStreamContext)

Creates a stream context that can be used to open a stream as a php stream resource.

`
    public
            static        createStreamContext(StreamInterface $stream) : resource`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

##### Return values

resource

#### getResource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_getResource)

Returns a resource representing the stream.

`
    public
            static        getResource(StreamInterface $stream) : resource`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

The stream to get a resource for

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_getResource\#tags)

throwsInvalidArgumentException

if stream is not readable or writable

##### Return values

resource

#### register()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_register)

Registers the stream wrapper if needed

`
    public
            static        register() : void`

#### stream\_cast()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_stream_cast)

`
    public
                    stream_cast(int $cast_as) : resource|false`

##### Parameters

$cast\_as
: int

##### Return values

resource\|false

#### stream\_eof()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_stream_eof)

`
    public
                    stream_eof() : bool`

##### Return values

bool

#### stream\_open()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_stream_open)

`
    public
                    stream_open(string $path, string $mode, int $options[, string|null &$opened_path = null ]) : bool`

##### Parameters

$path
: string$mode
: string$options
: int$opened\_path
: string\|null
= null

##### Return values

bool

#### stream\_read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_stream_read)

`
    public
                    stream_read(int $count) : string`

##### Parameters

$count
: int

##### Return values

string

#### stream\_seek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_stream_seek)

`
    public
                    stream_seek(int $offset, int $whence) : bool`

##### Parameters

$offset
: int$whence
: int

##### Return values

bool

#### stream\_stat()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_stream_stat)

`
    public
                    stream_stat() : array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}|false`

##### Return values

array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}\|false

#### stream\_tell()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_stream_tell)

`
    public
                    stream_tell() : int`

##### Return values

int

#### stream\_write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_stream_write)

`
    public
                    stream_write(string $data) : int`

##### Parameters

$data
: string

##### Return values

int

#### url\_stat()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html\#method_url_stat)

`
    public
                    url_stat(string $path, int $flags) : array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}`

##### Parameters

$path
: string$flags
: int

##### Return values

array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#toc-methods)
- Properties
  - [$context](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#property_context)
- Methods
  - [createStreamContext()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_createStreamContext)
  - [getResource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_getResource)
  - [register()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_register)
  - [stream\_cast()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_cast)
  - [stream\_eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_eof)
  - [stream\_open()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_open)
  - [stream\_read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_read)
  - [stream\_seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_seek)
  - [stream\_stat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_stat)
  - [stream\_tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_tell)
  - [stream\_write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_stream_write)
  - [url\_stat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#method_url_stat)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.StreamWrapper.html#top)
