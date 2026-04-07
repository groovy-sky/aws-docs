This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service

The `AWS::DevOpsAgent::Service` resource registers an external service for integration with the
AWS DevOps Agent service.

This resource does not support updates. To modify a registered service, you must replace the resource.

###### Note

Services that use OAuth authorization flows (ex: Datadog, GitHub, Slack) require interactive browser-based
registration through the AWS DevOps Agent console. These service types cannot be registered using
this resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DevOpsAgent::Service",
  "Properties" : {
      "KmsKeyArn" : String,
      "ServiceDetails" : ServiceDetails,
      "ServiceType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DevOpsAgent::Service
Properties:
  KmsKeyArn: String
  ServiceDetails:
    ServiceDetails
  ServiceType: String
  Tags:
    - Tag

```

## Properties

`KmsKeyArn`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceDetails`

Service-specific configuration details provided during registration. The structure of this property depends on
the value of `ServiceType`.

_Required_: No

_Type_: [ServiceDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsagent-service-servicedetails.html)

_Update requires_: Updates are not supported.

`ServiceType`

The type of external service to register.

_Required_: Yes

_Type_: String

_Allowed values_: `dynatrace | mcpserver | mcpserversplunk | mcpservernewrelic | gitlab | servicenow`

_Update requires_: Updates are not supported.

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devopsagent-service-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ServiceId.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccessibleResources`

List of accessible resources for this service.

`Arn`

Property description not available.

`ServiceId`

The unique identifier of the service.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SourceAwsConfiguration

AdditionalServiceDetails
