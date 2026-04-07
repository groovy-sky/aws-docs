Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## AssumeRoleCredentialProvider        in package    - [Aws](package-aws.md)

Credential provider that provides credentials via assuming a role
More Information, see: http://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sts-2011-06-15.html#assumerole

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#toc-constants)

[ERROR\_MSG](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#constant_ERROR_MSG)
= "Missing required 'AssumeRoleCredentialProvider' configuration option: "

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#method___construct)
: mixed The constructor requires following configure parameters:
\- client: a StsClient
\- assume\_role\_params: Parameters used to make assumeRole call[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#method___invoke)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Loads assume role credentials.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#constants)

#### ERROR\_MSG  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#constant_ERROR_MSG)

`
    public
        mixed
    ERROR_MSG
    = "Missing required 'AssumeRoleCredentialProvider' configuration option: "
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#method___construct)

The constructor requires following configure parameters:
\- client: a StsClient
\- assume\_role\_params: Parameters used to make assumeRole call

`
    public
                    __construct([array<string|int, mixed> $config = [] ]) : mixed`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Configuration options

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#method___construct\#tags)

throwsInvalidArgumentException

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html\#method___invoke)

Loads assume role credentials.

`
    public
                    __invoke() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#toc-methods)
- Constants
  - [ERROR\_MSG](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#constant_ERROR_MSG)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#method___invoke)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html#top)
