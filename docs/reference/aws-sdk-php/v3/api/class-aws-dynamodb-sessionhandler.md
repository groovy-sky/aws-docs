Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## SessionHandler        in package    - [Aws](package-aws.md)       implements  SessionHandlerInterface

Provides an interface for using Amazon DynamoDB as a session store by hooking
into PHP's session handler hooks. Once registered, You may use the native
\`$\_SESSION\` superglobal and session functions, and the sessions will be
stored automatically in DynamoDB. DynamoDB is a great session storage
solution due to its speed, scalability, and fault tolerance.

For maximum performance, we recommend that you keep the size of your sessions
small. Locking is disabled by default, since it can drive up latencies and
costs under high traffic. Only turn it on if you need it.

By far, the most expensive operation is garbage collection. Therefore, we
encourage you to carefully consider your session garbage collection strategy.
Note: the DynamoDB Session Handler does not allow garbage collection to be
triggered randomly. You must run garbage collection manually or through other
automated means using a cron job or similar scheduling technique.

### Table of Contents  [header link](class-aws-dynamodb-sessionhandler-toc.md)

#### Interfaces  [header link](class-aws-dynamodb-sessionhandler-toc-interfaces.md)

SessionHandlerInterface

#### Methods  [header link](class-aws-dynamodb-sessionhandler-toc-methods.md)

[\_\_construct()](class-aws-dynamodb-sessionhandler-method-construct.md)
: mixed [close()](class-aws-dynamodb-sessionhandler-method-close.md)
: bool Close a session from writing.[destroy()](class-aws-dynamodb-sessionhandler-method-destroy.md)
: bool Delete a session stored in DynamoDB.[fromClient()](class-aws-dynamodb-sessionhandler-method-fromclient.md)
: [SessionHandler](class-aws-dynamodb-sessionhandler.md)Creates a new DynamoDB Session Handler.[garbageCollect()](class-aws-dynamodb-sessionhandler-method-garbagecollect.md)
: mixed Triggers garbage collection on expired sessions.[gc()](class-aws-dynamodb-sessionhandler-method-gc.md)
: bool Satisfies the session handler interface, but does nothing. To do garbage
collection, you must manually call the garbageCollect() method.[open()](class-aws-dynamodb-sessionhandler-method-open.md)
: bool Open a session for writing. Triggered by session\_start().[read()](class-aws-dynamodb-sessionhandler-method-read.md)
: string Read a session stored in DynamoDB.[register()](class-aws-dynamodb-sessionhandler-method-register.md)
: bool Register the DynamoDB session handler.[write()](class-aws-dynamodb-sessionhandler-method-write.md)
: bool Write a session to DynamoDB.

### Methods  [header link](class-aws-dynamodb-sessionhandler-methods.md)

#### \_\_construct()  [header link](class-aws-dynamodb-sessionhandler-method-construct.md)

`
    public
                    __construct(SessionConnectionInterface $connection) : mixed`

##### Parameters

$connection
: [SessionConnectionInterface](class-aws-dynamodb-sessionconnectioninterface.md)

#### close()  [header link](class-aws-dynamodb-sessionhandler-method-close.md)

Close a session from writing.

`
    public
                    close() : bool`

##### Return values

bool
—

Success

#### destroy()  [header link](class-aws-dynamodb-sessionhandler-method-destroy.md)

Delete a session stored in DynamoDB.

`
    public
                    destroy(string $id) : bool`

##### Parameters

$id
: string

Session ID.

##### Return values

bool
—

Whether or not the operation succeeded.

#### fromClient()  [header link](class-aws-dynamodb-sessionhandler-method-fromclient.md)

Creates a new DynamoDB Session Handler.

`
    public
            static        fromClient(DynamoDbClient $client[, array<string|int, mixed> $config = [] ]) : SessionHandler`

The configuration array accepts the following array keys and values:

- table\_name: Name of table to store the sessions.
- hash\_key: Name of hash key in table. Default: "id".
- data\_attribute: Name of the data attribute in table. Default: "data".
- session\_lifetime: Lifetime of inactive sessions expiration.
- session\_lifetime\_attribute: Name of the session life time attribute in table. Default: "expires".
- consistent\_read: Whether or not to use consistent reads.
- batch\_config: Batch options used for garbage collection.
- locking: Whether or not to use session locking.
- max\_lock\_wait\_time: Max time (s) to wait for lock acquisition.
- min\_lock\_retry\_microtime: Min time (µs) to wait between lock attempts.
- max\_lock\_retry\_microtime: Max time (µs) to wait between lock attempts.

You can find the full list of parameters and defaults within the trait
Aws\\DynamoDb\\SessionConnectionConfigTrait

##### Parameters

$client
: [DynamoDbClient](class-aws-dynamodb-dynamodbclient.md)

Client for doing DynamoDB operations

$config
: array<string\|int, mixed>
= \[\]

Configuration for the Session Handler

##### Return values

[SessionHandler](class-aws-dynamodb-sessionhandler.md)

#### garbageCollect()  [header link](class-aws-dynamodb-sessionhandler-method-garbagecollect.md)

Triggers garbage collection on expired sessions.

`
    public
                    garbageCollect() : mixed`

##### Tags  [header link](class-aws-dynamodb-sessionhandler-method-garbagecollect-tags.md)

codeCoverageIgnore

#### gc()  [header link](class-aws-dynamodb-sessionhandler-method-gc.md)

Satisfies the session handler interface, but does nothing. To do garbage
collection, you must manually call the garbageCollect() method.

`
    public
                    gc(int $maxLifetime) : bool`

##### Parameters

$maxLifetime
: int

Ignored.

##### Tags  [header link](class-aws-dynamodb-sessionhandler-method-gc-tags.md)

codeCoverageIgnore

##### Return values

bool
—

Whether or not the operation succeeded.

#### open()  [header link](class-aws-dynamodb-sessionhandler-method-open.md)

Open a session for writing. Triggered by session\_start().

`
    public
                    open(string $savePath, string $sessionName) : bool`

##### Parameters

$savePath
: string

Session save path.

$sessionName
: string

Session name.

##### Return values

bool
—

Whether or not the operation succeeded.

#### read()  [header link](class-aws-dynamodb-sessionhandler-method-read.md)

Read a session stored in DynamoDB.

`
    public
                    read(string $id) : string`

##### Parameters

$id
: string

Session ID.

##### Return values

string
—

Session data.

#### register()  [header link](class-aws-dynamodb-sessionhandler-method-register.md)

Register the DynamoDB session handler.

`
    public
                    register() : bool`

##### Tags  [header link](class-aws-dynamodb-sessionhandler-method-register-tags.md)

codeCoverageIgnore

##### Return values

bool
—

Whether or not the handler was registered.

#### write()  [header link](class-aws-dynamodb-sessionhandler-method-write.md)

Write a session to DynamoDB.

`
    public
                    write(string $id, string $data) : bool`

##### Parameters

$id
: string

Session ID.

$data
: string

Serialized session data to write.

##### Return values

bool
—

Whether or not the operation succeeded.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-dynamodb-sessionhandler-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-dynamodb-sessionhandler-method-construct.md)
  - [close()](class-aws-dynamodb-sessionhandler-method-close.md)
  - [destroy()](class-aws-dynamodb-sessionhandler-method-destroy.md)
  - [fromClient()](class-aws-dynamodb-sessionhandler-method-fromclient.md)
  - [garbageCollect()](class-aws-dynamodb-sessionhandler-method-garbagecollect.md)
  - [gc()](class-aws-dynamodb-sessionhandler-method-gc.md)
  - [open()](class-aws-dynamodb-sessionhandler-method-open.md)
  - [read()](class-aws-dynamodb-sessionhandler-method-read.md)
  - [register()](class-aws-dynamodb-sessionhandler-method-register.md)
  - [write()](class-aws-dynamodb-sessionhandler-method-write.md)

[Back To Top](class-aws-dynamodb-sessionhandler-top.md)
