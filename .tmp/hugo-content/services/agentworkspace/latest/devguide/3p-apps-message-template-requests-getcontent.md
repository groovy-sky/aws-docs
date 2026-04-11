# Get content of a message template in Amazon Connect Agent Workspace

Gets the content of a message template. This includes plaintext or html content of the
body of the message template as a string, the subject of the message template, and
attachments if they are configured on the message template. Attributes in the message
template will be filled if they are system attributes, agent attributes, or custom
attributes set up in the contact flow. More information on the attributes can be found
here: [https://docs.aws.amazon.com/connect/latest/adminguide/personalize-templates.html](../../../connect/latest/adminguide/personalize-templates.md)

**Signature**

```

getContent(messageTemplateId: string, contactId: string): Promise<MessageTemplateContent>
```

**messageTemplateId**

The messageTemplateId can be either the ID or the ARN of a message template

- Passing in the ARN returned by searchMessageTemplates is recommended here,
since this will get the content of the active version of the message
template.

- Passing in the ID will return the content of the latest version of the
message template. A qualifier can be appended to the messageTemplateId to
get the content of a different version of the message template.

More information on qualifiers can be found here:

[https://docs.aws.amazon.com/connect/latest/APIReference/API\_amazon-q-connect\_GetMessageTemplate.html](../../../../reference/connect/latest/apireference/api-amazon-q-connect-getmessagetemplate.md)

More information on versioning can be found here:

[https://docs.aws.amazon.com/connect/latest/adminguide/about-version-message-templates.html](../../../connect/latest/adminguide/about-version-message-templates.md)

**contactId**

The current contact to add the message template to. This is used to populate
attributes in the message template

**MessageTemplateContent Properties**

**Parameter****Type****Description**subjectstringMessage subject populated in the templatebodyMessageTemplateBodyMessage body content populated in the template. This can include
plainText or html or bothattachmentsMessageTemplateAttachment\[\]Attachments populated in the templateattributesNotInterpolatedstring\[\]List of attributes that were not automatically populated in the
message template. If all attributes were automatically populated,
this list will be empty

**MessageTemplateBody Properties**

**Parameter****Type****Description**plainTextstringPlain text content of the message template as a string. It is
possible for both the plain text and html to be populated, or for
only the plain text or html content to be populatedhtmlstringHTML content of the message template as a string

**MessageTemplateAttachment Properties**

**Parameter****Type****Description**fileNamestringName of the attachmentfileIdstringID of the attachmentdownloadUrlstringURL to download the attachment from

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MessageTemplate

isEnabled()

All content copied from https://docs.aws.amazon.com/.
