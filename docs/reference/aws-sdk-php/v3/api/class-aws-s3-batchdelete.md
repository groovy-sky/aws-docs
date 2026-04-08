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

##### Tags  [header link](class-aws-s3-batchdelete-tags.md)

link[http://docs.aws.amazon.com/AmazonS3/latest/API/multiobjectdeleteapi.html](../../../../services/s3/latest/api/multiobjectdeleteapi.md)

### Table of Contents  [header link](class-aws-s3-batchdelete-toc.md)

#### Interfaces  [header link](class-aws-s3-batchdelete-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](class-aws-s3-batchdelete-toc-methods.md)

[delete()](class-aws-s3-batchdelete-method-delete.md)
: mixed Synchronously deletes all of the objects.[fromIterator()](class-aws-s3-batchdelete-method-fromiterator.md)
: [BatchDelete](class-aws-s3-batchdelete.md)Creates a BatchDelete object from an iterator that yields results.[fromListObjects()](class-aws-s3-batchdelete-method-fromlistobjects.md)
: [BatchDelete](class-aws-s3-batchdelete.md)Creates a BatchDelete object from all of the paginated results of a
ListObjects operation. Each result that is returned by the ListObjects
operation will be deleted.[promise()](class-aws-s3-batchdelete-method-promise.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise.

### Methods  [header link](class-aws-s3-batchdelete-methods.md)

#### delete()  [header link](class-aws-s3-batchdelete-method-delete.md)

Synchronously deletes all of the objects.

`
    public
                    delete() : mixed`

##### Tags  [header link](class-aws-s3-batchdelete-method-delete-tags.md)

throws[DeleteMultipleObjectsException](class-aws-s3-exception-deletemultipleobjectsexception.md)

on error.

#### fromIterator()  [header link](class-aws-s3-batchdelete-method-fromiterator.md)

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

[BatchDelete](class-aws-s3-batchdelete.md)

#### fromListObjects()  [header link](class-aws-s3-batchdelete-method-fromlistobjects.md)

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

[BatchDelete](class-aws-s3-batchdelete.md)

#### promise()  [header link](class-aws-s3-batchdelete-method-promise.md)

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
  - [Methods](class-aws-s3-batchdelete-toc-methods.md)
- Methods
  - [delete()](class-aws-s3-batchdelete-method-delete.md)
  - [fromIterator()](class-aws-s3-batchdelete-method-fromiterator.md)
  - [fromListObjects()](class-aws-s3-batchdelete-method-fromlistobjects.md)
  - [promise()](class-aws-s3-batchdelete-method-promise.md)

[Back To Top](class-aws-s3-batchdelete-top.md)
