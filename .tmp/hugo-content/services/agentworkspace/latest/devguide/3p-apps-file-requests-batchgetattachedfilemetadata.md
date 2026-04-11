# Get metadata about multiple attached files in Amazon Connect Agent Workspace

Get metadata about multiple attached files on an associated resource while
handling an active contact. The activeContactId is the id of the contact the agent
is actively viewing. Each attached file provided in the input list must be
associated with the associatedResourceArn in the RelatedAttachments request
object.

**Signature**

```

batchGetAttachedFileMetadata({ relatedAttachments, activeContactId }: { relatedAttachments: RelatedAttachments; activeContactId: string; }): Promise<BatchGetAttachedFileMetadataResponse>
```

**BatchGetAttachedFileMetadataResponse Properties**

**Parameter****Type****Description**filesAttachmentMetadata\[\]Array of file metadata objects that were successfully
retrievederrorsAttachmentError\[\]Array of errors of attached files that could not be
retrieved

**AttachmentMetadata Properties**

**Parameter****Type****Description**associatedResourceArnstringAmazon Connect ARN of the resource that the file is attached to.
Could be a Connect Email Contact ARN or a Connect Case ARNfileIdstringThe unique identifier of the attached file resourcefileArnstringThe unique identifier of the attached file resource
(ARN).fileNamestringA case-sensitive name of the attached file being
uploaded.fileStatusFileStatusThe current status of the attached file. Supported values:
"APPROVED", "REJECTED", "PROCESSING", "FAILED"fileSizeInBytesnumberThe size of the attached file in bytes.creationTimestringThe time of Creation of the file resource as an ISO timestamp.
It's specified in ISO 8601 format: yyyy-MM-ddThh:mm:ss.SSSZ. For
example, 2024-05-03T02:41:28.172Z.

**AttachmentError Properties**

**Parameter****Type****Description**errorCodestringStatus code describing the failureerrorMessagestringWhy the attached file couldn't be retrievedfileIdstringThe unique identifier of the attached file resource

**RelatedAttachments Properties**

**Parameter****Type****Description**associatedResourceArnstringAmazon Connect ARN of the resource that the file is attached to.
Could be a Connect Email Contact ARN or a Connect Case ARNfileIdsstring\[\]The unique identifiers of the attached file resources

**Usage**

```typescript

const relatedAttachments: RelatedAttachments = {
  fileIds: [sampleFileId1, sampleFileId2],
  associatedResourceArn: sampleAssociatedResourceArn,
};

const response: BatchGetAttachedFileMetadataResponse = await fileClient.batchGetAttachedFileMetadata({
  relatedAttachments,
  activeContactId: sampleActiveContactId, // The contact the agent is actively handling
});

const { files, errors } = response;

// Add logic to handle response
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

File

completeAttachedFileUpload()

All content copied from https://docs.aws.amazon.com/.
