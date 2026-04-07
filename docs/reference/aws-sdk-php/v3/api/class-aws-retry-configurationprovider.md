Menu

- [Aws](namespace-aws.md)
- [Retry](namespace-aws-retry.md)

## ConfigurationProvider     extends [AbstractConfigurationProvider](class-aws-abstractconfigurationprovider.md)   in package    - [Aws](package-aws.md)       implements  [ConfigurationProviderInterface](class-aws-configurationproviderinterface.md)

A configuration provider is a function that returns a promise that is
fulfilled with a {@see \\Aws\\Retry\\ConfigurationInterface}
or rejected with an {@see \\Aws\\Retry\\Exception\\ConfigurationException}.

`
use Aws\Sts\RegionalEndpoints\ConfigurationProvider;
$provider = ConfigurationProvider::defaultProvider();
// Returns a ConfigurationInterface or throws.
$config = $provider()->wait();
`

Configuration providers can be composed to create configuration using
conditional logic that can create different configurations in different
environments. You can compose multiple providers into a single provider using
ConfigurationProvider::chain. This function
accepts providers as variadic arguments and returns a new function that will
invoke each provider until a successful configuration is returned.

`
// First try an INI file at this location.
$a = ConfigurationProvider::ini(null, '/path/to/file.ini');
// Then try an INI file at this location.
$b = ConfigurationProvider::ini(null, '/path/to/other-file.ini');
// Then try loading from environment variables.
$c = ConfigurationProvider::env();
// Combine the three providers together.
$composed = ConfigurationProvider::chain($a, $b, $c);
// Returns a promise that is fulfilled with a configuration or throws.
$promise = $composed();
// Wait on the configuration to resolve.
$config = $promise->wait();
`

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#toc-interfaces)

[ConfigurationProviderInterface](class-aws-configurationproviderinterface.md)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#toc-constants)

[DEFAULT\_MAX\_ATTEMPTS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_DEFAULT_MAX_ATTEMPTS)
= 3 [DEFAULT\_MODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_DEFAULT_MODE)
= 'legacy' [ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider.md#constant_ENV_CONFIG_FILE)
= 'AWS\_CONFIG\_FILE' [ENV\_MAX\_ATTEMPTS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_ENV_MAX_ATTEMPTS)
= 'AWS\_MAX\_ATTEMPTS' [ENV\_MODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_ENV_MODE)
= 'AWS\_RETRY\_MODE' [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_ENV_PROFILE)
= 'AWS\_PROFILE' [INI\_MAX\_ATTEMPTS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_INI_MAX_ATTEMPTS)
= 'max\_attempts' [INI\_MODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_INI_MODE)
= 'retry\_mode'

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#toc-properties)

[$cacheKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#property_cacheKey)
: mixed

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#toc-methods)

[cache()](class-aws-abstractconfigurationprovider.md#method_cache)
: callable Wraps a config provider and saves provided configuration in an
instance of Aws\\CacheInterface. Forwards calls when no config found
in cache and updates cache with the results.[chain()](class-aws-abstractconfigurationprovider.md#method_chain)
: callable Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.[defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_defaultProvider)
: callable Create a default config provider that first checks for environment
variables, then checks for a specified profile in the environment-defined
config file location (env variable is 'AWS\_CONFIG\_FILE', file location
defaults to ~/.aws/config), then checks for the "default" profile in the
environment-defined config file location, and failing those uses a default
fallback set of configuration options.[env()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_env)
: callable Provider that creates config from environment variables.[fallback()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_fallback)
: callable Fallback config options when other sources are not set.[ini()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_ini)
: callable Config provider that creates config using a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified[memoize()](class-aws-abstractconfigurationprovider.md#method_memoize)
: callable Wraps a config provider and caches previously provided configuration.[unwrap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_unwrap)
: [ConfigurationInterface](class-aws-retry-configurationinterface.md)Unwraps a configuration object in whatever valid form it is in,
always returning a ConfigurationInterface object.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#constants)

#### DEFAULT\_MAX\_ATTEMPTS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#constant_DEFAULT_MAX_ATTEMPTS)

`
    public
        mixed
    DEFAULT_MAX_ATTEMPTS
    = 3
`

#### DEFAULT\_MODE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#constant_DEFAULT_MODE)

`
    public
        mixed
    DEFAULT_MODE
    = 'legacy'
`

#### ENV\_CONFIG\_FILE  [header link](class-aws-abstractconfigurationprovider.md\#constant_ENV_CONFIG_FILE)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_MAX\_ATTEMPTS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#constant_ENV_MAX_ATTEMPTS)

`
    public
        mixed
    ENV_MAX_ATTEMPTS
    = 'AWS_MAX_ATTEMPTS'
`

#### ENV\_MODE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#constant_ENV_MODE)

`
    public
        mixed
    ENV_MODE
    = 'AWS_RETRY_MODE'
`

#### ENV\_PROFILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#constant_ENV_PROFILE)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### INI\_MAX\_ATTEMPTS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#constant_INI_MAX_ATTEMPTS)

`
    public
        mixed
    INI_MAX_ATTEMPTS
    = 'max_attempts'
`

#### INI\_MODE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#constant_INI_MODE)

`
    public
        mixed
    INI_MODE
    = 'retry_mode'
`

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#properties)

#### $cacheKey  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#property_cacheKey)

`
    public
    static    mixed
    $cacheKey
     = 'aws_retries_config'`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#methods)

#### cache()  [header link](class-aws-abstractconfigurationprovider.md\#method_cache)

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

#### chain()  [header link](class-aws-abstractconfigurationprovider.md\#method_chain)

Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.

`
    public
            static        chain() : callable`

##### Return values

callable

#### defaultProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#method_defaultProvider)

Create a default config provider that first checks for environment
variables, then checks for a specified profile in the environment-defined
config file location (env variable is 'AWS\_CONFIG\_FILE', file location
defaults to ~/.aws/config), then checks for the "default" profile in the
environment-defined config file location, and failing those uses a default
fallback set of configuration options.

`
    public
            static        defaultProvider([array<string|int, mixed> $config = [] ]) : callable`

This provider is automatically wrapped in a memoize function that caches
previously provided config options.

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Return values

callable

#### env()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#method_env)

Provider that creates config from environment variables.

`
    public
            static        env() : callable`

##### Return values

callable

#### fallback()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#method_fallback)

Fallback config options when other sources are not set.

`
    public
            static        fallback() : callable`

##### Return values

callable

#### ini()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#method_ini)

Config provider that creates config using a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified

`
    public
            static        ini([string|null $profile = null ][, string|null $filename = null ]) : callable`

##### Parameters

$profile
: string\|null
= null

Profile to use. If not specified will use
the "default" profile.

$filename
: string\|null
= null

If provided, uses a custom filename rather
than looking in the default directory.

##### Return values

callable

#### memoize()  [header link](class-aws-abstractconfigurationprovider.md\#method_memoize)

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

#### unwrap()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#method_unwrap)

Unwraps a configuration object in whatever valid form it is in,
always returning a ConfigurationInterface object.

`
    public
            static        unwrap(mixed $config) : ConfigurationInterface`

##### Parameters

$config
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html\#method_unwrap\#tags)

throwsInvalidArgumentException

##### Return values

[ConfigurationInterface](class-aws-retry-configurationinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#toc-constants)
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#toc-methods)
- Constants
  - [DEFAULT\_MAX\_ATTEMPTS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_DEFAULT_MAX_ATTEMPTS)
  - [DEFAULT\_MODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_DEFAULT_MODE)
  - [ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider.md#constant_ENV_CONFIG_FILE)
  - [ENV\_MAX\_ATTEMPTS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_ENV_MAX_ATTEMPTS)
  - [ENV\_MODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_ENV_MODE)
  - [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_ENV_PROFILE)
  - [INI\_MAX\_ATTEMPTS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_INI_MAX_ATTEMPTS)
  - [INI\_MODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#constant_INI_MODE)
- Properties
  - [$cacheKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#property_cacheKey)
- Methods
  - [cache()](class-aws-abstractconfigurationprovider.md#method_cache)
  - [chain()](class-aws-abstractconfigurationprovider.md#method_chain)
  - [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_defaultProvider)
  - [env()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_env)
  - [fallback()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_fallback)
  - [ini()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_ini)
  - [memoize()](class-aws-abstractconfigurationprovider.md#method_memoize)
  - [unwrap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#method_unwrap)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationProvider.html#top)
