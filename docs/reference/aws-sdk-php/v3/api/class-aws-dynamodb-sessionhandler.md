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

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#toc-interfaces)

SessionHandlerInterface

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method___construct)
: mixed [close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_close)
: bool Close a session from writing.[destroy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_destroy)
: bool Delete a session stored in DynamoDB.[fromClient()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_fromClient)
: [SessionHandler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html)Creates a new DynamoDB Session Handler.[garbageCollect()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_garbageCollect)
: mixed Triggers garbage collection on expired sessions.[gc()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_gc)
: bool Satisfies the session handler interface, but does nothing. To do garbage
collection, you must manually call the garbageCollect() method.[open()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_open)
: bool Open a session for writing. Triggered by session\_start().[read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_read)
: string Read a session stored in DynamoDB.[register()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_register)
: bool Register the DynamoDB session handler.[write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_write)
: bool Write a session to DynamoDB.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method___construct)

`
    public
                    __construct(SessionConnectionInterface $connection) : mixed`

##### Parameters

$connection
: [SessionConnectionInterface](class-aws-dynamodb-sessionconnectioninterface.md)

#### close()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_close)

Close a session from writing.

`
    public
                    close() : bool`

##### Return values

bool
—

Success

#### destroy()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_destroy)

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

#### fromClient()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_fromClient)

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

[SessionHandler](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html)

#### garbageCollect()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_garbageCollect)

Triggers garbage collection on expired sessions.

`
    public
                    garbageCollect() : mixed`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_garbageCollect\#tags)

codeCoverageIgnore

#### gc()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_gc)

Satisfies the session handler interface, but does nothing. To do garbage
collection, you must manually call the garbageCollect() method.

`
    public
                    gc(int $maxLifetime) : bool`

##### Parameters

$maxLifetime
: int

Ignored.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_gc\#tags)

codeCoverageIgnore

##### Return values

bool
—

Whether or not the operation succeeded.

#### open()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_open)

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

#### read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_read)

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

#### register()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_register)

Register the DynamoDB session handler.

`
    public
                    register() : bool`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_register\#tags)

codeCoverageIgnore

##### Return values

bool
—

Whether or not the handler was registered.

#### write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html\#method_write)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method___construct)
  - [close()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_close)
  - [destroy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_destroy)
  - [fromClient()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_fromClient)
  - [garbageCollect()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_garbageCollect)
  - [gc()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_gc)
  - [open()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_open)
  - [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_read)
  - [register()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_register)
  - [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#method_write)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionHandler.html#top)
