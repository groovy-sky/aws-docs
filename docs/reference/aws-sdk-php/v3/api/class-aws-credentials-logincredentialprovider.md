Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## LoginCredentialProvider        in package    - [Aws](package-aws.md)

FinalYes

Credential provider for login using console credentials

### Table of Contents  [header link](class-aws-credentials-logincredentialprovider-toc.md)

#### Methods  [header link](class-aws-credentials-logincredentialprovider-toc-methods.md)

[\_\_construct()](class-aws-credentials-logincredentialprovider-method-construct.md)
: mixed [\_\_invoke()](class-aws-credentials-logincredentialprovider-method-invoke.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that resolves to AWS credentials

### Methods  [header link](class-aws-credentials-logincredentialprovider-methods.md)

#### \_\_construct()  [header link](class-aws-credentials-logincredentialprovider-method-construct.md)

`
    public
                    __construct(string $profileName[, string|null $region = null ]) : mixed`

##### Parameters

$profileName
: string

Profile containing your console login session information

$region
: string\|null
= null

Region used for refresh requests. If not provided,
attempts will be made to resolve a region using
`AWS_REGION`, then the profile specified for `login`.

#### \_\_invoke()  [header link](class-aws-credentials-logincredentialprovider-method-invoke.md)

Returns a promise that resolves to AWS credentials

`
    public
                    __invoke() : PromiseInterface`

This method loads the cached token, refreshes it if necessary,
and returns AWS credentials sourced from the access token.

##### Tags  [header link](class-aws-credentials-logincredentialprovider-method-invoke-tags.md)

throws[CredentialsException](class-aws-exception-credentialsexception.md)

If re-authentication is required or credentials cannot be loaded

throws[SigninException](class-aws-signin-exception-signinexception.md)

If the token refresh fails with a SigninException

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
—

A promise that resolves to a Credentials object

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-credentials-logincredentialprovider-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-credentials-logincredentialprovider-method-construct.md)
  - [\_\_invoke()](class-aws-credentials-logincredentialprovider-method-invoke.md)

[Back To Top](class-aws-credentials-logincredentialprovider-top.md)
