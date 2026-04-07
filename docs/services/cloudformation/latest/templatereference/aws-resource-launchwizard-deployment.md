This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LaunchWizard::Deployment

Creates a deployment for the given workload. Deployments created by this operation are
not available in the Launch Wizard console to use the `Clone deployment` action
on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LaunchWizard::Deployment",
  "Properties" : {
      "DeploymentPatternName" : String,
      "Name" : String,
      "Specifications" : {Key: Value, ...},
      "Tags" : [ Tags, ... ],
      "WorkloadName" : String
    }
}

```

### YAML

```yaml

Type: AWS::LaunchWizard::Deployment
Properties:
  DeploymentPatternName: String
  Name: String
  Specifications:
    Key: Value
  Tags:
    - Tags
  WorkloadName: String

```

## Properties

`DeploymentPatternName`

The name of the deployment pattern.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][a-zA-Z0-9-]*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the deployment.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9_\s\.-]+$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Specifications`

The settings specified for the deployment. These settings define how to deploy and configure your
resources created by the deployment. For more information about the specifications
required for creating a deployment for a SAP workload, see [SAP deployment\
specifications](https://docs.aws.amazon.com/launchwizard/latest/APIReference/launch-wizard-specifications-sap.html). To retrieve the specifications required to create a deployment for other workloads,
use the [`GetWorkloadDeploymentPattern`](https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_GetWorkloadDeploymentPattern.html) operation.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9-:]{3,256}$`

_Minimum_: `1`

_Maximum_: `1500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Information about the tags attached to a deployment.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-launchwizard-deployment-tags.html) of [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-launchwizard-deployment-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkloadName`

The name of the workload.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][a-zA-Z0-9-_]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the deployment. For
example:

`{ "Ref": "myLaunchWizardDeployment" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the deployment.

`CreatedAt`

The time the deployment was created.

`DeletedAt`

The time the deployment was deleted.

`DeploymentId`

The ID of the deployment.

`ResourceGroup`

The resource group of the deployment.

`Status`

The status of the deployment.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Launch Wizard

Tags
