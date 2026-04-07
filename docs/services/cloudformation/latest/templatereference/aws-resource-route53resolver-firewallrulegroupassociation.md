This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::FirewallRuleGroupAssociation

An association between a firewall rule group and a VPC, which enables DNS filtering for
the VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::FirewallRuleGroupAssociation",
  "Properties" : {
      "FirewallRuleGroupId" : String,
      "MutationProtection" : String,
      "Name" : String,
      "Priority" : Integer,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::FirewallRuleGroupAssociation
Properties:
  FirewallRuleGroupId: String
  MutationProtection: String
  Name: String
  Priority: Integer
  Tags:
    - Tag
  VpcId: String

```

## Properties

`FirewallRuleGroupId`

The unique identifier of the firewall rule group.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MutationProtection`

If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the association.

_Required_: No

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC. DNS Firewall
filters VPC traffic starting from rule group with the lowest numeric priority setting.

You must specify a unique priority for each rule group that you associate with a single VPC.
To make it easier to insert rule groups later, leave space between the numbers, for example, use 101, 200, and so on.
You can change the priority setting for a rule group association after you create it.

The allowed values for `Priority` are between 100 and 9900 (excluding 100 and 9900).

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of the tag keys and values that you want to associate with the rule group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53resolver-firewallrulegroupassociation-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The unique identifier of the VPC that is associated with the rule group.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `FirewallRuleGroupAssociation` ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the firewall rule group association.

`CreationTime`

The date and time that the association was created, in Unix time format and Coordinated Universal Time (UTC).

`CreatorRequestId`

A unique string defined by you to identify the request. This allows you to
retry failed requests without the risk of running the operation twice. This can be any
unique string, for example, a timestamp.

`Id`

The identifier for the association.

`ManagedOwnerName`

The owner of the association, used only for associations that are not managed by
you. If you use AWS Firewall Manager to manage your firewallls from DNS Firewall, then
this reports Firewall Manager as the managed owner.

`ModificationTime`

The date and time that the association was last modified, in Unix time format and Coordinated Universal Time (UTC).

`Status`

The current status of the association.

`StatusMessage`

Additional information about the status of the response, if available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
