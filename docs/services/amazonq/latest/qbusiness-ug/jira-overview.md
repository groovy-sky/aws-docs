# Jira connector overview

The following table gives an overview of the Amazon Q Business
Jira connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**Basic, Basic, OAuth 2.0 with Refresh Token Flow**Authentication credentials**

For Basic authentication

- Jira URL

- Jira username

- Password (Jira site
token)

For OAuth 2.0 authentication with Refresh
Token Flow

- App key

- App secret

- Access token

- [Refresh token](https://developer.atlassian.com/cloud/jira/platform/oauth-2-3lo-apps)

###### Note

Access and refresh tokens expire in 1 hour. For
information on regenerating tokens, see [Atlassian Developer\
Documentation](https://developer.atlassian.com/cloud/jira/platform/oauth-2-3lo-apps).

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/jira-user-management.html).
**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**No**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Custom objects**Yes**Custom metadata**Yes**Entities**Yes. The following entities are supported:

- Projects

- Issues

- Comments

- Attachments

- Worklogs

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports both default and custom field mappings. For more
information, see [Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/jira-field-mappings).**Filters**Yes. The following filters are supported:

- Include specific projects

- Include/exclude statuses

- Include/exclude comments

- Include/exclude attachments

- Include/exclude worklogs

- Include/exclude bugs

- Include/exclude epic

- Include/exclude story

- Include/exclude task

- Include/exclude by file name

- Include/exclude by file type

- Include/exclude by file path

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Known limitations

Prerequisites
