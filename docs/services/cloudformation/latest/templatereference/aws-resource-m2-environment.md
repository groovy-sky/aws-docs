This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::M2::Environment

Specifies a runtime environment for a given runtime engine.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::M2::Environment",
  "Properties" : {
      "Description" : String,
      "EngineType" : String,
      "EngineVersion" : String,
      "HighAvailabilityConfig" : HighAvailabilityConfig,
      "InstanceType" : String,
      "KmsKeyId" : String,
      "Name" : String,
      "NetworkType" : String,
      "PreferredMaintenanceWindow" : String,
      "PubliclyAccessible" : Boolean,
      "SecurityGroupIds" : [ String, ... ],
      "StorageConfigurations" : [ StorageConfiguration, ... ],
      "SubnetIds" : [ String, ... ],
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::M2::Environment
Properties:
  Description: String
  EngineType: String
  EngineVersion: String
  HighAvailabilityConfig:
    HighAvailabilityConfig
  InstanceType: String
  KmsKeyId: String
  Name: String
  NetworkType: String
  PreferredMaintenanceWindow: String
  PubliclyAccessible: Boolean
  SecurityGroupIds:
    - String
  StorageConfigurations:
    - StorageConfiguration
  SubnetIds:
    - String
  Tags:
    Key: Value

```

## Properties

`Description`

The description of the runtime environment.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineType`

The target platform for the runtime environment.

_Required_: Yes

_Type_: String

_Allowed values_: `microfocus | bluage`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineVersion`

The version of the runtime engine.

_Required_: No

_Type_: String

_Pattern_: `^\S{1,10}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HighAvailabilityConfig`

###### Important

AWS Mainframe Modernization Service (Managed Runtime Environment experience) will no longer be open to new customers starting on November 7, 2025. If you would like to use the service, please sign up prior to November 7, 2025. For capabilities similar to AWS Mainframe Modernization Service (Managed Runtime Environment experience) explore AWS Mainframe Modernization Service (Self-Managed Experience). Existing customers can
continue to use the service as normal. For more information, see
[AWS Mainframe Modernization availability change](../../../m2/latest/userguide/mainframe-modernization-availability-change.md).

Defines the details of a high availability configuration.

_Required_: No

_Type_: [HighAvailabilityConfig](aws-properties-m2-environment-highavailabilityconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The instance type of the runtime environment.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,20}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The identifier of a customer managed key.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the runtime environment.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9_\-]{1,59}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkType`

The network type supported by the runtime environment.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | dual`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreferredMaintenanceWindow`

Configures the maintenance window that you want for the runtime environment. The maintenance window must have the format `ddd:hh24:mi-ddd:hh24:mi` and must be less than 24 hours. The following two examples are valid maintenance windows: `sun:23:45-mon:00:15` or `sat:01:00-sat:03:00`.

If you do not provide a value, a random system-generated value will be assigned.

_Required_: No

_Type_: String

_Pattern_: `^\S{1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

Specifies whether the runtime environment is publicly accessible.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The list of security groups for the VPC associated with this runtime environment.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageConfigurations`

###### Important

AWS Mainframe Modernization Service (Managed Runtime Environment experience) will no longer be open to new customers starting on November 7, 2025. If you would like to use the service, please sign up prior to November 7, 2025. For capabilities similar to AWS Mainframe Modernization Service (Managed Runtime Environment experience) explore AWS Mainframe Modernization Service (Self-Managed Experience). Existing customers can
continue to use the service as normal. For more information, see
[AWS Mainframe Modernization availability change](../../../m2/latest/userguide/mainframe-modernization-availability-change.md).

Defines the storage configuration for a runtime environment.

_Required_: No

_Type_: Array of [StorageConfiguration](aws-properties-m2-environment-storageconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The list of subnets associated with the VPC for this runtime environment.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:).+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the environment Amazon Resource Name (ARN), such as the
following:

`{ "Ref": “SampleEnv” }`

Returns a value similar to the following:

`arn:aws:m2:us-west-2:1234567890:env/y3ca6bhaife2bcvxar3lpivfou`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EnvironmentArn`

The Amazon Resource Name (ARN) of the runtime environment.

`EnvironmentId`

The unique identifier of the runtime environment.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::M2::Deployment

EfsStorageConfiguration

All content copied from https://docs.aws.amazon.com/.
