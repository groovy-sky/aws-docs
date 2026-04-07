Menu

- [Aws](namespace-aws.md)
- [CloudFront](namespace-aws-cloudfront.md)

## UrlSigner        in package    - [Aws](package-aws.md)

Creates signed URLs for Amazon CloudFront resources.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html#method___construct)
: mixed [getSignedUrl()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html#method_getSignedUrl)
: string Create a signed Amazon CloudFront URL.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html\#method___construct)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html\#method___construct\#tags)

throwsRuntimeException

if the openssl extension is missing

throwsInvalidArgumentException

if the private key cannot be found.

#### getSignedUrl()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html\#method_getSignedUrl)

Create a signed Amazon CloudFront URL.

`
    public
                    getSignedUrl(string $url[, string|int|null $expires = null ][, string $policy = null ]) : string`

Keep in mind that URLs meant for use in media/flash players may have
different requirements for URL formats (e.g. some require that the
extension be removed, some require the file name to be prefixed

- mp4:, some require you to add "/cfx/st" into your URL).

##### Parameters

$url
: string

URL to sign (can include query
string string and wildcards)

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
creating a signed URL for a custom
policy.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html\#method_getSignedUrl\#tags)

throwsInvalidArgumentException

if the URL provided is invalid

link[http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html)

##### Return values

string
—

The file URL with authentication parameters

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html#method___construct)
  - [getSignedUrl()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html#method_getSignedUrl)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.UrlSigner.html#top)
