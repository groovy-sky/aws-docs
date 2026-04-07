Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## ConsoleProgressBar        in package    - [Aws](package-aws.md)       implements  [ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#toc-interfaces)

[ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)Progress bar base implementation.

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#toc-constants)

[DEFAULT\_PROGRESS\_BAR\_CHAR](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#constant_DEFAULT_PROGRESS_BAR_CHAR)
= '#' [DEFAULT\_PROGRESS\_BAR\_WIDTH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#constant_DEFAULT_PROGRESS_BAR_WIDTH)
= 50 [MAX\_PROGRESS\_BAR\_WIDTH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#constant_MAX_PROGRESS_BAR_WIDTH)
= 50

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method___construct)
: mixed [getPercentCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_getPercentCompleted)
: int [getProgressBarChar()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_getProgressBarChar)
: string [getProgressBarFormat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_getProgressBarFormat)
: [AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)[getProgressBarWidth()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_getProgressBarWidth)
: int [render()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_render)
: string [setPercentCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_setPercentCompleted)
: void Set current progress percent.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#constants)

#### DEFAULT\_PROGRESS\_BAR\_CHAR  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#constant_DEFAULT_PROGRESS_BAR_CHAR)

`
    public
        mixed
    DEFAULT_PROGRESS_BAR_CHAR
    = '#'
`

#### DEFAULT\_PROGRESS\_BAR\_WIDTH  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#constant_DEFAULT_PROGRESS_BAR_WIDTH)

`
    public
        mixed
    DEFAULT_PROGRESS_BAR_WIDTH
    = 50
`

#### MAX\_PROGRESS\_BAR\_WIDTH  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#constant_MAX_PROGRESS_BAR_WIDTH)

`
    public
        mixed
    MAX_PROGRESS_BAR_WIDTH
    = 50
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#method___construct)

`
    public
                    __construct([string $progressBarChar = self::DEFAULT_PROGRESS_BAR_CHAR ][, int $progressBarWidth = self::DEFAULT_PROGRESS_BAR_WIDTH ][, int $percentCompleted = 0 ][, AbstractProgressBarFormat $progressBarFormat = new ColoredTransferProgressBarFormat() ]) : mixed`

##### Parameters

$progressBarChar
: string
= self::DEFAULT\_PROGRESS\_BAR\_CHAR$progressBarWidth
: int
= self::DEFAULT\_PROGRESS\_BAR\_WIDTH$percentCompleted
: int
= 0$progressBarFormat
: [AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)
= new ColoredTransferProgressBarFormat()

#### getPercentCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#method_getPercentCompleted)

`
    public
                    getPercentCompleted() : int`

##### Return values

int

#### getProgressBarChar()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#method_getProgressBarChar)

`
    public
                    getProgressBarChar() : string`

##### Return values

string

#### getProgressBarFormat()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#method_getProgressBarFormat)

`
    public
                    getProgressBarFormat() : AbstractProgressBarFormat`

##### Return values

[AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)

#### getProgressBarWidth()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#method_getProgressBarWidth)

`
    public
                    getProgressBarWidth() : int`

##### Return values

int

#### render()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#method_render)

`
    public
                    render() : string`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#method_render\#tags)

inheritDoc

##### Return values

string

#### setPercentCompleted()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html\#method_setPercentCompleted)

Set current progress percent.

`
    public
                    setPercentCompleted(int $percent) : void`

##### Parameters

$percent
: int
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#toc-methods)
- Constants
  - [DEFAULT\_PROGRESS\_BAR\_CHAR](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#constant_DEFAULT_PROGRESS_BAR_CHAR)
  - [DEFAULT\_PROGRESS\_BAR\_WIDTH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#constant_DEFAULT_PROGRESS_BAR_WIDTH)
  - [MAX\_PROGRESS\_BAR\_WIDTH](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#constant_MAX_PROGRESS_BAR_WIDTH)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method___construct)
  - [getPercentCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_getPercentCompleted)
  - [getProgressBarChar()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_getProgressBarChar)
  - [getProgressBarFormat()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_getProgressBarFormat)
  - [getProgressBarWidth()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_getProgressBarWidth)
  - [render()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_render)
  - [setPercentCompleted()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#method_setPercentCompleted)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.ConsoleProgressBar.html#top)
