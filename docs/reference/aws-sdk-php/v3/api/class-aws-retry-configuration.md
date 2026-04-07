Menu

- [Aws](namespace-aws.md)
- [Retry](namespace-aws-retry.md)

## Configuration        in package    - [Aws](package-aws.md)       implements  [ConfigurationInterface](class-aws-retry-configurationinterface.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html\#toc-interfaces)

[ConfigurationInterface](class-aws-retry-configurationinterface.md)Provides access to retry configuration

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#method___construct)
: mixed [getMaxAttempts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#method_getMaxAttempts)
: string Returns the maximum number of attempts that will be used for a request[getMode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#method_getMode)
: string Returns the retry mode. Available modes include 'legacy', 'standard', and
'adapative'.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#method_toArray)
: array<string\|int, mixed> Returns the configuration as an associative array

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html\#method___construct)

`
    public
                    __construct([mixed $mode = 'legacy' ][, mixed $maxAttempts = 3 ]) : mixed`

##### Parameters

$mode
: mixed
= 'legacy'$maxAttempts
: mixed
= 3

#### getMaxAttempts()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html\#method_getMaxAttempts)

Returns the maximum number of attempts that will be used for a request

`
    public
                    getMaxAttempts() : string`

##### Return values

string

#### getMode()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html\#method_getMode)

Returns the retry mode. Available modes include 'legacy', 'standard', and
'adapative'.

`
    public
                    getMode() : string`

##### Return values

string

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html\#method_toArray)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#method___construct)
  - [getMaxAttempts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#method_getMaxAttempts)
  - [getMode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#method_getMode)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Retry.Configuration.html#top)
