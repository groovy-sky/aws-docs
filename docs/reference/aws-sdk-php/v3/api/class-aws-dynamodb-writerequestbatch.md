Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## WriteRequestBatch        in package    - [Aws](package-aws.md)

The WriteRequestBatch is an object that is capable of efficiently sending
DynamoDB BatchWriteItem requests from queued up put and delete item requests.

requests. The batch attempts to send the requests with the fewest requests
to DynamoDB as possible and also re-queues any unprocessed items to ensure
that all items are sent.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#method___construct)
: mixed Creates a WriteRequestBatch object that is capable of efficiently sending
DynamoDB BatchWriteItem requests from queued up Put and Delete requests.[delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#method_delete)
: $this Adds a delete item request to the batch.[flush()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#method_flush)
: $this Flushes the batch by combining all the queued put and delete requests
into BatchWriteItem commands and executing them. Unprocessed items are
automatically re-queued.[put()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#method_put)
: $this Adds a put item request to the batch.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html\#method___construct)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html\#method___construct\#tags)

throwsInvalidArgumentException

if the batch size is not between 2 and 25.

#### delete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html\#method_delete)

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

#### flush()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html\#method_flush)

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

#### put()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html\#method_put)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#method___construct)
  - [delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#method_delete)
  - [flush()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#method_flush)
  - [put()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#method_put)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.WriteRequestBatch.html#top)
