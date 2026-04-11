Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## MimeType        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-guzzlehttp-psr7-mimetype-toc.md)

#### Methods  [header link](class-guzzlehttp-psr7-mimetype-toc-methods.md)

[fromExtension()](class-guzzlehttp-psr7-mimetype-method-fromextension.md)
: string\|null Maps a file extensions to a mimetype.[fromFilename()](class-guzzlehttp-psr7-mimetype-method-fromfilename.md)
: string\|null Determines the mimetype of a file by looking at its extension.

### Methods  [header link](class-guzzlehttp-psr7-mimetype-methods.md)

#### fromExtension()  [header link](class-guzzlehttp-psr7-mimetype-method-fromextension.md)

Maps a file extensions to a mimetype.

`
    public
            static        fromExtension(string $extension) : string|null`

##### Parameters

$extension
: string

##### Tags  [header link](class-guzzlehttp-psr7-mimetype-method-fromextension-tags.md)

see[https://raw.githubusercontent.com/jshttp/mime-db/master/db.json](https://raw.githubusercontent.com/jshttp/mime-db/master/db.json)

##### Return values

string\|null

#### fromFilename()  [header link](class-guzzlehttp-psr7-mimetype-method-fromfilename.md)

Determines the mimetype of a file by looking at its extension.

`
    public
            static        fromFilename(string $filename) : string|null`

##### Parameters

$filename
: string

##### Tags  [header link](class-guzzlehttp-psr7-mimetype-method-fromfilename-tags.md)

see[https://raw.githubusercontent.com/jshttp/mime-db/master/db.json](https://raw.githubusercontent.com/jshttp/mime-db/master/db.json)

##### Return values

string\|null
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-psr7-mimetype-toc-methods.md)
- Methods
  - [fromExtension()](class-guzzlehttp-psr7-mimetype-method-fromextension.md)
  - [fromFilename()](class-guzzlehttp-psr7-mimetype-method-fromfilename.md)

[Back To Top](class-guzzlehttp-psr7-mimetype-top.md)
