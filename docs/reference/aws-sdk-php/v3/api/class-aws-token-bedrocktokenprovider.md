Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## BedrockTokenProvider     extends [TokenProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html)   in package    - [Aws](package-aws.md)

Token provider for Bedrock that sources bearer tokens from environment variables.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#toc-constants)

[BEARER\_AUTH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#constant_BEARER_AUTH)
= 'smithy.api#httpBearerAuth' [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#constant_ENV_PROFILE)
= 'AWS\_PROFILE' [TOKEN\_ENV\_KEY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#constant_TOKEN_ENV_KEY)
= 'bearer\_token\_bedrock'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#toc-methods)

[cache()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_cache)
: callable Wraps a token provider and saves provided token in an
instance of Aws\\CacheInterface. Forwards calls when no token found
in cache and updates cache with the results.[chain()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_chain)
: callable Creates an aggregate token provider that invokes the provided
variadic providers one after the other until a provider returns
a token.[createIfAvailable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#method_createIfAvailable)
: callable\|null Create a Bedrock token provider if the service is 'bedrock' and a token is available.[defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#method_defaultProvider)
: callable Create a default Bedrock token provider that checks for a bearer token
in the AWS\_BEARER\_TOKEN\_BEDROCK environment variable.[env()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#method_env)
: callable Token provider that creates a token from an environment variable.[fromToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_fromToken)
: callable Create a token provider function from a static token.[fromTokenValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#method_fromTokenValue)
: callable Create a token provider from a raw token value string.[memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_memoize)
: callable Wraps a token provider and caches a previously provided token.[sso()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_sso)
: [SsoTokenProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html)Token provider that creates a token from cached sso credentials

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#constants)

#### BEARER\_AUTH  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#constant_BEARER_AUTH)

`
    public
        mixed
    BEARER_AUTH
    = 'smithy.api#httpBearerAuth'
`

#### ENV\_PROFILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#constant_ENV_PROFILE)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### TOKEN\_ENV\_KEY  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#constant_TOKEN_ENV_KEY)

`
    public
        string
    TOKEN_ENV_KEY
    = 'bearer_token_bedrock'
`

used to resolve the AWS\_BEARER\_TOKEN\_BEDROCK env var

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#methods)

#### cache()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#method_cache)

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

#### chain()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#method_chain)

Creates an aggregate token provider that invokes the provided
variadic providers one after the other until a provider returns
a token.

`
    public
            static        chain() : callable`

##### Return values

callable

#### createIfAvailable()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#method_createIfAvailable)

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

#### defaultProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#method_defaultProvider)

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

#### env()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#method_env)

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

#### fromToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#method_fromToken)

Create a token provider function from a static token.

`
    public
            static        fromToken(TokenInterface $token) : callable`

##### Parameters

$token
: [TokenInterface](class-aws-token-tokeninterface.md)

##### Return values

callable

#### fromTokenValue()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html\#method_fromTokenValue)

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
: [TokenSource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenSource.html) \|null
= null

##### Return values

callable

#### memoize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#method_memoize)

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

#### sso()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#method_sso)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#method_sso\#tags)

see[SsoTokenProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html)

for $config details.

##### Return values

[SsoTokenProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#toc-methods)
- Constants
  - [BEARER\_AUTH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#constant_BEARER_AUTH)
  - [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#constant_ENV_PROFILE)
  - [TOKEN\_ENV\_KEY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#constant_TOKEN_ENV_KEY)
- Methods
  - [cache()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_cache)
  - [chain()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_chain)
  - [createIfAvailable()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#method_createIfAvailable)
  - [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#method_defaultProvider)
  - [env()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#method_env)
  - [fromToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_fromToken)
  - [fromTokenValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#method_fromTokenValue)
  - [memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_memoize)
  - [sso()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_sso)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html#top)
