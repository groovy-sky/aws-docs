This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# CloudFormation resource tagging

You can apply tags to resources by using the `Tags` property in your CloudFormation
template, which can help you identify and categorize those resources.

For information about which resources support the `Tags` property, see the
individual resources in the [AWS resource and property types reference](aws-template-resource-type-ref.md). If a resource doesn't yet support the
`Tags` property, the resource's service might support tagging using its own API
operations. For more information, refer to the documentation for that service.

In addition to any tags you define, CloudFormation automatically creates the following
stack-level tags with the `aws:` prefix:

- `aws:cloudformation:logical-id`

- `aws:cloudformation:stack-id`

- `aws:cloudformation:stack-name`

The `aws:` prefix is reserved for AWS use. This prefix is case-insensitive. If
you use this prefix in the `Key` or `Value` property, you can't update
or delete the tag. Tags with this prefix don't count toward the number of tags per
resource.

The propagation of stack-level tags to resources, including tags with the `aws:`
prefix, varies by resource type. For example, tags aren't propagated to Amazon EBS volumes that are
created from block device mappings.

###### Note

Some resources require explicit tag propagation settings. For example, the
`AWS::AutoScaling::AutoScalingGroup` resource must have its
`PropagateAtLaunch` property set to `true` to propagate tags to
its EC2 instances. However, stack-level tags are automatically applied to EC2 instances
regardless of the `PropagateAtLaunch` setting.

## Syntax

### JSON

```json

{
  "Key" : String,
  "Value" : String
}
```

### YAML

```yaml

Key: String
Value: String
```

## Properties

`Key`

The key name of the tag. You can specify a value that's 1 to 128 Unicode
characters in length and can't be prefixed with `aws:`. You can use any
of the following characters: the set of Unicode letters, digits, whitespace,
`_`, `.`, `:`, `/`, `=`,
`+`, `@`, `-`, and `"`.

_Required_: Yes

_Type_: String

`Value`

The value for the tag. You can specify a value that's 1 to 256 characters in
length. You can use any of the following characters: the set of Unicode letters,
digits, whitespace, `_`, `.`, `/`,
`=`, `+`, and `-`.

_Required_: Yes

_Type_: String

## Example

This example shows a `Tags` property. You specify this property within the
`Properties` section of a resource that supports it. When the resource is
created, the `Environment` tag is dynamically set using a parameter, and the
`Owner` tag is statically set to `MyName`.

### JSON

```json

"Tags" : [
   {
      "Key" : "Environment",
      "Value" : { "Ref": "Environment" }
   },
   {
      "Key" : "Owner",
      "Value" : "MyName"
   }
]
```

### YAML

```yaml

Tags:
  - Key: Environment
    Value: !Ref Environment
  - Key: Owner
    Value: MyName
```

## See also

- [Configure\
stack options](../userguide/cfn-console-create-stack.md#configure-stack-options) in the _AWS CloudFormation User Guide_

- [Viewing CloudFormation\
stack data and resources on the AWS Management Console](../userguide/cfn-console-view-stack-data-resources.md) in the
_AWS CloudFormation User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Name

Resource attributes

All content copied from https://docs.aws.amazon.com/.
