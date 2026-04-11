This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule RedirectConfig

Information about a redirect action.

A URI consists of the following components: protocol://hostname:port/path?query. You must
modify at least one of the following components to avoid a redirect loop: protocol, hostname,
port, or path. Any components that you do not modify retain their original values.

You can reuse URI components using the following reserved keywords:

- #{protocol}

- #{host}

- #{port}

- #{path} (the leading "/" is removed)

- #{query}

For example, you can change the path to "/new/#{path}", the hostname to "example.#{host}",
or the query to "#{query}&value=xyz".

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Host" : String,
  "Path" : String,
  "Port" : String,
  "Protocol" : String,
  "Query" : String,
  "StatusCode" : String
}

```

### YAML

```yaml

  Host: String
  Path: String
  Port: String
  Protocol: String
  Query: String
  StatusCode: String

```

## Properties

`Host`

The hostname. This component is not percent-encoded. The hostname can contain
#{host}.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The absolute path, starting with the leading "/". This component is not percent-encoded.
The path can contain #{host}, #{path}, and #{port}.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port. You can specify a value from 1 to 65535 or #{port}.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol. You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP,
HTTP to HTTPS, and HTTPS to HTTPS. You can't redirect HTTPS to HTTP.

_Required_: No

_Type_: String

_Pattern_: `^(HTTPS?|#\{protocol\})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Query`

The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include
the leading "?", as it is automatically added. You can specify any of the reserved
keywords.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusCode`

The HTTP redirect code. The redirect is either permanent (HTTP 301) or temporary (HTTP
302).

_Required_: Yes

_Type_: String

_Allowed values_: `HTTP_301 | HTTP_302`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

This example creates a listener rule with an action that redirects HTTPS requests on
port 443. Use this rule with an HTTP listener. At least one of the following components must
be modified (to avoid a redirect loop): protocol, hostname, port, or path. Components that
are not modified retain their original values.

#### YAML

```yaml

myRedirectListenerRule:
  Type: AWS::ElasticLoadBalancingV2::ListenerRule
  Properties:
    Actions:
      - Type: redirect
        RedirectConfig:
          Protocol: HTTPS
          Port: 443
          Host: '#{host}'
          Path: /#{path}
          Query: '#{query}'
          StatusCode: HTTP_301
    Conditions:
      - Field: path-pattern
        Values:
          - /path
    ListenerArn: !Ref 'myHTTPListener'
    Priority: '10'
```

#### JSON

```json

{
    "myHTTPSlistenerRule": {
        "Type": "AWS::ElasticLoadBalancingV2::ListenerRule",
        "Properties": {
            "Actions": [
                {
                   "Type": "redirect",
                   "RedirectConfig": {
                       "Protocol": "HTTPS",
                       "Port": 443,
                       "Host": "#{host}",
                       "Path": "/#{path}",
                       "Query": "#{query}",
                       "StatusCode": "HTTP_301"
                   }
                }
            ],
            "Conditions": [
                {
                    "Field" : "path-pattern",
                    "Values" : ["/path"]
                }
            ],
            "ListenerArn": {
                "Ref": "myHTTPListener"
            },
            "Priority": "10"
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryStringKeyValue

RewriteConfig

All content copied from https://docs.aws.amazon.com/.
