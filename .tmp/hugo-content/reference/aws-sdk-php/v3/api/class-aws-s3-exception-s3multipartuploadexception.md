Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Exception](namespace-aws-s3-exception.md)

## S3MultipartUploadException     extends [MultipartUploadException](class-aws-exception-multipartuploadexception.md)   in package    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-s3-exception-s3multipartuploadexception-toc.md)

#### Methods  [header link](class-aws-s3-exception-s3multipartuploadexception-toc-methods.md)

[\_\_construct()](class-aws-s3-exception-s3multipartuploadexception-method-construct.md)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[getBucket()](class-aws-s3-exception-s3multipartuploadexception-method-getbucket.md)
: string\|null Get the Bucket information of the transfer object[getKey()](class-aws-s3-exception-s3multipartuploadexception-method-getkey.md)
: string\|null Get the Key information of the transfer object[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getSourceFileName()](class-aws-s3-exception-s3multipartuploadexception-method-getsourcefilename.md)
: string\|null Get the source file name of the transfer object[getState()](class-aws-exception-multipartuploadexception.md#method_getState)
: [UploadState](class-aws-multipart-uploadstate.md)Get the state of the transfer[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](class-aws-s3-exception-s3multipartuploadexception-methods.md)

#### \_\_construct()  [header link](class-aws-s3-exception-s3multipartuploadexception-method-construct.md)

`
    public
                    __construct(UploadState $state[, Exception|array<string|int, mixed> $prev = null ]) : mixed`

##### Parameters

$state
: [UploadState](class-aws-multipart-uploadstate.md)

Upload state at time of the exception.

$prev
: Exception\|array<string\|int, mixed>
= null

Exception being thrown. Could be an array of
AwsExceptions being thrown when uploading parts
for one object, or an instance of AwsException
for a specific Multipart error being thrown in
the MultipartUpload process.

#### appendMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_appendMonitoringEvent)

Append a client-side monitoring event to this object's event list

`
    public
                    appendMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>

#### getBucket()  [header link](class-aws-s3-exception-s3multipartuploadexception-method-getbucket.md)

Get the Bucket information of the transfer object

`
    public
                    getBucket() : string|null`

##### Return values

string\|null
—

Returns null when 'Bucket' information
is unavailable.

#### getKey()  [header link](class-aws-s3-exception-s3multipartuploadexception-method-getkey.md)

Get the Key information of the transfer object

`
    public
                    getKey() : string|null`

##### Return values

string\|null
—

Returns null when 'Key' information
is unavailable.

#### getMonitoringEvents()  [header link](class-aws-hasmonitoringeventstrait.md\#method_getMonitoringEvents)

Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.

`
    public
                    getMonitoringEvents() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getSourceFileName()  [header link](class-aws-s3-exception-s3multipartuploadexception-method-getsourcefilename.md)

Get the source file name of the transfer object

`
    public
                    getSourceFileName() : string|null`

##### Return values

string\|null
—

Returns null when metadata of the stream
wrapped in 'Body' parameter is unavailable.

#### getState()  [header link](class-aws-exception-multipartuploadexception.md\#method_getState)

Get the state of the transfer

`
    public
                    getState() : UploadState`

##### Return values

[UploadState](class-aws-multipart-uploadstate.md)

#### prependMonitoringEvent()  [header link](class-aws-hasmonitoringeventstrait.md\#method_prependMonitoringEvent)

Prepend a client-side monitoring event to this object's event list

`
    public
                    prependMonitoringEvent(array<string|int, mixed> $event) : mixed`

##### Parameters

$event
: array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-exception-s3multipartuploadexception-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-exception-s3multipartuploadexception-method-construct.md)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [getBucket()](class-aws-s3-exception-s3multipartuploadexception-method-getbucket.md)
  - [getKey()](class-aws-s3-exception-s3multipartuploadexception-method-getkey.md)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getSourceFileName()](class-aws-s3-exception-s3multipartuploadexception-method-getsourcefilename.md)
  - [getState()](class-aws-exception-multipartuploadexception.md#method_getState)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](class-aws-s3-exception-s3multipartuploadexception-top.md)
