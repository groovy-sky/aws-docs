Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## EcsCredentialProvider        in package    - [Aws](package-aws.md)

Credential provider that fetches container credentials with GET request.

container environment variables are used in constructing request URI.

### Table of Contents  [header link](class-aws-credentials-ecscredentialprovider-toc.md)

#### Constants  [header link](class-aws-credentials-ecscredentialprovider-toc-constants.md)

[DEFAULT\_ENV\_RETRIES](class-aws-credentials-ecscredentialprovider-constant-default-env-retries.md)
= 3 [DEFAULT\_ENV\_TIMEOUT](class-aws-credentials-ecscredentialprovider-constant-default-env-timeout.md)
= 1.0 [EKS\_SERVER\_HOST\_IPV4](class-aws-credentials-ecscredentialprovider-constant-eks-server-host-ipv4.md)
= '169.254.170.23' [EKS\_SERVER\_HOST\_IPV6](class-aws-credentials-ecscredentialprovider-constant-eks-server-host-ipv6.md)
= 'fd00:ec2::23' [ENV\_AUTH\_TOKEN](class-aws-credentials-ecscredentialprovider-constant-env-auth-token.md)
= "AWS\_CONTAINER\_AUTHORIZATION\_TOKEN" [ENV\_AUTH\_TOKEN\_FILE](class-aws-credentials-ecscredentialprovider-constant-env-auth-token-file.md)
= "AWS\_CONTAINER\_AUTHORIZATION\_TOKEN\_FILE" [ENV\_FULL\_URI](class-aws-credentials-ecscredentialprovider-constant-env-full-uri.md)
= "AWS\_CONTAINER\_CREDENTIALS\_FULL\_URI" [ENV\_RETRIES](class-aws-credentials-ecscredentialprovider-constant-env-retries.md)
= 'AWS\_METADATA\_SERVICE\_NUM\_ATTEMPTS' [ENV\_TIMEOUT](class-aws-credentials-ecscredentialprovider-constant-env-timeout.md)
= 'AWS\_METADATA\_SERVICE\_TIMEOUT' [ENV\_URI](class-aws-credentials-ecscredentialprovider-constant-env-uri.md)
= "AWS\_CONTAINER\_CREDENTIALS\_RELATIVE\_URI" [SERVER\_URI](class-aws-credentials-ecscredentialprovider-constant-server-uri.md)
= 'http://169.254.170.2'

#### Methods  [header link](class-aws-credentials-ecscredentialprovider-toc-methods.md)

[\_\_construct()](class-aws-credentials-ecscredentialprovider-method-construct.md)
: mixed The constructor accepts following options:
\- timeout: (optional) Connection timeout, in seconds, default 1.0
\- retries: Optional number of retries to be attempted, default 3.[\_\_invoke()](class-aws-credentials-ecscredentialprovider-method-invoke.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Load container credentials.[getAttempts()](class-aws-credentials-ecscredentialprovider-method-getattempts.md)
: int Returns the number of attempts that have been done.[setHeaderForAuthToken()](class-aws-credentials-ecscredentialprovider-method-setheaderforauthtoken.md)
: mixed

### Constants  [header link](class-aws-credentials-ecscredentialprovider-constants.md)

#### DEFAULT\_ENV\_RETRIES  [header link](class-aws-credentials-ecscredentialprovider-constant-default-env-retries.md)

`
    public
        mixed
    DEFAULT_ENV_RETRIES
    = 3
`

#### DEFAULT\_ENV\_TIMEOUT  [header link](class-aws-credentials-ecscredentialprovider-constant-default-env-timeout.md)

`
    public
        mixed
    DEFAULT_ENV_TIMEOUT
    = 1.0
`

#### EKS\_SERVER\_HOST\_IPV4  [header link](class-aws-credentials-ecscredentialprovider-constant-eks-server-host-ipv4.md)

`
    public
        mixed
    EKS_SERVER_HOST_IPV4
    = '169.254.170.23'
`

#### EKS\_SERVER\_HOST\_IPV6  [header link](class-aws-credentials-ecscredentialprovider-constant-eks-server-host-ipv6.md)

`
    public
        mixed
    EKS_SERVER_HOST_IPV6
    = 'fd00:ec2::23'
`

#### ENV\_AUTH\_TOKEN  [header link](class-aws-credentials-ecscredentialprovider-constant-env-auth-token.md)

`
    public
        mixed
    ENV_AUTH_TOKEN
    = "AWS_CONTAINER_AUTHORIZATION_TOKEN"
`

#### ENV\_AUTH\_TOKEN\_FILE  [header link](class-aws-credentials-ecscredentialprovider-constant-env-auth-token-file.md)

`
    public
        mixed
    ENV_AUTH_TOKEN_FILE
    = "AWS_CONTAINER_AUTHORIZATION_TOKEN_FILE"
`

#### ENV\_FULL\_URI  [header link](class-aws-credentials-ecscredentialprovider-constant-env-full-uri.md)

`
    public
        mixed
    ENV_FULL_URI
    = "AWS_CONTAINER_CREDENTIALS_FULL_URI"
`

#### ENV\_RETRIES  [header link](class-aws-credentials-ecscredentialprovider-constant-env-retries.md)

`
    public
        mixed
    ENV_RETRIES
    = 'AWS_METADATA_SERVICE_NUM_ATTEMPTS'
`

#### ENV\_TIMEOUT  [header link](class-aws-credentials-ecscredentialprovider-constant-env-timeout.md)

`
    public
        mixed
    ENV_TIMEOUT
    = 'AWS_METADATA_SERVICE_TIMEOUT'
`

#### ENV\_URI  [header link](class-aws-credentials-ecscredentialprovider-constant-env-uri.md)

`
    public
        mixed
    ENV_URI
    = "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI"
`

#### SERVER\_URI  [header link](class-aws-credentials-ecscredentialprovider-constant-server-uri.md)

`
    public
        mixed
    SERVER_URI
    = 'http://169.254.170.2'
`

### Methods  [header link](class-aws-credentials-ecscredentialprovider-methods.md)

#### \_\_construct()  [header link](class-aws-credentials-ecscredentialprovider-method-construct.md)

The constructor accepts following options:
\- timeout: (optional) Connection timeout, in seconds, default 1.0
\- retries: Optional number of retries to be attempted, default 3.

`
    public
                    __construct([array<string|int, mixed> $config = [] ]) : mixed`

- client: An EcsClient to make request from

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

Configuration options

#### \_\_invoke()  [header link](class-aws-credentials-ecscredentialprovider-method-invoke.md)

Load container credentials.

`
    public
                    __invoke() : PromiseInterface`

##### Tags  [header link](class-aws-credentials-ecscredentialprovider-method-invoke-tags.md)

throwsGuzzleException

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getAttempts()  [header link](class-aws-credentials-ecscredentialprovider-method-getattempts.md)

Returns the number of attempts that have been done.

`
    public
                    getAttempts() : int`

##### Return values

int

#### setHeaderForAuthToken()  [header link](class-aws-credentials-ecscredentialprovider-method-setheaderforauthtoken.md)

`
    public
                    setHeaderForAuthToken() : mixed`

##### Tags  [header link](class-aws-credentials-ecscredentialprovider-method-setheaderforauthtoken-tags.md)

deprecated
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-credentials-ecscredentialprovider-toc-constants.md)
  - [Methods](class-aws-credentials-ecscredentialprovider-toc-methods.md)
- Constants
  - [DEFAULT\_ENV\_RETRIES](class-aws-credentials-ecscredentialprovider-constant-default-env-retries.md)
  - [DEFAULT\_ENV\_TIMEOUT](class-aws-credentials-ecscredentialprovider-constant-default-env-timeout.md)
  - [EKS\_SERVER\_HOST\_IPV4](class-aws-credentials-ecscredentialprovider-constant-eks-server-host-ipv4.md)
  - [EKS\_SERVER\_HOST\_IPV6](class-aws-credentials-ecscredentialprovider-constant-eks-server-host-ipv6.md)
  - [ENV\_AUTH\_TOKEN](class-aws-credentials-ecscredentialprovider-constant-env-auth-token.md)
  - [ENV\_AUTH\_TOKEN\_FILE](class-aws-credentials-ecscredentialprovider-constant-env-auth-token-file.md)
  - [ENV\_FULL\_URI](class-aws-credentials-ecscredentialprovider-constant-env-full-uri.md)
  - [ENV\_RETRIES](class-aws-credentials-ecscredentialprovider-constant-env-retries.md)
  - [ENV\_TIMEOUT](class-aws-credentials-ecscredentialprovider-constant-env-timeout.md)
  - [ENV\_URI](class-aws-credentials-ecscredentialprovider-constant-env-uri.md)
  - [SERVER\_URI](class-aws-credentials-ecscredentialprovider-constant-server-uri.md)
- Methods
  - [\_\_construct()](class-aws-credentials-ecscredentialprovider-method-construct.md)
  - [\_\_invoke()](class-aws-credentials-ecscredentialprovider-method-invoke.md)
  - [getAttempts()](class-aws-credentials-ecscredentialprovider-method-getattempts.md)
  - [setHeaderForAuthToken()](class-aws-credentials-ecscredentialprovider-method-setheaderforauthtoken.md)

[Back To Top](class-aws-credentials-ecscredentialprovider-top.md)
