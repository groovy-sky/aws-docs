Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## WriteRequestBatch        in package    - [Aws](package-aws.md)

The WriteRequestBatch is an object that is capable of efficiently sending
DynamoDB BatchWriteItem requests from queued up put and delete item requests.

requests. The batch attempts to send the requests with the fewest requests
to DynamoDB as possible and also re-queues any unprocessed items to ensure
that all items are sent.

### Table of Contents  [header link](class-aws-dynamodb-writerequestbatch-toc.md)

#### Methods  [header link](class-aws-dynamodb-writerequestbatch-toc-methods.md)

[\_\_construct()](class-aws-dynamodb-writerequestbatch-method-construct.md)
: mixed Creates a WriteRequestBatch object that is capable of efficiently sending
DynamoDB BatchWriteItem requests from queued up Put and Delete requests.[delete()](class-aws-dynamodb-writerequestbatch-method-delete.md)
: $this Adds a delete item request to the batch.[flush()](class-aws-dynamodb-writerequestbatch-method-flush.md)
: $this Flushes the batch by combining all the queued put and delete requests
into BatchWriteItem commands and executing them. Unprocessed items are
automatically re-queued.[put()](class-aws-dynamodb-writerequestbatch-method-put.md)
: $this Adds a put item request to the batch.

### Methods  [header link](class-aws-dynamodb-writerequestbatch-methods.md)

#### \_\_construct()  [header link](class-aws-dynamodb-writerequestbatch-method-construct.md)

Creates a WriteRequestBatch object that is capable of efficiently sending
DynamoDB BatchWriteItem requests from queued up Put and Delete requests.

`
    public
                    __construct(DynamoDbClient $client[, array<string|int, mixed> $config = [] ]) : mixed`

##### Parameters

$client
: [DynamoDbClient](class-aws-dynamodb-dynamodbclient.md)

DynamoDB client used to send batches.

$config
: array<string\|int, mixed>
= \[\]

Batch configuration options.

- table: (string) DynamoDB table used by the batch, this can be
overridden for each individual put() or delete() call.
- batch\_size: (int) The size of each batch (default: 25). The batch
size must be between 2 and 25. If you are sending batches of large
items, you may consider lowering the batch size, otherwise, you
should use 25.
- pool\_size: (int) This number dictates how many BatchWriteItem
requests you would like to do in parallel. For example, if the
"batch\_size" is 25, and "pool\_size" is 3, then you would send 3
BatchWriteItem requests at a time, each with 25 items. Please keep
your throughput in mind when choosing the "pool\_size" option.
- autoflush: (bool) This option allows the batch to automatically
flush once there are enough items (i.e., "batch\_size" \* "pool\_size")
in the queue. This defaults to true, so you must set this to false
to stop autoflush.
- before: (callable) Executed before every BatchWriteItem operation.
It should accept an \\Aws\\CommandInterface object as its argument.
- error: Executed if an error was encountered executing a,
BatchWriteItem operation, otherwise errors are ignored. It should
accept an \\Aws\\Exception\\AwsException as its argument.

##### Tags  [header link](class-aws-dynamodb-writerequestbatch-method-construct-tags.md)

throwsInvalidArgumentException

if the batch size is not between 2 and 25.

#### delete()  [header link](class-aws-dynamodb-writerequestbatch-method-delete.md)

Adds a delete item request to the batch.

`
    public
                    delete(array<string|int, mixed> $key[, string|null $table = null ]) : $this`

##### Parameters

$key
: array<string\|int, mixed>

Key of an item to delete. Format:
\[
'key1' => \['type' => 'value'\],
...
\]

$table
: string\|null
= null

The name of the table. This must be specified
unless the "table" option was provided in the
config of the WriteRequestBatch.

##### Return values

$this

#### flush()  [header link](class-aws-dynamodb-writerequestbatch-method-flush.md)

Flushes the batch by combining all the queued put and delete requests
into BatchWriteItem commands and executing them. Unprocessed items are
automatically re-queued.

`
    public
                    flush([bool $untilEmpty = true ]) : $this`

##### Parameters

$untilEmpty
: bool
= true

If true, flushing will continue until the queue
is completely empty. This will make sure that
unprocessed items are all eventually sent.

##### Return values

$this

#### put()  [header link](class-aws-dynamodb-writerequestbatch-method-put.md)

Adds a put item request to the batch.

`
    public
                    put(array<string|int, mixed> $item[, string|null $table = null ]) : $this`

##### Parameters

$item
: array<string\|int, mixed>

Data for an item to put. Format:
\[
'attribute1' => \['type' => 'value'\],
'attribute2' => \['type' => 'value'\],
...
\]

$table
: string\|null
= null

The name of the table. This must be specified
unless the "table" option was provided in the
config of the WriteRequestBatch.

##### Return values

$this
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-dynamodb-writerequestbatch-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-dynamodb-writerequestbatch-method-construct.md)
  - [delete()](class-aws-dynamodb-writerequestbatch-method-delete.md)
  - [flush()](class-aws-dynamodb-writerequestbatch-method-flush.md)
  - [put()](class-aws-dynamodb-writerequestbatch-method-put.md)

[Back To Top](class-aws-dynamodb-writerequestbatch-top.md)
