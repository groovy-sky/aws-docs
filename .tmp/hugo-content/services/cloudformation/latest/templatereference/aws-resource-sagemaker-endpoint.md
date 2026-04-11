This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Endpoint

Use the `AWS::SageMaker::Endpoint` resource to create an endpoint using the specified configuration
in the request. Amazon SageMaker uses the endpoint to provision resources and deploy models. You create the endpoint
configuration with the [AWS::SageMaker::EndpointConfig](../userguide/aws-resource-sagemaker-endpointconfig.md) resource. For more information, see [Deploy a Model on Amazon SageMaker Hosting Services](../../../sagemaker/latest/dg/how-it-works-hosting.md) in
the _Amazon SageMaker Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Endpoint",
  "Properties" : {
      "DeploymentConfig" : DeploymentConfig,
      "EndpointConfigName" : String,
      "EndpointName" : String,
      "ExcludeRetainedVariantProperties" : [ VariantProperty, ... ],
      "RetainAllVariantProperties" : Boolean,
      "RetainDeploymentConfig" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Endpoint
Properties:
  DeploymentConfig:
    DeploymentConfig
  EndpointConfigName: String
  EndpointName: String
  ExcludeRetainedVariantProperties:
    - VariantProperty
  RetainAllVariantProperties: Boolean
  RetainDeploymentConfig: Boolean
  Tags:
    - Tag

```

## Properties

`DeploymentConfig`

The deployment configuration for an endpoint, which contains the desired deployment
strategy and rollback configurations.

_Required_: No

_Type_: [DeploymentConfig](aws-properties-sagemaker-endpoint-deploymentconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointConfigName`

The name of the [AWS::SageMaker::EndpointConfig](../userguide/aws-resource-sagemaker-endpointconfig.md) resource that specifies the configuration for the endpoint. For more
information, see [CreateEndpointConfig](../../../sagemaker/latest/dg/api-createendpointconfig.md).

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointName`

The name of the endpoint. The name must be unique within an AWS Region in your AWS account. The name is case-insensitive in `CreateEndpoint`, but the case is preserved and
must be matched in [https://docs.aws.amazon.com/sagemaker/latest/APIReference/API\_runtime\_InvokeEndpoint.html](../../../../reference/sagemaker/latest/apireference/api-runtime-invokeendpoint.md).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExcludeRetainedVariantProperties`

When you are updating endpoint resources with [RetainAllVariantProperties](../../../sagemaker/latest/dg/api-updateendpoint.md#SageMaker-UpdateEndpoint-request-RetainAllVariantProperties) whose value is set to `true`,
`ExcludeRetainedVariantProperties` specifies the list of type [VariantProperty](../userguide/aws-properties-sagemaker-endpoint-variantproperty.md) to override with the values provided by `EndpointConfig`. If you don't specify a
value for `ExcludeAllVariantProperties`, no variant properties are overridden. Don't use this property
when creating new endpoint resources or when `RetainAllVariantProperties` is set to `false`.

_Required_: No

_Type_: Array of [VariantProperty](aws-properties-sagemaker-endpoint-variantproperty.md)

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetainAllVariantProperties`

When updating endpoint resources, enables or disables the retention of variant properties, such as the instance
count or the variant weight. To retain the variant properties of an endpoint when updating it, set
`RetainAllVariantProperties` to `true`. To use the variant properties specified in a new
`EndpointConfig` call when updating an endpoint, set `RetainAllVariantProperties` to
`false`. Use this property only when updating endpoint resources, not when creating new endpoint
resources.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetainDeploymentConfig`

Specifies whether to reuse the last deployment configuration. The default value is
false (the configuration is not reused).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs to apply to this resource.

For more information, see [Resource Tag](../userguide/aws-properties-resource-tags.md) and [Using Cost\
Allocation Tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md#allocation-what) in the _AWS Billing and Cost Management User_
_Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-endpoint-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the endpoint, such as
`arn:aws:sagemaker:us-west-2:012345678901:endpoint/myendpoint`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EndpointArn`

The Amazon Resource Name (ARN) of the endpoint.

`EndpointName`

The name of the endpoint, such as `MyEndpoint`.

## Examples

### SageMaker Endpoint Example

The following example creates an endpoint configuration from a trained model, and then creates an
endpoint.

#### JSON

```json

{
  "Description": "Basic Hosting entities test.  We need models to create endpoint configs.",
  "Mappings": {
    "RegionMap": {
      "us-west-2": {
        "NullTransformer": "12345678901.dkr.ecr.us-west-2.amazonaws.com/mymodel:latest"
      },
      "us-east-2": {
        "NullTransformer": "12345678901.dkr.ecr.us-east-2.amazonaws.com/mymodel:latest"
      },
      "us-east-1": {
        "NullTransformer": "12345678901.dkr.ecr.us-east-1.amazonaws.com/mymodel:latest"
      },
      "eu-west-1": {
        "NullTransformer": "12345678901.dkr.ecr.eu-west-1.amazonaws.com/mymodel:latest"
      },
      "ap-northeast-1": {
        "NullTransformer": "12345678901.dkr.ecr.ap-northeast-1.amazonaws.com/mymodel:latest"
      },
      "ap-northeast-2": {
        "NullTransformer": "12345678901.dkr.ecr.ap-northeast-2.amazonaws.com/mymodel:latest"
      },
      "ap-southeast-2": {
        "NullTransformer": "12345678901.dkr.ecr.ap-southeast-2.amazonaws.com/mymodel:latest"
      },
      "eu-central-1": {
        "NullTransformer": "12345678901.dkr.ecr.eu-central-1.amazonaws.com/mymodel:latest"
      }
    }
  },
  "Resources": {
    "Endpoint": {
      "Type": "AWS::SageMaker::Endpoint",
      "Properties": {
        "EndpointConfigName": { "Fn::GetAtt" : ["EndpointConfig", "EndpointConfigName" ] }
      }
    },
    "EndpointConfig": {
      "Type": "AWS::SageMaker::EndpointConfig",
      "Properties": {
        "ProductionVariants": [
          {
            "InitialInstanceCount": 1,
            "InitialVariantWeight": 1,
            "InstanceType": "ml.t2.large",
            "ModelName": { "Fn::GetAtt" : ["Model", "ModelName" ] },
            "VariantName": { "Fn::GetAtt" : ["Model", "ModelName" ] }
          }
        ]
      }
    },
    "Model": {
      "Type": "AWS::SageMaker::Model",
      "Properties": {
        "PrimaryContainer": {
          "Image": { "Fn::FindInMap" : [ "AWS::Region", "NullTransformer"] }
        },
        "ExecutionRoleArn": { "Fn::GetAtt" : [ "ExecutionRole", "Arn" ] }
      }
    },
    "ExecutionRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": [
                  "sagemaker.amazonaws.com"
                ]
              },
              "Action": [
                "sts:AssumeRole"
              ]
            }
          ]
        },
        "Path": "/",
        "Policies": [
          {
            "PolicyName": "root",
            "PolicyDocument": {
              "Version": "2012-10-17",
              "Statement": [
                {
                  "Effect": "Allow",
                  "Action": "*",
                  "Resource": "*"
                }
              ]
            }
          }
        ]
      }
    }
  },
  "Outputs": {
    "EndpointId": {
      "Value": { "Ref" : "Endpoint" }
    },
    "EndpointName": {
      "Value": { "Fn::GetAtt" : [ "Endpoint", "EndpointName" ] }
    }

  }

}
```

#### YAML

```yaml

Description: "Basic Hosting entities test.  We need models to create endpoint configs."
Mappings:
  RegionMap:
    "us-west-2":
      "NullTransformer": "123456789012.dkr.ecr.us-west-2.amazonaws.com/mymodel:latest"
    "us-east-2":
      "NullTransformer": "123456789012.dkr.ecr.us-east-2.amazonaws.com/mymodel:latest"
    "us-east-1":
      "NullTransformer": "123456789012.dkr.ecr.us-east-1.amazonaws.com/mymodel:latest"
    "eu-west-1":
      "NullTransformer": "123456789012.dkr.ecr.eu-west-1.amazonaws.com/mymodel:latest"
    "ap-northeast-1":
      "NullTransformer": "123456789012.dkr.ecr.ap-northeast-1.amazonaws.com/mymodel:latest"
    "ap-northeast-2":
      "NullTransformer": "123456789012.dkr.ecr.ap-northeast-2.amazonaws.com/mymodel:latest"
    "ap-southeast-2":
      "NullTransformer": "123456789012.dkr.ecr.ap-southeast-2.amazonaws.com/mymodel:latest"
    "eu-central-1":
      "NullTransformer": "123456789012.dkr.ecr.eu-central-1.amazonaws.com/mymodel:latest"
Resources:
  Endpoint:
    Type: "AWS::SageMaker::Endpoint"
    Properties:
      EndpointConfigName:
        !GetAtt EndpointConfig.EndpointConfigName
  EndpointConfig:
    Type: "AWS::SageMaker::EndpointConfig"
    Properties:
      ProductionVariants:
        - InitialInstanceCount: 1
          InitialVariantWeight: 1.0
          InstanceType: ml.t2.large
          ModelName: !GetAtt Model.ModelName
          VariantName: !GetAtt Model.ModelName
  Model:
    Type: "AWS::SageMaker::Model"
    Properties:
      PrimaryContainer:
        Image: !FindInMap [RegionMap, !Ref "AWS::Region", "NullTransformer"]
      ExecutionRoleArn: !GetAtt ExecutionRole.Arn

  ExecutionRole:
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          -
            Effect: "Allow"
            Principal:
              Service:
                - "sagemaker.amazonaws.com"
            Action:
              - "sts:AssumeRole"
      Path: "/"
      Policies:
        -
          PolicyName: "root"
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              -
                Effect: "Allow"
                Action: "*"
                Resource: "*"
Outputs:
  EndpointId:
    Value: !Ref Endpoint
  EndpointName:
    Value: !GetAtt Endpoint.EndpointName
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserSettings

Alarm

All content copied from https://docs.aws.amazon.com/.
