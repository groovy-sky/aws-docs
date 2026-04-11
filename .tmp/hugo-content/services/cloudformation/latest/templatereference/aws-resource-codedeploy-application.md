This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::Application

The `AWS::CodeDeploy::Application` resource creates an AWS CodeDeploy
application. In CodeDeploy, an application is a name that functions as a container
to ensure that the correct combination of revision, deployment configuration, and deployment
group are referenced during a deployment. You can use the
`AWS::CodeDeploy::DeploymentGroup` resource to associate the application with a
CodeDeploy deployment group. For more information, see [CodeDeploy\
Deployments](../../../codedeploy/latest/userguide/deployment-steps.md) in the _AWS CodeDeploy User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeDeploy::Application",
  "Properties" : {
      "ApplicationName" : String,
      "ComputePlatform" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeDeploy::Application
Properties:
  ApplicationName: String
  ComputePlatform: String
  Tags:
    - Tag

```

## Properties

`ApplicationName`

A name for the application. If you don't specify a name, CloudFormation generates a
unique physical ID and uses that ID for the application name. For more information, see [Name\
Type](../userguide/aws-properties-name.md).

###### Note

Updates to `ApplicationName` are not supported.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComputePlatform`

The compute platform that CodeDeploy deploys the application to.

_Required_: No

_Type_: String

_Allowed values_: `Server | Lambda | ECS | Kubernetes`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The metadata that you apply to CodeDeploy applications to help you organize and
categorize them. Each tag consists of a key and an optional value, both of which you
define.

_Required_: No

_Type_: Array of [Tag](aws-properties-codedeploy-application-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::CodeDeploy::Application` resource to
the intrinsic `Ref` function, the function returns the application name, such as
`myapplication-a123d0d1`.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

## Examples

- [Specify an application with a Lambda compute platform](#aws-resource-codedeploy-application--examples--Specify_an_application_with_a_compute_platform)

- [Specify an application with a Server compute platform](#aws-resource-codedeploy-application--examples--Specify_an_application_with_a_Server_compute_platform)

### Specify an application with a Lambda compute platform

The following example specifies a CodeDeploy application with a Lambda compute platform.

#### JSON

```json

"CodeDeployApplication": {
    "Type": "AWS::CodeDeploy::Application",
    "Properties": {
        "ComputePlatform": "Lambda"
    }
}
```

#### YAML

```yaml

CodeDeployApplication:
  Type: AWS::CodeDeploy::Application
  Properties:
    ComputePlatform: Lambda
```

### Specify an application with a Server compute platform

The following example creates a CodeDeploy application with a
`Server` compute platform.

#### JSON

```json

"CodeDeployApplication": {
    "Type": "AWS::CodeDeploy::Application",
    "Properties": {
        "ComputePlatform": "Server"
    }
}
```

#### YAML

```yaml

CodeDeployApplication:
  Type: AWS::CodeDeploy::Application
  Properties:
    ComputePlatform: Server
```

## See also

- For configuring your deployment and specifying your application revisions, see
[AWS::CodeDeploy::DeploymentConfig](../userguide/aws-resource-codedeploy-deploymentconfig.md) and [AWS::CodeDeploy::DeploymentGroup](../userguide/aws-resource-codedeploy-deploymentgroup.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS CodeDeploy

Tag

All content copied from https://docs.aws.amazon.com/.
