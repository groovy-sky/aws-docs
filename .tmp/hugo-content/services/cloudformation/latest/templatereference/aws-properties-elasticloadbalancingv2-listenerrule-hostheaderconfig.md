This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule HostHeaderConfig

Information about a host header condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RegexValues" : [ String, ... ],
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  RegexValues:
    - String
  Values:
    - String

```

## Properties

`RegexValues`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The host names. The maximum length of each string is 128 characters. The comparison is
case insensitive. The following wildcard characters are supported: \* (matches 0 or more
characters) and ? (matches exactly 1 character). You must include at least one "."
character. You can include only alphabetical characters after the final "." character.

If you specify multiple strings, the condition is satisfied if one of the strings matches
the host name.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

This example creates a listener rule with an action that forwards requests
destined for the specified domain and subdomain to the specified target group.

#### YAML

```yaml

myHostHeaderListenerRule:
   Type: 'AWS::ElasticLoadBalancingV2::ListenerRule'
   Properties:
     ListenerArn: !Ref myListener
     Priority: 10
     Conditions:
       - Field: host-header
         Values:
           - example.com
           - www.example.com
     Actions:
       - Type: forward
         TargetGroupArn: !Ref myTargetGroup
```

#### JSON

```json

{
    "myHostHeaderListenerRule": {
        "Type": "AWS::ElasticLoadBalancingV2::ListenerRule",
        "Properties": {
            "ListenerArn": {
                "Ref": "myListener"
            },
            "Priority": 10,
            "Conditions": [
                {
                    "Field": "host-header",
                    "Values": [
                        "example.com",
                        "www.example.com"
                    ]
                }
            ],
            "Actions": [
                {
                    "Type": "forward",
                    "TargetGroupArn": {
                        "Ref": "myTargetGroup"
                    }
                }
            ]
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ForwardConfig

HttpHeaderConfig

All content copied from https://docs.aws.amazon.com/.
