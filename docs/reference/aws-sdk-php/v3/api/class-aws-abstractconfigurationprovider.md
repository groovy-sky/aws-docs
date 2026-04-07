Menu

- [Aws](namespace-aws.md)

## AbstractConfigurationProvider        in package    - [Aws](package-aws.md)

AbstractYes

A configuration provider is a function that returns a promise that is
fulfilled with a configuration object. This class provides base functionality
usable by specific configuration provider implementations

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#toc-constants)

[ENV\_CONFIG\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#constant_ENV_CONFIG_FILE)
= 'AWS\_CONFIG\_FILE' [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#constant_ENV_PROFILE)
= 'AWS\_PROFILE'

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#toc-properties)

[$cacheKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#property_cacheKey)
: mixed

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#toc-methods)

[cache()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#method_cache)
: callable Wraps a config provider and saves provided configuration in an
instance of Aws\\CacheInterface. Forwards calls when no config found
in cache and updates cache with the results.[chain()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#method_chain)
: callable Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.[memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#method_memoize)
: callable Wraps a config provider and caches previously provided configuration.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#constants)

#### ENV\_CONFIG\_FILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#constant_ENV_CONFIG_FILE)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_PROFILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#constant_ENV_PROFILE)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#properties)

#### $cacheKey  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#property_cacheKey)

`
    public
    static    mixed
    $cacheKey
    `

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#methods)

#### cache()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#method_cache)

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

#### chain()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#method_chain)

Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.

`
    public
            static        chain() : callable`

##### Return values

callable

#### memoize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html\#method_memoize)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#toc-constants)
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#toc-methods)
- Constants
  - [ENV\_CONFIG\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#constant_ENV_CONFIG_FILE)
  - [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#constant_ENV_PROFILE)
- Properties
  - [$cacheKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#property_cacheKey)
- Methods
  - [cache()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#method_cache)
  - [chain()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#method_chain)
  - [memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#method_memoize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AbstractConfigurationProvider.html#top)
