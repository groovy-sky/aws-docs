This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::M2::Environment HighAvailabilityConfig

###### Important

AWS Mainframe Modernization Service (Managed Runtime Environment experience) will no longer be open to new customers starting on November 7, 2025. If you would like to use the service, please sign up prior to November 7, 2025. For capabilities similar to AWS Mainframe Modernization Service (Managed Runtime Environment experience) explore AWS Mainframe Modernization Service (Self-Managed Experience). Existing customers can
continue to use the service as normal. For more information, see
[AWS Mainframe Modernization availability change](../../../m2/latest/userguide/mainframe-modernization-availability-change.md).

Defines the details of a high availability configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DesiredCapacity" : Integer
}

```

### YAML

```yaml

  DesiredCapacity: Integer

```

## Properties

`DesiredCapacity`

The number of instances in a high availability configuration. The minimum possible value is 1 and the maximum is 100.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FsxStorageConfiguration

StorageConfiguration

All content copied from https://docs.aws.amazon.com/.
