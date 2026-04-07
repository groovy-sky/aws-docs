Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## TokenProvider        in package    - [Aws](package-aws.md)       Uses  [ParsesIniTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.ParsesIniTrait.html)

Token providers are functions that accept no arguments and return a
promise that is fulfilled with an {@see \\Aws\\Token\\TokenInterface}
or rejected with an {@see \\Aws\\Exception\\TokenException}.

`
use Aws\Token\TokenProvider;
$provider = TokenProvider::defaultProvider();
// Returns a TokenInterface or throws.
$token = $provider()->wait();
`

Token providers can be composed to create a token using conditional
logic that can create different tokens in different environments. You
can compose multiple providers into a single provider using
TokenProvider::chain. This function accepts
providers as variadic arguments and returns a new function that will invoke
each provider until a token is successfully returned.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#toc-constants)

[ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#constant_ENV_PROFILE)
= 'AWS\_PROFILE'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#toc-methods)

[cache()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_cache)
: callable Wraps a token provider and saves provided token in an
instance of Aws\\CacheInterface. Forwards calls when no token found
in cache and updates cache with the results.[chain()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_chain)
: callable Creates an aggregate token provider that invokes the provided
variadic providers one after the other until a provider returns
a token.[defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_defaultProvider)
: callable Create a default token provider tha checks for cached a SSO token from
the CLI[fromToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_fromToken)
: callable Create a token provider function from a static token.[memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_memoize)
: callable Wraps a token provider and caches a previously provided token.[sso()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_sso)
: [SsoTokenProvider](class-aws-token-ssotokenprovider.md)Token provider that creates a token from cached sso credentials

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#constants)

#### ENV\_PROFILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#constant_ENV_PROFILE)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#methods)

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

#### defaultProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html\#method_defaultProvider)

Create a default token provider tha checks for cached a SSO token from
the CLI

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

see[SsoTokenProvider](class-aws-token-ssotokenprovider.md)

for $config details.

##### Return values

[SsoTokenProvider](class-aws-token-ssotokenprovider.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#toc-methods)
- Constants
  - [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#constant_ENV_PROFILE)
- Methods
  - [cache()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_cache)
  - [chain()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_chain)
  - [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_defaultProvider)
  - [fromToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_fromToken)
  - [memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_memoize)
  - [sso()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#method_sso)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html#top)
