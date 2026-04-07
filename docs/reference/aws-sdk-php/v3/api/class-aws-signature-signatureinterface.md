Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## SignatureInterface     in    - [Aws](package-aws.md)

Interface used to provide interchangeable strategies for signing requests
using the various AWS signature protocols.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html\#toc-methods)

[presign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html#method_presign)
: [RequestInterface](class-psr-http-message-requestinterface.md)Create a pre-signed request.[signRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html#method_signRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)Signs the specified request with an AWS signing protocol by using the
provided AWS account credentials and adding the required headers to the
request.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html\#methods)

#### presign()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html\#method_presign)

Create a pre-signed request.

`
    public
                    presign(RequestInterface $request, CredentialsInterface $credentials, int|string|DateTimeInterface $expires[, array<string|int, mixed> $options = [] ]) : RequestInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to sign

$credentials
: [CredentialsInterface](class-aws-credentials-credentialsinterface.md)

Credentials used to sign

$expires
: int\|string\|DateTimeInterface

The time at which the URL should
expire. This can be a Unix timestamp, a PHP DateTime object, or a
string that can be evaluated by strtotime.

$options
: array<string\|int, mixed>
= \[\]

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### signRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html\#method_signRequest)

Signs the specified request with an AWS signing protocol by using the
provided AWS account credentials and adding the required headers to the
request.

`
    public
                    signRequest(RequestInterface $request, CredentialsInterface $credentials) : RequestInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to sign

$credentials
: [CredentialsInterface](class-aws-credentials-credentialsinterface.md)

Signing credentials

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)
—

Returns the modified request.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html#toc-methods)
- Methods
  - [presign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html#method_presign)
  - [signRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html#method_signRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.SignatureInterface.html#top)
