# Import stacks into CloudFormation StackSets

A stack import operation can import existing stacks into new or existing StackSets, so that
you can migrate existing stacks to a StackSet in one operation.

For _self-managed_ StackSets, the import operation can import stacks
in the administrator account or in different target accounts and AWS Regions. For
_service-managed_ StackSets, the import operation can import any stack in
the same AWS Organizations as the management account.

The following are considerations and limitations when importing stacks into
StackSets:

- The import operation can import up to 10 stacks using inline stack IDs or up to 200 stacks
using an Amazon S3 object.

- The `NoEcho` property is not supported. Stacks that contain `NoEcho`
won't be imported into new StackSets through StackSet import.

- Stacks can only belong to one StackSet.

- You can implement stack tags to the StackSet by specifying tags explicitly as parameters
in the stack import operation.

- A stack's custom parameter overrides aren't affected during the import operation.

- Parameters of imported stack instances can't be updated using the **Edit StackSet details**
option. You must use the **Override StackSet parameters**. For more information on
overriding parameters, see [Override parameters on\
stacks](stackinstances-override.md).

- The StackSets quotas and stack instances apply when importing stacks. For more
information about quotas, see [Understand CloudFormation quotas](cloudformation-limits.md).

###### Topics

- [Self-managed stack import for CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/self-managed-import.html)

- [Service-managed stack import for CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/service-managed-import.html)

- [Revert stack imports into CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/revert-stackset-import.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Detecting drift on StackSets

Self-managed stack import
