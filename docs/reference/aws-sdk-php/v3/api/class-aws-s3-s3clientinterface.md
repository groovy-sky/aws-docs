Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## S3ClientInterface    extends  [AwsClientInterface](class-aws-awsclientinterface.md)   in    - [Aws](package-aws.md)

**Amazon Simple Storage Service** client.

### Table of Contents  [header link](class-aws-s3-s3clientinterface-toc.md)

#### Methods  [header link](class-aws-s3-s3clientinterface-toc-methods.md)

[\_\_call()](class-aws-awsclientinterface.md#method___call)
: [ResultInterface](class-aws-resultinterface.md)Creates and executes a command for an operation by name.[copy()](class-aws-s3-s3clientinterface-method-copy.md)
: [ResultInterface](class-aws-resultinterface.md)Copy an object of any size to a different location.[copyAsync()](class-aws-s3-s3clientinterface-method-copyasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Copy an object of any size to a different location asynchronously.[createPresignedRequest()](class-aws-s3-s3clientinterface-method-createpresignedrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Create a pre-signed URL for the given S3 command object.[deleteMatchingObjects()](class-aws-s3-s3clientinterface-method-deletematchingobjects.md)
: mixed Deletes objects from Amazon S3 that match the result of a ListObjects
operation. For example, this allows you to do things like delete all
objects that match a specific key prefix.[deleteMatchingObjectsAsync()](class-aws-s3-s3clientinterface-method-deletematchingobjectsasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Deletes objects from Amazon S3 that match the result of a ListObjects
operation. For example, this allows you to do things like delete all
objects that match a specific key prefix.[determineBucketRegion()](class-aws-s3-s3clientinterface-method-determinebucketregion.md)
: string Returns the region in which a given bucket is located.[determineBucketRegionAsync()](class-aws-s3-s3clientinterface-method-determinebucketregionasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise fulfilled with the region in which a given bucket is
located.[doesBucketExist()](class-aws-s3-s3clientinterface-method-doesbucketexist.md)
: bool [doesBucketExistV2()](class-aws-s3-s3clientinterface-method-doesbucketexistv2.md)
: bool Determines whether or not a bucket exists by name. This method uses S3's
HeadBucket operation and requires the relevant bucket permissions in the
default case to prevent errors.[doesObjectExist()](class-aws-s3-s3clientinterface-method-doesobjectexist.md)
: bool [doesObjectExistV2()](class-aws-s3-s3clientinterface-method-doesobjectexistv2.md)
: bool Determines whether or not an object exists by name. This method uses S3's HeadObject
operation and requires the relevant bucket and object permissions to prevent errors.[downloadBucket()](class-aws-s3-s3clientinterface-method-downloadbucket.md)
: mixed Downloads a bucket to the local filesystem[downloadBucketAsync()](class-aws-s3-s3clientinterface-method-downloadbucketasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Downloads a bucket to the local filesystem[execute()](class-aws-awsclientinterface.md#method_execute)
: [ResultInterface](class-aws-resultinterface.md)Execute a single command.[executeAsync()](class-aws-awsclientinterface.md#method_executeAsync)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Execute a command asynchronously.[getApi()](class-aws-awsclientinterface.md#method_getApi)
: [Service](class-aws-api-service.md)Get the service description associated with the client.[getCommand()](class-aws-awsclientinterface.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)Create a command for an operation name.[getConfig()](class-aws-awsclientinterface.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclientinterface.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclientinterface.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getHandlerList()](class-aws-awsclientinterface.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclientinterface.md#method_getIterator)
: IteratorGet a resource iterator for the specified operation.[getObjectUrl()](class-aws-s3-s3clientinterface-method-getobjecturl.md)
: string Returns the URL to an object identified by its bucket and key.[getPaginator()](class-aws-awsclientinterface.md#method_getPaginator)
: [ResultPaginator](class-aws-resultpaginator.md)Get a result paginator for the specified operation.[getRegion()](class-aws-awsclientinterface.md#method_getRegion)
: string Get the region to which the client is configured to send requests.[getWaiter()](class-aws-awsclientinterface.md#method_getWaiter)
: [Waiter](class-aws-waiter.md)Get a waiter that waits until a resource is in a particular state.[registerStreamWrapper()](class-aws-s3-s3clientinterface-method-registerstreamwrapper.md)
: mixed Register the Amazon S3 stream wrapper with this client instance.[registerStreamWrapperV2()](class-aws-s3-s3clientinterface-method-registerstreamwrapperv2.md)
: mixed Registers the Amazon S3 stream wrapper with this client instance.[upload()](class-aws-s3-s3clientinterface-method-upload.md)
: [ResultInterface](class-aws-resultinterface.md)Upload a file, stream, or string to a bucket.[uploadAsync()](class-aws-s3-s3clientinterface-method-uploadasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Upload a file, stream, or string to a bucket asynchronously.[uploadDirectory()](class-aws-s3-s3clientinterface-method-uploaddirectory.md)
: mixed Recursively uploads all files in a given directory to a given bucket.[uploadDirectoryAsync()](class-aws-s3-s3clientinterface-method-uploaddirectoryasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Recursively uploads all files in a given directory to a given bucket.[waitUntil()](class-aws-awsclientinterface.md#method_waitUntil)
: void Wait until a resource is in a particular state.

### Methods  [header link](class-aws-s3-s3clientinterface-methods.md)

#### \_\_call()  [header link](class-aws-awsclientinterface.md\#method___call)

Creates and executes a command for an operation by name.

`
    public
                    __call(string $name, array<string|int, mixed> $arguments) : ResultInterface`

Suffixing an operation name with "Async" will return a
promise that can be used to execute commands asynchronously.

##### Parameters

$name
: string

Name of the command to execute.

$arguments
: array<string\|int, mixed>

Arguments to pass to the getCommand method.

##### Tags  [header link](class-aws-awsclientinterface.md\#method___call\#tags)

throwsException

##### Return values

[ResultInterface](class-aws-resultinterface.md)

#### copy()  [header link](class-aws-s3-s3clientinterface-method-copy.md)

Copy an object of any size to a different location.

`
    public
                    copy(string $fromBucket, string $fromKey, string $destBucket, string $destKey[, string $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : ResultInterface`

If the upload size exceeds the maximum allowable size for direct S3
copying, a multipart copy will be used.

The options array accepts the following options:

- before\_upload: (callable) Callback to invoke before any upload
operations during the upload process. The callback should have a
function signature like `function (Aws\Command $command) {...}`.
- concurrency: (int, default=int(5)) Maximum number of concurrent
`UploadPart` operations allowed during a multipart upload.
- params: (array, default=array(\[\])) Custom parameters to use with the
upload. For single uploads, they must correspond to those used for the
`CopyObject` operation. For multipart uploads, they correspond to the
parameters of the `CreateMultipartUpload` operation.
- part\_size: (int) Part size to use when doing a multipart upload.

##### Parameters

$fromBucket
: string

Bucket where the copy source resides.

$fromKey
: string

Key of the copy source.

$destBucket
: string

Bucket to which to copy the object.

$destKey
: string

Key to which to copy the object.

$acl
: string
= 'private'

ACL to apply to the copy (default: private).

$options
: array<string\|int, mixed>
= \[\]

Options used to configure the upload process.

##### Tags  [header link](class-aws-s3-s3clientinterface-method-copy-tags.md)

seeMultipartCopy

for more info about multipart uploads.

##### Return values

[ResultInterface](class-aws-resultinterface.md)
—

Returns the result of the copy.

#### copyAsync()  [header link](class-aws-s3-s3clientinterface-method-copyasync.md)

Copy an object of any size to a different location asynchronously.

`
    public
                    copyAsync(string $fromBucket, string $fromKey, string $destBucket, string $destKey[, string $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : PromiseInterface`

##### Parameters

$fromBucket
: string

Bucket where the copy source resides.

$fromKey
: string

Key of the copy source.

$destBucket
: string

Bucket to which to copy the object.

$destKey
: string

Key to which to copy the object.

$acl
: string
= 'private'

ACL to apply to the copy (default: private).

$options
: array<string\|int, mixed>
= \[\]

Options used to configure the upload process.

##### Tags  [header link](class-aws-s3-s3clientinterface-method-copyasync-tags.md)

seeself::copy

for more info about the parameters above.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
—

Returns a promise that will be fulfilled
with the result of the copy.

#### createPresignedRequest()  [header link](class-aws-s3-s3clientinterface-method-createpresignedrequest.md)

Create a pre-signed URL for the given S3 command object.

`
    public
                    createPresignedRequest(CommandInterface $command, int|string|DateTimeInterface $expires[, array<string|int, mixed> $options = [] ]) : RequestInterface`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

Command to create a pre-signed
URL for.

$expires
: int\|string\|DateTimeInterface

The time at which the URL should
expire. This can be a Unix
timestamp, a PHP DateTime object,
or a string that can be evaluated
by strtotime().

$options
: array<string\|int, mixed>
= \[\]

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### deleteMatchingObjects()  [header link](class-aws-s3-s3clientinterface-method-deletematchingobjects.md)

Deletes objects from Amazon S3 that match the result of a ListObjects
operation. For example, this allows you to do things like delete all
objects that match a specific key prefix.

`
    public
                    deleteMatchingObjects(string $bucket[, string $prefix = '' ][, string $regex = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: string

Bucket that contains the object keys

$prefix
: string
= ''

Optionally delete only objects under this key prefix

$regex
: string
= ''

Delete only objects that match this regex

$options
: array<string\|int, mixed>
= \[\]

Aws\\S3\\BatchDelete options array.

##### Tags  [header link](class-aws-s3-s3clientinterface-method-deletematchingobjects-tags.md)

seeS3Client::listObjectsthrowsRuntimeException

if no prefix and no regex is given

#### deleteMatchingObjectsAsync()  [header link](class-aws-s3-s3clientinterface-method-deletematchingobjectsasync.md)

Deletes objects from Amazon S3 that match the result of a ListObjects
operation. For example, this allows you to do things like delete all
objects that match a specific key prefix.

`
    public
                    deleteMatchingObjectsAsync(string $bucket[, string $prefix = '' ][, string $regex = '' ][, array<string|int, mixed> $options = [] ]) : PromiseInterface`

##### Parameters

$bucket
: string

Bucket that contains the object keys

$prefix
: string
= ''

Optionally delete only objects under this key prefix

$regex
: string
= ''

Delete only objects that match this regex

$options
: array<string\|int, mixed>
= \[\]

Aws\\S3\\BatchDelete options array.

##### Tags  [header link](class-aws-s3-s3clientinterface-method-deletematchingobjectsasync-tags.md)

seeS3Client::listObjects

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
—

A promise that is settled when matching
objects are deleted.

#### determineBucketRegion()  [header link](class-aws-s3-s3clientinterface-method-determinebucketregion.md)

Returns the region in which a given bucket is located.

`
    public
                    determineBucketRegion(string $bucketName) : string`

##### Parameters

$bucketName
: string

##### Return values

string

#### determineBucketRegionAsync()  [header link](class-aws-s3-s3clientinterface-method-determinebucketregionasync.md)

Returns a promise fulfilled with the region in which a given bucket is
located.

`
    public
                    determineBucketRegionAsync(string $bucketName) : PromiseInterface`

##### Parameters

$bucketName
: string

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### doesBucketExist()  [header link](class-aws-s3-s3clientinterface-method-doesbucketexist.md)

`
    public
                    doesBucketExist(string $bucket) : bool`

##### Parameters

$bucket
: string

The name of the bucket

##### Tags  [header link](class-aws-s3-s3clientinterface-method-doesbucketexist-tags.md)

deprecated

Use doesBucketExistV2() instead

Determines whether or not a bucket exists by name.

##### Return values

bool

#### doesBucketExistV2()  [header link](class-aws-s3-s3clientinterface-method-doesbucketexistv2.md)

Determines whether or not a bucket exists by name. This method uses S3's
HeadBucket operation and requires the relevant bucket permissions in the
default case to prevent errors.

`
    public
                    doesBucketExistV2(string $bucket, bool $accept403) : bool`

##### Parameters

$bucket
: string

The name of the bucket

$accept403
: bool

Set to true for this method to return true in the case of
invalid bucket-level permissions. Credentials MUST be valid
to avoid inaccuracies. Using the default value of false will
cause an exception to be thrown instead.

##### Tags  [header link](class-aws-s3-s3clientinterface-method-doesbucketexistv2-tags.md)

throws[S3Exception](class-aws-s3-exception-s3exception.md) \|Exception

if there is an unhandled exception

##### Return values

bool

#### doesObjectExist()  [header link](class-aws-s3-s3clientinterface-method-doesobjectexist.md)

`
    public
                    doesObjectExist(string $bucket, string $key[, array<string|int, mixed> $options = [] ]) : bool`

##### Parameters

$bucket
: string

The name of the bucket

$key
: string

The key of the object

$options
: array<string\|int, mixed>
= \[\]

Additional options available in the HeadObject
operation (e.g., VersionId).

##### Tags  [header link](class-aws-s3-s3clientinterface-method-doesobjectexist-tags.md)

deprecated

Use doesObjectExistV2() instead

Determines whether or not an object exists by name.

##### Return values

bool

#### doesObjectExistV2()  [header link](class-aws-s3-s3clientinterface-method-doesobjectexistv2.md)

Determines whether or not an object exists by name. This method uses S3's HeadObject
operation and requires the relevant bucket and object permissions to prevent errors.

`
    public
                    doesObjectExistV2(string $bucket, string $key[, bool $includeDeleteMarkers = false ][, array<string|int, mixed> $options = [] ]) : bool`

##### Parameters

$bucket
: string

The name of the bucket

$key
: string

The key of the object

$includeDeleteMarkers
: bool
= false

Set to true to consider delete markers
existing objects. Using the default value
of false will ignore delete markers and
return false.

$options
: array<string\|int, mixed>
= \[\]

Additional options available in the HeadObject
operation (e.g., VersionId).

##### Tags  [header link](class-aws-s3-s3clientinterface-method-doesobjectexistv2-tags.md)

throws[S3Exception](class-aws-s3-exception-s3exception.md) \|Exception

if there is an unhandled exception

##### Return values

bool

#### downloadBucket()  [header link](class-aws-s3-s3clientinterface-method-downloadbucket.md)

Downloads a bucket to the local filesystem

`
    public
                    downloadBucket(string $directory, string $bucket[, string $keyPrefix = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: string

Directory to download to

$bucket
: string

Bucket to download from

$keyPrefix
: string
= ''

Only download objects that use this key prefix

$options
: array<string\|int, mixed>
= \[\]

Options available in Aws\\S3\\Transfer::\_\_construct

#### downloadBucketAsync()  [header link](class-aws-s3-s3clientinterface-method-downloadbucketasync.md)

Downloads a bucket to the local filesystem

`
    public
                    downloadBucketAsync(string $directory, string $bucket[, string $keyPrefix = '' ][, array<string|int, mixed> $options = [] ]) : PromiseInterface`

##### Parameters

$directory
: string

Directory to download to

$bucket
: string

Bucket to download from

$keyPrefix
: string
= ''

Only download objects that use this key prefix

$options
: array<string\|int, mixed>
= \[\]

Options available in Aws\\S3\\Transfer::\_\_construct

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
—

A promise that is settled when the download is
complete.

#### execute()  [header link](class-aws-awsclientinterface.md\#method_execute)

Execute a single command.

`
    public
                    execute(CommandInterface $command) : ResultInterface`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

Command to execute

##### Tags  [header link](class-aws-awsclientinterface.md\#method_execute\#tags)

throwsException

##### Return values

[ResultInterface](class-aws-resultinterface.md)

#### executeAsync()  [header link](class-aws-awsclientinterface.md\#method_executeAsync)

Execute a command asynchronously.

`
    public
                    executeAsync(CommandInterface $command) : PromiseInterface`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

Command to execute

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getApi()  [header link](class-aws-awsclientinterface.md\#method_getApi)

Get the service description associated with the client.

`
    public
                    getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getCommand()  [header link](class-aws-awsclientinterface.md\#method_getCommand)

Create a command for an operation name.

`
    public
                    getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

Special keys may be set on the command to control how it behaves,
including:

- @http: Associative array of transfer specific options to apply to the
request that is serialized for this command. Available keys include
"proxy", "verify", "timeout", "connect\_timeout", "debug", "delay", and
"headers".

##### Parameters

$name
: string

Name of the operation to use in the command

$args
: array<string\|int, mixed>
= \[\]

Arguments to pass to the command

##### Tags  [header link](class-aws-awsclientinterface.md\#method_getCommand\#tags)

throwsInvalidArgumentException

if no command can be found by name

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-awsclientinterface.md\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([string|null $option = null ]) : mixed|null`

##### Parameters

$option
: string\|null
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](class-aws-awsclientinterface.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclientinterface.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getHandlerList()  [header link](class-aws-awsclientinterface.md\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclientinterface.md\#method_getIterator)

Get a resource iterator for the specified operation.

`
    public
                    getIterator(string $name[, array<string|int, mixed> $args = [] ]) : Iterator`

##### Parameters

$name
: string

Name of the iterator to retrieve.

$args
: array<string\|int, mixed>
= \[\]

Command arguments to use with each command.

##### Tags  [header link](class-aws-awsclientinterface.md\#method_getIterator\#tags)

throwsUnexpectedValueException

if the iterator config is invalid.

##### Return values

Iterator

#### getObjectUrl()  [header link](class-aws-s3-s3clientinterface-method-getobjecturl.md)

Returns the URL to an object identified by its bucket and key.

`
    public
                    getObjectUrl(string $bucket, string $key) : string`

The URL returned by this method is not signed nor does it ensure that the
bucket and key given to the method exist. If you need a signed URL, then
use the S3Client::createPresignedRequest method and get
the URI of the signed request.

##### Parameters

$bucket
: string

The name of the bucket where the object is located

$key
: string

The key of the object

##### Return values

string
—

The URL to the object

#### getPaginator()  [header link](class-aws-awsclientinterface.md\#method_getPaginator)

Get a result paginator for the specified operation.

`
    public
                    getPaginator(string $name[, array<string|int, mixed> $args = [] ]) : ResultPaginator`

##### Parameters

$name
: string

Name of the operation used for iterator

$args
: array<string\|int, mixed>
= \[\]

Command args to be used with each command

##### Tags  [header link](class-aws-awsclientinterface.md\#method_getPaginator\#tags)

throwsUnexpectedValueException

if the iterator config is invalid.

##### Return values

[ResultPaginator](class-aws-resultpaginator.md)

#### getRegion()  [header link](class-aws-awsclientinterface.md\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getWaiter()  [header link](class-aws-awsclientinterface.md\#method_getWaiter)

Get a waiter that waits until a resource is in a particular state.

`
    public
                    getWaiter(string|callable $name[, array<string|int, mixed> $args = [] ]) : Waiter`

Retrieving a waiter can be useful when you wish to wait asynchronously:

$waiter = $client->getWaiter('foo', \['bar' => 'baz'\]);
$waiter->promise()->then(function () { echo 'Done!'; });

##### Parameters

$name
: string\|callable

Name of the waiter that defines the wait
configuration and conditions.

$args
: array<string\|int, mixed>
= \[\]

Args to be used with each command executed
by the waiter. Waiter configuration options
can be provided in an associative array in
the @waiter key.

##### Tags  [header link](class-aws-awsclientinterface.md\#method_getWaiter\#tags)

throwsUnexpectedValueException

if the waiter is invalid.

##### Return values

[Waiter](class-aws-waiter.md)

#### registerStreamWrapper()  [header link](class-aws-s3-s3clientinterface-method-registerstreamwrapper.md)

Register the Amazon S3 stream wrapper with this client instance.

`
    public
                    registerStreamWrapper() : mixed`

#### registerStreamWrapperV2()  [header link](class-aws-s3-s3clientinterface-method-registerstreamwrapperv2.md)

Registers the Amazon S3 stream wrapper with this client instance.

`
    public
                    registerStreamWrapperV2() : mixed`

This version uses doesObjectExistV2 and doesBucketExistV2 to check
resource existence.

#### upload()  [header link](class-aws-s3-s3clientinterface-method-upload.md)

Upload a file, stream, or string to a bucket.

`
    public
                    upload(string $bucket, string $key, mixed $body[, string $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : ResultInterface`

If the upload size exceeds the specified threshold, the upload will be
performed using concurrent multipart uploads.

The options array accepts the following options:

- before\_upload: (callable) Callback to invoke before any upload
operations during the upload process. The callback should have a
function signature like `function (Aws\Command $command) {...}`.
- concurrency: (int, default=int(3)) Maximum number of concurrent
`UploadPart` operations allowed during a multipart upload.
- mup\_threshold: (int, default=int(16777216)) The size, in bytes, allowed
before the upload must be sent via a multipart upload. Default: 16 MB.
- params: (array, default=array(\[\])) Custom parameters to use with the
upload. For single uploads, they must correspond to those used for the
`PutObject` operation. For multipart uploads, they correspond to the
parameters of the `CreateMultipartUpload` operation.
- part\_size: (int) Part size to use when doing a multipart upload.

##### Parameters

$bucket
: string

Bucket to upload the object.

$key
: string

Key of the object.

$body
: mixed

Object data to upload. Can be a
StreamInterface, PHP stream resource, or a
string of data to upload.

$acl
: string
= 'private'

ACL to apply to the object (default: private).

$options
: array<string\|int, mixed>
= \[\]

Options used to configure the upload process.

##### Tags  [header link](class-aws-s3-s3clientinterface-method-upload-tags.md)

seeMultipartUploader

for more info about multipart uploads.

##### Return values

[ResultInterface](class-aws-resultinterface.md)
—

Returns the result of the upload.

#### uploadAsync()  [header link](class-aws-s3-s3clientinterface-method-uploadasync.md)

Upload a file, stream, or string to a bucket asynchronously.

`
    public
                    uploadAsync(string $bucket, string $key, mixed $body[, string $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : PromiseInterface`

##### Parameters

$bucket
: string

Bucket to upload the object.

$key
: string

Key of the object.

$body
: mixed

Object data to upload. Can be a
StreamInterface, PHP stream resource, or a
string of data to upload.

$acl
: string
= 'private'

ACL to apply to the object (default: private).

$options
: array<string\|int, mixed>
= \[\]

Options used to configure the upload process.

##### Tags  [header link](class-aws-s3-s3clientinterface-method-uploadasync-tags.md)

seeself::upload

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
—

Returns a promise that will be fulfilled
with the result of the upload.

#### uploadDirectory()  [header link](class-aws-s3-s3clientinterface-method-uploaddirectory.md)

Recursively uploads all files in a given directory to a given bucket.

`
    public
                    uploadDirectory(string $directory, string $bucket[, string $keyPrefix = null ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: string

Full path to a directory to upload

$bucket
: string

Name of the bucket

$keyPrefix
: string
= null

Virtual directory key prefix to add to each upload

$options
: array<string\|int, mixed>
= \[\]

Options available in Aws\\S3\\Transfer::\_\_construct

##### Tags  [header link](class-aws-s3-s3clientinterface-method-uploaddirectory-tags.md)

seeTransfer

for more options and customization

#### uploadDirectoryAsync()  [header link](class-aws-s3-s3clientinterface-method-uploaddirectoryasync.md)

Recursively uploads all files in a given directory to a given bucket.

`
    public
                    uploadDirectoryAsync(string $directory, string $bucket[, string $keyPrefix = null ][, array<string|int, mixed> $options = [] ]) : PromiseInterface`

##### Parameters

$directory
: string

Full path to a directory to upload

$bucket
: string

Name of the bucket

$keyPrefix
: string
= null

Virtual directory key prefix to add to each upload

$options
: array<string\|int, mixed>
= \[\]

Options available in Aws\\S3\\Transfer::\_\_construct

##### Tags  [header link](class-aws-s3-s3clientinterface-method-uploaddirectoryasync-tags.md)

seeTransfer

for more options and customization

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
—

A promise that is settled when the upload is
complete.

#### waitUntil()  [header link](class-aws-awsclientinterface.md\#method_waitUntil)

Wait until a resource is in a particular state.

`
    public
                    waitUntil(string|callable $name[, array<string|int, mixed> $args = [] ]) : void`

##### Parameters

$name
: string\|callable

Name of the waiter that defines the wait
configuration and conditions.

$args
: array<string\|int, mixed>
= \[\]

Args to be used with each command executed
by the waiter. Waiter configuration options
can be provided in an associative array in
the @waiter key.

##### Tags  [header link](class-aws-awsclientinterface.md\#method_waitUntil\#tags)

throwsUnexpectedValueException

if the waiter is invalid.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-s3-s3clientinterface-toc-constants.md)
  - [Methods](class-aws-s3-s3clientinterface-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclientinterface.md#method___call)
  - [copy()](class-aws-s3-s3clientinterface-method-copy.md)
  - [copyAsync()](class-aws-s3-s3clientinterface-method-copyasync.md)
  - [createPresignedRequest()](class-aws-s3-s3clientinterface-method-createpresignedrequest.md)
  - [deleteMatchingObjects()](class-aws-s3-s3clientinterface-method-deletematchingobjects.md)
  - [deleteMatchingObjectsAsync()](class-aws-s3-s3clientinterface-method-deletematchingobjectsasync.md)
  - [determineBucketRegion()](class-aws-s3-s3clientinterface-method-determinebucketregion.md)
  - [determineBucketRegionAsync()](class-aws-s3-s3clientinterface-method-determinebucketregionasync.md)
  - [doesBucketExist()](class-aws-s3-s3clientinterface-method-doesbucketexist.md)
  - [doesBucketExistV2()](class-aws-s3-s3clientinterface-method-doesbucketexistv2.md)
  - [doesObjectExist()](class-aws-s3-s3clientinterface-method-doesobjectexist.md)
  - [doesObjectExistV2()](class-aws-s3-s3clientinterface-method-doesobjectexistv2.md)
  - [downloadBucket()](class-aws-s3-s3clientinterface-method-downloadbucket.md)
  - [downloadBucketAsync()](class-aws-s3-s3clientinterface-method-downloadbucketasync.md)
  - [execute()](class-aws-awsclientinterface.md#method_execute)
  - [executeAsync()](class-aws-awsclientinterface.md#method_executeAsync)
  - [getApi()](class-aws-awsclientinterface.md#method_getApi)
  - [getCommand()](class-aws-awsclientinterface.md#method_getCommand)
  - [getConfig()](class-aws-awsclientinterface.md#method_getConfig)
  - [getCredentials()](class-aws-awsclientinterface.md#method_getCredentials)
  - [getEndpoint()](class-aws-awsclientinterface.md#method_getEndpoint)
  - [getHandlerList()](class-aws-awsclientinterface.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclientinterface.md#method_getIterator)
  - [getObjectUrl()](class-aws-s3-s3clientinterface-method-getobjecturl.md)
  - [getPaginator()](class-aws-awsclientinterface.md#method_getPaginator)
  - [getRegion()](class-aws-awsclientinterface.md#method_getRegion)
  - [getWaiter()](class-aws-awsclientinterface.md#method_getWaiter)
  - [registerStreamWrapper()](class-aws-s3-s3clientinterface-method-registerstreamwrapper.md)
  - [registerStreamWrapperV2()](class-aws-s3-s3clientinterface-method-registerstreamwrapperv2.md)
  - [upload()](class-aws-s3-s3clientinterface-method-upload.md)
  - [uploadAsync()](class-aws-s3-s3clientinterface-method-uploadasync.md)
  - [uploadDirectory()](class-aws-s3-s3clientinterface-method-uploaddirectory.md)
  - [uploadDirectoryAsync()](class-aws-s3-s3clientinterface-method-uploaddirectoryasync.md)
  - [waitUntil()](class-aws-awsclientinterface.md#method_waitUntil)

[Back To Top](class-aws-s3-s3clientinterface-top.md)
