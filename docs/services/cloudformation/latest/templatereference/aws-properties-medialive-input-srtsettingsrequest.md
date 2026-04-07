This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Input SrtSettingsRequest

The `SrtSettingsRequest` property type specifies Property description not available. for an [AWS::MediaLive::Input](aws-resource-medialive-input.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SrtCallerSources" : [ SrtCallerSourceRequest, ... ],
  "SrtListenerSettings" : SrtListenerSettingsRequest
}

```

### YAML

```yaml

  SrtCallerSources:
    - SrtCallerSourceRequest
  SrtListenerSettings:
    SrtListenerSettingsRequest

```

## Properties

`SrtCallerSources`

Property description not available.

_Required_: No

_Type_: Array of [SrtCallerSourceRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-srtcallersourcerequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtListenerSettings`

Property description not available.

_Required_: No

_Type_: [SrtListenerSettingsRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-srtlistenersettingsrequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SrtListenerSettingsRequest

AWS::MediaLive::InputSecurityGroup
