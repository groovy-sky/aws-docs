This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment

Creates a continuous deployment for a target, which is a AWS IoT Greengrass core device
or group of core devices. When you add a new core device to a group of core devices that has a
deployment, AWS IoT Greengrass deploys that group's deployment to the new device.

You can define one deployment for each target. When you create a new deployment for a
target that has an existing deployment, you replace the previous deployment. AWS IoT Greengrass applies the new deployment to the target devices.

You can only add, update, or delete up to 10 deployments at a time to a single
target.

Every deployment has a revision number that indicates how many deployment revisions you
define for a target. Use this operation to create a new revision of an existing deployment.
This operation returns the revision number of the new deployment when you create it.

For more information, see the [Create deployments](https://docs.aws.amazon.com/greengrass/v2/developerguide/create-deployments.html) in the
_AWS IoT Greengrass V2 Developer Guide_.

###### Important

Deployment resources are deleted when you delete stacks. To keep the deployments in a
stack, you must specify `"DeletionPolicy": "Retain"` on each deployment resource
in the stack template that you want to keep. For more information, see [DeletionPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).

You can only delete up to 10 deployment resources at a time. If you delete more than 10
resources, you receive an error.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GreengrassV2::Deployment",
  "Properties" : {
      "Components" : {Key: Value, ...},
      "DeploymentName" : String,
      "DeploymentPolicies" : DeploymentPolicies,
      "IotJobConfiguration" : DeploymentIoTJobConfiguration,
      "ParentTargetArn" : String,
      "Tags" : {Key: Value, ...},
      "TargetArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::GreengrassV2::Deployment
Properties:
  Components:
    Key: Value
  DeploymentName: String
  DeploymentPolicies:
    DeploymentPolicies
  IotJobConfiguration:
    DeploymentIoTJobConfiguration
  ParentTargetArn: String
  Tags:
    Key: Value
  TargetArn: String

```

## Properties

`Components`

The components to deploy. This is a dictionary, where each key is the name of a component,
and each key's value is the version and configuration to deploy for that component.

_Required_: No

_Type_: Object of [ComponentDeploymentSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrassv2-deployment-componentdeploymentspecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeploymentName`

The name of the deployment.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeploymentPolicies`

The deployment policies for the deployment. These policies define how the deployment
updates components and handles failure.

_Required_: No

_Type_: [DeploymentPolicies](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrassv2-deployment-deploymentpolicies.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IotJobConfiguration`

The job configuration for the deployment configuration. The job configuration specifies
the rollout, timeout, and stop configurations for the deployment configuration.

_Required_: No

_Type_: [DeploymentIoTJobConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrassv2-deployment-deploymentiotjobconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParentTargetArn`

The parent deployment's [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) for a subdeployment.

_Required_: No

_Type_: String

_Pattern_: `arn:[^:]*:iot:[^:]*:[0-9]+:thinggroup/.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Application-specific metadata to attach to the deployment. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use
tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2\
resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the _AWS IoT Greengrass V2 Developer Guide_.

This `Json` property type is processed as a map of key-value pairs. It uses the
following format, which is different from most `Tags` implementations in CloudFormation templates.

```json

"Tags": {
    "KeyName0": "value",
    "KeyName1": "value",
    "KeyName2": "value"
}
```

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The ARN of the target AWS IoT thing or thing group.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[^:]*:iot:[^:]*:[0-9]+:(thing|thinggroup)/.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `DeploymentId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DeploymentId`

The ID of the deployment.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LambdaVolumeMount

ComponentConfigurationUpdate
