Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## CredentialProvider        in package    - [Aws](package-aws.md)

Credential providers are functions that accept no arguments and return a
promise that is fulfilled with an {@see \\Aws\\Credentials\\CredentialsInterface}
or rejected with an {@see \\Aws\\Exception\\CredentialsException}.

`
use Aws\Credentials\CredentialProvider;
$provider = CredentialProvider::defaultProvider();
// Returns a CredentialsInterface or throws.
$creds = $provider()->wait();
`

Credential providers can be composed to create credentials using conditional
logic that can create different credentials in different environments. You
can compose multiple providers into a single provider using
CredentialProvider::chain. This function accepts
providers as variadic arguments and returns a new function that will invoke
each provider until a successful set of credentials is returned.

`
// First try an INI file at this location.
$a = CredentialProvider::ini(null, '/path/to/file.ini');
// Then try an INI file at this location.
$b = CredentialProvider::ini(null, '/path/to/other-file.ini');
// Then try loading from environment variables.
$c = CredentialProvider::env();
// Combine the three providers together.
$composed = CredentialProvider::chain($a, $b, $c);
// Returns a promise that is fulfilled with credentials or throws.
$promise = $composed();
// Wait on the credentials to resolve.
$creds = $promise->wait();
`

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#toc-constants)

[ENV\_ACCOUNT\_ID](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_ACCOUNT_ID)
= 'AWS\_ACCOUNT\_ID' [ENV\_ARN](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_ARN)
= 'AWS\_ROLE\_ARN' [ENV\_CONFIG\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_CONFIG_FILE)
= 'AWS\_CONFIG\_FILE' [ENV\_KEY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_KEY)
= 'AWS\_ACCESS\_KEY\_ID' [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_PROFILE)
= 'AWS\_PROFILE' [ENV\_REGION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_REGION)
= 'AWS\_REGION' [ENV\_ROLE\_SESSION\_NAME](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_ROLE_SESSION_NAME)
= 'AWS\_ROLE\_SESSION\_NAME' [ENV\_SECRET](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_SECRET)
= 'AWS\_SECRET\_ACCESS\_KEY' [ENV\_SESSION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_SESSION)
= 'AWS\_SESSION\_TOKEN' [ENV\_SHARED\_CREDENTIALS\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_SHARED_CREDENTIALS_FILE)
= 'AWS\_SHARED\_CREDENTIALS\_FILE' [ENV\_TOKEN\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_TOKEN_FILE)
= 'AWS\_WEB\_IDENTITY\_TOKEN\_FILE' [FALLBACK\_REGION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_FALLBACK_REGION)
= 'us-east-1' [REFRESH\_WINDOW](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_REFRESH_WINDOW)
= 60

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#toc-methods)

[assumeRole()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_assumeRole)
: callable Credential provider that creates credentials using assume role[assumeRoleWithWebIdentityCredentialProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_assumeRoleWithWebIdentityCredentialProvider)
: callable Credential provider that creates credentials by assuming role from a
Web Identity Token[cache()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_cache)
: callable Wraps a credential provider and saves provided credentials in an
instance of Aws\\CacheInterface. Forwards calls when no credentials found
in cache and updates cache with the results.[chain()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_chain)
: callable Creates an aggregate credentials provider that invokes the provided
variadic providers one after the other until a provider returns
credentials.[defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_defaultProvider)
: callable Create a default credential provider that
first checks for environment variables,
then checks for assumed role via web identity,
then checks for cached SSO credentials from the CLI,
then check for credential\_process in the "default" profile in ~/.aws/credentials,
then checks for the "default" profile in ~/.aws/credentials,
then for credential\_process in the "default profile" profile in ~/.aws/config,
then checks for "profile default" profile in ~/.aws/config (which is
the default profile of AWS CLI),
then tries to make a GET Request to fetch credentials if ECS environment variable is presented,
finally checks for EC2 instance profile credentials.[ecsCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_ecsCredentials)
: [EcsCredentialProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html)Credential provider that creates credentials using
ecs credentials by a GET request, whose uri is specified
by environment variable[env()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_env)
: callable Provider that creates credentials from environment variables
AWS\_ACCESS\_KEY\_ID, AWS\_SECRET\_ACCESS\_KEY, and AWS\_SESSION\_TOKEN.[fromCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_fromCredentials)
: callable Create a credential provider function from a set of static credentials.[getConfigFileName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_getConfigFileName)
: string Locates shared configuration file by first checking for AWS\_CONFIG,
then falling back to the default location. Returns the path of the
resolved configuration file.[getCredentialsFileName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_getCredentialsFileName)
: string Locates credentials file by first checking for AWS\_SHARED\_CREDENTIALS\_FILE,
then falling back to the default location. Returns the path of the
resolved credentials file.[getCredentialsFromSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_getCredentialsFromSource)
: mixed [getHomeDir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_getHomeDir)
: null\|string Gets the environment's HOME directory if available.[ini()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_ini)
: callable Credentials provider that creates credentials using an ini file stored
in the current user's home directory. A source can be provided
in this file for assuming a role using the credential\_source config option.[instanceProfile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_instanceProfile)
: [InstanceProfileProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html)Credential provider that creates credentials using instance profile
credentials.[loadProfiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_loadProfiles)
: mixed Gets profiles from specified $filename, or default ini files.[login()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_login)
: callable Login credential provider for AWS local development using console credentials[memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_memoize)
: callable Wraps a credential provider and caches previously provided credentials.[process()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_process)
: callable Credentials provider that creates credentials using a process configured in
ini file stored in the current user's home directory.[shouldUseEcs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_shouldUseEcs)
: bool [sso()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_sso)
: callable Credential provider that retrieves cached SSO credentials from the CLI

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constants)

#### ENV\_ACCOUNT\_ID  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_ACCOUNT_ID)

`
    public
        mixed
    ENV_ACCOUNT_ID
    = 'AWS_ACCOUNT_ID'
`

#### ENV\_ARN  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_ARN)

`
    public
        mixed
    ENV_ARN
    = 'AWS_ROLE_ARN'
`

#### ENV\_CONFIG\_FILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_CONFIG_FILE)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_KEY  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_KEY)

`
    public
        mixed
    ENV_KEY
    = 'AWS_ACCESS_KEY_ID'
`

#### ENV\_PROFILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_PROFILE)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### ENV\_REGION  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_REGION)

`
    public
        mixed
    ENV_REGION
    = 'AWS_REGION'
`

#### ENV\_ROLE\_SESSION\_NAME  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_ROLE_SESSION_NAME)

`
    public
        mixed
    ENV_ROLE_SESSION_NAME
    = 'AWS_ROLE_SESSION_NAME'
`

#### ENV\_SECRET  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_SECRET)

`
    public
        mixed
    ENV_SECRET
    = 'AWS_SECRET_ACCESS_KEY'
`

#### ENV\_SESSION  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_SESSION)

`
    public
        mixed
    ENV_SESSION
    = 'AWS_SESSION_TOKEN'
`

#### ENV\_SHARED\_CREDENTIALS\_FILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_SHARED_CREDENTIALS_FILE)

`
    public
        mixed
    ENV_SHARED_CREDENTIALS_FILE
    = 'AWS_SHARED_CREDENTIALS_FILE'
`

#### ENV\_TOKEN\_FILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_ENV_TOKEN_FILE)

`
    public
        mixed
    ENV_TOKEN_FILE
    = 'AWS_WEB_IDENTITY_TOKEN_FILE'
`

#### FALLBACK\_REGION  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_FALLBACK_REGION)

`
    public
        mixed
    FALLBACK_REGION
    = 'us-east-1'
`

#### REFRESH\_WINDOW  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#constant_REFRESH_WINDOW)

`
    public
        mixed
    REFRESH_WINDOW
    = 60
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#methods)

#### assumeRole()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_assumeRole)

Credential provider that creates credentials using assume role

`
    public
            static        assumeRole([array<string|int, mixed> $config = [] ]) : callable`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Array of configuration data

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_assumeRole\#tags)

see[AssumeRoleCredentialProvider](class-aws-credentials-assumerolecredentialprovider.md)

for $config details.

##### Return values

callable

#### assumeRoleWithWebIdentityCredentialProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_assumeRoleWithWebIdentityCredentialProvider)

Credential provider that creates credentials by assuming role from a
Web Identity Token

`
    public
            static        assumeRoleWithWebIdentityCredentialProvider([array<string|int, mixed> $config = [] ]) : callable`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Array of configuration data

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_assumeRoleWithWebIdentityCredentialProvider\#tags)

see[AssumeRoleWithWebIdentityCredentialProvider](class-aws-credentials-assumerolewithwebidentitycredentialprovider.md)

for
$config details.

##### Return values

callable

#### cache()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_cache)

Wraps a credential provider and saves provided credentials in an
instance of Aws\\CacheInterface. Forwards calls when no credentials found
in cache and updates cache with the results.

`
    public
            static        cache(callable $provider, CacheInterface $cache[, string|null $cacheKey = null ]) : callable`

##### Parameters

$provider
: callable

Credentials provider function to wrap

$cache
: [CacheInterface](class-aws-cacheinterface.md)

Cache to store credentials

$cacheKey
: string\|null
= null

(optional) Cache key to use

##### Return values

callable

#### chain()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_chain)

Creates an aggregate credentials provider that invokes the provided
variadic providers one after the other until a provider returns
credentials.

`
    public
            static        chain() : callable`

##### Return values

callable

#### defaultProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_defaultProvider)

Create a default credential provider that
first checks for environment variables,
then checks for assumed role via web identity,
then checks for cached SSO credentials from the CLI,
then check for credential\_process in the "default" profile in ~/.aws/credentials,
then checks for the "default" profile in ~/.aws/credentials,
then for credential\_process in the "default profile" profile in ~/.aws/config,
then checks for "profile default" profile in ~/.aws/config (which is
the default profile of AWS CLI),
then tries to make a GET Request to fetch credentials if ECS environment variable is presented,
finally checks for EC2 instance profile credentials.

`
    public
            static        defaultProvider([array<string|int, mixed> $config = [] ]) : callable`

This provider is automatically wrapped in a memoize function that caches
previously provided credentials.

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Optional array of ecs/instance profile credentials
provider options.

##### Return values

callable

#### ecsCredentials()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_ecsCredentials)

Credential provider that creates credentials using
ecs credentials by a GET request, whose uri is specified
by environment variable

`
    public
            static        ecsCredentials([array<string|int, mixed> $config = [] ]) : EcsCredentialProvider`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Array of configuration data.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_ecsCredentials\#tags)

see[EcsCredentialProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html)

for $config details.

##### Return values

[EcsCredentialProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html)

#### env()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_env)

Provider that creates credentials from environment variables
AWS\_ACCESS\_KEY\_ID, AWS\_SECRET\_ACCESS\_KEY, and AWS\_SESSION\_TOKEN.

`
    public
            static        env() : callable`

##### Return values

callable

#### fromCredentials()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_fromCredentials)

Create a credential provider function from a set of static credentials.

`
    public
            static        fromCredentials(CredentialsInterface $creds) : callable`

##### Parameters

$creds
: [CredentialsInterface](class-aws-credentials-credentialsinterface.md)

##### Return values

callable

#### getConfigFileName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_getConfigFileName)

Locates shared configuration file by first checking for AWS\_CONFIG,
then falling back to the default location. Returns the path of the
resolved configuration file.

`
    public
            static        getConfigFileName() : string`

##### Return values

string

#### getCredentialsFileName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_getCredentialsFileName)

Locates credentials file by first checking for AWS\_SHARED\_CREDENTIALS\_FILE,
then falling back to the default location. Returns the path of the
resolved credentials file.

`
    public
            static        getCredentialsFileName( $filename) : string`

##### Parameters

$filename
:

##### Return values

string

#### getCredentialsFromSource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_getCredentialsFromSource)

`
    public
            static        getCredentialsFromSource([mixed $profileName = '' ][, mixed $filename = '' ][, mixed $config = [] ]) : mixed`

##### Parameters

$profileName
: mixed
= ''$filename
: mixed
= ''$config
: mixed
= \[\]

#### getHomeDir()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_getHomeDir)

Gets the environment's HOME directory if available.

`
    public
            static        getHomeDir() : null|string`

##### Return values

null\|string

#### ini()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_ini)

Credentials provider that creates credentials using an ini file stored
in the current user's home directory. A source can be provided
in this file for assuming a role using the credential\_source config option.

`
    public
            static        ini([string|null $profile = null ][, string|null $filename = null ][, array<string|int, mixed>|null $config = [] ]) : callable`

##### Parameters

$profile
: string\|null
= null

Profile to use. If not specified will use
the "default" profile in "~/.aws/credentials".

$filename
: string\|null
= null

If provided, uses a custom filename rather
than looking in the home directory.

$config
: array<string\|int, mixed>\|null
= \[\]

If provided, may contain the following:
preferStaticCredentials: If true, prefer static
credentials to role\_arn if both are present
disableAssumeRole: If true, disable support for
roles that assume an IAM role. If true and role profile
is selected, an error is raised.
stsClient: StsClient used to assume role specified in profile

##### Return values

callable

#### instanceProfile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_instanceProfile)

Credential provider that creates credentials using instance profile
credentials.

`
    public
            static        instanceProfile([array<string|int, mixed> $config = [] ]) : InstanceProfileProvider`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Array of configuration data.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_instanceProfile\#tags)

see[InstanceProfileProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html)

for $config details.

##### Return values

[InstanceProfileProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html)

#### loadProfiles()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_loadProfiles)

Gets profiles from specified $filename, or default ini files.

`
    public
            static        loadProfiles(mixed $filename) : mixed`

##### Parameters

$filename
: mixed

#### login()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_login)

Login credential provider for AWS local development using console credentials

`
    public
            static        login([string|null $profileName = null ][, array<string|int, mixed> $config = [] ]) : callable`

##### Parameters

$profileName
: string\|null
= null

profile containing your console login session information

$config
: array<string\|int, mixed>
= \[\]

region used for refresh requests.
pass `'region' => <your_region>` to configure a region,
otherwise, provider construction falls back to AWS\_REGION,
then the profile specified for `login`

##### Return values

callable

#### memoize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_memoize)

Wraps a credential provider and caches previously provided credentials.

`
    public
            static        memoize(callable $provider) : callable`

Ensures that cached credentials are refreshed when they expire.

##### Parameters

$provider
: callable

Credentials provider function to wrap.

##### Return values

callable

#### process()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_process)

Credentials provider that creates credentials using a process configured in
ini file stored in the current user's home directory.

`
    public
            static        process([string|null $profile = null ][, string|null $filename = null ]) : callable`

##### Parameters

$profile
: string\|null
= null

Profile to use. If not specified will use
the "default" profile in "~/.aws/credentials".

$filename
: string\|null
= null

If provided, uses a custom filename rather
than looking in the home directory.

##### Return values

callable

#### shouldUseEcs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_shouldUseEcs)

`
    public
            static        shouldUseEcs() : bool`

##### Return values

bool

#### sso()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html\#method_sso)

Credential provider that retrieves cached SSO credentials from the CLI

`
    public
            static        sso([mixed $ssoProfileName = 'default' ][, mixed $filename = null ][, mixed $config = [] ]) : callable`

##### Parameters

$ssoProfileName
: mixed
= 'default'$filename
: mixed
= null$config
: mixed
= \[\]

##### Return values

callable
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#toc-methods)
- Constants
  - [ENV\_ACCOUNT\_ID](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_ACCOUNT_ID)
  - [ENV\_ARN](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_ARN)
  - [ENV\_CONFIG\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_CONFIG_FILE)
  - [ENV\_KEY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_KEY)
  - [ENV\_PROFILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_PROFILE)
  - [ENV\_REGION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_REGION)
  - [ENV\_ROLE\_SESSION\_NAME](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_ROLE_SESSION_NAME)
  - [ENV\_SECRET](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_SECRET)
  - [ENV\_SESSION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_SESSION)
  - [ENV\_SHARED\_CREDENTIALS\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_SHARED_CREDENTIALS_FILE)
  - [ENV\_TOKEN\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_ENV_TOKEN_FILE)
  - [FALLBACK\_REGION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_FALLBACK_REGION)
  - [REFRESH\_WINDOW](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#constant_REFRESH_WINDOW)
- Methods
  - [assumeRole()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_assumeRole)
  - [assumeRoleWithWebIdentityCredentialProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_assumeRoleWithWebIdentityCredentialProvider)
  - [cache()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_cache)
  - [chain()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_chain)
  - [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_defaultProvider)
  - [ecsCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_ecsCredentials)
  - [env()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_env)
  - [fromCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_fromCredentials)
  - [getConfigFileName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_getConfigFileName)
  - [getCredentialsFileName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_getCredentialsFileName)
  - [getCredentialsFromSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_getCredentialsFromSource)
  - [getHomeDir()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_getHomeDir)
  - [ini()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_ini)
  - [instanceProfile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_instanceProfile)
  - [loadProfiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_loadProfiles)
  - [login()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_login)
  - [memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_memoize)
  - [process()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_process)
  - [shouldUseEcs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_shouldUseEcs)
  - [sso()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#method_sso)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html#top)
