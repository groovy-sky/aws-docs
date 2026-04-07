Menu

- [Aws](namespace-aws.md)
- [Configuration](namespace-aws-configuration.md)

## ConfigurationResolver        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#toc-constants)

[ENV\_CONFIG\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#constant_ENV_CONFIG_FILE)
= 'AWS\_CONFIG\_FILE' [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#constant_ENV_PROFILE)
= 'AWS\_PROFILE'

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#toc-properties)

[$envPrefix](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#property_envPrefix)
: mixed

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#toc-methods)

[env()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#method_env)
: null Resolves config values from environment variables.[ini()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#method_ini)
: null Gets config values from a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified[resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#method_resolve)
: mixed Generic configuration resolver that first checks for environment
variables, then checks for a specified profile in the environment-defined
config file location (env variable is 'AWS\_CONFIG\_FILE', file location
defaults to ~/.aws/config), then checks for the "default" profile in the
environment-defined config file location, and failing those uses a default
fallback value.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#constants)

#### ENV\_CONFIG\_FILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#constant_ENV_CONFIG_FILE)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_PROFILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#constant_ENV_PROFILE)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#properties)

#### $envPrefix  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#property_envPrefix)

`
    public
    static    mixed
    $envPrefix
     = 'AWS_'`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#methods)

#### env()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#method_env)

Resolves config values from environment variables.

`
    public
            static        env(string $key[, string $expectedType = 'string' ]) : null`

##### Parameters

$key
: string

Configuration key to be used when attempting
to retrieve value from the environment.

$expectedType
: string
= 'string'

The expected type of the retrieved value.

##### Return values

null
—

\| mixed

#### ini()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#method_ini)

Gets config values from a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified

`
    public
            static        ini(string $key, string $expectedType[, string|null $profile = null ][, string|null $filename = null ][, mixed $options = [] ]) : null`

##### Parameters

$key
: string

Configuration key to be used when attempting
to retrieve value from ini file.

$expectedType
: string

The expected type of the retrieved value.

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

$options
: mixed
= \[\]

##### Return values

null
—

\| mixed

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html\#method_resolve)

Generic configuration resolver that first checks for environment
variables, then checks for a specified profile in the environment-defined
config file location (env variable is 'AWS\_CONFIG\_FILE', file location
defaults to ~/.aws/config), then checks for the "default" profile in the
environment-defined config file location, and failing those uses a default
fallback value.

`
    public
            static        resolve(string $key, mixed $defaultValue, string $expectedType[, array<string|int, mixed> $config = [] ]) : mixed`

##### Parameters

$key
: string

Configuration key to be used when attempting
to retrieve value from the environment or ini file.

$defaultValue
: mixed$expectedType
: string

The expected type of the retrieved value.

$config
: array<string\|int, mixed>
= \[\]

additional configuration options.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#toc-constants)
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#toc-methods)
- Constants
  - [ENV\_CONFIG\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#constant_ENV_CONFIG_FILE)
  - [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#constant_ENV_PROFILE)
- Properties
  - [$envPrefix](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#property_envPrefix)
- Methods
  - [env()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#method_env)
  - [ini()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#method_ini)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#method_resolve)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Configuration.ConfigurationResolver.html#top)
