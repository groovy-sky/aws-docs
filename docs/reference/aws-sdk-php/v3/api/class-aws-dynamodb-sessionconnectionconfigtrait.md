Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## SessionConnectionConfigTrait

### Table of Contents  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-toc.md)

#### Methods  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-toc-methods.md)

[getBatchConfig()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getbatchconfig.md)
: mixed [getDataAttribute()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getdataattribute.md)
: string [getDataAttributeType()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getdataattributetype.md)
: string [getHashKey()](class-aws-dynamodb-sessionconnectionconfigtrait-method-gethashkey.md)
: string [getMaxLockRetryMicrotime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getmaxlockretrymicrotime.md)
: number[getMaxLockWaitTime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getmaxlockwaittime.md)
: number[getMinLockRetryMicrotime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getminlockretrymicrotime.md)
: number[getSessionLifetime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getsessionlifetime.md)
: number[getSessionLifetimeAttribute()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getsessionlifetimeattribute.md)
: string [getTableName()](class-aws-dynamodb-sessionconnectionconfigtrait-method-gettablename.md)
: string [initConfig()](class-aws-dynamodb-sessionconnectionconfigtrait-method-initconfig.md)
: mixed It initialize the Config class and
it sets values in case of valid configurations.[isConsistentRead()](class-aws-dynamodb-sessionconnectionconfigtrait-method-isconsistentread.md)
: bool [isLocking()](class-aws-dynamodb-sessionconnectionconfigtrait-method-islocking.md)
: bool [setBatchConfig()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setbatchconfig.md)
: mixed [setConsistentRead()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setconsistentread.md)
: mixed [setDataAttribute()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setdataattribute.md)
: mixed [setDataAttributeType()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setdataattributetype.md)
: mixed [setHashKey()](class-aws-dynamodb-sessionconnectionconfigtrait-method-sethashkey.md)
: mixed [setLocking()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setlocking.md)
: mixed [setMaxLockRetryMicrotime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setmaxlockretrymicrotime.md)
: mixed [setMaxLockWaitTime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setmaxlockwaittime.md)
: mixed [setMinLockRetryMicrotime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setminlockretrymicrotime.md)
: mixed [setSessionLifetime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setsessionlifetime.md)
: mixed [setSessionLifetimeAttribute()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setsessionlifetimeattribute.md)
: mixed [setTableName()](class-aws-dynamodb-sessionconnectionconfigtrait-method-settablename.md)
: mixed

### Methods  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-methods.md)

#### getBatchConfig()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-getbatchconfig.md)

`
    public
                    getBatchConfig() : mixed`

#### getDataAttribute()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-getdataattribute.md)

`
    public
                    getDataAttribute() : string`

##### Return values

string

#### getDataAttributeType()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-getdataattributetype.md)

`
    public
                    getDataAttributeType() : string`

##### Return values

string

#### getHashKey()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-gethashkey.md)

`
    public
                    getHashKey() : string`

##### Return values

string

#### getMaxLockRetryMicrotime()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-getmaxlockretrymicrotime.md)

`
    public
                    getMaxLockRetryMicrotime() : number`

##### Return values

number

#### getMaxLockWaitTime()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-getmaxlockwaittime.md)

`
    public
                    getMaxLockWaitTime() : number`

##### Return values

number

#### getMinLockRetryMicrotime()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-getminlockretrymicrotime.md)

`
    public
                    getMinLockRetryMicrotime() : number`

##### Return values

number

#### getSessionLifetime()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-getsessionlifetime.md)

`
    public
                    getSessionLifetime() : number`

##### Return values

number

#### getSessionLifetimeAttribute()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-getsessionlifetimeattribute.md)

`
    public
                    getSessionLifetimeAttribute() : string`

##### Return values

string

#### getTableName()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-gettablename.md)

`
    public
                    getTableName() : string`

##### Return values

string

#### initConfig()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-initconfig.md)

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

#### isConsistentRead()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-isconsistentread.md)

`
    public
                    isConsistentRead() : bool`

##### Return values

bool

#### isLocking()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-islocking.md)

`
    public
                    isLocking() : bool`

##### Return values

bool

#### setBatchConfig()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setbatchconfig.md)

`
    public
                    setBatchConfig(mixed $batchConfig) : mixed`

##### Parameters

$batchConfig
: mixed

#### setConsistentRead()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setconsistentread.md)

`
    public
                    setConsistentRead(bool $consistentRead) : mixed`

##### Parameters

$consistentRead
: bool

#### setDataAttribute()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setdataattribute.md)

`
    public
                    setDataAttribute(string $dataAttribute) : mixed`

##### Parameters

$dataAttribute
: string

#### setDataAttributeType()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setdataattributetype.md)

`
    public
                    setDataAttributeType(string $dataAttributeType) : mixed`

##### Parameters

$dataAttributeType
: string

#### setHashKey()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-sethashkey.md)

`
    public
                    setHashKey(string $hashKey) : mixed`

##### Parameters

$hashKey
: string

#### setLocking()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setlocking.md)

`
    public
                    setLocking(bool $locking) : mixed`

##### Parameters

$locking
: bool

#### setMaxLockRetryMicrotime()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setmaxlockretrymicrotime.md)

`
    public
                    setMaxLockRetryMicrotime(number $maxLockRetryMicrotime) : mixed`

##### Parameters

$maxLockRetryMicrotime
: number

#### setMaxLockWaitTime()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setmaxlockwaittime.md)

`
    public
                    setMaxLockWaitTime(number $maxLockWaitTime) : mixed`

##### Parameters

$maxLockWaitTime
: number

#### setMinLockRetryMicrotime()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setminlockretrymicrotime.md)

`
    public
                    setMinLockRetryMicrotime(number $minLockRetryMicrotime) : mixed`

##### Parameters

$minLockRetryMicrotime
: number

#### setSessionLifetime()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setsessionlifetime.md)

`
    public
                    setSessionLifetime(number $sessionLifetime) : mixed`

##### Parameters

$sessionLifetime
: number

#### setSessionLifetimeAttribute()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-setsessionlifetimeattribute.md)

`
    public
                    setSessionLifetimeAttribute(string $sessionLifetimeAttribute) : mixed`

##### Parameters

$sessionLifetimeAttribute
: string

#### setTableName()  [header link](class-aws-dynamodb-sessionconnectionconfigtrait-method-settablename.md)

`
    public
                    setTableName(string $tableName) : mixed`

##### Parameters

$tableName
: string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-dynamodb-sessionconnectionconfigtrait-toc-methods.md)
- Methods
  - [getBatchConfig()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getbatchconfig.md)
  - [getDataAttribute()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getdataattribute.md)
  - [getDataAttributeType()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getdataattributetype.md)
  - [getHashKey()](class-aws-dynamodb-sessionconnectionconfigtrait-method-gethashkey.md)
  - [getMaxLockRetryMicrotime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getmaxlockretrymicrotime.md)
  - [getMaxLockWaitTime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getmaxlockwaittime.md)
  - [getMinLockRetryMicrotime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getminlockretrymicrotime.md)
  - [getSessionLifetime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getsessionlifetime.md)
  - [getSessionLifetimeAttribute()](class-aws-dynamodb-sessionconnectionconfigtrait-method-getsessionlifetimeattribute.md)
  - [getTableName()](class-aws-dynamodb-sessionconnectionconfigtrait-method-gettablename.md)
  - [initConfig()](class-aws-dynamodb-sessionconnectionconfigtrait-method-initconfig.md)
  - [isConsistentRead()](class-aws-dynamodb-sessionconnectionconfigtrait-method-isconsistentread.md)
  - [isLocking()](class-aws-dynamodb-sessionconnectionconfigtrait-method-islocking.md)
  - [setBatchConfig()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setbatchconfig.md)
  - [setConsistentRead()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setconsistentread.md)
  - [setDataAttribute()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setdataattribute.md)
  - [setDataAttributeType()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setdataattributetype.md)
  - [setHashKey()](class-aws-dynamodb-sessionconnectionconfigtrait-method-sethashkey.md)
  - [setLocking()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setlocking.md)
  - [setMaxLockRetryMicrotime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setmaxlockretrymicrotime.md)
  - [setMaxLockWaitTime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setmaxlockwaittime.md)
  - [setMinLockRetryMicrotime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setminlockretrymicrotime.md)
  - [setSessionLifetime()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setsessionlifetime.md)
  - [setSessionLifetimeAttribute()](class-aws-dynamodb-sessionconnectionconfigtrait-method-setsessionlifetimeattribute.md)
  - [setTableName()](class-aws-dynamodb-sessionconnectionconfigtrait-method-settablename.md)

[Back To Top](class-aws-dynamodb-sessionconnectionconfigtrait-top.md)
