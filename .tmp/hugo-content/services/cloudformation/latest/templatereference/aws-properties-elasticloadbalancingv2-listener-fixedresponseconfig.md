This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener FixedResponseConfig

Specifies information required when returning a custom HTTP response.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentType" : String,
  "MessageBody" : String,
  "StatusCode" : String
}

```

### YAML

```yaml

  ContentType: String
  MessageBody: String
  StatusCode: String

```

## Properties

`ContentType`

The content type.

Valid Values: text/plain \| text/css \| text/html \| application/javascript \|
application/json

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageBody`

The message.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusCode`

The HTTP response code (2XX, 4XX, or 5XX).

_Required_: Yes

_Type_: String

_Pattern_: `^(2|4|5)\d\d$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example defines a listener with a default action that returns the
specified HTTP response.

#### YAML

```yaml

myHTTPListener:
  Type: 'AWS::ElasticLoadBalancingV2::Listener'
  Properties:
    LoadBalancerArn: !Ref myLoadBalancer
    Protocol: HTTP
    Port: 80
    DefaultActions:
      - Type: fixed-response
        FixedResponseConfig:
          ContentType: text/plain
          MessageBody: "Hello world"
          StatusCode: 200
```

#### JSON

```json

{
    "myHTTPListener": {
        "Type": "AWS::ElasticLoadBalancingV2::Listener",
        "Properties": {
            "LoadBalancerArn": {
                "Ref": "myLoadBalancer"
            },
            "Protocol": "HTTP",
            "Port": 80,
            "DefaultActions": [
                {
                    "Type": "fixed-response",
                    "FixedResponseConfig": {
                        "ContentType": "text/plain",
                        "MessageBody": "Hello world",
                        "StatusCode": 200
                    }
                }
            ]
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Certificate

ForwardConfig

All content copied from https://docs.aws.amazon.com/.
