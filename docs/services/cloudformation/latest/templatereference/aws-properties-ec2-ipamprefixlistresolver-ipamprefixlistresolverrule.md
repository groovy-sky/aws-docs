This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMPrefixListResolver IpamPrefixListResolverRule

Describes a CIDR selection rule.

CIDR selection rules define the business logic for selecting CIDRs from IPAM. If a CIDR matches any of the rules, it will be included. If a rule has multiple conditions, the CIDR has to match every condition of that rule. You can create a prefix list resolver without any CIDR selection rules, but it will generate empty versions (containing no CIDRs) until you add rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Conditions" : [ IpamPrefixListResolverRuleCondition, ... ],
  "IpamScopeId" : String,
  "ResourceType" : String,
  "RuleType" : String,
  "StaticCidr" : String
}

```

### YAML

```yaml

  Conditions:
    - IpamPrefixListResolverRuleCondition
  IpamScopeId: String
  ResourceType: String
  RuleType: String
  StaticCidr: String

```

## Properties

`Conditions`

The conditions that determine which CIDRs are selected by this rule. Conditions specify criteria such as resource type, tags, account IDs, and Regions.

_Required_: No

_Type_: Array of [IpamPrefixListResolverRuleCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamScopeId`

The ID of the IPAM scope from which to select CIDRs. This determines whether to select from public or private IP address space.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

For rules of type `ipam-resource-cidr`, this is the resource type.

_Required_: No

_Type_: String

_Allowed values_: `vpc | subnet | eip | public-ipv4-pool`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleType`

The type of CIDR selection rule. Valid values include `static-cidr` for specifying a fixed CIDR, `ipam-resource-cidr` for CIDRs from resources such as VPCs and subnets, and `ipam-pool-cidr` for CIDRs from IPAM pools.

_Required_: Yes

_Type_: String

_Allowed values_: `static-cidr | ipam-resource-cidr | ipam-pool-cidr`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticCidr`

A fixed CIDR that does not change. This property can only be specified for rules of type `static-cidr`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::IPAMPrefixListResolver

IpamPrefixListResolverRuleCondition
