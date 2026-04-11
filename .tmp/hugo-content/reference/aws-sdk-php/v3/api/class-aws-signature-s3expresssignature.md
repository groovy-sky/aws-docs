Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## S3ExpressSignature     extends [S3SignatureV4](class-aws-signature-s3signaturev4.md)   in package    - [Aws](package-aws.md)

Amazon S3 signature version 4 support.

### Table of Contents  [header link](class-aws-signature-s3expresssignature-toc.md)

#### Constants  [header link](class-aws-signature-s3expresssignature-toc-constants.md)

[AMZ\_CONTENT\_SHA256\_HEADER](class-aws-signature-signaturev4-constant-amz-content-sha256-header.md)
= 'X-Amz-Content-Sha256' [ISO8601\_BASIC](class-aws-signature-signaturev4-constant-iso8601-basic.md)
= 'Ymd\\\THis\\\Z' [UNSIGNED\_PAYLOAD](class-aws-signature-signaturev4-constant-unsigned-payload.md)
= 'UNSIGNED-PAYLOAD'

#### Methods  [header link](class-aws-signature-s3expresssignature-toc-methods.md)

[\_\_construct()](class-aws-signature-signaturev4-method-construct.md)
: mixed [convertPostToGet()](class-aws-signature-signaturev4-method-convertposttoget.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Converts a POST request to a GET request by moving POST fields into the
query string.[presign()](class-aws-signature-s3expresssignature-method-presign.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Always add a x-amz-content-sha-256 for data integrity.[signRequest()](class-aws-signature-s3expresssignature-method-signrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Signs the specified request with an AWS signing protocol by using the
provided AWS account credentials and adding the required headers to the
request.

### Constants  [header link](class-aws-signature-s3expresssignature-constants.md)

#### AMZ\_CONTENT\_SHA256\_HEADER  [header link](class-aws-signature-signaturev4-constant-amz-content-sha256-header.md)

`
    public
        mixed
    AMZ_CONTENT_SHA256_HEADER
    = 'X-Amz-Content-Sha256'
`

#### ISO8601\_BASIC  [header link](class-aws-signature-signaturev4-constant-iso8601-basic.md)

`
    public
        mixed
    ISO8601_BASIC
    = 'Ymd\\THis\\Z'
`

#### UNSIGNED\_PAYLOAD  [header link](class-aws-signature-signaturev4-constant-unsigned-payload.md)

`
    public
        mixed
    UNSIGNED_PAYLOAD
    = 'UNSIGNED-PAYLOAD'
`

### Methods  [header link](class-aws-signature-s3expresssignature-methods.md)

#### \_\_construct()  [header link](class-aws-signature-signaturev4-method-construct.md)

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

#### convertPostToGet()  [header link](class-aws-signature-signaturev4-method-convertposttoget.md)

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

##### Tags  [header link](class-aws-signature-signaturev4-method-convertposttoget-tags.md)

throwsInvalidArgumentException

if the method is not POST

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### presign()  [header link](class-aws-signature-s3expresssignature-method-presign.md)

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

#### signRequest()  [header link](class-aws-signature-s3expresssignature-method-signrequest.md)

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
  - [Constants](class-aws-signature-s3expresssignature-toc-constants.md)
  - [Methods](class-aws-signature-s3expresssignature-toc-methods.md)
- Constants
  - [AMZ\_CONTENT\_SHA256\_HEADER](class-aws-signature-signaturev4-constant-amz-content-sha256-header.md)
  - [ISO8601\_BASIC](class-aws-signature-signaturev4-constant-iso8601-basic.md)
  - [UNSIGNED\_PAYLOAD](class-aws-signature-signaturev4-constant-unsigned-payload.md)
- Methods
  - [\_\_construct()](class-aws-signature-signaturev4-method-construct.md)
  - [convertPostToGet()](class-aws-signature-signaturev4-method-convertposttoget.md)
  - [presign()](class-aws-signature-s3expresssignature-method-presign.md)
  - [signRequest()](class-aws-signature-s3expresssignature-method-signrequest.md)

[Back To Top](class-aws-signature-s3expresssignature-top.md)
