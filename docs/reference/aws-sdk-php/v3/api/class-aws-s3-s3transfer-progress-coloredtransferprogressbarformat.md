Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## ColoredTransferProgressBarFormat     extends [AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)   in package    - [Aws](package-aws.md)

FinalYes

Defines a progress bar format.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#toc-constants)

[BLACK\_COLOR\_CODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_BLACK_COLOR_CODE)
= '\[30m' [BLUE\_COLOR\_CODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_BLUE_COLOR_CODE)
= '\[34m' [FORMAT\_PARAMETERS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_FORMAT_PARAMETERS)
= \['progress\_bar', 'percent', 'transferred', 'to\_be\_transferred', 'unit', 'color\_code', 'message', 'object\_name'\] [FORMAT\_TEMPLATE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_FORMAT_TEMPLATE)
= "\|object\_name\|:\\n" . "\\x1b\|color\_code\|\[\|progress\_bar\|\] \|percent\|% " . "\|transferred\|/\|to\_be\_transferred\| \|unit\| \|message\|\\x1b\[0m" [GREEN\_COLOR\_CODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_GREEN_COLOR_CODE)
= '\[32m' [RED\_COLOR\_CODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_RED_COLOR_CODE)
= '\[31m'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#toc-methods)

[\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method___construct)
: mixed [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_format)
: string [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_getArgs)
: array<string\|int, mixed> [getFormatParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#method_getFormatParameters)
: array<string\|int, mixed> [getFormatTemplate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#method_getFormatTemplate)
: string [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArg)
: void [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArgs)
: void To set multiple arguments at once.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#constants)

#### BLACK\_COLOR\_CODE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#constant_BLACK_COLOR_CODE)

`
    public
        mixed
    BLACK_COLOR_CODE
    = '[30m'
`

#### BLUE\_COLOR\_CODE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#constant_BLUE_COLOR_CODE)

`
    public
        mixed
    BLUE_COLOR_CODE
    = '[34m'
`

#### FORMAT\_PARAMETERS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#constant_FORMAT_PARAMETERS)

`
    public
        mixed
    FORMAT_PARAMETERS
    = ['progress_bar', 'percent', 'transferred', 'to_be_transferred', 'unit', 'color_code', 'message', 'object_name']
`

#### FORMAT\_TEMPLATE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#constant_FORMAT_TEMPLATE)

`
    public
        mixed
    FORMAT_TEMPLATE
    = "|object_name|:\n" . "\x1b|color_code|[|progress_bar|] |percent|% " . "|transferred|/|to_be_transferred| |unit| |message|\x1b[0m"
`

#### GREEN\_COLOR\_CODE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#constant_GREEN_COLOR_CODE)

`
    public
        mixed
    GREEN_COLOR_CODE
    = '[32m'
`

#### RED\_COLOR\_CODE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#constant_RED_COLOR_CODE)

`
    public
        mixed
    RED_COLOR_CODE
    = '[31m'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#methods)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md\#method___construct)

`
    public
                    __construct([array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

#### format()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md\#method_format)

`
    public
                    format() : string`

##### Return values

string

#### getArgs()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md\#method_getArgs)

`
    public
                    getArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormatParameters()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#method_getFormatParameters)

`
    public
                    getFormatParameters() : array<string|int, mixed>`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#method_getFormatParameters\#tags)

inheritDoc

##### Return values

array<string\|int, mixed>

#### getFormatTemplate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#method_getFormatTemplate)

`
    public
                    getFormatTemplate() : string`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html\#method_getFormatTemplate\#tags)

inheritDoc

##### Return values

string

#### setArg()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md\#method_setArg)

`
    public
                    setArg(string $key, mixed $value) : void`

##### Parameters

$key
: string$value
: mixed

#### setArgs()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md\#method_setArgs)

To set multiple arguments at once.

`
    public
                    setArgs(array<string|int, mixed> $args) : void`

It does not override all the values, instead
it adds the arguments individually and if a value
already exists then that value will be overridden.

##### Parameters

$args
: array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#toc-methods)
- Constants
  - [BLACK\_COLOR\_CODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_BLACK_COLOR_CODE)
  - [BLUE\_COLOR\_CODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_BLUE_COLOR_CODE)
  - [FORMAT\_PARAMETERS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_FORMAT_PARAMETERS)
  - [FORMAT\_TEMPLATE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_FORMAT_TEMPLATE)
  - [GREEN\_COLOR\_CODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_GREEN_COLOR_CODE)
  - [RED\_COLOR\_CODE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#constant_RED_COLOR_CODE)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method___construct)
  - [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_format)
  - [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_getArgs)
  - [getFormatParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#method_getFormatParameters)
  - [getFormatTemplate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#method_getFormatTemplate)
  - [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArg)
  - [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArgs)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ColoredTransferProgressBarFormat.html#top)
