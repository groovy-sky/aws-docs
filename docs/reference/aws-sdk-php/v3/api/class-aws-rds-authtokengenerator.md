Menu

- [Aws](namespace-aws.md)
- [Rds](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.rds.html)

## AuthTokenGenerator        in package    - [Aws](package-aws.md)

Generates RDS auth tokens for use with IAM authentication.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html#method___construct)
: mixed The constructor takes an instance of Credentials or a CredentialProvider[createToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html#method_createToken)
: string Create the token for database login

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html\#method___construct)

The constructor takes an instance of Credentials or a CredentialProvider

`
    public
                    __construct(callable|Credentials $creds) : mixed`

##### Parameters

$creds
: callable\| [Credentials](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html)

#### createToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html\#method_createToken)

Create the token for database login

`
    public
                    createToken(string $endpoint, string $region, string $username[, int $lifetime = 15 ]) : string`

##### Parameters

$endpoint
: string

The database hostname with port number specified
(e.g., host:port)

$region
: string

The region where the database is located

$username
: string

The username to login as

$lifetime
: int
= 15

The lifetime of the token in minutes

##### Return values

string
—

Token generated

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html#method___construct)
  - [createToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html#method_createToken)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Rds.AuthTokenGenerator.html#top)
