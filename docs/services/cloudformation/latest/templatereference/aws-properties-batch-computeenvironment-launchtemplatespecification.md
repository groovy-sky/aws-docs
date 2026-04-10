This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ComputeEnvironment LaunchTemplateSpecification

An object that represents a launch template that's associated with a compute resource. You
must specify either the launch template ID or launch template name in the request, but not
both.

If security groups are specified using both the `securityGroupIds` parameter of
`CreateComputeEnvironment` and the launch template, the values in the
`securityGroupIds` parameter of `CreateComputeEnvironment` will be
used.

###### Note

This object isn't applicable to jobs that are running on Fargate resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateId" : String,
  "LaunchTemplateName" : String,
  "Overrides" : [ LaunchTemplateSpecificationOverride, ... ],
  "UserdataType" : String,
  "Version" : String
}

```

### YAML

```yaml

  LaunchTemplateId: String
  LaunchTemplateName: String
  Overrides:
    - LaunchTemplateSpecificationOverride
  UserdataType: String
  Version: String

```

## Properties

`LaunchTemplateId`

The ID of the launch template.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LaunchTemplateName`

The name of the launch template.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Overrides`

A launch template to use in place of the default launch template. You must specify either the launch template ID or launch template name in the request, but not both.

You can specify up to ten (10) launch template overrides that are associated to unique instance types or families for each compute environment.

###### Note

To unset all override templates for a compute environment, you can pass an empty array to the [UpdateComputeEnvironment.overrides](../../../../reference/batch/latest/apireference/api-updatecomputeenvironment.md) parameter, or not include the `overrides` parameter when submitting the `UpdateComputeEnvironment` API operation.

_Required_: No

_Type_: Array of [LaunchTemplateSpecificationOverride](aws-properties-batch-computeenvironment-launchtemplatespecificationoverride.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UserdataType`

The EKS node initialization process to use. You only need to specify this value if you are
using a custom AMI. The default value is `EKS_BOOTSTRAP_SH`. If
_imageType_ is a custom AMI based on EKS\_AL2023 or EKS\_AL2023\_NVIDIA then you
must choose `EKS_NODEADM`.

_Required_: No

_Type_: String

_Allowed values_: `EKS_BOOTSTRAP_SH | EKS_NODEADM`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Version`

The version number of the launch template,
`$Default`, or `$Latest`.

If the value is `$Default`, the default version of the launch template is used. If the value is `$Latest`, the latest version of the launch template is used.

###### Important

If the AMI ID that's used in a compute environment is from the launch template, the AMI
isn't changed when the compute environment is updated. It's only changed if the
`updateToLatestImageVersion` parameter for the compute environment is set to
`true`. During an infrastructure update, if either `$Default` or
`$Latest` is specified, AWS Batch re-evaluates the launch template version, and it
might use a different version of the launch template. This is the case even if the launch
template isn't specified in the update. When updating a compute environment, changing the launch
template requires an infrastructure update of the compute environment. For more information, see
[Updating compute\
environments](../../../batch/latest/userguide/updating-compute-environments.md) in the _AWS Batch User Guide_.

Default: `$Default`

Latest: `$Latest`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## See also

- [Launch Template Support](../../../batch/latest/userguide/launch-templates.md) in the _AWS Batch User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksConfiguration

LaunchTemplateSpecificationOverride

All content copied from https://docs.aws.amazon.com/.
