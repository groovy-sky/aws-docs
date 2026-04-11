Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## ConsoleProgressBar        in package    - [Aws](package-aws.md)       implements  [ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)

FinalYes

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-toc.md)

#### Interfaces  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-toc-interfaces.md)

[ProgressBarInterface](class-aws-s3-s3transfer-progress-progressbarinterface.md)Progress bar base implementation.

#### Constants  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-toc-constants.md)

[DEFAULT\_PROGRESS\_BAR\_CHAR](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-default-progress-bar-char.md)
= '#' [DEFAULT\_PROGRESS\_BAR\_WIDTH](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-default-progress-bar-width.md)
= 50 [MAX\_PROGRESS\_BAR\_WIDTH](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-max-progress-bar-width.md)
= 50

#### Methods  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-construct.md)
: mixed [getPercentCompleted()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getpercentcompleted.md)
: int [getProgressBarChar()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarchar.md)
: string [getProgressBarFormat()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarformat.md)
: [AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)[getProgressBarWidth()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarwidth.md)
: int [render()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-render.md)
: string [setPercentCompleted()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-setpercentcompleted.md)
: void Set current progress percent.

### Constants  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-constants.md)

#### DEFAULT\_PROGRESS\_BAR\_CHAR  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-default-progress-bar-char.md)

`
    public
        mixed
    DEFAULT_PROGRESS_BAR_CHAR
    = '#'
`

#### DEFAULT\_PROGRESS\_BAR\_WIDTH  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-default-progress-bar-width.md)

`
    public
        mixed
    DEFAULT_PROGRESS_BAR_WIDTH
    = 50
`

#### MAX\_PROGRESS\_BAR\_WIDTH  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-max-progress-bar-width.md)

`
    public
        mixed
    MAX_PROGRESS_BAR_WIDTH
    = 50
`

### Methods  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-method-construct.md)

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

#### getPercentCompleted()  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getpercentcompleted.md)

`
    public
                    getPercentCompleted() : int`

##### Return values

int

#### getProgressBarChar()  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarchar.md)

`
    public
                    getProgressBarChar() : string`

##### Return values

string

#### getProgressBarFormat()  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarformat.md)

`
    public
                    getProgressBarFormat() : AbstractProgressBarFormat`

##### Return values

[AbstractProgressBarFormat](class-aws-s3-s3transfer-progress-abstractprogressbarformat.md)

#### getProgressBarWidth()  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarwidth.md)

`
    public
                    getProgressBarWidth() : int`

##### Return values

int

#### render()  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-method-render.md)

`
    public
                    render() : string`

##### Tags  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-method-render-tags.md)

inheritDoc

##### Return values

string

#### setPercentCompleted()  [header link](class-aws-s3-s3transfer-progress-consoleprogressbar-method-setpercentcompleted.md)

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
  - [Constants](class-aws-s3-s3transfer-progress-consoleprogressbar-toc-constants.md)
  - [Methods](class-aws-s3-s3transfer-progress-consoleprogressbar-toc-methods.md)
- Constants
  - [DEFAULT\_PROGRESS\_BAR\_CHAR](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-default-progress-bar-char.md)
  - [DEFAULT\_PROGRESS\_BAR\_WIDTH](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-default-progress-bar-width.md)
  - [MAX\_PROGRESS\_BAR\_WIDTH](class-aws-s3-s3transfer-progress-consoleprogressbar-constant-max-progress-bar-width.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-construct.md)
  - [getPercentCompleted()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getpercentcompleted.md)
  - [getProgressBarChar()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarchar.md)
  - [getProgressBarFormat()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarformat.md)
  - [getProgressBarWidth()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-getprogressbarwidth.md)
  - [render()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-render.md)
  - [setPercentCompleted()](class-aws-s3-s3transfer-progress-consoleprogressbar-method-setpercentcompleted.md)

[Back To Top](class-aws-s3-s3transfer-progress-consoleprogressbar-top.md)
