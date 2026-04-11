# Amazon Connect Agent Workspace Voice API

The Amazon Connect SDK provides an `VoiceClient` which serves as an interface that your app in the
Amazon Connect agent workspace can use to make data requests on voice contact.

The `VoiceClient` accepts an optional constructor argument, `
        ConnectClientConfig` which itself is defined as:

```typescript

export type ConnectClientConfig = {
    context?: ModuleContext;
    provider?: AmazonConnectProvider;
};

```

If you do not provide a value for this config, then the client will default to using the **AmazonConnectProvider** set in the global provider scope. You can also manually
configure this using **setGlobalProvider**.

You can instantiate the agent client as follows:

```typescript

import { VoiceClient } from "@amazon-connect/voice";

const voiceClient = new VoiceClient({ provider });

```

###### Note

You must first instantiate the [AmazonConnectApp](../../../../reference/agentworkspace/latest/devguide/getting-started-initialize-sdk.md) which initializes the default AmazonConnectProvider and returns `
            { provider } `. This is the recommended option.

Alternatively, providing a constructor argument:

```typescript

import { VoiceClient } from "@amazon-connect/voice";

const voiceClient = new VoiceClient({
    context: sampleContext,
    provider: sampleProvider
});

```

###### Note

Third-party applications must be configured with \* permission in order to utilize the VoiceClient APIs.

The following sections describe API calls for working with the Voice API.

###### Contents

- [canResumeParticipant()](3p-apps-voice-requests-canresumeparticipant.md)

- [canResumeSelf()](3p-apps-voice-requests-canresumeself.md)

- [conferenceParticipants()](3p-apps-voice-requests-conferenceparticipants.md)

- [createOutboundCall()](3p-apps-voice-requests-createoutboundcall.md)

- [getInitialCustomerPhoneNumber()](3p-apps-voice-requests-getinitialcustomerphonenumber.md)

- [getOutboundCallPermission()](3p-apps-voice-requests-getoutboundcallpermission.md)

- [holdParticipant()](3p-apps-voice-requests-holdparticipant.md)

- [getVoiceEnhancementMode()](3p-apps-voice-requests-getvoiceenhancementmode.md)

- [getVoiceEnhancementPaths()](3p-apps-voice-requests-getvoiceenhancementpaths.md)

- [isParticipantOnHold()](3p-apps-voice-requests-isparticipantonhold.md)

- [listDialableCountries()](3p-apps-voice-requests-listdialablecountries.md)

- [offCanResumeParticipantChanged()](3p-apps-voice-requests-offcanresumeparticipantchanged.md)

- [offCanResumeSelfChanged()](3p-apps-voice-requests-offcanresumeselfchanged.md)

- [offParticipantHold()](3p-apps-voice-requests-offparticipanthold.md)

- [offParticipantResume()](3p-apps-voice-requests-offparticipantresume.md)

- [offSelfHold()](3p-apps-voice-requests-offselfhold.md)

- [offSelfResume()](3p-apps-voice-requests-offselfresume.md)

- [offVoiceEnhancementModeChanged()](3p-apps-voice-requests-offvoiceenhancementmodechanged.md)

- [onCanResumeParticipantChanged()](3p-apps-voice-requests-oncanresumeparticipantchanged.md)

- [onCanResumeSelfChanged()](3p-apps-voice-requests-oncanresumeselfchanged.md)

- [onParticipantHold()](3p-apps-voice-requests-onparticipanthold.md)

- [onParticipantResume()](3p-apps-voice-requests-onparticipantresume.md)

- [onSelfHold()](3p-apps-voice-requests-onselfhold.md)

- [onSelfResume()](3p-apps-voice-requests-onselfresume.md)

- [onVoiceEnhancementModeChanged()](3p-apps-voice-requests-onvoiceenhancementmodechanged.md)

- [resumeParticipant()](3p-apps-voice-requests-resumeparticipant.md)

- [setVoiceEnhancementMode()](3p-apps-voice-requests-setvoiceenhancementmode.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

offLanguageChanged()

canResumeParticipant()

All content copied from https://docs.aws.amazon.com/.
