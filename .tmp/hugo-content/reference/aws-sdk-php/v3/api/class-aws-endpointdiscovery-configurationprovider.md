Menu

- [Aws](namespace-aws.md)
- [EndpointDiscovery](namespace-aws-endpointdiscovery.md)

## ConfigurationProvider     extends [AbstractConfigurationProvider](class-aws-abstractconfigurationprovider.md)   in package    - [Aws](package-aws.md)       implements  [ConfigurationProviderInterface](class-aws-configurationproviderinterface.md)

A configuration provider is a function that returns a promise that is
fulfilled with a {@see \\Aws\\EndpointDiscovery\\ConfigurationInterface}
or rejected with an {@see \\Aws\\EndpointDiscovery\\Exception\\ConfigurationException}.

`
use Aws\EndpointDiscovery\ConfigurationProvider;
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

### Table of Contents  [header link](class-aws-endpointdiscovery-configurationprovider-toc.md)

#### Interfaces  [header link](class-aws-endpointdiscovery-configurationprovider-toc-interfaces.md)

[ConfigurationProviderInterface](class-aws-configurationproviderinterface.md)

#### Constants  [header link](class-aws-endpointdiscovery-configurationprovider-toc-constants.md)

[DEFAULT\_CACHE\_LIMIT](class-aws-endpointdiscovery-configurationprovider-constant-default-cache-limit.md)
= 1000 [DEFAULT\_ENABLED](class-aws-endpointdiscovery-configurationprovider-constant-default-enabled.md)
= false [ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider.md#constant_ENV_CONFIG_FILE)
= 'AWS\_CONFIG\_FILE' [ENV\_ENABLED](class-aws-endpointdiscovery-configurationprovider-constant-env-enabled.md)
= 'AWS\_ENDPOINT\_DISCOVERY\_ENABLED' [ENV\_ENABLED\_ALT](class-aws-endpointdiscovery-configurationprovider-constant-env-enabled-alt.md)
= 'AWS\_ENABLE\_ENDPOINT\_DISCOVERY' [ENV\_PROFILE](class-aws-endpointdiscovery-configurationprovider-constant-env-profile.md)
= 'AWS\_PROFILE'

#### Properties  [header link](class-aws-endpointdiscovery-configurationprovider-toc-properties.md)

[$cacheKey](class-aws-endpointdiscovery-configurationprovider-property-cachekey.md)
: mixed

#### Methods  [header link](class-aws-endpointdiscovery-configurationprovider-toc-methods.md)

[cache()](class-aws-abstractconfigurationprovider.md#method_cache)
: callable Wraps a config provider and saves provided configuration in an
instance of Aws\\CacheInterface. Forwards calls when no config found
in cache and updates cache with the results.[chain()](class-aws-abstractconfigurationprovider.md#method_chain)
: callable Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.[defaultProvider()](class-aws-endpointdiscovery-configurationprovider-method-defaultprovider.md)
: callable Create a default config provider that first checks for environment
variables, then checks for a specified profile in the environment-defined
config file location (env variable is 'AWS\_CONFIG\_FILE', file location
defaults to ~/.aws/config), then checks for the "default" profile in the
environment-defined config file location, and failing those uses a default
fallback set of configuration options.[env()](class-aws-endpointdiscovery-configurationprovider-method-env.md)
: callable Provider that creates config from environment variables.[fallback()](class-aws-endpointdiscovery-configurationprovider-method-fallback.md)
: callable Fallback config options when other sources are not set. Will check the
service model for any endpoint discovery required operations, and enable
endpoint discovery in that case. If no required operations found, will use
the class default values.[ini()](class-aws-endpointdiscovery-configurationprovider-method-ini.md)
: callable Config provider that creates config using a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified[memoize()](class-aws-abstractconfigurationprovider.md#method_memoize)
: callable Wraps a config provider and caches previously provided configuration.[unwrap()](class-aws-endpointdiscovery-configurationprovider-method-unwrap.md)
: [ConfigurationInterface](class-aws-endpointdiscovery-configurationinterface.md)Unwraps a configuration object in whatever valid form it is in,
always returning a ConfigurationInterface object.

### Constants  [header link](class-aws-endpointdiscovery-configurationprovider-constants.md)

#### DEFAULT\_CACHE\_LIMIT  [header link](class-aws-endpointdiscovery-configurationprovider-constant-default-cache-limit.md)

`
    public
        mixed
    DEFAULT_CACHE_LIMIT
    = 1000
`

#### DEFAULT\_ENABLED  [header link](class-aws-endpointdiscovery-configurationprovider-constant-default-enabled.md)

`
    public
        mixed
    DEFAULT_ENABLED
    = false
`

#### ENV\_CONFIG\_FILE  [header link](class-aws-abstractconfigurationprovider.md\#constant_ENV_CONFIG_FILE)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_ENABLED  [header link](class-aws-endpointdiscovery-configurationprovider-constant-env-enabled.md)

`
    public
        mixed
    ENV_ENABLED
    = 'AWS_ENDPOINT_DISCOVERY_ENABLED'
`

#### ENV\_ENABLED\_ALT  [header link](class-aws-endpointdiscovery-configurationprovider-constant-env-enabled-alt.md)

`
    public
        mixed
    ENV_ENABLED_ALT
    = 'AWS_ENABLE_ENDPOINT_DISCOVERY'
`

#### ENV\_PROFILE  [header link](class-aws-endpointdiscovery-configurationprovider-constant-env-profile.md)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

### Properties  [header link](class-aws-endpointdiscovery-configurationprovider-properties.md)

#### $cacheKey  [header link](class-aws-endpointdiscovery-configurationprovider-property-cachekey.md)

`
    public
    static    mixed
    $cacheKey
     = 'aws_cached_endpoint_discovery_config'`

### Methods  [header link](class-aws-endpointdiscovery-configurationprovider-methods.md)

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

#### defaultProvider()  [header link](class-aws-endpointdiscovery-configurationprovider-method-defaultprovider.md)

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

#### env()  [header link](class-aws-endpointdiscovery-configurationprovider-method-env.md)

Provider that creates config from environment variables.

`
    public
            static        env([ $cacheLimit = self::DEFAULT_CACHE_LIMIT ]) : callable`

##### Parameters

$cacheLimit
:
= self::DEFAULT\_CACHE\_LIMIT

##### Return values

callable

#### fallback()  [header link](class-aws-endpointdiscovery-configurationprovider-method-fallback.md)

Fallback config options when other sources are not set. Will check the
service model for any endpoint discovery required operations, and enable
endpoint discovery in that case. If no required operations found, will use
the class default values.

`
    public
            static        fallback([array<string|int, mixed> $config = [] ]) : callable`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Return values

callable

#### ini()  [header link](class-aws-endpointdiscovery-configurationprovider-method-ini.md)

Config provider that creates config using a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified

`
    public
            static        ini([string|null $profile = null ][, string|null $filename = null ][, int $cacheLimit = self::DEFAULT_CACHE_LIMIT ]) : callable`

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

$cacheLimit
: int
= self::DEFAULT\_CACHE\_LIMIT

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

#### unwrap()  [header link](class-aws-endpointdiscovery-configurationprovider-method-unwrap.md)

Unwraps a configuration object in whatever valid form it is in,
always returning a ConfigurationInterface object.

`
    public
            static        unwrap(mixed $config) : ConfigurationInterface`

##### Parameters

$config
: mixed

##### Tags  [header link](class-aws-endpointdiscovery-configurationprovider-method-unwrap-tags.md)

throwsInvalidArgumentException

##### Return values

[ConfigurationInterface](class-aws-endpointdiscovery-configurationinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-endpointdiscovery-configurationprovider-toc-constants.md)
  - [Properties](class-aws-endpointdiscovery-configurationprovider-toc-properties.md)
  - [Methods](class-aws-endpointdiscovery-configurationprovider-toc-methods.md)
- Constants
  - [DEFAULT\_CACHE\_LIMIT](class-aws-endpointdiscovery-configurationprovider-constant-default-cache-limit.md)
  - [DEFAULT\_ENABLED](class-aws-endpointdiscovery-configurationprovider-constant-default-enabled.md)
  - [ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider.md#constant_ENV_CONFIG_FILE)
  - [ENV\_ENABLED](class-aws-endpointdiscovery-configurationprovider-constant-env-enabled.md)
  - [ENV\_ENABLED\_ALT](class-aws-endpointdiscovery-configurationprovider-constant-env-enabled-alt.md)
  - [ENV\_PROFILE](class-aws-endpointdiscovery-configurationprovider-constant-env-profile.md)
- Properties
  - [$cacheKey](class-aws-endpointdiscovery-configurationprovider-property-cachekey.md)
- Methods
  - [cache()](class-aws-abstractconfigurationprovider.md#method_cache)
  - [chain()](class-aws-abstractconfigurationprovider.md#method_chain)
  - [defaultProvider()](class-aws-endpointdiscovery-configurationprovider-method-defaultprovider.md)
  - [env()](class-aws-endpointdiscovery-configurationprovider-method-env.md)
  - [fallback()](class-aws-endpointdiscovery-configurationprovider-method-fallback.md)
  - [ini()](class-aws-endpointdiscovery-configurationprovider-method-ini.md)
  - [memoize()](class-aws-abstractconfigurationprovider.md#method_memoize)
  - [unwrap()](class-aws-endpointdiscovery-configurationprovider-method-unwrap.md)

[Back To Top](class-aws-endpointdiscovery-configurationprovider-top.md)
