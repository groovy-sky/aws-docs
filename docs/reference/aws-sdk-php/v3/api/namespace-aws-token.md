Menu

- [Aws](namespace-aws.md)

## Token

### Table of Contents  [header link](namespace-aws-token-toc.md)

#### Interfaces  [header link](namespace-aws-token-toc-interfaces.md)

[RefreshableTokenProviderInterface](class-aws-token-refreshabletokenproviderinterface.md)Provides access to an AWS token used for accessing AWS services[TokenAuthorization](class-aws-token-tokenauthorization.md)Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.[TokenInterface](class-aws-token-tokeninterface.md)Provides access to an AWS token used for accessing AWS services

#### Classes  [header link](namespace-aws-token-toc-classes.md)

[BearerTokenAuthorization](class-aws-token-bearertokenauthorization.md)Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.[BedrockTokenProvider](class-aws-token-bedrocktokenprovider.md)Token provider for Bedrock that sources bearer tokens from environment variables.[SsoToken](class-aws-token-ssotoken.md)Token that comes from the SSO provider[SsoTokenProvider](class-aws-token-ssotokenprovider.md)Token that comes from the SSO provider[Token](class-aws-token-token.md)Basic implementation of the AWS Token interface that allows callers to
pass in an AWS token in the constructor.[TokenProvider](class-aws-token-tokenprovider.md)Token providers are functions that accept no arguments and return a
promise that is fulfilled with an {@see \\Aws\\Token\\TokenInterface}
or rejected with an {@see \\Aws\\Exception\\TokenException}.

#### Traits  [header link](namespace-aws-token-toc-traits.md)

[ParsesIniTrait](class-aws-token-parsesinitrait.md)

#### Enums  [header link](namespace-aws-token-toc-enums.md)

[TokenSource](class-aws-token-tokensource.md)

```

```

×

**On this page**

- Table Of Contents
  - [Interfaces](namespace-aws-token-toc-interfaces.md)
  - [Classes](namespace-aws-token-toc-classes.md)
  - [Traits](namespace-aws-token-toc-traits.md)
  - [Enums](namespace-aws-token-toc-enums.md)

[Back To Top](namespace-aws-token-top.md)
