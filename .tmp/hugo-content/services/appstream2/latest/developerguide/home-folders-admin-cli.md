# Using the AWS Command Line Interface or AWS SDKs

You can enable and disable home folders for a stack by using the AWS CLI or
AWS SDKs.

Use the following [create-stack](../../../cli/latest/reference/appstream/create-stack.md) command to enable home folders while creating a new
stack:

```nohighlight

aws appstream create-stack --name ExampleStack --storage-connectors ConnectorType=HOMEFOLDERS
```

Use the following [update-stack](../../../cli/latest/reference/appstream/update-stack.md) command to enable home folders for an existing
stack:

```nohighlight

aws appstream update-stack --name ExistingStack --storage-connectors ConnectorType=HOMEFOLDERS
```

Use the following command to disable home folders for an existing stack. This
command does not delete any user data.

```nohighlight

aws appstream update-stack --name ExistingStack --delete-storage-connectors
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Home Folder Formats

Additional Resources

All content copied from https://docs.aws.amazon.com/.
