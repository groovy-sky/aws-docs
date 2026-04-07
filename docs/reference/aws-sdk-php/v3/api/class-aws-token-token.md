Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## Token     extends BearerTokenIdentity   in package    - [Aws](package-aws.md)       implements  [TokenInterface](class-aws-token-tokeninterface.md), Serializable

Basic implementation of the AWS Token interface that allows callers to
pass in an AWS token in the constructor.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#toc-interfaces)

[TokenInterface](class-aws-token-tokeninterface.md)Provides access to an AWS token used for accessing AWS servicesSerializable

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___construct)
: mixed Constructs a new basic token object, with the specified AWS
token[\_\_serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___serialize)
: array<string\|int, mixed> [\_\_set\_state()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___set_state)
: mixed Sets the state of a token object[\_\_unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___unserialize)
: mixed Sets the state of this object from an array[getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getExpiration)
: int Get the UNIX timestamp in which the token will expire[getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getSource)
: string\|null [getToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getToken)
: string Returns the token this token object.[isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_isExpired)
: bool Check if the token are expired[serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_serialize)
: string [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_toArray)
: array<string\|int, mixed> Converts the token to an associative array.[unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_unserialize)
: mixed Sets the state of the object from serialized json data

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method___construct)

Constructs a new basic token object, with the specified AWS
token

`
    public
                    __construct(string $token[, int $expires = null ][, TokenSource|null $source = null ]) : mixed`

##### Parameters

$token
: string

Security token to use

$expires
: int
= null

UNIX timestamp for when the token expires

$source
: [TokenSource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenSource.html) \|null
= null

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

#### getExpiration()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_getExpiration)

Get the UNIX timestamp in which the token will expire

`
    public
                    getExpiration() : int`

##### Return values

int

#### getSource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_getSource)

`
    public
                    getSource() : string|null`

##### Return values

string\|null

#### getToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_getToken)

Returns the token this token object.

`
    public
                    getToken() : string`

##### Return values

string

#### isExpired()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html\#method_isExpired)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___construct)
  - [\_\_serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___serialize)
  - [\_\_set\_state()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___set_state)
  - [\_\_unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method___unserialize)
  - [getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getExpiration)
  - [getSource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getSource)
  - [getToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_getToken)
  - [isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_isExpired)
  - [serialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_serialize)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_toArray)
  - [unserialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#method_unserialize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.Token.html#top)
