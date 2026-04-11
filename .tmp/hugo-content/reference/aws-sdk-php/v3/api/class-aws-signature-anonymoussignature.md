Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## AnonymousSignature        in package    - [Aws](package-aws.md)       implements  [SignatureInterface](class-aws-signature-signatureinterface.md)

Provides anonymous client access (does not sign requests).

### Table of Contents  [header link](class-aws-signature-anonymoussignature-toc.md)

#### Interfaces  [header link](class-aws-signature-anonymoussignature-toc-interfaces.md)

[SignatureInterface](class-aws-signature-signatureinterface.md)Interface used to provide interchangeable strategies for signing requests
using the various AWS signature protocols.

#### Methods  [header link](class-aws-signature-anonymoussignature-toc-methods.md)

[presign()](class-aws-signature-anonymoussignature-method-presign.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)/\\*\\* {@inheritdoc}[signRequest()](class-aws-signature-anonymoussignature-method-signrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)/\\*\\* {@inheritdoc}

### Methods  [header link](class-aws-signature-anonymoussignature-methods.md)

#### presign()  [header link](class-aws-signature-anonymoussignature-method-presign.md)

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

#### signRequest()  [header link](class-aws-signature-anonymoussignature-method-signrequest.md)

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
  - [Methods](class-aws-signature-anonymoussignature-toc-methods.md)
- Methods
  - [presign()](class-aws-signature-anonymoussignature-method-presign.md)
  - [signRequest()](class-aws-signature-anonymoussignature-method-signrequest.md)

[Back To Top](class-aws-signature-anonymoussignature-top.md)
