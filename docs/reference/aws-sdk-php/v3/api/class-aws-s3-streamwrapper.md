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

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#toc)

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#toc-properties)

[$context](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#property_context)
: resource\|null

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#toc-methods)

[dir\_closedir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_dir_closedir)
: bool Close the directory listing handles[dir\_opendir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_dir_opendir)
: bool Support for opendir().[dir\_readdir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_dir_readdir)
: string This method is called in response to readdir()[dir\_rewinddir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_dir_rewinddir)
: bool This method is called in response to rewinddir()[mkdir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_mkdir)
: bool Support for mkdir().[register()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_register)
: mixed Register the 's3://' stream wrapper[rename()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_rename)
: bool Called in response to rename() to rename a file or directory. Currently
only supports renaming objects.[rmdir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_rmdir)
: mixed [stream\_cast()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_cast)
: mixed [stream\_close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_close)
: mixed [stream\_eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_eof)
: mixed [stream\_flush()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_flush)
: mixed [stream\_lock()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_lock)
: mixed [stream\_metadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_metadata)
: mixed [stream\_open()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_open)
: mixed [stream\_read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_read)
: mixed [stream\_seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_seek)
: mixed [stream\_set\_option()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_set_option)
: mixed [stream\_stat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_stat)
: mixed [stream\_tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_tell)
: mixed [stream\_truncate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_truncate)
: mixed [stream\_write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_write)
: mixed [unlink()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_unlink)
: mixed [url\_stat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_url_stat)
: mixed Provides information for is\_dir, is\_file, filesize, etc. Works on
buckets, keys, and prefixes.

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#properties)

#### $context  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#property_context)

`
    public
        resource|null
    $context
    `

Stream context (this is set by PHP)

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#methods)

#### dir\_closedir()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_dir_closedir)

Close the directory listing handles

`
    public
                    dir_closedir() : bool`

##### Return values

bool
—

true on success

#### dir\_opendir()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_dir_opendir)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_dir_opendir\#tags)

see[http://www.php.net/manual/en/function.opendir.php](http://www.php.net/manual/en/function.opendir.php)

##### Return values

bool
—

true on success

#### dir\_readdir()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_dir_readdir)

This method is called in response to readdir()

`
    public
                    dir_readdir() : string`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_dir_readdir\#tags)

link[http://www.php.net/manual/en/function.readdir.php](http://www.php.net/manual/en/function.readdir.php)

##### Return values

string
—

Should return a string representing the next filename, or
false if there is no next file.

#### dir\_rewinddir()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_dir_rewinddir)

This method is called in response to rewinddir()

`
    public
                    dir_rewinddir() : bool`

##### Return values

bool
—

true on success

#### mkdir()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_mkdir)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_mkdir\#tags)

link[http://www.php.net/manual/en/streamwrapper.mkdir.php](http://www.php.net/manual/en/streamwrapper.mkdir.php)

##### Return values

bool

#### register()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_register)

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

#### rename()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_rename)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_rename\#tags)

link[http://www.php.net/manual/en/function.rename.php](http://www.php.net/manual/en/function.rename.php)

##### Return values

bool
—

true if file was successfully renamed

#### rmdir()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_rmdir)

`
    public
                    rmdir(mixed $path, mixed $options) : mixed`

##### Parameters

$path
: mixed$options
: mixed

#### stream\_cast()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_cast)

`
    public
                    stream_cast(mixed $cast_as) : mixed`

##### Parameters

$cast\_as
: mixed

#### stream\_close()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_close)

`
    public
                    stream_close() : mixed`

#### stream\_eof()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_eof)

`
    public
                    stream_eof() : mixed`

#### stream\_flush()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_flush)

`
    public
                    stream_flush() : mixed`

#### stream\_lock()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_lock)

`
    public
                    stream_lock(mixed $operation) : mixed`

##### Parameters

$operation
: mixed

#### stream\_metadata()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_metadata)

`
    public
                    stream_metadata(mixed $path, mixed $option, mixed $value) : mixed`

##### Parameters

$path
: mixed$option
: mixed$value
: mixed

#### stream\_open()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_open)

`
    public
                    stream_open(mixed $path, mixed $mode, mixed $options, mixed &$opened_path) : mixed`

##### Parameters

$path
: mixed$mode
: mixed$options
: mixed$opened\_path
: mixed

#### stream\_read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_read)

`
    public
                    stream_read(mixed $count) : mixed`

##### Parameters

$count
: mixed

#### stream\_seek()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_seek)

`
    public
                    stream_seek(mixed $offset[, mixed $whence = SEEK_SET ]) : mixed`

##### Parameters

$offset
: mixed$whence
: mixed
= SEEK\_SET

#### stream\_set\_option()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_set_option)

`
    public
                    stream_set_option(mixed $option, mixed $arg1, mixed $arg2) : mixed`

##### Parameters

$option
: mixed$arg1
: mixed$arg2
: mixed

#### stream\_stat()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_stat)

`
    public
                    stream_stat() : mixed`

#### stream\_tell()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_tell)

`
    public
                    stream_tell() : mixed`

#### stream\_truncate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_truncate)

`
    public
                    stream_truncate(mixed $new_size) : mixed`

##### Parameters

$new\_size
: mixed

#### stream\_write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_stream_write)

`
    public
                    stream_write(mixed $data) : mixed`

##### Parameters

$data
: mixed

#### unlink()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_unlink)

`
    public
                    unlink(mixed $path) : mixed`

##### Parameters

$path
: mixed

#### url\_stat()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_url_stat)

Provides information for is\_dir, is\_file, filesize, etc. Works on
buckets, keys, and prefixes.

`
    public
                    url_stat(mixed $path, mixed $flags) : mixed`

##### Parameters

$path
: mixed$flags
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html\#method_url_stat\#tags)

link[http://www.php.net/manual/en/streamwrapper.url-stat.php](http://www.php.net/manual/en/streamwrapper.url-stat.php)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#toc-methods)
- Properties
  - [$context](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#property_context)
- Methods
  - [dir\_closedir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_dir_closedir)
  - [dir\_opendir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_dir_opendir)
  - [dir\_readdir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_dir_readdir)
  - [dir\_rewinddir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_dir_rewinddir)
  - [mkdir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_mkdir)
  - [register()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_register)
  - [rename()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_rename)
  - [rmdir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_rmdir)
  - [stream\_cast()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_cast)
  - [stream\_close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_close)
  - [stream\_eof()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_eof)
  - [stream\_flush()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_flush)
  - [stream\_lock()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_lock)
  - [stream\_metadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_metadata)
  - [stream\_open()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_open)
  - [stream\_read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_read)
  - [stream\_seek()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_seek)
  - [stream\_set\_option()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_set_option)
  - [stream\_stat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_stat)
  - [stream\_tell()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_tell)
  - [stream\_truncate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_truncate)
  - [stream\_write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_stream_write)
  - [unlink()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_unlink)
  - [url\_stat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#method_url_stat)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html#top)
