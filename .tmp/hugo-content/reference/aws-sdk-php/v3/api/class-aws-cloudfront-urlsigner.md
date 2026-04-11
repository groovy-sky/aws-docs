Menu

- [Aws](namespace-aws.md)
- [CloudFront](namespace-aws-cloudfront.md)

## UrlSigner        in package    - [Aws](package-aws.md)

Creates signed URLs for Amazon CloudFront resources.

### Table of Contents  [header link](class-aws-cloudfront-urlsigner-toc.md)

#### Methods  [header link](class-aws-cloudfront-urlsigner-toc-methods.md)

[\_\_construct()](class-aws-cloudfront-urlsigner-method-construct.md)
: mixed [getSignedUrl()](class-aws-cloudfront-urlsigner-method-getsignedurl.md)
: string Create a signed Amazon CloudFront URL.

### Methods  [header link](class-aws-cloudfront-urlsigner-methods.md)

#### \_\_construct()  [header link](class-aws-cloudfront-urlsigner-method-construct.md)

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

##### Tags  [header link](class-aws-cloudfront-urlsigner-method-construct-tags.md)

throwsRuntimeException

if the openssl extension is missing

throwsInvalidArgumentException

if the private key cannot be found.

#### getSignedUrl()  [header link](class-aws-cloudfront-urlsigner-method-getsignedurl.md)

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

##### Tags  [header link](class-aws-cloudfront-urlsigner-method-getsignedurl-tags.md)

throwsInvalidArgumentException

if the URL provided is invalid

link[http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html](../../../../services/amazoncloudfront/latest/developerguide/workingwithstreamingdistributions.md)

##### Return values

string
—

The file URL with authentication parameters

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-cloudfront-urlsigner-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-cloudfront-urlsigner-method-construct.md)
  - [getSignedUrl()](class-aws-cloudfront-urlsigner-method-getsignedurl.md)

[Back To Top](class-aws-cloudfront-urlsigner-top.md)
