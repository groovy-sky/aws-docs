Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## S3ClientTrait

A trait providing S3-specific functionality. This is meant to be used in
classes implementing \\Aws\\S3\\S3ClientInterface

### Table of Contents  [header link](class-aws-s3-s3clienttrait-toc.md)

#### Methods  [header link](class-aws-s3-s3clienttrait-toc-methods.md)

[copy()](class-aws-s3-s3clienttrait-method-copy.md)
: mixed [copyAsync()](class-aws-s3-s3clienttrait-method-copyasync.md)
: mixed [deleteMatchingObjects()](class-aws-s3-s3clienttrait-method-deletematchingobjects.md)
: mixed [deleteMatchingObjectsAsync()](class-aws-s3-s3clienttrait-method-deletematchingobjectsasync.md)
: mixed [determineBucketRegion()](class-aws-s3-s3clienttrait-method-determinebucketregion.md)
: mixed [determineBucketRegionAsync()](class-aws-s3-s3clienttrait-method-determinebucketregionasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[doesBucketExist()](class-aws-s3-s3clienttrait-method-doesbucketexist.md)
: mixed [doesBucketExistV2()](class-aws-s3-s3clienttrait-method-doesbucketexistv2.md)
: mixed [doesObjectExist()](class-aws-s3-s3clienttrait-method-doesobjectexist.md)
: mixed [doesObjectExistV2()](class-aws-s3-s3clienttrait-method-doesobjectexistv2.md)
: mixed [downloadBucket()](class-aws-s3-s3clienttrait-method-downloadbucket.md)
: mixed [downloadBucketAsync()](class-aws-s3-s3clienttrait-method-downloadbucketasync.md)
: mixed [execute()](class-aws-s3-s3clienttrait-method-execute.md)
: mixed [getCommand()](class-aws-s3-s3clienttrait-method-getcommand.md)
: mixed [getHandlerList()](class-aws-s3-s3clienttrait-method-gethandlerlist.md)
: [HandlerList](class-aws-handlerlist.md)[getIterator()](class-aws-s3-s3clienttrait-method-getiterator.md)
: Iterator[registerStreamWrapper()](class-aws-s3-s3clienttrait-method-registerstreamwrapper.md)
: mixed [registerStreamWrapperV2()](class-aws-s3-s3clienttrait-method-registerstreamwrapperv2.md)
: mixed [upload()](class-aws-s3-s3clienttrait-method-upload.md)
: mixed [uploadAsync()](class-aws-s3-s3clienttrait-method-uploadasync.md)
: mixed [uploadDirectory()](class-aws-s3-s3clienttrait-method-uploaddirectory.md)
: mixed [uploadDirectoryAsync()](class-aws-s3-s3clienttrait-method-uploaddirectoryasync.md)
: mixed

### Methods  [header link](class-aws-s3-s3clienttrait-methods.md)

#### copy()  [header link](class-aws-s3-s3clienttrait-method-copy.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-copy-tags.md)

see[S3ClientInterface::copy()](class-aws-s3-s3clientinterface.md#method_copy)

#### copyAsync()  [header link](class-aws-s3-s3clienttrait-method-copyasync.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-copyasync-tags.md)

see[S3ClientInterface::copyAsync()](class-aws-s3-s3clientinterface.md#method_copyAsync)

#### deleteMatchingObjects()  [header link](class-aws-s3-s3clienttrait-method-deletematchingobjects.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-deletematchingobjects-tags.md)

see[S3ClientInterface::deleteMatchingObjects()](class-aws-s3-s3clientinterface.md#method_deleteMatchingObjects)

#### deleteMatchingObjectsAsync()  [header link](class-aws-s3-s3clienttrait-method-deletematchingobjectsasync.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-deletematchingobjectsasync-tags.md)

see[S3ClientInterface::deleteMatchingObjectsAsync()](class-aws-s3-s3clientinterface.md#method_deleteMatchingObjectsAsync)

#### determineBucketRegion()  [header link](class-aws-s3-s3clienttrait-method-determinebucketregion.md)

`
    public
                    determineBucketRegion(mixed $bucketName) : mixed`

##### Parameters

$bucketName
: mixed

##### Tags  [header link](class-aws-s3-s3clienttrait-method-determinebucketregion-tags.md)

see[S3ClientInterface::determineBucketRegion()](class-aws-s3-s3clientinterface.md#method_determineBucketRegion)

#### determineBucketRegionAsync()  [header link](class-aws-s3-s3clienttrait-method-determinebucketregionasync.md)

`
    public
                    determineBucketRegionAsync(string $bucketName) : PromiseInterface`

##### Parameters

$bucketName
: string

##### Tags  [header link](class-aws-s3-s3clienttrait-method-determinebucketregionasync-tags.md)

see[S3ClientInterface::determineBucketRegionAsync()](class-aws-s3-s3clientinterface.md#method_determineBucketRegionAsync)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### doesBucketExist()  [header link](class-aws-s3-s3clienttrait-method-doesbucketexist.md)

`
    public
                    doesBucketExist(mixed $bucket) : mixed`

##### Parameters

$bucket
: mixed

##### Tags  [header link](class-aws-s3-s3clienttrait-method-doesbucketexist-tags.md)

see[S3ClientInterface::doesBucketExist()](class-aws-s3-s3clientinterface.md#method_doesBucketExist)

#### doesBucketExistV2()  [header link](class-aws-s3-s3clienttrait-method-doesbucketexistv2.md)

`
    public
                    doesBucketExistV2(mixed $bucket[, mixed $accept403 = false ]) : mixed`

##### Parameters

$bucket
: mixed$accept403
: mixed
= false

##### Tags  [header link](class-aws-s3-s3clienttrait-method-doesbucketexistv2-tags.md)

see[S3ClientInterface::doesBucketExistV2()](class-aws-s3-s3clientinterface.md#method_doesBucketExistV2)

#### doesObjectExist()  [header link](class-aws-s3-s3clienttrait-method-doesobjectexist.md)

`
    public
                    doesObjectExist(mixed $bucket, mixed $key[, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-doesobjectexist-tags.md)

see[S3ClientInterface::doesObjectExist()](class-aws-s3-s3clientinterface.md#method_doesObjectExist)

#### doesObjectExistV2()  [header link](class-aws-s3-s3clienttrait-method-doesobjectexistv2.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-doesobjectexistv2-tags.md)

see[S3ClientInterface::doesObjectExistV2()](class-aws-s3-s3clientinterface.md#method_doesObjectExistV2)

#### downloadBucket()  [header link](class-aws-s3-s3clienttrait-method-downloadbucket.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-downloadbucket-tags.md)

see[S3ClientInterface::downloadBucket()](class-aws-s3-s3clientinterface.md#method_downloadBucket)

#### downloadBucketAsync()  [header link](class-aws-s3-s3clienttrait-method-downloadbucketasync.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-downloadbucketasync-tags.md)

see[S3ClientInterface::downloadBucketAsync()](class-aws-s3-s3clientinterface.md#method_downloadBucketAsync)

#### execute()  [header link](class-aws-s3-s3clienttrait-method-execute.md)

`
    public
    abstract                execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

##### Tags  [header link](class-aws-s3-s3clienttrait-method-execute-tags.md)

seeS3ClientInterface::execute()

#### getCommand()  [header link](class-aws-s3-s3clienttrait-method-getcommand.md)

`
    public
    abstract                getCommand(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-getcommand-tags.md)

seeS3ClientInterface::getCommand()

#### getHandlerList()  [header link](class-aws-s3-s3clienttrait-method-gethandlerlist.md)

`
    public
    abstract                getHandlerList() : HandlerList`

##### Tags  [header link](class-aws-s3-s3clienttrait-method-gethandlerlist-tags.md)

seeS3ClientInterface::getHandlerList()

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-s3-s3clienttrait-method-getiterator.md)

`
    public
    abstract                getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : Iterator`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-getiterator-tags.md)

seeS3ClientInterface::getIterator()

##### Return values

Iterator

#### registerStreamWrapper()  [header link](class-aws-s3-s3clienttrait-method-registerstreamwrapper.md)

`
    public
                    registerStreamWrapper() : mixed`

##### Tags  [header link](class-aws-s3-s3clienttrait-method-registerstreamwrapper-tags.md)

see[S3ClientInterface::registerStreamWrapper()](class-aws-s3-s3clientinterface.md#method_registerStreamWrapper)

#### registerStreamWrapperV2()  [header link](class-aws-s3-s3clienttrait-method-registerstreamwrapperv2.md)

`
    public
                    registerStreamWrapperV2() : mixed`

##### Tags  [header link](class-aws-s3-s3clienttrait-method-registerstreamwrapperv2-tags.md)

see[S3ClientInterface::registerStreamWrapperV2()](class-aws-s3-s3clientinterface.md#method_registerStreamWrapperV2)

#### upload()  [header link](class-aws-s3-s3clienttrait-method-upload.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-upload-tags.md)

see[S3ClientInterface::upload()](class-aws-s3-s3clientinterface.md#method_upload)

#### uploadAsync()  [header link](class-aws-s3-s3clienttrait-method-uploadasync.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-uploadasync-tags.md)

see[S3ClientInterface::uploadAsync()](class-aws-s3-s3clientinterface.md#method_uploadAsync)

#### uploadDirectory()  [header link](class-aws-s3-s3clienttrait-method-uploaddirectory.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-uploaddirectory-tags.md)

see[S3ClientInterface::uploadDirectory()](class-aws-s3-s3clientinterface.md#method_uploadDirectory)

#### uploadDirectoryAsync()  [header link](class-aws-s3-s3clienttrait-method-uploaddirectoryasync.md)

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

##### Tags  [header link](class-aws-s3-s3clienttrait-method-uploaddirectoryasync-tags.md)

see[S3ClientInterface::uploadDirectoryAsync()](class-aws-s3-s3clientinterface.md#method_uploadDirectoryAsync)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3clienttrait-toc-methods.md)
- Methods
  - [copy()](class-aws-s3-s3clienttrait-method-copy.md)
  - [copyAsync()](class-aws-s3-s3clienttrait-method-copyasync.md)
  - [deleteMatchingObjects()](class-aws-s3-s3clienttrait-method-deletematchingobjects.md)
  - [deleteMatchingObjectsAsync()](class-aws-s3-s3clienttrait-method-deletematchingobjectsasync.md)
  - [determineBucketRegion()](class-aws-s3-s3clienttrait-method-determinebucketregion.md)
  - [determineBucketRegionAsync()](class-aws-s3-s3clienttrait-method-determinebucketregionasync.md)
  - [doesBucketExist()](class-aws-s3-s3clienttrait-method-doesbucketexist.md)
  - [doesBucketExistV2()](class-aws-s3-s3clienttrait-method-doesbucketexistv2.md)
  - [doesObjectExist()](class-aws-s3-s3clienttrait-method-doesobjectexist.md)
  - [doesObjectExistV2()](class-aws-s3-s3clienttrait-method-doesobjectexistv2.md)
  - [downloadBucket()](class-aws-s3-s3clienttrait-method-downloadbucket.md)
  - [downloadBucketAsync()](class-aws-s3-s3clienttrait-method-downloadbucketasync.md)
  - [execute()](class-aws-s3-s3clienttrait-method-execute.md)
  - [getCommand()](class-aws-s3-s3clienttrait-method-getcommand.md)
  - [getHandlerList()](class-aws-s3-s3clienttrait-method-gethandlerlist.md)
  - [getIterator()](class-aws-s3-s3clienttrait-method-getiterator.md)
  - [registerStreamWrapper()](class-aws-s3-s3clienttrait-method-registerstreamwrapper.md)
  - [registerStreamWrapperV2()](class-aws-s3-s3clienttrait-method-registerstreamwrapperv2.md)
  - [upload()](class-aws-s3-s3clienttrait-method-upload.md)
  - [uploadAsync()](class-aws-s3-s3clienttrait-method-uploadasync.md)
  - [uploadDirectory()](class-aws-s3-s3clienttrait-method-uploaddirectory.md)
  - [uploadDirectoryAsync()](class-aws-s3-s3clienttrait-method-uploaddirectoryasync.md)

[Back To Top](class-aws-s3-s3clienttrait-top.md)
