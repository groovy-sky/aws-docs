Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## ColoredTransferProgressBarFormat     extends [AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)   in package    - [Aws](package-aws.md)

FinalYes

Defines a progress bar format.

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-toc-constants.md)

[BLACK\_COLOR\_CODE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-black-color-code.md)
= '\[30m' [BLUE\_COLOR\_CODE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-blue-color-code.md)
= '\[34m' [FORMAT\_PARAMETERS](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-format-parameters.md)
= \['progress\_bar', 'percent', 'transferred', 'to\_be\_transferred', 'unit', 'color\_code', 'message', 'object\_name'\] [FORMAT\_TEMPLATE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-format-template.md)
= "\|object\_name\|:\\n" . "\\x1b\|color\_code\|\[\|progress\_bar\|\] \|percent\|% " . "\|transferred\|/\|to\_be\_transferred\| \|unit\| \|message\|\\x1b\[0m" [GREEN\_COLOR\_CODE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-green-color-code.md)
= '\[32m' [RED\_COLOR\_CODE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-red-color-code.md)
= '\[31m'

#### Methods  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method___construct)
: mixed [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_format)
: string [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_getArgs)
: array<string\|int, mixed> [getFormatParameters()](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-method-getformatparameters.md)
: array<string\|int, mixed> [getFormatTemplate()](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-method-getformattemplate.md)
: string [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArg)
: void [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArgs)
: void To set multiple arguments at once.

### Constants  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constants.md)

#### BLACK\_COLOR\_CODE  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-black-color-code.md)

`
    public
        mixed
    BLACK_COLOR_CODE
    = '[30m'
`

#### BLUE\_COLOR\_CODE  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-blue-color-code.md)

`
    public
        mixed
    BLUE_COLOR_CODE
    = '[34m'
`

#### FORMAT\_PARAMETERS  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-format-parameters.md)

`
    public
        mixed
    FORMAT_PARAMETERS
    = ['progress_bar', 'percent', 'transferred', 'to_be_transferred', 'unit', 'color_code', 'message', 'object_name']
`

#### FORMAT\_TEMPLATE  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-format-template.md)

`
    public
        mixed
    FORMAT_TEMPLATE
    = "|object_name|:\n" . "\x1b|color_code|[|progress_bar|] |percent|% " . "|transferred|/|to_be_transferred| |unit| |message|\x1b[0m"
`

#### GREEN\_COLOR\_CODE  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-green-color-code.md)

`
    public
        mixed
    GREEN_COLOR_CODE
    = '[32m'
`

#### RED\_COLOR\_CODE  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-red-color-code.md)

`
    public
        mixed
    RED_COLOR_CODE
    = '[31m'
`

### Methods  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-methods.md)

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

#### getFormatParameters()  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-method-getformatparameters.md)

`
    public
                    getFormatParameters() : array<string|int, mixed>`

##### Tags  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-method-getformatparameters-tags.md)

inheritDoc

##### Return values

array<string\|int, mixed>

#### getFormatTemplate()  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-method-getformattemplate.md)

`
    public
                    getFormatTemplate() : string`

##### Tags  [header link](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-method-getformattemplate-tags.md)

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
  - [Constants](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-toc-methods.md)
- Constants
  - [BLACK\_COLOR\_CODE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-black-color-code.md)
  - [BLUE\_COLOR\_CODE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-blue-color-code.md)
  - [FORMAT\_PARAMETERS](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-format-parameters.md)
  - [FORMAT\_TEMPLATE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-format-template.md)
  - [GREEN\_COLOR\_CODE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-green-color-code.md)
  - [RED\_COLOR\_CODE](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-constant-red-color-code.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method___construct)
  - [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_format)
  - [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_getArgs)
  - [getFormatParameters()](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-method-getformatparameters.md)
  - [getFormatTemplate()](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-method-getformattemplate.md)
  - [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArg)
  - [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArgs)

[Back To Top](class-aws-s3-s3transfer-progress-coloredtransferprogressbarformat-top.md)
