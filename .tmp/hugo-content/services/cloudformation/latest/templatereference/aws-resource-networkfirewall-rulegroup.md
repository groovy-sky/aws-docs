This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup

Use the [https://docs.aws.amazon.com/RuleGroup](../../../rulegroup/index.md) to define a reusable collection of stateless or stateful network traffic filtering rules.
You use rule groups in an firewall policy to specify the filtering behavior of an firewall.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkFirewall::RuleGroup",
  "Properties" : {
      "Capacity" : Integer,
      "Description" : String,
      "RuleGroup" : RuleGroup,
      "RuleGroupName" : String,
      "SummaryConfiguration" : SummaryConfiguration,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkFirewall::RuleGroup
Properties:
  Capacity: Integer
  Description: String
  RuleGroup:
    RuleGroup
  RuleGroupName: String
  SummaryConfiguration:
    SummaryConfiguration
  Tags:
    - Tag
  Type: String

```

## Properties

`Capacity`

The maximum operating resources that this rule group can use. You can't change a rule group's capacity setting
after you create the rule group. When you update a rule group, you are limited to this capacity. When you reference a rule group
from a firewall policy, Network Firewall reserves this capacity for the rule group.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the rule group.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleGroup`

An object that defines the rule group rules.

_Required_: No

_Type_: [RuleGroup](aws-properties-networkfirewall-rulegroup-rulegroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleGroupName`

The descriptive name of the rule group. You can't change the name of a rule group after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SummaryConfiguration`

A complex type containing the currently selected rule option fields that will be displayed for rule summarization returned by `DescribeRuleGroupSummary`.

- The `RuleOptions` specified in `SummaryConfiguration`

- Rule metadata organization preferences

_Required_: No

_Type_: [SummaryConfiguration](aws-properties-networkfirewall-rulegroup-summaryconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-networkfirewall-rulegroup-tag.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Indicates whether the rule group is stateless or stateful. If the rule group is stateless, it contains
stateless rules. If it is stateful, it contains stateful rules.

_Required_: Yes

_Type_: String

_Allowed values_: `STATELESS | STATEFUL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the rule group. For example:

`{ "Ref": "arn:aws:network-firewall:us-east-1:012345678901:stateful-rulegroup/myStatefulRuleGroupName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RuleGroupArn`

The Amazon Resource Name (ARN) of the rule group.

`RuleGroupId`

The unique ID of the rule group resource.

## Examples

- [Create a stateful rule group](#aws-resource-networkfirewall-rulegroup--examples--Create_a_stateful_rule_group)

- [Create a stateless rule group](#aws-resource-networkfirewall-rulegroup--examples--Create_a_stateless_rule_group)

### Create a stateful rule group

The following shows example stateful rule group specifications.

#### JSON

```json

"SampleStatefulRulegroup": {
    "Type": "AWS::NetworkFirewall::RuleGroup",
    "Properties": {
        "RuleGroupName": "SampleStatefulRulegroupName",
        "Type": "STATEFUL",
        "RuleGroup": {
            "RulesSource": {
                "RulesString": "pass tcp 10.20.20.0/24 45400:45500 <> 10.10.10.0/24 5203 (msg:\"test\";sid:1;rev:1;)"
            }
        },
        "Capacity": 100,
        "Description": "Rulegroup description goes here",
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

SampleStatefulRulegroup:
  Type: 'AWS::NetworkFirewall::RuleGroup'
  Properties:
    RuleGroupName: SampleStatefulRulegroupName
    Type: STATEFUL
    RuleGroup:
      RulesSource:
        RulesString: >-
          pass tcp 10.20.20.0/24 45400:45500 <> 10.10.10.0/24 5203
          (msg:"test";sid:1;rev:1;)
    Capacity: 100
    Description: Rulegroup description goes here
    Tags:
      - Key: Foo
        Value: Bar

```

### Create a stateless rule group

The following shows example stateless rule group specifications.

#### JSON

```json

"SampleStatelessRulegroup": {
    "Type": "AWS::NetworkFirewall::RuleGroup",
    "Properties": {
        "RuleGroupName": "SampleStatelessRulegroupName",
        "Type": "STATELESS",
        "RuleGroup": {
            "RulesSource": {
                "StatelessRulesAndCustomActions": {
                    "StatelessRules": [
                        {
                            "RuleDefinition": {
                                "MatchAttributes": {
                                    "Sources": [
                                        {
                                            "AddressDefinition": "0.0.0.0/0"
                                        }
                                    ],
                                    "Destinations": [
                                        {
                                            "AddressDefinition": "10.0.0.0/8"
                                        }
                                    ],
                                    "SourcePorts": [
                                        {
                                            "FromPort": 15000
                                        },
                                        {
                                            "ToPort": 30000
                                        }
                                    ],
                                    "DestinationPorts": [
                                        {
                                            "FromPort": 443
                                        },
                                        {
                                            "ToPort": 443
                                        }
                                    ],
                                    "Protocols": [
                                        6
                                    ]
                                },
                                "Actions": [
                                    "aws:pass"
                                ]
                            },
                            "Priority": 1
                        }
                    ]
                }
            }
        },
        "Capacity": 100,
        "Description": "Rulegroup description goes here",
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

SampleStatelessRulegroup:
              Type: 'AWS::NetworkFirewall::RuleGroup'
              Properties:
                RuleGroupName: SampleStatelessRulegroupName
                Type: STATELESS
                RuleGroup:
                  RulesSource:
                    StatelessRulesAndCustomActions:
                      StatelessRules:
                        - RuleDefinition:
                            MatchAttributes:
                              Sources:
                                - AddressDefinition: 0.0.0.0/0
                              Destinations:
                                - AddressDefinition: 10.0.0.0/8
                              SourcePorts:
                                - FromPort: 15000
                                  ToPort: 30000
                              DestinationPorts:
                                - FromPort: 443
                                  ToPort: 443
                              Protocols:
                                - 6
                            Actions:
                              - 'aws:pass'
                          Priority: 1
                Capacity: 100
                Description: Rulegroup description goes here
                Tags:
                  - Key: Foo
                    Value: Bar
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingConfiguration

ActionDefinition

All content copied from https://docs.aws.amazon.com/.
