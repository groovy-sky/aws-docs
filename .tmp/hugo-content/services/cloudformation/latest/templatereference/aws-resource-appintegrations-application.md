This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppIntegrations::Application

Creates and persists an Application resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppIntegrations::Application",
  "Properties" : {
      "ApplicationConfig" : ApplicationConfig,
      "ApplicationSourceConfig" : ApplicationSourceConfig,
      "ApplicationType" : String,
      "Description" : String,
      "IframeConfig" : IframeConfig,
      "InitializationTimeout" : Integer,
      "IsService" : Boolean,
      "Name" : String,
      "Namespace" : String,
      "Permissions" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppIntegrations::Application
Properties:
  ApplicationConfig:
    ApplicationConfig
  ApplicationSourceConfig:
    ApplicationSourceConfig
  ApplicationType: String
  Description: String
  IframeConfig:
    IframeConfig
  InitializationTimeout: Integer
  IsService: Boolean
  Name: String
  Namespace: String
  Permissions:
    - String
  Tags:
    - Tag

```

## Properties

`ApplicationConfig`

Property description not available.

_Required_: No

_Type_: [ApplicationConfig](aws-properties-appintegrations-application-applicationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationSourceConfig`

The configuration for where the application should be loaded from.

_Required_: Yes

_Type_: [ApplicationSourceConfig](aws-properties-appintegrations-application-applicationsourceconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationType`

The type of application.

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | SERVICE | MCP_SERVER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the application.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IframeConfig`

Property description not available.

_Required_: No

_Type_: [IframeConfig](aws-properties-appintegrations-application-iframeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InitializationTimeout`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsService`

Indicates whether the application is a service.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the application.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\/\._ \-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the application.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9/\._\-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, {
"tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](aws-properties-appintegrations-application-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ApplicationArn`

The Amazon Resource Name (ARN) of the Application.

`Id`

A unique identifier for the Application.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon AppIntegrations

ApplicationConfig

All content copied from https://docs.aws.amazon.com/.
