# Gmail connector overview

The following table gives an overview of the Gmail connector and its
supported features.

CategoryFeatureLatest ConnectorLegacy Connector**Security****Authentication type**Google Service AccountGoogle Service Account**Authentication credentials**

**Google service account**

- Google service account

- Admin account email

- Client email

- Private key

**Google service account**

- Google service account

- Admin account email

- Client email

- Private key

**[Access Control List (ACL)](connector-concepts.md#connector-authorization)**
**crawling**Yes (Automatic)Yes (Manual configuration)**[Identity\**
**crawling](connector-concepts.md#connector-identity-crawler)**Yes (Automatic)Yes (Manual configuration)**[VPC](connector-concepts.md#connector-vpc)**NoYes**Customer Managed Key (CMK) support**NoYes**Crawl features****Custom metadata**NoNo**Entities**Messages (automatic), Draft emails (configurable)Messages, Attachments (configurable)**[Field mappings](connector-concepts.md#connector-field-mappings)**No (Uses optimized defaults)Yes (Configurable)**Filters**NoDate range, Attachments (regex), Domains, Keywords, Labels**[Sync mode](connector-concepts.md#connector-sync-mode)**Sync modeNo (Optimized automatic sync)Yes (Full and incremental sync)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Connector versions

Prerequisites

All content copied from https://docs.aws.amazon.com/.
