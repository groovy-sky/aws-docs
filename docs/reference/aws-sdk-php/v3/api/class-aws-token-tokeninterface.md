Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## TokenInterface     in    - [Aws](package-aws.md)

Provides access to an AWS token used for accessing AWS services

### Table of Contents  [header link](class-aws-token-tokeninterface-toc.md)

#### Methods  [header link](class-aws-token-tokeninterface-toc-methods.md)

[getExpiration()](class-aws-token-tokeninterface-method-getexpiration.md)
: int\|null Get the UNIX timestamp in which the token will expire[getToken()](class-aws-token-tokeninterface-method-gettoken.md)
: string Returns the token this token object.[isExpired()](class-aws-token-tokeninterface-method-isexpired.md)
: bool Check if the token are expired[toArray()](class-aws-token-tokeninterface-method-toarray.md)
: array<string\|int, mixed> Converts the token to an associative array.

### Methods  [header link](class-aws-token-tokeninterface-methods.md)

#### getExpiration()  [header link](class-aws-token-tokeninterface-method-getexpiration.md)

Get the UNIX timestamp in which the token will expire

`
    public
                    getExpiration() : int|null`

##### Return values

int\|null

#### getToken()  [header link](class-aws-token-tokeninterface-method-gettoken.md)

Returns the token this token object.

`
    public
                    getToken() : string`

##### Return values

string

#### isExpired()  [header link](class-aws-token-tokeninterface-method-isexpired.md)

Check if the token are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### toArray()  [header link](class-aws-token-tokeninterface-method-toarray.md)

Converts the token to an associative array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-token-tokeninterface-toc-constants.md)
  - [Methods](class-aws-token-tokeninterface-toc-methods.md)
- Methods
  - [getExpiration()](class-aws-token-tokeninterface-method-getexpiration.md)
  - [getToken()](class-aws-token-tokeninterface-method-gettoken.md)
  - [isExpired()](class-aws-token-tokeninterface-method-isexpired.md)
  - [toArray()](class-aws-token-tokeninterface-method-toarray.md)

[Back To Top](class-aws-token-tokeninterface-top.md)
