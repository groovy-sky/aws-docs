Menu

- [Aws](namespace-aws.md)
- [Identity](namespace-aws-identity.md)
- [S3](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.identity.s3.html)

## S3ExpressIdentity     extends [Credentials](class-aws-credentials-credentials.md)   in package    - [Aws](package-aws.md)

Basic implementation of the AWS Credentials interface that allows callers to
pass in the AWS Access Key and AWS Secret Access Key in the constructor.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Identity.S3.S3ExpressIdentity.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Identity.S3.S3ExpressIdentity.html\#toc-methods)

[\_\_construct()](class-aws-credentials-credentials.md#method___construct)
: mixed Constructs a new BasicAWSCredentials object, with the specified AWS
access key and AWS secret key[\_\_serialize()](class-aws-credentials-credentials.md#method___serialize)
: mixed [\_\_set\_state()](class-aws-credentials-credentials.md#method___set_state)
: mixed [\_\_unserialize()](class-aws-credentials-credentials.md#method___unserialize)
: mixed [getAccessKeyId()](class-aws-credentials-credentials.md#method_getAccessKeyId)
: string Returns the AWS access key ID for this credentials object.[getAccountId()](class-aws-credentials-credentials.md#method_getAccountId)
: mixed [getExpiration()](class-aws-credentials-credentials.md#method_getExpiration)
: int\|null Get the UNIX timestamp in which the credentials will expire[getSecretKey()](class-aws-credentials-credentials.md#method_getSecretKey)
: string Returns the AWS secret access key for this credentials object.[getSecurityToken()](class-aws-credentials-credentials.md#method_getSecurityToken)
: string\|null Get the associated security token if available[getSource()](class-aws-credentials-credentials.md#method_getSource)
: mixed [isExpired()](class-aws-credentials-credentials.md#method_isExpired)
: bool Check if the credentials are expired[serialize()](class-aws-credentials-credentials.md#method_serialize)
: mixed [toArray()](class-aws-credentials-credentials.md#method_toArray)
: array<string\|int, mixed> Converts the credentials to an associative array.[unserialize()](class-aws-credentials-credentials.md#method_unserialize)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Identity.S3.S3ExpressIdentity.html\#methods)

#### \_\_construct()  [header link](class-aws-credentials-credentials.md\#method___construct)

Constructs a new BasicAWSCredentials object, with the specified AWS
access key and AWS secret key

`
    public
                    __construct(string $key, string $secret[, string $token = null ][, int $expires = null ][, mixed $accountId = null ][, mixed $source = CredentialSources::STATIC ]) : mixed`

##### Parameters

$key
: string

AWS access key ID

$secret
: string

AWS secret access key

$token
: string
= null

Security token to use

$expires
: int
= null

UNIX timestamp for when credentials expire

$accountId
: mixed
= null$source
: mixed
= CredentialSources::STATIC

#### \_\_serialize()  [header link](class-aws-credentials-credentials.md\#method___serialize)

`
    public
                    __serialize() : mixed`

#### \_\_set\_state()  [header link](class-aws-credentials-credentials.md\#method___set_state)

`
    public
            static        __set_state(array<string|int, mixed> $state) : mixed`

##### Parameters

$state
: array<string\|int, mixed>

#### \_\_unserialize()  [header link](class-aws-credentials-credentials.md\#method___unserialize)

`
    public
                    __unserialize(mixed $data) : mixed`

##### Parameters

$data
: mixed

#### getAccessKeyId()  [header link](class-aws-credentials-credentials.md\#method_getAccessKeyId)

Returns the AWS access key ID for this credentials object.

`
    public
                    getAccessKeyId() : string`

##### Return values

string

#### getAccountId()  [header link](class-aws-credentials-credentials.md\#method_getAccountId)

`
    public
                    getAccountId() : mixed`

#### getExpiration()  [header link](class-aws-credentials-credentials.md\#method_getExpiration)

Get the UNIX timestamp in which the credentials will expire

`
    public
                    getExpiration() : int|null`

##### Return values

int\|null

#### getSecretKey()  [header link](class-aws-credentials-credentials.md\#method_getSecretKey)

Returns the AWS secret access key for this credentials object.

`
    public
                    getSecretKey() : string`

##### Return values

string

#### getSecurityToken()  [header link](class-aws-credentials-credentials.md\#method_getSecurityToken)

Get the associated security token if available

`
    public
                    getSecurityToken() : string|null`

##### Return values

string\|null

#### getSource()  [header link](class-aws-credentials-credentials.md\#method_getSource)

`
    public
                    getSource() : mixed`

#### isExpired()  [header link](class-aws-credentials-credentials.md\#method_isExpired)

Check if the credentials are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### serialize()  [header link](class-aws-credentials-credentials.md\#method_serialize)

`
    public
                    serialize() : mixed`

#### toArray()  [header link](class-aws-credentials-credentials.md\#method_toArray)

Converts the credentials to an associative array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### unserialize()  [header link](class-aws-credentials-credentials.md\#method_unserialize)

`
    public
                    unserialize(mixed $serialized) : mixed`

##### Parameters

$serialized
: mixed
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Identity.S3.S3ExpressIdentity.html#toc-methods)
- Methods
  - [\_\_construct()](class-aws-credentials-credentials.md#method___construct)
  - [\_\_serialize()](class-aws-credentials-credentials.md#method___serialize)
  - [\_\_set\_state()](class-aws-credentials-credentials.md#method___set_state)
  - [\_\_unserialize()](class-aws-credentials-credentials.md#method___unserialize)
  - [getAccessKeyId()](class-aws-credentials-credentials.md#method_getAccessKeyId)
  - [getAccountId()](class-aws-credentials-credentials.md#method_getAccountId)
  - [getExpiration()](class-aws-credentials-credentials.md#method_getExpiration)
  - [getSecretKey()](class-aws-credentials-credentials.md#method_getSecretKey)
  - [getSecurityToken()](class-aws-credentials-credentials.md#method_getSecurityToken)
  - [getSource()](class-aws-credentials-credentials.md#method_getSource)
  - [isExpired()](class-aws-credentials-credentials.md#method_isExpired)
  - [serialize()](class-aws-credentials-credentials.md#method_serialize)
  - [toArray()](class-aws-credentials-credentials.md#method_toArray)
  - [unserialize()](class-aws-credentials-credentials.md#method_unserialize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Identity.S3.S3ExpressIdentity.html#top)
