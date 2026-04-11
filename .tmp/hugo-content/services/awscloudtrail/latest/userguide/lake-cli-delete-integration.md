# Delete a channel to delete an integration with the AWS CLI

This section describes how to run the `delete-channel` command to delete the channel for a CloudTrail Lake integration.
You would delete a channel, if you wanted to stop ingesting partner or other activity events outside of AWS. The ARN or channel
ID (the ARN suffix) of the channel that you want to delete is required.

The following example shows how to delete the channel.

```JSON

aws cloudtrail delete-channel \
--channel EXAMPLE8-0558-4f7e-a06a-43969EXAMPLE
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Update a channel with the AWS CLI

CloudTrail Lake integrations event schema

All content copied from https://docs.aws.amazon.com/.
