Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## MimeType        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html\#toc-methods)

[fromExtension()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html#method_fromExtension)
: string\|null Maps a file extensions to a mimetype.[fromFilename()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html#method_fromFilename)
: string\|null Determines the mimetype of a file by looking at its extension.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html\#methods)

#### fromExtension()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html\#method_fromExtension)

Maps a file extensions to a mimetype.

`
    public
            static        fromExtension(string $extension) : string|null`

##### Parameters

$extension
: string

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html\#method_fromExtension\#tags)

see[https://raw.githubusercontent.com/jshttp/mime-db/master/db.json](https://raw.githubusercontent.com/jshttp/mime-db/master/db.json)

##### Return values

string\|null

#### fromFilename()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html\#method_fromFilename)

Determines the mimetype of a file by looking at its extension.

`
    public
            static        fromFilename(string $filename) : string|null`

##### Parameters

$filename
: string

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html\#method_fromFilename\#tags)

see[https://raw.githubusercontent.com/jshttp/mime-db/master/db.json](https://raw.githubusercontent.com/jshttp/mime-db/master/db.json)

##### Return values

string\|null
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html#toc-methods)
- Methods
  - [fromExtension()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html#method_fromExtension)
  - [fromFilename()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html#method_fromFilename)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MimeType.html#top)
