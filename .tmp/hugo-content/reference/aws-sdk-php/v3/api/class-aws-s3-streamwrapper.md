Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## StreamWrapper        in package    - [Aws](package-aws.md)

Amazon S3 stream wrapper to use "s3://<bucket>/<key>" files with PHP
streams, supporting "r", "w", "a", "x".

# Opening "r" (read only) streams:

Read only streams are truly streaming by default and will not allow you to
seek. This is because data read from the stream is not kept in memory or on
the local filesystem. You can force a "r" stream to be seekable by setting
the "seekable" stream context option true. This will allow true streaming of
data from Amazon S3, but will maintain a buffer of previously read bytes in
a 'php://temp' stream to allow seeking to previously read bytes from the
stream.

You may pass any GetObject parameters as 's3' stream context options. These
options will affect how the data is downloaded from Amazon S3.

# Opening "w" and "x" (write only) streams:

Because Amazon S3 requires a Content-Length header, write only streams will
maintain a 'php://temp' stream to buffer data written to the stream until
the stream is flushed (usually by closing the stream with fclose).

You may pass any PutObject parameters as 's3' stream context options. These
options will affect how the data is uploaded to Amazon S3.

When opening an "x" stream, the file must exist on Amazon S3 for the stream
to open successfully.

# Opening "a" (write only append) streams:

Similar to "w" streams, opening append streams requires that the data be
buffered in a "php://temp" stream. Append streams will attempt to download
the contents of an object in Amazon S3, seek to the end of the object, then
allow you to append to the contents of the object. The data will then be
uploaded using a PutObject operation when the stream is flushed (usually
with fclose).

You may pass any GetObject and/or PutObject parameters as 's3' stream
context options. These options will affect how the data is downloaded and
uploaded from Amazon S3.

Stream context options:

- "seekable": Set to true to create a seekable "r" (read only) stream by
using a php://temp stream buffer
- For "unlink" only: Any option that can be passed to the DeleteObject
operation

### Table of Contents  [header link](class-aws-s3-streamwrapper-toc.md)

#### Properties  [header link](class-aws-s3-streamwrapper-toc-properties.md)

[$context](class-aws-s3-streamwrapper-property-context.md)
: resource\|null

#### Methods  [header link](class-aws-s3-streamwrapper-toc-methods.md)

[dir\_closedir()](class-aws-s3-streamwrapper-method-dir-closedir.md)
: bool Close the directory listing handles[dir\_opendir()](class-aws-s3-streamwrapper-method-dir-opendir.md)
: bool Support for opendir().[dir\_readdir()](class-aws-s3-streamwrapper-method-dir-readdir.md)
: string This method is called in response to readdir()[dir\_rewinddir()](class-aws-s3-streamwrapper-method-dir-rewinddir.md)
: bool This method is called in response to rewinddir()[mkdir()](class-aws-s3-streamwrapper-method-mkdir.md)
: bool Support for mkdir().[register()](class-aws-s3-streamwrapper-method-register.md)
: mixed Register the 's3://' stream wrapper[rename()](class-aws-s3-streamwrapper-method-rename.md)
: bool Called in response to rename() to rename a file or directory. Currently
only supports renaming objects.[rmdir()](class-aws-s3-streamwrapper-method-rmdir.md)
: mixed [stream\_cast()](class-aws-s3-streamwrapper-method-stream-cast.md)
: mixed [stream\_close()](class-aws-s3-streamwrapper-method-stream-close.md)
: mixed [stream\_eof()](class-aws-s3-streamwrapper-method-stream-eof.md)
: mixed [stream\_flush()](class-aws-s3-streamwrapper-method-stream-flush.md)
: mixed [stream\_lock()](class-aws-s3-streamwrapper-method-stream-lock.md)
: mixed [stream\_metadata()](class-aws-s3-streamwrapper-method-stream-metadata.md)
: mixed [stream\_open()](class-aws-s3-streamwrapper-method-stream-open.md)
: mixed [stream\_read()](class-aws-s3-streamwrapper-method-stream-read.md)
: mixed [stream\_seek()](class-aws-s3-streamwrapper-method-stream-seek.md)
: mixed [stream\_set\_option()](class-aws-s3-streamwrapper-method-stream-set-option.md)
: mixed [stream\_stat()](class-aws-s3-streamwrapper-method-stream-stat.md)
: mixed [stream\_tell()](class-aws-s3-streamwrapper-method-stream-tell.md)
: mixed [stream\_truncate()](class-aws-s3-streamwrapper-method-stream-truncate.md)
: mixed [stream\_write()](class-aws-s3-streamwrapper-method-stream-write.md)
: mixed [unlink()](class-aws-s3-streamwrapper-method-unlink.md)
: mixed [url\_stat()](class-aws-s3-streamwrapper-method-url-stat.md)
: mixed Provides information for is\_dir, is\_file, filesize, etc. Works on
buckets, keys, and prefixes.

### Properties  [header link](class-aws-s3-streamwrapper-properties.md)

#### $context  [header link](class-aws-s3-streamwrapper-property-context.md)

`
    public
        resource|null
    $context
    `

Stream context (this is set by PHP)

### Methods  [header link](class-aws-s3-streamwrapper-methods.md)

#### dir\_closedir()  [header link](class-aws-s3-streamwrapper-method-dir-closedir.md)

Close the directory listing handles

`
    public
                    dir_closedir() : bool`

##### Return values

bool
—

true on success

#### dir\_opendir()  [header link](class-aws-s3-streamwrapper-method-dir-opendir.md)

Support for opendir().

`
    public
                    dir_opendir(string $path, string $options) : bool`

The opendir() method of the Amazon S3 stream wrapper supports a stream
context option of "listFilter". listFilter must be a callable that
accepts an associative array of object data and returns true if the
object should be yielded when iterating the keys in a bucket.

##### Parameters

$path
: string

The path to the directory
(e.g. "s3://dir\[\]")

$options
: string

Unused option variable

##### Tags  [header link](class-aws-s3-streamwrapper-method-dir-opendir-tags.md)

see[http://www.php.net/manual/en/function.opendir.php](http://www.php.net/manual/en/function.opendir.php)

##### Return values

bool
—

true on success

#### dir\_readdir()  [header link](class-aws-s3-streamwrapper-method-dir-readdir.md)

This method is called in response to readdir()

`
    public
                    dir_readdir() : string`

##### Tags  [header link](class-aws-s3-streamwrapper-method-dir-readdir-tags.md)

link[http://www.php.net/manual/en/function.readdir.php](http://www.php.net/manual/en/function.readdir.php)

##### Return values

string
—

Should return a string representing the next filename, or
false if there is no next file.

#### dir\_rewinddir()  [header link](class-aws-s3-streamwrapper-method-dir-rewinddir.md)

This method is called in response to rewinddir()

`
    public
                    dir_rewinddir() : bool`

##### Return values

bool
—

true on success

#### mkdir()  [header link](class-aws-s3-streamwrapper-method-mkdir.md)

Support for mkdir().

`
    public
                    mkdir(string $path, int $mode, int $options) : bool`

##### Parameters

$path
: string

Directory which should be created.

$mode
: int

Permissions. 700-range permissions map to
ACL\_PUBLIC. 600-range permissions map to
ACL\_AUTH\_READ. All other permissions map to
ACL\_PRIVATE. Expects octal form.

$options
: int

A bitwise mask of values, such as
STREAM\_MKDIR\_RECURSIVE.

##### Tags  [header link](class-aws-s3-streamwrapper-method-mkdir-tags.md)

link[http://www.php.net/manual/en/streamwrapper.mkdir.php](http://www.php.net/manual/en/streamwrapper.mkdir.php)

##### Return values

bool

#### register()  [header link](class-aws-s3-streamwrapper-method-register.md)

Register the 's3://' stream wrapper

`
    public
            static        register(S3ClientInterface $client[, string $protocol = 's3' ][, CacheInterface $cache = null ][, mixed $v2Existence = false ]) : mixed`

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

Client to use with the stream wrapper

$protocol
: string
= 's3'

Protocol to register as.

$cache
: [CacheInterface](class-aws-cacheinterface.md)
= null

Default cache for the protocol.

$v2Existence
: mixed
= false

#### rename()  [header link](class-aws-s3-streamwrapper-method-rename.md)

Called in response to rename() to rename a file or directory. Currently
only supports renaming objects.

`
    public
                    rename(string $path_from, string $path_to) : bool`

##### Parameters

$path\_from
: string

the path to the file to rename

$path\_to
: string

the new path to the file

##### Tags  [header link](class-aws-s3-streamwrapper-method-rename-tags.md)

link[http://www.php.net/manual/en/function.rename.php](http://www.php.net/manual/en/function.rename.php)

##### Return values

bool
—

true if file was successfully renamed

#### rmdir()  [header link](class-aws-s3-streamwrapper-method-rmdir.md)

`
    public
                    rmdir(mixed $path, mixed $options) : mixed`

##### Parameters

$path
: mixed$options
: mixed

#### stream\_cast()  [header link](class-aws-s3-streamwrapper-method-stream-cast.md)

`
    public
                    stream_cast(mixed $cast_as) : mixed`

##### Parameters

$cast\_as
: mixed

#### stream\_close()  [header link](class-aws-s3-streamwrapper-method-stream-close.md)

`
    public
                    stream_close() : mixed`

#### stream\_eof()  [header link](class-aws-s3-streamwrapper-method-stream-eof.md)

`
    public
                    stream_eof() : mixed`

#### stream\_flush()  [header link](class-aws-s3-streamwrapper-method-stream-flush.md)

`
    public
                    stream_flush() : mixed`

#### stream\_lock()  [header link](class-aws-s3-streamwrapper-method-stream-lock.md)

`
    public
                    stream_lock(mixed $operation) : mixed`

##### Parameters

$operation
: mixed

#### stream\_metadata()  [header link](class-aws-s3-streamwrapper-method-stream-metadata.md)

`
    public
                    stream_metadata(mixed $path, mixed $option, mixed $value) : mixed`

##### Parameters

$path
: mixed$option
: mixed$value
: mixed

#### stream\_open()  [header link](class-aws-s3-streamwrapper-method-stream-open.md)

`
    public
                    stream_open(mixed $path, mixed $mode, mixed $options, mixed &$opened_path) : mixed`

##### Parameters

$path
: mixed$mode
: mixed$options
: mixed$opened\_path
: mixed

#### stream\_read()  [header link](class-aws-s3-streamwrapper-method-stream-read.md)

`
    public
                    stream_read(mixed $count) : mixed`

##### Parameters

$count
: mixed

#### stream\_seek()  [header link](class-aws-s3-streamwrapper-method-stream-seek.md)

`
    public
                    stream_seek(mixed $offset[, mixed $whence = SEEK_SET ]) : mixed`

##### Parameters

$offset
: mixed$whence
: mixed
= SEEK\_SET

#### stream\_set\_option()  [header link](class-aws-s3-streamwrapper-method-stream-set-option.md)

`
    public
                    stream_set_option(mixed $option, mixed $arg1, mixed $arg2) : mixed`

##### Parameters

$option
: mixed$arg1
: mixed$arg2
: mixed

#### stream\_stat()  [header link](class-aws-s3-streamwrapper-method-stream-stat.md)

`
    public
                    stream_stat() : mixed`

#### stream\_tell()  [header link](class-aws-s3-streamwrapper-method-stream-tell.md)

`
    public
                    stream_tell() : mixed`

#### stream\_truncate()  [header link](class-aws-s3-streamwrapper-method-stream-truncate.md)

`
    public
                    stream_truncate(mixed $new_size) : mixed`

##### Parameters

$new\_size
: mixed

#### stream\_write()  [header link](class-aws-s3-streamwrapper-method-stream-write.md)

`
    public
                    stream_write(mixed $data) : mixed`

##### Parameters

$data
: mixed

#### unlink()  [header link](class-aws-s3-streamwrapper-method-unlink.md)

`
    public
                    unlink(mixed $path) : mixed`

##### Parameters

$path
: mixed

#### url\_stat()  [header link](class-aws-s3-streamwrapper-method-url-stat.md)

Provides information for is\_dir, is\_file, filesize, etc. Works on
buckets, keys, and prefixes.

`
    public
                    url_stat(mixed $path, mixed $flags) : mixed`

##### Parameters

$path
: mixed$flags
: mixed

##### Tags  [header link](class-aws-s3-streamwrapper-method-url-stat-tags.md)

link[http://www.php.net/manual/en/streamwrapper.url-stat.php](http://www.php.net/manual/en/streamwrapper.url-stat.php)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](class-aws-s3-streamwrapper-toc-properties.md)
  - [Methods](class-aws-s3-streamwrapper-toc-methods.md)
- Properties
  - [$context](class-aws-s3-streamwrapper-property-context.md)
- Methods
  - [dir\_closedir()](class-aws-s3-streamwrapper-method-dir-closedir.md)
  - [dir\_opendir()](class-aws-s3-streamwrapper-method-dir-opendir.md)
  - [dir\_readdir()](class-aws-s3-streamwrapper-method-dir-readdir.md)
  - [dir\_rewinddir()](class-aws-s3-streamwrapper-method-dir-rewinddir.md)
  - [mkdir()](class-aws-s3-streamwrapper-method-mkdir.md)
  - [register()](class-aws-s3-streamwrapper-method-register.md)
  - [rename()](class-aws-s3-streamwrapper-method-rename.md)
  - [rmdir()](class-aws-s3-streamwrapper-method-rmdir.md)
  - [stream\_cast()](class-aws-s3-streamwrapper-method-stream-cast.md)
  - [stream\_close()](class-aws-s3-streamwrapper-method-stream-close.md)
  - [stream\_eof()](class-aws-s3-streamwrapper-method-stream-eof.md)
  - [stream\_flush()](class-aws-s3-streamwrapper-method-stream-flush.md)
  - [stream\_lock()](class-aws-s3-streamwrapper-method-stream-lock.md)
  - [stream\_metadata()](class-aws-s3-streamwrapper-method-stream-metadata.md)
  - [stream\_open()](class-aws-s3-streamwrapper-method-stream-open.md)
  - [stream\_read()](class-aws-s3-streamwrapper-method-stream-read.md)
  - [stream\_seek()](class-aws-s3-streamwrapper-method-stream-seek.md)
  - [stream\_set\_option()](class-aws-s3-streamwrapper-method-stream-set-option.md)
  - [stream\_stat()](class-aws-s3-streamwrapper-method-stream-stat.md)
  - [stream\_tell()](class-aws-s3-streamwrapper-method-stream-tell.md)
  - [stream\_truncate()](class-aws-s3-streamwrapper-method-stream-truncate.md)
  - [stream\_write()](class-aws-s3-streamwrapper-method-stream-write.md)
  - [unlink()](class-aws-s3-streamwrapper-method-unlink.md)
  - [url\_stat()](class-aws-s3-streamwrapper-method-url-stat.md)

[Back To Top](class-aws-s3-streamwrapper-top.md)
