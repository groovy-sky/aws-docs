Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## SsoToken     extends [Token](class-aws-token-token.md)   in package    - [Aws](package-aws.md)

Token that comes from the SSO provider

### Table of Contents  [header link](class-aws-token-ssotoken-toc.md)

#### Methods  [header link](class-aws-token-ssotoken-toc-methods.md)

[\_\_construct()](class-aws-token-ssotoken-method-construct.md)
: mixed Constructs a new SSO token object, with the specified AWS
token[\_\_serialize()](class-aws-token-token-method-serialize.md)
: array<string\|int, mixed> [\_\_set\_state()](class-aws-token-token-method-set-state.md)
: mixed Sets the state of a token object[\_\_unserialize()](class-aws-token-token-method-unserialize.md)
: mixed Sets the state of this object from an array[fromTokenData()](class-aws-token-ssotoken-method-fromtokendata.md)
: [SsoToken](class-aws-token-ssotoken.md)Creates an instance of SsoToken from a token data.[getClientId()](class-aws-token-ssotoken-method-getclientid.md)
: string\|null [getClientSecret()](class-aws-token-ssotoken-method-getclientsecret.md)
: string\|null [getExpiration()](class-aws-token-token-method-getexpiration.md)
: int Get the UNIX timestamp in which the token will expire[getRefreshToken()](class-aws-token-ssotoken-method-getrefreshtoken.md)
: string\|null [getRegion()](class-aws-token-ssotoken-method-getregion.md)
: string\|null [getRegistrationExpiresAt()](class-aws-token-ssotoken-method-getregistrationexpiresat.md)
: int\|null [getSource()](class-aws-token-token-method-getsource.md)
: string\|null [getStartUrl()](class-aws-token-ssotoken-method-getstarturl.md)
: string\|null [getToken()](class-aws-token-token-method-gettoken.md)
: string Returns the token this token object.[isExpired()](class-aws-token-ssotoken-method-isexpired.md)
: bool Check if the token are expired[serialize()](class-aws-token-token-method-serialize.md)
: string [toArray()](class-aws-token-token-method-toarray.md)
: array<string\|int, mixed> Converts the token to an associative array.[unserialize()](class-aws-token-token-method-unserialize.md)
: mixed Sets the state of the object from serialized json data

### Methods  [header link](class-aws-token-ssotoken-methods.md)

#### \_\_construct()  [header link](class-aws-token-ssotoken-method-construct.md)

Constructs a new SSO token object, with the specified AWS
token

`
    public
                    __construct(string $token, int $expires[, int $refreshToken = null ][, int $clientId = null ][, int $clientSecret = null ][, int $registrationExpiresAt = null ][, int $region = null ][, int $startUrl = null ]) : mixed`

##### Parameters

$token
: string

Security token to use

$expires
: int

UNIX timestamp for when the token expires

$refreshToken
: int
= null

An opaque string returned by the sso-oidc service

$clientId
: int
= null

The client ID generated when performing the registration portion of the OIDC authorization flow

$clientSecret
: int
= null

The client secret generated when performing the registration portion of the OIDC authorization flow

$registrationExpiresAt
: int
= null

The expiration time of the client registration (clientId and clientSecret)

$region
: int
= null

The configured sso\_region for the profile that credentials are being resolved for

$startUrl
: int
= null

The configured sso\_start\_url for the profile that credentials are being resolved for

#### \_\_serialize()  [header link](class-aws-token-token-method-serialize.md)

`
    public
                    __serialize() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### \_\_set\_state()  [header link](class-aws-token-token-method-set-state.md)

Sets the state of a token object

`
    public
            static        __set_state(array<string|int, mixed> $state) : mixed`

##### Parameters

$state
: array<string\|int, mixed>

array containing 'token' and 'expires'

#### \_\_unserialize()  [header link](class-aws-token-token-method-unserialize.md)

Sets the state of this object from an array

`
    public
                    __unserialize(mixed $data) : mixed`

##### Parameters

$data
: mixed

#### fromTokenData()  [header link](class-aws-token-ssotoken-method-fromtokendata.md)

Creates an instance of SsoToken from a token data.

`
    public
            static        fromTokenData( $tokenData) : SsoToken`

##### Parameters

$tokenData
:

##### Return values

[SsoToken](class-aws-token-ssotoken.md)

#### getClientId()  [header link](class-aws-token-ssotoken-method-getclientid.md)

`
    public
                    getClientId() : string|null`

##### Return values

string\|null

#### getClientSecret()  [header link](class-aws-token-ssotoken-method-getclientsecret.md)

`
    public
                    getClientSecret() : string|null`

##### Return values

string\|null

#### getExpiration()  [header link](class-aws-token-token-method-getexpiration.md)

Get the UNIX timestamp in which the token will expire

`
    public
                    getExpiration() : int`

##### Return values

int

#### getRefreshToken()  [header link](class-aws-token-ssotoken-method-getrefreshtoken.md)

`
    public
                    getRefreshToken() : string|null`

##### Return values

string\|null

#### getRegion()  [header link](class-aws-token-ssotoken-method-getregion.md)

`
    public
                    getRegion() : string|null`

##### Return values

string\|null

#### getRegistrationExpiresAt()  [header link](class-aws-token-ssotoken-method-getregistrationexpiresat.md)

`
    public
                    getRegistrationExpiresAt() : int|null`

##### Return values

int\|null

#### getSource()  [header link](class-aws-token-token-method-getsource.md)

`
    public
                    getSource() : string|null`

##### Return values

string\|null

#### getStartUrl()  [header link](class-aws-token-ssotoken-method-getstarturl.md)

`
    public
                    getStartUrl() : string|null`

##### Return values

string\|null

#### getToken()  [header link](class-aws-token-token-method-gettoken.md)

Returns the token this token object.

`
    public
                    getToken() : string`

##### Return values

string

#### isExpired()  [header link](class-aws-token-ssotoken-method-isexpired.md)

Check if the token are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### serialize()  [header link](class-aws-token-token-method-serialize.md)

`
    public
                    serialize() : string`

##### Return values

string

#### toArray()  [header link](class-aws-token-token-method-toarray.md)

Converts the token to an associative array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### unserialize()  [header link](class-aws-token-token-method-unserialize.md)

Sets the state of the object from serialized json data

`
    public
                    unserialize(mixed $serialized) : mixed`

##### Parameters

$serialized
: mixed
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-token-ssotoken-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-token-ssotoken-method-construct.md)
  - [\_\_serialize()](class-aws-token-token-method-serialize.md)
  - [\_\_set\_state()](class-aws-token-token-method-set-state.md)
  - [\_\_unserialize()](class-aws-token-token-method-unserialize.md)
  - [fromTokenData()](class-aws-token-ssotoken-method-fromtokendata.md)
  - [getClientId()](class-aws-token-ssotoken-method-getclientid.md)
  - [getClientSecret()](class-aws-token-ssotoken-method-getclientsecret.md)
  - [getExpiration()](class-aws-token-token-method-getexpiration.md)
  - [getRefreshToken()](class-aws-token-ssotoken-method-getrefreshtoken.md)
  - [getRegion()](class-aws-token-ssotoken-method-getregion.md)
  - [getRegistrationExpiresAt()](class-aws-token-ssotoken-method-getregistrationexpiresat.md)
  - [getSource()](class-aws-token-token-method-getsource.md)
  - [getStartUrl()](class-aws-token-ssotoken-method-getstarturl.md)
  - [getToken()](class-aws-token-token-method-gettoken.md)
  - [isExpired()](class-aws-token-ssotoken-method-isexpired.md)
  - [serialize()](class-aws-token-token-method-serialize.md)
  - [toArray()](class-aws-token-token-method-toarray.md)
  - [unserialize()](class-aws-token-token-method-unserialize.md)

[Back To Top](class-aws-token-ssotoken-top.md)
