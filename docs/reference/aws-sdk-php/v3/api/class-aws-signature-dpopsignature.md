Menu

- [Aws](namespace-aws.md)
- [Signature](namespace-aws-signature.md)

## DpopSignature        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#toc-constants)

[ALLOW\_LISTED\_SERVICES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#constant_ALLOW_LISTED_SERVICES)
= \['signin' => true\]

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#method___construct)
: mixed Creates a new DpopSignature instance for the specified service[signRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#method_signRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)Signs an HTTP request with a DPoP header

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#constants)

#### ALLOW\_LISTED\_SERVICES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#constant_ALLOW_LISTED_SERVICES)

`
    public
        mixed
    ALLOW_LISTED_SERVICES
    = ['signin' => true]
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#method___construct)

Creates a new DpopSignature instance for the specified service

`
    public
                    __construct(string $serviceName) : mixed`

##### Parameters

$serviceName
: string

The name of the AWS service (must be in the allow list)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#method___construct\#tags)

throwsRuntimeException

If the OpenSSL extension is not loaded

throwsInvalidArgumentException

If the service is not in the allow list

#### signRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#method_signRequest)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html\#method_signRequest\#tags)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#toc-methods)
- Constants
  - [ALLOW\_LISTED\_SERVICES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#constant_ALLOW_LISTED_SERVICES)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#method___construct)
  - [signRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#method_signRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Signature.DpopSignature.html#top)
