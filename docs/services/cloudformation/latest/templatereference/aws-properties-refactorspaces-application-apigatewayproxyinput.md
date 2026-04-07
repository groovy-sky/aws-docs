This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RefactorSpaces::Application ApiGatewayProxyInput

A wrapper object holding the Amazon API Gateway endpoint input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndpointType" : String,
  "StageName" : String
}

```

### YAML

```yaml

  EndpointType: String
  StageName: String

```

## Properties

`EndpointType`

The type of endpoint to use for the API Gateway proxy. If no value is specified in
the request, the value is set to `REGIONAL` by default.

If the value is set to `PRIVATE` in the request, this creates a private API
endpoint that is isolated from the public internet. The private endpoint can only be accessed
by using Amazon Virtual Private Cloud (Amazon VPC) interface endpoints for the Amazon API Gateway that has been granted access. For more information about creating a private
connection with Refactor Spaces and interface endpoint (AWS PrivateLink) availability,
see [Access\
Refactor Spaces using an interface endpoint (AWS PrivateLink)](https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/userguide/vpc-interface-endpoints.html).

_Required_: No

_Type_: String

_Allowed values_: `REGIONAL | PRIVATE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StageName`

The name of the API Gateway stage. The name defaults to `prod`.

_Required_: No

_Type_: String

_Pattern_: `^[-a-zA-Z0-9_]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::RefactorSpaces::Application

Tag
