This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::WaitConditionHandle

The `AWS::CloudFormation::WaitConditionHandle` type has no properties. When
you reference the `WaitConditionHandle` resource by using the
`Ref` function, CloudFormation returns a presigned URL. You pass
this URL to applications or scripts that are running on your Amazon EC2
instances to send signals to that URL. An associated
`AWS::CloudFormation::WaitCondition` resource checks the URL for the
required number of success signals or for a failure signal.

For more information, see [Create wait\
conditions in a CloudFormation template](../userguide/using-cfn-waitcondition.md) in the _CloudFormation User Guide_.

Anytime you add a `WaitCondition` resource during a stack update or update
a resource with a wait condition, you must associate the wait condition with a new
`WaitConditionHandle` resource. Don't reuse an old wait condition handle
that has already been defined in the template. If you reuse a wait condition handle, the
wait condition might evaluate old signals from a previous create or update stack
command.

Updates aren't supported for this resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::WaitConditionHandle"
}

```

### YAML

```yaml

Type: AWS::CloudFormation::WaitConditionHandle
```

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

Returns a unique identifier for the resource.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFormation::WaitCondition

Next

All content copied from https://docs.aws.amazon.com/.
