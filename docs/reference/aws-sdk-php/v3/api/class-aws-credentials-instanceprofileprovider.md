Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## InstanceProfileProvider        in package    - [Aws](package-aws.md)

Credential provider that provides credentials from the EC2 metadata service.

### Table of Contents  [header link](class-aws-credentials-instanceprofileprovider-toc.md)

#### Constants  [header link](class-aws-credentials-instanceprofileprovider-toc-constants.md)

[CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-service-endpoint.md)
= 'ec2\_metadata\_service\_endpoint' [CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT\_MODE](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-service-endpoint-mode.md)
= 'ec2\_metadata\_service\_endpoint\_mode' [CFG\_EC2\_METADATA\_V1\_DISABLED](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-v1-disabled.md)
= 'ec2\_metadata\_v1\_disabled' [CRED\_PATH](class-aws-credentials-instanceprofileprovider-constant-cred-path.md)
= 'meta-data/iam/security-credentials/' [DEFAULT\_AWS\_EC2\_METADATA\_V1\_DISABLED](class-aws-credentials-instanceprofileprovider-constant-default-aws-ec2-metadata-v1-disabled.md)
= false [DEFAULT\_METADATA\_SERVICE\_IPv4\_ENDPOINT](class-aws-credentials-instanceprofileprovider-constant-default-metadata-service-ipv4-endpoint.md)
= 'http://169.254.169.254' [DEFAULT\_METADATA\_SERVICE\_IPv6\_ENDPOINT](class-aws-credentials-instanceprofileprovider-constant-default-metadata-service-ipv6-endpoint.md)
= 'http://\[fd00:ec2::254\]' [DEFAULT\_RETRIES](class-aws-credentials-instanceprofileprovider-constant-default-retries.md)
= 3 [DEFAULT\_TIMEOUT](class-aws-credentials-instanceprofileprovider-constant-default-timeout.md)
= 1.0 [DEFAULT\_TOKEN\_TTL\_SECONDS](class-aws-credentials-instanceprofileprovider-constant-default-token-ttl-seconds.md)
= 21600 [ENDPOINT\_MODE\_IPv4](class-aws-credentials-instanceprofileprovider-constant-endpoint-mode-ipv4.md)
= 'IPv4' [ENDPOINT\_MODE\_IPv6](class-aws-credentials-instanceprofileprovider-constant-endpoint-mode-ipv6.md)
= 'IPv6' [ENV\_DISABLE](class-aws-credentials-instanceprofileprovider-constant-env-disable.md)
= 'AWS\_EC2\_METADATA\_DISABLED' [ENV\_RETRIES](class-aws-credentials-instanceprofileprovider-constant-env-retries.md)
= 'AWS\_METADATA\_SERVICE\_NUM\_ATTEMPTS' [ENV\_TIMEOUT](class-aws-credentials-instanceprofileprovider-constant-env-timeout.md)
= 'AWS\_METADATA\_SERVICE\_TIMEOUT' [TOKEN\_PATH](class-aws-credentials-instanceprofileprovider-constant-token-path.md)
= 'api/token'

#### Methods  [header link](class-aws-credentials-instanceprofileprovider-toc-methods.md)

[\_\_construct()](class-aws-credentials-instanceprofileprovider-method-construct.md)
: mixed The constructor accepts the following options:[\_\_invoke()](class-aws-credentials-instanceprofileprovider-method-invoke.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Loads instance profile credentials.

### Constants  [header link](class-aws-credentials-instanceprofileprovider-constants.md)

#### CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT  [header link](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-service-endpoint.md)

`
    public
        mixed
    CFG_EC2_METADATA_SERVICE_ENDPOINT
    = 'ec2_metadata_service_endpoint'
`

#### CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT\_MODE  [header link](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-service-endpoint-mode.md)

`
    public
        mixed
    CFG_EC2_METADATA_SERVICE_ENDPOINT_MODE
    = 'ec2_metadata_service_endpoint_mode'
`

#### CFG\_EC2\_METADATA\_V1\_DISABLED  [header link](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-v1-disabled.md)

`
    public
        mixed
    CFG_EC2_METADATA_V1_DISABLED
    = 'ec2_metadata_v1_disabled'
`

#### CRED\_PATH  [header link](class-aws-credentials-instanceprofileprovider-constant-cred-path.md)

`
    public
        mixed
    CRED_PATH
    = 'meta-data/iam/security-credentials/'
`

#### DEFAULT\_AWS\_EC2\_METADATA\_V1\_DISABLED  [header link](class-aws-credentials-instanceprofileprovider-constant-default-aws-ec2-metadata-v1-disabled.md)

`
    public
        mixed
    DEFAULT_AWS_EC2_METADATA_V1_DISABLED
    = false
`

#### DEFAULT\_METADATA\_SERVICE\_IPv4\_ENDPOINT  [header link](class-aws-credentials-instanceprofileprovider-constant-default-metadata-service-ipv4-endpoint.md)

`
    public
        mixed
    DEFAULT_METADATA_SERVICE_IPv4_ENDPOINT
    = 'http://169.254.169.254'
`

#### DEFAULT\_METADATA\_SERVICE\_IPv6\_ENDPOINT  [header link](class-aws-credentials-instanceprofileprovider-constant-default-metadata-service-ipv6-endpoint.md)

`
    public
        mixed
    DEFAULT_METADATA_SERVICE_IPv6_ENDPOINT
    = 'http://[fd00:ec2::254]'
`

#### DEFAULT\_RETRIES  [header link](class-aws-credentials-instanceprofileprovider-constant-default-retries.md)

`
    public
        mixed
    DEFAULT_RETRIES
    = 3
`

#### DEFAULT\_TIMEOUT  [header link](class-aws-credentials-instanceprofileprovider-constant-default-timeout.md)

`
    public
        mixed
    DEFAULT_TIMEOUT
    = 1.0
`

#### DEFAULT\_TOKEN\_TTL\_SECONDS  [header link](class-aws-credentials-instanceprofileprovider-constant-default-token-ttl-seconds.md)

`
    public
        mixed
    DEFAULT_TOKEN_TTL_SECONDS
    = 21600
`

#### ENDPOINT\_MODE\_IPv4  [header link](class-aws-credentials-instanceprofileprovider-constant-endpoint-mode-ipv4.md)

`
    public
        mixed
    ENDPOINT_MODE_IPv4
    = 'IPv4'
`

#### ENDPOINT\_MODE\_IPv6  [header link](class-aws-credentials-instanceprofileprovider-constant-endpoint-mode-ipv6.md)

`
    public
        mixed
    ENDPOINT_MODE_IPv6
    = 'IPv6'
`

#### ENV\_DISABLE  [header link](class-aws-credentials-instanceprofileprovider-constant-env-disable.md)

`
    public
        mixed
    ENV_DISABLE
    = 'AWS_EC2_METADATA_DISABLED'
`

#### ENV\_RETRIES  [header link](class-aws-credentials-instanceprofileprovider-constant-env-retries.md)

`
    public
        mixed
    ENV_RETRIES
    = 'AWS_METADATA_SERVICE_NUM_ATTEMPTS'
`

#### ENV\_TIMEOUT  [header link](class-aws-credentials-instanceprofileprovider-constant-env-timeout.md)

`
    public
        mixed
    ENV_TIMEOUT
    = 'AWS_METADATA_SERVICE_TIMEOUT'
`

#### TOKEN\_PATH  [header link](class-aws-credentials-instanceprofileprovider-constant-token-path.md)

`
    public
        mixed
    TOKEN_PATH
    = 'api/token'
`

### Methods  [header link](class-aws-credentials-instanceprofileprovider-methods.md)

#### \_\_construct()  [header link](class-aws-credentials-instanceprofileprovider-method-construct.md)

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

#### \_\_invoke()  [header link](class-aws-credentials-instanceprofileprovider-method-invoke.md)

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
  - [Constants](class-aws-credentials-instanceprofileprovider-toc-constants.md)
  - [Methods](class-aws-credentials-instanceprofileprovider-toc-methods.md)
- Constants
  - [CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-service-endpoint.md)
  - [CFG\_EC2\_METADATA\_SERVICE\_ENDPOINT\_MODE](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-service-endpoint-mode.md)
  - [CFG\_EC2\_METADATA\_V1\_DISABLED](class-aws-credentials-instanceprofileprovider-constant-cfg-ec2-metadata-v1-disabled.md)
  - [CRED\_PATH](class-aws-credentials-instanceprofileprovider-constant-cred-path.md)
  - [DEFAULT\_AWS\_EC2\_METADATA\_V1\_DISABLED](class-aws-credentials-instanceprofileprovider-constant-default-aws-ec2-metadata-v1-disabled.md)
  - [DEFAULT\_METADATA\_SERVICE\_IPv4\_ENDPOINT](class-aws-credentials-instanceprofileprovider-constant-default-metadata-service-ipv4-endpoint.md)
  - [DEFAULT\_METADATA\_SERVICE\_IPv6\_ENDPOINT](class-aws-credentials-instanceprofileprovider-constant-default-metadata-service-ipv6-endpoint.md)
  - [DEFAULT\_RETRIES](class-aws-credentials-instanceprofileprovider-constant-default-retries.md)
  - [DEFAULT\_TIMEOUT](class-aws-credentials-instanceprofileprovider-constant-default-timeout.md)
  - [DEFAULT\_TOKEN\_TTL\_SECONDS](class-aws-credentials-instanceprofileprovider-constant-default-token-ttl-seconds.md)
  - [ENDPOINT\_MODE\_IPv4](class-aws-credentials-instanceprofileprovider-constant-endpoint-mode-ipv4.md)
  - [ENDPOINT\_MODE\_IPv6](class-aws-credentials-instanceprofileprovider-constant-endpoint-mode-ipv6.md)
  - [ENV\_DISABLE](class-aws-credentials-instanceprofileprovider-constant-env-disable.md)
  - [ENV\_RETRIES](class-aws-credentials-instanceprofileprovider-constant-env-retries.md)
  - [ENV\_TIMEOUT](class-aws-credentials-instanceprofileprovider-constant-env-timeout.md)
  - [TOKEN\_PATH](class-aws-credentials-instanceprofileprovider-constant-token-path.md)
- Methods
  - [\_\_construct()](class-aws-credentials-instanceprofileprovider-method-construct.md)
  - [\_\_invoke()](class-aws-credentials-instanceprofileprovider-method-invoke.md)

[Back To Top](class-aws-credentials-instanceprofileprovider-top.md)
