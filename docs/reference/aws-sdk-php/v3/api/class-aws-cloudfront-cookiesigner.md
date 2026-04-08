Menu

- [Aws](namespace-aws.md)
- [CloudFront](namespace-aws-cloudfront.md)

## CookieSigner        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-cloudfront-cookiesigner-toc.md)

#### Methods  [header link](class-aws-cloudfront-cookiesigner-toc-methods.md)

[\_\_construct()](class-aws-cloudfront-cookiesigner-method-construct.md)
: mixed [getSignedCookie()](class-aws-cloudfront-cookiesigner-method-getsignedcookie.md)
: array<string\|int, mixed> Create a signed Amazon CloudFront Cookie.

### Methods  [header link](class-aws-cloudfront-cookiesigner-methods.md)

#### \_\_construct()  [header link](class-aws-cloudfront-cookiesigner-method-construct.md)

`
    public
                    __construct( $keyPairId,  $privateKey) : mixed`

##### Parameters

$keyPairId
:

string ID of the key pair

$privateKey
:

string Path to the private key used for signing

##### Tags  [header link](class-aws-cloudfront-cookiesigner-method-construct-tags.md)

throwsRuntimeException

if the openssl extension is missing

throwsInvalidArgumentException

if the private key cannot be found.

#### getSignedCookie()  [header link](class-aws-cloudfront-cookiesigner-method-getsignedcookie.md)

Create a signed Amazon CloudFront Cookie.

`
    public
                    getSignedCookie([string $url = null ][, string|int|null $expires = null ][, string $policy = null ]) : array<string|int, mixed>`

##### Parameters

$url
: string
= null

URL to sign (can include query string
and wildcards). Not required
when passing a custom $policy.

$expires
: string\|int\|null
= null

UTC Unix timestamp used when signing
with a canned policy. Not required
when passing a custom $policy.

$policy
: string
= null

JSON policy. Use this option when
creating a signed cookie for a custom
policy.

##### Tags  [header link](class-aws-cloudfront-cookiesigner-method-getsignedcookie-tags.md)

throwsInvalidArgumentException

if the URL provided is invalid

link[http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-signed-cookies.html](../../../../services/amazoncloudfront/latest/developerguide/private-content-signed-cookies.md)

##### Return values

array<string\|int, mixed>
—

The authenticated cookie parameters

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-cloudfront-cookiesigner-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-cloudfront-cookiesigner-method-construct.md)
  - [getSignedCookie()](class-aws-cloudfront-cookiesigner-method-getsignedcookie.md)

[Back To Top](class-aws-cloudfront-cookiesigner-top.md)
