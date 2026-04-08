# Get the metadata for an email contact in Amazon Connect Agent Workspace

Returns the metadata for an email contact id while handling an active contact. The
activeContactId is the id of the email contact the agent is actively viewing while
contactId is the id of the email contact whose metadata should be retrieved.

**Signature**

```

getEmailData({ contactId, activeContactId }: { contactId: string; activeContactId: string; }): Promise<EmailContact>
```

**Output - _EmailContact_**
**_Properties_**

**Parameter****Type****Description**contactIdstringThe id of the email contactcontactArnstringThe ARN of the email contactcontactAssociationIdstringThe root contactId which is used as a unique identifier for all
subsequent contacts in a contact tree. Use this value with the
EmailClient.getEmailThread api.relatedContactIdstringThis contact is in response or related to the specified related
contact.initialContactIdstringIf this contact is related to other contacts, this is the id of
the initial contact.subjectstringThe subject of the emailfromEmailAddressAn object that includes the from email address; this value could
be undefined when the email contact has not been sent.toEmailAddress\[\]An array of objects, each including an email address the email
contact was sent toccEmailAddress\[\]An array of objects, each including an email address that was
carbon copied on the email contactdeliveredToEmailAddressAn object that includes the email address associated with Amazon
Connect that received this message; this is only applicable when
direction=INBOUND.direction"INBOUND" \| "OUTBOUND"INBOUND means the email contact was delivered to Amazon Connect;
OUTBOUND means the email contact is from Amazon ConnectbodyLocationEmailArtifactLocationAn object that includes the id and associated resource ARN of the
file that the email contact's body is stored in; this value could be
undefined when the email contact has not been sent.attachmentLocationsEmailArtifactLocation\[\]An array of objects, each including the id and associated
resource ARN, of files that have been attached to the email
contact

**EmailAddress Properties**

**Parameter****Type****Description**emailAddressstringThe email addressdisplayNamestringThe name that is displayed inside the recipient's mailbox

**EmailArtifactLocation Properties**

**Parameter****Type****Description**fileIdstringThe id of the attached fileassociatedResourceArnstringThe Amazon Connect resource to which the attached file is related
to

**Usage**

```typescript

const activeContactId: string = "exampleActiveContactId"; // The contact the agent is actively handling
const contactId: string = "contactIdToDescribe"; // The email contact id whose metadata should be retrieved

const emailMetadata: EmailContact = await emailClient.getEmailData({ contactId, activeContactId });

// Get the body of the email through the File Client
const bodyLocation = emailMetadata.bodyLocation;
if (bodyLocation) {
   const body: DownloadableAttachment = await fileClient.getAttachedFileUrl({
        attachment: bodyLocation,
        activeContactId,
   });

   const { downloadUrl } = body;
   const response: Response = await fetch(downloadUrl, { method: "GET" });
   const bodyContent: string = (await response.json()).messageContent;
}
```

###### Note

The EmailContact object will contain bodyLocation and attachmentLocations,
both of which will require use of the FileClient's getAttachedFileUrl to get the
relevant data for those objects.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offDraftEmailCreated()

getEmailThread()

All content copied from https://docs.aws.amazon.com/.
