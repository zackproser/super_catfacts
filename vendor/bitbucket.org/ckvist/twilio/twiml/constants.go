package twiml

const (
	TwiMan     = "man"
	TwiWoman   = "woman"
	TwiEnglish = "en"
	TwiBritish = "en-gb"
	TwiSpanish = "es"
	TwiFrench  = "fr"
	TwiGerman  = "de"
	// --------------------------
	TwiAlice              = "alice"
	TwiDanishDenmark      = "da-DK"
	TwiGermanGermany      = "de-DE"
	TwiEnglishAustralia   = "en-AU"
	TwiEnglishCanada      = "en-CA"
	TwiEnglishUK          = "en-UK"
	TwiEnglishIndia       = "en-IN"
	TwiEnglishUSA         = "en-US"
	TwiSpanishCatalan     = "ca-ES"
	TwiSpanishSpain       = "es-ES"
	TwiSpanishMexico      = "es-MX"
	TwiFinishFinland      = "fi-FI"
	TwiFrenchCanada       = "fr-CA"
	TwiFrenchFrance       = "fr-FR"
	TwiItalianItaly       = "it-IT"
	TwiJapaneseJapan      = "ja-JP"
	TwiKoreanKorea        = "ko-KR"
	TwiNorwegianNorway    = "nb-NO"
	TwiDutchNetherlands   = "nl-NL"
	TwiPolishPoland       = "pl-PL"
	TwiPortugueseBrazil   = "pt-BR"
	TwiPortuguesePortugal = "pt-PT"
	TwiRussianRussia      = "ru-RU"
	TwiSwedishSweden      = "sv-SE"
	TwiChineseMandarin    = "zh-CH"
	TwiChineseCantonese   = "zh-HK"
	TwiChineseTaiwanese   = "zh-TW"
)

// Twilio callback status parameters for: Call End Callback (StatusCallback),
// Voice Request
const (
	TwiCallSid       = "CallSid"
	TwiAccountSid    = "AccountSid"
	TwiFrom          = "From"
	TwiTo            = "To"
	TwiCallStatus    = "CallStatus"
	TwiApiVersion    = "ApiVersion"
	TwiDirection     = "Direction"
	TwiForwardedFrom = "ForwardedFrom"
	TwiCallerName    = "CallerName"
	// Geographic data
	TwiFromCity    = "FromCity"
	TwiFromState   = "FromState"
	TwiFromZip     = "FromZip"
	TwiFromCountry = "FromCountry"
	TwiToCity      = "ToCity"
	TwiToState     = "ToState"
	TwiToZip       = "ToZip"
	TwiToCountry   = "Tocountry"
	//  Status callback
	TwiCallDuration      = "CallDuration"
	TwiRecordingUrl      = "RecordingUrl"
	TwiRecordingSid      = "RecordingSid"
	TwiRecordingDuration = "RecordingDuration"
	// Below parameters are included in AddCallerId request response
	TwiVerificationStatus  = "VerificationStatus"
	TwiOutgoingCallerIdSid = "OutgoingCallerIdSid"
)

// Call status
const (
	TwiQueued     = "queued"
	TwiRinging    = "ringing"
	TwiInProgress = "in-progress"
	TwiCompleted  = "completed"
	TwiBusy       = "busy"
	TwiFailed     = "failed"
	TwiNoAnswer   = "no-answer"
	TwiCanceled   = "canceled"
)
