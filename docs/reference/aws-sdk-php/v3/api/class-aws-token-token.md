Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## Token     extends BearerTokenIdentity   in package    - [Aws](package-aws.md)       implements  [TokenInterface](class-aws-token-tokeninterface.md), Serializable

Basic implementation of the AWS Token interface that allows callers to
pass in an AWS token in the constructor.

### Table of Contents  [header link](class-aws-token-token-toc.md)

#### Interfaces  [header link](class-aws-token-token-toc-interfaces.md)

[TokenInterface](class-aws-token-tokeninterface.md)Provides access to an AWS token used for accessing AWS servicesSerializable

#### Methods  [header link](class-aws-token-token-toc-methods.md)

[\_\_construct()](class-aws-token-token-method-construct.md)
: mixed Constructs a new basic token object, with the specified AWS
token[\_\_serialize()](class-aws-token-token-method-serialize.md)
: array<string\|int, mixed> [\_\_set\_state()](class-aws-token-token-method-set-state.md)
: mixed Sets the state of a token object[\_\_unserialize()](class-aws-token-token-method-unserialize.md)
: mixed Sets the state of this object from an array[getExpiration()](class-aws-token-token-method-getexpiration.md)
: int Get the UNIX timestamp in which the token will expire[getSource()](class-aws-token-token-method-getsource.md)
: string\|null [getToken()](class-aws-token-token-method-gettoken.md)
: string Returns the token this token object.[isExpired()](class-aws-token-token-method-isexpired.md)
: bool Check if the token are expired[serialize()](class-aws-token-token-method-serialize.md)
: string [toArray()](class-aws-token-token-method-toarray.md)
: array<string\|int, mixed> Converts the token to an associative array.[unserialize()](class-aws-token-token-method-unserialize.md)
: mixed Sets the state of the object from serialized json data

### Methods  [header link](class-aws-token-token-methods.md)

#### \_\_construct()  [header link](class-aws-token-token-method-construct.md)

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
: [TokenSource](class-aws-token-tokensource.md) \|null
= null

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

#### getExpiration()  [header link](class-aws-token-token-method-getexpiration.md)

Get the UNIX timestamp in which the token will expire

`
    public
                    getExpiration() : int`

##### Return values

int

#### getSource()  [header link](class-aws-token-token-method-getsource.md)

`
    public
                    getSource() : string|null`

##### Return values

string\|null

#### getToken()  [header link](class-aws-token-token-method-gettoken.md)

Returns the token this token object.

`
    public
                    getToken() : string`

##### Return values

string

#### isExpired()  [header link](class-aws-token-token-method-isexpired.md)

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
  - [Methods](class-aws-token-token-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-token-token-method-construct.md)
  - [\_\_serialize()](class-aws-token-token-method-serialize.md)
  - [\_\_set\_state()](class-aws-token-token-method-set-state.md)
  - [\_\_unserialize()](class-aws-token-token-method-unserialize.md)
  - [getExpiration()](class-aws-token-token-method-getexpiration.md)
  - [getSource()](class-aws-token-token-method-getsource.md)
  - [getToken()](class-aws-token-token-method-gettoken.md)
  - [isExpired()](class-aws-token-token-method-isexpired.md)
  - [serialize()](class-aws-token-token-method-serialize.md)
  - [toArray()](class-aws-token-token-method-toarray.md)
  - [unserialize()](class-aws-token-token-method-unserialize.md)

[Back To Top](class-aws-token-token-top.md)
