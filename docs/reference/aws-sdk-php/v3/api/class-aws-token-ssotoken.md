Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## SsoToken     extends [Token](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html)   in package    - [Aws](package-aws.md)

Token that comes from the SSO provider

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method___construct)
: mixed Constructs a new SSO token object, with the specified AWS
token[\_\_serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___serialize)
: array<string\|int, mixed> [\_\_set\_state()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___set_state)
: mixed Sets the state of a token object[\_\_unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___unserialize)
: mixed Sets the state of this object from an array[fromTokenData()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_fromTokenData)
: [SsoToken](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html)Creates an instance of SsoToken from a token data.[getClientId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getClientId)
: string\|null [getClientSecret()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getClientSecret)
: string\|null [getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getExpiration)
: int Get the UNIX timestamp in which the token will expire[getRefreshToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getRefreshToken)
: string\|null [getRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getRegion)
: string\|null [getRegistrationExpiresAt()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getRegistrationExpiresAt)
: int\|null [getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getSource)
: string\|null [getStartUrl()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getStartUrl)
: string\|null [getToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getToken)
: string Returns the token this token object.[isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_isExpired)
: bool Check if the token are expired[serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_serialize)
: string [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_toArray)
: array<string\|int, mixed> Converts the token to an associative array.[unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_unserialize)
: mixed Sets the state of the object from serialized json data

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method___construct)

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

#### \_\_serialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method___serialize)

`
    public
                    __serialize() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### \_\_set\_state()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method___set_state)

Sets the state of a token object

`
    public
            static        __set_state(array<string|int, mixed> $state) : mixed`

##### Parameters

$state
: array<string\|int, mixed>

array containing 'token' and 'expires'

#### \_\_unserialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method___unserialize)

Sets the state of this object from an array

`
    public
                    __unserialize(mixed $data) : mixed`

##### Parameters

$data
: mixed

#### fromTokenData()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method_fromTokenData)

Creates an instance of SsoToken from a token data.

`
    public
            static        fromTokenData( $tokenData) : SsoToken`

##### Parameters

$tokenData
:

##### Return values

[SsoToken](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html)

#### getClientId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method_getClientId)

`
    public
                    getClientId() : string|null`

##### Return values

string\|null

#### getClientSecret()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method_getClientSecret)

`
    public
                    getClientSecret() : string|null`

##### Return values

string\|null

#### getExpiration()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_getExpiration)

Get the UNIX timestamp in which the token will expire

`
    public
                    getExpiration() : int`

##### Return values

int

#### getRefreshToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method_getRefreshToken)

`
    public
                    getRefreshToken() : string|null`

##### Return values

string\|null

#### getRegion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method_getRegion)

`
    public
                    getRegion() : string|null`

##### Return values

string\|null

#### getRegistrationExpiresAt()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method_getRegistrationExpiresAt)

`
    public
                    getRegistrationExpiresAt() : int|null`

##### Return values

int\|null

#### getSource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_getSource)

`
    public
                    getSource() : string|null`

##### Return values

string\|null

#### getStartUrl()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method_getStartUrl)

`
    public
                    getStartUrl() : string|null`

##### Return values

string\|null

#### getToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_getToken)

Returns the token this token object.

`
    public
                    getToken() : string`

##### Return values

string

#### isExpired()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html\#method_isExpired)

Check if the token are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### serialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_serialize)

`
    public
                    serialize() : string`

##### Return values

string

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_toArray)

Converts the token to an associative array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### unserialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_unserialize)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method___construct)
  - [\_\_serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___serialize)
  - [\_\_set\_state()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___set_state)
  - [\_\_unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___unserialize)
  - [fromTokenData()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_fromTokenData)
  - [getClientId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getClientId)
  - [getClientSecret()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getClientSecret)
  - [getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getExpiration)
  - [getRefreshToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getRefreshToken)
  - [getRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getRegion)
  - [getRegistrationExpiresAt()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getRegistrationExpiresAt)
  - [getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getSource)
  - [getStartUrl()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_getStartUrl)
  - [getToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getToken)
  - [isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#method_isExpired)
  - [serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_serialize)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_toArray)
  - [unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_unserialize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.SsoToken.html#top)
