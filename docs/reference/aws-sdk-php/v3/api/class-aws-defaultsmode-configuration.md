Menu

- [Aws](namespace-aws.md)
- [DefaultsMode](namespace-aws-defaultsmode.md)

## Configuration        in package    - [Aws](package-aws.md)       implements  [ConfigurationInterface](class-aws-defaultsmode-configurationinterface.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#toc-interfaces)

[ConfigurationInterface](class-aws-defaultsmode-configurationinterface.md)Provides access to defaultsMode configuration

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method___construct)
: mixed [getConnectTimeoutInMillis()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getConnectTimeoutInMillis)
: int Returns the connection timeout in milliseconds[getHttpRequestTimeoutInMillis()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getHttpRequestTimeoutInMillis)
: int Returns the http request timeout in milliseconds[getMode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getMode)
: string Returns the configuration mode. Available modes include 'legacy', 'standard', and
'adapative'.[getRetryMode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getRetryMode)
: mixed {@inheritdoc}[getS3UsEast1RegionalEndpoints()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getS3UsEast1RegionalEndpoints)
: bool Returns the s3 us-east-1 regional endpoints option[getStsRegionalEndpoints()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getStsRegionalEndpoints)
: bool Returns the sts regional endpoints option[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_toArray)
: array<string\|int, mixed> Returns the configuration as an associative array

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#method___construct)

`
    public
                    __construct([mixed $mode = 'legacy' ]) : mixed`

##### Parameters

$mode
: mixed
= 'legacy'

#### getConnectTimeoutInMillis()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#method_getConnectTimeoutInMillis)

Returns the connection timeout in milliseconds

`
    public
                    getConnectTimeoutInMillis() : int`

##### Return values

int

#### getHttpRequestTimeoutInMillis()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#method_getHttpRequestTimeoutInMillis)

Returns the http request timeout in milliseconds

`
    public
                    getHttpRequestTimeoutInMillis() : int`

##### Return values

int

#### getMode()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#method_getMode)

Returns the configuration mode. Available modes include 'legacy', 'standard', and
'adapative'.

`
    public
                    getMode() : string`

##### Return values

string

#### getRetryMode()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#method_getRetryMode)

{@inheritdoc}

`
    public
                    getRetryMode() : mixed`

#### getS3UsEast1RegionalEndpoints()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#method_getS3UsEast1RegionalEndpoints)

Returns the s3 us-east-1 regional endpoints option

`
    public
                    getS3UsEast1RegionalEndpoints() : bool`

##### Return values

bool

#### getStsRegionalEndpoints()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#method_getStsRegionalEndpoints)

Returns the sts regional endpoints option

`
    public
                    getStsRegionalEndpoints() : bool`

##### Return values

bool

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html\#method_toArray)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method___construct)
  - [getConnectTimeoutInMillis()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getConnectTimeoutInMillis)
  - [getHttpRequestTimeoutInMillis()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getHttpRequestTimeoutInMillis)
  - [getMode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getMode)
  - [getRetryMode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getRetryMode)
  - [getS3UsEast1RegionalEndpoints()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getS3UsEast1RegionalEndpoints)
  - [getStsRegionalEndpoints()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_getStsRegionalEndpoints)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DefaultsMode.Configuration.html#top)
