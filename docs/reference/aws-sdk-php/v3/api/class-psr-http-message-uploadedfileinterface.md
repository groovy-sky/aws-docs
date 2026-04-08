Menu

- [Psr](namespace-psr.md)
- [Http](namespace-psr-http.md)
- [Message](namespace-psr-http-message.md)

## UploadedFileInterface     in    - [Aws](package-aws.md)

Value object representing a file uploaded through an HTTP request.

Instances of this interface are considered immutable; all methods that
might change state MUST be implemented such that they retain the internal
state of the current instance and return an instance that contains the
changed state.

### Table of Contents  [header link](class-psr-http-message-uploadedfileinterface-toc.md)

#### Methods  [header link](class-psr-http-message-uploadedfileinterface-toc-methods.md)

[getClientFilename()](class-psr-http-message-uploadedfileinterface-method-getclientfilename.md)
: string\|null Retrieve the filename sent by the client.[getClientMediaType()](class-psr-http-message-uploadedfileinterface-method-getclientmediatype.md)
: string\|null Retrieve the media type sent by the client.[getError()](class-psr-http-message-uploadedfileinterface-method-geterror.md)
: int Retrieve the error associated with the uploaded file.[getSize()](class-psr-http-message-uploadedfileinterface-method-getsize.md)
: int\|null Retrieve the file size.[getStream()](class-psr-http-message-uploadedfileinterface-method-getstream.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)Retrieve a stream representing the uploaded file.[moveTo()](class-psr-http-message-uploadedfileinterface-method-moveto.md)
: void Move the uploaded file to a new location.

### Methods  [header link](class-psr-http-message-uploadedfileinterface-methods.md)

#### getClientFilename()  [header link](class-psr-http-message-uploadedfileinterface-method-getclientfilename.md)

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

#### getClientMediaType()  [header link](class-psr-http-message-uploadedfileinterface-method-getclientmediatype.md)

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

#### getError()  [header link](class-psr-http-message-uploadedfileinterface-method-geterror.md)

Retrieve the error associated with the uploaded file.

`
    public
                    getError() : int`

The return value MUST be one of PHP's UPLOAD\_ERR\_XXX constants.

If the file was uploaded successfully, this method MUST return
UPLOAD\_ERR\_OK.

Implementations SHOULD return the value stored in the "error" key of
the file in the $\_FILES array.

##### Tags  [header link](class-psr-http-message-uploadedfileinterface-method-geterror-tags.md)

see[http://php.net/manual/en/features.file-upload.errors.php](http://php.net/manual/en/features.file-upload.errors.php)

##### Return values

int
—

One of PHP's UPLOAD\_ERR\_XXX constants.

#### getSize()  [header link](class-psr-http-message-uploadedfileinterface-method-getsize.md)

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

#### getStream()  [header link](class-psr-http-message-uploadedfileinterface-method-getstream.md)

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

##### Tags  [header link](class-psr-http-message-uploadedfileinterface-method-getstream-tags.md)

throwsRuntimeException

in cases when no stream is available or can be
created.

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)
—

Stream representation of the uploaded file.

#### moveTo()  [header link](class-psr-http-message-uploadedfileinterface-method-moveto.md)

Move the uploaded file to a new location.

`
    public
                    moveTo(string $targetPath) : void`

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
: string

Path to which to move the uploaded file.

##### Tags  [header link](class-psr-http-message-uploadedfileinterface-method-moveto-tags.md)

see[http://php.net/is\_uploaded\_file](http://php.net/is_uploaded_file)see[http://php.net/move\_uploaded\_file](http://php.net/move_uploaded_file)throwsInvalidArgumentException

if the $targetPath specified is invalid.

throwsRuntimeException

on any error during the move operation, or on
the second or subsequent call to the method.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-psr-http-message-uploadedfileinterface-toc-constants.md)
  - [Methods](class-psr-http-message-uploadedfileinterface-toc-methods.md)
- Methods
  - [getClientFilename()](class-psr-http-message-uploadedfileinterface-method-getclientfilename.md)
  - [getClientMediaType()](class-psr-http-message-uploadedfileinterface-method-getclientmediatype.md)
  - [getError()](class-psr-http-message-uploadedfileinterface-method-geterror.md)
  - [getSize()](class-psr-http-message-uploadedfileinterface-method-getsize.md)
  - [getStream()](class-psr-http-message-uploadedfileinterface-method-getstream.md)
  - [moveTo()](class-psr-http-message-uploadedfileinterface-method-moveto.md)

[Back To Top](class-psr-http-message-uploadedfileinterface-top.md)
