Menu

- [Aws](namespace-aws.md)
- [Auth](namespace-aws-auth.md)

## AuthSchemeResolver        in package    - [Aws](package-aws.md)       implements  [AuthSchemeResolverInterface](class-aws-auth-authschemeresolverinterface.md)

Houses logic for selecting an auth scheme modeled in a service's \`auth\` trait.

The `auth` trait can be modeled either in a service's metadata, or at the operation level.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#toc-interfaces)

[AuthSchemeResolverInterface](class-aws-auth-authschemeresolverinterface.md)An AuthSchemeResolver object determines which auth scheme will be used for request signing.

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#toc-constants)

[UNSIGNED\_BODY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#constant_UNSIGNED_BODY)
= '-unsigned-body'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#method___construct)
: mixed [selectAuthScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#method_selectAuthScheme)
: string Accepts a priority-ordered list of auth schemes and an Identity
and selects the first compatible auth schemes, returning a normalized
signature version. For example, based on the default auth scheme mapping,
if \`aws.auth#sigv4\` is selected, \`v4\` will be returned.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#constants)

#### UNSIGNED\_BODY  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#constant_UNSIGNED_BODY)

`
    public
        mixed
    UNSIGNED_BODY
    = '-unsigned-body'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#method___construct)

`
    public
                    __construct(callable $credentialProvider[, callable|null $tokenProvider = null ][, array<string|int, mixed> $authSchemeMap = [] ]) : mixed`

##### Parameters

$credentialProvider
: callable$tokenProvider
: callable\|null
= null$authSchemeMap
: array<string\|int, mixed>
= \[\]

#### selectAuthScheme()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#method_selectAuthScheme)

Accepts a priority-ordered list of auth schemes and an Identity
and selects the first compatible auth schemes, returning a normalized
signature version. For example, based on the default auth scheme mapping,
if \`aws.auth#sigv4\` is selected, \`v4\` will be returned.

`
    public
                    selectAuthScheme(array<string|int, mixed> $authSchemes[, array<string|int, mixed> $args = [] ]) : string`

##### Parameters

$authSchemes
: array<string\|int, mixed>$args
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html\#method_selectAuthScheme\#tags)

throws[UnresolvedAuthSchemeException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.Exception.UnresolvedAuthSchemeException.html)

##### Return values

string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#toc-methods)
- Constants
  - [UNSIGNED\_BODY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#constant_UNSIGNED_BODY)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#method___construct)
  - [selectAuthScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#method_selectAuthScheme)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Auth.AuthSchemeResolver.html#top)
