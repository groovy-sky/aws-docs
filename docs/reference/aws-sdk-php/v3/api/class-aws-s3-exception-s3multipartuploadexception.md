Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [Exception](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.exception.html)

## S3MultipartUploadException     extends [MultipartUploadException](class-aws-exception-multipartuploadexception.md)   in package    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#method___construct)
: mixed [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
: mixed Append a client-side monitoring event to this object's event list[getBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#method_getBucket)
: string\|null Get the Bucket information of the transfer object[getKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#method_getKey)
: string\|null Get the Key information of the transfer object[getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
: array<string\|int, mixed> Get client-side monitoring events attached to this object. Each event is
represented as an associative array within the returned array.[getSourceFileName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#method_getSourceFileName)
: string\|null Get the source file name of the transfer object[getState()](class-aws-exception-multipartuploadexception.md#method_getState)
: [UploadState](class-aws-multipart-uploadstate.md)Get the state of the transfer[prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)
: mixed Prepend a client-side monitoring event to this object's event list

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html\#method___construct)

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

#### getBucket()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html\#method_getBucket)

Get the Bucket information of the transfer object

`
    public
                    getBucket() : string|null`

##### Return values

string\|null
—

Returns null when 'Bucket' information
is unavailable.

#### getKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html\#method_getKey)

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

#### getSourceFileName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html\#method_getSourceFileName)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#method___construct)
  - [appendMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_appendMonitoringEvent)
  - [getBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#method_getBucket)
  - [getKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#method_getKey)
  - [getMonitoringEvents()](class-aws-hasmonitoringeventstrait.md#method_getMonitoringEvents)
  - [getSourceFileName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#method_getSourceFileName)
  - [getState()](class-aws-exception-multipartuploadexception.md#method_getState)
  - [prependMonitoringEvent()](class-aws-hasmonitoringeventstrait.md#method_prependMonitoringEvent)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.S3MultipartUploadException.html#top)
