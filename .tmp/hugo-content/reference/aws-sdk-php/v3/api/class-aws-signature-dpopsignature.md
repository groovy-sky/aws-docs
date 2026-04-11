Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## DpopSignature        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-signature-dpopsignature-toc.md)

#### Constants  [header link](class-aws-signature-dpopsignature-toc-constants.md)

[ALLOW\_LISTED\_SERVICES](class-aws-signature-dpopsignature-constant-allow-listed-services.md)
= \['signin' => true\]

#### Methods  [header link](class-aws-signature-dpopsignature-toc-methods.md)

[\_\_construct()](class-aws-signature-dpopsignature-method-construct.md)
: mixed Creates a new DpopSignature instance for the specified service[signRequest()](class-aws-signature-dpopsignature-method-signrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Signs an HTTP request with a DPoP header

### Constants  [header link](class-aws-signature-dpopsignature-constants.md)

#### ALLOW\_LISTED\_SERVICES  [header link](class-aws-signature-dpopsignature-constant-allow-listed-services.md)

`
    public
        mixed
    ALLOW_LISTED_SERVICES
    = ['signin' => true]
`

### Methods  [header link](class-aws-signature-dpopsignature-methods.md)

#### \_\_construct()  [header link](class-aws-signature-dpopsignature-method-construct.md)

Creates a new DpopSignature instance for the specified service

`
    public
                    __construct(string $serviceName) : mixed`

##### Parameters

$serviceName
: string

The name of the AWS service (must be in the allow list)

##### Tags  [header link](class-aws-signature-dpopsignature-method-construct-tags.md)

throwsRuntimeException

If the OpenSSL extension is not loaded

throwsInvalidArgumentException

If the service is not in the allow list

#### signRequest()  [header link](class-aws-signature-dpopsignature-method-signrequest.md)

Signs an HTTP request with a DPoP header

`
    public
                    signRequest(RequestInterface $request, OpenSSLAsymmetricKey $key) : RequestInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

The HTTP request to sign

$key
: OpenSSLAsymmetricKey

The private key for signing

##### Tags  [header link](class-aws-signature-dpopsignature-method-signrequest-tags.md)

throwsRuntimeException\|Exception

If signature generation fails

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)
—

The request with the DPoP header added

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-signature-dpopsignature-toc-constants.md)
  - [Methods](class-aws-signature-dpopsignature-toc-methods.md)
- Constants
  - [ALLOW\_LISTED\_SERVICES](class-aws-signature-dpopsignature-constant-allow-listed-services.md)
- Methods
  - [\_\_construct()](class-aws-signature-dpopsignature-method-construct.md)
  - [signRequest()](class-aws-signature-dpopsignature-method-signrequest.md)

[Back To Top](class-aws-signature-dpopsignature-top.md)
