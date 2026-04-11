Menu

- [Aws](namespace-aws.md)
- [Rds](namespace-aws-rds.md)

## AuthTokenGenerator        in package    - [Aws](package-aws.md)

Generates RDS auth tokens for use with IAM authentication.

### Table of Contents  [header link](class-aws-rds-authtokengenerator-toc.md)

#### Methods  [header link](class-aws-rds-authtokengenerator-toc-methods.md)

[\_\_construct()](class-aws-rds-authtokengenerator-method-construct.md)
: mixed The constructor takes an instance of Credentials or a CredentialProvider[createToken()](class-aws-rds-authtokengenerator-method-createtoken.md)
: string Create the token for database login

### Methods  [header link](class-aws-rds-authtokengenerator-methods.md)

#### \_\_construct()  [header link](class-aws-rds-authtokengenerator-method-construct.md)

The constructor takes an instance of Credentials or a CredentialProvider

`
    public
                    __construct(callable|Credentials $creds) : mixed`

##### Parameters

$creds
: callable\| [Credentials](class-aws-credentials-credentials.md)

#### createToken()  [header link](class-aws-rds-authtokengenerator-method-createtoken.md)

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
  - [Methods](class-aws-rds-authtokengenerator-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-rds-authtokengenerator-method-construct.md)
  - [createToken()](class-aws-rds-authtokengenerator-method-createtoken.md)

[Back To Top](class-aws-rds-authtokengenerator-top.md)
