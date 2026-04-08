Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## CredentialsInterface     in    - [Aws](package-aws.md)

Provides access to the AWS credentials used for accessing AWS services: AWS
access key ID, secret access key, and security token. These credentials are
used to securely sign requests to AWS services.

### Table of Contents  [header link](class-aws-credentials-credentialsinterface-toc.md)

#### Methods  [header link](class-aws-credentials-credentialsinterface-toc-methods.md)

[getAccessKeyId()](class-aws-credentials-credentialsinterface-method-getaccesskeyid.md)
: string Returns the AWS access key ID for this credentials object.[getExpiration()](class-aws-credentials-credentialsinterface-method-getexpiration.md)
: int\|null Get the UNIX timestamp in which the credentials will expire[getSecretKey()](class-aws-credentials-credentialsinterface-method-getsecretkey.md)
: string Returns the AWS secret access key for this credentials object.[getSecurityToken()](class-aws-credentials-credentialsinterface-method-getsecuritytoken.md)
: string\|null Get the associated security token if available[isExpired()](class-aws-credentials-credentialsinterface-method-isexpired.md)
: bool Check if the credentials are expired[toArray()](class-aws-credentials-credentialsinterface-method-toarray.md)
: array<string\|int, mixed> Converts the credentials to an associative array.

### Methods  [header link](class-aws-credentials-credentialsinterface-methods.md)

#### getAccessKeyId()  [header link](class-aws-credentials-credentialsinterface-method-getaccesskeyid.md)

Returns the AWS access key ID for this credentials object.

`
    public
                    getAccessKeyId() : string`

##### Return values

string

#### getExpiration()  [header link](class-aws-credentials-credentialsinterface-method-getexpiration.md)

Get the UNIX timestamp in which the credentials will expire

`
    public
                    getExpiration() : int|null`

##### Return values

int\|null

#### getSecretKey()  [header link](class-aws-credentials-credentialsinterface-method-getsecretkey.md)

Returns the AWS secret access key for this credentials object.

`
    public
                    getSecretKey() : string`

##### Return values

string

#### getSecurityToken()  [header link](class-aws-credentials-credentialsinterface-method-getsecuritytoken.md)

Get the associated security token if available

`
    public
                    getSecurityToken() : string|null`

##### Return values

string\|null

#### isExpired()  [header link](class-aws-credentials-credentialsinterface-method-isexpired.md)

Check if the credentials are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### toArray()  [header link](class-aws-credentials-credentialsinterface-method-toarray.md)

Converts the credentials to an associative array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-credentials-credentialsinterface-toc-constants.md)
  - [Methods](class-aws-credentials-credentialsinterface-toc-methods.md)
- Methods
  - [getAccessKeyId()](class-aws-credentials-credentialsinterface-method-getaccesskeyid.md)
  - [getExpiration()](class-aws-credentials-credentialsinterface-method-getexpiration.md)
  - [getSecretKey()](class-aws-credentials-credentialsinterface-method-getsecretkey.md)
  - [getSecurityToken()](class-aws-credentials-credentialsinterface-method-getsecuritytoken.md)
  - [isExpired()](class-aws-credentials-credentialsinterface-method-isexpired.md)
  - [toArray()](class-aws-credentials-credentialsinterface-method-toarray.md)

[Back To Top](class-aws-credentials-credentialsinterface-top.md)
