Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## SsoTokenProvider        in package    - [Aws](package-aws.md)       implements  [RefreshableTokenProviderInterface](class-aws-token-refreshabletokenproviderinterface.md)  Uses  [ParsesIniTrait](class-aws-token-parsesinitrait.md)

Token that comes from the SSO provider

### Table of Contents  [header link](class-aws-token-ssotokenprovider-toc.md)

#### Interfaces  [header link](class-aws-token-ssotokenprovider-toc-interfaces.md)

[RefreshableTokenProviderInterface](class-aws-token-refreshabletokenproviderinterface.md)Provides access to an AWS token used for accessing AWS services

#### Constants  [header link](class-aws-token-ssotokenprovider-toc-constants.md)

[ENV\_PROFILE](class-aws-token-ssotokenprovider-constant-env-profile.md)
= 'AWS\_PROFILE' [REFRESH\_ATTEMPT\_WINDOW\_IN\_SECS](class-aws-token-ssotokenprovider-constant-refresh-attempt-window-in-secs.md)
= 30 [REFRESH\_WINDOW\_IN\_SECS](class-aws-token-ssotokenprovider-constant-refresh-window-in-secs.md)
= 300

#### Methods  [header link](class-aws-token-ssotokenprovider-toc-methods.md)

[\_\_construct()](class-aws-token-ssotokenprovider-method-construct.md)
: mixed Constructs a new SsoTokenProvider object, which will fetch a token from an authenticated SSO profile[\_\_invoke()](class-aws-token-ssotokenprovider-method-invoke.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Loads cached sso credentials.[getTokenData()](class-aws-token-ssotokenprovider-method-gettokendata.md)
: array<string\|int, mixed> [getTokenLocation()](class-aws-token-ssotokenprovider-method-gettokenlocation.md)
: string [refresh()](class-aws-token-ssotokenprovider-method-refresh.md)
: array<string\|int, mixed> This method attempt to refresh when possible.[shouldAttemptRefresh()](class-aws-token-ssotokenprovider-method-shouldattemptrefresh.md)
: bool This method checks for whether a token refresh should happen.

### Constants  [header link](class-aws-token-ssotokenprovider-constants.md)

#### ENV\_PROFILE  [header link](class-aws-token-ssotokenprovider-constant-env-profile.md)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### REFRESH\_ATTEMPT\_WINDOW\_IN\_SECS  [header link](class-aws-token-ssotokenprovider-constant-refresh-attempt-window-in-secs.md)

`
    public
        mixed
    REFRESH_ATTEMPT_WINDOW_IN_SECS
    = 30
`

#### REFRESH\_WINDOW\_IN\_SECS  [header link](class-aws-token-ssotokenprovider-constant-refresh-window-in-secs.md)

`
    public
        mixed
    REFRESH_WINDOW_IN_SECS
    = 300
`

### Methods  [header link](class-aws-token-ssotokenprovider-methods.md)

#### \_\_construct()  [header link](class-aws-token-ssotokenprovider-method-construct.md)

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

#### \_\_invoke()  [header link](class-aws-token-ssotokenprovider-method-invoke.md)

Loads cached sso credentials.

`
    public
                    __invoke() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getTokenData()  [header link](class-aws-token-ssotokenprovider-method-gettokendata.md)

`
    public
                    getTokenData( $tokenLocation) : array<string|int, mixed>`

##### Parameters

$tokenLocation
:

##### Return values

array<string\|int, mixed>

#### getTokenLocation()  [header link](class-aws-token-ssotokenprovider-method-gettokenlocation.md)

`
    public
            static        getTokenLocation( $sso_session) : string`

##### Parameters

$sso\_session
:

##### Return values

string

#### refresh()  [header link](class-aws-token-ssotokenprovider-method-refresh.md)

This method attempt to refresh when possible.

`
    public
                    refresh() : array<string|int, mixed>`

If a refresh is not possible then it just returns
the current token data as it is.

##### Tags  [header link](class-aws-token-ssotokenprovider-method-refresh-tags.md)

throws[TokenException](class-aws-exception-tokenexception.md)

##### Return values

array<string\|int, mixed>

#### shouldAttemptRefresh()  [header link](class-aws-token-ssotokenprovider-method-shouldattemptrefresh.md)

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
  - [Constants](class-aws-token-ssotokenprovider-toc-constants.md)
  - [Methods](class-aws-token-ssotokenprovider-toc-methods.md)
- Constants
  - [ENV\_PROFILE](class-aws-token-ssotokenprovider-constant-env-profile.md)
  - [REFRESH\_ATTEMPT\_WINDOW\_IN\_SECS](class-aws-token-ssotokenprovider-constant-refresh-attempt-window-in-secs.md)
  - [REFRESH\_WINDOW\_IN\_SECS](class-aws-token-ssotokenprovider-constant-refresh-window-in-secs.md)
- Methods
  - [\_\_construct()](class-aws-token-ssotokenprovider-method-construct.md)
  - [\_\_invoke()](class-aws-token-ssotokenprovider-method-invoke.md)
  - [getTokenData()](class-aws-token-ssotokenprovider-method-gettokendata.md)
  - [getTokenLocation()](class-aws-token-ssotokenprovider-method-gettokenlocation.md)
  - [refresh()](class-aws-token-ssotokenprovider-method-refresh.md)
  - [shouldAttemptRefresh()](class-aws-token-ssotokenprovider-method-shouldattemptrefresh.md)

[Back To Top](class-aws-token-ssotokenprovider-top.md)
