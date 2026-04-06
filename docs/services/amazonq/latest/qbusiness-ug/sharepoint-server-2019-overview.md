# SharePoint Server 2019 connector overview

The following table gives an overview of the Amazon Q Business
SharePoint Server 2019 connector and its supported features.

CategoryFeatureSupport**Security****Authentication type**NTLM, Kerberos, SharePoint App-Only (Client Credentials Flow)**Authentication credentials**

**NTLM**

- SharePoint admin username

- SharePoint admin password

If you're using **Email ID with Domain from**
**IDP** to crawl ACLs, then you also need to add
a:

- LDAP Server Endpoint

- LDAP Search Base

- LDAP username

- LDAP password

**Kerberos**

- SharePoint admin username

- SharePoint admin password

If you're using **Email ID with Domain from**
**IDP** to crawl ACLs, then you also need to add
a:

- LDAP Server Endpoint

- LDAP Search Base

- LDAP username

- LDAP password

**SharePoint App-Only (Client Credentials**
**Flow)**

- Tenant ID

- SharePoint App-Only client ID

- SharePoint App-Only client secret

If you're using **Email ID with Domain from**
**IDP** to crawl ACLs, then you also need to add
a:

- LDAP Server Endpoint

- LDAP Search Base

- LDAP username

- LDAP password

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes. For more information, see [ACL crawling](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/sharepoint-server-2019-user-management.html).
**Integration with Identity Provider**
**(IdP)**Yes. LDAP.**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes**[VPC](connector-concepts.md#connector-vpc)**Yes**Crawl features****Entities**Yes. The following entities are supported:

- Files

- Attachments

- Link

- Pages

- Events

- List data

See [What is a document?](connector-doc-crawl.md) for more
details on what each connector crawls as a document.

**[Field mappings](connector-concepts.md#connector-field-mappings)**Yes. Supports both default and custom field mappings. For more
information, see [Field mappings](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/sharepoint-server-2019-field-mappings.html).**Filters**Yes. The following filters are supported:

- Include/exclude by **Links**

- Include/exclude by **Pages**

- Include/exclude by **Events**

- Include/exclude by file name

- Include/exclude by file path

- Include/exclude by file type

- Include/exclude by **OneNote Section**
name

- Include/exclude by **OneNote Page**
name

**[Sync mode](connector-concepts.md#connector-sync-mode)**Supports full and incremental sync.**[File types](doc-types.md)**Supports all files supported by Amazon Q.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Known limitations

Prerequisites
