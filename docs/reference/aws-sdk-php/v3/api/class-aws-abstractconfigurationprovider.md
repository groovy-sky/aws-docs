Menu

- [Aws](namespace-aws.md)

## AbstractConfigurationProvider        in package    - [Aws](package-aws.md)

AbstractYes

A configuration provider is a function that returns a promise that is
fulfilled with a configuration object. This class provides base functionality
usable by specific configuration provider implementations

### Table of Contents  [header link](class-aws-abstractconfigurationprovider-toc.md)

#### Constants  [header link](class-aws-abstractconfigurationprovider-toc-constants.md)

[ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider-constant-env-config-file.md)
= 'AWS\_CONFIG\_FILE' [ENV\_PROFILE](class-aws-abstractconfigurationprovider-constant-env-profile.md)
= 'AWS\_PROFILE'

#### Properties  [header link](class-aws-abstractconfigurationprovider-toc-properties.md)

[$cacheKey](class-aws-abstractconfigurationprovider-property-cachekey.md)
: mixed

#### Methods  [header link](class-aws-abstractconfigurationprovider-toc-methods.md)

[cache()](class-aws-abstractconfigurationprovider-method-cache.md)
: callable Wraps a config provider and saves provided configuration in an
instance of Aws\\CacheInterface. Forwards calls when no config found
in cache and updates cache with the results.[chain()](class-aws-abstractconfigurationprovider-method-chain.md)
: callable Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.[memoize()](class-aws-abstractconfigurationprovider-method-memoize.md)
: callable Wraps a config provider and caches previously provided configuration.

### Constants  [header link](class-aws-abstractconfigurationprovider-constants.md)

#### ENV\_CONFIG\_FILE  [header link](class-aws-abstractconfigurationprovider-constant-env-config-file.md)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_PROFILE  [header link](class-aws-abstractconfigurationprovider-constant-env-profile.md)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

### Properties  [header link](class-aws-abstractconfigurationprovider-properties.md)

#### $cacheKey  [header link](class-aws-abstractconfigurationprovider-property-cachekey.md)

`
    public
    static    mixed
    $cacheKey
    `

### Methods  [header link](class-aws-abstractconfigurationprovider-methods.md)

#### cache()  [header link](class-aws-abstractconfigurationprovider-method-cache.md)

Wraps a config provider and saves provided configuration in an
instance of Aws\\CacheInterface. Forwards calls when no config found
in cache and updates cache with the results.

`
    public
            static        cache(callable $provider, CacheInterface $cache[, string|null $cacheKey = null ]) : callable`

##### Parameters

$provider
: callable

Configuration provider function to wrap

$cache
: [CacheInterface](class-aws-cacheinterface.md)

Cache to store configuration

$cacheKey
: string\|null
= null

(optional) Cache key to use

##### Return values

callable

#### chain()  [header link](class-aws-abstractconfigurationprovider-method-chain.md)

Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.

`
    public
            static        chain() : callable`

##### Return values

callable

#### memoize()  [header link](class-aws-abstractconfigurationprovider-method-memoize.md)

Wraps a config provider and caches previously provided configuration.

`
    public
            static        memoize(callable $provider) : callable`

##### Parameters

$provider
: callable

Config provider function to wrap.

##### Return values

callable
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-abstractconfigurationprovider-toc-constants.md)
  - [Properties](class-aws-abstractconfigurationprovider-toc-properties.md)
  - [Methods](class-aws-abstractconfigurationprovider-toc-methods.md)
- Constants
  - [ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider-constant-env-config-file.md)
  - [ENV\_PROFILE](class-aws-abstractconfigurationprovider-constant-env-profile.md)
- Properties
  - [$cacheKey](class-aws-abstractconfigurationprovider-property-cachekey.md)
- Methods
  - [cache()](class-aws-abstractconfigurationprovider-method-cache.md)
  - [chain()](class-aws-abstractconfigurationprovider-method-chain.md)
  - [memoize()](class-aws-abstractconfigurationprovider-method-memoize.md)

[Back To Top](class-aws-abstractconfigurationprovider-top.md)
