Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## TokenAuthorization     in    - [Aws](package-aws.md)

Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.

### Table of Contents  [header link](class-aws-token-tokenauthorization-toc.md)

#### Methods  [header link](class-aws-token-tokenauthorization-toc-methods.md)

[authorizeRequest()](class-aws-token-tokenauthorization-method-authorizerequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Adds the specified token to a request by adding the required headers.

### Methods  [header link](class-aws-token-tokenauthorization-methods.md)

#### authorizeRequest()  [header link](class-aws-token-tokenauthorization-method-authorizerequest.md)

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
  - [Constants](class-aws-token-tokenauthorization-toc-constants.md)
  - [Methods](class-aws-token-tokenauthorization-toc-methods.md)
- Methods
  - [authorizeRequest()](class-aws-token-tokenauthorization-method-authorizerequest.md)

[Back To Top](class-aws-token-tokenauthorization-top.md)
