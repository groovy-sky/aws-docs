# Delete an attached file in Amazon Connect Agent Workspace

Deletes an attached file along with the underlying S3 Object. The attached file is
permanently deleted if S3 bucket versioning is not enabled. The request accepts an
Attachment object, which has the following properties:

- `associatedResourceArn: string`: Amazon Connect ARN of the resource
that the file is attached to. Could be a Connect Email Contact ARN or a Connect
Case ARN

- `fileId: string`: ID in Connect's File record

**Signature**

```

deleteAttachedFile(data: Attachment): Promise<void>
```

**Usage**

```typescript

await fileClient.deleteAttachedFile({
  associatedResourceArn: sampleAssociatedResourceArn, // Get this from the response from `startAttachedFileUpload`
  fileId: sampleFileId // Get this from the response from `startAttachedFileUpload`
});
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

completeAttachedFileUpload()

getAttachedFileUrl()

All content copied from https://docs.aws.amazon.com/.
