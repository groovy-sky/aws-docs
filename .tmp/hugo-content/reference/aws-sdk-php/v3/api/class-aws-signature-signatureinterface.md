Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## SignatureInterface     in    - [Aws](package-aws.md)

Interface used to provide interchangeable strategies for signing requests
using the various AWS signature protocols.

### Table of Contents  [header link](class-aws-signature-signatureinterface-toc.md)

#### Methods  [header link](class-aws-signature-signatureinterface-toc-methods.md)

[presign()](class-aws-signature-signatureinterface-method-presign.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Create a pre-signed request.[signRequest()](class-aws-signature-signatureinterface-method-signrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Signs the specified request with an AWS signing protocol by using the
provided AWS account credentials and adding the required headers to the
request.

### Methods  [header link](class-aws-signature-signatureinterface-methods.md)

#### presign()  [header link](class-aws-signature-signatureinterface-method-presign.md)

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

#### signRequest()  [header link](class-aws-signature-signatureinterface-method-signrequest.md)

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
  - [Constants](class-aws-signature-signatureinterface-toc-constants.md)
  - [Methods](class-aws-signature-signatureinterface-toc-methods.md)
- Methods
  - [presign()](class-aws-signature-signatureinterface-method-presign.md)
  - [signRequest()](class-aws-signature-signatureinterface-method-signrequest.md)

[Back To Top](class-aws-signature-signatureinterface-top.md)
