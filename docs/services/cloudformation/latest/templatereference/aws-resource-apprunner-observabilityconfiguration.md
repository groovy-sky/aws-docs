This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::ObservabilityConfiguration

The `AWS::AppRunner::ObservabilityConfiguration` resource is an AWS App Runner resource type that specifies an App Runner
observability configuration.

App Runner requires this resource when you specify App Runner services and you want to enable non-default observability features.
You can share an observability configuration across multiple services.

Create multiple revisions of a configuration by specifying this resource multiple times using the same `ObservabilityConfigurationName`.
App Runner creates multiple resources with incremental `ObservabilityConfigurationRevision` values. When you specify a service and
configure an observability configuration resource, the service uses the latest active revision of the observability configuration by default. You can
optionally configure the service to use a specific revision.

The observability configuration resource is designed to configure multiple features (currently one feature, tracing). This resource takes optional
parameters that describe the configuration of these features (currently one parameter, `TraceConfiguration`). If you don't specify a feature
parameter, App Runner doesn't enable the feature.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppRunner::ObservabilityConfiguration",
  "Properties" : {
      "ObservabilityConfigurationName" : String,
      "Tags" : [ Tag, ... ],
      "TraceConfiguration" : TraceConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::AppRunner::ObservabilityConfiguration
Properties:
  ObservabilityConfigurationName: String
  Tags:
    - Tag
  TraceConfiguration:
    TraceConfiguration

```

## Properties

`ObservabilityConfigurationName`

A name for the observability configuration. When you use it for the first time in an AWS Region, App Runner creates revision number
`1` of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.

###### Note

The name `DefaultConfiguration` is reserved. You can't use it to create a new observability configuration, and you can't create a
revision of it.

When you want to use your own observability configuration for your App Runner service, _create a configuration with a different name_,
and then provide it when you create or update your service.

If you don't specify a name, CloudFormation generates a name for your observability configuration.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9][A-Za-z0-9\-_]{3,31}`

_Minimum_: `4`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of metadata items that you can associate with your observability configuration resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apprunner-observabilityconfiguration-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TraceConfiguration`

The configuration of the tracing feature within this observability configuration. If you don't specify it, App Runner doesn't enable tracing.

_Required_: No

_Type_: [TraceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apprunner-observabilityconfiguration-traceconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Latest`

It's set to `true` for the configuration with the highest `Revision` among all configurations that share the same
`ObservabilityConfigurationName`. It's set to `false` otherwise.

`ObservabilityConfigurationArn`

The Amazon Resource Name (ARN) of this observability configuration.

`ObservabilityConfigurationRevision`

The revision of this observability configuration. It's unique among all the active configurations ( `"Status": "ACTIVE"`) that share the
same `ObservabilityConfigurationName`.

## Examples

### Observability configuration to enable tracing

This example illustrates creating an observability configuration that enables tracing using AWS X-Ray.

#### JSON

```json

{
  "Type": "AWS::AppRunner::ObservabilityConfiguration",
  "Properties": {
    "ObservabilityConfigurationName": "xray-tracing",
    "TraceConfiguration": {
      "Vendor": "AWSXRAY"
    }
  }
}
```

#### YAML

```yaml

Type: AWS::AppRunner::ObservabilityConfiguration
Properties:
  ObservabilityConfigurationName: xray-tracing
  TraceConfiguration:
    Vendor: AWSXRAY
```

## See also

- [Tracing for your App Runner application with X-Ray](https://docs.aws.amazon.com/apprunner/latest/dg/monitor-xray.html) in the _AWS App Runner Developer Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
