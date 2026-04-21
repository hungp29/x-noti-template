package telegramtpl

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"path"
	"strings"
	"sync"
	texttpl "text/template"
)

//go:embed templates
var telegramTemplatesFS embed.FS

var (
	tgTplOnce sync.Once
	tgTplByID map[string]*texttpl.Template // key: template_id|lang
	tgTplErr  error
)

func loadTelegramTemplates() {
	tgTplByID = make(map[string]*texttpl.Template)
	paths := make(map[string]string) // key -> full path under FS

	err := fs.WalkDir(telegramTemplatesFS, "templates", func(p string, d fs.DirEntry, err error) error {
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
			return fmt.Errorf("unexpected telegram template path %q", p)
		}
		segs := strings.Split(rel, "/")
		if len(segs) != 3 {
			return fmt.Errorf("expected templates/<template_id>/<lang>/body.gotmpl, got %q", rel)
		}
		tplID := strings.TrimSpace(segs[0])
		lang := NormalizeLanguage(segs[1])
		partFile := segs[2]
		part := strings.TrimSuffix(partFile, ".gotmpl")
		if part != "body" {
			return fmt.Errorf("telegram template file must be body.gotmpl, got %q", rel)
		}
		if tplID == "" {
			return fmt.Errorf("empty template_id in %q", rel)
		}
		key := tplID + "|" + lang
		if paths[key] != "" {
			return fmt.Errorf("duplicate telegram template for %q", key)
		}
		paths[key] = p
		return nil
	})
	if err != nil {
		tgTplErr = err
		return
	}

	for key, p := range paths {
		b, err := telegramTemplatesFS.ReadFile(p)
		if err != nil {
			tgTplErr = err
			return
		}
		t, err := texttpl.New(path.Base(p)).Parse(string(b))
		if err != nil {
			tgTplErr = fmt.Errorf("parse %s: %w", p, err)
			return
		}
		tgTplByID[key] = t
	}
}

func getTelegramTemplate(templateID, lang string) (*texttpl.Template, error) {
	tgTplOnce.Do(loadTelegramTemplates)
	if tgTplErr != nil {
		return nil, tgTplErr
	}
	key := strings.TrimSpace(templateID) + "|" + NormalizeLanguage(lang)
	t := tgTplByID[key]
	if t == nil {
		return nil, fmt.Errorf("missing telegram template for template_id=%q lang=%q", templateID, lang)
	}
	return t, nil
}

func execTelegramTpl(t *texttpl.Template, data any) (string, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}
