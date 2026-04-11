Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## UploadedFile        in package    - [Aws](package-aws.md)       implements  [UploadedFileInterface](class-psr-http-message-uploadedfileinterface.md)

### Table of Contents  [header link](class-guzzlehttp-psr7-uploadedfile-toc.md)

#### Interfaces  [header link](class-guzzlehttp-psr7-uploadedfile-toc-interfaces.md)

[UploadedFileInterface](class-psr-http-message-uploadedfileinterface.md)Value object representing a file uploaded through an HTTP request.

#### Methods  [header link](class-guzzlehttp-psr7-uploadedfile-toc-methods.md)

[\_\_construct()](class-guzzlehttp-psr7-uploadedfile-method-construct.md)
: mixed [getClientFilename()](class-guzzlehttp-psr7-uploadedfile-method-getclientfilename.md)
: string\|null Retrieve the filename sent by the client.[getClientMediaType()](class-guzzlehttp-psr7-uploadedfile-method-getclientmediatype.md)
: string\|null Retrieve the media type sent by the client.[getError()](class-guzzlehttp-psr7-uploadedfile-method-geterror.md)
: int Retrieve the error associated with the uploaded file.[getSize()](class-guzzlehttp-psr7-uploadedfile-method-getsize.md)
: int\|null Retrieve the file size.[getStream()](class-guzzlehttp-psr7-uploadedfile-method-getstream.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)Retrieve a stream representing the uploaded file.[isMoved()](class-guzzlehttp-psr7-uploadedfile-method-ismoved.md)
: bool [moveTo()](class-guzzlehttp-psr7-uploadedfile-method-moveto.md)
: void Move the uploaded file to a new location.

### Methods  [header link](class-guzzlehttp-psr7-uploadedfile-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-psr7-uploadedfile-method-construct.md)

`
    public
                    __construct(StreamInterface|string|resource $streamOrFile, int|null $size, int $errorStatus[, string|null $clientFilename = null ][, string|null $clientMediaType = null ]) : mixed`

##### Parameters

$streamOrFile
: [StreamInterface](class-psr-http-message-streaminterface.md) \|string\|resource$size
: int\|null$errorStatus
: int$clientFilename
: string\|null
= null$clientMediaType
: string\|null
= null

#### getClientFilename()  [header link](class-guzzlehttp-psr7-uploadedfile-method-getclientfilename.md)

Retrieve the filename sent by the client.

`
    public
                    getClientFilename() : string|null`

Do not trust the value returned by this method. A client could send
a malicious filename with the intention to corrupt or hack your
application.

Implementations SHOULD return the value stored in the "name" key of
the file in the $\_FILES array.

##### Return values

string\|null
—

The filename sent by the client or null if none
was provided.

#### getClientMediaType()  [header link](class-guzzlehttp-psr7-uploadedfile-method-getclientmediatype.md)

Retrieve the media type sent by the client.

`
    public
                    getClientMediaType() : string|null`

Do not trust the value returned by this method. A client could send
a malicious media type with the intention to corrupt or hack your
application.

Implementations SHOULD return the value stored in the "type" key of
the file in the $\_FILES array.

##### Return values

string\|null
—

The media type sent by the client or null if none
was provided.

#### getError()  [header link](class-guzzlehttp-psr7-uploadedfile-method-geterror.md)

Retrieve the error associated with the uploaded file.

`
    public
                    getError() : int`

The return value MUST be one of PHP's UPLOAD\_ERR\_XXX constants.

If the file was uploaded successfully, this method MUST return
UPLOAD\_ERR\_OK.

Implementations SHOULD return the value stored in the "error" key of
the file in the $\_FILES array.

##### Return values

int
—

One of PHP's UPLOAD\_ERR\_XXX constants.

#### getSize()  [header link](class-guzzlehttp-psr7-uploadedfile-method-getsize.md)

Retrieve the file size.

`
    public
                    getSize() : int|null`

Implementations SHOULD return the value stored in the "size" key of
the file in the $\_FILES array if available, as PHP calculates this based
on the actual size transmitted.

##### Return values

int\|null
—

The file size in bytes or null if unknown.

#### getStream()  [header link](class-guzzlehttp-psr7-uploadedfile-method-getstream.md)

Retrieve a stream representing the uploaded file.

`
    public
                    getStream() : StreamInterface`

This method MUST return a StreamInterface instance, representing the
uploaded file. The purpose of this method is to allow utilizing native PHP
stream functionality to manipulate the file upload, such as
stream\_copy\_to\_stream() (though the result will need to be decorated in a
native PHP stream wrapper to work with such functions).

If the moveTo() method has been called previously, this method MUST raise
an exception.

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)
—

Stream representation of the uploaded file.

#### isMoved()  [header link](class-guzzlehttp-psr7-uploadedfile-method-ismoved.md)

`
    public
                    isMoved() : bool`

##### Return values

bool

#### moveTo()  [header link](class-guzzlehttp-psr7-uploadedfile-method-moveto.md)

Move the uploaded file to a new location.

`
    public
                    moveTo(mixed $targetPath) : void`

Use this method as an alternative to move\_uploaded\_file(). This method is
guaranteed to work in both SAPI and non-SAPI environments.
Implementations must determine which environment they are in, and use the
appropriate method (move\_uploaded\_file(), rename(), or a stream
operation) to perform the operation.

$targetPath may be an absolute path, or a relative path. If it is a
relative path, resolution should be the same as used by PHP's rename()
function.

The original file or stream MUST be removed on completion.

If this method is called more than once, any subsequent calls MUST raise
an exception.

When used in an SAPI environment where $\_FILES is populated, when writing
files via moveTo(), is\_uploaded\_file() and move\_uploaded\_file() SHOULD be
used to ensure permissions and upload status are verified correctly.

If you wish to move to a stream, use getStream(), as SAPI operations
cannot guarantee writing to stream destinations.

##### Parameters

$targetPath
: mixed

Path to which to move the uploaded file.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-psr7-uploadedfile-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-psr7-uploadedfile-method-construct.md)
  - [getClientFilename()](class-guzzlehttp-psr7-uploadedfile-method-getclientfilename.md)
  - [getClientMediaType()](class-guzzlehttp-psr7-uploadedfile-method-getclientmediatype.md)
  - [getError()](class-guzzlehttp-psr7-uploadedfile-method-geterror.md)
  - [getSize()](class-guzzlehttp-psr7-uploadedfile-method-getsize.md)
  - [getStream()](class-guzzlehttp-psr7-uploadedfile-method-getstream.md)
  - [isMoved()](class-guzzlehttp-psr7-uploadedfile-method-ismoved.md)
  - [moveTo()](class-guzzlehttp-psr7-uploadedfile-method-moveto.md)

[Back To Top](class-guzzlehttp-psr7-uploadedfile-top.md)
