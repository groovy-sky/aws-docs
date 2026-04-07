This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::SessionLogger

The session logger resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::SessionLogger",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "CustomerManagedKey" : String,
      "DisplayName" : String,
      "EventFilter" : EventFilter,
      "LogConfiguration" : LogConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::SessionLogger
Properties:
  AdditionalEncryptionContext:
    Key: Value
  CustomerManagedKey: String
  DisplayName: String
  EventFilter:
    EventFilter
  LogConfiguration:
    LogConfiguration
  Tags:
    - Tag

```

## Properties

`AdditionalEncryptionContext`

The additional encryption context of the session logger.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomerManagedKey`

The custom managed key of the session logger.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisplayName`

The human-readable display name.

_Required_: No

_Type_: String

_Pattern_: `^[ _\-\d\w]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventFilter`

The filter that specifies which events to monitor.

_Required_: Yes

_Type_: [EventFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-sessionlogger-eventfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

The configuration that specifies where logs are fowarded.

_Required_: Yes

_Type_: [LogConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-sessionlogger-logconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags of the session logger.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-sessionlogger-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AssociatedPortalArns`

The associated portal ARN.

`CreationDate`

The date the session logger resource was created.

`SessionLoggerArn`

The ARN of the session logger resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

EventFilter
