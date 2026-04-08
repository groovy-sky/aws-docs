Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## Credentials     extends AwsCredentialIdentity   in package    - [Aws](package-aws.md)       implements  [CredentialsInterface](class-aws-credentials-credentialsinterface.md), Serializable

Basic implementation of the AWS Credentials interface that allows callers to
pass in the AWS Access Key and AWS Secret Access Key in the constructor.

### Table of Contents  [header link](class-aws-credentials-credentials-toc.md)

#### Interfaces  [header link](class-aws-credentials-credentials-toc-interfaces.md)

[CredentialsInterface](class-aws-credentials-credentialsinterface.md)Provides access to the AWS credentials used for accessing AWS services: AWS
access key ID, secret access key, and security token. These credentials are
used to securely sign requests to AWS services.Serializable

#### Methods  [header link](class-aws-credentials-credentials-toc-methods.md)

[\_\_construct()](class-aws-credentials-credentials-method-construct.md)
: mixed Constructs a new BasicAWSCredentials object, with the specified AWS
access key and AWS secret key[\_\_serialize()](class-aws-credentials-credentials-method-serialize.md)
: mixed [\_\_set\_state()](class-aws-credentials-credentials-method-set-state.md)
: mixed [\_\_unserialize()](class-aws-credentials-credentials-method-unserialize.md)
: mixed [getAccessKeyId()](class-aws-credentials-credentials-method-getaccesskeyid.md)
: string Returns the AWS access key ID for this credentials object.[getAccountId()](class-aws-credentials-credentials-method-getaccountid.md)
: mixed [getExpiration()](class-aws-credentials-credentials-method-getexpiration.md)
: int\|null Get the UNIX timestamp in which the credentials will expire[getSecretKey()](class-aws-credentials-credentials-method-getsecretkey.md)
: string Returns the AWS secret access key for this credentials object.[getSecurityToken()](class-aws-credentials-credentials-method-getsecuritytoken.md)
: string\|null Get the associated security token if available[getSource()](class-aws-credentials-credentials-method-getsource.md)
: mixed [isExpired()](class-aws-credentials-credentials-method-isexpired.md)
: bool Check if the credentials are expired[serialize()](class-aws-credentials-credentials-method-serialize.md)
: mixed [toArray()](class-aws-credentials-credentials-method-toarray.md)
: array<string\|int, mixed> Converts the credentials to an associative array.[unserialize()](class-aws-credentials-credentials-method-unserialize.md)
: mixed

### Methods  [header link](class-aws-credentials-credentials-methods.md)

#### \_\_construct()  [header link](class-aws-credentials-credentials-method-construct.md)

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

#### \_\_serialize()  [header link](class-aws-credentials-credentials-method-serialize.md)

`
    public
                    __serialize() : mixed`

#### \_\_set\_state()  [header link](class-aws-credentials-credentials-method-set-state.md)

`
    public
            static        __set_state(array<string|int, mixed> $state) : mixed`

##### Parameters

$state
: array<string\|int, mixed>

#### \_\_unserialize()  [header link](class-aws-credentials-credentials-method-unserialize.md)

`
    public
                    __unserialize(mixed $data) : mixed`

##### Parameters

$data
: mixed

#### getAccessKeyId()  [header link](class-aws-credentials-credentials-method-getaccesskeyid.md)

Returns the AWS access key ID for this credentials object.

`
    public
                    getAccessKeyId() : string`

##### Return values

string

#### getAccountId()  [header link](class-aws-credentials-credentials-method-getaccountid.md)

`
    public
                    getAccountId() : mixed`

#### getExpiration()  [header link](class-aws-credentials-credentials-method-getexpiration.md)

Get the UNIX timestamp in which the credentials will expire

`
    public
                    getExpiration() : int|null`

##### Return values

int\|null

#### getSecretKey()  [header link](class-aws-credentials-credentials-method-getsecretkey.md)

Returns the AWS secret access key for this credentials object.

`
    public
                    getSecretKey() : string`

##### Return values

string

#### getSecurityToken()  [header link](class-aws-credentials-credentials-method-getsecuritytoken.md)

Get the associated security token if available

`
    public
                    getSecurityToken() : string|null`

##### Return values

string\|null

#### getSource()  [header link](class-aws-credentials-credentials-method-getsource.md)

`
    public
                    getSource() : mixed`

#### isExpired()  [header link](class-aws-credentials-credentials-method-isexpired.md)

Check if the credentials are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### serialize()  [header link](class-aws-credentials-credentials-method-serialize.md)

`
    public
                    serialize() : mixed`

#### toArray()  [header link](class-aws-credentials-credentials-method-toarray.md)

Converts the credentials to an associative array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### unserialize()  [header link](class-aws-credentials-credentials-method-unserialize.md)

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
  - [Methods](class-aws-credentials-credentials-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-credentials-credentials-method-construct.md)
  - [\_\_serialize()](class-aws-credentials-credentials-method-serialize.md)
  - [\_\_set\_state()](class-aws-credentials-credentials-method-set-state.md)
  - [\_\_unserialize()](class-aws-credentials-credentials-method-unserialize.md)
  - [getAccessKeyId()](class-aws-credentials-credentials-method-getaccesskeyid.md)
  - [getAccountId()](class-aws-credentials-credentials-method-getaccountid.md)
  - [getExpiration()](class-aws-credentials-credentials-method-getexpiration.md)
  - [getSecretKey()](class-aws-credentials-credentials-method-getsecretkey.md)
  - [getSecurityToken()](class-aws-credentials-credentials-method-getsecuritytoken.md)
  - [getSource()](class-aws-credentials-credentials-method-getsource.md)
  - [isExpired()](class-aws-credentials-credentials-method-isexpired.md)
  - [serialize()](class-aws-credentials-credentials-method-serialize.md)
  - [toArray()](class-aws-credentials-credentials-method-toarray.md)
  - [unserialize()](class-aws-credentials-credentials-method-unserialize.md)

[Back To Top](class-aws-credentials-credentials-top.md)
