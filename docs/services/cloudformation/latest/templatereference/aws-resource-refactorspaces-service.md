This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RefactorSpaces::Service

###### Note

AWS Migration Hub is no longer open to new customers as of November 7, 2025. For capabilities similar to AWS Migration Hub, explore [AWS Migration Hub](https://aws.amazon.com/transform).

Creates an AWS Migration Hub Refactor Spaces service. The account owner of the service is always the
environment owner, regardless of which account in the environment creates the service.
Services have either a URL endpoint in a virtual private cloud (VPC), or a Lambda
function endpoint.

###### Important

If an AWS resource is launched in a service VPC, and you want it to be
accessible to all of an environment’s services with VPCs and routes, apply the
`RefactorSpacesSecurityGroup` to the resource. Alternatively, to add more
cross-account constraints, apply your own security group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RefactorSpaces::Service",
  "Properties" : {
      "ApplicationIdentifier" : String,
      "Description" : String,
      "EndpointType" : String,
      "EnvironmentIdentifier" : String,
      "LambdaEndpoint" : LambdaEndpointInput,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "UrlEndpoint" : UrlEndpointInput,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::RefactorSpaces::Service
Properties:
  ApplicationIdentifier: String
  Description: String
  EndpointType: String
  EnvironmentIdentifier: String
  LambdaEndpoint:
    LambdaEndpointInput
  Name: String
  Tags:
    - Tag
  UrlEndpoint:
    UrlEndpointInput
  VpcId: String

```

## Properties

`ApplicationIdentifier`

The unique identifier of the application.

_Required_: Yes

_Type_: String

_Pattern_: `^app-([0-9A-Za-z]{10}$)`

_Minimum_: `14`

_Maximum_: `14`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the service.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_\s\.\!\*\#\@\']+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointType`

The endpoint type of the service.

_Required_: Yes

_Type_: String

_Allowed values_: `LAMBDA | URL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentIdentifier`

The unique identifier of the environment.

_Required_: Yes

_Type_: String

_Pattern_: `^env-([0-9A-Za-z]{10}$)`

_Minimum_: `14`

_Maximum_: `14`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LambdaEndpoint`

A summary of the configuration for the AWS Lambda endpoint type.

_Required_: No

_Type_: [LambdaEndpointInput](aws-properties-refactorspaces-service-lambdaendpointinput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the service.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!svc-)[a-zA-Z0-9]+[a-zA-Z0-9-_ ]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the service.

_Required_: No

_Type_: Array of [Tag](aws-properties-refactorspaces-service-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UrlEndpoint`

The summary of the configuration for the URL endpoint type.

_Required_: No

_Type_: [UrlEndpointInput](aws-properties-refactorspaces-service-urlendpointinput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The ID of the virtual private cloud (VPC).

_Required_: No

_Type_: String

_Pattern_: `^vpc-[-a-f0-9]{8}([-a-f0-9]{9})?$`

_Minimum_: `12`

_Maximum_: `21`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a composite ID following this format:
`<EnvironmentId>|<ApplicationId>|<ServiceId>`. For example,
`env-1234654123|app-1234654123|svc-1234654123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the service.

`ServiceIdentifier`

The unique identifier of the service.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UriPathRouteInput

LambdaEndpointInput

All content copied from https://docs.aws.amazon.com/.
