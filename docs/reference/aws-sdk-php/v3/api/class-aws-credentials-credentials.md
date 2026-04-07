Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## Credentials     extends AwsCredentialIdentity   in package    - [Aws](package-aws.md)       implements  [CredentialsInterface](class-aws-credentials-credentialsinterface.md), Serializable

Basic implementation of the AWS Credentials interface that allows callers to
pass in the AWS Access Key and AWS Secret Access Key in the constructor.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#toc-interfaces)

[CredentialsInterface](class-aws-credentials-credentialsinterface.md)Provides access to the AWS credentials used for accessing AWS services: AWS
access key ID, secret access key, and security token. These credentials are
used to securely sign requests to AWS services.Serializable

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method___construct)
: mixed Constructs a new BasicAWSCredentials object, with the specified AWS
access key and AWS secret key[\_\_serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method___serialize)
: mixed [\_\_set\_state()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method___set_state)
: mixed [\_\_unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method___unserialize)
: mixed [getAccessKeyId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getAccessKeyId)
: string Returns the AWS access key ID for this credentials object.[getAccountId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getAccountId)
: mixed [getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getExpiration)
: int\|null Get the UNIX timestamp in which the credentials will expire[getSecretKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getSecretKey)
: string Returns the AWS secret access key for this credentials object.[getSecurityToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getSecurityToken)
: string\|null Get the associated security token if available[getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getSource)
: mixed [isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_isExpired)
: bool Check if the credentials are expired[serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_serialize)
: mixed [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_toArray)
: array<string\|int, mixed> Converts the credentials to an associative array.[unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_unserialize)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method___construct)

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

#### \_\_serialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method___serialize)

`
    public
                    __serialize() : mixed`

#### \_\_set\_state()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method___set_state)

`
    public
            static        __set_state(array<string|int, mixed> $state) : mixed`

##### Parameters

$state
: array<string\|int, mixed>

#### \_\_unserialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method___unserialize)

`
    public
                    __unserialize(mixed $data) : mixed`

##### Parameters

$data
: mixed

#### getAccessKeyId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_getAccessKeyId)

Returns the AWS access key ID for this credentials object.

`
    public
                    getAccessKeyId() : string`

##### Return values

string

#### getAccountId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_getAccountId)

`
    public
                    getAccountId() : mixed`

#### getExpiration()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_getExpiration)

Get the UNIX timestamp in which the credentials will expire

`
    public
                    getExpiration() : int|null`

##### Return values

int\|null

#### getSecretKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_getSecretKey)

Returns the AWS secret access key for this credentials object.

`
    public
                    getSecretKey() : string`

##### Return values

string

#### getSecurityToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_getSecurityToken)

Get the associated security token if available

`
    public
                    getSecurityToken() : string|null`

##### Return values

string\|null

#### getSource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_getSource)

`
    public
                    getSource() : mixed`

#### isExpired()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_isExpired)

Check if the credentials are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### serialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_serialize)

`
    public
                    serialize() : mixed`

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_toArray)

Converts the credentials to an associative array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### unserialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html\#method_unserialize)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method___construct)
  - [\_\_serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method___serialize)
  - [\_\_set\_state()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method___set_state)
  - [\_\_unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method___unserialize)
  - [getAccessKeyId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getAccessKeyId)
  - [getAccountId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getAccountId)
  - [getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getExpiration)
  - [getSecretKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getSecretKey)
  - [getSecurityToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getSecurityToken)
  - [getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_getSource)
  - [isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_isExpired)
  - [serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_serialize)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_toArray)
  - [unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#method_unserialize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html#top)
