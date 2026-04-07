Menu

- [Aws](namespace-aws.md)

## Token

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html\#toc-interfaces)

[RefreshableTokenProviderInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.RefreshableTokenProviderInterface.html)Provides access to an AWS token used for accessing AWS services[TokenAuthorization](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenAuthorization.html)Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.[TokenInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html)Provides access to an AWS token used for accessing AWS services

#### Classes  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html\#toc-classes)

[BearerTokenAuthorization](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BearerTokenAuthorization.html)Interface used to provide interchangeable strategies for adding authorization
to requests using the various AWS signature protocols.[BedrockTokenProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.BedrockTokenProvider.html)Token provider for Bedrock that sources bearer tokens from environment variables.[SsoToken](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html)Token that comes from the SSO provider[SsoTokenProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoTokenProvider.html)Token that comes from the SSO provider[Token](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html)Basic implementation of the AWS Token interface that allows callers to
pass in an AWS token in the constructor.[TokenProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenProvider.html)Token providers are functions that accept no arguments and return a
promise that is fulfilled with an {@see \\Aws\\Token\\TokenInterface}
or rejected with an {@see \\Aws\\Exception\\TokenException}.

#### Traits  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html\#toc-traits)

[ParsesIniTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.ParsesIniTrait.html)

#### Enums  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html\#toc-enums)

[TokenSource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenSource.html)

```

```

×

**On this page**

- Table Of Contents
  - [Interfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html#toc-interfaces)
  - [Classes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html#toc-classes)
  - [Traits](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html#toc-traits)
  - [Enums](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html#toc-enums)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.token.html#top)
