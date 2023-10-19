Volcanos(chat.ONSYNTAX, {
	md: {func: function(can, push, text, indent) {
		if (indent == 0 && text.indexOf("# ") == 0) { push(text) }
		if (indent == 0 && text.indexOf("## ") == 0) { push(text) }
	}},
})

