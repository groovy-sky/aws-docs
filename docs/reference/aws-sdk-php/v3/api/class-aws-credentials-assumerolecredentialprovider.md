Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## AssumeRoleCredentialProvider        in package    - [Aws](package-aws.md)

Credential provider that provides credentials via assuming a role
More Information, see: http://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sts-2011-06-15.html#assumerole

### Table of Contents  [header link](class-aws-credentials-assumerolecredentialprovider-toc.md)

#### Constants  [header link](class-aws-credentials-assumerolecredentialprovider-toc-constants.md)

[ERROR\_MSG](class-aws-credentials-assumerolecredentialprovider-constant-error-msg.md)
= "Missing required 'AssumeRoleCredentialProvider' configuration option: "

#### Methods  [header link](class-aws-credentials-assumerolecredentialprovider-toc-methods.md)

[\_\_construct()](class-aws-credentials-assumerolecredentialprovider-method-construct.md)
: mixed The constructor requires following configure parameters:
\- client: a StsClient
\- assume\_role\_params: Parameters used to make assumeRole call[\_\_invoke()](class-aws-credentials-assumerolecredentialprovider-method-invoke.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Loads assume role credentials.

### Constants  [header link](class-aws-credentials-assumerolecredentialprovider-constants.md)

#### ERROR\_MSG  [header link](class-aws-credentials-assumerolecredentialprovider-constant-error-msg.md)

`
    public
        mixed
    ERROR_MSG
    = "Missing required 'AssumeRoleCredentialProvider' configuration option: "
`

### Methods  [header link](class-aws-credentials-assumerolecredentialprovider-methods.md)

#### \_\_construct()  [header link](class-aws-credentials-assumerolecredentialprovider-method-construct.md)

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

##### Tags  [header link](class-aws-credentials-assumerolecredentialprovider-method-construct-tags.md)

throwsInvalidArgumentException

#### \_\_invoke()  [header link](class-aws-credentials-assumerolecredentialprovider-method-invoke.md)

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
  - [Constants](class-aws-credentials-assumerolecredentialprovider-toc-constants.md)
  - [Methods](class-aws-credentials-assumerolecredentialprovider-toc-methods.md)
- Constants
  - [ERROR\_MSG](class-aws-credentials-assumerolecredentialprovider-constant-error-msg.md)
- Methods
  - [\_\_construct()](class-aws-credentials-assumerolecredentialprovider-method-construct.md)
  - [\_\_invoke()](class-aws-credentials-assumerolecredentialprovider-method-invoke.md)

[Back To Top](class-aws-credentials-assumerolecredentialprovider-top.md)
