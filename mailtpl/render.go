package mailtpl

import (
	"fmt"
	"strings"

	"github.com/hungp29/x-noti-template/wishn"
)

type passwordInviteTmplData struct {
	Name string
	URL  string
}

type emailChangeOTPTmplData struct {
	Name string
	OTP  string
}

type familyInviteTmplData struct {
	Invitee     string
	Family      string
	Inviter     string
	ProfileLink string
}

// Render selects the template for templateID and locale lang ("vi" or "en"),
// substituting values from data. lang should be normalized with NormalizeLanguage first.
func Render(templateID, lang string, data map[string]string) (subject string, bodyText string, bodyHTML string, err error) {
	if data == nil {
		data = map[string]string{}
	}
	lang = NormalizeLanguage(lang)
	tid := strings.TrimSpace(templateID)
	switch tid {
	case TemplatePasswordInvite:
		url, err := reqKey(data, DataSetPasswordURL)
		if err != nil {
			return "", "", "", err
		}
		name := strings.TrimSpace(data[DataRecipientName])
		if lang == "vi" && name == "" {
			name = "bạn"
		}
		if lang != "vi" && name == "" {
			name = "there"
		}
		return renderMailBundle(tid, lang, passwordInviteTmplData{Name: name, URL: url})

	case TemplateEmailChangeOTP:
		otp, err := reqKey(data, DataOTP)
		if err != nil {
			return "", "", "", err
		}
		name := strings.TrimSpace(data[DataDisplayName])
		if lang == "vi" && name == "" {
			name = "bạn"
		}
		if lang != "vi" && name == "" {
			name = "there"
		}
		return renderMailBundle(tid, lang, emailChangeOTPTmplData{Name: name, OTP: otp})

	case TemplateFamilyInvite:
		invitee := strings.TrimSpace(data[DataInviteeDisplayName])
		fam := strings.TrimSpace(data[DataFamilyName])
		inv := strings.TrimSpace(data[DataInviterName])
		link, err := reqKey(data, DataProfileLink)
		if err != nil {
			return "", "", "", err
		}
		if lang == "vi" {
			if invitee == "" {
				invitee = "bạn"
			}
			if fam == "" {
				fam = "gia đình"
			}
			if inv == "" {
				inv = "một thành viên"
			}
		} else {
			if invitee == "" {
				invitee = "there"
			}
			if fam == "" {
				fam = "your family"
			}
			if inv == "" {
				inv = "a member"
			}
		}
		return renderMailBundle(tid, lang, familyInviteTmplData{
			Invitee:     invitee,
			Family:      fam,
			Inviter:     inv,
			ProfileLink: link,
		})

	case wishn.TemplateWishPriceChange:
		tmplData, err := buildWishPriceMailData(lang, data)
		if err != nil {
			return "", "", "", err
		}
		return renderMailBundle(tid, lang, tmplData)

	default:
		return "", "", "", fmt.Errorf("unknown template_id %q", templateID)
	}
}

func reqKey(data map[string]string, key string) (string, error) {
	v := strings.TrimSpace(data[key])
	if v == "" {
		return "", fmt.Errorf("template data missing key %q", key)
	}
	return v, nil
}
