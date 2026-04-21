package mailtpl

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"path"
	"strings"
	"sync"
	texttpl "text/template"
)

//go:embed templates
var mailTemplatesFS embed.FS

type mailPartSet struct {
	subject *texttpl.Template
	text    *texttpl.Template
	html    *template.Template
}

var (
	mailTplOnce sync.Once
	mailTplByID map[string]*mailPartSet // key: template_id|lang
	mailTplErr  error
)

func loadMailTemplates() {
	mailTplByID = make(map[string]*mailPartSet)
	paths := make(map[string]map[string]string) // bundleKey -> part -> rel path under FS root "templates"

	err := fs.WalkDir(mailTemplatesFS, "templates", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(p, ".gotmpl") {
			return nil
		}
		rel, ok := strings.CutPrefix(p, "templates/")
		if !ok {
			return fmt.Errorf("unexpected template path %q", p)
		}
		segs := strings.Split(rel, "/")
		if len(segs) != 3 {
			return fmt.Errorf("expected templates/<template_id>/<lang>/<part>.gotmpl, got %q", rel)
		}
		tplID := strings.TrimSpace(segs[0])
		lang := NormalizeLanguage(segs[1])
		partFile := segs[2]
		part := strings.TrimSuffix(partFile, ".gotmpl")
		switch part {
		case "subject", "text", "html":
		default:
			return fmt.Errorf("unknown template part %q in %q", part, rel)
		}
		if tplID == "" {
			return fmt.Errorf("empty template_id in %q", rel)
		}
		key := tplID + "|" + lang
		if paths[key] == nil {
			paths[key] = make(map[string]string)
		}
		paths[key][part] = p
		return nil
	})
	if err != nil {
		mailTplErr = err
		return
	}

	for key, parts := range paths {
		for _, need := range []string{"subject", "text", "html"} {
			if parts[need] == "" {
				mailTplErr = fmt.Errorf("mail templates: bundle %q missing %q", key, need)
				return
			}
		}
		subjBytes, err := mailTemplatesFS.ReadFile(parts["subject"])
		if err != nil {
			mailTplErr = err
			return
		}
		textBytes, err := mailTemplatesFS.ReadFile(parts["text"])
		if err != nil {
			mailTplErr = err
			return
		}
		htmlBytes, err := mailTemplatesFS.ReadFile(parts["html"])
		if err != nil {
			mailTplErr = err
			return
		}
		subjT, err := texttpl.New(path.Base(parts["subject"])).Parse(string(subjBytes))
		if err != nil {
			mailTplErr = fmt.Errorf("parse %s: %w", parts["subject"], err)
			return
		}
		textT, err := texttpl.New(path.Base(parts["text"])).Parse(string(textBytes))
		if err != nil {
			mailTplErr = fmt.Errorf("parse %s: %w", parts["text"], err)
			return
		}
		htmlT, err := template.New(path.Base(parts["html"])).Parse(string(htmlBytes))
		if err != nil {
			mailTplErr = fmt.Errorf("parse %s: %w", parts["html"], err)
			return
		}
		mailTplByID[key] = &mailPartSet{subject: subjT, text: textT, html: htmlT}
	}
}

func getMailPartSet(templateID, lang string) (*mailPartSet, error) {
	mailTplOnce.Do(loadMailTemplates)
	if mailTplErr != nil {
		return nil, mailTplErr
	}
	key := strings.TrimSpace(templateID) + "|" + NormalizeLanguage(lang)
	set := mailTplByID[key]
	if set == nil {
		return nil, fmt.Errorf("missing mail templates for template_id=%q lang=%q", templateID, lang)
	}
	return set, nil
}

func execTextTpl(t *texttpl.Template, data any) (string, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}

func execHTMLTpl(t *template.Template, data any) (string, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}

func renderMailBundle(templateID, lang string, data any) (subject string, bodyText string, bodyHTML string, err error) {
	set, err := getMailPartSet(templateID, lang)
	if err != nil {
		return "", "", "", err
	}
	subject, err = execTextTpl(set.subject, data)
	if err != nil {
		return "", "", "", err
	}
	bodyText, err = execTextTpl(set.text, data)
	if err != nil {
		return "", "", "", err
	}
	bodyHTML, err = execHTMLTpl(set.html, data)
	if err != nil {
		return "", "", "", err
	}
	return subject, bodyText, bodyHTML, nil
}
