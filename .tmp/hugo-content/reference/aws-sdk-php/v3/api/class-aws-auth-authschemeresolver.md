Menu

- [Aws](namespace-aws.md)
- [Auth](namespace-aws-auth.md)

## AuthSchemeResolver        in package    - [Aws](package-aws.md)       implements  [AuthSchemeResolverInterface](class-aws-auth-authschemeresolverinterface.md)

Houses logic for selecting an auth scheme modeled in a service's \`auth\` trait.

The `auth` trait can be modeled either in a service's metadata, or at the operation level.

### Table of Contents  [header link](class-aws-auth-authschemeresolver-toc.md)

#### Interfaces  [header link](class-aws-auth-authschemeresolver-toc-interfaces.md)

[AuthSchemeResolverInterface](class-aws-auth-authschemeresolverinterface.md)An AuthSchemeResolver object determines which auth scheme will be used for request signing.

#### Constants  [header link](class-aws-auth-authschemeresolver-toc-constants.md)

[UNSIGNED\_BODY](class-aws-auth-authschemeresolver-constant-unsigned-body.md)
= '-unsigned-body'

#### Methods  [header link](class-aws-auth-authschemeresolver-toc-methods.md)

[\_\_construct()](class-aws-auth-authschemeresolver-method-construct.md)
: mixed [selectAuthScheme()](class-aws-auth-authschemeresolver-method-selectauthscheme.md)
: string Accepts a priority-ordered list of auth schemes and an Identity
and selects the first compatible auth schemes, returning a normalized
signature version. For example, based on the default auth scheme mapping,
if \`aws.auth#sigv4\` is selected, \`v4\` will be returned.

### Constants  [header link](class-aws-auth-authschemeresolver-constants.md)

#### UNSIGNED\_BODY  [header link](class-aws-auth-authschemeresolver-constant-unsigned-body.md)

`
    public
        mixed
    UNSIGNED_BODY
    = '-unsigned-body'
`

### Methods  [header link](class-aws-auth-authschemeresolver-methods.md)

#### \_\_construct()  [header link](class-aws-auth-authschemeresolver-method-construct.md)

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

#### selectAuthScheme()  [header link](class-aws-auth-authschemeresolver-method-selectauthscheme.md)

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

##### Tags  [header link](class-aws-auth-authschemeresolver-method-selectauthscheme-tags.md)

throws[UnresolvedAuthSchemeException](class-aws-auth-exception-unresolvedauthschemeexception.md)

##### Return values

string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-auth-authschemeresolver-toc-constants.md)
  - [Methods](class-aws-auth-authschemeresolver-toc-methods.md)
- Constants
  - [UNSIGNED\_BODY](class-aws-auth-authschemeresolver-constant-unsigned-body.md)
- Methods
  - [\_\_construct()](class-aws-auth-authschemeresolver-method-construct.md)
  - [selectAuthScheme()](class-aws-auth-authschemeresolver-method-selectauthscheme.md)

[Back To Top](class-aws-auth-authschemeresolver-top.md)
