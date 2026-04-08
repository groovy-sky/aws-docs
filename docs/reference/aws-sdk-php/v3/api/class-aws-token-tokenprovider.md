Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## TokenProvider        in package    - [Aws](package-aws.md)       Uses  [ParsesIniTrait](class-aws-token-parsesinitrait.md)

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

### Table of Contents  [header link](class-aws-token-tokenprovider-toc.md)

#### Constants  [header link](class-aws-token-tokenprovider-toc-constants.md)

[ENV\_PROFILE](class-aws-token-tokenprovider-constant-env-profile.md)
= 'AWS\_PROFILE'

#### Methods  [header link](class-aws-token-tokenprovider-toc-methods.md)

[cache()](class-aws-token-tokenprovider-method-cache.md)
: callable Wraps a token provider and saves provided token in an
instance of Aws\\CacheInterface. Forwards calls when no token found
in cache and updates cache with the results.[chain()](class-aws-token-tokenprovider-method-chain.md)
: callable Creates an aggregate token provider that invokes the provided
variadic providers one after the other until a provider returns
a token.[defaultProvider()](class-aws-token-tokenprovider-method-defaultprovider.md)
: callable Create a default token provider tha checks for cached a SSO token from
the CLI[fromToken()](class-aws-token-tokenprovider-method-fromtoken.md)
: callable Create a token provider function from a static token.[memoize()](class-aws-token-tokenprovider-method-memoize.md)
: callable Wraps a token provider and caches a previously provided token.[sso()](class-aws-token-tokenprovider-method-sso.md)
: [SsoTokenProvider](class-aws-token-ssotokenprovider.md)Token provider that creates a token from cached sso credentials

### Constants  [header link](class-aws-token-tokenprovider-constants.md)

#### ENV\_PROFILE  [header link](class-aws-token-tokenprovider-constant-env-profile.md)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

### Methods  [header link](class-aws-token-tokenprovider-methods.md)

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

#### defaultProvider()  [header link](class-aws-token-tokenprovider-method-defaultprovider.md)

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
  - [Constants](class-aws-token-tokenprovider-toc-constants.md)
  - [Methods](class-aws-token-tokenprovider-toc-methods.md)
- Constants
  - [ENV\_PROFILE](class-aws-token-tokenprovider-constant-env-profile.md)
- Methods
  - [cache()](class-aws-token-tokenprovider-method-cache.md)
  - [chain()](class-aws-token-tokenprovider-method-chain.md)
  - [defaultProvider()](class-aws-token-tokenprovider-method-defaultprovider.md)
  - [fromToken()](class-aws-token-tokenprovider-method-fromtoken.md)
  - [memoize()](class-aws-token-tokenprovider-method-memoize.md)
  - [sso()](class-aws-token-tokenprovider-method-sso.md)

[Back To Top](class-aws-token-tokenprovider-top.md)
