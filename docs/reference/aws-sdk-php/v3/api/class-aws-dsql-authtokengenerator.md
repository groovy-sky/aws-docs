Menu

- [Aws](namespace-aws.md)
- [DSQL](namespace-aws-dsql.md)

## AuthTokenGenerator        in package    - [Aws](package-aws.md)

Generates auth tokens to connect to DSQL database clusters.

### Table of Contents  [header link](class-aws-dsql-authtokengenerator-toc.md)

#### Methods  [header link](class-aws-dsql-authtokengenerator-toc-methods.md)

[\_\_construct()](class-aws-dsql-authtokengenerator-method-construct.md)
: mixed The constructor takes an instance of Credentials or a CredentialProvider[generateDbConnectAdminAuthToken()](class-aws-dsql-authtokengenerator-method-generatedbconnectadminauthtoken.md)
: string [generateDbConnectAuthToken()](class-aws-dsql-authtokengenerator-method-generatedbconnectauthtoken.md)
: string

### Methods  [header link](class-aws-dsql-authtokengenerator-methods.md)

#### \_\_construct()  [header link](class-aws-dsql-authtokengenerator-method-construct.md)

The constructor takes an instance of Credentials or a CredentialProvider

`
    public
                    __construct(Credentials|callable $creds) : mixed`

##### Parameters

$creds
: [Credentials](class-aws-credentials-credentials.md) \|callable

#### generateDbConnectAdminAuthToken()  [header link](class-aws-dsql-authtokengenerator-method-generatedbconnectadminauthtoken.md)

`
    public
                    generateDbConnectAdminAuthToken(string $endpoint, string $region[, int $expiration = self::DEFAULT_EXPIRATION_TIME_SECONDS ]) : string`

##### Parameters

$endpoint
: string$region
: string$expiration
: int
= self::DEFAULT\_EXPIRATION\_TIME\_SECONDS

##### Return values

string

#### generateDbConnectAuthToken()  [header link](class-aws-dsql-authtokengenerator-method-generatedbconnectauthtoken.md)

`
    public
                    generateDbConnectAuthToken(string $endpoint, string $region[, int $expiration = self::DEFAULT_EXPIRATION_TIME_SECONDS ]) : string`

##### Parameters

$endpoint
: string$region
: string$expiration
: int
= self::DEFAULT\_EXPIRATION\_TIME\_SECONDS

##### Return values

string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-dsql-authtokengenerator-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-dsql-authtokengenerator-method-construct.md)
  - [generateDbConnectAdminAuthToken()](class-aws-dsql-authtokengenerator-method-generatedbconnectadminauthtoken.md)
  - [generateDbConnectAuthToken()](class-aws-dsql-authtokengenerator-method-generatedbconnectauthtoken.md)

[Back To Top](class-aws-dsql-authtokengenerator-top.md)
