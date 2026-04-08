Menu

- [Aws](namespace-aws.md)
- [Auth](namespace-aws-auth.md)

## AuthSchemeResolverInterface     in    - [Aws](package-aws.md)

An AuthSchemeResolver object determines which auth scheme will be used for request signing.

### Table of Contents  [header link](class-aws-auth-authschemeresolverinterface-toc.md)

#### Methods  [header link](class-aws-auth-authschemeresolverinterface-toc-methods.md)

[selectAuthScheme()](class-aws-auth-authschemeresolverinterface-method-selectauthscheme.md)
: string\|null Selects an auth scheme for request signing.

### Methods  [header link](class-aws-auth-authschemeresolverinterface-methods.md)

#### selectAuthScheme()  [header link](class-aws-auth-authschemeresolverinterface-method-selectauthscheme.md)

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
  - [Constants](class-aws-auth-authschemeresolverinterface-toc-constants.md)
  - [Methods](class-aws-auth-authschemeresolverinterface-toc-methods.md)
- Methods
  - [selectAuthScheme()](class-aws-auth-authschemeresolverinterface-method-selectauthscheme.md)

[Back To Top](class-aws-auth-authschemeresolverinterface-top.md)
