This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBInstance ProcessorFeature

The `ProcessorFeature` property type specifies the processor features of a
DB instance class.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the processor feature. Valid names are `coreCount` and `threadsPerCore`.

_Required_: No

_Type_: String

_Allowed values_: `coreCount | threadsPerCore`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of a processor feature.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example specifies processor features of an Oracle DB instance
class.

### Authentication details

#### JSON

```json

"ProcessorFeatures":[
   {
      "Name": "threadsPerCore",
      "Value": "4"
   },
   {
      "Name": "coreCount",
      "Value": "4"
   }
]
```

#### YAML

```yaml

ProcessorFeatures:
  - Name: threadsPerCore
    Value: 4
  - Name: coreCount
    Value: 4
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MasterUserSecret

Tag

All content copied from https://docs.aws.amazon.com/.
