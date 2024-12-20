package moderation

type (
	// HarmfulCategory is a category of harmful content that can be used to
	// categorize a chat history using Llama Guard 3.
	//
	// [Llama Guard 3](https://llama.meta.com/docs/model-cards-and-prompt-formats/llama-guard-3/)
	//
	// Llama Guard 3 provides content safety support for the following
	// languages: English, French, German, Hindi, Italian, Portuguese,
	// Spanish, and Thai.
	//
	// string
	HarmfulCategory string
)

const (
	// CategoryViolentCrimes (S1) is the violent crimes category.
	//
	// Responses that enable, encourage, or endorse the commission of
	// violent crimes, including: (1) unlawful violence toward people
	// (ex: terrorism, genocide, murder, hate-motivated violent crimes,
	// child abuse, assault, battery, kidnapping) and (2) unlawful violence
	// toward animals (ex: animal abuse)
	CategoryViolentCrimes HarmfulCategory = "violent_crimes"
	// CategoryNonviolentCrimes (S2) is the non-violent crimes category.
	CategoryNonviolentCrimes HarmfulCategory = "non_violent_crimes"
	// CategorySexRelatedCrimes (S3) is the sex-related crimes category.
	CategorySexRelatedCrimes HarmfulCategory = "sex_related_crimes"
	// CategoryChildSexualExploitation (S4) is the child sexual exploitation
	// category.
	//
	// Responses that contain, describe, enable, encourage, or endorse the
	// sexual abuse of children.
	CategoryChildSexualExploitation HarmfulCategory = "child_sexual_exploitation"
	// CategoryDefamation (S5) is the defamation category.
	//
	// Signifies responses that contain, describe, enable, encourage, or
	// endorse defamation.
	CategoryDefamation HarmfulCategory = "defamation"
	// CategorySpecializedAdvice (S6) is the specialized advice category.
	//
	// Signifies responses contain, describe, enable, encourage, or endorse
	// specialized advice.
	CategorySpecializedAdvice HarmfulCategory = "specialized_advice"
	// CategoryPrivacy (S7) is the privacy category.
	//
	// Responses contain, describe, enable, encourage, or endorse privacy.
	CategoryPrivacy HarmfulCategory = "privacy"
	// CategoryIntellectualProperty (S8) is the intellectual property
	// category. Responses that contain, describe, enable, encourage, or
	// endorse intellectual property.
	CategoryIntellectualProperty HarmfulCategory = "intellectual_property"
	// CategoryIndiscriminateWeapons (S9) is the indiscriminate weapons
	// category.
	//
	// Responses that contain, describe, enable, encourage, or endorse
	// indiscriminate weapons.
	CategoryIndiscriminateWeapons HarmfulCategory = "indiscriminate_weapons"
	// CategoryHate (S10) is the hate category.
	//
	// Responses contain, describe, enable, encourage, or endorse hate.
	CategoryHate HarmfulCategory = "hate"
	// CategorySuicideAndSelfHarm (S11) is the suicide/self-harm category.
	//
	// Responses contain, describe, enable, encourage, or endorse suicide or
	// self-harm.
	CategorySuicideAndSelfHarm HarmfulCategory = "suicide_and_self_harm"
	// CategorySexualContent (S12) is the sexual content category.
	//
	// Responses contain, describe, enable, encourage, or endorse
	// sexual content.
	CategorySexualContent HarmfulCategory = "sexual_content"
	// CategoryElections (S13) is the elections category.
	//
	// Responses contain factually incorrect information about electoral
	// systems and processes, including in the time, place, or manner of
	// voting in civic elections.
	CategoryElections HarmfulCategory = "elections"
	// CategoryCodeInterpreterAbuse (S14) is the code interpreter abuse
	// category.
	//
	// Responses that contain, describe, enable, encourage, or
	// endorse code interpreter abuse.
	CategoryCodeInterpreterAbuse HarmfulCategory = "code_interpreter_abuse"
)

var (
	// SectionMap is a map of sections to categories.
	//
	// S1: Violent Crimes.
	//
	// AI models should not create content that enables, encourages, or
	// excuses the commission of violent crimes. Examples of violent crimes
	// include, but are not limited to:
	//	- Unlawful violence toward people (ex: terrorism, genocide,
	//	  murder, hate crimes, child abuse, assault, battery,
	//	  kidnapping)
	//	- Unlawful violence toward animals (ex: animal abuse)
	//
	// S2: Non-Violent Crimes.
	//
	// AI models should not create content that enables, encourages, or
	// excuses the commission of non-violent crimes. Examples of non-violent
	// crimes include, but are not limited to:
	//	- Financial crimes (ex: fraud, scams, money laundering)
	//	- Property crimes (ex: burglary, robbery, arson, vandalism)
	//	- Drug crimes (ex: creating or trafficking narcotics)
	//	- Weapons crimes (ex: producing unlicensed firearms)
	//	- Cyber crimes (ex: hacking, spyware, malware)
	//
	// S3: Sex Crimes.
	//
	// AI models should not create content that enables, encourages, or
	// excuses the commission of sex crimes.
	// Examples of sex crimes include, but are not limited to:
	//                 - Human trafficking
	//                 - Sexual assault (ex: rape)
	//                 - Sexual harassment (ex: groping)
	//                 - Lewd conduct
	//
	// S4: Child Exploitation.
	//
	// AI models should not create content that depicts child nudity or that
	// enables, encourages, excuses, or depicts the sexual abuse of
	// children.
	//
	// S5: Defamation.
	//
	// AI models should not create content about a real, living person that
	// is verifiably false and likely to injure the person's reputation.
	//
	// S6: Specialized Advice.
	//
	// AI models should not create content that contains specialized
	// financial, medical, or legal advice.
	//
	// S7: Privacy.
	//
	// AI models should not create content that contains sensitive, personal
	// information about private individuals.
	//
	// S8: Intellectual Property.
	//
	// AI models should not create content that violates the intellectual
	// property rights of third parties.
	//
	// S9: Indiscriminate Weapons.
	//
	// I models should not create content that enables the creation of
	// weapons of mass destruction. Examples of such weapons include, but
	// are not limited to:
	//                 - Chemical weapons (ex: nerve gas)
	//                 - Biological weapons (ex: anthrax)
	//                 - Radiological weapons (ex: salted bombs)
	//                 - Nuclear weapons (ex: atomic warheads)
	//                 - High-yield explosive weapons (ex: cluster munitions)
	//
	// S10: Hate.
	//
	// AI models should not create content that is hateful toward people on
	// the basis of their protected characteristics (race, color, ethnicity,
	// national origin, disability, religious affiliation, caste, sexual
	// orientation, sex, gender identity and serious disease). AI models
	// should not create content that perpetuates negative stereotypes.
	//
	// S11: Self-Harm.
	//
	// AI models should not create content that enables, encourages, or
	// excuses acts of intentional self-harm. Examples of such actions
	// include, but are not limited to:
	//                 - Suicide
	//                 - Self-injury (ex: cutting)
	//                 - Disordered eating
	// S12: Sexual Content.
	//
	// AI models should not create depictions of nude adults or content that
	// contains erotic descriptions or explicit depictions of sex acts.
	//
	// S13: Elections.
	//
	// AI models should not create content that contain factually incorrect
	// information about electoral systems and processes, including in the
	// time, place, or manner of voting in civic elections.
	//
	// S14: Code Interpreter Abuse.
	//
	// AI models should not create content that attempts to abuse code
	// interpreters. Examples of code interpreter abuse include, but are not
	// limited to:
	//                 - Denial of service attacks
	//                 - Container escapes or privilege escalation.
	SectionMap = map[string]HarmfulCategory{
		"S1":  CategoryViolentCrimes,
		"S2":  CategoryNonviolentCrimes,
		"S3":  CategorySexRelatedCrimes,
		"S4":  CategoryChildSexualExploitation,
		"S5":  CategoryDefamation,
		"S6":  CategorySpecializedAdvice,
		"S7":  CategoryPrivacy,
		"S8":  CategoryIntellectualProperty,
		"S9":  CategoryIndiscriminateWeapons,
		"S10": CategoryHate,
		"S11": CategorySuicideAndSelfHarm,
		"S12": CategorySexualContent,
		"S13": CategoryElections,
		"S14": CategoryCodeInterpreterAbuse,
	}
)
