Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Models](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.models.html)

## AbstractTransferRequest        in package    - [Aws](package-aws.md)

AbstractYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#toc)

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#toc-properties)

[$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#property_configKeys)
: array<string\|int, mixed>

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method___construct)
: mixed [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_getConfig)
: array<string\|int, mixed> [getListeners()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_getListeners)
: array<string\|int, mixed> Get current listeners.[getProgressTracker()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_getProgressTracker)
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null Get the progress tracker.[getSingleObjectListeners()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_getSingleObjectListeners)
: array<string\|int, mixed> Get listeners that should receive single-object events.[updateConfigWithDefaults()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_updateConfigWithDefaults)
: void [validateConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_validateConfig)
: void For validating config. By default, it provides an empty
implementation.

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#properties)

#### $configKeys  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#property_configKeys)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['track_progress' => 'bool']`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $listeners, AbstractTransferListener|null $progressTracker, array<string|int, mixed> $config[, array<string|int, mixed> $singleObjectListeners = [] ]) : mixed`

##### Parameters

$listeners
: array<string\|int, mixed>$progressTracker
: [AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null$config
: array<string\|int, mixed>$singleObjectListeners
: array<string\|int, mixed>
= \[\]

#### getConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#method_getConfig)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getListeners()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#method_getListeners)

Get current listeners.

`
    public
                    getListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getProgressTracker()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#method_getProgressTracker)

Get the progress tracker.

`
    public
                    getProgressTracker() : AbstractTransferListener|null`

##### Return values

[AbstractTransferListener](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractTransferListener.html) \|null

#### getSingleObjectListeners()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#method_getSingleObjectListeners)

Get listeners that should receive single-object events.

`
    public
                    getSingleObjectListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### updateConfigWithDefaults()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#method_updateConfigWithDefaults)

`
    public
                    updateConfigWithDefaults(array<string|int, mixed> $defaultConfig) : void`

##### Parameters

$defaultConfig
: array<string\|int, mixed>

#### validateConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html\#method_validateConfig)

For validating config. By default, it provides an empty
implementation.

`
    public
                    validateConfig() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#toc-methods)
- Properties
  - [$configKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#property_configKeys)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method___construct)
  - [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_getConfig)
  - [getListeners()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_getListeners)
  - [getProgressTracker()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_getProgressTracker)
  - [getSingleObjectListeners()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_getSingleObjectListeners)
  - [updateConfigWithDefaults()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_updateConfigWithDefaults)
  - [validateConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#method_validateConfig)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Models.AbstractTransferRequest.html#top)
