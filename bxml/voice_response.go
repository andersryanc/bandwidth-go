package bxml

func Voice(verbs []Element) (string, error) {
	doc, response := CreateDocument()
	if verbs != nil {
		AddAllVerbs(response, verbs)
	}
	return ToXML(doc)
}

// VoiceRefer <Refer> BXML Verb
type VoiceRefer struct {
	// action: Action URL
	// method: Action URL method
	// OptionalAttributes: additional attributes
	Action             string
	Method             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceRefer) GetName() string {
	return "Refer"
}

func (m VoiceRefer) GetText() string {
	return ""
}

func (m VoiceRefer) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action": m.Action,
		"Method": m.Method,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceRefer) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceReferSip <Sip> BXML Noun used in <Refer>
type VoiceReferSip struct {
	// sip_url: SIP URL
	// OptionalAttributes: additional attributes
	SipUrl             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceReferSip) GetName() string {
	return "Sip"
}

func (m VoiceReferSip) GetText() string {
	return m.SipUrl
}

func (m VoiceReferSip) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m VoiceReferSip) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceStopStream <StopStream> BXML Verb
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
type VoiceStartStream struct {
	// name: Friendly name given to the Stream
	// destination: URL of the remote service where the Stream is routed
	// tracks: Track to be streamed to remote service
	// stream_event_url: Stream Event URL
	// stream_event_method: Stream Event URL method
	// username: Stream Event URL username
	// password: Stream Event URL password
	// OptionalAttributes: additional attributes
	Name               string
	Destination        string
	Tracks             string
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
		"Destination":       m.Destination,
		"Tracks":            m.Tracks,
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

// VoicePlay <Play> BXML Verb
type VoicePlay struct {
	// url: Media URL
	// loop: Times to loop media
	// digits: Play DTMF tones for digits
	// OptionalAttributes: additional attributes
	Url                string
	Loop               string
	Digits             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoicePlay) GetName() string {
	return "Play"
}

func (m VoicePlay) GetText() string {
	return m.Url
}

func (m VoicePlay) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Loop":   m.Loop,
		"Digits": m.Digits,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoicePlay) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceSpeakSentence <SpeakSentence> BXML Verb
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

// VoiceReject <Reject> BXML Verb
type VoiceReject struct {
	// reason: Rejection reason
	// OptionalAttributes: additional attributes
	Reason             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceReject) GetName() string {
	return "Reject"
}

func (m VoiceReject) GetText() string {
	return ""
}

func (m VoiceReject) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Reason": m.Reason,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceReject) GetInnerElements() []Element {
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
type VoiceRecord struct {
	// action: Action URL
	// method: Action URL method
	// finish_on_key: Finish recording on key
	// max_length: Max time to record in seconds
	// play_beep: Play beep
	// trim: Trim the recording
	// recording_status_callback: Status callback URL
	// recording_status_callback_method: Status callback URL method
	// OptionalAttributes: additional attributes
	Action                        string
	Method                        string
	FinishOnKey                   string
	MaxLength                     string
	PlayBeep                      string
	Trim                          string
	RecordingStatusCallback       string
	RecordingStatusCallbackMethod string
	InnerElements                 []Element
	OptionalAttributes            map[string]string
}

func (m VoiceRecord) GetName() string {
	return "Record"
}

func (m VoiceRecord) GetText() string {
	return ""
}

func (m VoiceRecord) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action":                        m.Action,
		"Method":                        m.Method,
		"FinishOnKey":                   m.FinishOnKey,
		"MaxLength":                     m.MaxLength,
		"PlayBeep":                      m.PlayBeep,
		"Trim":                          m.Trim,
		"RecordingStatusCallback":       m.RecordingStatusCallback,
		"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceRecord) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceQueue <Queue> BXML Noun
type VoiceQueue struct {
	// name: Queue name
	// url: Action URL
	// method: Action URL method
	// reservation_sid: TaskRouter Reservation SID
	// post_work_activity_sid: TaskRouter Activity SID
	// OptionalAttributes: additional attributes
	Name                string
	Url                 string
	Method              string
	ReservationSid      string
	PostWorkActivitySid string
	InnerElements       []Element
	OptionalAttributes  map[string]string
}

func (m VoiceQueue) GetName() string {
	return "Queue"
}

func (m VoiceQueue) GetText() string {
	return m.Name
}

func (m VoiceQueue) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Url":                 m.Url,
		"Method":              m.Method,
		"ReservationSid":      m.ReservationSid,
		"PostWorkActivitySid": m.PostWorkActivitySid,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceQueue) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceLeave <Leave> BXML Verb
type VoiceLeave struct {
	// OptionalAttributes: additional attributes
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceLeave) GetName() string {
	return "Leave"
}

func (m VoiceLeave) GetText() string {
	return ""
}

func (m VoiceLeave) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m VoiceLeave) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceHangup <Hangup> BXML Verb
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

// VoiceEnqueue <Enqueue> BXML Noun
type VoiceEnqueue struct {
	// name: Friendly name
	// action: Action URL
	// max_queue_size: Maximum size of queue
	// method: Action URL method
	// wait_url: Wait URL
	// wait_url_method: Wait URL method
	// workflow_sid: TaskRouter Workflow SID
	// OptionalAttributes: additional attributes
	Name               string
	Action             string
	MaxQueueSize       string
	Method             string
	WaitUrl            string
	WaitUrlMethod      string
	WorkflowSid        string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceEnqueue) GetName() string {
	return "Enqueue"
}

func (m VoiceEnqueue) GetText() string {
	return m.Name
}

func (m VoiceEnqueue) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action":        m.Action,
		"MaxQueueSize":  m.MaxQueueSize,
		"Method":        m.Method,
		"WaitUrl":       m.WaitUrl,
		"WaitUrlMethod": m.WaitUrlMethod,
		"WorkflowSid":   m.WorkflowSid,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceEnqueue) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceTransfer <Transfer> BXML Verb
type VoiceTransfer struct {
	// transfer_caller_id: The caller ID to use when the call is transferred, if different. Must be in E.164 format (e.g. +15555555555) or be one of the following strings Restricted, Anonymous, Private, or Unavailable.
	// transfer_caller_display_name: The caller display name to use when the call is transferred. May not exceed 256 characters nor contain control characters such as new lines.
	// call_timeout: The timeout (in seconds) for the callee to answer the call after it starts ringing.
	// transfer_complete_url: URL to send the Transfer Complete event to and request new BXML. Optional but recommended.
	// transfer_complete_method:  The HTTP method to use for the request to transferCompleteUrl. GET or POST. Default value is POST.
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
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceTransfer) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceSipUri <SipUri> BXML Noun
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
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceSipUri) GetInnerElements() []Element {
	return m.InnerElements
}

// VoicePhoneNumber <PhoneNumber> BXML Noun
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

// VoiceConnect <Connect> BXML Verb
type VoiceConnect struct {
	// action: Action URL
	// method: Action URL method
	// OptionalAttributes: additional attributes
	Action             string
	Method             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceConnect) GetName() string {
	return "Connect"
}

func (m VoiceConnect) GetText() string {
	return ""
}

func (m VoiceConnect) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action": m.Action,
		"Method": m.Method,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceConnect) GetInnerElements() []Element {
	return m.InnerElements
}
