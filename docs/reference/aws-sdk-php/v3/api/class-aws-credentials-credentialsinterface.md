Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## CredentialsInterface     in    - [Aws](package-aws.md)

Provides access to the AWS credentials used for accessing AWS services: AWS
access key ID, secret access key, and security token. These credentials are
used to securely sign requests to AWS services.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#toc-methods)

[getAccessKeyId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_getAccessKeyId)
: string Returns the AWS access key ID for this credentials object.[getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_getExpiration)
: int\|null Get the UNIX timestamp in which the credentials will expire[getSecretKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_getSecretKey)
: string Returns the AWS secret access key for this credentials object.[getSecurityToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_getSecurityToken)
: string\|null Get the associated security token if available[isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_isExpired)
: bool Check if the credentials are expired[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_toArray)
: array<string\|int, mixed> Converts the credentials to an associative array.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#methods)

#### getAccessKeyId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#method_getAccessKeyId)

Returns the AWS access key ID for this credentials object.

`
    public
                    getAccessKeyId() : string`

##### Return values

string

#### getExpiration()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#method_getExpiration)

Get the UNIX timestamp in which the credentials will expire

`
    public
                    getExpiration() : int|null`

##### Return values

int\|null

#### getSecretKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#method_getSecretKey)

Returns the AWS secret access key for this credentials object.

`
    public
                    getSecretKey() : string`

##### Return values

string

#### getSecurityToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#method_getSecurityToken)

Get the associated security token if available

`
    public
                    getSecurityToken() : string|null`

##### Return values

string\|null

#### isExpired()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#method_isExpired)

Check if the credentials are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html\#method_toArray)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#toc-methods)
- Methods
  - [getAccessKeyId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_getAccessKeyId)
  - [getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_getExpiration)
  - [getSecretKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_getSecretKey)
  - [getSecurityToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_getSecurityToken)
  - [isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_isExpired)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html#top)
