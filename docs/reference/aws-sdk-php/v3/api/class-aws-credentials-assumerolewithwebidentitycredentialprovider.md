Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## AssumeRoleWithWebIdentityCredentialProvider        in package    - [Aws](package-aws.md)

Credential provider that provides credentials via assuming a role with a web identity
More Information, see: https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sts-2011-06-15.html#assumerolewithwebidentity

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#toc-constants)

[ENV\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#constant_ENV_RETRIES)
= 'AWS\_METADATA\_SERVICE\_NUM\_ATTEMPTS' [ERROR\_MSG](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#constant_ERROR_MSG)
= "Missing required 'AssumeRoleWithWebIdentityCredentialProvider' configuration option: "

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#method___construct)
: mixed The constructor attempts to load config from environment variables.[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#method___invoke)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Loads assume role with web identity credentials.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#constants)

#### ENV\_RETRIES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#constant_ENV_RETRIES)

`
    public
        mixed
    ENV_RETRIES
    = 'AWS_METADATA_SERVICE_NUM_ATTEMPTS'
`

#### ERROR\_MSG  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#constant_ERROR_MSG)

`
    public
        mixed
    ERROR_MSG
    = "Missing required 'AssumeRoleWithWebIdentityCredentialProvider' configuration option: "
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#method___construct)

The constructor attempts to load config from environment variables.

`
    public
                    __construct([array<string|int, mixed> $config = [] ]) : mixed`

If not set, the following config options are used:

- WebIdentityTokenFile: full path of token filename
- RoleArn: arn of role to be assumed
- SessionName: (optional) set by SDK if not provided
- source: To identify if the provider was sourced by a profile or
from environment definition. Default will be `sts_web_id_token`.

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Configuration options

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#method___construct\#tags)

throwsInvalidArgumentException

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html\#method___invoke)

Loads assume role with web identity credentials.

`
    public
                    __invoke() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#toc-methods)
- Constants
  - [ENV\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#constant_ENV_RETRIES)
  - [ERROR\_MSG](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#constant_ERROR_MSG)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#method___invoke)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html#top)
