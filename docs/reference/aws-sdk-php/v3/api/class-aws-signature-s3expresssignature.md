Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## S3ExpressSignature     extends [S3SignatureV4](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3SignatureV4.html)   in package    - [Aws](package-aws.md)

Amazon S3 signature version 4 support.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html\#toc-constants)

[AMZ\_CONTENT\_SHA256\_HEADER](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#constant_AMZ_CONTENT_SHA256_HEADER)
= 'X-Amz-Content-Sha256' [ISO8601\_BASIC](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#constant_ISO8601_BASIC)
= 'Ymd\\\THis\\\Z' [UNSIGNED\_PAYLOAD](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#constant_UNSIGNED_PAYLOAD)
= 'UNSIGNED-PAYLOAD'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#method___construct)
: mixed [convertPostToGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#method_convertPostToGet)
: [RequestInterface](class-psr-http-message-requestinterface.md)Converts a POST request to a GET request by moving POST fields into the
query string.[presign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html#method_presign)
: [RequestInterface](class-psr-http-message-requestinterface.md)Always add a x-amz-content-sha-256 for data integrity.[signRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html#method_signRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)Signs the specified request with an AWS signing protocol by using the
provided AWS account credentials and adding the required headers to the
request.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html\#constants)

#### AMZ\_CONTENT\_SHA256\_HEADER  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html\#constant_AMZ_CONTENT_SHA256_HEADER)

`
    public
        mixed
    AMZ_CONTENT_SHA256_HEADER
    = 'X-Amz-Content-Sha256'
`

#### ISO8601\_BASIC  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html\#constant_ISO8601_BASIC)

`
    public
        mixed
    ISO8601_BASIC
    = 'Ymd\\THis\\Z'
`

#### UNSIGNED\_PAYLOAD  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html\#constant_UNSIGNED_PAYLOAD)

`
    public
        mixed
    UNSIGNED_PAYLOAD
    = 'UNSIGNED-PAYLOAD'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html\#method___construct)

`
    public
                    __construct(string $service, string $region[, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$service
: string

Service name to use when signing

$region
: string

Region name to use when signing

$options
: array<string\|int, mixed>
= \[\]

Array of configuration options used when signing

- unsigned-body: Flag to make request have unsigned payload.
Unsigned body is used primarily for streaming requests.

#### convertPostToGet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html\#method_convertPostToGet)

Converts a POST request to a GET request by moving POST fields into the
query string.

`
    public
            static        convertPostToGet(RequestInterface $request[, mixed $additionalQueryParams = "" ]) : RequestInterface`

Useful for pre-signing query protocol requests.

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to clone

$additionalQueryParams
: mixed
= ""

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html\#method_convertPostToGet\#tags)

throwsInvalidArgumentException

if the method is not POST

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### presign()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html\#method_presign)

Always add a x-amz-content-sha-256 for data integrity.

`
    public
                    presign(RequestInterface $request, CredentialsInterface $credentials, mixed $expires[, array<string|int, mixed> $options = [] ]) : RequestInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to sign

$credentials
: [CredentialsInterface](class-aws-credentials-credentialsinterface.md)

Credentials used to sign

$expires
: mixed

The time at which the URL should
expire. This can be a Unix timestamp, a PHP DateTime object, or a
string that can be evaluated by strtotime.

$options
: array<string\|int, mixed>
= \[\]

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### signRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html\#method_signRequest)

Signs the specified request with an AWS signing protocol by using the
provided AWS account credentials and adding the required headers to the
request.

`
    public
                    signRequest(RequestInterface $request, CredentialsInterface $credentials[, mixed $signingService = 's3express' ]) : RequestInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to sign

$credentials
: [CredentialsInterface](class-aws-credentials-credentialsinterface.md)

Signing credentials

$signingService
: mixed
= 's3express'

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)
—

Returns the modified request.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html#toc-methods)
- Constants
  - [AMZ\_CONTENT\_SHA256\_HEADER](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#constant_AMZ_CONTENT_SHA256_HEADER)
  - [ISO8601\_BASIC](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#constant_ISO8601_BASIC)
  - [UNSIGNED\_PAYLOAD](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#constant_UNSIGNED_PAYLOAD)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#method___construct)
  - [convertPostToGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureV4.html#method_convertPostToGet)
  - [presign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html#method_presign)
  - [signRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html#method_signRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.S3ExpressSignature.html#top)
