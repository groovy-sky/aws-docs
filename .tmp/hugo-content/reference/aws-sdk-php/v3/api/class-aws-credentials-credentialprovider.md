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

### Table of Contents  [header link](class-aws-credentials-credentialprovider-toc.md)

#### Constants  [header link](class-aws-credentials-credentialprovider-toc-constants.md)

[ENV\_ACCOUNT\_ID](class-aws-credentials-credentialprovider-constant-env-account-id.md)
= 'AWS\_ACCOUNT\_ID' [ENV\_ARN](class-aws-credentials-credentialprovider-constant-env-arn.md)
= 'AWS\_ROLE\_ARN' [ENV\_CONFIG\_FILE](class-aws-credentials-credentialprovider-constant-env-config-file.md)
= 'AWS\_CONFIG\_FILE' [ENV\_KEY](class-aws-credentials-credentialprovider-constant-env-key.md)
= 'AWS\_ACCESS\_KEY\_ID' [ENV\_PROFILE](class-aws-credentials-credentialprovider-constant-env-profile.md)
= 'AWS\_PROFILE' [ENV\_REGION](class-aws-credentials-credentialprovider-constant-env-region.md)
= 'AWS\_REGION' [ENV\_ROLE\_SESSION\_NAME](class-aws-credentials-credentialprovider-constant-env-role-session-name.md)
= 'AWS\_ROLE\_SESSION\_NAME' [ENV\_SECRET](class-aws-credentials-credentialprovider-constant-env-secret.md)
= 'AWS\_SECRET\_ACCESS\_KEY' [ENV\_SESSION](class-aws-credentials-credentialprovider-constant-env-session.md)
= 'AWS\_SESSION\_TOKEN' [ENV\_SHARED\_CREDENTIALS\_FILE](class-aws-credentials-credentialprovider-constant-env-shared-credentials-file.md)
= 'AWS\_SHARED\_CREDENTIALS\_FILE' [ENV\_TOKEN\_FILE](class-aws-credentials-credentialprovider-constant-env-token-file.md)
= 'AWS\_WEB\_IDENTITY\_TOKEN\_FILE' [FALLBACK\_REGION](class-aws-credentials-credentialprovider-constant-fallback-region.md)
= 'us-east-1' [REFRESH\_WINDOW](class-aws-credentials-credentialprovider-constant-refresh-window.md)
= 60

#### Methods  [header link](class-aws-credentials-credentialprovider-toc-methods.md)

[assumeRole()](class-aws-credentials-credentialprovider-method-assumerole.md)
: callable Credential provider that creates credentials using assume role[assumeRoleWithWebIdentityCredentialProvider()](class-aws-credentials-credentialprovider-method-assumerolewithwebidentitycredentialprovider.md)
: callable Credential provider that creates credentials by assuming role from a
Web Identity Token[cache()](class-aws-credentials-credentialprovider-method-cache.md)
: callable Wraps a credential provider and saves provided credentials in an
instance of Aws\\CacheInterface. Forwards calls when no credentials found
in cache and updates cache with the results.[chain()](class-aws-credentials-credentialprovider-method-chain.md)
: callable Creates an aggregate credentials provider that invokes the provided
variadic providers one after the other until a provider returns
credentials.[defaultProvider()](class-aws-credentials-credentialprovider-method-defaultprovider.md)
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
finally checks for EC2 instance profile credentials.[ecsCredentials()](class-aws-credentials-credentialprovider-method-ecscredentials.md)
: [EcsCredentialProvider](class-aws-credentials-ecscredentialprovider.md)Credential provider that creates credentials using
ecs credentials by a GET request, whose uri is specified
by environment variable[env()](class-aws-credentials-credentialprovider-method-env.md)
: callable Provider that creates credentials from environment variables
AWS\_ACCESS\_KEY\_ID, AWS\_SECRET\_ACCESS\_KEY, and AWS\_SESSION\_TOKEN.[fromCredentials()](class-aws-credentials-credentialprovider-method-fromcredentials.md)
: callable Create a credential provider function from a set of static credentials.[getConfigFileName()](class-aws-credentials-credentialprovider-method-getconfigfilename.md)
: string Locates shared configuration file by first checking for AWS\_CONFIG,
then falling back to the default location. Returns the path of the
resolved configuration file.[getCredentialsFileName()](class-aws-credentials-credentialprovider-method-getcredentialsfilename.md)
: string Locates credentials file by first checking for AWS\_SHARED\_CREDENTIALS\_FILE,
then falling back to the default location. Returns the path of the
resolved credentials file.[getCredentialsFromSource()](class-aws-credentials-credentialprovider-method-getcredentialsfromsource.md)
: mixed [getHomeDir()](class-aws-credentials-credentialprovider-method-gethomedir.md)
: null\|string Gets the environment's HOME directory if available.[ini()](class-aws-credentials-credentialprovider-method-ini.md)
: callable Credentials provider that creates credentials using an ini file stored
in the current user's home directory. A source can be provided
in this file for assuming a role using the credential\_source config option.[instanceProfile()](class-aws-credentials-credentialprovider-method-instanceprofile.md)
: [InstanceProfileProvider](class-aws-credentials-instanceprofileprovider.md)Credential provider that creates credentials using instance profile
credentials.[loadProfiles()](class-aws-credentials-credentialprovider-method-loadprofiles.md)
: mixed Gets profiles from specified $filename, or default ini files.[login()](class-aws-credentials-credentialprovider-method-login.md)
: callable Login credential provider for AWS local development using console credentials[memoize()](class-aws-credentials-credentialprovider-method-memoize.md)
: callable Wraps a credential provider and caches previously provided credentials.[process()](class-aws-credentials-credentialprovider-method-process.md)
: callable Credentials provider that creates credentials using a process configured in
ini file stored in the current user's home directory.[shouldUseEcs()](class-aws-credentials-credentialprovider-method-shoulduseecs.md)
: bool [sso()](class-aws-credentials-credentialprovider-method-sso.md)
: callable Credential provider that retrieves cached SSO credentials from the CLI

### Constants  [header link](class-aws-credentials-credentialprovider-constants.md)

#### ENV\_ACCOUNT\_ID  [header link](class-aws-credentials-credentialprovider-constant-env-account-id.md)

`
    public
        mixed
    ENV_ACCOUNT_ID
    = 'AWS_ACCOUNT_ID'
`

#### ENV\_ARN  [header link](class-aws-credentials-credentialprovider-constant-env-arn.md)

`
    public
        mixed
    ENV_ARN
    = 'AWS_ROLE_ARN'
`

#### ENV\_CONFIG\_FILE  [header link](class-aws-credentials-credentialprovider-constant-env-config-file.md)

`
    public
        mixed
    ENV_CONFIG_FILE
    = 'AWS_CONFIG_FILE'
`

#### ENV\_KEY  [header link](class-aws-credentials-credentialprovider-constant-env-key.md)

`
    public
        mixed
    ENV_KEY
    = 'AWS_ACCESS_KEY_ID'
`

#### ENV\_PROFILE  [header link](class-aws-credentials-credentialprovider-constant-env-profile.md)

`
    public
        mixed
    ENV_PROFILE
    = 'AWS_PROFILE'
`

#### ENV\_REGION  [header link](class-aws-credentials-credentialprovider-constant-env-region.md)

`
    public
        mixed
    ENV_REGION
    = 'AWS_REGION'
`

#### ENV\_ROLE\_SESSION\_NAME  [header link](class-aws-credentials-credentialprovider-constant-env-role-session-name.md)

`
    public
        mixed
    ENV_ROLE_SESSION_NAME
    = 'AWS_ROLE_SESSION_NAME'
`

#### ENV\_SECRET  [header link](class-aws-credentials-credentialprovider-constant-env-secret.md)

`
    public
        mixed
    ENV_SECRET
    = 'AWS_SECRET_ACCESS_KEY'
`

#### ENV\_SESSION  [header link](class-aws-credentials-credentialprovider-constant-env-session.md)

`
    public
        mixed
    ENV_SESSION
    = 'AWS_SESSION_TOKEN'
`

#### ENV\_SHARED\_CREDENTIALS\_FILE  [header link](class-aws-credentials-credentialprovider-constant-env-shared-credentials-file.md)

`
    public
        mixed
    ENV_SHARED_CREDENTIALS_FILE
    = 'AWS_SHARED_CREDENTIALS_FILE'
`

#### ENV\_TOKEN\_FILE  [header link](class-aws-credentials-credentialprovider-constant-env-token-file.md)

`
    public
        mixed
    ENV_TOKEN_FILE
    = 'AWS_WEB_IDENTITY_TOKEN_FILE'
`

#### FALLBACK\_REGION  [header link](class-aws-credentials-credentialprovider-constant-fallback-region.md)

`
    public
        mixed
    FALLBACK_REGION
    = 'us-east-1'
`

#### REFRESH\_WINDOW  [header link](class-aws-credentials-credentialprovider-constant-refresh-window.md)

`
    public
        mixed
    REFRESH_WINDOW
    = 60
`

### Methods  [header link](class-aws-credentials-credentialprovider-methods.md)

#### assumeRole()  [header link](class-aws-credentials-credentialprovider-method-assumerole.md)

Credential provider that creates credentials using assume role

`
    public
            static        assumeRole([array<string|int, mixed> $config = [] ]) : callable`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Array of configuration data

##### Tags  [header link](class-aws-credentials-credentialprovider-method-assumerole-tags.md)

see[AssumeRoleCredentialProvider](class-aws-credentials-assumerolecredentialprovider.md)

for $config details.

##### Return values

callable

#### assumeRoleWithWebIdentityCredentialProvider()  [header link](class-aws-credentials-credentialprovider-method-assumerolewithwebidentitycredentialprovider.md)

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

##### Tags  [header link](class-aws-credentials-credentialprovider-method-assumerolewithwebidentitycredentialprovider-tags.md)

see[AssumeRoleWithWebIdentityCredentialProvider](class-aws-credentials-assumerolewithwebidentitycredentialprovider.md)

for
$config details.

##### Return values

callable

#### cache()  [header link](class-aws-credentials-credentialprovider-method-cache.md)

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

#### chain()  [header link](class-aws-credentials-credentialprovider-method-chain.md)

Creates an aggregate credentials provider that invokes the provided
variadic providers one after the other until a provider returns
credentials.

`
    public
            static        chain() : callable`

##### Return values

callable

#### defaultProvider()  [header link](class-aws-credentials-credentialprovider-method-defaultprovider.md)

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

#### ecsCredentials()  [header link](class-aws-credentials-credentialprovider-method-ecscredentials.md)

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

##### Tags  [header link](class-aws-credentials-credentialprovider-method-ecscredentials-tags.md)

see[EcsCredentialProvider](class-aws-credentials-ecscredentialprovider.md)

for $config details.

##### Return values

[EcsCredentialProvider](class-aws-credentials-ecscredentialprovider.md)

#### env()  [header link](class-aws-credentials-credentialprovider-method-env.md)

Provider that creates credentials from environment variables
AWS\_ACCESS\_KEY\_ID, AWS\_SECRET\_ACCESS\_KEY, and AWS\_SESSION\_TOKEN.

`
    public
            static        env() : callable`

##### Return values

callable

#### fromCredentials()  [header link](class-aws-credentials-credentialprovider-method-fromcredentials.md)

Create a credential provider function from a set of static credentials.

`
    public
            static        fromCredentials(CredentialsInterface $creds) : callable`

##### Parameters

$creds
: [CredentialsInterface](class-aws-credentials-credentialsinterface.md)

##### Return values

callable

#### getConfigFileName()  [header link](class-aws-credentials-credentialprovider-method-getconfigfilename.md)

Locates shared configuration file by first checking for AWS\_CONFIG,
then falling back to the default location. Returns the path of the
resolved configuration file.

`
    public
            static        getConfigFileName() : string`

##### Return values

string

#### getCredentialsFileName()  [header link](class-aws-credentials-credentialprovider-method-getcredentialsfilename.md)

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

#### getCredentialsFromSource()  [header link](class-aws-credentials-credentialprovider-method-getcredentialsfromsource.md)

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

#### getHomeDir()  [header link](class-aws-credentials-credentialprovider-method-gethomedir.md)

Gets the environment's HOME directory if available.

`
    public
            static        getHomeDir() : null|string`

##### Return values

null\|string

#### ini()  [header link](class-aws-credentials-credentialprovider-method-ini.md)

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

#### instanceProfile()  [header link](class-aws-credentials-credentialprovider-method-instanceprofile.md)

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

##### Tags  [header link](class-aws-credentials-credentialprovider-method-instanceprofile-tags.md)

see[InstanceProfileProvider](class-aws-credentials-instanceprofileprovider.md)

for $config details.

##### Return values

[InstanceProfileProvider](class-aws-credentials-instanceprofileprovider.md)

#### loadProfiles()  [header link](class-aws-credentials-credentialprovider-method-loadprofiles.md)

Gets profiles from specified $filename, or default ini files.

`
    public
            static        loadProfiles(mixed $filename) : mixed`

##### Parameters

$filename
: mixed

#### login()  [header link](class-aws-credentials-credentialprovider-method-login.md)

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

#### memoize()  [header link](class-aws-credentials-credentialprovider-method-memoize.md)

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

#### process()  [header link](class-aws-credentials-credentialprovider-method-process.md)

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

#### shouldUseEcs()  [header link](class-aws-credentials-credentialprovider-method-shoulduseecs.md)

`
    public
            static        shouldUseEcs() : bool`

##### Return values

bool

#### sso()  [header link](class-aws-credentials-credentialprovider-method-sso.md)

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
  - [Constants](class-aws-credentials-credentialprovider-toc-constants.md)
  - [Methods](class-aws-credentials-credentialprovider-toc-methods.md)
- Constants
  - [ENV\_ACCOUNT\_ID](class-aws-credentials-credentialprovider-constant-env-account-id.md)
  - [ENV\_ARN](class-aws-credentials-credentialprovider-constant-env-arn.md)
  - [ENV\_CONFIG\_FILE](class-aws-credentials-credentialprovider-constant-env-config-file.md)
  - [ENV\_KEY](class-aws-credentials-credentialprovider-constant-env-key.md)
  - [ENV\_PROFILE](class-aws-credentials-credentialprovider-constant-env-profile.md)
  - [ENV\_REGION](class-aws-credentials-credentialprovider-constant-env-region.md)
  - [ENV\_ROLE\_SESSION\_NAME](class-aws-credentials-credentialprovider-constant-env-role-session-name.md)
  - [ENV\_SECRET](class-aws-credentials-credentialprovider-constant-env-secret.md)
  - [ENV\_SESSION](class-aws-credentials-credentialprovider-constant-env-session.md)
  - [ENV\_SHARED\_CREDENTIALS\_FILE](class-aws-credentials-credentialprovider-constant-env-shared-credentials-file.md)
  - [ENV\_TOKEN\_FILE](class-aws-credentials-credentialprovider-constant-env-token-file.md)
  - [FALLBACK\_REGION](class-aws-credentials-credentialprovider-constant-fallback-region.md)
  - [REFRESH\_WINDOW](class-aws-credentials-credentialprovider-constant-refresh-window.md)
- Methods
  - [assumeRole()](class-aws-credentials-credentialprovider-method-assumerole.md)
  - [assumeRoleWithWebIdentityCredentialProvider()](class-aws-credentials-credentialprovider-method-assumerolewithwebidentitycredentialprovider.md)
  - [cache()](class-aws-credentials-credentialprovider-method-cache.md)
  - [chain()](class-aws-credentials-credentialprovider-method-chain.md)
  - [defaultProvider()](class-aws-credentials-credentialprovider-method-defaultprovider.md)
  - [ecsCredentials()](class-aws-credentials-credentialprovider-method-ecscredentials.md)
  - [env()](class-aws-credentials-credentialprovider-method-env.md)
  - [fromCredentials()](class-aws-credentials-credentialprovider-method-fromcredentials.md)
  - [getConfigFileName()](class-aws-credentials-credentialprovider-method-getconfigfilename.md)
  - [getCredentialsFileName()](class-aws-credentials-credentialprovider-method-getcredentialsfilename.md)
  - [getCredentialsFromSource()](class-aws-credentials-credentialprovider-method-getcredentialsfromsource.md)
  - [getHomeDir()](class-aws-credentials-credentialprovider-method-gethomedir.md)
  - [ini()](class-aws-credentials-credentialprovider-method-ini.md)
  - [instanceProfile()](class-aws-credentials-credentialprovider-method-instanceprofile.md)
  - [loadProfiles()](class-aws-credentials-credentialprovider-method-loadprofiles.md)
  - [login()](class-aws-credentials-credentialprovider-method-login.md)
  - [memoize()](class-aws-credentials-credentialprovider-method-memoize.md)
  - [process()](class-aws-credentials-credentialprovider-method-process.md)
  - [shouldUseEcs()](class-aws-credentials-credentialprovider-method-shoulduseecs.md)
  - [sso()](class-aws-credentials-credentialprovider-method-sso.md)

[Back To Top](class-aws-credentials-credentialprovider-top.md)
