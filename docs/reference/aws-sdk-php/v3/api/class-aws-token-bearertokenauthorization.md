Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## BearerTokenAuthorization        in package    - [Aws](package-aws.md)       implements  [TokenAuthorization](class-aws-token-tokenauthorization.md)

Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.

### Table of Contents  [header link](class-aws-token-bearertokenauthorization-toc.md)

#### Interfaces  [header link](class-aws-token-bearertokenauthorization-toc-interfaces.md)

[TokenAuthorization](class-aws-token-tokenauthorization.md)Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.

#### Methods  [header link](class-aws-token-bearertokenauthorization-toc-methods.md)

[authorizeRequest()](class-aws-token-bearertokenauthorization-method-authorizerequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Adds the specified token to a request by adding the required headers.

### Methods  [header link](class-aws-token-bearertokenauthorization-methods.md)

#### authorizeRequest()  [header link](class-aws-token-bearertokenauthorization-method-authorizerequest.md)

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
  - [Methods](class-aws-token-bearertokenauthorization-toc-methods.md)
- Methods
  - [authorizeRequest()](class-aws-token-bearertokenauthorization-method-authorizerequest.md)

[Back To Top](class-aws-token-bearertokenauthorization-top.md)
