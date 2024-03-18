package logic

import (
	"github.com/gogf/gf/v2/i18n/gi18n"
)

func i18nA() {
	var (
		i18n = gi18n.New()
	)

	i18n.T(ctx, "")
	i18n.Tf(ctx, "gf.gvalid.rule.domain")
	i18n.Translate(ctx, "gf.gvalid.rule.mac")
	i18n.TranslateFormat(ctx, "gf.gvalid.rule.mac")
}
