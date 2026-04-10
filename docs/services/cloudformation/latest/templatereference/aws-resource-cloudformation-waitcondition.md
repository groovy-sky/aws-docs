This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::WaitCondition

The `AWS::CloudFormation::WaitCondition` resource provides a way to
coordinate stack resource creation with configuration actions that are external to the
stack creation or to track the status of a configuration process. In these situations,
we recommend that you associate a `CreationPolicy` attribute with the wait
condition instead of using a wait condition handle. For more information and an example,
see [CreationPolicy attribute](aws-attribute-creationpolicy.md) in the _CloudFormation User_
_Guide_. If you use a `CreationPolicy` with a wait condition,
don't specify any of the wait condition's properties.

###### Note

If you use AWS PrivateLink, resources in the VPC that respond to wait
conditions must have access to CloudFormation, specific Amazon S3
buckets. Resources must send wait condition responses to a presigned Amazon S3 URL. If they can't send responses to Amazon S3, CloudFormation
won't receive a response and the stack operation fails. For more information, see
[Access\
CloudFormation using an interface endpoint (AWS PrivateLink)](../userguide/vpc-interface-endpoints.md) in the _CloudFormation User_
_Guide_.

###### Important

For Amazon EC2 and Auto Scaling resources, we recommend that you use
a `CreationPolicy` attribute instead of wait conditions. Add a
`CreationPolicy` attribute to those resources, and use the
`cfn-signal` helper script to signal when an instance creation
process has completed successfully.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::WaitCondition",
  "Properties" : {
      "Count" : Integer,
      "Handle" : String,
      "Timeout" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::WaitCondition
Properties:
  Count: Integer
  Handle: String
  Timeout: String

```

## Properties

`Count`

The number of success signals that CloudFormation must receive before it
continues the stack creation process. When the wait condition receives the requisite
number of success signals, CloudFormation resumes the creation of the stack. If
the wait condition doesn't receive the specified number of success signals before the
Timeout period expires, CloudFormation assumes that the wait condition has
failed and rolls the stack back.

Updates aren't supported.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Handle`

A reference to the wait condition handle used to signal this wait condition. Use the
`Ref` intrinsic function to specify an [AWS::CloudFormation::WaitConditionHandle](aws-resource-cloudformation-waitconditionhandle.md) resource.

Anytime you add a `WaitCondition` resource during a stack update, you must
associate the wait condition with a new WaitConditionHandle resource. Don't reuse an old
wait condition handle that has already been defined in the template. If you reuse a wait
condition handle, the wait condition might evaluate old signals from a previous create
or update stack command.

Updates aren't supported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The length of time (in seconds) to wait for the number of signals that the
`Count` property specifies. `Timeout` is a minimum-bound
property, meaning the timeout occurs no sooner than the time you specify, but can occur
shortly thereafter. The maximum time that can be specified for this property is 12 hours
(43200 seconds).

Updates aren't supported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

Returns a unique identifier for the resource.

## Examples

### WaitCondition that has a timeout of 300 seconds

#### JSON

```json

{
    "WaitHandle": {
        "Type": "AWS::CloudFormation::WaitConditionHandle"
    },
    "WaitCondition": {
        "Type": "AWS::CloudFormation::WaitCondition",
        "DependsOn": "MyResource",
        "Properties": {
            "Handle": {
                "Ref": "WaitHandle"
            },
            "Timeout": "300"
        }
    }
}
```

#### YAML

```yaml

WaitHandle:
  Type: AWS::CloudFormation::WaitConditionHandle
WaitCondition:
  Type: AWS::CloudFormation::WaitCondition
  DependsOn: MyResource
  Properties:
    Handle: !Ref 'WaitHandle'
    Timeout: '300'
```

## See also

- [Create wait conditions in a CloudFormation template](../userguide/using-cfn-waitcondition.md) in
the _CloudFormation User Guide_

- [DependsOn attribute](aws-attribute-dependson.md) in the _CloudFormation User_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingConfig

AWS::CloudFormation::WaitConditionHandle

All content copied from https://docs.aws.amazon.com/.
