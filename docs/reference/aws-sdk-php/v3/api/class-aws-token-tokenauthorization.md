Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## TokenAuthorization     in    - [Aws](package-aws.md)

Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html\#toc-methods)

[authorizeRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html#method_authorizeRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)Adds the specified token to a request by adding the required headers.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html\#methods)

#### authorizeRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html\#method_authorizeRequest)

Adds the specified token to a request by adding the required headers.

`
    public
                    authorizeRequest(RequestInterface $request, TokenInterface $token) : RequestInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to sign

$token
: [TokenInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html)

Token

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)
—

Returns the modified request.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html#toc-methods)
- Methods
  - [authorizeRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html#method_authorizeRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html#top)
