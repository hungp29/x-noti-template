package mailtpl

import "testing"

func TestRenderPasswordInviteEN(t *testing.T) {
	subj, text, html, err := Render(TemplatePasswordInvite, "en", map[string]string{
		DataRecipientName:      "Pat",
		DataSetPasswordURL:     "https://example.com/set?token=1",
	})
	if err != nil {
		t.Fatal(err)
	}
	if subj == "" || text == "" || html == "" {
		t.Fatalf("empty parts: subj=%q", subj)
	}
	if subj != "Set your password" {
		t.Fatalf("subject: %q", subj)
	}
}

func TestRenderPasswordInviteVI(t *testing.T) {
	_, _, _, err := Render(TemplatePasswordInvite, "vi", map[string]string{
		DataRecipientName:      "",
		DataSetPasswordURL:     "https://example.com/set?token=1",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderUnknownTemplate(t *testing.T) {
	_, _, _, err := Render("nope", "en", nil)
	if err == nil {
		t.Fatal("expected error")
	}
}
