Menu

- [Aws](namespace-aws.md)
- [Endpoint](namespace-aws-endpoint.md)
- [UseDualstackEndpoint](namespace-aws-endpoint-usedualstackendpoint.md)

## ConfigurationProvider     extends [AbstractConfigurationProvider](class-aws-abstractconfigurationprovider.md)   in package    - [Aws](package-aws.md)       implements  [ConfigurationProviderInterface](class-aws-configurationproviderinterface.md)

A configuration provider is a function that returns a promise that is
fulfilled with a {@see \\Aws\\Endpoint\\UseDualstackEndpoint\\onfigurationInterface}
or rejected with an {@see \\Aws\\Endpoint\\UseDualstackEndpoint\\ConfigurationException}.

`
use Aws\Endpoint\UseDualstackEndpoint\ConfigurationProvider;
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

### Table of Contents  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-toc.md)

#### Interfaces  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-toc-interfaces.md)

[ConfigurationProviderInterface](class-aws-configurationproviderinterface.md)

#### Constants  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-toc-constants.md)

[ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider.md#constant_ENV_CONFIG_FILE)
= 'AWS\_CONFIG\_FILE' [ENV\_PROFILE](class-aws-abstractconfigurationprovider.md#constant_ENV_PROFILE)
= 'AWS\_PROFILE' [ENV\_USE\_DUAL\_STACK\_ENDPOINT](class-aws-endpoint-usedualstackendpoint-configurationprovider-constant-env-use-dual-stack-endpoint.md)
= 'AWS\_USE\_DUALSTACK\_ENDPOINT' [INI\_USE\_DUAL\_STACK\_ENDPOINT](class-aws-endpoint-usedualstackendpoint-configurationprovider-constant-ini-use-dual-stack-endpoint.md)
= 'use\_dualstack\_endpoint'

#### Properties  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-toc-properties.md)

[$cacheKey](class-aws-endpoint-usedualstackendpoint-configurationprovider-property-cachekey.md)
: mixed

#### Methods  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-toc-methods.md)

[cache()](class-aws-abstractconfigurationprovider.md#method_cache)
: callable Wraps a config provider and saves provided configuration in an
instance of Aws\\CacheInterface. Forwards calls when no config found
in cache and updates cache with the results.[chain()](class-aws-abstractconfigurationprovider.md#method_chain)
: callable Creates an aggregate configuration provider that invokes the provided
variadic providers one after the other until a provider returns
configuration.[defaultProvider()](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-defaultprovider.md)
: callable Create a default config provider that first checks for environment
variables, then checks for a specified profile in the environment-defined
config file location (env variable is 'AWS\_CONFIG\_FILE', file location
defaults to ~/.aws/config), then checks for the "default" profile in the
environment-defined config file location, and failing those uses a default
fallback set of configuration options.[env()](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-env.md)
: callable Provider that creates config from environment variables.[fallback()](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-fallback.md)
: callable Fallback config options when other sources are not set.[ini()](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-ini.md)
: callable Config provider that creates config using a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified[memoize()](class-aws-abstractconfigurationprovider.md#method_memoize)
: callable Wraps a config provider and caches previously provided configuration.

### Constants  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-constants.md)

#### ENV\_CONFIG\_FILE  [header link](class-aws-abstractconfigurationprovider.md\#constant_ENV_CONFIG_FILE)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_PROFILE  [header link](class-aws-abstractconfigurationprovider.md\#constant_ENV_PROFILE)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### ENV\_USE\_DUAL\_STACK\_ENDPOINT  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-constant-env-use-dual-stack-endpoint.md)

`
    public
        mixed
    ENV_USE_DUAL_STACK_ENDPOINT
    = 'AWS_USE_DUALSTACK_ENDPOINT'
`

#### INI\_USE\_DUAL\_STACK\_ENDPOINT  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-constant-ini-use-dual-stack-endpoint.md)

`
    public
        mixed
    INI_USE_DUAL_STACK_ENDPOINT
    = 'use_dualstack_endpoint'
`

### Properties  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-properties.md)

#### $cacheKey  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-property-cachekey.md)

`
    public
    static    mixed
    $cacheKey
     = 'aws_cached_use_dualstack_endpoint_config'`

### Methods  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-methods.md)

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

#### defaultProvider()  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-defaultprovider.md)

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

#### env()  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-env.md)

Provider that creates config from environment variables.

`
    public
            static        env(mixed $region) : callable`

##### Parameters

$region
: mixed

##### Return values

callable

#### fallback()  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-fallback.md)

Fallback config options when other sources are not set.

`
    public
            static        fallback(mixed $region) : callable`

##### Parameters

$region
: mixed

##### Return values

callable

#### ini()  [header link](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-ini.md)

Config provider that creates config using a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified

`
    public
            static        ini(mixed $region[, string|null $profile = null ][, string|null $filename = null ]) : callable`

##### Parameters

$region
: mixed$profile
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
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-endpoint-usedualstackendpoint-configurationprovider-toc-constants.md)
  - [Properties](class-aws-endpoint-usedualstackendpoint-configurationprovider-toc-properties.md)
  - [Methods](class-aws-endpoint-usedualstackendpoint-configurationprovider-toc-methods.md)
- Constants
  - [ENV\_CONFIG\_FILE](class-aws-abstractconfigurationprovider.md#constant_ENV_CONFIG_FILE)
  - [ENV\_PROFILE](class-aws-abstractconfigurationprovider.md#constant_ENV_PROFILE)
  - [ENV\_USE\_DUAL\_STACK\_ENDPOINT](class-aws-endpoint-usedualstackendpoint-configurationprovider-constant-env-use-dual-stack-endpoint.md)
  - [INI\_USE\_DUAL\_STACK\_ENDPOINT](class-aws-endpoint-usedualstackendpoint-configurationprovider-constant-ini-use-dual-stack-endpoint.md)
- Properties
  - [$cacheKey](class-aws-endpoint-usedualstackendpoint-configurationprovider-property-cachekey.md)
- Methods
  - [cache()](class-aws-abstractconfigurationprovider.md#method_cache)
  - [chain()](class-aws-abstractconfigurationprovider.md#method_chain)
  - [defaultProvider()](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-defaultprovider.md)
  - [env()](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-env.md)
  - [fallback()](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-fallback.md)
  - [ini()](class-aws-endpoint-usedualstackendpoint-configurationprovider-method-ini.md)
  - [memoize()](class-aws-abstractconfigurationprovider.md#method_memoize)

[Back To Top](class-aws-endpoint-usedualstackendpoint-configurationprovider-top.md)
