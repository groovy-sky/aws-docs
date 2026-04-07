Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## InstanceProfileProvider        in package    - [Aws](package-aws.md)

Credential provider that provides credentials from the EC2 metadata service.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#toc-constants)

[CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_CFG_EC2_METADATA_SERVICE_ENDPOINT)
= 'ec2\_metadata\_service\_endpoint' [CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT\_MODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_CFG_EC2_METADATA_SERVICE_ENDPOINT_MODE)
= 'ec2\_metadata\_service\_endpoint\_mode' [CFG\_EC2\_METADATA\_V1\_DISABLED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_CFG_EC2_METADATA_V1_DISABLED)
= 'ec2\_metadata\_v1\_disabled' [CRED\_PATH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_CRED_PATH)
= 'meta-data/iam/security-credentials/' [DEFAULT\_AWS\_EC2\_METADATA\_V1\_DISABLED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_AWS_EC2_METADATA_V1_DISABLED)
= false [DEFAULT\_METADATA\_SERVICE\_IPv4\_ENDPOINT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_METADATA_SERVICE_IPv4_ENDPOINT)
= 'http://169.254.169.254' [DEFAULT\_METADATA\_SERVICE\_IPv6\_ENDPOINT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_METADATA_SERVICE_IPv6_ENDPOINT)
= 'http://\[fd00:ec2::254\]' [DEFAULT\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_RETRIES)
= 3 [DEFAULT\_TIMEOUT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_TIMEOUT)
= 1.0 [DEFAULT\_TOKEN\_TTL\_SECONDS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_TOKEN_TTL_SECONDS)
= 21600 [ENDPOINT\_MODE\_IPv4](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENDPOINT_MODE_IPv4)
= 'IPv4' [ENDPOINT\_MODE\_IPv6](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENDPOINT_MODE_IPv6)
= 'IPv6' [ENV\_DISABLE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENV_DISABLE)
= 'AWS\_EC2\_METADATA\_DISABLED' [ENV\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENV_RETRIES)
= 'AWS\_METADATA\_SERVICE\_NUM\_ATTEMPTS' [ENV\_TIMEOUT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENV_TIMEOUT)
= 'AWS\_METADATA\_SERVICE\_TIMEOUT' [TOKEN\_PATH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_TOKEN_PATH)
= 'api/token'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#method___construct)
: mixed The constructor accepts the following options:[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#method___invoke)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Loads instance profile credentials.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constants)

#### CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_CFG_EC2_METADATA_SERVICE_ENDPOINT)

`
    public
        mixed
    CFG_EC2_METADATA_SERVICE_ENDPOINT
    = 'ec2_metadata_service_endpoint'
`

#### CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT\_MODE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_CFG_EC2_METADATA_SERVICE_ENDPOINT_MODE)

`
    public
        mixed
    CFG_EC2_METADATA_SERVICE_ENDPOINT_MODE
    = 'ec2_metadata_service_endpoint_mode'
`

#### CFG\_EC2\_METADATA\_V1\_DISABLED  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_CFG_EC2_METADATA_V1_DISABLED)

`
    public
        mixed
    CFG_EC2_METADATA_V1_DISABLED
    = 'ec2_metadata_v1_disabled'
`

#### CRED\_PATH  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_CRED_PATH)

`
    public
        mixed
    CRED_PATH
    = 'meta-data/iam/security-credentials/'
`

#### DEFAULT\_AWS\_EC2\_METADATA\_V1\_DISABLED  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_DEFAULT_AWS_EC2_METADATA_V1_DISABLED)

`
    public
        mixed
    DEFAULT_AWS_EC2_METADATA_V1_DISABLED
    = false
`

#### DEFAULT\_METADATA\_SERVICE\_IPv4\_ENDPOINT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_DEFAULT_METADATA_SERVICE_IPv4_ENDPOINT)

`
    public
        mixed
    DEFAULT_METADATA_SERVICE_IPv4_ENDPOINT
    = 'http://169.254.169.254'
`

#### DEFAULT\_METADATA\_SERVICE\_IPv6\_ENDPOINT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_DEFAULT_METADATA_SERVICE_IPv6_ENDPOINT)

`
    public
        mixed
    DEFAULT_METADATA_SERVICE_IPv6_ENDPOINT
    = 'http://[fd00:ec2::254]'
`

#### DEFAULT\_RETRIES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_DEFAULT_RETRIES)

`
    public
        mixed
    DEFAULT_RETRIES
    = 3
`

#### DEFAULT\_TIMEOUT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_DEFAULT_TIMEOUT)

`
    public
        mixed
    DEFAULT_TIMEOUT
    = 1.0
`

#### DEFAULT\_TOKEN\_TTL\_SECONDS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_DEFAULT_TOKEN_TTL_SECONDS)

`
    public
        mixed
    DEFAULT_TOKEN_TTL_SECONDS
    = 21600
`

#### ENDPOINT\_MODE\_IPv4  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_ENDPOINT_MODE_IPv4)

`
    public
        mixed
    ENDPOINT_MODE_IPv4
    = 'IPv4'
`

#### ENDPOINT\_MODE\_IPv6  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_ENDPOINT_MODE_IPv6)

`
    public
        mixed
    ENDPOINT_MODE_IPv6
    = 'IPv6'
`

#### ENV\_DISABLE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_ENV_DISABLE)

`
    public
        mixed
    ENV_DISABLE
    = 'AWS_EC2_METADATA_DISABLED'
`

#### ENV\_RETRIES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_ENV_RETRIES)

`
    public
        mixed
    ENV_RETRIES
    = 'AWS_METADATA_SERVICE_NUM_ATTEMPTS'
`

#### ENV\_TIMEOUT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_ENV_TIMEOUT)

`
    public
        mixed
    ENV_TIMEOUT
    = 'AWS_METADATA_SERVICE_TIMEOUT'
`

#### TOKEN\_PATH  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#constant_TOKEN_PATH)

`
    public
        mixed
    TOKEN_PATH
    = 'api/token'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#method___construct)

The constructor accepts the following options:

`
    public
                    __construct([array<string|int, mixed> $config = [] ]) : mixed`

- timeout: Connection timeout, in seconds.
- profile: Optional EC2 profile name, if known.
- retries: Optional number of retries to be attempted.
- ec2\_metadata\_v1\_disabled: Optional for disabling the fallback to IMDSv1.
- endpoint: Optional for overriding the default endpoint to be used for fetching credentials.
The value must contain a valid URI scheme. If the URI scheme is not https, it must
resolve to a loopback address.
- endpoint\_mode: Optional for overriding the default endpoint mode (IPv4\|IPv6) to be used for
resolving the default endpoint.
- use\_aws\_shared\_config\_files: Decides whether the shared config file should be considered when
using the ConfigurationResolver::resolve method.

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Configuration options.

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html\#method___invoke)

Loads instance profile credentials.

`
    public
                    __invoke([mixed $previousCredentials = null ]) : PromiseInterface`

##### Parameters

$previousCredentials
: mixed
= null

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#toc-methods)
- Constants
  - [CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_CFG_EC2_METADATA_SERVICE_ENDPOINT)
  - [CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT\_MODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_CFG_EC2_METADATA_SERVICE_ENDPOINT_MODE)
  - [CFG\_EC2\_METADATA\_V1\_DISABLED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_CFG_EC2_METADATA_V1_DISABLED)
  - [CRED\_PATH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_CRED_PATH)
  - [DEFAULT\_AWS\_EC2\_METADATA\_V1\_DISABLED](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_AWS_EC2_METADATA_V1_DISABLED)
  - [DEFAULT\_METADATA\_SERVICE\_IPv4\_ENDPOINT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_METADATA_SERVICE_IPv4_ENDPOINT)
  - [DEFAULT\_METADATA\_SERVICE\_IPv6\_ENDPOINT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_METADATA_SERVICE_IPv6_ENDPOINT)
  - [DEFAULT\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_RETRIES)
  - [DEFAULT\_TIMEOUT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_TIMEOUT)
  - [DEFAULT\_TOKEN\_TTL\_SECONDS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_DEFAULT_TOKEN_TTL_SECONDS)
  - [ENDPOINT\_MODE\_IPv4](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENDPOINT_MODE_IPv4)
  - [ENDPOINT\_MODE\_IPv6](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENDPOINT_MODE_IPv6)
  - [ENV\_DISABLE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENV_DISABLE)
  - [ENV\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENV_RETRIES)
  - [ENV\_TIMEOUT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_ENV_TIMEOUT)
  - [TOKEN\_PATH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#constant_TOKEN_PATH)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#method___invoke)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html#top)
