package mailtpl

import (
	"fmt"
	"html"
	"strings"
)

// FamilyInviteEmail builds subject and bodies for a pending family membership invite.
// Vietnamese copy appears first, then English, per product requirement.
func FamilyInviteEmail(
	inviteeDisplayName, familyName, inviterName, profileLink string,
) (subject string, bodyText string, bodyHTML string) {
	invitee := strings.TrimSpace(inviteeDisplayName)
	viInvitee := invitee
	if viInvitee == "" {
		viInvitee = "bạn"
	}
	enInvitee := invitee
	if enInvitee == "" {
		enInvitee = "there"
	}
	fam := strings.TrimSpace(familyName)
	if fam == "" {
		fam = "gia đình"
	}
	inv := strings.TrimSpace(inviterName)
	if inv == "" {
		inv = "một thành viên"
	}
	enInv := strings.TrimSpace(inviterName)
	if enInv == "" {
		enInv = "a member"
	}

	subject = "Lời mời tham gia gia đình / Family invitation"

	bodyText = strings.TrimSpace(fmt.Sprintf(`
Xin chào %s,

%s đã mời bạn tham gia gia đình "%s" trên MMFamily.

Mở liên kết sau để xem chi tiết và chấp nhận hoặc từ chối trên trang thông tin cá nhân:
%s

Nếu bạn không mong đợi lời mời này, bạn có thể bỏ qua email.

---

Hello %s,

%s has invited you to join the family "%s" on MMFamily.

Open the link below on your profile page to review and accept or decline:
%s

If you were not expecting this invitation, you can ignore this message.`,
		viInvitee, inv, fam, profileLink,
		enInvitee, enInv, fam, profileLink,
	))

	escLink := html.EscapeString(profileLink)
	bodyHTML = fmt.Sprintf(
		`<div style="font-family:system-ui,sans-serif;font-size:15px;line-height:1.5;color:#1a1a1a;">
<p>Xin chào <strong>%s</strong>,</p>
<p><strong>%s</strong> đã mời bạn tham gia gia đình <strong>%s</strong> trên MMFamily.</p>
<p>Mở liên kết sau để xem chi tiết và chấp nhận hoặc từ chối trên trang thông tin cá nhân:</p>
<p><a href="%s" style="color:#0066cc;word-break:break-all;">%s</a></p>
<p style="color:#666;font-size:13px;">Nếu bạn không mong đợi lời mời này, bạn có thể bỏ qua email.</p>
<hr style="border:none;border-top:1px solid #e0e0e0;margin:24px 0;" />
<p>Hello <strong>%s</strong>,</p>
<p><strong>%s</strong> has invited you to join the family <strong>%s</strong> on MMFamily.</p>
<p>Open the link below on your profile page to review and accept or decline:</p>
<p><a href="%s" style="color:#0066cc;word-break:break-all;">%s</a></p>
<p style="color:#666;font-size:13px;">If you were not expect this invitation, you can ignore this message.</p>
</div>`,
		html.EscapeString(viInvitee), html.EscapeString(inv), html.EscapeString(fam), escLink, escLink,
		html.EscapeString(enInvitee), html.EscapeString(enInv), html.EscapeString(fam), escLink, escLink,
	)
	// Fix typo "were not expect" -> "were not expecting"
	bodyHTML = strings.Replace(bodyHTML, "were not expect this", "were not expecting this", 1)

	return subject, bodyText, bodyHTML
}
