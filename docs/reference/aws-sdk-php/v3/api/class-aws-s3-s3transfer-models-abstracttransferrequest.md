Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Models](namespace-aws-s3-s3transfer-models.md)

## AbstractTransferRequest        in package    - [Aws](package-aws.md)

AbstractYes

### Table of Contents  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-toc.md)

#### Properties  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-toc-properties.md)

[$configKeys](class-aws-s3-s3transfer-models-abstracttransferrequest-property-configkeys.md)
: array<string\|int, mixed>

#### Methods  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-construct.md)
: mixed [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getconfig.md)
: array<string\|int, mixed> [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getlisteners.md)
: array<string\|int, mixed> Get current listeners.[getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getprogresstracker.md)
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null Get the progress tracker.[getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getsingleobjectlisteners.md)
: array<string\|int, mixed> Get listeners that should receive single-object events.[updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-updateconfigwithdefaults.md)
: void [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-validateconfig.md)
: void For validating config. By default, it provides an empty
implementation.

### Properties  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-properties.md)

#### $configKeys  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-property-configkeys.md)

`
    public
    static    array<string|int, mixed>
    $configKeys
     = ['track_progress' => 'bool']`

### Methods  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $listeners, AbstractTransferListener|null $progressTracker, array<string|int, mixed> $config[, array<string|int, mixed> $singleObjectListeners = [] ]) : mixed`

##### Parameters

$listeners
: array<string\|int, mixed>$progressTracker
: [AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null$config
: array<string\|int, mixed>$singleObjectListeners
: array<string\|int, mixed>
= \[\]

#### getConfig()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getconfig.md)

`
    public
                    getConfig() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getlisteners.md)

Get current listeners.

`
    public
                    getListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getProgressTracker()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getprogresstracker.md)

Get the progress tracker.

`
    public
                    getProgressTracker() : AbstractTransferListener|null`

##### Return values

[AbstractTransferListener](class-aws-s3-s3transfer-progress-abstracttransferlistener.md) \|null

#### getSingleObjectListeners()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getsingleobjectlisteners.md)

Get listeners that should receive single-object events.

`
    public
                    getSingleObjectListeners() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### updateConfigWithDefaults()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-method-updateconfigwithdefaults.md)

`
    public
                    updateConfigWithDefaults(array<string|int, mixed> $defaultConfig) : void`

##### Parameters

$defaultConfig
: array<string\|int, mixed>

#### validateConfig()  [header link](class-aws-s3-s3transfer-models-abstracttransferrequest-method-validateconfig.md)

For validating config. By default, it provides an empty
implementation.

`
    public
                    validateConfig() : void`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](class-aws-s3-s3transfer-models-abstracttransferrequest-toc-properties.md)
  - [Methods](class-aws-s3-s3transfer-models-abstracttransferrequest-toc-methods.md)
- Properties
  - [$configKeys](class-aws-s3-s3transfer-models-abstracttransferrequest-property-configkeys.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-construct.md)
  - [getConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getconfig.md)
  - [getListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getlisteners.md)
  - [getProgressTracker()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getprogresstracker.md)
  - [getSingleObjectListeners()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-getsingleobjectlisteners.md)
  - [updateConfigWithDefaults()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-updateconfigwithdefaults.md)
  - [validateConfig()](class-aws-s3-s3transfer-models-abstracttransferrequest-method-validateconfig.md)

[Back To Top](class-aws-s3-s3transfer-models-abstracttransferrequest-top.md)
