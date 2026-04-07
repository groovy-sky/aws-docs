Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## LoginCredentialProvider        in package    - [Aws](package-aws.md)

FinalYes

Credential provider for login using console credentials

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html#method___construct)
: mixed [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html#method___invoke)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that resolves to AWS credentials

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html\#method___construct)

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

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html\#method___invoke)

Returns a promise that resolves to AWS credentials

`
    public
                    __invoke() : PromiseInterface`

This method loads the cached token, refreshes it if necessary,
and returns AWS credentials sourced from the access token.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html\#method___invoke\#tags)

throws[CredentialsException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.CredentialsException.html)

If re-authentication is required or credentials cannot be loaded

throws[SigninException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signin.Exception.SigninException.html)

If the token refresh fails with a SigninException

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
—

A promise that resolves to a Credentials object

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html#method___invoke)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html#top)
