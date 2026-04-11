This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMPrefixListResolver IpamPrefixListResolverRuleCondition

Describes a condition within a CIDR selection rule. Conditions define the criteria for selecting CIDRs from IPAM's database based on resource attributes.

CIDR selection rules define the business logic for selecting CIDRs from IPAM. If a CIDR matches any of the rules, it will be included. If a rule has multiple conditions, the CIDR has to match every condition of that rule. You can create a prefix list resolver without any CIDR selection rules, but it will generate empty versions (containing no CIDRs) until you add rules.

There are three rule types. Only 2 of the 3 rule types support conditions - **IPAM pool CIDR** and **Scope resource CIDR**. **Static CIDR** rules cannot have conditions.

- **Static CIDR**: A fixed list of CIDRs that do not change (like a manual list replicated across Regions)

- **IPAM pool CIDR**: CIDRs from specific IPAM pools (like all CIDRs from your IPAM production pool)

If you choose this option, choose the following:

- **IPAM scope**: Select the IPAM scope to search for resources

- **Conditions:**

- **Property**

- **IPAM pool ID**: Select an IPAM pool that contains the resources

- **CIDR** (like 10.24.34.0/23)

- **Operation**: Equals/Not equals

- **Value**: The value on which to match the condition

- **Scope resource CIDR**: CIDRs from AWS resources like VPCs, subnets, EIPs within an IPAM scope

If you choose this option, choose the following:

- **IPAM scope**: Select the IPAM scope to search for resources

- **Resource type**: Select a resource, like a VPC or subnet.

- **Conditions**:

- **Property**:

- Resource ID: The unique ID of a resource (like vpc-1234567890abcdef0)

- Resource owner (like 111122223333)

- Resource region (like us-east-1)

- Resource tag (like key: name, value: dev-vpc-1)

- CIDR (like 10.24.34.0/23)

- **Operation**: Equals/Not equals

- **Value**: The value on which to match the condition

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidr" : String,
  "IpamPoolId" : String,
  "Operation" : String,
  "ResourceId" : String,
  "ResourceOwner" : String,
  "ResourceRegion" : String,
  "ResourceTag" : Tag
}

```

### YAML

```yaml

  Cidr: String
  IpamPoolId: String
  Operation: String
  ResourceId: String
  ResourceOwner: String
  ResourceRegion: String
  ResourceTag:
    Tag

```

## Properties

`Cidr`

A CIDR block to match against. This condition selects CIDRs that fall within or match the specified CIDR range.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamPoolId`

The ID of the IPAM pool to match against. This condition selects CIDRs that belong to the specified IPAM pool.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operation`

The operation to perform when evaluating this condition. Valid values include `equals`, `not-equals`, and `subnet-of`.

_Required_: No

_Type_: String

_Allowed values_: `equals | not-equals | subnet-of`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

The ID of the AWS resource to match against. This condition selects CIDRs associated with the specified resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceOwner`

The AWS account ID that owns the resources to match against. This condition selects CIDRs from resources owned by the specified account.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceRegion`

The AWS Region where the resources are located. This condition selects CIDRs from resources in the specified Region.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTag`

A tag key-value pair to match against. This condition selects CIDRs from resources that have the specified tag.

_Required_: No

_Type_: [Tag](aws-properties-ec2-ipamprefixlistresolver-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IpamPrefixListResolverRule

Tag

All content copied from https://docs.aws.amazon.com/.
