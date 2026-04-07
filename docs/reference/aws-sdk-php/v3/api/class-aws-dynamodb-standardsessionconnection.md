Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## StandardSessionConnection        in package    - [Aws](package-aws.md)       implements  [SessionConnectionInterface](class-aws-dynamodb-sessionconnectioninterface.md)  Uses  [SessionConnectionConfigTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html)

The standard connection performs the read and write operations to DynamoDB.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#toc-interfaces)

[SessionConnectionInterface](class-aws-dynamodb-sessionconnectioninterface.md)The session connection provides the underlying logic for interacting with
Amazon DynamoDB and performs all of the reading and writing operations.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method___construct)
: mixed [delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method_delete)
: bool Deletes session record from DynamoDB[deleteExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method_deleteExpired)
: bool Performs garbage collection on the sessions stored in the DynamoDB[getBatchConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getBatchConfig)
: mixed [getDataAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getDataAttribute)
: string [getDataAttributeType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getDataAttributeType)
: string [getHashKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getHashKey)
: string [getMaxLockRetryMicrotime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getMaxLockRetryMicrotime)
: number[getMaxLockWaitTime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getMaxLockWaitTime)
: number[getMinLockRetryMicrotime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getMinLockRetryMicrotime)
: number[getSessionLifetime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getSessionLifetime)
: number[getSessionLifetimeAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getSessionLifetimeAttribute)
: string [getTableName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getTableName)
: string [initConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_initConfig)
: mixed It initialize the Config class and
it sets values in case of valid configurations.[isConsistentRead()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_isConsistentRead)
: bool [isLocking()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_isLocking)
: bool [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method_read)
: array<string\|int, mixed> Reads session data from DynamoDB[setBatchConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setBatchConfig)
: mixed [setConsistentRead()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setConsistentRead)
: mixed [setDataAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setDataAttribute)
: mixed [setDataAttributeType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setDataAttributeType)
: mixed [setHashKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setHashKey)
: mixed [setLocking()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setLocking)
: mixed [setMaxLockRetryMicrotime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setMaxLockRetryMicrotime)
: mixed [setMaxLockWaitTime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setMaxLockWaitTime)
: mixed [setMinLockRetryMicrotime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setMinLockRetryMicrotime)
: mixed [setSessionLifetime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setSessionLifetime)
: mixed [setSessionLifetimeAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setSessionLifetimeAttribute)
: mixed [setTableName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setTableName)
: mixed [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method_write)
: bool Writes session data to DynamoDB

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#method___construct)

`
    public
                    __construct(DynamoDbClient $client[, array<string|int, mixed> $config = [] ]) : mixed`

##### Parameters

$client
: [DynamoDbClient](class-aws-dynamodb-dynamodbclient.md)

DynamoDB client

$config
: array<string\|int, mixed>
= \[\]

Session handler config

#### delete()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#method_delete)

Deletes session record from DynamoDB

`
    public
                    delete(mixed $id) : bool`

##### Parameters

$id
: mixed

Session ID

##### Return values

bool

#### deleteExpired()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#method_deleteExpired)

Performs garbage collection on the sessions stored in the DynamoDB

`
    public
                    deleteExpired() : bool`

##### Return values

bool

#### getBatchConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getBatchConfig)

`
    public
                    getBatchConfig() : mixed`

#### getDataAttribute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getDataAttribute)

`
    public
                    getDataAttribute() : string`

##### Return values

string

#### getDataAttributeType()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getDataAttributeType)

`
    public
                    getDataAttributeType() : string`

##### Return values

string

#### getHashKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getHashKey)

`
    public
                    getHashKey() : string`

##### Return values

string

#### getMaxLockRetryMicrotime()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getMaxLockRetryMicrotime)

`
    public
                    getMaxLockRetryMicrotime() : number`

##### Return values

number

#### getMaxLockWaitTime()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getMaxLockWaitTime)

`
    public
                    getMaxLockWaitTime() : number`

##### Return values

number

#### getMinLockRetryMicrotime()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getMinLockRetryMicrotime)

`
    public
                    getMinLockRetryMicrotime() : number`

##### Return values

number

#### getSessionLifetime()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getSessionLifetime)

`
    public
                    getSessionLifetime() : number`

##### Return values

number

#### getSessionLifetimeAttribute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getSessionLifetimeAttribute)

`
    public
                    getSessionLifetimeAttribute() : string`

##### Return values

string

#### getTableName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_getTableName)

`
    public
                    getTableName() : string`

##### Return values

string

#### initConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_initConfig)

It initialize the Config class and
it sets values in case of valid configurations.

`
    public
                    initConfig([array<string|int, mixed> $config = [] ]) : mixed`

It transforms parameters underscore separated in camelcase "this\_is\_a\_test" => ThisIsATest
and it uses it in order to set the values.

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

#### isConsistentRead()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_isConsistentRead)

`
    public
                    isConsistentRead() : bool`

##### Return values

bool

#### isLocking()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_isLocking)

`
    public
                    isLocking() : bool`

##### Return values

bool

#### read()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#method_read)

Reads session data from DynamoDB

`
    public
                    read(mixed $id) : array<string|int, mixed>`

##### Parameters

$id
: mixed

Session ID

##### Return values

array<string\|int, mixed>

#### setBatchConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setBatchConfig)

`
    public
                    setBatchConfig(mixed $batchConfig) : mixed`

##### Parameters

$batchConfig
: mixed

#### setConsistentRead()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setConsistentRead)

`
    public
                    setConsistentRead(bool $consistentRead) : mixed`

##### Parameters

$consistentRead
: bool

#### setDataAttribute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setDataAttribute)

`
    public
                    setDataAttribute(string $dataAttribute) : mixed`

##### Parameters

$dataAttribute
: string

#### setDataAttributeType()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setDataAttributeType)

`
    public
                    setDataAttributeType(string $dataAttributeType) : mixed`

##### Parameters

$dataAttributeType
: string

#### setHashKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setHashKey)

`
    public
                    setHashKey(string $hashKey) : mixed`

##### Parameters

$hashKey
: string

#### setLocking()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setLocking)

`
    public
                    setLocking(bool $locking) : mixed`

##### Parameters

$locking
: bool

#### setMaxLockRetryMicrotime()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setMaxLockRetryMicrotime)

`
    public
                    setMaxLockRetryMicrotime(number $maxLockRetryMicrotime) : mixed`

##### Parameters

$maxLockRetryMicrotime
: number

#### setMaxLockWaitTime()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setMaxLockWaitTime)

`
    public
                    setMaxLockWaitTime(number $maxLockWaitTime) : mixed`

##### Parameters

$maxLockWaitTime
: number

#### setMinLockRetryMicrotime()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setMinLockRetryMicrotime)

`
    public
                    setMinLockRetryMicrotime(number $minLockRetryMicrotime) : mixed`

##### Parameters

$minLockRetryMicrotime
: number

#### setSessionLifetime()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setSessionLifetime)

`
    public
                    setSessionLifetime(number $sessionLifetime) : mixed`

##### Parameters

$sessionLifetime
: number

#### setSessionLifetimeAttribute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setSessionLifetimeAttribute)

`
    public
                    setSessionLifetimeAttribute(string $sessionLifetimeAttribute) : mixed`

##### Parameters

$sessionLifetimeAttribute
: string

#### setTableName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html\#method_setTableName)

`
    public
                    setTableName(string $tableName) : mixed`

##### Parameters

$tableName
: string

#### write()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html\#method_write)

Writes session data to DynamoDB

`
    public
                    write(mixed $id, mixed $data, mixed $isChanged) : bool`

##### Parameters

$id
: mixed

Session ID

$data
: mixed

Serialized session data

$isChanged
: mixed

Whether or not the data has changed

##### Return values

bool
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method___construct)
  - [delete()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method_delete)
  - [deleteExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method_deleteExpired)
  - [getBatchConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getBatchConfig)
  - [getDataAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getDataAttribute)
  - [getDataAttributeType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getDataAttributeType)
  - [getHashKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getHashKey)
  - [getMaxLockRetryMicrotime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getMaxLockRetryMicrotime)
  - [getMaxLockWaitTime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getMaxLockWaitTime)
  - [getMinLockRetryMicrotime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getMinLockRetryMicrotime)
  - [getSessionLifetime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getSessionLifetime)
  - [getSessionLifetimeAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getSessionLifetimeAttribute)
  - [getTableName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_getTableName)
  - [initConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_initConfig)
  - [isConsistentRead()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_isConsistentRead)
  - [isLocking()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_isLocking)
  - [read()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method_read)
  - [setBatchConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setBatchConfig)
  - [setConsistentRead()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setConsistentRead)
  - [setDataAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setDataAttribute)
  - [setDataAttributeType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setDataAttributeType)
  - [setHashKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setHashKey)
  - [setLocking()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setLocking)
  - [setMaxLockRetryMicrotime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setMaxLockRetryMicrotime)
  - [setMaxLockWaitTime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setMaxLockWaitTime)
  - [setMinLockRetryMicrotime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setMinLockRetryMicrotime)
  - [setSessionLifetime()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setSessionLifetime)
  - [setSessionLifetimeAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setSessionLifetimeAttribute)
  - [setTableName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SessionConnectionConfigTrait.html#method_setTableName)
  - [write()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#method_write)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.StandardSessionConnection.html#top)
