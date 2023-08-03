Volcanos(chat.ONSYNTAX, {
	yaml: {func: function(can, push, text, indent) {
		if (indent == 0) { push(text) }
	}}, yml: {include: ["yaml"]}
})
