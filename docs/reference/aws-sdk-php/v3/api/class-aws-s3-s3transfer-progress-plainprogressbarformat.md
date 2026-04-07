Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## PlainProgressBarFormat     extends [AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)   in package    - [Aws](package-aws.md)

FinalYes

Defines a progress bar format.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#toc-constants)

[FORMAT\_PARAMETERS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#constant_FORMAT_PARAMETERS)
= \['object\_name', 'progress\_bar', 'percent'\] [FORMAT\_TEMPLATE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#constant_FORMAT_TEMPLATE)
= "\|object\_name\|:\\n\[\|progress\_bar\|\] \|percent\|%"

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#toc-methods)

[\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method___construct)
: mixed [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_format)
: string [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_getArgs)
: array<string\|int, mixed> [getFormatParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#method_getFormatParameters)
: array<string\|int, mixed> [getFormatTemplate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#method_getFormatTemplate)
: string [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArg)
: void [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArgs)
: void To set multiple arguments at once.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#constants)

#### FORMAT\_PARAMETERS  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#constant_FORMAT_PARAMETERS)

`
    public
        mixed
    FORMAT_PARAMETERS
    = ['object_name', 'progress_bar', 'percent']
`

#### FORMAT\_TEMPLATE  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#constant_FORMAT_TEMPLATE)

`
    public
        mixed
    FORMAT_TEMPLATE
    = "|object_name|:\n[|progress_bar|] |percent|%"
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#methods)

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

#### getFormatParameters()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#method_getFormatParameters)

`
    public
                    getFormatParameters() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormatTemplate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html\#method_getFormatTemplate)

`
    public
                    getFormatTemplate() : string`

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#toc-methods)
- Constants
  - [FORMAT\_PARAMETERS](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#constant_FORMAT_PARAMETERS)
  - [FORMAT\_TEMPLATE](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#constant_FORMAT_TEMPLATE)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method___construct)
  - [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_format)
  - [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_getArgs)
  - [getFormatParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#method_getFormatParameters)
  - [getFormatTemplate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#method_getFormatTemplate)
  - [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArg)
  - [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md#method_setArgs)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.PlainProgressBarFormat.html#top)
