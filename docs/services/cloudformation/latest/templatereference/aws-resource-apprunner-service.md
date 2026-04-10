This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service

The `AWS::AppRunner::Service` resource is an AWS App Runner resource type that specifies an App Runner service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppRunner::Service",
  "Properties" : {
      "AutoScalingConfigurationArn" : String,
      "EncryptionConfiguration" : EncryptionConfiguration,
      "HealthCheckConfiguration" : HealthCheckConfiguration,
      "InstanceConfiguration" : InstanceConfiguration,
      "NetworkConfiguration" : NetworkConfiguration,
      "ObservabilityConfiguration" : ServiceObservabilityConfiguration,
      "ServiceName" : String,
      "SourceConfiguration" : SourceConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppRunner::Service
Properties:
  AutoScalingConfigurationArn: String
  EncryptionConfiguration:
    EncryptionConfiguration
  HealthCheckConfiguration:
    HealthCheckConfiguration
  InstanceConfiguration:
    InstanceConfiguration
  NetworkConfiguration:
    NetworkConfiguration
  ObservabilityConfiguration:
    ServiceObservabilityConfiguration
  ServiceName: String
  SourceConfiguration:
    SourceConfiguration
  Tags:
    - Tag

```

## Properties

`AutoScalingConfigurationArn`

The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service. If not provided, App Runner
associates the latest revision of a default auto scaling configuration.

Specify an ARN with a name and a revision number to associate that revision. For example:
`arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/3`

Specify just the name to associate the latest revision. For example:
`arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability`

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs. By default,
App Runner uses an AWS managed key.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-apprunner-service-encryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HealthCheckConfiguration`

The settings for the health check that AWS App Runner performs to monitor the health of the App Runner service.

_Required_: No

_Type_: [HealthCheckConfiguration](aws-properties-apprunner-service-healthcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceConfiguration`

The runtime configuration of instances (scaling units) of your service.

_Required_: No

_Type_: [InstanceConfiguration](aws-properties-apprunner-service-instanceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

Configuration settings related to network traffic of the web application that the App Runner service runs.

_Required_: No

_Type_: [NetworkConfiguration](aws-properties-apprunner-service-networkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObservabilityConfiguration`

The observability configuration of your service.

_Required_: No

_Type_: [ServiceObservabilityConfiguration](aws-properties-apprunner-service-serviceobservabilityconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

A name for the App Runner service. It must be unique across all the running App Runner services in your AWS account in the AWS Region.

If you don't specify a name, CloudFormation generates a name for your service.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9][A-Za-z0-9-_]{3,39}`

_Minimum_: `4`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceConfiguration`

The source to deploy to the App Runner service. It can be a code or an image repository.

_Required_: Yes

_Type_: [SourceConfiguration](aws-properties-apprunner-service-sourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An optional list of metadata items that you can associate with the App Runner service resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](aws-properties-apprunner-service-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the ARN of the App Runner service.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ServiceArn`

The Amazon Resource Name (ARN) of this service.

`ServiceId`

An ID that App Runner generated for this service. It's unique within the AWS Region.

`ServiceUrl`

A subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.

`Status`

The current state of the App Runner service. These particular values mean the following.

- `CREATE_FAILED` – The service failed to create. The failed service isn't usable, and still counts towards your service quota. To
troubleshoot this failure, read the failure events and logs, change any parameters that need to be fixed, and rebuild your service using
`UpdateService`.

- `DELETE_FAILED` – The service failed to delete and can't be successfully recovered. Retry the service deletion call to ensure
that all related resources are removed.

## Examples

- [Service based on source code](#aws-resource-apprunner-service--examples--Service_based_on_source_code)

- [Service based on source image](#aws-resource-apprunner-service--examples--Service_based_on_source_image)

### Service based on source code

This example illustrates creating a service based on a Python source code repository.

#### JSON

```json

{
  "Type": "AWS::AppRunner::Service",
  "Properties": {
    "ServiceName": "python-app",
    "SourceConfiguration": {
      "AuthenticationConfiguration": {
        "ConnectionArn": "arn:aws:apprunner:us-east-1:123456789012:connection/my-github-connection/e7656250f67242d7819feade6800f59e"
      },
      "AutoDeploymentsEnabled": true,
      "CodeRepository": {
        "RepositoryUrl": "https://github.com/my-account/python-hello",
        "SourceCodeVersion": {
          "Type": "BRANCH",
          "Value": "main"
        },
        "CodeConfiguration": {
          "ConfigurationSource": "API",
          "CodeConfigurationValues": {
            "Runtime": "PYTHON_3",
            "BuildCommand": "pip install -r requirements.txt",
            "StartCommand": "python server.py",
            "Port": "8080",
            "RuntimeEnvironmentVariables": [
              {
                "Name": "NAME",
                "Value": "Jane"
              }
            ]
          }
        }
      }
    },
    "InstanceConfiguration": {
      "Cpu": "1 vCPU",
      "Memory": "3 GB"
    }
  }
}
```

#### YAML

```yaml

Type: AWS::AppRunner::Service
Properties:
  ServiceName: python-app
  SourceConfiguration:
    AuthenticationConfiguration:
      ConnectionArn: "arn:aws:apprunner:us-east-1:123456789012:connection/my-github-connection/e7656250f67242d7819feade6800f59e"
    AutoDeploymentsEnabled: true
    CodeRepository:
      RepositoryUrl: "https://github.com/my-account/python-hello"
      SourceCodeVersion:
        Type: BRANCH
        Value: main
      CodeConfiguration:
        ConfigurationSource: API
        CodeConfigurationValues:
          Runtime: PYTHON_3
          BuildCommand: "pip install -r requirements.txt"
          StartCommand: "python server.py"
          Port: 8080
          RuntimeEnvironmentVariables:
            -
              Name: NAME
              Value: Jane
  InstanceConfiguration:
    Cpu: 1 vCPU
    Memory: 3 GB
```

### Service based on source image

This example illustrates creating a service based on an image stored in Amazon Elastic Container Registry (Amazon ECR).

#### JSON

```json

{
  "Type": "AWS::AppRunner::Service",
  "Properties": {
    "ServiceName": "golang-container-app",
    "SourceConfiguration": {
      "AuthenticationConfiguration": {
        "AccessRoleArn": "arn:aws:iam::123456789012:role/my-ecr-role"
      },
      "AutoDeploymentsEnabled": true,
      "ImageRepository": {
        "ImageIdentifier": "123456789012.dkr.ecr.us-east-1.amazonaws.com/golang-app:latest",
        "ImageRepositoryType": "ECR",
        "ImageConfiguration": {
          "Port": "8080",
          "RuntimeEnvironmentVariables": [
            {
              "Name": "NAME",
              "Value": "Jane"
            }
          ]
        }
      }
    },
    "InstanceConfiguration": {
      "Cpu": "1 vCPU",
      "Memory": "3 GB"
    }
  }
}
```

#### YAML

```yaml

Type: AWS::AppRunner::Service
Properties:
  ServiceName: golang-container-app
  SourceConfiguration:
    AuthenticationConfiguration:
      AccessRoleArn: "arn:aws:iam::123456789012:role/my-ecr-role"
    AutoDeploymentsEnabled: true
    ImageRepository:
      ImageIdentifier: "123456789012.dkr.ecr.us-east-1.amazonaws.com/golang-app:latest"
      ImageRepositoryType: ECR
      ImageConfiguration:
        Port: 8080
        RuntimeEnvironmentVariables:
          -
            Name: NAME
            Value: Jane
  InstanceConfiguration:
    Cpu: 1 vCPU
    Memory: 3 GB
```

## See also

- [Creating an App Runner service](../../../apprunner/latest/dg/manage-create.md) in the _AWS App Runner_
_Developer Guide_

- [Deploying a new application version to App Runner](../../../apprunner/latest/dg/manage-deploy.md) in the
_AWS App Runner Developer Guide_

- [Configuring an App Runner service](../../../apprunner/latest/dg/manage-configure.md) in the
_AWS App Runner Developer Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TraceConfiguration

AuthenticationConfiguration

All content copied from https://docs.aws.amazon.com/.
