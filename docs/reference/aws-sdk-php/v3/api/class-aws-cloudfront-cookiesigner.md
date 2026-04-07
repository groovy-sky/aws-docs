Menu

- [Aws](namespace-aws.md)
- [CloudFront](namespace-aws-cloudfront.md)

## CookieSigner        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html#method___construct)
: mixed [getSignedCookie()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html#method_getSignedCookie)
: array<string\|int, mixed> Create a signed Amazon CloudFront Cookie.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html\#method___construct)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html\#method___construct\#tags)

throwsRuntimeException

if the openssl extension is missing

throwsInvalidArgumentException

if the private key cannot be found.

#### getSignedCookie()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html\#method_getSignedCookie)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html\#method_getSignedCookie\#tags)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html#method___construct)
  - [getSignedCookie()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html#method_getSignedCookie)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CookieSigner.html#top)
