# Configurable parser-type processors

This section contains information about the configurable data parser processors that you can use in a log event
transformer.

###### Contents

- [parseJSON](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#CloudWatch-Logs-Transformation-parseJSON)

- [grok](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#CloudWatch-Logs-Transformation-Grok)

  - [Grok examples](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#Grok-Examples)

    - [Example 1: Use grok to extract a field from unstructured logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#Grok-Example1)

    - [Example 2: Use grok in combination with parseJSON to extract fields from a JSON log event](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#Grok-Example3)

    - [Example 3: Grok pattern with dotted annotation in FIELD\_NAME](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#Grok-Example4)
  - [Supported grok patterns](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#Grok-Patterns)

    - [Common log format examples](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#Common-Log-Examples)

      - [Apache log example](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#Apache-Log-Example)

      - [NGINX log example](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#NGINX-Log-Example)

      - [Syslog Protocol (RFC 5424) log example](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#syslog5424-Log-Example)
- [csv](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#CloudWatch-Logs-Transformation-csv)

- [parseKeyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Configurable.html#CloudWatch-Logs-Transformation-parseKeyValue)

## parseJSON

The **parseJSON** processor parses JSON log events and
inserts extracted JSON key-value pairs under the destination. If you don't
specify a destination, the processor places the key-value pair under the root
node. When using `parseJSON` as the first processor, you must parse
the entire log event using `@message` as the source field. After the
initial JSON parsing, you can then manipulate specific fields in subsequent
processors.

The original `@message` content is not changed, the new keys are
added to the message.

FieldDescriptionRequired?DefaultLimits

source

Path to the field in the log event that will be parsed. Use
dot notation to access child fields. For example,
`store.book`

No

`@message`

Maximum length: 128

Maximum nested key depth: 3

destination

The destination field of the parsed JSON

No

`Parent JSON node`

Maximum length: 128

Maximum nested key depth: 3

**Example**

Suppose an ingested log event looks like this:

```json

{
    "outer_key": {
        "inner_key": "inner_value"
    }
}
```

Then if we have this **parseJSON** processor:

```json

[
   {
        "parseJSON": {
            "destination": "new_key"
        }
   }
]
```

The transformed log event would be the following.

```json

{
    "new_key": {
        "outer_key": {
            "inner_key": "inner_value"
        }
    }
}
```

## grok

Use the grok processor to parse and structure unstructured data using pattern
matching. This processor can also extract fields from log messages.

FieldDescriptionRequired?DefaultLimitsNotes

source

Path of the field to apply Grok matching on

No

`@message`

Maximum length: 128

Maximum nested key depth: 3

match

The grok pattern to match against the log event

Yes

Maximum length: 512

Maximum grok patterns: 20

Some grok pattern types have individual usage limits. Any
combination of the following patterns can be used as many as
five times: {URI, URIPARAM, URIPATHPARAM, SPACE, DATA,
GREEDYDATA, GREEDYDATA\_MULTILINE}

Grok patterns don't support type conversions.

For common log format patterns (APACHE\_ACCESS\_LOG,
NGINX\_ACCESS\_LOG, SYSLOG5424), only DATA, GREEDYDATA, or
GREEDYDATA\_MULTILINE patterns are supported to be included
after the common log pattern.

[See all supported Grok\
patterns](#Grok-Patterns)

**Structure of a Grok Pattern**

This is the supported grok pattern structure:

```nohighlight

%{PATTERN_NAME:FIELD_NAME}
```

- **PATTERN\_NAME**: Refers to a pre-defined
regular expression for matching a specific type of data. Only predefined
[grok patterns](#Grok-Patterns) are supported. Creating custom
patterns is not allowed.

- **FIELD\_NAME**: Assigns a name to the
extracted value. `FIELD_NAME` is optional, but if you don't
specify this value then the extracted data will be dropped from the
transformed log event. If `FIELD_NAME` uses dotted notation
(e.g., "parent.child"), it is considered as a JSON path.

- **Type Conversion**: Explicit type
conversions are not be supported. Use [TypeConverter processor](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Datatype.html#CloudWatch-Logs-Transformation-typeConverter) to convert the datatype of any
value extracted by grok.

To create more complex matching expressions, you can combine several grok
patterns. As many as 20 grok patterns can be combined to match a log event. For
example, this combination of patterns `%{NUMBER:timestamp} [%{NUMBER:db}
                        %{IP:client_ip}:%{NUMBER:client_port}] %{GREEDYDATA:data}` can be used
to extract fields from a Redis slow log entry like this:

`1629860738.123456 [0 127.0.0.1:6379] "SET" "key1" "value1"`

### Grok examples

#### Example 1: Use grok to extract a field from unstructured logs

Sample log:

```json

293750 server-01.internal-network.local OK "[Thread-000] token generated"
```

Transformer used:

```json

[
     {
         "grok": {
             "match": "%{NUMBER:version} %{HOSTNAME:hostname} %{NOTSPACE:status} %{QUOTEDSTRING:logMsg}"
         }
    }
]
```

Output:

```json

{
  "version": "293750",
  "hostname": "server-01.internal-network.local",
  "status": "OK",
  "logMsg": "[Thread-000] token generated"
}
```

Sample log:

```nohighlight

23/Nov/2024:10:25:15 -0900 172.16.0.1 200
```

Transformer used:

```json

[
    {
        "grok": {
            "match": "%{HTTPDATE:timestamp} %{IPORHOST:clientip} %{NUMBER:response_status}"
        }
    }
]
```

Output:

```json

{
  "timestamp": "23/Nov/2024:10:25:15 -0900",
  "clientip": "172.16.0.1",
  "response_status": "200"
}
```

#### Example 2: Use grok in combination with parseJSON to extract fields from a JSON log event

Sample log:

```json

{
    "timestamp": "2024-11-23T16:03:12Z",
    "level": "ERROR",
    "logMsg": "GET /page.html HTTP/1.1"
}
```

Transformer used:

```json

[
     {
        "parseJSON": {}
    },
    {
         "grok": {
            "source": "logMsg",
             "match": "%{WORD:http_method} %{NOTSPACE:request} HTTP/%{NUMBER:http_version}"
         }
    }
]
```

Output:

```json

{
  "timestamp": "2024-11-23T16:03:12Z",
  "level": "ERROR",
  "logMsg": "GET /page.html HTTP/1.1",
  "http_method": "GET",
  "request": "/page.html",
  "http_version": "1.1"
}
```

#### Example 3: Grok pattern with dotted annotation in FIELD\_NAME

Sample log:

```nohighlight

192.168.1.1 GET /index.html?param=value 200 1234
```

Transformer used:

```json

[
    {
        "grok": {
            "match": "%{IP:client.ip} %{WORD:method} %{URIPATHPARAM:request.uri} %{NUMBER:response.status} %{NUMBER:response.bytes}"
        }
    }
]
```

Output:

```json

{
  "client": {
    "ip": "192.168.1.1"
  },
  "method": "GET",
  "request": {
    "uri": "/index.html?param=value"
  },
  "response": {
    "status": "200",
    "bytes": "1234"
  }
}
```

### Supported grok patterns

The following tables list the patterns that are supported by the
`grok` processor.

**General grok patterns**

Grok PatternDescriptionMaximum pattern limitExampleUSERNAME or USERMatches one or more characters that can include lowercase
letters (a-z), uppercase letters (A-Z), digits (0-9), dots
(.), underscores (\_), or hyphens (-)20

Input: `user123.name-TEST`

Pattern: `%{USERNAME:name}`

Output: `{"name":
                                            "user123.name-TEST"}`

INTMatches an optional plus or minus sign followed by one or
more digits.20

Input: `-456`

Pattern: `%{INT:num}`

Output: `{"num": "-456"}`

BASE10NUMMatches an integer or a floating-point number with
optional sign and decimal point20

Input: `-0.67`

Pattern: `%{BASE10NUM:num}`

Output: `{"num": "-0.67"}`

BASE16NUMMatches decimal and hexadecimal numbers with an optional
sign (+ or -) and an optional 0x prefix20

Input: `+0xA1B2`

Pattern: `%{BASE16NUM:num}`

Output: `{"num": "+0xA1B2"}`

POSINTMatches whole positive integers without leading zeros,
consisting of one or more digits (1-9 followed by
0-9)20

Input: `123`

Pattern: `%{POSINT:num}`

Output: `{"num": "123"}`

NONNEGINTMatches any whole numbers (consisting of one or more
digits 0-9) including zero and numbers with leading
zeros.20

Input: `007`

Pattern: `%{NONNEGINT:num}`

Output: `{"num": "007"}`

WORDMatches whole words composed of one or more word
characters (\\w), including letters, digits, and
underscores20

Input: `user_123`

Pattern: `%{WORD:user}`

Output: `{"user": "user_123"}`

NOTSPACEMatches one or more non-whitespace characters.5

Input: `hello_world123`

Pattern: `%{NOTSPACE:msg}`

Output: `{"msg": "hello_world123"}`

SPACEMatches zero or more whitespace characters.5

Input: `" "`

Pattern: `%{SPACE:extra}`

Output: `{"extra": " "}`

DATAMatches any character (except newline) zero or more
times, non-greedy.5

Input: `abc def ghi`

Pattern: `%{DATA:x} %{DATA:y}`

Output: `{"x": "abc", "y": "def
                                            ghi"}`

GREEDYDATAMatches any character (except newline) zero or more
times, greedy.5

Input: `abc def ghi`

Pattern: `%{GREEDYDATA:x}
                                            %{GREEDYDATA:y}`

Output: `{"x": "abc def", "y":
                                            "ghi"}`

GREEDYDATA\_MULTILINEMatches any character (including newline) zero or more
times, greedy.1

Input:

`abc`

`def`

`ghi`

Pattern:
`%{GREEDYDATA_MULTILINE:data}`

Output: `{"data": "abc\ndef\nghi"}`

QUOTEDSTRINGMatches quoted strings (single or double quotes) with
escaped characters.20

Input: `"Hello, world!"`

Pattern: `%{QUOTEDSTRING:msg}`

Output: `{"msg": "Hello, world!"}`

UUIDMatches a standard UUID format: 8 hexadecimal characters,
followed by three groups of 4 hexadecimal characters, and
ending with 12 hexadecimal characters, all separated by
hyphens.20

Input:
`550e8400-e29b-41d4-a716-446655440000`

Pattern: `%{UUID:id}`

Output: `{"id":
                                                "550e8400-e29b-41d4-a716-446655440000"}`

URNMatches URN (Uniform Resource Name) syntax.20

Input: `urn:isbn:0451450523`

Pattern: `%{URN:urn}`

Output: `{"urn":
                                            "urn:isbn:0451450523"}`

**AWS grok patterns**

PatternDescriptionMaximum pattern limitExample

ARN

Matches AWS Amazon Resource Names (ARNs), capturing
the partition ( `aws`, `aws-cn`, or
`aws-us-gov`), service, Region, account
ID, and up to 5 hierarchical resource identifiers
separated by slashes. It will not match ARNs that are
missing information between colons.

5

Input:
`arn:aws:iam:us-east-1:123456789012:user/johndoe`

Pattern: `%{ARN:arn}`

Output: `{"arn":
                                                "arn:aws:iam:us-east-1:123456789012:user/johndoe"}`

**Networking grok patterns**

Grok PatternDescriptionMaximum pattern limitExampleCISCOMACMatches a MAC address in 4-4-4 hexadecimal
format.20

Input: `0123.4567.89AB`

Pattern: `%{CISCOMAC:MacAddress}`

Output: `{"MacAddress":
                                            "0123.4567.89AB"}`

WINDOWSMACMatches a MAC address in hexadecimal format with
hyphens20

Input: `01-23-45-67-89-AB`

Pattern: `%{WINDOWSMAC:MacAddress}`

Output: `{"MacAddress":
                                                "01-23-45-67-89-AB"}`

COMMONMACMatches a MAC address in hexadecimal format with
colons.20

Input: `01:23:45:67:89:AB`

Pattern: `%{COMMONMAC:MacAddress}`

Output: `{"MacAddress":
                                                "01:23:45:67:89:AB"}`

MACMatches one of CISCOMAC, WINDOWSMAC or COMMONMAC grok
patterns20

Input: `01:23:45:67:89:AB`

Pattern: `%{MAC:m1}`

Output: `{"m1":"01:23:45:67:89:AB"}`

IPV6Matches IPv6 addresses, including compressed forms and
IPv4-mapped IPv6 addresses.5

Input:
`2001:db8:3333:4444:5555:6666:7777:8888`

Pattern: `%{IPV6:ip}`

Output: `{"ip":
                                                "2001:db8:3333:4444:5555:6666:7777:8888"}`

IPV4Matches an IPv4 address.20

Input: `192.168.0.1`

Pattern: `%{IPV4:ip}`

Output: `{"ip": "192.168.0.1"}`

IPMatches either IPv6 addresses as supported by %{IPv6} or
IPv4 addresses as supported by %{IPv4}5

Input: `192.168.0.1`

Pattern: `%{IP:ip}`

Output: `{"ip": "192.168.0.1"}`

HOSTNAME or HOSTMatches domain names, including subdomains5

Input:
`server-01.internal-network.local`

Pattern: `%{HOST:host}`

Output: `{"host":
                                                "server-01.internal-network.local"}`

IPORHOSTMatches either a hostname or an IP address5

Input:
`2001:db8:3333:4444:5555:6666:7777:8888`

Pattern: `%{IPORHOST:ip}`

Output: `{"ip":
                                                "2001:db8:3333:4444:5555:6666:7777:8888"}`

HOSTPORTMatches an IP address or hostname as supported by
%{IPORHOST} pattern followed by a colon and a port number,
capturing the port as "PORT" in the output.5

Input: `192.168.1.1:8080`

Pattern: `%{HOSTPORT:ip}`

Output:
`{"ip":"192.168.1.1:8080","PORT":"8080"}`

URIHOSTMatches an IP address or hostname as supported by
%{IPORHOST} pattern, optionally followed by a colon and a
port number, capturing the port as "port" if
present.5

Input: `example.com:443 10.0.0.1`

Pattern: `%{URIHOST:host}
                                            %{URIHOST:ip}`

Output:
`{"host":"example.com:443","port":"443","ip":"10.0.0.1"}`

**Path grok patterns**

Grok PatternDescriptionMaximum pattern limitExampleUNIXPATHMatches URL paths, potentially including query
parameters.20

Input: `/search?q=regex`

Pattern: `%{UNIXPATH:path}`

Output: `{"path":"/search?q=regex"}`

WINPATHMatches Windows file paths.5

Input:
`C:\Users\John\Documents\file.txt`

Pattern: `%{WINPATH:path}`

Output: `{"path":
                                                "C:\\Users\\John\\Documents\\file.txt"}`

PATHMatches either URL or Windows file paths5

Input: `/search?q=regex`

Pattern: `%{PATH:path}`

Output: `{"path":"/search?q=regex"}`

TTYMatches Unix device paths for terminals and
pseudo-terminals.20

Input: `/dev/tty1`

Pattern: `%{TTY:path}`

Output: `{"path":"/dev/tty1"}`

URIPROTOMatches letters, optionally followed by a plus (+)
character and additional letters or plus (+)
characters20

Input: `web+transformer`

Pattern: `%{URIPROTO:protocol}`

Output:
`{"protocol":"web+transformer"}`

URIPATHMatches the path component of a URI20

Input:
`/category/sub-category/product_name`

Pattern: `%{URIPATH:path}`

Output:
`{"path":"/category/sub-category/product_name"}`

URIPARAMMatches URL query parameters5

Input:
`?param1=value1&param2=value2`

Pattern: `%{URIPARAM:url}`

Output:
`{"url":"?param1=value1&param2=value2"}`

URIPATHPARAMMatches a URI path optionally followed by query
parameters5

Input:
`/category/sub-category/product?id=12345&color=red`

Pattern: `%{URIPATHPARAM:path}`

Output:
`{"path":"/category/sub-category/product?id=12345&color=red"}`

URIMatches a complete URI5

Input:
`https://user:password@example.com/path/to/resource?param1=value1&param2=value2`

Pattern: `%{URI:uri}`

Output:
`{"path":"https://user:password@example.com/path/to/resource?param1=value1&param2=value2"}`

**Date and time grok patterns**

Grok PatternDescriptionMaximum pattern limitExampleMONTHMatches full or abbreviated english month names as whole
words20

Input: `Jan`

Pattern: `%{MONTH:month}`

Output: `{"month":"Jan"}`

Input: `January`

Pattern: `%{MONTH:month}`

Output: `{"month":"January"}`

MONTHNUMMatches month numbers from 1 to 12, with optional leading
zero for single-digit months.20

Input: `5`

Pattern: `%{MONTHNUM:month}`

Output: `{"month":"5"}`

Input: `05`

Pattern: `%{MONTHNUM:month}`

Output: `{"month":"05"}`

MONTHNUM2Matches two-digit month numbers from 01 to 12.20

Input: `05`

Pattern: `%{MONTHNUM2:month}`

Output: `{"month":"05"}`

MONTHDAYMatches day of the month from 1 to 31, with optional
leading zero.20

Input: `31`

Pattern: `%{MONTHDAY:monthDay}`

Output: `{"monthDay":"31"}`

YEARMatches year in two or four digits20

Input: `2024`

Pattern: `%{YEAR:year}`

Output: `{"year":"2024"}`

Input: `24`

Pattern: `%{YEAR:year}`

Output: `{"year":"24"}`

DAYMatches full or abbreviated day names.20

Input: `Tuesday`

Pattern: `%{DAY:day}`

Output: `{"day":"Tuesday"}`

HOURMatches hour in 24-hour format with an optional leading
zero (0)0-23.20

Input: `22`

Pattern: `%{HOUR:hour}`

Output: `{"hour":"22"}`

MINUTEMatches minutes (00-59).20

Input: `59`

Pattern: `%{MINUTE:min}`

Output: `{"min":"59"}`

SECONDMatches a number representing seconds (0)0-60, optionally
followed by a decimal point or colon and one or more digits
for fractional minutes20

Input: `3`

Pattern: `%{SECOND:second}`

Output: `{"second":"3"}`

Input: `30.5`

Pattern: `%{SECOND:minSec}`

Output: `{"minSec":"30.5"}`

Input: `30:5`

Pattern: `%{SECOND:minSec}`

Output: `{"minSec":"30:5"}`

TIMEMatches a time format with hours, minutes, and seconds in
the format (H)H:mm:(s)s. Seconds include leap second
(0)0-60.20

Input: `09:45:32`

Pattern: `%{TIME:time}`

Output: `{"time":"09:45:32"}`

DATE\_USMatches a date in the format of (M)M/(d)d/(yy)yy or
(M)M-(d)d-(yy)yy.20

Input: `11/23/2024`

Pattern: `%{DATE_US:date}`

Output: `{"date":"11/23/2024"}`

Input: `1-01-24`

Pattern: `%{DATE_US:date}`

Output: `{"date":"1-01-24"}`

DATE\_EUMatches date in format of (d)d/(M)M/(yy)yy,
(d)d-(M)M-(yy)yy, or (d)d.(M)M.(yy)yy.20

Input: `23/11/2024`

Pattern: `%{DATE_EU:date}`

Output: `{"date":"23/11/2024"}`

Input: `1.01.24`

Pattern: `%{DATE_EU:date}`

Output: `{"date":"1.01.24"}`

ISO8601\_TIMEZONEMatches UTC offset 'Z' or time zone offset with optional
colon in format \[+-\](H)H(:)mm.20

Input: `+05:30`

Pattern: `%{ISO8601_TIMEZONE:tz}`

Output: `{"tz":"+05:30"}`

Input: `-530`

Pattern: `%{ISO8601_TIMEZONE:tz}`

Output: `{"tz":"-530"}`

Input: `Z`

Pattern: `%{ISO8601_TIMEZONE:tz}`

Output: `{"tz":"Z"}`

ISO8601\_SECONDMatches a number representing seconds (0)0-60, optionally
followed by a decimal point or colon and one or more digits
for fractional seconds20

Input: `60`

Pattern: `%{ISO8601_SECOND:second}`

Output: `{"second":"60"}`

TIMESTAMP\_ISO8601Matches ISO8601 datetime format
(yy)yy-(M)M-(d)dT(H)H:mm:((s)s)(Z\|\[+-\](H)H:mm) with optional
seconds and timezone.20

Input: `2023-05-15T14:30:00+05:30`

Pattern:
`%{TIMESTAMP_ISO8601:timestamp}`

Output:
`{"timestamp":"2023-05-15T14:30:00+05:30"}`

Input: `23-5-1T1:25+5:30`

Pattern:
`%{TIMESTAMP_ISO8601:timestamp}`

Output:
`{"timestamp":"23-5-1T1:25+5:30"}`

Input: `23-5-1T1:25Z`

Pattern:
`%{TIMESTAMP_ISO8601:timestamp}`

Output:
`{"timestamp":"23-5-1T1:25Z"}`

DATEMatches either a date in the US format using %{DATE\_US}
or in the EU format using %{DATE\_EU}20

Input: `11/29/2024`

Pattern: `%{DATE:date}`

Output: `{"date":"11/29/2024"}`

Input: `29.11.2024`

Pattern: `%{DATE:date}`

Output: `{"date":"29.11.2024"}`

DATESTAMPMatches %{DATE} followed by %{TIME} pattern, separated by
space or hyphen.20

Input: `29-11-2024 14:30:00`

Pattern: `%{DATESTAMP:dateTime}`

Output: `{"dateTime":"29-11-2024
                                                14:30:00"}`

TZMatches common time zone abbreviations (PST, PDT, MST,
MDT, CST CDT, EST, EDT, UTC).20

Input: `PDT`

Pattern: `%{TZ:tz}`

Output: `{"tz":"PDT"}`

DATESTAMP\_RFC822Matches date and time in format: Day MonthName (D)D
(YY)YY (H)H:mm:(s)s Timezone20

Input: `Monday Jan 5 23 1:30:00 CDT`

Pattern:
`%{DATESTAMP_RFC822:dateTime}`

Output: `{"dateTime":"Monday Jan 5 23 1:30:00
                                                CDT"}`

Input: `Mon January 15 2023 14:30:00
                                            PST`

Pattern:
`%{DATESTAMP_RFC822:dateTime}`

Output: `{"dateTime":"Mon January 15 2023
                                                14:30:00 PST"}`

DATESTAMP\_RFC2822Matches RFC2822 date-time format: Day, (d)d MonthName
(yy)yy (H)H:mm:(s)s Z\|\[+-\](H)H:mm20

Input: `Mon, 15 May 2023 14:30:00
                                            +0530`

Pattern:
`%{DATESTAMP_RFC2822:dateTime}`

Output: `{"dateTime":"Mon, 15 May 2023 14:30:00
                                                +0530"}`

Input: `Monday, 15 Jan 23 14:30:00
                                            Z`

Pattern:
`%{DATESTAMP_RFC2822:dateTime}`

Output: `{"dateTime":"Monday, 15 Jan 23 14:30:00
                                                Z"}`

DATESTAMP\_OTHERMatches date and time in format: Day MonthName (d)d
(H)H:mm:(s)s Timezone (yy)yy20

Input: `Mon May 15 14:30:00 PST
                                            2023`

Pattern:
`%{DATESTAMP_OTHER:dateTime}`

Output: `{"dateTime":"Mon May 15 14:30:00 PST
                                                2023"}`

DATESTAMP\_EVENTLOGMatches compact datetime format without separators:
(yy)yyMM(d)d(H)Hmm(s)s20

Input: `20230515143000`

Pattern:
`%{DATESTAMP_EVENTLOG:dateTime}`

Output:
`{"dateTime":"20230515143000"}`

**Log grok patterns**

Grok PatternDescriptionMaximum pattern limitExampleLOGLEVELMatches standard log levels in different capitalizations
and abbreviations, including the following:
`Alert/ALERT`, `Trace/TRACE`,
`Debug/DEBUG`, `Notice/NOTICE`,
`Info/INFO`,
`Warn/Warning/WARN/WARNING`,
`Err/Error/ERR/ERROR`,
`Crit/Critical/CRIT/CRITICAL`,
`Fatal/FATAL`, `Severe/SEVERE`,
`Emerg/Emergency/EMERG/EMERGENCY`20

Input: `INFO`

Pattern: `%{LOGLEVEL:logLevel}`

Output: `{"logLevel":"INFO"}`

HTTPDATEMatches date and time format often used in log files.
Format: (d)d/MonthName/(yy)yy:(H)H:mm:(s)s Timezone
MonthName: Matches full or abbreviated english month names
(Example: "Jan" or "January") Timezone: Matches %{INT} grok
pattern20

Input: `23/Nov/2024:14:30:00 +0640`

Pattern: `%{HTTPDATE:date}`

Output: `{"date":"23/Nov/2024:14:30:00
                                                +0640"}`

SYSLOGTIMESTAMPMatches date format with MonthName (d)d (H)H:mm:(s)s
MonthName: Matches full or abbreviated english month names
(Example: "Jan" or "January")20

Input: `Nov 29 14:30:00`

Pattern:
`%{SYSLOGTIMESTAMP:dateTime}`

Output: `{"dateTime":"Nov 29
                                            14:30:00"}`

PROGMatches a program name consisting of string of letters,
digits, dot, underscore, forward slash, percent sign, and
hyphen characters.20

Input: `user.profile/settings-page`

Pattern: `%{PROG:program}`

Output:
`{"program":"user.profile/settings-page"}`

SYSLOGPROGMatches PROG grok pattern optionally followed by a
process ID in square brackets.20

Input:
`user.profile/settings-page[1234]`

Pattern:
`%{SYSLOGPROG:programWithId}`

Output:
`{"programWithId":"user.profile/settings-page[1234]","program":"user.profile/settings-page","pid":"1234"}`

SYSLOGHOSTMatches either a %{HOST} or %{IP} pattern5

Input:
`2001:db8:3333:4444:5555:6666:7777:8888`

Pattern: `%{SYSLOGHOST:ip}`

Output: `{"ip":
                                                "2001:db8:3333:4444:5555:6666:7777:8888"}`

SYSLOGFACILITYMatches syslog priority in decimal format. The value
should be enclosed in angular brackets (<>).20

Input: `<13.6>`

Pattern: `%{SYSLOGFACILITY:syslog}`

Output:
`{"syslog":"<13.6>","facility":"13","priority":"6"}`

**Common log grok patterns**

You can use pre-defined custom grok patterns to match Apache, NGINX and
Syslog Protocol (RFC 5424) log formats. When you use these specific
patterns, they must be the first patterns in your matching configuration,
and no other patterns can precede them. Also, you can follow them only with
exactly one **DATA**. **GREEDYDATA** or
**GREEDYDATA\_MULTILINE** pattern.

Grok patternDescriptionMaximum pattern limit

APACHE\_ACCESS\_LOG

Matches Apache access logs

1

NGINX\_ACCESS\_LOG

Matches NGINX access logs

1

SYSLOG5424

Matches Syslog Protocol (RFC 5424)
logs

1

The following shows valid and invalid examples for using these common log
format patterns.

```nohighlight

"%{NGINX_ACCESS_LOG} %{DATA}" // Valid
"%{SYSLOG5424}%{DATA:logMsg}" // Valid
"%{APACHE_ACCESS_LOG} %{GREEDYDATA:logMsg}" // Valid
"%{APACHE_ACCESS_LOG} %{SYSLOG5424}" // Invalid (multiple common log patterns used)
"%{NGINX_ACCESS_LOG} %{NUMBER:num}" // Invalid (Only GREEDYDATA and DATA patterns are supported with common log patterns)
"%{GREEDYDATA:logMsg} %{SYSLOG5424}" // Invalid (GREEDYDATA and DATA patterns are supported only after common log patterns)
```

#### Common log format examples

##### Apache log example

Sample log:

```nohighlight

127.0.0.1 - - [03/Aug/2023:12:34:56 +0000] "GET /page.html HTTP/1.1" 200 1234
```

Transformer:

```json

[
     {
        "grok": {
            "match": "%{APACHE_ACCESS_LOG}"
        }
    }
]
```

Output:

```json

{
    "request": "/page.html",
    "http_method": "GET",
    "status_code": 200,
    "http_version": "1.1",
    "response_size": 1234,
    "remote_host": "127.0.0.1",
    "timestamp": "2023-08-03T12:34:56Z"
}
```

##### NGINX log example

Sample log:

```nohighlight

192.168.1.100 - Foo [03/Aug/2023:12:34:56 +0000] "GET /account/login.html HTTP/1.1" 200 42 "https://www.amazon.com/" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"
```

Transformer:

```json

[
     {
        "grok": {
            "match": "%{NGINX_ACCESS_LOG}"
        }
    }
]
```

Output:

```json

{
    "request": "/account/login.html",
    "referrer": "https://www.amazon.com/",
    "agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",
    "http_method": "GET",
    "status_code": 200,
    "auth_user": "Foo",
    "http_version": "1.1",
    "response_size": 42,
    "remote_host": "192.168.1.100",
    "timestamp": "2023-08-03T12:34:56Z"
}
```

##### Syslog Protocol (RFC 5424) log example

Sample log:

```nohighlight

<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource= "Application" eventID="1011"][examplePriority@32473 class="high"]
```

Transformer:

```json

[
     {
        "grok": {
            "match": "%{SYSLOG5424}"
        }
    }
]
```

Output:

```json

{
  "pri": 165,
  "version": 1,
  "timestamp": "2003-10-11T22:14:15.003Z",
  "hostname": "mymachine.example.com",
  "app": "evntslog",
  "msg_id": "ID47",
  "structured_data": "exampleSDID@32473 iut=\"3\" eventSource= \"Application\" eventID=\"1011\"",
  "message": "[examplePriority@32473 class=\"high\"]"
}
```

## csv

The **csv** processor parses comma-separated values (CSV)
from the log events into columns.

FieldDescriptionRequired?DefaultLimits

source

Path to the field in the log event that will be
parsed

No

`@message`

Maximum length: 128

Maximum nested key depth: 3

delimiter

The character used to separate each column in the original
comma-separated value log event

No

`,`

Maximum length: 1 unless the value is `\t`
or `\s`

quoteCharacter

Character used as a text qualifier for a single column of
data

No

`"`

Maximum length: 1

columns

List of names to use for the columns in the transformed log
event.

No

`[column_1, column_2 ...]`

Maximum CSV columns: 100

Maximum length: 128

Maximum nested key depth: 3

destination

The parent field to put transformed key value pairs under

No

`Root node`

Maximum length: 128

Maximum nested key depth: 3

Setting `delimiter` to `\t` will separate each column on
a tab character, and `\t` will separate each column on a single space
character.

**Example**

Suppose part of an ingested log event looks like this:

```json

'Akua Mansa':28:'New York: USA'
```

Suppose we use only the **csv** processor:

```json

[
     {
        "csv": {
            "delimiter": ":",
            "quoteCharacter": "'"
        }
    }
]
```

The transformed log event would be the following.

```json

{
  "column_1": "Akua Mansa",
  "column_2": "28",
  "column_3": "New York: USA"
}
```

**Example 2**

Suppose an ingested log event looks like this:

```json

{
    "timestamp": "2024-11-23T16:03:12Z",
    "type": "user_data",
    "logMsg": "'Akua Mansa':28:'New York: USA'"
}
```

Suppose we parse the event as JSON, they parse a JSON field with the **csv** processor, specifying column names and destination:

```json

[
    {
        "parseJSON": {}
    },
    {
        "csv": {
            "source": "logMsg",
            "delimiter": ":",
            "quoteCharacter": "'",
            "columns":["name","age","location"],
            "destination": "msg"
        }
    }
]
```

The transformed log event would be the following.

```json

{
    "timestamp": "2024-11-23T16:03:12Z",
    "logMsg": "'Akua Mansa':28:'New York: USA'",
    "type": "user_data",
    "msg": {
        "name": "Akua Mansa",
        "age": "28",
        "location": "New York: USA"
    }
}
```

## parseKeyValue

Use the **parseKeyValue** processor to parse a specified
field into key-value pairs. You can customize the processor to parse field
information with the following options.

FieldDescriptionRequired?DefaultLimits

source

Path to the field in the log event that will be
parsed

No

`@message`

Maximum length: 128

Maximum nested key depth: 3

destination

The destination field to put the extracted key-value pairs
into

No

Maximum length: 128

fieldDelimiter

The field delimiter string that is used between key-value
pairs in the original log events

No

`&`

Maximum length: 128

keyValueDelimiter

The delimiter string to use between the key and value in each
pair in the transformed log event

No

`=`

Maximum length: 128

nonMatchValue

A value to insert into the value field in the result,
when a key-value pair is not successfully
split.

No

Maximum length: 128

keyPrefix

If you want to add a prefix toall transformed keys,
specify it here.

No

Maximum length: 128

overwriteIfExists

Whether to overwrite the value if the destination key already
exists

No

`false`

**Example**

Take the following example log event:

```json

key1:value1!key2:value2!key3:value3!key4
```

Suppose we use the following processor configuration:

```json

[
    {
        "parseKeyValue": {
            "destination": "new_key",
            "fieldDelimiter": "!",
            "keyValueDelimiter": ":",
            "nonMatchValue": "defaultValue",
            "keyPrefix": "parsed_"
        }
    }
]
```

The transformed log event would be the following.

```json

{
  "new_key": {
    "parsed_key1": "value1",
    "parsed_key2": "value2",
    "parsed_key3": "value3",
    "parsed_key4": "defaultValue"
  }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Delete a log-group-level transformer

Built-in processors for AWS vended logs
