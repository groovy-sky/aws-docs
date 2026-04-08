# Get a pre-signed S3 URL to download an approved attached file in Amazon Connect Agent Workspace

Returns a pre-signed URL to download an approved attached file while handling an
active contact. The activeContactId is the id of the contact the agent is actively
viewing. This API also returns metadata about the attached file and it will only
return a downloadUrl if the status of the attached file is APPROVED.

**Signature**

```

getAttachedFileUrl({ attachment, activeContactId }: { attachment: Attachment; activeContactId: string; }): Promise<DownloadableAttachment>
```

**DownloadableAttachment Properties**

**Parameter****Type****Description**associatedResourceArnstringAmazon Connect ARN of the resource that the file is attached to.
Could be a Connect Email Contact ARN or a Connect Case ARNfileIdstringThe unique identifier of the attached file resource.downloadUrlstringA pre-signed URL that should be used to download the attached
file.fileArnstringThe unique identifier of the attached file resource
(ARN).fileNamestringA case-sensitive name of the attached file being
uploaded.fileStatusFileStatusThe current status of the attached file. Supported values:
"APPROVED", "REJECTED", "PROCESSING", "FAILED"fileSizeInBytesnumberThe size of the attached file in bytes.creationTimestringThe time of Creation of the file resource as an ISO timestamp.
It's specified in ISO 8601 format: yyyy-MM-ddThh:mm:ss.SSSZ. For
example, 2024-05-03T02:41:28.172Z.

**Attachment Properties**

**Parameter****Type****Description**associatedResourceArnstringAmazon Connect ARN of the resource that the file is attached to.
Could be a Connect Email Contact ARN or a Connect Case ARNfileIdstringThe unique identifier of the attached file resource.

**Usage**

```typescript

const downloadableAttachment = await fileClient.getAttachedFileUrl({
  attachment: {
    associatedResourceArn: sampleAssociatedResourceArn,
    fileId: sampleFileId,
  },
  activeContactId: sampleActiveContactId, // The contact the agent is actively handling
});

const { downloadUrl } = downloadableAttachment;
const response: Response = await fetch(downloadUrl, { method: "GET" });
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

deleteAttachedFile()

startAttachedFileUpload()

All content copied from https://docs.aws.amazon.com/.
