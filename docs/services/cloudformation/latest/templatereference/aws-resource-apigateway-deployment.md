This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Deployment

The `AWS::ApiGateway::Deployment` resource deploys an API Gateway `RestApi` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::Deployment",
  "Properties" : {
      "DeploymentCanarySettings" : DeploymentCanarySettings,
      "Description" : String,
      "RestApiId" : String,
      "StageDescription" : StageDescription,
      "StageName" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::Deployment
Properties:
  DeploymentCanarySettings:
    DeploymentCanarySettings
  Description: String
  RestApiId: String
  StageDescription:
    StageDescription
  StageName: String

```

## Properties

`DeploymentCanarySettings`

The input configuration for a canary deployment.

_Required_: No

_Type_: [DeploymentCanarySettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-deployment-deploymentcanarysettings.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description for the Deployment resource to create.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StageDescription`

The description of the Stage resource for the Deployment resource to create. To specify a stage description, you must also provide a stage name.

_Required_: Conditional

_Type_: [StageDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-deployment-stagedescription.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StageName`

The name of the Stage resource for the Deployment resource to create.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the deployment ID, such as `123abc`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DeploymentId`

The ID for the deployment. For example: `abc123`.

## Examples

The following sections provide examples for declaring API Gateway deployments.

- [Deployment with an empty embedded stage](#aws-resource-apigateway-deployment--examples--Deployment_with_an_empty_embedded_stage)

- [AWS::ApiGateway::Method Dependency](#aws-resource-apigateway-deployment--examples--AWS::ApiGateway::Method_Dependency)

### Deployment with an empty embedded stage

The following example deploys the `MyApi` API to a stage named `DummyStage`.

#### JSON

```json

{
    "Deployment": {
        "Type": "AWS::ApiGateway::Deployment",
        "Properties": {
            "RestApiId": {
                "Ref": "MyApi"
            },
            "Description": "My deployment",
            "StageName": "DummyStage"
        }
    }
}
```

#### YAML

```yaml

Deployment:
  Type: 'AWS::ApiGateway::Deployment'
  Properties:
    RestApiId: !Ref MyApi
    Description: My deployment
    StageName: DummyStage

```

### AWS::ApiGateway::Method Dependency

If you create an [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource and its methods (using [AWS::ApiGateway::Method](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html)) in the same template as your deployment, the deployment must depend on the `RestApi`'s methods. To create a dependency, add a `DependsOn` attribute to the deployment. If you don't, CloudFormation creates the deployment right after it creates the `RestApi` resource that doesn't contain any methods, and CloudFormation encounters the following error: `The REST API doesn't contain any methods`.

#### JSON

```json

{
    "Deployment": {
        "DependsOn": "MyMethod",
        "Type": "AWS::ApiGateway::Deployment",
        "Properties": {
            "RestApiId": {
                "Ref": "MyApi"
            },
            "Description": "My deployment",
            "StageName": "DummyStage"
        }
    }
}
```

#### YAML

```yaml

Deployment:
  DependsOn: MyMethod
  Type: 'AWS::ApiGateway::Deployment'
  Properties:
    RestApiId: !Ref MyApi
    Description: My deployment
    StageName: DummyStage

```

## See also

- [deployment:create](../../../apigateway/latest/api/api-createdeployment.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AccessLogSetting
