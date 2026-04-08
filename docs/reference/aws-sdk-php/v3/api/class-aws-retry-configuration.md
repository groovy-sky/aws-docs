Menu

- [Aws](namespace-aws.md)
- [Retry](namespace-aws-retry.md)

## Configuration        in package    - [Aws](package-aws.md)       implements  [ConfigurationInterface](class-aws-retry-configurationinterface.md)

### Table of Contents  [header link](class-aws-retry-configuration-toc.md)

#### Interfaces  [header link](class-aws-retry-configuration-toc-interfaces.md)

[ConfigurationInterface](class-aws-retry-configurationinterface.md)Provides access to retry configuration

#### Methods  [header link](class-aws-retry-configuration-toc-methods.md)

[\_\_construct()](class-aws-retry-configuration-method-construct.md)
: mixed [getMaxAttempts()](class-aws-retry-configuration-method-getmaxattempts.md)
: string Returns the maximum number of attempts that will be used for a request[getMode()](class-aws-retry-configuration-method-getmode.md)
: string Returns the retry mode. Available modes include 'legacy', 'standard', and
'adapative'.[toArray()](class-aws-retry-configuration-method-toarray.md)
: array<string\|int, mixed> Returns the configuration as an associative array

### Methods  [header link](class-aws-retry-configuration-methods.md)

#### \_\_construct()  [header link](class-aws-retry-configuration-method-construct.md)

`
    public
                    __construct([mixed $mode = 'legacy' ][, mixed $maxAttempts = 3 ]) : mixed`

##### Parameters

$mode
: mixed
= 'legacy'$maxAttempts
: mixed
= 3

#### getMaxAttempts()  [header link](class-aws-retry-configuration-method-getmaxattempts.md)

Returns the maximum number of attempts that will be used for a request

`
    public
                    getMaxAttempts() : string`

##### Return values

string

#### getMode()  [header link](class-aws-retry-configuration-method-getmode.md)

Returns the retry mode. Available modes include 'legacy', 'standard', and
'adapative'.

`
    public
                    getMode() : string`

##### Return values

string

#### toArray()  [header link](class-aws-retry-configuration-method-toarray.md)

Returns the configuration as an associative array

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-retry-configuration-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-retry-configuration-method-construct.md)
  - [getMaxAttempts()](class-aws-retry-configuration-method-getmaxattempts.md)
  - [getMode()](class-aws-retry-configuration-method-getmode.md)
  - [toArray()](class-aws-retry-configuration-method-toarray.md)

[Back To Top](class-aws-retry-configuration-top.md)
