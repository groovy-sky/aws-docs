Menu

- [Aws](namespace-aws.md)
- [Credentials](namespace-aws-credentials.md)

## EcsCredentialProvider        in package    - [Aws](package-aws.md)

Credential provider that fetches container credentials with GET request.

container environment variables are used in constructing request URI.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#toc-constants)

[DEFAULT\_ENV\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_DEFAULT_ENV_RETRIES)
= 3 [DEFAULT\_ENV\_TIMEOUT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_DEFAULT_ENV_TIMEOUT)
= 1.0 [EKS\_SERVER\_HOST\_IPV4](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_EKS_SERVER_HOST_IPV4)
= '169.254.170.23' [EKS\_SERVER\_HOST\_IPV6](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_EKS_SERVER_HOST_IPV6)
= 'fd00:ec2::23' [ENV\_AUTH\_TOKEN](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_AUTH_TOKEN)
= "AWS\_CONTAINER\_AUTHORIZATION\_TOKEN" [ENV\_AUTH\_TOKEN\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_AUTH_TOKEN_FILE)
= "AWS\_CONTAINER\_AUTHORIZATION\_TOKEN\_FILE" [ENV\_FULL\_URI](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_FULL_URI)
= "AWS\_CONTAINER\_CREDENTIALS\_FULL\_URI" [ENV\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_RETRIES)
= 'AWS\_METADATA\_SERVICE\_NUM\_ATTEMPTS' [ENV\_TIMEOUT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_TIMEOUT)
= 'AWS\_METADATA\_SERVICE\_TIMEOUT' [ENV\_URI](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_URI)
= "AWS\_CONTAINER\_CREDENTIALS\_RELATIVE\_URI" [SERVER\_URI](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_SERVER_URI)
= 'http://169.254.170.2'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#method___construct)
: mixed The constructor accepts following options:
\- timeout: (optional) Connection timeout, in seconds, default 1.0
\- retries: Optional number of retries to be attempted, default 3.[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#method___invoke)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Load container credentials.[getAttempts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#method_getAttempts)
: int Returns the number of attempts that have been done.[setHeaderForAuthToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#method_setHeaderForAuthToken)
: mixed

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constants)

#### DEFAULT\_ENV\_RETRIES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_DEFAULT_ENV_RETRIES)

`
    public
        mixed
    DEFAULT_ENV_RETRIES
    = 3
`

#### DEFAULT\_ENV\_TIMEOUT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_DEFAULT_ENV_TIMEOUT)

`
    public
        mixed
    DEFAULT_ENV_TIMEOUT
    = 1.0
`

#### EKS\_SERVER\_HOST\_IPV4  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_EKS_SERVER_HOST_IPV4)

`
    public
        mixed
    EKS_SERVER_HOST_IPV4
    = '169.254.170.23'
`

#### EKS\_SERVER\_HOST\_IPV6  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_EKS_SERVER_HOST_IPV6)

`
    public
        mixed
    EKS_SERVER_HOST_IPV6
    = 'fd00:ec2::23'
`

#### ENV\_AUTH\_TOKEN  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_ENV_AUTH_TOKEN)

`
    public
        mixed
    ENV_AUTH_TOKEN
    = "AWS_CONTAINER_AUTHORIZATION_TOKEN"
`

#### ENV\_AUTH\_TOKEN\_FILE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_ENV_AUTH_TOKEN_FILE)

`
    public
        mixed
    ENV_AUTH_TOKEN_FILE
    = "AWS_CONTAINER_AUTHORIZATION_TOKEN_FILE"
`

#### ENV\_FULL\_URI  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_ENV_FULL_URI)

`
    public
        mixed
    ENV_FULL_URI
    = "AWS_CONTAINER_CREDENTIALS_FULL_URI"
`

#### ENV\_RETRIES  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_ENV_RETRIES)

`
    public
        mixed
    ENV_RETRIES
    = 'AWS_METADATA_SERVICE_NUM_ATTEMPTS'
`

#### ENV\_TIMEOUT  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_ENV_TIMEOUT)

`
    public
        mixed
    ENV_TIMEOUT
    = 'AWS_METADATA_SERVICE_TIMEOUT'
`

#### ENV\_URI  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_ENV_URI)

`
    public
        mixed
    ENV_URI
    = "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI"
`

#### SERVER\_URI  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#constant_SERVER_URI)

`
    public
        mixed
    SERVER_URI
    = 'http://169.254.170.2'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#method___construct)

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

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#method___invoke)

Load container credentials.

`
    public
                    __invoke() : PromiseInterface`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#method___invoke\#tags)

throwsGuzzleException

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getAttempts()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#method_getAttempts)

Returns the number of attempts that have been done.

`
    public
                    getAttempts() : int`

##### Return values

int

#### setHeaderForAuthToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#method_setHeaderForAuthToken)

`
    public
                    setHeaderForAuthToken() : mixed`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html\#method_setHeaderForAuthToken\#tags)

deprecated
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#toc-methods)
- Constants
  - [DEFAULT\_ENV\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_DEFAULT_ENV_RETRIES)
  - [DEFAULT\_ENV\_TIMEOUT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_DEFAULT_ENV_TIMEOUT)
  - [EKS\_SERVER\_HOST\_IPV4](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_EKS_SERVER_HOST_IPV4)
  - [EKS\_SERVER\_HOST\_IPV6](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_EKS_SERVER_HOST_IPV6)
  - [ENV\_AUTH\_TOKEN](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_AUTH_TOKEN)
  - [ENV\_AUTH\_TOKEN\_FILE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_AUTH_TOKEN_FILE)
  - [ENV\_FULL\_URI](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_FULL_URI)
  - [ENV\_RETRIES](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_RETRIES)
  - [ENV\_TIMEOUT](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_TIMEOUT)
  - [ENV\_URI](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_ENV_URI)
  - [SERVER\_URI](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#constant_SERVER_URI)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#method___invoke)
  - [getAttempts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#method_getAttempts)
  - [setHeaderForAuthToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#method_setHeaderForAuthToken)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html#top)
