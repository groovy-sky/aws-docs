Menu

- [Aws](namespace-aws.md)
- [Sts](namespace-aws-sts.md)
- [RegionalEndpoints](namespace-aws-sts-regionalendpoints.md)

## ConfigurationProvider     extends [AbstractConfigurationProvider](class-aws-abstractconfigurationprovider.md)   in package    - [Aws](package-aws.md)       implements  [ConfigurationProviderInterface](class-aws-configurationproviderinterface.md)

A configuration provider is a function that returns a promise that is
fulfilled with a {@see \\Aws\\Sts\\RegionalEndpoints\\ConfigurationInterface}
or rejected with an {@see \\Aws\\Sts\\RegionalEndpoints\\Exception\\ConfigurationException}.

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

### Table of Contents  [header link](class-aws-sts-regionalendpoints-configurationprovider-toc.md)

#### Interfaces  [header link](class-aws-sts-regionalendpoints-configurationprovider-toc-interfaces.md)

[ConfigurationProviderInterface](class-aws-configurationproviderinterface.md)

#### Constants  [header link](class-aws-sts-regionalendpoints-configurationprovider-toc-constants.md)

[DEFAULT\_ENDPOINTS\_TYPE](class-aws-sts-regionalendpoints-configurationprovider-constant-default-endpoints-type.md)
= 'regional' [ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider.md#constant_ENV_CONFIG_FILE)
= 'AWS\_CONFIG\_FILE' [ENV\_ENDPOINTS\_TYPE](class-aws-sts-regionalendpoints-configurationprovider-constant-env-endpoints-type.md)
= 'AWS\_STS\_REGIONAL\_ENDPOINTS' [ENV\_PROFILE](class-aws-sts-regionalendpoints-configurationprovider-constant-env-profile.md)
= 'AWS\_PROFILE' [INI\_ENDPOINTS\_TYPE](class-aws-sts-regionalendpoints-configurationprovider-constant-ini-endpoints-type.md)
= 'sts\_regional\_endpoints'

#### Properties  [header link](class-aws-sts-regionalendpoints-configurationprovider-toc-properties.md)

[$cacheKey](class-aws-sts-regionalendpoints-configurationprovider-property-cachekey.md)
: mixed

#### Methods  [header link](class-aws-sts-regionalendpoints-configurationprovider-toc-methods.md)

[cache()](class-aws-abstractconfigurationprovider.md#method_cache)
: callable Wraps a config provider and saves provided configuration in an
instance of Aws\\CacheInterface. Forwards calls when no config found
in cache and updates cache with the results.[chain()](class-aws-abstractconfigurationprovider.md#method_chain)
: callable Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.[defaultProvider()](class-aws-sts-regionalendpoints-configurationprovider-method-defaultprovider.md)
: callable Create a default config provider that first checks for environment
variables, then checks for a specified profile in the environment-defined
config file location (env variable is 'AWS\_CONFIG\_FILE', file location
defaults to ~/.aws/config), then checks for the "default" profile in the
environment-defined config file location, and failing those uses a default
fallback set of configuration options.[env()](class-aws-sts-regionalendpoints-configurationprovider-method-env.md)
: callable Provider that creates config from environment variables.[fallback()](class-aws-sts-regionalendpoints-configurationprovider-method-fallback.md)
: callable Fallback config options when other sources are not set.[ini()](class-aws-sts-regionalendpoints-configurationprovider-method-ini.md)
: callable Config provider that creates config using a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified[memoize()](class-aws-abstractconfigurationprovider.md#method_memoize)
: callable Wraps a config provider and caches previously provided configuration.[unwrap()](class-aws-sts-regionalendpoints-configurationprovider-method-unwrap.md)
: [ConfigurationInterface](class-aws-sts-regionalendpoints-configurationinterface.md)Unwraps a configuration object in whatever valid form it is in,
always returning a ConfigurationInterface object.

### Constants  [header link](class-aws-sts-regionalendpoints-configurationprovider-constants.md)

#### DEFAULT\_ENDPOINTS\_TYPE  [header link](class-aws-sts-regionalendpoints-configurationprovider-constant-default-endpoints-type.md)

`
    public
        mixed
    DEFAULT_ENDPOINTS_TYPE
    = 'regional'
`

#### ENV\_CONFIG\_FILE  [header link](class-aws-abstractconfigurationprovider.md\#constant_ENV_CONFIG_FILE)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_ENDPOINTS\_TYPE  [header link](class-aws-sts-regionalendpoints-configurationprovider-constant-env-endpoints-type.md)

`
    public
        mixed
    ENV_ENDPOINTS_TYPE
    = 'AWS_STS_REGIONAL_ENDPOINTS'
`

#### ENV\_PROFILE  [header link](class-aws-sts-regionalendpoints-configurationprovider-constant-env-profile.md)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### INI\_ENDPOINTS\_TYPE  [header link](class-aws-sts-regionalendpoints-configurationprovider-constant-ini-endpoints-type.md)

`
    public
        mixed
    INI_ENDPOINTS_TYPE
    = 'sts_regional_endpoints'
`

### Properties  [header link](class-aws-sts-regionalendpoints-configurationprovider-properties.md)

#### $cacheKey  [header link](class-aws-sts-regionalendpoints-configurationprovider-property-cachekey.md)

`
    public
    static    mixed
    $cacheKey
     = 'aws_sts_regional_endpoints_config'`

### Methods  [header link](class-aws-sts-regionalendpoints-configurationprovider-methods.md)

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

#### defaultProvider()  [header link](class-aws-sts-regionalendpoints-configurationprovider-method-defaultprovider.md)

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

#### env()  [header link](class-aws-sts-regionalendpoints-configurationprovider-method-env.md)

Provider that creates config from environment variables.

`
    public
            static        env() : callable`

##### Return values

callable

#### fallback()  [header link](class-aws-sts-regionalendpoints-configurationprovider-method-fallback.md)

Fallback config options when other sources are not set.

`
    public
            static        fallback() : callable`

##### Return values

callable

#### ini()  [header link](class-aws-sts-regionalendpoints-configurationprovider-method-ini.md)

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

#### unwrap()  [header link](class-aws-sts-regionalendpoints-configurationprovider-method-unwrap.md)

Unwraps a configuration object in whatever valid form it is in,
always returning a ConfigurationInterface object.

`
    public
            static        unwrap(mixed $config) : ConfigurationInterface`

##### Parameters

$config
: mixed

##### Tags  [header link](class-aws-sts-regionalendpoints-configurationprovider-method-unwrap-tags.md)

throwsInvalidArgumentException

##### Return values

[ConfigurationInterface](class-aws-sts-regionalendpoints-configurationinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-sts-regionalendpoints-configurationprovider-toc-constants.md)
  - [Properties](class-aws-sts-regionalendpoints-configurationprovider-toc-properties.md)
  - [Methods](class-aws-sts-regionalendpoints-configurationprovider-toc-methods.md)
- Constants
  - [DEFAULT\_ENDPOINTS\_TYPE](class-aws-sts-regionalendpoints-configurationprovider-constant-default-endpoints-type.md)
  - [ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider.md#constant_ENV_CONFIG_FILE)
  - [ENV\_ENDPOINTS\_TYPE](class-aws-sts-regionalendpoints-configurationprovider-constant-env-endpoints-type.md)
  - [ENV\_PROFILE](class-aws-sts-regionalendpoints-configurationprovider-constant-env-profile.md)
  - [INI\_ENDPOINTS\_TYPE](class-aws-sts-regionalendpoints-configurationprovider-constant-ini-endpoints-type.md)
- Properties
  - [$cacheKey](class-aws-sts-regionalendpoints-configurationprovider-property-cachekey.md)
- Methods
  - [cache()](class-aws-abstractconfigurationprovider.md#method_cache)
  - [chain()](class-aws-abstractconfigurationprovider.md#method_chain)
  - [defaultProvider()](class-aws-sts-regionalendpoints-configurationprovider-method-defaultprovider.md)
  - [env()](class-aws-sts-regionalendpoints-configurationprovider-method-env.md)
  - [fallback()](class-aws-sts-regionalendpoints-configurationprovider-method-fallback.md)
  - [ini()](class-aws-sts-regionalendpoints-configurationprovider-method-ini.md)
  - [memoize()](class-aws-abstractconfigurationprovider.md#method_memoize)
  - [unwrap()](class-aws-sts-regionalendpoints-configurationprovider-method-unwrap.md)

[Back To Top](class-aws-sts-regionalendpoints-configurationprovider-top.md)
