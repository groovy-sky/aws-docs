This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application

The `AWS::EMRServerless::Application` resource specifies an EMR Serverless
application. An application uses open source analytics frameworks to run jobs that
process data. To create an application, you must specify the release version for the
open source framework version you want to use and the type of application you want, such
as Apache Spark or Apache Hive. After you create an application, you can submit data
processing jobs or interactive requests to it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMRServerless::Application",
  "Properties" : {
      "Architecture" : String,
      "AutoStartConfiguration" : AutoStartConfiguration,
      "AutoStopConfiguration" : AutoStopConfiguration,
      "IdentityCenterConfiguration" : IdentityCenterConfiguration,
      "ImageConfiguration" : ImageConfigurationInput,
      "InitialCapacity" : [ InitialCapacityConfigKeyValuePair, ... ],
      "InteractiveConfiguration" : InteractiveConfiguration,
      "MaximumCapacity" : MaximumAllowedResources,
      "MonitoringConfiguration" : MonitoringConfiguration,
      "Name" : String,
      "NetworkConfiguration" : NetworkConfiguration,
      "ReleaseLabel" : String,
      "RuntimeConfiguration" : [ ConfigurationObject, ... ],
      "SchedulerConfiguration" : SchedulerConfiguration,
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "WorkerTypeSpecifications" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::EMRServerless::Application
Properties:
  Architecture: String
  AutoStartConfiguration:
    AutoStartConfiguration
  AutoStopConfiguration:
    AutoStopConfiguration
  IdentityCenterConfiguration:
    IdentityCenterConfiguration
  ImageConfiguration:
    ImageConfigurationInput
  InitialCapacity:
    - InitialCapacityConfigKeyValuePair
  InteractiveConfiguration:
    InteractiveConfiguration
  MaximumCapacity:
    MaximumAllowedResources
  MonitoringConfiguration:
    MonitoringConfiguration
  Name: String
  NetworkConfiguration:
    NetworkConfiguration
  ReleaseLabel: String
  RuntimeConfiguration:
    - ConfigurationObject
  SchedulerConfiguration:
    SchedulerConfiguration
  Tags:
    - Tag
  Type: String
  WorkerTypeSpecifications:
    Key: Value

```

## Properties

`Architecture`

The CPU architecture of an application.

_Required_: No

_Type_: String

_Allowed values_: `ARM64 | X86_64`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`AutoStartConfiguration`

The configuration for an application to automatically start on job submission.

_Required_: No

_Type_: [AutoStartConfiguration](aws-properties-emrserverless-application-autostartconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`AutoStopConfiguration`

The configuration for an application to automatically stop after a certain amount of
time being idle.

_Required_: No

_Type_: [AutoStopConfiguration](aws-properties-emrserverless-application-autostopconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`IdentityCenterConfiguration`

The IAM Identity Center configuration applied to enable trusted identity propagation.

_Required_: No

_Type_: [IdentityCenterConfiguration](aws-properties-emrserverless-application-identitycenterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageConfiguration`

The image configuration applied to all worker types.

_Required_: No

_Type_: [ImageConfigurationInput](aws-properties-emrserverless-application-imageconfigurationinput.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`InitialCapacity`

The initial capacity of the application.

_Required_: No

_Type_: Array of [InitialCapacityConfigKeyValuePair](aws-properties-emrserverless-application-initialcapacityconfigkeyvaluepair.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`InteractiveConfiguration`

The interactive configuration object that enables the interactive use cases for an application.

_Required_: No

_Type_: [InteractiveConfiguration](aws-properties-emrserverless-application-interactiveconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MaximumCapacity`

The maximum capacity of the application. This is cumulative across all workers at any
given point in time during the lifespan of the application is created. No new resources
will be created once any one of the defined limits is hit.

_Required_: No

_Type_: [MaximumAllowedResources](aws-properties-emrserverless-application-maximumallowedresources.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MonitoringConfiguration`

A configuration specification to be used when provisioning an application. A
configuration consists of a classification, properties, and optional nested configurations.
A classification refers to an application-specific configuration file. Properties are the
settings you want to change in that file.

_Required_: No

_Type_: [MonitoringConfiguration](aws-properties-emrserverless-application-monitoringconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Name`

The name of the application.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9._\/#-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkConfiguration`

The network configuration for customer VPC connectivity for the application.

_Required_: No

_Type_: [NetworkConfiguration](aws-properties-emrserverless-application-networkconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ReleaseLabel`

The EMR release associated with the application.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9._/-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`RuntimeConfiguration`

The [Configuration](../../../../reference/emr-serverless/latest/apireference/api-configuration.md)
specifications of an application. Each configuration consists of a classification and properties. You use this
parameter when creating or updating an application. To see the runtimeConfiguration object of an application,
run the [GetApplication](../../../../reference/emr-serverless/latest/apireference/api-getapplication.md) API operation.

_Required_: No

_Type_: Array of [ConfigurationObject](aws-properties-emrserverless-application-configurationobject.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SchedulerConfiguration`

The scheduler configuration for batch and streaming jobs running on this application. Supported with release labels emr-7.0.0 and above.

_Required_: No

_Type_: [SchedulerConfiguration](aws-properties-emrserverless-application-schedulerconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

The tags assigned to the application.

_Required_: No

_Type_: Array of [Tag](aws-properties-emrserverless-application-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of application, such as Spark or Hive.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkerTypeSpecifications`

The specification applied to each worker type.

_Required_: No

_Type_: Object of [WorkerTypeSpecificationInput](aws-properties-emrserverless-application-workertypespecificationinput.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the application.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationId`

The ID of the application, such as `ab4rp1abcs8xz47n3x0example`.

`Arn`

The Amazon Resource Name (ARN) of the EMR Serverless Application.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon EMR Serverless

AutoStartConfiguration

All content copied from https://docs.aws.amazon.com/.
