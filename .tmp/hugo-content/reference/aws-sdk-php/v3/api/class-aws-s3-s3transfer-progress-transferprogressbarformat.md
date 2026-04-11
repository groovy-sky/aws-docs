Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## TransferProgressBarFormat     extends [AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)   in package    - [Aws](package-aws.md)

FinalYes

Defines a progress bar format.

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-toc.md)

#### Constants  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-toc-constants.md)

[FORMAT\_PARAMETERS](class-aws-s3-s3transfer-progress-transferprogressbarformat-constant-format-parameters.md)
= \['object\_name', 'progress\_bar', 'percent', 'transferred', 'to\_be\_transferred', 'unit'\] [FORMAT\_TEMPLATE](class-aws-s3-s3transfer-progress-transferprogressbarformat-constant-format-template.md)
= "\|object\_name\|:\\n\[\|progress\_bar\|\]" . " \|percent\|% \|transferred\|/\|to\_be\_transferred\| \|unit\|"

#### Methods  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method___construct)
: mixed [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_format)
: string [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_getArgs)
: array<string\|int, mixed> [getFormatParameters()](class-aws-s3-s3transfer-progress-transferprogressbarformat-method-getformatparameters.md)
: array<string\|int, mixed> [getFormatTemplate()](class-aws-s3-s3transfer-progress-transferprogressbarformat-method-getformattemplate.md)
: string [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArg)
: void [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArgs)
: void To set multiple arguments at once.

### Constants  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-constants.md)

#### FORMAT\_PARAMETERS  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-constant-format-parameters.md)

`
    public
        mixed
    FORMAT_PARAMETERS
    = ['object_name', 'progress_bar', 'percent', 'transferred', 'to_be_transferred', 'unit']
`

#### FORMAT\_TEMPLATE  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-constant-format-template.md)

`
    public
        mixed
    FORMAT_TEMPLATE
    = "|object_name|:\n[|progress_bar|]" . " |percent|% |transferred|/|to_be_transferred| |unit|"
`

### Methods  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-methods.md)

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

#### getFormatParameters()  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-method-getformatparameters.md)

`
    public
                    getFormatParameters() : array<string|int, mixed>`

##### Tags  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-method-getformatparameters-tags.md)

inheritDoc

##### Return values

array<string\|int, mixed>

#### getFormatTemplate()  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-method-getformattemplate.md)

`
    public
                    getFormatTemplate() : string`

##### Tags  [header link](class-aws-s3-s3transfer-progress-transferprogressbarformat-method-getformattemplate-tags.md)

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
  - [Constants](class-aws-s3-s3transfer-progress-transferprogressbarformat-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-transferprogressbarformat-toc-methods.md)
- Constants
  - [FORMAT\_PARAMETERS](class-aws-s3-s3transfer-progress-transferprogressbarformat-constant-format-parameters.md)
  - [FORMAT\_TEMPLATE](class-aws-s3-s3transfer-progress-transferprogressbarformat-constant-format-template.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method___construct)
  - [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_format)
  - [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_getArgs)
  - [getFormatParameters()](class-aws-s3-s3transfer-progress-transferprogressbarformat-method-getformatparameters.md)
  - [getFormatTemplate()](class-aws-s3-s3transfer-progress-transferprogressbarformat-method-getformattemplate.md)
  - [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArg)
  - [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArgs)

[Back To Top](class-aws-s3-s3transfer-progress-transferprogressbarformat-top.md)
