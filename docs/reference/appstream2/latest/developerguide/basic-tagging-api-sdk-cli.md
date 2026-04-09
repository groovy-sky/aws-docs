# Working with Tags by Using the Amazon WorkSpaces Applications API, an AWS SDK, or AWS CLI

If you're using the WorkSpaces Applications API, an AWS SDK, or the AWS Command Line Interface (AWS
CLI), you can use the following WorkSpaces Applications operations with the `tags` parameter
to add tags when you create new resources.

###### Note

You can use spaces in tag keys and values. To indicate a space when you use the AWS CLI, use "\\s" (without the quotation marks).

TaskAWS CLIAPI OperationAdd one or more tags for a new fleet[create-fleet](../../../../services/cli/latest/reference/appstream/create-fleet.md)

[CreateFleet](../apireference/api-createfleet.md#AppStream2-CreateFleet-request-Tags)

Add one or more tags for a new image builder[create-imagebuilder](../../../../services/cli/latest/reference/appstream/create-imagebuilder.md)

[CreateImageBuilder](../apireference/api-createimagebuilder.md#AppStream2-CreateImageBuilder-request-Tags)

Add one or more tags for a new stack

[create-stack](../../../../services/cli/latest/reference/appstream/create-stack.md)

[CreateStack](../apireference/api-createstack.md#AppStream2-CreateStack-request-Tags)

You can use the following WorkSpaces Applications operations to add, edit, remove, or list tags for existing
resources:

TaskAWS CLIAPI OperationAdd or overwrite one or more tags for a resource[tag-resource](../../../../services/cli/latest/reference/appstream/tag-resource.md)

[TagResource](../apireference/api-tagresource.md)

Remove one or more tags for a resource[untag-resource](../../../../services/cli/latest/reference/appstream/untag-resource.md)

[UntagResource](../apireference/api-untagresource.md)

List one or more tags for a resource

[list-tags-for-resource](../../../../services/cli/latest/reference/appstream/list-tags-for-resource.md)

[ListTagsForResource](../apireference/api-listtagsforresource.md)

When you use the WorkSpaces Applications API, an AWS SDK, or AWS CLI actions to add, edit, remove, or
list tags for an existing WorkSpaces Applications resource, specify the resource by using its Amazon Resource Name
(ARN). An ARN uniquely identifies an AWS resource and uses the following general
syntax.

```nohighlight

arn:aws:appstream:region:account:resourceType/resourceName
```

**`region`**

The AWS Region in which the resource was created (for example,
`us-east-1`).

**`account`**

The AWS account ID, with no hyphens (for example, `123456789012`).

**`resourceType`**

The type of resource. You can tag the following WorkSpaces Applications resource types:
`image-builder`, `image`, `fleet`, and
`stack`.

**`resourceName`**

The name of the resource.

For example, you can obtain the ARN for an WorkSpaces Applications fleet by using the AWS CLI [describe-fleets](../../../../services/cli/latest/reference/appstream/describe-fleets.md) command.
Copy the following command.

```nohighlight

aws appstream describe-fleets
```

For an environment that contains a single fleet named `TestFleet`, the ARN
for this resource would appear in the JSON output similar to the following.

```json

"Arn": "arn:aws:appstream:us-east-1:123456789012:fleet/TestFleet"
```

After you obtain the ARN for this resource, you can add two tags by using the [tag-resource](../../../../services/cli/latest/reference/appstream/tag-resource.md) command:

```nohighlight

aws appstream tag-resource --resource arn:awsappstream:us-east-1:123456789012:fleet/TestFleet --tags Environment=Test,Department=IT
```

The first tag, `Environment=Test`, indicates that the fleet is in a test
environment. The second tag, `Department=IT`, indicates that the fleet is in
the IT department.

You can use the following command to list the two tags that you added to the
fleet.

```nohighlight

aws appstream list-tags-for-resource --resource arn:aws:appstream:us-east-1:123456789012:fleet/TestFleet
```

For this example, the JSON output appears as follows:

```json

{
    "Tags": {
       "Environment" : "Test",
       "Department" : "IT"
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Adding, Editing, and Deleting Tags

Monitoring and Reporting

All content copied from https://docs.aws.amazon.com/.
