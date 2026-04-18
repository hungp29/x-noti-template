package mailtpl

import (
	"fmt"
	"html"
	"strings"
)

// EmailChangeOTPEmail builds subject and bodies for verifying a new account email.
// Vietnamese copy appears first, then English, per product requirement.
func EmailChangeOTPEmail(displayName, otp string) (subject string, bodyText string, bodyHTML string) {
	name := strings.TrimSpace(displayName)
	viName := name
	if viName == "" {
		viName = "bạn"
	}
	enName := name
	if enName == "" {
		enName = "there"
	}

	subject = "Mã xác minh đổi email / Your email verification code"

	bodyText = strings.TrimSpace(fmt.Sprintf(`
Xin chào %s,

Dùng mã sau để xác nhận địa chỉ email mới của bạn: %s

Nếu bạn không yêu cầu đổi email, vui lòng bỏ qua email này.

---

Hello %s,

Use this code to confirm your new email address: %s

If you did not request this, you can ignore this message.
`, viName, otp, enName, otp))

	escVi := html.EscapeString(viName)
	escEn := html.EscapeString(enName)
	escOTP := html.EscapeString(otp)
	bodyHTML = fmt.Sprintf(
		`<div style="font-family:system-ui,sans-serif;font-size:15px;line-height:1.5;color:#1a1a1a;">
<p>Xin chào <strong>%s</strong>,</p>
<p>Dùng mã sau để xác nhận địa chỉ email mới của bạn:</p>
<p style="font-size:24px;font-weight:bold;letter-spacing:4px;">%s</p>
<p style="color:#666;font-size:13px;">Nếu bạn không yêu cầu đổi email, vui lòng bỏ qua email này.</p>
<hr style="border:none;border-top:1px solid #e0e0e0;margin:24px 0;" />
<p>Hello <strong>%s</strong>,</p>
<p>Use this code to confirm your new email address:</p>
<p style="font-size:24px;font-weight:bold;letter-spacing:4px;">%s</p>
<p style="color:#666;font-size:13px;">If you did not request this, you can ignore this message.</p>
</div>`,
		escVi, escOTP, escEn, escOTP,
	)
	return subject, bodyText, bodyHTML
}
