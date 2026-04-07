Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## BatchDelete        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Efficiently deletes many objects from a single Amazon S3 bucket using an
iterator that yields keys. Deletes are made using the DeleteObjects API
operation.

$s3 = new Aws\\S3\\Client(\[
'region' => 'us-west-2',
'version' => 'latest'
\]);

```prettyprint
$listObjectsParams = ['Bucket' => 'foo', 'Prefix' => 'starts/with/'];
$delete = Aws\S3\BatchDelete::fromListObjects($s3, $listObjectsParams);
// Asynchronously delete
$promise = $delete->promise();
// Force synchronous completion
$delete->delete();

```

When using one of the batch delete creational static methods, you can supply
an associative array of options:

- before: Function invoked before executing a command. The function is
passed the command that is about to be executed. This can be useful
for logging, adding custom request headers, etc.
- batch\_size: The size of each delete batch. Defaults to 1000.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#tags)

link[http://docs.aws.amazon.com/AmazonS3/latest/API/multiobjectdeleteapi.html](https://docs.aws.amazon.com/AmazonS3/latest/API/multiobjectdeleteapi.html)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#toc-interfaces)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#toc-methods)

[delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#method_delete)
: mixed Synchronously deletes all of the objects.[fromIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#method_fromIterator)
: [BatchDelete](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html)Creates a BatchDelete object from an iterator that yields results.[fromListObjects()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#method_fromListObjects)
: [BatchDelete](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html)Creates a BatchDelete object from all of the paginated results of a
ListObjects operation. Each result that is returned by the ListObjects
operation will be deleted.[promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#methods)

#### delete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#method_delete)

Synchronously deletes all of the objects.

`
    public
                    delete() : mixed`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#method_delete\#tags)

throws[DeleteMultipleObjectsException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Exception.DeleteMultipleObjectsException.html)

on error.

#### fromIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#method_fromIterator)

Creates a BatchDelete object from an iterator that yields results.

`
    public
            static        fromIterator(AwsClientInterface $client, string $bucket, Iterator $iter[, array<string|int, mixed> $options = [] ]) : BatchDelete`

##### Parameters

$client
: [AwsClientInterface](class-aws-awsclientinterface.md)

AWS Client to use to execute commands

$bucket
: string

Bucket where the objects are stored

$iter
: Iterator

Iterator that yields assoc arrays

$options
: array<string\|int, mixed>
= \[\]

BatchDelete options

##### Return values

[BatchDelete](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html)

#### fromListObjects()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#method_fromListObjects)

Creates a BatchDelete object from all of the paginated results of a
ListObjects operation. Each result that is returned by the ListObjects
operation will be deleted.

`
    public
            static        fromListObjects(AwsClientInterface $client, array<string|int, mixed> $listObjectsParams[, array<string|int, mixed> $options = [] ]) : BatchDelete`

##### Parameters

$client
: [AwsClientInterface](class-aws-awsclientinterface.md)

AWS Client to use.

$listObjectsParams
: array<string\|int, mixed>

ListObjects API parameters

$options
: array<string\|int, mixed>
= \[\]

BatchDelete options.

##### Return values

[BatchDelete](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html)

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html\#method_promise)

Returns a promise.

`
    public
                    promise() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#toc-methods)
- Methods
  - [delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#method_delete)
  - [fromIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#method_fromIterator)
  - [fromListObjects()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#method_fromListObjects)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#method_promise)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html#top)
