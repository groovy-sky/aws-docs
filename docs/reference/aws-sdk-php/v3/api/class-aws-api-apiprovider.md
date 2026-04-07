Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## ApiProvider        in package    - [Aws](package-aws.md)

API providers.

An API provider is a function that accepts a type, service, and version and
returns an array of API data on success or NULL if no API data can be created
for the provided arguments.

You can wrap your calls to an API provider with the
ApiProvider::resolve method to ensure that API data is created. If the
API data is not created, then the resolve() method will throw a
UnresolvedApiException.

```prettyprint
use Aws\Api\ApiProvider;
$provider = ApiProvider::defaultProvider();
// Returns an array or NULL.
$data = $provider('api', 's3', '2006-03-01');
// Returns an array or throws.
$data = ApiProvider::resolve($provider, 'api', 'elasticfood', '2020-01-01');

```

You can compose multiple providers into a single provider using
or\_chain. This method accepts providers as arguments and
returns a new function that will invoke each provider until a non-null value
is returned.

```prettyprint
$a = ApiProvider::filesystem(sys_get_temp_dir() . '/aws-beta-models');
$b = ApiProvider::manifest();

$c = \Aws\or_chain($a, $b);
$data = $c('api', 'betaservice', '2015-08-08'); // $a handles this.
$data = $c('api', 's3', '2006-03-01');          // $b handles this.
$data = $c('api', 'invalid', '2014-12-15');     // Neither handles this.

```

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#toc-methods)

[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method___invoke)
: array<string\|int, mixed>\|null Execute the provider.[defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_defaultProvider)
: self Default SDK API provider.[filesystem()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_filesystem)
: self Loads API data from the specified directory.[getVersions()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_getVersions)
: array<string\|int, mixed> Retrieves a list of valid versions for the specified service.[manifest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_manifest)
: self Loads API data after resolving the version to the latest, compatible,
available version based on the provided manifest data.[resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_resolve)
: array<string\|int, mixed> Resolves an API provider and ensures a non-null return value.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#methods)

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#method___invoke)

Execute the provider.

`
    public
                    __invoke(string $type, string $service, string $version) : array<string|int, mixed>|null`

##### Parameters

$type
: string

Type of data ('api', 'waiter', 'paginator').

$service
: string

Service name.

$version
: string

API version.

##### Return values

array<string\|int, mixed>\|null

#### defaultProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#method_defaultProvider)

Default SDK API provider.

`
    public
            static        defaultProvider() : self`

This provider loads pre-built manifest data from the `data` directory.

##### Return values

self

#### filesystem()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#method_filesystem)

Loads API data from the specified directory.

`
    public
            static        filesystem(string $dir) : self`

If "latest" is specified as the version, this provider must glob the
directory to find which is the latest available version.

##### Parameters

$dir
: string

Directory containing service models.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#method_filesystem\#tags)

throwsInvalidArgumentException

if the provided `$dir` is invalid.

##### Return values

self

#### getVersions()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#method_getVersions)

Retrieves a list of valid versions for the specified service.

`
    public
                    getVersions(string $service) : array<string|int, mixed>`

##### Parameters

$service
: string

Service name

##### Return values

array<string\|int, mixed>

#### manifest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#method_manifest)

Loads API data after resolving the version to the latest, compatible,
available version based on the provided manifest data.

`
    public
            static        manifest(string $dir, array<string|int, mixed> $manifest) : self`

Manifest data is essentially an associative array of service names to
associative arrays of API version aliases.

\[
...
'ec2' => \[
'latest' => '2014-10-01',
'2014-10-01' => '2014-10-01',
'2014-09-01' => '2014-10-01',
'2014-06-15' => '2014-10-01',
...
\],
'ecs' => \[...\],
'elasticache' => \[...\],
...
\]

##### Parameters

$dir
: string

Directory containing service models.

$manifest
: array<string\|int, mixed>

The API version manifest data.

##### Return values

self

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#method_resolve)

Resolves an API provider and ensures a non-null return value.

`
    public
            static        resolve(callable $provider, string $type, string $service, string $version) : array<string|int, mixed>`

##### Parameters

$provider
: callable

Provider function to invoke.

$type
: string

Type of data ('api', 'waiter', 'paginator').

$service
: string

Service name.

$version
: string

API version.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html\#method_resolve\#tags)

throws[UnresolvedApiException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.UnresolvedApiException.html)

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#toc-methods)
- Methods
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method___invoke)
  - [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_defaultProvider)
  - [filesystem()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_filesystem)
  - [getVersions()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_getVersions)
  - [manifest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_manifest)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#method_resolve)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ApiProvider.html#top)
