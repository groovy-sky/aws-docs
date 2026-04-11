# Get the application catalog in Amazon Connect Agent Workspace

Returns all the applications that are available in the Amazon Connect agent workspace for
the current logged-in user.

**Signature**

```typescript

getAppCatalog(): Promise<AppConfig[]>
```

**Usage**

```typescript

const applications: AppConfig[] = await appControllerClient.getAppCatalog();
```

**Output - AppConfig**

**Parameter****Type****Description**arnstringThe AmazonResourceName(ARN) of the applicationnamespacestringThe immutable application namespace used at the time of
registrationidstringThe unique identifier of the applicationnamestringName of the applicationdescriptionstringDescription of the applicationaccessUrlstringURL to access the applicationinitializationTimeoutnumberThe maximum time allowed in milliseconds to establish a
connection with the workspacecontactHandling.scopePER\_CONTACT \| CROSS\_CONTACTSIndicates whether the application refreshes for each
contact

**Permissions required:**

```typescript

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getApp()

getAppConfig()

All content copied from https://docs.aws.amazon.com/.
