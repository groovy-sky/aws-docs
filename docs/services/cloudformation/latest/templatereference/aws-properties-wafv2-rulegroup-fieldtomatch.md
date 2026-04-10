This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup FieldToMatch

Specifies a web request component to be used in a rule match statement or in a logging configuration.

- In a rule statement, this is the part of the web request that you want AWS WAF to inspect. Include the single
`FieldToMatch` type that you want to inspect, with additional specifications
as needed, according to the type. You specify a single request component in
`FieldToMatch` for each rule statement that requires it. To inspect more than
one component of the web request, create a separate rule statement for each
component.

Example JSON for a `QueryString` field to match:

` "FieldToMatch": { "QueryString": {} }`

Example JSON for a `Method` field to match specification:

` "FieldToMatch": { "Method": { "Name": "DELETE" } }`

- In a logging configuration, this is used in the `RedactedFields` property to specify a field to
redact from the logging records. For this use case, note the following:

- Even though all `FieldToMatch` settings
are available, the only valid settings for field redaction are `UriPath`, `QueryString`, `SingleHeader`, and `Method`.

- In this documentation, the descriptions of the individual fields talk about specifying the web request component to inspect,
but for field redaction, you are specifying the component type to redact from the logs.

- If you have request sampling enabled, the redacted fields configuration for logging has no impact on sampling.
You can only exclude fields from request sampling by disabling sampling in the web ACL visibility configuration
or by configuring data protection for the web ACL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllQueryArguments" : Json,
  "Body" : Body,
  "Cookies" : Cookies,
  "HeaderOrder" : HeaderOrder,
  "Headers" : Headers,
  "JA3Fingerprint" : JA3Fingerprint,
  "JA4Fingerprint" : JA4Fingerprint,
  "JsonBody" : JsonBody,
  "Method" : Json,
  "QueryString" : Json,
  "SingleHeader" : SingleHeader,
  "SingleQueryArgument" : SingleQueryArgument,
  "UriFragment" : UriFragment,
  "UriPath" : Json
}

```

### YAML

```yaml

  AllQueryArguments: Json
  Body:
    Body
  Cookies:
    Cookies
  HeaderOrder:
    HeaderOrder
  Headers:
    Headers
  JA3Fingerprint:
    JA3Fingerprint
  JA4Fingerprint:
    JA4Fingerprint
  JsonBody:
    JsonBody
  Method: Json
  QueryString: Json
  SingleHeader:
    SingleHeader
  SingleQueryArgument:
    SingleQueryArgument
  UriFragment:
    UriFragment
  UriPath: Json

```

## Properties

`AllQueryArguments`

Inspect all query arguments.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Body`

Inspect the request body as plain text. The request body immediately follows the request
headers. This is the part of a request that contains any additional data that you want to
send to your web server as the HTTP request body, such as data from a form.

AWS WAF does not support inspecting the entire contents of the web request body if the body
exceeds the limit for the resource type. When a web request body is larger than the limit, the underlying host service
only forwards the contents that are within the limit to AWS WAF for inspection.

- For Application Load Balancer and AWS AppSync, the limit is fixed at 8 KB (8,192 bytes).

- For CloudFront, API Gateway, Amazon Cognito, App Runner, and Verified Access, the default limit is 16 KB (16,384 bytes), and
you can increase the limit for each resource type in the web ACL `AssociationConfig`, for additional processing fees.

- For AWS Amplify, use the CloudFront limit.

For information about how to handle oversized
request bodies, see the `Body` object configuration.

_Required_: No

_Type_: [Body](aws-properties-wafv2-rulegroup-body.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cookies`

Inspect the request cookies. You must configure scope and pattern matching filters in
the `Cookies` object, to define the set of cookies and the parts of the cookies
that AWS WAF inspects.

Only the first 8 KB (8192 bytes) of a request's cookies and only the first 200 cookies
are forwarded to AWS WAF for inspection by the underlying host service. You must
configure how to handle any oversize cookie content in the `Cookies` object.
AWS WAF applies the pattern matching filters to the cookies that it receives from the
underlying host service.

_Required_: No

_Type_: [Cookies](aws-properties-wafv2-rulegroup-cookies.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderOrder`

Inspect a string containing the list of the request's header names, ordered as they appear in the web request
that AWS WAF receives for inspection.
AWS WAF generates the string and then uses that as the field to match component in its inspection.
AWS WAF separates the header names in the string using colons and no added spaces, for example `host:user-agent:accept:authorization:referer`.

_Required_: No

_Type_: [HeaderOrder](aws-properties-wafv2-rulegroup-headerorder.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Headers`

Inspect the request headers. You must configure scope and pattern matching filters in
the `Headers` object, to define the set of headers to and the parts of the
headers that AWS WAF inspects.

Only the first 8 KB (8192 bytes) of a request's headers and only the first 200 headers
are forwarded to AWS WAF for inspection by the underlying host service. You must
configure how to handle any oversize header content in the `Headers` object.
AWS WAF applies the pattern matching filters to the headers that it receives from the
underlying host service.

_Required_: No

_Type_: [Headers](aws-properties-wafv2-rulegroup-headers.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JA3Fingerprint`

Available for use with Amazon CloudFront distributions and Application Load Balancers. Match against the request's JA3 fingerprint. The JA3 fingerprint is a 32-character hash derived from the TLS Client Hello of an incoming request. This fingerprint serves as a unique identifier for the client's TLS configuration. AWS WAF calculates and logs this fingerprint for each
request that has enough TLS Client Hello information for the calculation. Almost
all web requests include this information.

###### Note

You can use this choice only with a string match `ByteMatchStatement` with the `PositionalConstraint` set to
`EXACTLY`.

You can obtain the JA3 fingerprint for client requests from the web ACL logs.
If AWS WAF is able to calculate the fingerprint, it includes it in the logs.
For information about the logging fields,
see [Log fields](../../../waf/latest/developerguide/logging-fields.md) in the _AWS WAF Developer Guide_.

Provide the JA3 fingerprint string from the logs in your string match statement
specification, to match with any future requests that have the same TLS configuration.

_Required_: No

_Type_: [JA3Fingerprint](aws-properties-wafv2-rulegroup-ja3fingerprint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JA4Fingerprint`

Available for use with Amazon CloudFront distributions and Application Load Balancers. Match against the request's JA4 fingerprint. The JA4 fingerprint is a 36-character hash derived from the TLS Client Hello of an incoming request. This fingerprint serves as a unique identifier for the client's TLS configuration. AWS WAF calculates and logs this fingerprint for each
request that has enough TLS Client Hello information for the calculation. Almost
all web requests include this information.

###### Note

You can use this choice only with a string match `ByteMatchStatement` with the `PositionalConstraint` set to
`EXACTLY`.

You can obtain the JA4 fingerprint for client requests from the web ACL logs.
If AWS WAF is able to calculate the fingerprint, it includes it in the logs.
For information about the logging fields,
see [Log fields](../../../waf/latest/developerguide/logging-fields.md) in the _AWS WAF Developer Guide_.

Provide the JA4 fingerprint string from the logs in your string match statement
specification, to match with any future requests that have the same TLS configuration.

_Required_: No

_Type_: [JA4Fingerprint](aws-properties-wafv2-rulegroup-ja4fingerprint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JsonBody`

Inspect the request body as JSON. The request body immediately follows the request
headers. This is the part of a request that contains any additional data that you want to
send to your web server as the HTTP request body, such as data from a form.

AWS WAF does not support inspecting the entire contents of the web request body if the body
exceeds the limit for the resource type. When a web request body is larger than the limit, the underlying host service
only forwards the contents that are within the limit to AWS WAF for inspection.

- For Application Load Balancer and AWS AppSync, the limit is fixed at 8 KB (8,192 bytes).

- For CloudFront, API Gateway, Amazon Cognito, App Runner, and Verified Access, the default limit is 16 KB (16,384 bytes), and
you can increase the limit for each resource type in the web ACL `AssociationConfig`, for additional processing fees.

- For AWS Amplify, use the CloudFront limit.

For information about how to handle oversized
request bodies, see the `JsonBody` object configuration.

_Required_: No

_Type_: [JsonBody](aws-properties-wafv2-rulegroup-jsonbody.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Method`

Inspect the HTTP method. The method indicates the type of operation that the request is
asking the origin to perform.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

Inspect the query string. This is the part of a URL that appears after a `?`
character, if any.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleHeader`

Inspect a single header. Provide the name of the header to inspect, for example,
`User-Agent` or `Referer`. This setting isn't case
sensitive.

Example JSON: `"SingleHeader": { "Name": "haystack" }`

Alternately, you can filter and inspect all headers with the `Headers` `FieldToMatch` setting.

_Required_: No

_Type_: [SingleHeader](aws-properties-wafv2-rulegroup-singleheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleQueryArgument`

Inspect a single query argument. Provide the name of the query argument to inspect, such
as _UserName_ or _SalesRegion_. The name can be up to
30 characters long and isn't case sensitive.

Example JSON: `"SingleQueryArgument": { "Name": "myArgument" }`

_Required_: No

_Type_: [SingleQueryArgument](aws-properties-wafv2-rulegroup-singlequeryargument.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UriFragment`

Inspect fragments of the request URI. You must configure scope and pattern matching filters in
the `UriFragment` object, to define the fragment of a URI that AWS WAF inspects.

Only the first 8 KB (8192 bytes) of a request's URI fragments and only the first 200 URI fragments
are forwarded to AWS WAF for inspection by the underlying host service. You must
configure how to handle any oversize URI fragment content in the `UriFragment` object.
AWS WAF applies the pattern matching filters to the cookies that it receives from the
underlying host service.

_Required_: No

_Type_: [UriFragment](aws-properties-wafv2-rulegroup-urifragment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UriPath`

Inspect the request URI path. This is the part of the web request that identifies a
resource, for example, `/images/daily-ad.jpg`.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Set the field to match to QueryString](#aws-properties-wafv2-rulegroup-fieldtomatch--examples--Set_the_field_to_match_to_QueryString)

- [Set the field to match to Method](#aws-properties-wafv2-rulegroup-fieldtomatch--examples--Set_the_field_to_match_to_Method)

### Set the field to match to `QueryString`

The following shows an example field to match specification for a setting that
doesn't requires additional configuration.

#### YAML

```yaml

FieldToMatch:
  QueryString: {}
```

#### JSON

```json

"FieldToMatch":
    {
         "QueryString": {}
    }
```

### Set the field to match to `Method`

The following shows an example field to match specification for a setting that has
additional configuration requirements.

#### YAML

```yaml

FieldToMatch:
  Method:
     Name: DELETE
```

#### JSON

```json

"FieldToMatch":
{
    "Method":
    {
         "Name": "DELETE"
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomResponseBody

ForwardedIPConfiguration

All content copied from https://docs.aws.amazon.com/.
