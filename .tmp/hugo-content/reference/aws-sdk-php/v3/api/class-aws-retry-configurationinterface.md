Menu

- [Aws](namespace-aws.md)
- [Retry](namespace-aws-retry.md)

## ConfigurationInterface     in    - [Aws](package-aws.md)

Provides access to retry configuration

### Table of Contents  [header link](class-aws-retry-configurationinterface-toc.md)

#### Methods  [header link](class-aws-retry-configurationinterface-toc-methods.md)

[getMaxAttempts()](class-aws-retry-configurationinterface-method-getmaxattempts.md)
: string Returns the maximum number of attempts that will be used for a request[getMode()](class-aws-retry-configurationinterface-method-getmode.md)
: string Returns the retry mode. Available modes include 'legacy', 'standard', and
'adapative'.[toArray()](class-aws-retry-configurationinterface-method-toarray.md)
: array<string\|int, mixed> Returns the configuration as an associative array

### Methods  [header link](class-aws-retry-configurationinterface-methods.md)

#### getMaxAttempts()  [header link](class-aws-retry-configurationinterface-method-getmaxattempts.md)

Returns the maximum number of attempts that will be used for a request

`
    public
                    getMaxAttempts() : string`

##### Return values

string

#### getMode()  [header link](class-aws-retry-configurationinterface-method-getmode.md)

Returns the retry mode. Available modes include 'legacy', 'standard', and
'adapative'.

`
    public
                    getMode() : string`

##### Return values

string

#### toArray()  [header link](class-aws-retry-configurationinterface-method-toarray.md)

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
  - [Constants](class-aws-retry-configurationinterface-toc-constants.md)
  - [Methods](class-aws-retry-configurationinterface-toc-methods.md)
- Methods
  - [getMaxAttempts()](class-aws-retry-configurationinterface-method-getmaxattempts.md)
  - [getMode()](class-aws-retry-configurationinterface-method-getmode.md)
  - [toArray()](class-aws-retry-configurationinterface-method-toarray.md)

[Back To Top](class-aws-retry-configurationinterface-top.md)
