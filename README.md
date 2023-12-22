# bandwidth-go

This is a Go client library for [Bandwidth's Communication APIs](https://dev.bandwidth.com/).

At the time of writing, Bandwidth does not provide an official Go SDK for their APIs. Please check the official [Brandwidth SDKs](https://dev.bandwidth.com/sdks) page for updates and a possible release of an official Go SDK in the future.

For now, this library only contains a package for generating [BXML](https://dev.bandwidth.com/docs/voice/bxml).

## ⚠️ WARNING! ⚠️ 

This library is still in active development. Be aware that not all BXML verbs and nouns have been fully implemented yet. Please see below for a full list of the completed verbs and nouns. If you would like to contribute, please see feel free to submit a pull request.

## Installation

The recommended way to install `bandwidth-go` is by using [Go modules](https://go.dev/ref/mod#go-get).

If you already have an initialized project, you can run the command below from your terminal in the project directory to install the library:

```shell
go get github.com/andersryanc/bandwidth-go
```

If you are starting from scratch in a new directory, you will first need to create a go.mod file for tracking dependencies such as bandwidth-go. This is similar to using package.json in a Node.js project or requirements.txt in a Python project. [You can read more about mod files in the Go documentation](https://golang.org/doc/modules/managing-dependencies). To create the file, run the following command in your terminal:

```shell
go mod init bandwidth-example
```

Once the module is initialized, you may run the installation command from above, which will update your go.mod file to include bandwidth-go.

## Completed Verbs and Nouns

- [ ] [`<Conference>`](https://dev.bandwidth.com/docs/voice/bxml/conference)	The Conference verb is used to add a call to a conference.
- [ ] [`<Bridge>`](https://dev.bandwidth.com/docs/voice/bxml/bridge)	The Bridge verb is used to bridge two calls.
- [x] [`<Pause>`](https://dev.bandwidth.com/docs/voice/bxml/pause)	The Pause verb is used to delay verb execution for a period of time.
- [ ] [`<Forward>`](https://dev.bandwidth.com/docs/voice/bxml/forward)	The Forward verb is used to forward an unanswered incoming call to another number.
- [x] [`<Transfer>`](https://dev.bandwidth.com/docs/voice/bxml/transfer)	The Transfer verb is used to transfer the call to another number.
    - [x] [`<PhoneNumber>`](https://dev.bandwidth.com/docs/voice/bxml/transfer/#nested-tags)	A phone number to transfer the call to. Value must be in E.164 format (e.g. `+15555555555`).
    - [x] [`<SipUri>`](https://dev.bandwidth.com/docs/voice/bxml/transfer/#nested-tags)	A SIP URI to transfer the call to (e.g. `sip:user@server.com`).
- [ ] [`<Ring>`](https://dev.bandwidth.com/docs/voice/bxml/ring)	The Ring verb is used to play ringing audio on a call.
- [x] [`<Hangup>`](https://dev.bandwidth.com/docs/voice/bxml/hangup)	The Hangup verb is used to hang up or reject a call.
- [ ] [`<Redirect>`](https://dev.bandwidth.com/docs/voice/bxml/redirect)	The Redirect verb is used to redirect the current XML execution to another URL.
- [x] [`<PlayAudio>`](https://dev.bandwidth.com/docs/voice/bxml/playAudio/)	The PlayAudio verb is used to play an audio file in the call.
- [x] [`<SpeakSentence>`](https://dev.bandwidth.com/docs/voice/bxml/speakSentence)	The SpeakSentence verb converts text into audible speech.
- [x] [`<Record>`](https://dev.bandwidth.com/docs/voice/bxml/record)	The Record verb allows a segment of audio to be recorded during a call.
- [ ] [`<StartRecording>`](https://dev.bandwidth.com/docs/voice/bxml/startRecording)	The StartRecording verb allows a segment of a call to be recorded while other verbs are executing.
- [ ] [`<PauseRecording>`](https://dev.bandwidth.com/docs/voice/bxml/pauseRecording)	The PauseRecording verb is used to pause a recording previously started by a <StartRecording> verb.
- [ ] [`<ResumeRecording>`](https://dev.bandwidth.com/docs/voice/bxml/resumeRecording)	The ResumeRecording verb is used to resume a recording previously paused by a <PauseRecording> verb.
- [ ] [`<StopRecording>`](https://dev.bandwidth.com/docs/voice/bxml/stopRecording)	The StopRecording verb stops a recording that was previously started by a <StartRecording>.
- [x] [`<Gather>`](https://dev.bandwidth.com/docs/voice/bxml/gather)	The Gather verb is used to collect DTMF digits.
- [ ] [`<StartGather>`](https://dev.bandwidth.com/docs/voice/bxml/startGather)	The StartGather verb is used to collect DTMF digits during the execution of other verbs.
- [ ] [`<StopGather>`](https://dev.bandwidth.com/docs/voice/bxml/stopGather)	The StopGather verb stops the DTMF collection initiated by <StartGather>.
- [x] [`<StartStream>`](https://dev.bandwidth.com/docs/voice/bxml/startStream)	The StartStream verb allows a segment of a call to be streamed to an external destination.
- [x] [`<StopStream>`](https://dev.bandwidth.com/docs/voice/bxml/stopStream)	The StopStream verb is used to stop a stream previously started by a <StartStream> verb.
    - [x] [`<StreamParam>`](https://dev.bandwidth.com/docs/voice/bxml/startStream#nested-tags)	These elements define optional user specified parameters that will be sent to the destination URL when the stream is first started.
- [ ] [`<StartTranscription>`](https://dev.bandwidth.com/docs/voice/bxml/startTranscription)	The StartTranscription verb allows a segment of a call to be transcribed during the execution of other verbs.
- [ ] [`<StopTranscription>`](https://dev.bandwidth.com/docs/voice/bxml/stopTranscription)	The StopTranscription verb is used to stop a transcription previously started by a <StartTranscription> verb.
- [ ] [`<SendDtmf>`](https://dev.bandwidth.com/docs/voice/bxml/sendDtmf)	The SendDtmf verb is used to play DTMF digits in the call.
- [ ] [`<Tag>`](https://dev.bandwidth.com/docs/voice/bxml/tag)	The Tag verb is used to set a new tag value without executing a webhook.