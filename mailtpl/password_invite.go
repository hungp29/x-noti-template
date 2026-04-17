package mailtpl

import (
	"fmt"
	"html"
	"strings"
)

// PasswordInviteEmail returns subject, plain text, and HTML for inviting a user to set their password.
// Vietnamese copy appears first, then English, per product requirement.
func PasswordInviteEmail(recipientName, setPasswordURL string) (subject string, bodyText string, bodyHTML string) {
	name := strings.TrimSpace(recipientName)
	if name == "" {
		name = "bạn"
	}
	subject = "Đặt mật khẩu tài khoản / Set your password"

	bodyText = strings.TrimSpace(fmt.Sprintf(`
Xin chào %s,

Bạn đã được mời tham gia. Vui lòng nhấp vào liên kết sau để đặt mật khẩu cho tài khoản của bạn:
%s

Liên kết có thời hạn. Nếu bạn không yêu cầu email này, hãy bỏ qua.

---

Hello %s,

You have been invited to join. Please use the link below to set your password:
%s

This link will expire. If you did not expect this message, you can ignore it.
`, name, setPasswordURL, name, setPasswordURL))

	escName := html.EscapeString(name)
	escURL := html.EscapeString(setPasswordURL)
	bodyHTML = fmt.Sprintf(
		`<div style="font-family:system-ui,sans-serif;font-size:15px;line-height:1.5;color:#1a1a1a;">
<p>Xin chào <strong>%s</strong>,</p>
<p>Bạn đã được mời tham gia. Vui lòng nhấp vào liên kết sau để đặt mật khẩu cho tài khoản của bạn:</p>
<p><a href="%s">%s</a></p>
<p style="color:#666;font-size:13px;">Liên kết có thời hạn. Nếu bạn không yêu cầu email này, hãy bỏ qua.</p>
<hr style="border:none;border-top:1px solid #e0e0e0;margin:24px 0;" />
<p>Hello <strong>%s</strong>,</p>
<p>You have been invited to join. Please use the link below to set your password:</p>
<p><a href="%s">%s</a></p>
<p style="color:#666;font-size:13px;">This link will expire. If you did not expect this message, you can ignore it.</p>
</div>`,
		escName, escURL, escURL, escName, escURL, escURL,
	)
	return subject, bodyText, bodyHTML
}
