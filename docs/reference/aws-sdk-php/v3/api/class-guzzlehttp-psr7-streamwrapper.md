Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## StreamWrapper        in package    - [Aws](package-aws.md)

FinalYes

Converts Guzzle streams into PHP stream resources.

##### Tags  [header link](class-guzzlehttp-psr7-streamwrapper-tags.md)

see[https://www.php.net/streamwrapper](https://www.php.net/streamwrapper)

### Table of Contents  [header link](class-guzzlehttp-psr7-streamwrapper-toc.md)

#### Properties  [header link](class-guzzlehttp-psr7-streamwrapper-toc-properties.md)

[$context](class-guzzlehttp-psr7-streamwrapper-property-context.md)
: resource

#### Methods  [header link](class-guzzlehttp-psr7-streamwrapper-toc-methods.md)

[createStreamContext()](class-guzzlehttp-psr7-streamwrapper-method-createstreamcontext.md)
: resource Creates a stream context that can be used to open a stream as a php stream resource.[getResource()](class-guzzlehttp-psr7-streamwrapper-method-getresource.md)
: resource Returns a resource representing the stream.[register()](class-guzzlehttp-psr7-streamwrapper-method-register.md)
: void Registers the stream wrapper if needed[stream\_cast()](class-guzzlehttp-psr7-streamwrapper-method-stream-cast.md)
: resource\|false [stream\_eof()](class-guzzlehttp-psr7-streamwrapper-method-stream-eof.md)
: bool [stream\_open()](class-guzzlehttp-psr7-streamwrapper-method-stream-open.md)
: bool [stream\_read()](class-guzzlehttp-psr7-streamwrapper-method-stream-read.md)
: string [stream\_seek()](class-guzzlehttp-psr7-streamwrapper-method-stream-seek.md)
: bool [stream\_stat()](class-guzzlehttp-psr7-streamwrapper-method-stream-stat.md)
: array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}\|false [stream\_tell()](class-guzzlehttp-psr7-streamwrapper-method-stream-tell.md)
: int [stream\_write()](class-guzzlehttp-psr7-streamwrapper-method-stream-write.md)
: int [url\_stat()](class-guzzlehttp-psr7-streamwrapper-method-url-stat.md)
: array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}

### Properties  [header link](class-guzzlehttp-psr7-streamwrapper-properties.md)

#### $context  [header link](class-guzzlehttp-psr7-streamwrapper-property-context.md)

`
    public
        resource
    $context
    `

### Methods  [header link](class-guzzlehttp-psr7-streamwrapper-methods.md)

#### createStreamContext()  [header link](class-guzzlehttp-psr7-streamwrapper-method-createstreamcontext.md)

Creates a stream context that can be used to open a stream as a php stream resource.

`
    public
            static        createStreamContext(StreamInterface $stream) : resource`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

##### Return values

resource

#### getResource()  [header link](class-guzzlehttp-psr7-streamwrapper-method-getresource.md)

Returns a resource representing the stream.

`
    public
            static        getResource(StreamInterface $stream) : resource`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)

The stream to get a resource for

##### Tags  [header link](class-guzzlehttp-psr7-streamwrapper-method-getresource-tags.md)

throwsInvalidArgumentException

if stream is not readable or writable

##### Return values

resource

#### register()  [header link](class-guzzlehttp-psr7-streamwrapper-method-register.md)

Registers the stream wrapper if needed

`
    public
            static        register() : void`

#### stream\_cast()  [header link](class-guzzlehttp-psr7-streamwrapper-method-stream-cast.md)

`
    public
                    stream_cast(int $cast_as) : resource|false`

##### Parameters

$cast\_as
: int

##### Return values

resource\|false

#### stream\_eof()  [header link](class-guzzlehttp-psr7-streamwrapper-method-stream-eof.md)

`
    public
                    stream_eof() : bool`

##### Return values

bool

#### stream\_open()  [header link](class-guzzlehttp-psr7-streamwrapper-method-stream-open.md)

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

#### stream\_read()  [header link](class-guzzlehttp-psr7-streamwrapper-method-stream-read.md)

`
    public
                    stream_read(int $count) : string`

##### Parameters

$count
: int

##### Return values

string

#### stream\_seek()  [header link](class-guzzlehttp-psr7-streamwrapper-method-stream-seek.md)

`
    public
                    stream_seek(int $offset, int $whence) : bool`

##### Parameters

$offset
: int$whence
: int

##### Return values

bool

#### stream\_stat()  [header link](class-guzzlehttp-psr7-streamwrapper-method-stream-stat.md)

`
    public
                    stream_stat() : array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}|false`

##### Return values

array{dev: int, ino: int, mode: int, nlink: int, uid: int, gid: int, rdev: int, size: int, atime: int, mtime: int, ctime: int, blksize: int, blocks: int}\|false

#### stream\_tell()  [header link](class-guzzlehttp-psr7-streamwrapper-method-stream-tell.md)

`
    public
                    stream_tell() : int`

##### Return values

int

#### stream\_write()  [header link](class-guzzlehttp-psr7-streamwrapper-method-stream-write.md)

`
    public
                    stream_write(string $data) : int`

##### Parameters

$data
: string

##### Return values

int

#### url\_stat()  [header link](class-guzzlehttp-psr7-streamwrapper-method-url-stat.md)

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
  - [Properties](class-guzzlehttp-psr7-streamwrapper-toc-properties.md)
  - [Methods](class-guzzlehttp-psr7-streamwrapper-toc-methods.md)
- Properties
  - [$context](class-guzzlehttp-psr7-streamwrapper-property-context.md)
- Methods
  - [createStreamContext()](class-guzzlehttp-psr7-streamwrapper-method-createstreamcontext.md)
  - [getResource()](class-guzzlehttp-psr7-streamwrapper-method-getresource.md)
  - [register()](class-guzzlehttp-psr7-streamwrapper-method-register.md)
  - [stream\_cast()](class-guzzlehttp-psr7-streamwrapper-method-stream-cast.md)
  - [stream\_eof()](class-guzzlehttp-psr7-streamwrapper-method-stream-eof.md)
  - [stream\_open()](class-guzzlehttp-psr7-streamwrapper-method-stream-open.md)
  - [stream\_read()](class-guzzlehttp-psr7-streamwrapper-method-stream-read.md)
  - [stream\_seek()](class-guzzlehttp-psr7-streamwrapper-method-stream-seek.md)
  - [stream\_stat()](class-guzzlehttp-psr7-streamwrapper-method-stream-stat.md)
  - [stream\_tell()](class-guzzlehttp-psr7-streamwrapper-method-stream-tell.md)
  - [stream\_write()](class-guzzlehttp-psr7-streamwrapper-method-stream-write.md)
  - [url\_stat()](class-guzzlehttp-psr7-streamwrapper-method-url-stat.md)

[Back To Top](class-guzzlehttp-psr7-streamwrapper-top.md)
