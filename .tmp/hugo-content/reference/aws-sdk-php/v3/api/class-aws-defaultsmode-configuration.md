Menu

- [Aws](namespace-aws.md)
- [DefaultsMode](namespace-aws-defaultsmode.md)

## Configuration        in package    - [Aws](package-aws.md)       implements  [ConfigurationInterface](class-aws-defaultsmode-configurationinterface.md)

### Table of Contents  [header link](class-aws-defaultsmode-configuration-toc.md)

#### Interfaces  [header link](class-aws-defaultsmode-configuration-toc-interfaces.md)

[ConfigurationInterface](class-aws-defaultsmode-configurationinterface.md)Provides access to defaultsMode configuration

#### Methods  [header link](class-aws-defaultsmode-configuration-toc-methods.md)

[\_\_construct()](class-aws-defaultsmode-configuration-method-construct.md)
: mixed [getConnectTimeoutInMillis()](class-aws-defaultsmode-configuration-method-getconnecttimeoutinmillis.md)
: int Returns the connection timeout in milliseconds[getHttpRequestTimeoutInMillis()](class-aws-defaultsmode-configuration-method-gethttprequesttimeoutinmillis.md)
: int Returns the http request timeout in milliseconds[getMode()](class-aws-defaultsmode-configuration-method-getmode.md)
: string Returns the configuration mode. Available modes include 'legacy', 'standard', and
'adapative'.[getRetryMode()](class-aws-defaultsmode-configuration-method-getretrymode.md)
: mixed {@inheritdoc}[getS3UsEast1RegionalEndpoints()](class-aws-defaultsmode-configuration-method-gets3useast1regionalendpoints.md)
: bool Returns the s3 us-east-1 regional endpoints option[getStsRegionalEndpoints()](class-aws-defaultsmode-configuration-method-getstsregionalendpoints.md)
: bool Returns the sts regional endpoints option[toArray()](class-aws-defaultsmode-configuration-method-toarray.md)
: array<string\|int, mixed> Returns the configuration as an associative array

### Methods  [header link](class-aws-defaultsmode-configuration-methods.md)

#### \_\_construct()  [header link](class-aws-defaultsmode-configuration-method-construct.md)

`
    public
                    __construct([mixed $mode = 'legacy' ]) : mixed`

##### Parameters

$mode
: mixed
= 'legacy'

#### getConnectTimeoutInMillis()  [header link](class-aws-defaultsmode-configuration-method-getconnecttimeoutinmillis.md)

Returns the connection timeout in milliseconds

`
    public
                    getConnectTimeoutInMillis() : int`

##### Return values

int

#### getHttpRequestTimeoutInMillis()  [header link](class-aws-defaultsmode-configuration-method-gethttprequesttimeoutinmillis.md)

Returns the http request timeout in milliseconds

`
    public
                    getHttpRequestTimeoutInMillis() : int`

##### Return values

int

#### getMode()  [header link](class-aws-defaultsmode-configuration-method-getmode.md)

Returns the configuration mode. Available modes include 'legacy', 'standard', and
'adapative'.

`
    public
                    getMode() : string`

##### Return values

string

#### getRetryMode()  [header link](class-aws-defaultsmode-configuration-method-getretrymode.md)

{@inheritdoc}

`
    public
                    getRetryMode() : mixed`

#### getS3UsEast1RegionalEndpoints()  [header link](class-aws-defaultsmode-configuration-method-gets3useast1regionalendpoints.md)

Returns the s3 us-east-1 regional endpoints option

`
    public
                    getS3UsEast1RegionalEndpoints() : bool`

##### Return values

bool

#### getStsRegionalEndpoints()  [header link](class-aws-defaultsmode-configuration-method-getstsregionalendpoints.md)

Returns the sts regional endpoints option

`
    public
                    getStsRegionalEndpoints() : bool`

##### Return values

bool

#### toArray()  [header link](class-aws-defaultsmode-configuration-method-toarray.md)

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
  - [Methods](class-aws-defaultsmode-configuration-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-defaultsmode-configuration-method-construct.md)
  - [getConnectTimeoutInMillis()](class-aws-defaultsmode-configuration-method-getconnecttimeoutinmillis.md)
  - [getHttpRequestTimeoutInMillis()](class-aws-defaultsmode-configuration-method-gethttprequesttimeoutinmillis.md)
  - [getMode()](class-aws-defaultsmode-configuration-method-getmode.md)
  - [getRetryMode()](class-aws-defaultsmode-configuration-method-getretrymode.md)
  - [getS3UsEast1RegionalEndpoints()](class-aws-defaultsmode-configuration-method-gets3useast1regionalendpoints.md)
  - [getStsRegionalEndpoints()](class-aws-defaultsmode-configuration-method-getstsregionalendpoints.md)
  - [toArray()](class-aws-defaultsmode-configuration-method-toarray.md)

[Back To Top](class-aws-defaultsmode-configuration-top.md)
