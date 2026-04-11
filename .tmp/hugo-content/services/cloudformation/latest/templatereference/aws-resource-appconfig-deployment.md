This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::Deployment

The `AWS::AppConfig::Deployment` resource starts a deployment. Starting a
deployment in AWS AppConfig calls the `StartDeployment` API action. This
call includes the IDs of the AWS AppConfig application, the environment, the
configuration profile, and (optionally) the configuration data version to deploy. The call
also includes the ID of the deployment strategy to use, which determines how the configuration
data is deployed.

AWS AppConfig monitors the distribution to all hosts and reports status. If a
distribution fails, then AWS AppConfig rolls back the configuration.

AWS AppConfig requires that you create resources and deploy a configuration in the
following order:

1. Create an application

2. Create an environment

3. Create a configuration profile

4. Choose a pre-defined deployment strategy or create your own

5. Deploy the configuration

For more information, see [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md) in the _AWS AppConfig User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppConfig::Deployment",
  "Properties" : {
      "ApplicationId" : String,
      "ConfigurationProfileId" : String,
      "ConfigurationVersion" : String,
      "DeploymentStrategyId" : String,
      "Description" : String,
      "DynamicExtensionParameters" : [ DynamicExtensionParameters, ... ],
      "EnvironmentId" : String,
      "KmsKeyIdentifier" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppConfig::Deployment
Properties:
  ApplicationId: String
  ConfigurationProfileId: String
  ConfigurationVersion: String
  DeploymentStrategyId: String
  Description: String
  DynamicExtensionParameters:
    - DynamicExtensionParameters
  EnvironmentId: String
  KmsKeyIdentifier: String
  Tags:
    - Tag

```

## Properties

`ApplicationId`

The application ID.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9]{4,7}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationProfileId`

The configuration profile ID.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9]{4,7}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationVersion`

The configuration version to deploy. If deploying an AWS AppConfig hosted
configuration version, you can specify either the version number or version label. For all
other configurations, you must specify the version number.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeploymentStrategyId`

The deployment strategy ID.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9]{4,7}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the deployment.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DynamicExtensionParameters`

A map of dynamic extension parameter names to values to pass to associated extensions
with `PRE_START_DEPLOYMENT` actions.

_Required_: No

_Type_: [Array](aws-properties-appconfig-deployment-dynamicextensionparameters.md) of [DynamicExtensionParameters](aws-properties-appconfig-deployment-dynamicextensionparameters.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentId`

The environment ID.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9]{4,7}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyIdentifier`

The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when
the resource was created or updated.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata to assign to the deployment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which
you define.

_Required_: No

_Type_: Array of [Tag](aws-properties-appconfig-deployment-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`DeploymentNumber`

The sequence number of the deployment.

`State`

The state of the deployment.

## Examples

### AWS AppConfig deployment example

The following example creates an AWS AppConfig deployment. Starting a
deployment in AWS AppConfig calls the StartDeployment API action. This call
includes the IDs of the AWS AppConfig application, the environment, the
configuration profile, and (optionally) the configuration data version to deploy. The call
also includes the ID of the deployment strategy to use, which determines how the
configuration data is deployed.

AWS AppConfig monitors the distribution to all hosts and reports status. If a
distribution fails, then AWS AppConfig rolls back the configuration.

#### JSON

```json

Resources": {
    "BasicDeployment": {
      "Type": "AWS::AppConfig::Deployment",
      "DependsOn": [
        "MyTestApplication",
        "MyTestConfigurationProfile",
        "MyTestEnvironment",
        "MyTestDeploymentStrategy"
      ],
      "Properties": {
        "ApplicationId": MyTestApplication,
        "EnvironmentId": MyTestEnvironment,
        "DeploymentStrategyId": MyTestDeploymentStrategy,
        "ConfigurationProfileId": MyTestConfigurationProfile,
        "ConfigurationVersion": "1",
        "Description": "My test deployment",
        "Tags": [
          {
            "Key": "Env",
            "Value": "test"
          }
        ]
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  BasicDeployment:
    Type: AWS::AppConfig::Deployment
    Properties:
      ApplicationId: !Ref MyTestApplication
      EnvironmentId: !Ref MyTestEnvironment
      DeploymentStrategyId: !Ref MyTestDeploymentStrategy
      ConfigurationProfileId: !Ref MyTestConfigurationProfile
      ConfigurationVersion: '1'
      Description: 'My test deployment'
      Tags:
        - Key: Env
          Value: test

```

## See also

- [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md)

- [Deploying a\
configuration](../../../systems-manager/latest/userguide/appconfig-deploying.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Validators

DynamicExtensionParameters

All content copied from https://docs.aws.amazon.com/.
