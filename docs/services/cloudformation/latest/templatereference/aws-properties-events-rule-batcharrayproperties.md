This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule BatchArrayProperties

The array properties for the submitted job, such as the size of the array. The array size
can be between 2 and 10,000. If you specify array properties for a job, it becomes an array
job. This parameter is used only if the target is an AWS Batch job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Size" : Integer
}

```

### YAML

```yaml

  Size: Integer

```

## Properties

`Size`

The size of the array, if this is an array batch job. Valid values are integers between 2
and 10,000.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Set the BatchArrayProperties property

The following example sets the `BatchArrayProperties` property to 950.

#### JSON

```json

"ArrayProperties": {
  "Size": 950
}
```

#### YAML

```yaml

ArrayProperties:
  Size: 950
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsVpcConfiguration

BatchParameters

All content copied from https://docs.aws.amazon.com/.
