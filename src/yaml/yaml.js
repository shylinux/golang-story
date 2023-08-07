Volcanos(chat.ONSYNTAX, {
	yaml: {keyword: {
		"true": code.CONSTANT,
		"false": code.CONSTANT,
	},func: function(can, push, text, indent, opts) {
		var ls = can.core.Split(text, "\t :")
		if (indent == 0) { push(ls[0]), opts.block = ls[0] }
		if (indent == 2) { push(opts.block+"."+ls[0]) }
	}}, yml: {include: ["yaml"]}
})
