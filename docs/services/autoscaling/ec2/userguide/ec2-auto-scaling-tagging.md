# Tag Auto Scaling groups and instances

A _tag_ is a custom attribute label that you assign
or that AWS assigns to an AWS resource. Each tag has two parts:

- A tag key (for example, `costcenter`, `environment`, or
`project`)

- An optional field known as a tag value (for example, `111122223333`
or `production`)

Tags help you do the following:

- Track your AWS costs. You activate these tags on the AWS Billing and Cost Management dashboard.
AWS uses the tags to categorize your costs and deliver a monthly cost
allocation report to you. For more information, see [Using cost allocation tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md)
in the _AWS Billing User Guide_.

- Control access to Auto Scaling groups based on tags. You can use conditions in your
IAM policies to control access to Auto Scaling groups based on the tags on that group.
For more information, see [Tags for security](https://docs.aws.amazon.com/autoscaling/ec2/userguide/tag-security.html).

- Filter and search for Auto Scaling groups based on the tags that you add. For more
information, see [Use tags to filter Auto Scaling groups](https://docs.aws.amazon.com/autoscaling/ec2/userguide/use-tag-filters-aws-cli.html).

- Identify and organize your AWS resources. Many AWS services support
tagging, so you can assign the same tag to resources from different services to
indicate that the resources are related.

You can tag new or existing Auto Scaling groups. You can also propagate tags from an Auto Scaling
group to the EC2 instances that it launches.

Tags are not propagated to Amazon EBS volumes. To add tags to Amazon EBS volumes, specify the
tags in a launch template. For more information, see [Create a launch template for an Auto Scaling group](create-launch-template.md).

You can create and manage tags through the AWS Management Console, AWS CLI, or SDKs.

###### Contents

- [Tag naming and usage restrictions](#tag_restrictions)

- [EC2 instance tagging lifecycle](#tag-lifecycle)

- [Tag your Auto Scaling groups](https://docs.aws.amazon.com/autoscaling/ec2/userguide/add-tags.html)

- [Delete tags](https://docs.aws.amazon.com/autoscaling/ec2/userguide/delete-tag.html)

- [Tags for security](https://docs.aws.amazon.com/autoscaling/ec2/userguide/tag-security.html)

- [Control access to tags](https://docs.aws.amazon.com/autoscaling/ec2/userguide/tag-permissions.html)

- [Use tags to filter Auto Scaling groups](https://docs.aws.amazon.com/autoscaling/ec2/userguide/use-tag-filters-aws-cli.html)

## Tag naming and usage restrictions

The following basic restrictions apply to tags:

- The maximum number of tags per resource is 50.

- The maximum number of tags that you can add or remove using a single call
is 25.

- The maximum key length is 128 Unicode characters.

- The maximum value length is 256 Unicode characters.

- Tag keys and values are case-sensitive. As a best practice, decide on a
strategy for capitalizing tags, and consistently implement that strategy
across all resource types.

- Do not use the `aws:` prefix in your tag names or values,
because it is reserved for AWS use. You can't edit or delete tag names or
values with this prefix, and they do not count toward your tags per resource
quota.

## EC2 instance tagging lifecycle

If you have opted to propagate tags to your EC2 instances, the tags are managed as
follows:

- When an Auto Scaling group launches instances, it adds tags to the instances
during resource creation rather than after the resource is created.

- The Auto Scaling group automatically adds a tag to instances with a key of
`aws:autoscaling:groupName` and a value of the Auto Scaling group
name.

- If you specify instance tags in your launch template and you opted to
propagate your group's tags to its instances, all the tags are merged. If
the same tag key is specified for a tag in your launch template and a tag in
your Auto Scaling group, then the tag value from the group takes precedence.

- When you attach existing instances, the Auto Scaling group adds the tags to the
instances, overwriting any existing tags with the same tag key. It also adds
a tag with a key of `aws:autoscaling:groupName` and a value of
the Auto Scaling group name.

- When you detach an instance from an Auto Scaling group, it removes only the
`aws:autoscaling:groupName` tag.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Update an Auto Scaling group

Tag your Auto Scaling groups
