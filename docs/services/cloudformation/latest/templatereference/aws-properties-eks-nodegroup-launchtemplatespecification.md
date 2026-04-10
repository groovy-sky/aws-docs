This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Nodegroup LaunchTemplateSpecification

An object representing a node group launch template specification. The launch template
can't include [`SubnetId`](../../../../reference/awsec2/latest/apireference/api-createnetworkinterface.md), [`IamInstanceProfile`](../../../../reference/awsec2/latest/apireference/api-iaminstanceprofile.md), [`RequestSpotInstances`](../../../../reference/awsec2/latest/apireference/api-requestspotinstances.md), [`HibernationOptions`](../../../../reference/awsec2/latest/apireference/api-hibernationoptionsrequest.md), or [`TerminateInstances`](../../../../reference/awsec2/latest/apireference/api-terminateinstances.md), or the node group deployment or
update will fail. For more information about launch templates, see [`CreateLaunchTemplate`](../../../../reference/awsec2/latest/apireference/api-createlaunchtemplate.md) in the Amazon EC2 API Reference.
For more information about using launch templates with Amazon EKS, see [Customizing managed nodes with launch templates](../../../eks/latest/userguide/launch-templates.md) in the _Amazon EKS User Guide_.

You must specify either the launch template ID or the launch template name in the
request, but not both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Name" : String,
  "Version" : String
}

```

### YAML

```yaml

  Id: String
  Name: String
  Version: String

```

## Properties

`Id`

The ID of the launch template.

You must specify either the launch template ID or the launch template name in the
request, but not both. After node group creation, you cannot use a different ID.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the launch template.

You must specify either the launch template name or the launch template ID in the
request, but not both. After node group creation, you cannot use a different
name.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version number of the launch template to use. If no version is specified, then the
template's default version is used. You can use a different version for node group
updates.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EKS::Nodegroup

NodeRepairConfig

All content copied from https://docs.aws.amazon.com/.
