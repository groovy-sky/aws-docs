This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule HttpHeaderConfig

Information about an HTTP header condition.

There is a set of standard HTTP header fields. You can also define custom HTTP header
fields.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HttpHeaderName" : String,
  "RegexValues" : [ String, ... ],
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  HttpHeaderName: String
  RegexValues:
    - String
  Values:
    - String

```

## Properties

`HttpHeaderName`

The name of the HTTP header field. The maximum size is 40 characters. The header name is
case insensitive. The allowed characters are specified by RFC 7230. Wildcards are not
supported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegexValues`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The strings to compare against the value of the HTTP header. The maximum length of
each string is 128 characters. The comparison strings are case insensitive. The following
wildcard characters are supported: \* (matches 0 or more characters) and ? (matches exactly 1
character).

If the same header appears multiple times in the request, we search them in order until a
match is found.

If you specify multiple strings, the condition is satisfied if one of the strings matches
the value of the HTTP header. To require that all of the strings are a match, create one
condition per string.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

This example creates a listener rule with an action that redirects requests
with the specified values for the `User-Agent` header to the mobile
version of the website.

#### YAML

```yaml

myHttpHeaderListenerRule:
   Type: 'AWS::ElasticLoadBalancingV2::ListenerRule'
   Properties:
     ListenerArn: !Ref myListener
     Priority: 30
     Conditions:
       - Field: http-header
         HttpHeaderConfig:
           HttpHeaderName: User-Agent
           Values:
             - "*Mobile*"
             - "*Android*"
             - "*iPhone*"
     Actions:
       - Type: redirect
         RedirectConfig:
           Host: m.example.com
           StatusCode: HTTP_302
```

#### JSON

```json

{
    "myHttpHeaderListenerRule": {
        "Type": "AWS::ElasticLoadBalancingV2::ListenerRule",
        "Properties": {
            "ListenerArn": {
                "Ref": "myListener"
            },
            "Priority": 30,
            "Conditions": [
                {
                    "Field": "http-header",
                    "HttpHeaderConfig": {
                        "HttpHeaderName": "User-Agent",
                        "Values": [
                            "*Mobile*",
                            "*Android*",
                            "*iPhone*"
                        ]
                    }
                }
            ],
            "Actions": [
                {
                    "Type": "redirect",
                    "RedirectConfig": {
                        "Host": "m.example.com",
                        "StatusCode": "HTTP_302"
                    }
                }
            ]
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HostHeaderConfig

HttpRequestMethodConfig

All content copied from https://docs.aws.amazon.com/.
