package bxml

func Voice(verbs []Element) (string, error) {
	doc, response := CreateDocument()
	if verbs != nil {
		AddAllVerbs(response, verbs)
	}
	return ToXML(doc)
}

// VoicePlayAudio <PlayAudio> BXML Verb
// https://dev.bandwidth.com/docs/voice/bxml/playAudio/
type VoicePlayAudio struct {
	// audio_uri: The URL of the audio file to play. May be a relative URL.
	// username: The username to send in the HTTP request to audioUri.
	// password: The password to send in the HTTP request to audioUri.
	// OptionalAttributes: additional attributes
	AudioUri           string
	Username           string
	Password           string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoicePlayAudio) GetName() string {
	return "PlayAudio"
}

func (m VoicePlayAudio) GetText() string {
	return m.AudioUri
}

func (m VoicePlayAudio) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Username": m.Username,
		"Password": m.Password,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoicePlayAudio) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceStopStream <StopStream> BXML Verb
// https://dev.bandwidth.com/docs/voice/bxml/stopStream
type VoiceStopStream struct {
	// name: Friendly name given to the Stream
	// OptionalAttributes: additional attributes
	Name               string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceStopStream) GetName() string {
	return "StopStream"
}

func (m VoiceStopStream) GetText() string {
	return ""
}

func (m VoiceStopStream) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Name": m.Name,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceStopStream) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceStreamParam <StreamParam> BXML Noun
// https://dev.bandwidth.com/docs/voice/bxml/startStream#nested-tags
type VoiceStreamParam struct {
	// name: The name of the custom parameter
	// value: The value of the custom parameter
	// OptionalAttributes: additional attributes
	Name               string
	Value              string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceStreamParam) GetName() string {
	return "StreamParam"
}

func (m VoiceStreamParam) GetText() string {
	return ""
}

func (m VoiceStreamParam) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Name":  m.Name,
		"Value": m.Value,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceStreamParam) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceStartStream <StartStream> BXML Noun
// https://dev.bandwidth.com/docs/voice/bxml/startStream
type VoiceStartStream struct {
	// name: Friendly name given to the Stream
	// tracks: Track to be streamed to remote service
	// destination: URL of the remote service where the Stream is routed
	// stream_event_url: Stream Event URL
	// stream_event_method: Stream Event URL method
	// username: Stream Event URL username
	// password: Stream Event URL password
	// OptionalAttributes: additional attributes
	Name               string
	Tracks             string
	Destination        string
	StreamEventUrl     string
	StreamEventMethod  string
	Username           string
	Password           string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceStartStream) GetName() string {
	return "StartStream"
}

func (m VoiceStartStream) GetText() string {
	return ""
}

func (m VoiceStartStream) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Name":              m.Name,
		"Tracks":            m.Tracks,
		"Destination":       m.Destination,
		"StreamEventUrl":    m.StreamEventUrl,
		"StreamEventMethod": m.StreamEventMethod,
		"Username":          m.Username,
		"Password":          m.Password,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceStartStream) GetInnerElements() []Element {
	return m.InnerElements
}

// VoicePause <Pause> BXML Verb
// https://dev.bandwidth.com/docs/voice/bxml/pause
type VoicePause struct {
	// duration: Length in seconds to pause
	// OptionalAttributes: additional attributes
	Duration           string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoicePause) GetName() string {
	return "Pause"
}

func (m VoicePause) GetText() string {
	return ""
}

func (m VoicePause) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Duration": m.Duration,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoicePause) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceSpeakSentence <SpeakSentence> BXML Verb
// https://dev.bandwidth.com/docs/voice/bxml/speakSentence
type VoiceSpeakSentence struct {
	// text: The text to speak. Cannot be blank. Can be a mixture of plain text and SSML tags (see below for list of supported tags).
	// voice: Selects the voice of the speaker. If the voice attribute is present, gender and locale are ignored.
	// gender: Selects the gender of the speaker. Valid values are "male" or "female". Default "female". If the chosen gender does not exist for the region, the opposite gender will be used by default.
	// locale: Selects the locale of the speaker. Default "en_US".
	// OptionalAttributes: additional attributes
	Text               string
	Voice              string
	Gender             string
	Locale             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceSpeakSentence) GetName() string {
	return "SpeakSentence"
}

func (m VoiceSpeakSentence) GetText() string {
	return m.Text
}

func (m VoiceSpeakSentence) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Voice":  m.Voice,
		"Gender": m.Gender,
		"Locale": m.Locale,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceSpeakSentence) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceRedirect <Redirect> BXML Verb
type VoiceRedirect struct {
	// url: Redirect URL
	// method: Redirect URL method
	// OptionalAttributes: additional attributes
	Url                string
	Method             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceRedirect) GetName() string {
	return "Redirect"
}

func (m VoiceRedirect) GetText() string {
	return m.Url
}

func (m VoiceRedirect) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Method": m.Method,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceRedirect) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceRecord <Record> BXML Verb
// https://dev.bandwidth.com/docs/voice/bxml/record
type VoiceRecord struct {
	// record_complete_url: URL to send the Record Complete event to once the recording has ended. Accepts BXML, and may be a relative URL. This callback will not be sent if the recording ended due to the call hanging up.
	// record_complete_method: The HTTP method to use for the request to recordCompleteUrl. GET or POST. Default value is POST.
	// ... see docs above
	// OptionalAttributes: additional attributes
	RecordCompleteUrl            string
	RecordCompleteMethod         string
	RecordCompleteFallbackUrl    string
	RecordCompleteFallbackMethod string
	RecordingAvailableUrl        string
	RecordingAvailableMethod     string
	Transcribe                   string
	DetectLanguage               string
	TranscriptionAvailableUrl    string
	TranscriptionAvailableMethod string
	Username                     string
	Password                     string
	FallbackUsername             string
	FallbackPassword             string
	Tag                          string
	TerminatingDigits            string
	MaxDuration                  string
	SilenceTimeout               string
	FileFormat                   string
	InnerElements                []Element
	OptionalAttributes           map[string]string
}

func (m VoiceRecord) GetName() string {
	return "Record"
}

func (m VoiceRecord) GetText() string {
	return ""
}

func (m VoiceRecord) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"RecordCompleteUrl":            m.RecordCompleteUrl,
		"RecordCompleteMethod":         m.RecordCompleteMethod,
		"RecordCompleteFallbackUrl":    m.RecordCompleteFallbackUrl,
		"RecordCompleteFallbackMethod": m.RecordCompleteFallbackMethod,
		"RecordingAvailableUrl":        m.RecordingAvailableUrl,
		"RecordingAvailableMethod":     m.RecordingAvailableMethod,
		"Transcribe":                   m.Transcribe,
		"DetectLanguage":               m.DetectLanguage,
		"TranscriptionAvailableUrl":    m.TranscriptionAvailableUrl,
		"TranscriptionAvailableMethod": m.TranscriptionAvailableMethod,
		"Username":                     m.Username,
		"Password":                     m.Password,
		"FallbackUsername":             m.FallbackUsername,
		"FallbackPassword":             m.FallbackPassword,
		"Tag":                          m.Tag,
		"TerminatingDigits":            m.TerminatingDigits,
		"MaxDuration":                  m.MaxDuration,
		"SilenceTimeout":               m.SilenceTimeout,
		"FileFormat":                   m.FileFormat,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceRecord) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceHangup <Hangup> BXML Verb
// https://dev.bandwidth.com/docs/voice/bxml/hangup/
type VoiceHangup struct {
	// OptionalAttributes: additional attributes
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceHangup) GetName() string {
	return "Hangup"
}

func (m VoiceHangup) GetText() string {
	return ""
}

func (m VoiceHangup) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m VoiceHangup) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceGather <Gather> BXML Verb
type VoiceGather struct {
	// gather_url: Gather URL
	// gather_method: Gather URL method
	// inter_digit_timeout: Time to wait between gathering input(s)
	// first_digit_timeout: Time to wait to gather input
	// terminating_digits: Finish gather on key
	// max_digits: Number of digits to collect
	// OptionalAttributes: additional attributes
	GatherUrl          string
	GatherMethod       string
	InterDigitTimeout  string
	FirstDigitTimeout  string
	TerminatingDigits  string
	MaxDigits          string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceGather) GetName() string {
	return "Gather"
}

func (m VoiceGather) GetText() string {
	return ""
}

func (m VoiceGather) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"GatherUrl":         m.GatherUrl,
		"GatherMethod":      m.GatherMethod,
		"InterDigitTimeout": m.InterDigitTimeout,
		"FirstDigitTimeout": m.FirstDigitTimeout,
		"TerminatingDigits": m.TerminatingDigits,
		"MaxDigits":         m.MaxDigits,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceGather) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceTransfer <Transfer> BXML Verb
// https://dev.bandwidth.com/docs/voice/bxml/transfer/
type VoiceTransfer struct {
	// transfer_caller_id: The caller ID to use when the call is transferred, if different. Must be in E.164 format (e.g. +15555555555) or be one of the following strings Restricted, Anonymous, Private, or Unavailable.
	// transfer_caller_display_name: The caller display name to use when the call is transferred. May not exceed 256 characters nor contain control characters such as new lines.
	// call_timeout: The timeout (in seconds) for the callee to answer the call after it starts ringing.
	// transfer_complete_url: URL to send the Transfer Complete event to and request new BXML. Optional but recommended.
	// transfer_complete_method:  The HTTP method to use for the request to transferCompleteUrl. GET or POST. Default value is POST.
	// ... see docs above
	// OptionalAttributes: additional attributes
	TransferCallerId               string
	TransferCallerDisplayName      string
	CallTimeout                    string
	TransferCompleteUrl            string
	TransferCompleteMethod         string
	TransferCompleteFallbackUrl    string
	TransferCompleteFallbackMethod string
	Username                       string
	Password                       string
	FallbackUsername               string
	FallbackPassword               string
	Tag                            string
	DiversionTreatment             string
	DiversionReason                string
	InnerElements                  []Element
	OptionalAttributes             map[string]string
}

func (m VoiceTransfer) GetName() string {
	return "Transfer"
}

func (m VoiceTransfer) GetText() string {
	return ""
}

func (m VoiceTransfer) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"TransferCallerId":               m.TransferCallerId,
		"TransferCallerDisplayName":      m.TransferCallerDisplayName,
		"CallTimeout":                    m.CallTimeout,
		"TransferCompleteUrl":            m.TransferCompleteUrl,
		"TransferCompleteMethod":         m.TransferCompleteMethod,
		"TransferCompleteFallbackUrl":    m.TransferCompleteFallbackUrl,
		"TransferCompleteFallbackMethod": m.TransferCompleteFallbackMethod,
		"Username":                       m.Username,
		"Password":                       m.Password,
		"FallbackUsername":               m.FallbackUsername,
		"FallbackPassword":               m.FallbackPassword,
		"Tag":                            m.Tag,
		"DiversionTreatment":             m.DiversionTreatment,
		"DiversionReason":                m.DiversionReason,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceTransfer) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceSipUri <SipUri> BXML Noun
// https://dev.bandwidth.com/docs/voice/bxml/transfer/#nested-tags
type VoiceSipUri struct {
	// sip_url: SIP URL
	// uui: The value of the User-To-User header to send within the initial INVITE. Must include the encoding parameter as specified in RFC 7433. Only base64 and jwt encoding are currently allowed. This value, including the encoding specifier, may not exceed 256 characters.
	// transfer_answer_url: URL, if any, to send the Transfer Answer event to and request BXML to be executed for the called party before the call is bridged. May be a relative URL.
	// transfer_answer_method: The HTTP method to use for the request to transferAnswerUrl. GET or POST. Default value is POST.
	// ... See BW Voice Docs: https://dev.bandwidth.com/docs/voice/bxml/transfer#sipuri-attributes
	// OptionalAttributes: additional attributes
	SipUrl                       string
	Uui                          string
	TransferAnswerUrl            string
	TransferAnswerMethod         string
	TransferAnswerFallbackUrl    string
	TransferAnswerFallbackMethod string
	TransferDisconnectUrl        string
	TransferDisconnectMethod     string
	Username                     string
	Password                     string
	FallbackUsername             string
	FallbackPassword             string
	Tag                          string
	InnerElements                []Element
	OptionalAttributes           map[string]string
}

func (m VoiceSipUri) GetName() string {
	return "SipUri"
}

func (m VoiceSipUri) GetText() string {
	return m.SipUrl
}

func (m VoiceSipUri) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Uui":                          m.Uui,
		"TransferAnswerUrl":            m.TransferAnswerUrl,
		"TransferAnswerMethod":         m.TransferAnswerMethod,
		"TransferAnswerFallbackUrl":    m.TransferAnswerFallbackUrl,
		"TransferAnswerFallbackMethod": m.TransferAnswerFallbackMethod,
		"TransferDisconnectUrl":        m.TransferDisconnectUrl,
		"TransferDisconnectMethod":     m.TransferDisconnectMethod,
		"Username":                     m.Username,
		"Password":                     m.Password,
		"FallbackUsername":             m.FallbackUsername,
		"FallbackPassword":             m.FallbackPassword,
		"Tag":                          m.Tag,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceSipUri) GetInnerElements() []Element {
	return m.InnerElements
}

// VoicePhoneNumber <PhoneNumber> BXML Noun
// https://dev.bandwidth.com/docs/voice/bxml/transfer/#nested-tags
type VoicePhoneNumber struct {
	// phone_number: Phone Number to dial
	// transfer_answer_url: URL, if any, to send the Transfer Answer event to and request BXML to be executed for the called party before the call is bridged. May be a relative URL.
	// transfer_answer_method: The HTTP method to use for the request to transferAnswerUrl. GET or POST. Default value is POST.
	// ... See BW Voice Docs: https://dev.bandwidth.com/docs/voice/bxml/transfer#phonenumber-attributes
	// OptionalAttributes: additional attributes
	PhoneNumber                  string
	TransferAnswerUrl            string
	TransferAnswerMethod         string
	TransferAnswerFallbackUrl    string
	TransferAnswerFallbackMethod string
	TransferDisconnectUrl        string
	TransferDisconnectMethod     string
	Username                     string
	Password                     string
	FallbackUsername             string
	FallbackPassword             string
	Tag                          string
	InnerElements                []Element
	OptionalAttributes           map[string]string
}

func (m VoicePhoneNumber) GetName() string {
	return "PhoneNumber"
}

func (m VoicePhoneNumber) GetText() string {
	return m.PhoneNumber
}

func (m VoicePhoneNumber) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"TransferAnswerUrl":            m.TransferAnswerUrl,
		"TransferAnswerMethod":         m.TransferAnswerMethod,
		"TransferAnswerFallbackUrl":    m.TransferAnswerFallbackUrl,
		"TransferAnswerFallbackMethod": m.TransferAnswerFallbackMethod,
		"TransferDisconnectUrl":        m.TransferDisconnectUrl,
		"TransferDisconnectMethod":     m.TransferDisconnectMethod,
		"Username":                     m.Username,
		"Password":                     m.Password,
		"FallbackUsername":             m.FallbackUsername,
		"FallbackPassword":             m.FallbackPassword,
		"Tag":                          m.Tag,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoicePhoneNumber) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceConference <Conference> BXML Noun
type VoiceConference struct {
	// name: Conference name
	// muted: Join the conference muted
	// beep: Play beep when joining
	// start_conference_on_enter: Start the conference on enter
	// end_conference_on_exit: End the conferenceon exit
	// wait_url: Wait URL
	// wait_method: Wait URL method
	// max_participants: Maximum number of participants
	// record: Record the conference
	// region: Conference region
	// coach: Call coach
	// trim: Trim the conference recording
	// status_callback_event: Events to call status callback URL
	// status_callback: Status callback URL
	// status_callback_method: Status callback URL method
	// recording_status_callback: Recording status callback URL
	// recording_status_callback_method: Recording status callback URL method
	// recording_status_callback_event: Recording status callback events
	// event_callback_url: Event callback URL
	// jitter_buffer_size: Size of jitter buffer for participant
	// participant_label: A label for participant
	// OptionalAttributes: additional attributes
	Name                          string
	Muted                         string
	Beep                          string
	StartConferenceOnEnter        string
	EndConferenceOnExit           string
	WaitUrl                       string
	WaitMethod                    string
	MaxParticipants               string
	Record                        string
	Region                        string
	Coach                         string
	Trim                          string
	StatusCallbackEvent           string
	StatusCallback                string
	StatusCallbackMethod          string
	RecordingStatusCallback       string
	RecordingStatusCallbackMethod string
	RecordingStatusCallbackEvent  string
	EventCallbackUrl              string
	JitterBufferSize              string
	ParticipantLabel              string
	InnerElements                 []Element
	OptionalAttributes            map[string]string
}

func (m VoiceConference) GetName() string {
	return "Conference"
}

func (m VoiceConference) GetText() string {
	return m.Name
}

func (m VoiceConference) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Muted":                         m.Muted,
		"Beep":                          m.Beep,
		"StartConferenceOnEnter":        m.StartConferenceOnEnter,
		"EndConferenceOnExit":           m.EndConferenceOnExit,
		"WaitUrl":                       m.WaitUrl,
		"WaitMethod":                    m.WaitMethod,
		"MaxParticipants":               m.MaxParticipants,
		"Record":                        m.Record,
		"Region":                        m.Region,
		"Coach":                         m.Coach,
		"Trim":                          m.Trim,
		"StatusCallbackEvent":           m.StatusCallbackEvent,
		"StatusCallback":                m.StatusCallback,
		"StatusCallbackMethod":          m.StatusCallbackMethod,
		"RecordingStatusCallback":       m.RecordingStatusCallback,
		"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
		"RecordingStatusCallbackEvent":  m.RecordingStatusCallbackEvent,
		"EventCallbackUrl":              m.EventCallbackUrl,
		"JitterBufferSize":              m.JitterBufferSize,
		"ParticipantLabel":              m.ParticipantLabel,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceConference) GetInnerElements() []Element {
	return m.InnerElements
}
