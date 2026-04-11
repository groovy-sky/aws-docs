Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## BedrockTokenProvider     extends [TokenProvider](class-aws-token-tokenprovider.md)   in package    - [Aws](package-aws.md)

Token provider for Bedrock that sources bearer tokens from environment variables.

### Table of Contents  [header link](class-aws-token-bedrocktokenprovider-toc.md)

#### Constants  [header link](class-aws-token-bedrocktokenprovider-toc-constants.md)

[BEARER\_AUTH](class-aws-token-bedrocktokenprovider-constant-bearer-auth.md)
= 'smithy.api#httpBearerAuth' [ENV\_PROFILE](class-aws-token-tokenprovider-constant-env-profile.md)
= 'AWS\_PROFILE' [TOKEN\_ENV\_KEY](class-aws-token-bedrocktokenprovider-constant-token-env-key.md)
= 'bearer\_token\_bedrock'

#### Methods  [header link](class-aws-token-bedrocktokenprovider-toc-methods.md)

[cache()](class-aws-token-tokenprovider-method-cache.md)
: callable Wraps a token provider and saves provided token in an
instance of Aws\\CacheInterface. Forwards calls when no token found
in cache and updates cache with the results.[chain()](class-aws-token-tokenprovider-method-chain.md)
: callable Creates an aggregate token provider that invokes the provided
variadic providers one after the other until a provider returns
a token.[createIfAvailable()](class-aws-token-bedrocktokenprovider-method-createifavailable.md)
: callable\|null Create a Bedrock token provider if the service is 'bedrock' and a token is available.[defaultProvider()](class-aws-token-bedrocktokenprovider-method-defaultprovider.md)
: callable Create a default Bedrock token provider that checks for a bearer token
in the AWS\_BEARER\_TOKEN\_BEDROCK environment variable.[env()](class-aws-token-bedrocktokenprovider-method-env.md)
: callable Token provider that creates a token from an environment variable.[fromToken()](class-aws-token-tokenprovider-method-fromtoken.md)
: callable Create a token provider function from a static token.[fromTokenValue()](class-aws-token-bedrocktokenprovider-method-fromtokenvalue.md)
: callable Create a token provider from a raw token value string.[memoize()](class-aws-token-tokenprovider-method-memoize.md)
: callable Wraps a token provider and caches a previously provided token.[sso()](class-aws-token-tokenprovider-method-sso.md)
: [SsoTokenProvider](class-aws-token-ssotokenprovider.md)Token provider that creates a token from cached sso credentials

### Constants  [header link](class-aws-token-bedrocktokenprovider-constants.md)

#### BEARER\_AUTH  [header link](class-aws-token-bedrocktokenprovider-constant-bearer-auth.md)

`
    public
        mixed
    BEARER_AUTH
    = 'smithy.api#httpBearerAuth'
`

#### ENV\_PROFILE  [header link](class-aws-token-tokenprovider-constant-env-profile.md)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### TOKEN\_ENV\_KEY  [header link](class-aws-token-bedrocktokenprovider-constant-token-env-key.md)

`
    public
        string
    TOKEN_ENV_KEY
    = 'bearer_token_bedrock'
`

used to resolve the AWS\_BEARER\_TOKEN\_BEDROCK env var

### Methods  [header link](class-aws-token-bedrocktokenprovider-methods.md)

#### cache()  [header link](class-aws-token-tokenprovider-method-cache.md)

Wraps a token provider and saves provided token in an
instance of Aws\\CacheInterface. Forwards calls when no token found
in cache and updates cache with the results.

`
    public
            static        cache(callable $provider, CacheInterface $cache[, string|null $cacheKey = null ]) : callable`

##### Parameters

$provider
: callable

Token provider function to wrap

$cache
: [CacheInterface](class-aws-cacheinterface.md)

Cache to store the token

$cacheKey
: string\|null
= null

(optional) Cache key to use

##### Return values

callable

#### chain()  [header link](class-aws-token-tokenprovider-method-chain.md)

Creates an aggregate token provider that invokes the provided
variadic providers one after the other until a provider returns
a token.

`
    public
            static        chain() : callable`

##### Return values

callable

#### createIfAvailable()  [header link](class-aws-token-bedrocktokenprovider-method-createifavailable.md)

Create a Bedrock token provider if the service is 'bedrock' and a token is available.

`
    public
            static        createIfAvailable(array<string|int, mixed> &$args) : callable|null`

Sets auth scheme preference to `bearer` auth.

##### Parameters

$args
: array<string\|int, mixed>

Configuration arguments containing 'config' array

##### Return values

callable\|null
—

Returns a token provider if conditions are met, null otherwise

#### defaultProvider()  [header link](class-aws-token-bedrocktokenprovider-method-defaultprovider.md)

Create a default Bedrock token provider that checks for a bearer token
in the AWS\_BEARER\_TOKEN\_BEDROCK environment variable.

`
    public
            static        defaultProvider([array<string|int, mixed> $config = [] ]) : callable`

This provider is automatically wrapped in a memoize function that caches
previously provided tokens.

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Optional array of token provider options.

##### Return values

callable

#### env()  [header link](class-aws-token-bedrocktokenprovider-method-env.md)

Token provider that creates a token from an environment variable.

`
    public
            static        env(string $configKey) : callable`

##### Parameters

$configKey
: string

The configuration key that will be transformed
to an environment variable name by ConfigurationResolver

##### Return values

callable

#### fromToken()  [header link](class-aws-token-tokenprovider-method-fromtoken.md)

Create a token provider function from a static token.

`
    public
            static        fromToken(TokenInterface $token) : callable`

##### Parameters

$token
: [TokenInterface](class-aws-token-tokeninterface.md)

##### Return values

callable

#### fromTokenValue()  [header link](class-aws-token-bedrocktokenprovider-method-fromtokenvalue.md)

Create a token provider from a raw token value string.

`
    public
            static        fromTokenValue(string $tokenValue[, TokenSource|null $source = null ]) : callable`

Bedrock bearer tokens sourced from env do not have an expiration

##### Parameters

$tokenValue
: string

The bearer token value

$source
: [TokenSource](class-aws-token-tokensource.md) \|null
= null

##### Return values

callable

#### memoize()  [header link](class-aws-token-tokenprovider-method-memoize.md)

Wraps a token provider and caches a previously provided token.

`
    public
            static        memoize(callable $provider) : callable`

Ensures that cached tokens are refreshed when they expire.

##### Parameters

$provider
: callable

Token provider function to wrap.

##### Return values

callable

#### sso()  [header link](class-aws-token-tokenprovider-method-sso.md)

Token provider that creates a token from cached sso credentials

`
    public
            static        sso(string $profileName, string $filename[, array<string|int, mixed> $config = [] ]) : SsoTokenProvider`

##### Parameters

$profileName
: string

the name of the ini profile name

$filename
: string

the location of the ini file

$config
: array<string\|int, mixed>
= \[\]

configuration options

##### Tags  [header link](class-aws-token-tokenprovider-method-sso-tags.md)

see[SsoTokenProvider](class-aws-token-ssotokenprovider.md)

for $config details.

##### Return values

[SsoTokenProvider](class-aws-token-ssotokenprovider.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-token-bedrocktokenprovider-toc-constants.md)
  - [Methods](class-aws-token-bedrocktokenprovider-toc-methods.md)
- Constants
  - [BEARER\_AUTH](class-aws-token-bedrocktokenprovider-constant-bearer-auth.md)
  - [ENV\_PROFILE](class-aws-token-tokenprovider-constant-env-profile.md)
  - [TOKEN\_ENV\_KEY](class-aws-token-bedrocktokenprovider-constant-token-env-key.md)
- Methods
  - [cache()](class-aws-token-tokenprovider-method-cache.md)
  - [chain()](class-aws-token-tokenprovider-method-chain.md)
  - [createIfAvailable()](class-aws-token-bedrocktokenprovider-method-createifavailable.md)
  - [defaultProvider()](class-aws-token-bedrocktokenprovider-method-defaultprovider.md)
  - [env()](class-aws-token-bedrocktokenprovider-method-env.md)
  - [fromToken()](class-aws-token-tokenprovider-method-fromtoken.md)
  - [fromTokenValue()](class-aws-token-bedrocktokenprovider-method-fromtokenvalue.md)
  - [memoize()](class-aws-token-tokenprovider-method-memoize.md)
  - [sso()](class-aws-token-tokenprovider-method-sso.md)

[Back To Top](class-aws-token-bedrocktokenprovider-top.md)
