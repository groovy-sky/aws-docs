Menu

- [Aws](namespace-aws.md)
- [Retry](namespace-aws-retry.md)

## ConfigurationInterface     in    - [Aws](package-aws.md)

Provides access to retry configuration

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html\#toc-methods)

[getMaxAttempts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#method_getMaxAttempts)
: string Returns the maximum number of attempts that will be used for a request[getMode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#method_getMode)
: string Returns the retry mode. Available modes include 'legacy', 'standard', and
'adapative'.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#method_toArray)
: array<string\|int, mixed> Returns the configuration as an associative array

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html\#methods)

#### getMaxAttempts()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html\#method_getMaxAttempts)

Returns the maximum number of attempts that will be used for a request

`
    public
                    getMaxAttempts() : string`

##### Return values

string

#### getMode()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html\#method_getMode)

Returns the retry mode. Available modes include 'legacy', 'standard', and
'adapative'.

`
    public
                    getMode() : string`

##### Return values

string

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html\#method_toArray)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#toc-methods)
- Methods
  - [getMaxAttempts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#method_getMaxAttempts)
  - [getMode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#method_getMode)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.ConfigurationInterface.html#top)
