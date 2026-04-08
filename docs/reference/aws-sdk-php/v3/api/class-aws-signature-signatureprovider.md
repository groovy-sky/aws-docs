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

### Table of Contents  [header link](class-aws-signature-signatureprovider-toc.md)

#### Methods  [header link](class-aws-signature-signatureprovider-toc-methods.md)

[defaultProvider()](class-aws-signature-signatureprovider-method-defaultprovider.md)
: callable Default SDK signature provider.[memoize()](class-aws-signature-signatureprovider-method-memoize.md)
: callable Creates a signature provider that caches previously created signature
objects. The computed cache key is the concatenation of the version,
service, and region.[resolve()](class-aws-signature-signatureprovider-method-resolve.md)
: [SignatureInterface](class-aws-signature-signatureinterface.md)Resolves and signature provider and ensures a non-null return value.[version()](class-aws-signature-signatureprovider-method-version.md)
: callable Creates signature objects from known signature versions.

### Methods  [header link](class-aws-signature-signatureprovider-methods.md)

#### defaultProvider()  [header link](class-aws-signature-signatureprovider-method-defaultprovider.md)

Default SDK signature provider.

`
    public
            static        defaultProvider() : callable`

##### Return values

callable

#### memoize()  [header link](class-aws-signature-signatureprovider-method-memoize.md)

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

#### resolve()  [header link](class-aws-signature-signatureprovider-method-resolve.md)

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

##### Tags  [header link](class-aws-signature-signatureprovider-method-resolve-tags.md)

throws[UnresolvedSignatureException](class-aws-exception-unresolvedsignatureexception.md)

##### Return values

[SignatureInterface](class-aws-signature-signatureinterface.md)

#### version()  [header link](class-aws-signature-signatureprovider-method-version.md)

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
  - [Methods](class-aws-signature-signatureprovider-toc-methods.md)
- Methods
  - [defaultProvider()](class-aws-signature-signatureprovider-method-defaultprovider.md)
  - [memoize()](class-aws-signature-signatureprovider-method-memoize.md)
  - [resolve()](class-aws-signature-signatureprovider-method-resolve.md)
  - [version()](class-aws-signature-signatureprovider-method-version.md)

[Back To Top](class-aws-signature-signatureprovider-top.md)
