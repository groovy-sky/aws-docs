Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## AnonymousSignature        in package    - [Aws](package-aws.md)       implements  [SignatureInterface](class-aws-signature-signatureinterface.md)

Provides anonymous client access (does not sign requests).

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html\#toc-interfaces)

[SignatureInterface](class-aws-signature-signatureinterface.md)Interface used to provide interchangeable strategies for signing requests
using the various AWS signature protocols.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html\#toc-methods)

[presign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html#method_presign)
: [RequestInterface](class-psr-http-message-requestinterface.md)/\\*\\* {@inheritdoc}[signRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html#method_signRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)/\\*\\* {@inheritdoc}

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html\#methods)

#### presign()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html\#method_presign)

/\\*\\* {@inheritdoc}

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

#### signRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html\#method_signRequest)

/\\*\\* {@inheritdoc}

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html#toc-methods)
- Methods
  - [presign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html#method_presign)
  - [signRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html#method_signRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.AnonymousSignature.html#top)
