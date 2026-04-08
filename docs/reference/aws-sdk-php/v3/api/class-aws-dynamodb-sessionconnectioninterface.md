Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## SessionConnectionInterface     in    - [Aws](package-aws.md)

The session connection provides the underlying logic for interacting with
Amazon DynamoDB and performs all of the reading and writing operations.

### Table of Contents  [header link](class-aws-dynamodb-sessionconnectioninterface-toc.md)

#### Methods  [header link](class-aws-dynamodb-sessionconnectioninterface-toc-methods.md)

[delete()](class-aws-dynamodb-sessionconnectioninterface-method-delete.md)
: bool Deletes session record from DynamoDB[deleteExpired()](class-aws-dynamodb-sessionconnectioninterface-method-deleteexpired.md)
: bool Performs garbage collection on the sessions stored in the DynamoDB[read()](class-aws-dynamodb-sessionconnectioninterface-method-read.md)
: array<string\|int, mixed> Reads session data from DynamoDB[write()](class-aws-dynamodb-sessionconnectioninterface-method-write.md)
: bool Writes session data to DynamoDB

### Methods  [header link](class-aws-dynamodb-sessionconnectioninterface-methods.md)

#### delete()  [header link](class-aws-dynamodb-sessionconnectioninterface-method-delete.md)

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

#### deleteExpired()  [header link](class-aws-dynamodb-sessionconnectioninterface-method-deleteexpired.md)

Performs garbage collection on the sessions stored in the DynamoDB

`
    public
                    deleteExpired() : bool`

##### Return values

bool

#### read()  [header link](class-aws-dynamodb-sessionconnectioninterface-method-read.md)

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

#### write()  [header link](class-aws-dynamodb-sessionconnectioninterface-method-write.md)

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
  - [Constants](class-aws-dynamodb-sessionconnectioninterface-toc-constants.md)
  - [Methods](class-aws-dynamodb-sessionconnectioninterface-toc-methods.md)
- Methods
  - [delete()](class-aws-dynamodb-sessionconnectioninterface-method-delete.md)
  - [deleteExpired()](class-aws-dynamodb-sessionconnectioninterface-method-deleteexpired.md)
  - [read()](class-aws-dynamodb-sessionconnectioninterface-method-read.md)
  - [write()](class-aws-dynamodb-sessionconnectioninterface-method-write.md)

[Back To Top](class-aws-dynamodb-sessionconnectioninterface-top.md)
