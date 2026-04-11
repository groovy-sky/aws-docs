This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::FirewallRuleGroup

High-level information for a firewall rule group. A firewall
rule group is a collection of rules that DNS Firewall uses to filter DNS network traffic for a VPC. To retrieve the rules for the rule group, call
[ListFirewallRules](../../../../reference/route53/latest/apireference/api-route53resolver-listfirewallrules.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::FirewallRuleGroup",
  "Properties" : {
      "FirewallRules" : [ FirewallRule, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::FirewallRuleGroup
Properties:
  FirewallRules:
    - FirewallRule
  Name: String
  Tags:
    - Tag

```

## Properties

`FirewallRules`

A list of the rules that you have defined.

_Required_: No

_Type_: Array of [FirewallRule](aws-properties-route53resolver-firewallrulegroup-firewallrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the rule group.

_Required_: No

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of the tag keys and values that you want to associate with the rule group.

_Required_: No

_Type_: Array of [Tag](aws-properties-route53resolver-firewallrulegroup-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returnsthe `FirewallRuleGroupId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN (Amazon Resource Name) of the rule group.

`CreationTime`

The date and time that the rule group was created, in Unix time format and Coordinated Universal Time (UTC).

`CreatorRequestId`

A unique string defined by you to identify the request. This allows you to
retry failed requests without the risk of running the operation twice. This can be any
unique string, for example, a timestamp.

`Id`

The ID of the rule group.

`ModificationTime`

The date and time that the rule group was last modified, in Unix time format and Coordinated Universal Time (UTC).

`OwnerId`

The AWS account ID for the account that created the rule group. When a rule group is
shared with your account, this is the account that has shared the rule group with you.

`RuleCount`

The number of rules in the rule group.

`ShareStatus`

Whether the rule group is shared with other AWS accounts, or was shared with the current account by
another AWS account. Sharing is configured through AWS Resource Access Manager (AWS RAM).

`Status`

The status of the domain list.

`StatusMessage`

Additional information about the status of the rule group, if available.

## Examples

### Create Firewall rule group

The following example creates a DNS Firewall rule group with
associated rules for `ALLOW`, `ALERT`, and `BLOCK`.

#### JSON

```json

{"Type": "AWS::Route53Resolver::FirewallRuleGroup",
"Properties": {
     "FirewallRules": [
         {
            "Action": "ALERT",
            "FirewallDomainListId": "rslvr-fdl-sampleID1",
            "Priority": 1
         },
         {
            "Action": "BLOCK",
            "BlockResponse": "NODATA",
            "FirewallDomainListId": "rslvr-fdl-sampleID2",
            "Priority": 2
          },
          {
            "Action": "BLOCK",
            "BlockResponse": "NXDOMAIN",
            "FirewallDomainListId": "rslvr-fdl-sampleID3",
            "Priority": 3
          },
          {
            "Action": "BLOCK",
            "BlockResponse": "OVERRIDE",
            "BlockOverrideDnsType": "CNAME",
            "BlockOverrideDomain": "www.example.com",
            "BlockOverrideTtl": 300,
            "FirewallDomainListId": "rslvr-fdl-sampleID4",
            "Priority": 4
          },
          {
            "Action": "ALLOW",
            "FirewallDomainListId": "rslvr-fdl-sampleID5",
            "Priority": 5
          }
        ],
        "Name": "SampleFirewallRuleGroup",
        "Tags": [
          {
            "Key": "LineOfBusiness",
            "Value": "Engineering"
          }
      ]
   }
}
```

#### YAML

```yaml

Type: AWS::Route53Resolver::FirewallRuleGroup
Properties:
   FirewallRules:
      -
        Action: ALERT
        FirewallDomainListId: rslvr-fdl-sampleID1
        Priority: 1
      -
        Action: BLOCK
        BlockResponse: NODATA
        FirewallDomainListId: rslvr-fdl-sampleID2
        Priority: 2
      -
        Action: BLOCK
        BlockResponse: NXDOMAIN
        FirewallDomainListId: rslvr-fdl-sampleID3
        Priority: 3
      -
        Action: BLOCK
        BlockResponse: OVERRIDE
        BlockOverrideDnsType: CNAME
        BlockOverrideDomain: "www.example.com"
        BlockOverrideTtl: 300
        FirewallDomainListId: rslvr-fdl-sampleID4
        Priority: 4
      -
        Action: ALLOW
        FirewallDomainListId: rslvr-fdl-sampleID5
        Priority: 5
        Name: SampleFirewallRuleGroup
    Tags:
      -
         Key: LineOfBusiness
         Value: Engineering
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

FirewallRule

All content copied from https://docs.aws.amazon.com/.
