Menu

- [Aws](namespace-aws.md)
- [Auth](namespace-aws-auth.md)

## AuthSchemeResolverInterface     in    - [Aws](package-aws.md)

An AuthSchemeResolver object determines which auth scheme will be used for request signing.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html\#toc-methods)

[selectAuthScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html#method_selectAuthScheme)
: string\|null Selects an auth scheme for request signing.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html\#methods)

#### selectAuthScheme()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html\#method_selectAuthScheme)

Selects an auth scheme for request signing.

`
    public
                    selectAuthScheme(array<string|int, mixed> $authSchemes, array<string|int, mixed> $args) : string|null`

##### Parameters

$authSchemes
: array<string\|int, mixed>

a priority-ordered list of authentication schemes.

$args
: array<string\|int, mixed>

##### Return values

string\|null
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html#toc-methods)
- Methods
  - [selectAuthScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html#method_selectAuthScheme)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolverInterface.html#top)
