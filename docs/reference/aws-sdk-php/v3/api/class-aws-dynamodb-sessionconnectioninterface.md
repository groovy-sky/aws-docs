Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## SessionConnectionInterface     in    - [Aws](package-aws.md)

The session connection provides the underlying logic for interacting with
Amazon DynamoDB and performs all of the reading and writing operations.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html\#toc-methods)

[delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#method_delete)
: bool Deletes session record from DynamoDB[deleteExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#method_deleteExpired)
: bool Performs garbage collection on the sessions stored in the DynamoDB[read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#method_read)
: array<string\|int, mixed> Reads session data from DynamoDB[write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#method_write)
: bool Writes session data to DynamoDB

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html\#methods)

#### delete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html\#method_delete)

Deletes session record from DynamoDB

`
    public
                    delete(string $id) : bool`

##### Parameters

$id
: string

Session ID

##### Return values

bool

#### deleteExpired()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html\#method_deleteExpired)

Performs garbage collection on the sessions stored in the DynamoDB

`
    public
                    deleteExpired() : bool`

##### Return values

bool

#### read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html\#method_read)

Reads session data from DynamoDB

`
    public
                    read(string $id) : array<string|int, mixed>`

##### Parameters

$id
: string

Session ID

##### Return values

array<string\|int, mixed>

#### write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html\#method_write)

Writes session data to DynamoDB

`
    public
                    write(string $id, string $data, bool $isChanged) : bool`

##### Parameters

$id
: string

Session ID

$data
: string

Serialized session data

$isChanged
: bool

Whether or not the data has changed

##### Return values

bool
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#toc-methods)
- Methods
  - [delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#method_delete)
  - [deleteExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#method_deleteExpired)
  - [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#method_read)
  - [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#method_write)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionInterface.html#top)
