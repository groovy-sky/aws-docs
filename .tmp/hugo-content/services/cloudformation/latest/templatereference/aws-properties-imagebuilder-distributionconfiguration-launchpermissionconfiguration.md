This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration LaunchPermissionConfiguration

Describes the configuration for a launch permission. The launch permission
modification request is sent to the [Amazon EC2\
ModifyImageAttribute](../../../../reference/awsec2/latest/apireference/api-modifyimageattribute.md) API on behalf of the user for each Region they have
selected to distribute the AMI. To make an AMI public, set the launch permission
authorized accounts to `all`. See the examples for making an AMI public at
[Amazon EC2\
ModifyImageAttribute](../../../../reference/awsec2/latest/apireference/api-modifyimageattribute.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OrganizationalUnitArns" : [ String, ... ],
  "OrganizationArns" : [ String, ... ],
  "UserGroups" : [ String, ... ],
  "UserIds" : [ String, ... ]
}

```

### YAML

```yaml

  OrganizationalUnitArns:
    - String
  OrganizationArns:
    - String
  UserGroups:
    - String
  UserIds:
    - String

```

## Properties

`OrganizationalUnitArns`

The ARN for an AWS Organizations organizational unit (OU) that you want to share your AMI with.
For more information about key concepts for AWS Organizations, see [AWS Organizations\
terminology and concepts](../../../organizations/latest/userguide/orgs-getting-started-concepts.md).

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationArns`

The ARN for an AWS Organization that you want to share your AMI with. For more
information, see [What is\
AWS Organizations?](../../../organizations/latest/userguide/orgs-introduction.md).

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserGroups`

The name of the group.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserIds`

The AWS account ID.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FastLaunchSnapshotConfiguration

LaunchTemplateConfiguration

All content copied from https://docs.aws.amazon.com/.
