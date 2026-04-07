Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## SsoTokenProvider        in package    - [Aws](package-aws.md)       implements  [RefreshableTokenProviderInterface](class-aws-token-refreshabletokenproviderinterface.md)  Uses  [ParsesIniTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.ParsesIniTrait.html)

Token that comes from the SSO provider

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#toc-interfaces)

[RefreshableTokenProviderInterface](class-aws-token-refreshabletokenproviderinterface.md)Provides access to an AWS token used for accessing AWS services

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#toc-constants)

[ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#constant_ENV_PROFILE)
= 'AWS\_PROFILE' [REFRESH\_ATTEMPT\_WINDOW\_IN\_SECS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#constant_REFRESH_ATTEMPT_WINDOW_IN_SECS)
= 30 [REFRESH\_WINDOW\_IN\_SECS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#constant_REFRESH_WINDOW_IN_SECS)
= 300

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method___construct)
: mixed Constructs a new SsoTokenProvider object, which will fetch a token from an authenticated SSO profile[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method___invoke)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Loads cached sso credentials.[getTokenData()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method_getTokenData)
: array<string\|int, mixed> [getTokenLocation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method_getTokenLocation)
: string [refresh()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method_refresh)
: array<string\|int, mixed> This method attempt to refresh when possible.[shouldAttemptRefresh()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method_shouldAttemptRefresh)
: bool This method checks for whether a token refresh should happen.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#constants)

#### ENV\_PROFILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#constant_ENV_PROFILE)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### REFRESH\_ATTEMPT\_WINDOW\_IN\_SECS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#constant_REFRESH_ATTEMPT_WINDOW_IN_SECS)

`
    public
        mixed
    REFRESH_ATTEMPT_WINDOW_IN_SECS
    = 30
`

#### REFRESH\_WINDOW\_IN\_SECS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#constant_REFRESH_WINDOW_IN_SECS)

`
    public
        mixed
    REFRESH_WINDOW_IN_SECS
    = 300
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#method___construct)

Constructs a new SsoTokenProvider object, which will fetch a token from an authenticated SSO profile

`
    public
                    __construct(string $profileName[, string|null $configFilePath = null ][, SSOOIDCClient|null $ssoOidcClient = null ]) : mixed`

##### Parameters

$profileName
: string

The name of the profile that contains the sso\_session key

$configFilePath
: string\|null
= null

Name of the config file to sso profile from

$ssoOidcClient
: [SSOOIDCClient](class-aws-ssooidc-ssooidcclient.md) \|null
= null

The sso client for generating a new token

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#method___invoke)

Loads cached sso credentials.

`
    public
                    __invoke() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getTokenData()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#method_getTokenData)

`
    public
                    getTokenData( $tokenLocation) : array<string|int, mixed>`

##### Parameters

$tokenLocation
:

##### Return values

array<string\|int, mixed>

#### getTokenLocation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#method_getTokenLocation)

`
    public
            static        getTokenLocation( $sso_session) : string`

##### Parameters

$sso\_session
:

##### Return values

string

#### refresh()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#method_refresh)

This method attempt to refresh when possible.

`
    public
                    refresh() : array<string|int, mixed>`

If a refresh is not possible then it just returns
the current token data as it is.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#method_refresh\#tags)

throws[TokenException](class-aws-exception-tokenexception.md)

##### Return values

array<string\|int, mixed>

#### shouldAttemptRefresh()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html\#method_shouldAttemptRefresh)

This method checks for whether a token refresh should happen.

`
    public
                    shouldAttemptRefresh() : bool`

It will return true just if more than 30 seconds has happened
since last refresh, and if the expiration is within a 5-minutes
window from the current time.

##### Return values

bool
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#toc-methods)
- Constants
  - [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#constant_ENV_PROFILE)
  - [REFRESH\_ATTEMPT\_WINDOW\_IN\_SECS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#constant_REFRESH_ATTEMPT_WINDOW_IN_SECS)
  - [REFRESH\_WINDOW\_IN\_SECS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#constant_REFRESH_WINDOW_IN_SECS)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method___invoke)
  - [getTokenData()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method_getTokenData)
  - [getTokenLocation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method_getTokenLocation)
  - [refresh()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method_refresh)
  - [shouldAttemptRefresh()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#method_shouldAttemptRefresh)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html#top)
