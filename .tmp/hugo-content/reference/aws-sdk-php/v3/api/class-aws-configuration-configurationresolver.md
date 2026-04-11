Menu

- [Aws](namespace-aws.md)
- [Configuration](namespace-aws-configuration.md)

## ConfigurationResolver        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-configuration-configurationresolver-toc.md)

#### Constants  [header link](class-aws-configuration-configurationresolver-toc-constants.md)

[ENV\_CONFIG\_FILE](class-aws-configuration-configurationresolver-constant-env-config-file.md)
= 'AWS\_CONFIG\_FILE' [ENV\_PROFILE](class-aws-configuration-configurationresolver-constant-env-profile.md)
= 'AWS\_PROFILE'

#### Properties  [header link](class-aws-configuration-configurationresolver-toc-properties.md)

[$envPrefix](class-aws-configuration-configurationresolver-property-envprefix.md)
: mixed

#### Methods  [header link](class-aws-configuration-configurationresolver-toc-methods.md)

[env()](class-aws-configuration-configurationresolver-method-env.md)
: null Resolves config values from environment variables.[ini()](class-aws-configuration-configurationresolver-method-ini.md)
: null Gets config values from a config file whose location
is specified by an environment variable 'AWS\_CONFIG\_FILE', defaulting to
~/.aws/config if not specified[resolve()](class-aws-configuration-configurationresolver-method-resolve.md)
: mixed Generic configuration resolver that first checks for environment
variables, then checks for a specified profile in the environment-defined
config file location (env variable is 'AWS\_CONFIG\_FILE', file location
defaults to ~/.aws/config), then checks for the "default" profile in the
environment-defined config file location, and failing those uses a default
fallback value.

### Constants  [header link](class-aws-configuration-configurationresolver-constants.md)

#### ENV\_CONFIG\_FILE  [header link](class-aws-configuration-configurationresolver-constant-env-config-file.md)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_PROFILE  [header link](class-aws-configuration-configurationresolver-constant-env-profile.md)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

### Properties  [header link](class-aws-configuration-configurationresolver-properties.md)

#### $envPrefix  [header link](class-aws-configuration-configurationresolver-property-envprefix.md)

`
    public
    static    mixed
    $envPrefix
     = 'AWS_'`

### Methods  [header link](class-aws-configuration-configurationresolver-methods.md)

#### env()  [header link](class-aws-configuration-configurationresolver-method-env.md)

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

#### ini()  [header link](class-aws-configuration-configurationresolver-method-ini.md)

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

#### resolve()  [header link](class-aws-configuration-configurationresolver-method-resolve.md)

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
  - [Constants](class-aws-configuration-configurationresolver-toc-constants.md)
  - [Properties](class-aws-configuration-configurationresolver-toc-properties.md)
  - [Methods](class-aws-configuration-configurationresolver-toc-methods.md)
- Constants
  - [ENV\_CONFIG\_FILE](class-aws-configuration-configurationresolver-constant-env-config-file.md)
  - [ENV\_PROFILE](class-aws-configuration-configurationresolver-constant-env-profile.md)
- Properties
  - [$envPrefix](class-aws-configuration-configurationresolver-property-envprefix.md)
- Methods
  - [env()](class-aws-configuration-configurationresolver-method-env.md)
  - [ini()](class-aws-configuration-configurationresolver-method-ini.md)
  - [resolve()](class-aws-configuration-configurationresolver-method-resolve.md)

[Back To Top](class-aws-configuration-configurationresolver-top.md)
