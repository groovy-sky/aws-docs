This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy

Use the firewall policy to define the stateless and stateful network traffic filtering behavior for your firewall. You can use one firewall policy for multiple firewalls.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkFirewall::FirewallPolicy",
  "Properties" : {
      "Description" : String,
      "FirewallPolicy" : FirewallPolicy,
      "FirewallPolicyName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::NetworkFirewall::FirewallPolicy
Properties:
  Description: String
  FirewallPolicy:
    FirewallPolicy
  FirewallPolicyName: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the firewall policy.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallPolicy`

The traffic filtering behavior of a firewall policy, defined in a collection of stateless
and stateful rule groups and other settings.

_Required_: Yes

_Type_: [FirewallPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-firewallpolicy-firewallpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallPolicyName`

The descriptive name of the firewall policy. You can't change the name of a firewall policy after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkfirewall-firewallpolicy-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the firewall policy. For example:

`{ "Ref": "arn:aws:network-firewall:us-east-1:012345678901:firewall-policy/myFirewallPolicyName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`FirewallPolicyArn`

The Amazon Resource Name (ARN) of the firewall policy.

`FirewallPolicyId`

The unique ID of the firewall policy resource.

## Examples

### Create a firewall policy

The following shows example firewall policy specifications.

#### JSON

```json

"SampleFirewallPolicy": {
    "Type": "AWS::NetworkFirewall::FirewallPolicy",
    "Properties": {
        "FirewallPolicyName": "SampleFirewallPolicyName",
        "FirewallPolicy": {
            "StatelessDefaultActions": [
                "aws:pass"
            ],
            "StatelessFragmentDefaultActions": [
                "aws:drop"
            ],
            "StatefulRuleGroupReferences": [
                {
                    "ResourceArn": {
                        "Ref": "SampleStatefulRuleGroup"
                    }
                }
            ],
            "StatelessRuleGroupReferences": [
                {
                    "ResourceArn": {
                        "Ref": "SampleStatelessRuleGroup"
                    },
                    "Priority": 100
                }
            ]
        },
        "Description": "FirewallPolicy description goes here",
        "Tags": [
            {
                "Key": "Foo",
                "Value": "Bar"
            }
        ]
    }
}
```

#### YAML

```yaml

SampleFirewallPolicy:
  Type: 'AWS::NetworkFirewall::FirewallPolicy'
  Properties:
    FirewallPolicyName: SampleFirewallPolicyName
    FirewallPolicy:
      StatelessDefaultActions:
        - 'aws:pass'
      StatelessFragmentDefaultActions:
        - 'aws:drop'
      StatefulRuleGroupReferences:
        - ResourceArn: !Ref SampleStatefulRuleGroup1
      StatelessRuleGroupReferences:
        - ResourceArn: !Ref SampleStatelessRuleGroup
          Priority: 100
    Description: FirewallPolicy description goes here
    Tags:
      - Key: Foo
        Value: Bar

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ActionDefinition
