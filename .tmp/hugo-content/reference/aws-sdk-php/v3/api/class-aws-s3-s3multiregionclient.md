Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## S3MultiRegionClient     extends [MultiRegionClient](class-aws-multiregionclient.md)   in package    - [Aws](package-aws.md)       implements  [S3ClientInterface](class-aws-s3-s3clientinterface.md)  Uses  [S3ClientTrait](class-aws-s3-s3clienttrait.md)

**Amazon Simple Storage Service** multi-region client.

### Table of Contents  [header link](class-aws-s3-s3multiregionclient-toc.md)

#### Interfaces  [header link](class-aws-s3-s3multiregionclient-toc-interfaces.md)

[S3ClientInterface](class-aws-s3-s3clientinterface.md)\*\*Amazon Simple Storage Service\*\* client.

#### Methods  [header link](class-aws-s3-s3multiregionclient-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-s3-s3multiregionclient-method-construct.md)
: mixed The multi-region client constructor accepts the following options:[copy()](class-aws-s3-s3clienttrait.md#method_copy)
: mixed [copyAsync()](class-aws-s3-s3clienttrait.md#method_copyAsync)
: mixed [createPresignedRequest()](class-aws-s3-s3multiregionclient-method-createpresignedrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Create a pre-signed URL for the given S3 command object.[deleteMatchingObjects()](class-aws-s3-s3clienttrait.md#method_deleteMatchingObjects)
: mixed [deleteMatchingObjectsAsync()](class-aws-s3-s3clienttrait.md#method_deleteMatchingObjectsAsync)
: mixed [determineBucketRegion()](class-aws-s3-s3clienttrait.md#method_determineBucketRegion)
: mixed [determineBucketRegionAsync()](class-aws-s3-s3multiregionclient-method-determinebucketregionasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise fulfilled with the region in which a given bucket is
located.[doesBucketExist()](class-aws-s3-s3clienttrait.md#method_doesBucketExist)
: mixed [doesBucketExistV2()](class-aws-s3-s3clienttrait.md#method_doesBucketExistV2)
: mixed [doesObjectExist()](class-aws-s3-s3clienttrait.md#method_doesObjectExist)
: mixed [doesObjectExistV2()](class-aws-s3-s3clienttrait.md#method_doesObjectExistV2)
: mixed [downloadBucket()](class-aws-s3-s3clienttrait.md#method_downloadBucket)
: mixed [downloadBucketAsync()](class-aws-s3-s3clienttrait.md#method_downloadBucketAsync)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-s3-s3multiregionclient-method-getarguments.md)
: mixed [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-multiregionclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-multiregionclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-multiregionclient.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getHandlerList()](class-aws-multiregionclient.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getObjectUrl()](class-aws-s3-s3multiregionclient-method-getobjecturl.md)
: string Returns the URL to an object identified by its bucket and key.[getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-multiregionclient.md#method_getRegion)
: string Get the region to which the client is configured to send requests by
default.[getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [registerStreamWrapper()](class-aws-s3-s3clienttrait.md#method_registerStreamWrapper)
: mixed [registerStreamWrapperV2()](class-aws-s3-s3clienttrait.md#method_registerStreamWrapperV2)
: mixed [upload()](class-aws-s3-s3clienttrait.md#method_upload)
: mixed [uploadAsync()](class-aws-s3-s3clienttrait.md#method_uploadAsync)
: mixed [uploadDirectory()](class-aws-s3-s3clienttrait.md#method_uploadDirectory)
: mixed [uploadDirectoryAsync()](class-aws-s3-s3clienttrait.md#method_uploadDirectoryAsync)
: mixed [useCustomHandler()](class-aws-multiregionclient.md#method_useCustomHandler)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](class-aws-s3-s3multiregionclient-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-s3-s3multiregionclient-method-construct.md)

The multi-region client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

- client\_factory: (callable) An optional callable that takes an array of
client configuration arguments and returns a regionalized client.
- partition: (Aws\\Endpoint\\Partition\|string) AWS partition to connect to.
Valid partitions include "aws," "aws-cn," and "aws-us-gov." Used to
restrict the scope of the mapRegions method.
- region: (string) Region to connect to when no override is provided.
Used to create the default client factory and determine the appropriate
AWS partition when present.

##### Parameters

$args
: array<string\|int, mixed>

Client configuration arguments.

#### copy()  [header link](class-aws-s3-s3clienttrait.md\#method_copy)

`
    public
                    copy(mixed $fromB, mixed $fromK, mixed $destB, mixed $destK[, mixed $acl = 'private' ][, array<string|int, mixed> $opts = [] ]) : mixed`

##### Parameters

$fromB
: mixed$fromK
: mixed$destB
: mixed$destK
: mixed$acl
: mixed
= 'private'$opts
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_copy\#tags)

see[S3ClientInterface::copy()](class-aws-s3-s3clientinterface.md#method_copy)

#### copyAsync()  [header link](class-aws-s3-s3clienttrait.md\#method_copyAsync)

`
    public
                    copyAsync(mixed $fromB, mixed $fromK, mixed $destB, mixed $destK[, mixed $acl = 'private' ][, array<string|int, mixed> $opts = [] ]) : mixed`

##### Parameters

$fromB
: mixed$fromK
: mixed$destB
: mixed$destK
: mixed$acl
: mixed
= 'private'$opts
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_copyAsync\#tags)

see[S3ClientInterface::copyAsync()](class-aws-s3-s3clientinterface.md#method_copyAsync)

#### createPresignedRequest()  [header link](class-aws-s3-s3multiregionclient-method-createpresignedrequest.md)

Create a pre-signed URL for the given S3 command object.

`
    public
                    createPresignedRequest(CommandInterface $command, mixed $expires[, array<string|int, mixed> $options = [] ]) : RequestInterface`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

Command to create a pre-signed
URL for.

$expires
: mixed

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

#### deleteMatchingObjects()  [header link](class-aws-s3-s3clienttrait.md\#method_deleteMatchingObjects)

`
    public
                    deleteMatchingObjects(mixed $bucket[, mixed $prefix = '' ][, mixed $regex = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$prefix
: mixed
= ''$regex
: mixed
= ''$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_deleteMatchingObjects\#tags)

see[S3ClientInterface::deleteMatchingObjects()](class-aws-s3-s3clientinterface.md#method_deleteMatchingObjects)

#### deleteMatchingObjectsAsync()  [header link](class-aws-s3-s3clienttrait.md\#method_deleteMatchingObjectsAsync)

`
    public
                    deleteMatchingObjectsAsync(mixed $bucket[, mixed $prefix = '' ][, mixed $regex = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$prefix
: mixed
= ''$regex
: mixed
= ''$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_deleteMatchingObjectsAsync\#tags)

see[S3ClientInterface::deleteMatchingObjectsAsync()](class-aws-s3-s3clientinterface.md#method_deleteMatchingObjectsAsync)

#### determineBucketRegion()  [header link](class-aws-s3-s3clienttrait.md\#method_determineBucketRegion)

`
    public
                    determineBucketRegion(mixed $bucketName) : mixed`

##### Parameters

$bucketName
: mixed

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_determineBucketRegion\#tags)

see[S3ClientInterface::determineBucketRegion()](class-aws-s3-s3clientinterface.md#method_determineBucketRegion)

#### determineBucketRegionAsync()  [header link](class-aws-s3-s3multiregionclient-method-determinebucketregionasync.md)

Returns a promise fulfilled with the region in which a given bucket is
located.

`
    public
                    determineBucketRegionAsync(mixed $bucketName) : PromiseInterface`

##### Parameters

$bucketName
: mixed

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### doesBucketExist()  [header link](class-aws-s3-s3clienttrait.md\#method_doesBucketExist)

`
    public
                    doesBucketExist(mixed $bucket) : mixed`

##### Parameters

$bucket
: mixed

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_doesBucketExist\#tags)

see[S3ClientInterface::doesBucketExist()](class-aws-s3-s3clientinterface.md#method_doesBucketExist)

#### doesBucketExistV2()  [header link](class-aws-s3-s3clienttrait.md\#method_doesBucketExistV2)

`
    public
                    doesBucketExistV2(mixed $bucket[, mixed $accept403 = false ]) : mixed`

##### Parameters

$bucket
: mixed$accept403
: mixed
= false

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_doesBucketExistV2\#tags)

see[S3ClientInterface::doesBucketExistV2()](class-aws-s3-s3clientinterface.md#method_doesBucketExistV2)

#### doesObjectExist()  [header link](class-aws-s3-s3clienttrait.md\#method_doesObjectExist)

`
    public
                    doesObjectExist(mixed $bucket, mixed $key[, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_doesObjectExist\#tags)

see[S3ClientInterface::doesObjectExist()](class-aws-s3-s3clientinterface.md#method_doesObjectExist)

#### doesObjectExistV2()  [header link](class-aws-s3-s3clienttrait.md\#method_doesObjectExistV2)

`
    public
                    doesObjectExistV2(mixed $bucket, mixed $key[, mixed $includeDeleteMarkers = false ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$includeDeleteMarkers
: mixed
= false$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_doesObjectExistV2\#tags)

see[S3ClientInterface::doesObjectExistV2()](class-aws-s3-s3clientinterface.md#method_doesObjectExistV2)

#### downloadBucket()  [header link](class-aws-s3-s3clienttrait.md\#method_downloadBucket)

`
    public
                    downloadBucket(mixed $directory, mixed $bucket[, mixed $keyPrefix = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: mixed$bucket
: mixed$keyPrefix
: mixed
= ''$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_downloadBucket\#tags)

see[S3ClientInterface::downloadBucket()](class-aws-s3-s3clientinterface.md#method_downloadBucket)

#### downloadBucketAsync()  [header link](class-aws-s3-s3clienttrait.md\#method_downloadBucketAsync)

`
    public
                    downloadBucketAsync(mixed $directory, mixed $bucket[, mixed $keyPrefix = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: mixed$bucket
: mixed$keyPrefix
: mixed
= ''$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_downloadBucketAsync\#tags)

see[S3ClientInterface::downloadBucketAsync()](class-aws-s3-s3clientinterface.md#method_downloadBucketAsync)

#### execute()  [header link](class-aws-awsclienttrait.md\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait.md\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### getApi()  [header link](class-aws-awsclienttrait.md\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-s3-s3multiregionclient-method-getarguments.md)

`
    public
            static        getArguments() : mixed`

#### getCommand()  [header link](class-aws-awsclienttrait.md\#method_getCommand)

`
    public
    abstract                getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

##### Parameters

$name
: string$args
: array<string\|int, mixed>
= \[\]

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-multiregionclient.md\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([mixed $option = null ]) : mixed|null`

##### Parameters

$option
: mixed
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](class-aws-multiregionclient.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-multiregionclient.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getHandlerList()  [header link](class-aws-multiregionclient.md\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait.md\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getObjectUrl()  [header link](class-aws-s3-s3multiregionclient-method-getobjecturl.md)

Returns the URL to an object identified by its bucket and key.

`
    public
                    getObjectUrl(mixed $bucket, mixed $key) : string`

The URL returned by this method is not signed nor does it ensure that the
bucket and key given to the method exist. If you need a signed URL, then
use the S3Client::createPresignedRequest method and get
the URI of the signed request.

##### Parameters

$bucket
: mixed

The name of the bucket where the object is located

$key
: mixed

The key of the object

##### Return values

string
—

The URL to the object

#### getPaginator()  [header link](class-aws-awsclienttrait.md\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](class-aws-multiregionclient.md\#method_getRegion)

Get the region to which the client is configured to send requests by
default.

`
    public
                    getRegion() : string`

##### Return values

string

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### registerStreamWrapper()  [header link](class-aws-s3-s3clienttrait.md\#method_registerStreamWrapper)

`
    public
                    registerStreamWrapper() : mixed`

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_registerStreamWrapper\#tags)

see[S3ClientInterface::registerStreamWrapper()](class-aws-s3-s3clientinterface.md#method_registerStreamWrapper)

#### registerStreamWrapperV2()  [header link](class-aws-s3-s3clienttrait.md\#method_registerStreamWrapperV2)

`
    public
                    registerStreamWrapperV2() : mixed`

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_registerStreamWrapperV2\#tags)

see[S3ClientInterface::registerStreamWrapperV2()](class-aws-s3-s3clientinterface.md#method_registerStreamWrapperV2)

#### upload()  [header link](class-aws-s3-s3clienttrait.md\#method_upload)

`
    public
                    upload(mixed $bucket, mixed $key, mixed $body[, mixed $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$body
: mixed$acl
: mixed
= 'private'$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_upload\#tags)

see[S3ClientInterface::upload()](class-aws-s3-s3clientinterface.md#method_upload)

#### uploadAsync()  [header link](class-aws-s3-s3clienttrait.md\#method_uploadAsync)

`
    public
                    uploadAsync(mixed $bucket, mixed $key, mixed $body[, mixed $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$body
: mixed$acl
: mixed
= 'private'$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_uploadAsync\#tags)

see[S3ClientInterface::uploadAsync()](class-aws-s3-s3clientinterface.md#method_uploadAsync)

#### uploadDirectory()  [header link](class-aws-s3-s3clienttrait.md\#method_uploadDirectory)

`
    public
                    uploadDirectory(mixed $directory, mixed $bucket[, mixed $keyPrefix = null ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: mixed$bucket
: mixed$keyPrefix
: mixed
= null$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_uploadDirectory\#tags)

see[S3ClientInterface::uploadDirectory()](class-aws-s3-s3clientinterface.md#method_uploadDirectory)

#### uploadDirectoryAsync()  [header link](class-aws-s3-s3clienttrait.md\#method_uploadDirectoryAsync)

`
    public
                    uploadDirectoryAsync(mixed $directory, mixed $bucket[, mixed $keyPrefix = null ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: mixed$bucket
: mixed$keyPrefix
: mixed
= null$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait.md\#method_uploadDirectoryAsync\#tags)

see[S3ClientInterface::uploadDirectoryAsync()](class-aws-s3-s3clientinterface.md#method_uploadDirectoryAsync)

#### useCustomHandler()  [header link](class-aws-multiregionclient.md\#method_useCustomHandler)

`
    public
                    useCustomHandler(callable $handler) : mixed`

##### Parameters

$handler
: callable

#### waitUntil()  [header link](class-aws-awsclienttrait.md\#method_waitUntil)

`
    public
                    waitUntil(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3multiregionclient-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-s3-s3multiregionclient-method-construct.md)
  - [copy()](class-aws-s3-s3clienttrait.md#method_copy)
  - [copyAsync()](class-aws-s3-s3clienttrait.md#method_copyAsync)
  - [createPresignedRequest()](class-aws-s3-s3multiregionclient-method-createpresignedrequest.md)
  - [deleteMatchingObjects()](class-aws-s3-s3clienttrait.md#method_deleteMatchingObjects)
  - [deleteMatchingObjectsAsync()](class-aws-s3-s3clienttrait.md#method_deleteMatchingObjectsAsync)
  - [determineBucketRegion()](class-aws-s3-s3clienttrait.md#method_determineBucketRegion)
  - [determineBucketRegionAsync()](class-aws-s3-s3multiregionclient-method-determinebucketregionasync.md)
  - [doesBucketExist()](class-aws-s3-s3clienttrait.md#method_doesBucketExist)
  - [doesBucketExistV2()](class-aws-s3-s3clienttrait.md#method_doesBucketExistV2)
  - [doesObjectExist()](class-aws-s3-s3clienttrait.md#method_doesObjectExist)
  - [doesObjectExistV2()](class-aws-s3-s3clienttrait.md#method_doesObjectExistV2)
  - [downloadBucket()](class-aws-s3-s3clienttrait.md#method_downloadBucket)
  - [downloadBucketAsync()](class-aws-s3-s3clienttrait.md#method_downloadBucketAsync)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-s3-s3multiregionclient-method-getarguments.md)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-multiregionclient.md#method_getConfig)
  - [getCredentials()](class-aws-multiregionclient.md#method_getCredentials)
  - [getEndpoint()](class-aws-multiregionclient.md#method_getEndpoint)
  - [getHandlerList()](class-aws-multiregionclient.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getObjectUrl()](class-aws-s3-s3multiregionclient-method-getobjecturl.md)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-multiregionclient.md#method_getRegion)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [registerStreamWrapper()](class-aws-s3-s3clienttrait.md#method_registerStreamWrapper)
  - [registerStreamWrapperV2()](class-aws-s3-s3clienttrait.md#method_registerStreamWrapperV2)
  - [upload()](class-aws-s3-s3clienttrait.md#method_upload)
  - [uploadAsync()](class-aws-s3-s3clienttrait.md#method_uploadAsync)
  - [uploadDirectory()](class-aws-s3-s3clienttrait.md#method_uploadDirectory)
  - [uploadDirectoryAsync()](class-aws-s3-s3clienttrait.md#method_uploadDirectoryAsync)
  - [useCustomHandler()](class-aws-multiregionclient.md#method_useCustomHandler)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](class-aws-s3-s3multiregionclient-top.md)
