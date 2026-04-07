Menu

- [Aws](namespace-aws.md)
- [DSQL](namespace-aws-dsql.md)

## AuthTokenGenerator        in package    - [Aws](package-aws.md)

Generates auth tokens to connect to DSQL database clusters.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html#method___construct)
: mixed The constructor takes an instance of Credentials or a CredentialProvider[generateDbConnectAdminAuthToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html#method_generateDbConnectAdminAuthToken)
: string [generateDbConnectAuthToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html#method_generateDbConnectAuthToken)
: string

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html\#method___construct)

The constructor takes an instance of Credentials or a CredentialProvider

`
    public
                    __construct(Credentials|callable $creds) : mixed`

##### Parameters

$creds
: [Credentials](class-aws-credentials-credentials.md) \|callable

#### generateDbConnectAdminAuthToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html\#method_generateDbConnectAdminAuthToken)

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

#### generateDbConnectAuthToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html\#method_generateDbConnectAuthToken)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html#method___construct)
  - [generateDbConnectAdminAuthToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html#method_generateDbConnectAdminAuthToken)
  - [generateDbConnectAuthToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html#method_generateDbConnectAuthToken)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DSQL.AuthTokenGenerator.html#top)
