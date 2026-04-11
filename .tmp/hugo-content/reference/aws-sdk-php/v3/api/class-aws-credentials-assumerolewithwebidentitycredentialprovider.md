Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## AssumeRoleWithWebIdentityCredentialProvider        in package    - [Aws](package-aws.md)

Credential provider that provides credentials via assuming a role with a web identity
More Information, see: https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sts-2011-06-15.html#assumerolewithwebidentity

### Table of Contents  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-toc.md)

#### Constants  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-toc-constants.md)

[ENV\_RETRIES](class-aws-credentials-assumerolewithwebidentitycredentialprovider-constant-env-retries.md)
= 'AWS\_METADATA\_SERVICE\_NUM\_ATTEMPTS' [ERROR\_MSG](class-aws-credentials-assumerolewithwebidentitycredentialprovider-constant-error-msg.md)
= "Missing required 'AssumeRoleWithWebIdentityCredentialProvider' configuration option: "

#### Methods  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-toc-methods.md)

[\_\_construct()](class-aws-credentials-assumerolewithwebidentitycredentialprovider-method-construct.md)
: mixed The constructor attempts to load config from environment variables.[\_\_invoke()](class-aws-credentials-assumerolewithwebidentitycredentialprovider-method-invoke.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Loads assume role with web identity credentials.

### Constants  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-constants.md)

#### ENV\_RETRIES  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-constant-env-retries.md)

`
    public
        mixed
    ENV_RETRIES
    = 'AWS_METADATA_SERVICE_NUM_ATTEMPTS'
`

#### ERROR\_MSG  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-constant-error-msg.md)

`
    public
        mixed
    ERROR_MSG
    = "Missing required 'AssumeRoleWithWebIdentityCredentialProvider' configuration option: "
`

### Methods  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-methods.md)

#### \_\_construct()  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-method-construct.md)

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

##### Tags  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-method-construct-tags.md)

throwsInvalidArgumentException

#### \_\_invoke()  [header link](class-aws-credentials-assumerolewithwebidentitycredentialprovider-method-invoke.md)

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
  - [Constants](class-aws-credentials-assumerolewithwebidentitycredentialprovider-toc-constants.md)
  - [Methods](class-aws-credentials-assumerolewithwebidentitycredentialprovider-toc-methods.md)
- Constants
  - [ENV\_RETRIES](class-aws-credentials-assumerolewithwebidentitycredentialprovider-constant-env-retries.md)
  - [ERROR\_MSG](class-aws-credentials-assumerolewithwebidentitycredentialprovider-constant-error-msg.md)
- Methods
  - [\_\_construct()](class-aws-credentials-assumerolewithwebidentitycredentialprovider-method-construct.md)
  - [\_\_invoke()](class-aws-credentials-assumerolewithwebidentitycredentialprovider-method-invoke.md)

[Back To Top](class-aws-credentials-assumerolewithwebidentitycredentialprovider-top.md)
