Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## S3ClientTrait

A trait providing S3-specific functionality. This is meant to be used in
classes implementing \\Aws\\S3\\S3ClientInterface

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#toc-methods)

[copy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_copy)
: mixed [copyAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_copyAsync)
: mixed [deleteMatchingObjects()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_deleteMatchingObjects)
: mixed [deleteMatchingObjectsAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_deleteMatchingObjectsAsync)
: mixed [determineBucketRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_determineBucketRegion)
: mixed [determineBucketRegionAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_determineBucketRegionAsync)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[doesBucketExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesBucketExist)
: mixed [doesBucketExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesBucketExistV2)
: mixed [doesObjectExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesObjectExist)
: mixed [doesObjectExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesObjectExistV2)
: mixed [downloadBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_downloadBucket)
: mixed [downloadBucketAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_downloadBucketAsync)
: mixed [execute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_execute)
: mixed [getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_getCommand)
: mixed [getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)[getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_getIterator)
: Iterator[registerStreamWrapper()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_registerStreamWrapper)
: mixed [registerStreamWrapperV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_registerStreamWrapperV2)
: mixed [upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_upload)
: mixed [uploadAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadAsync)
: mixed [uploadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadDirectory)
: mixed [uploadDirectoryAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadDirectoryAsync)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#methods)

#### copy()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_copy)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_copy\#tags)

see[S3ClientInterface::copy()](class-aws-s3-s3clientinterface.md#method_copy)

#### copyAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_copyAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_copyAsync\#tags)

see[S3ClientInterface::copyAsync()](class-aws-s3-s3clientinterface.md#method_copyAsync)

#### deleteMatchingObjects()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_deleteMatchingObjects)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_deleteMatchingObjects\#tags)

see[S3ClientInterface::deleteMatchingObjects()](class-aws-s3-s3clientinterface.md#method_deleteMatchingObjects)

#### deleteMatchingObjectsAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_deleteMatchingObjectsAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_deleteMatchingObjectsAsync\#tags)

see[S3ClientInterface::deleteMatchingObjectsAsync()](class-aws-s3-s3clientinterface.md#method_deleteMatchingObjectsAsync)

#### determineBucketRegion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_determineBucketRegion)

`
    public
                    determineBucketRegion(mixed $bucketName) : mixed`

##### Parameters

$bucketName
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_determineBucketRegion\#tags)

see[S3ClientInterface::determineBucketRegion()](class-aws-s3-s3clientinterface.md#method_determineBucketRegion)

#### determineBucketRegionAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_determineBucketRegionAsync)

`
    public
                    determineBucketRegionAsync(string $bucketName) : PromiseInterface`

##### Parameters

$bucketName
: string

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_determineBucketRegionAsync\#tags)

see[S3ClientInterface::determineBucketRegionAsync()](class-aws-s3-s3clientinterface.md#method_determineBucketRegionAsync)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### doesBucketExist()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesBucketExist)

`
    public
                    doesBucketExist(mixed $bucket) : mixed`

##### Parameters

$bucket
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesBucketExist\#tags)

see[S3ClientInterface::doesBucketExist()](class-aws-s3-s3clientinterface.md#method_doesBucketExist)

#### doesBucketExistV2()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesBucketExistV2)

`
    public
                    doesBucketExistV2(mixed $bucket[, mixed $accept403 = false ]) : mixed`

##### Parameters

$bucket
: mixed$accept403
: mixed
= false

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesBucketExistV2\#tags)

see[S3ClientInterface::doesBucketExistV2()](class-aws-s3-s3clientinterface.md#method_doesBucketExistV2)

#### doesObjectExist()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesObjectExist)

`
    public
                    doesObjectExist(mixed $bucket, mixed $key[, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesObjectExist\#tags)

see[S3ClientInterface::doesObjectExist()](class-aws-s3-s3clientinterface.md#method_doesObjectExist)

#### doesObjectExistV2()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesObjectExistV2)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesObjectExistV2\#tags)

see[S3ClientInterface::doesObjectExistV2()](class-aws-s3-s3clientinterface.md#method_doesObjectExistV2)

#### downloadBucket()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_downloadBucket)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_downloadBucket\#tags)

see[S3ClientInterface::downloadBucket()](class-aws-s3-s3clientinterface.md#method_downloadBucket)

#### downloadBucketAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_downloadBucketAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_downloadBucketAsync\#tags)

see[S3ClientInterface::downloadBucketAsync()](class-aws-s3-s3clientinterface.md#method_downloadBucketAsync)

#### execute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_execute)

`
    public
    abstract                execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_execute\#tags)

seeS3ClientInterface::execute()

#### getCommand()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_getCommand)

`
    public
    abstract                getCommand(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_getCommand\#tags)

seeS3ClientInterface::getCommand()

#### getHandlerList()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_getHandlerList)

`
    public
    abstract                getHandlerList() : HandlerList`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_getHandlerList\#tags)

seeS3ClientInterface::getHandlerList()

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_getIterator)

`
    public
    abstract                getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : Iterator`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_getIterator\#tags)

seeS3ClientInterface::getIterator()

##### Return values

Iterator

#### registerStreamWrapper()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_registerStreamWrapper)

`
    public
                    registerStreamWrapper() : mixed`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_registerStreamWrapper\#tags)

see[S3ClientInterface::registerStreamWrapper()](class-aws-s3-s3clientinterface.md#method_registerStreamWrapper)

#### registerStreamWrapperV2()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_registerStreamWrapperV2)

`
    public
                    registerStreamWrapperV2() : mixed`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_registerStreamWrapperV2\#tags)

see[S3ClientInterface::registerStreamWrapperV2()](class-aws-s3-s3clientinterface.md#method_registerStreamWrapperV2)

#### upload()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_upload)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_upload\#tags)

see[S3ClientInterface::upload()](class-aws-s3-s3clientinterface.md#method_upload)

#### uploadAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadAsync\#tags)

see[S3ClientInterface::uploadAsync()](class-aws-s3-s3clientinterface.md#method_uploadAsync)

#### uploadDirectory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadDirectory)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadDirectory\#tags)

see[S3ClientInterface::uploadDirectory()](class-aws-s3-s3clientinterface.md#method_uploadDirectory)

#### uploadDirectoryAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadDirectoryAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadDirectoryAsync\#tags)

see[S3ClientInterface::uploadDirectoryAsync()](class-aws-s3-s3clientinterface.md#method_uploadDirectoryAsync)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#toc-methods)
- Methods
  - [copy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_copy)
  - [copyAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_copyAsync)
  - [deleteMatchingObjects()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_deleteMatchingObjects)
  - [deleteMatchingObjectsAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_deleteMatchingObjectsAsync)
  - [determineBucketRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_determineBucketRegion)
  - [determineBucketRegionAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_determineBucketRegionAsync)
  - [doesBucketExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesBucketExist)
  - [doesBucketExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesBucketExistV2)
  - [doesObjectExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesObjectExist)
  - [doesObjectExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesObjectExistV2)
  - [downloadBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_downloadBucket)
  - [downloadBucketAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_downloadBucketAsync)
  - [execute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_execute)
  - [getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_getCommand)
  - [getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_getHandlerList)
  - [getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_getIterator)
  - [registerStreamWrapper()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_registerStreamWrapper)
  - [registerStreamWrapperV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_registerStreamWrapperV2)
  - [upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_upload)
  - [uploadAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadAsync)
  - [uploadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadDirectory)
  - [uploadDirectoryAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadDirectoryAsync)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#top)
