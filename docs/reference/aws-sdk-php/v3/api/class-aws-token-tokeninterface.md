Menu

- [Aws](namespace-aws.md)
- [Token](namespace-aws-token.md)

## TokenInterface     in    - [Aws](package-aws.md)

Provides access to an AWS token used for accessing AWS services

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html\#toc-methods)

[getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#method_getExpiration)
: int\|null Get the UNIX timestamp in which the token will expire[getToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#method_getToken)
: string Returns the token this token object.[isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#method_isExpired)
: bool Check if the token are expired[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#method_toArray)
: array<string\|int, mixed> Converts the token to an associative array.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html\#methods)

#### getExpiration()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html\#method_getExpiration)

Get the UNIX timestamp in which the token will expire

`
    public
                    getExpiration() : int|null`

##### Return values

int\|null

#### getToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html\#method_getToken)

Returns the token this token object.

`
    public
                    getToken() : string`

##### Return values

string

#### isExpired()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html\#method_isExpired)

Check if the token are expired

`
    public
                    isExpired() : bool`

##### Return values

bool

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html\#method_toArray)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#toc-methods)
- Methods
  - [getExpiration()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#method_getExpiration)
  - [getToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#method_getToken)
  - [isExpired()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#method_isExpired)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Token.TokenInterface.html#top)
