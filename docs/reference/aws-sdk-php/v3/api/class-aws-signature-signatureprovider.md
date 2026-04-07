Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## SignatureProvider        in package    - [Aws](package-aws.md)

Signature providers.

A signature provider is a function that accepts a version, service, and
region and returns a [SignatureInterface](class-aws-signature-signatureinterface.md) object on success or NULL if
no signature can be created from the provided arguments.

You can wrap your calls to a signature provider with the
SignatureProvider::resolve function to ensure that a signature object
is created. If a signature object is not created, then the resolve()
function will throw a UnresolvedSignatureException.

```prettyprint
use Aws\Signature\SignatureProvider;
$provider = SignatureProvider::defaultProvider();
// Returns a SignatureInterface or NULL.
$signer = $provider('v4', 's3', 'us-west-2');
// Returns a SignatureInterface or throws.
$signer = SignatureProvider::resolve($provider, 'no', 's3', 'foo');

```

You can compose multiple providers into a single provider using
or\_chain. This function accepts providers as arguments and
returns a new function that will invoke each provider until a non-null value
is returned.

```prettyprint
$a = SignatureProvider::defaultProvider();
$b = function ($version, $service, $region) {
    if ($version === 'foo') {
        return new MyFooSignature();
    }
};
$c = \Aws\or_chain($a, $b);
$signer = $c('v4', 'abc', '123');     // $a handles this.
$signer = $c('foo', 'abc', '123');    // $b handles this.
$nullValue = $c('???', 'abc', '123'); // Neither can handle this.

```

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html\#toc-methods)

[defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#method_defaultProvider)
: callable Default SDK signature provider.[memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#method_memoize)
: callable Creates a signature provider that caches previously created signature
objects. The computed cache key is the concatenation of the version,
service, and region.[resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#method_resolve)
: [SignatureInterface](class-aws-signature-signatureinterface.md)Resolves and signature provider and ensures a non-null return value.[version()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#method_version)
: callable Creates signature objects from known signature versions.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html\#methods)

#### defaultProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html\#method_defaultProvider)

Default SDK signature provider.

`
    public
            static        defaultProvider() : callable`

##### Return values

callable

#### memoize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html\#method_memoize)

Creates a signature provider that caches previously created signature
objects. The computed cache key is the concatenation of the version,
service, and region.

`
    public
            static        memoize(callable $provider) : callable`

##### Parameters

$provider
: callable

Signature provider to wrap.

##### Return values

callable

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html\#method_resolve)

Resolves and signature provider and ensures a non-null return value.

`
    public
            static        resolve(callable $provider, string $version, string $service, string $region) : SignatureInterface`

##### Parameters

$provider
: callable

Provider function to invoke.

$version
: string

Signature version.

$service
: string

Service name.

$region
: string

Region name.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html\#method_resolve\#tags)

throws[UnresolvedSignatureException](class-aws-exception-unresolvedsignatureexception.md)

##### Return values

[SignatureInterface](class-aws-signature-signatureinterface.md)

#### version()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html\#method_version)

Creates signature objects from known signature versions.

`
    public
            static        version() : callable`

This provider currently recognizes the following signature versions:

- v4: Signature version 4.
- anonymous: Does not sign requests.

##### Return values

callable
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#toc-methods)
- Methods
  - [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#method_defaultProvider)
  - [memoize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#method_memoize)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#method_resolve)
  - [version()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#method_version)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureProvider.html#top)
