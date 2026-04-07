Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## BearerTokenAuthorization        in package    - [Aws](package-aws.md)       implements  [TokenAuthorization](class-aws-token-tokenauthorization.md)

Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html\#toc-interfaces)

[TokenAuthorization](class-aws-token-tokenauthorization.md)Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html\#toc-methods)

[authorizeRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html#method_authorizeRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)Adds the specified token to a request by adding the required headers.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html\#methods)

#### authorizeRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html\#method_authorizeRequest)

Adds the specified token to a request by adding the required headers.

`
    public
                    authorizeRequest(RequestInterface $request, TokenInterface $token) : RequestInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to sign

$token
: [TokenInterface](class-aws-token-tokeninterface.md)

Token

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)
—

Returns the modified request.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html#toc-methods)
- Methods
  - [authorizeRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html#method_authorizeRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html#top)
